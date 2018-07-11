package transformer

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"fmt"

	"reflect"

	"github.com/emicklei/go-restful"
	"gopkg.in/mgo.v2/bson"
)

func RequestToFilter(request *restful.Request) (bson.M, error) {
	params, err := url.ParseQuery(request.Request.URL.RawQuery)

	if err != nil {
		return nil, err
	}

	filters := make(bson.M)

	if rawFilters, ok := params["filter[]"]; ok {
		for _, filterString := range rawFilters {
			if len(strings.TrimSpace(filterString)) == 0 {
				return nil, errors.New("filter[] value should not be blank.")
			}

			filter, err := parseFilter(filterString)

			if err != nil {
				return nil, err
			}

			for k, v := range filter {
				filters[k] = v
			}
		}
	}

	return filters, nil
}

func parseFilter(data string) (bson.M, error) {
	if !strings.Contains(data, ":") {
		return nil, errors.New("filter[] has bad formant (field:value).")
	}

	dataSlice := strings.SplitN(data, ":", 2)

	if len(strings.TrimSpace(dataSlice[0])) == 0 {
		return nil, errors.New("Filter kay should not be blank.")
	}

	if len(strings.TrimSpace(dataSlice[1])) == 0 {
		return nil, errors.New("Filter value should not be blank.")
	}

	jsonFormat := `{"%s":%s}`

	if reflect.TypeOf(dataSlice[1]).String() == "string" {
		jsonFormat = `{"%s":"%s"}`
	}

	jsonData := fmt.Sprintf(jsonFormat, dataSlice[0], dataSlice[1])

	var filter map[string]interface{}

	if err := json.Unmarshal([]byte(jsonData), &filter); err != nil {
		return nil, err
	}

	return convertToBson(filter), nil
}

func convertToBson(data map[string]interface{}) bson.M {
	filter := make(bson.M)

	if regex, ok := data["$regex"]; ok {
		var options string
		if opts, ok := data["$options"]; ok {
			options = opts.(string)
		}

		filter["$regex"] = bson.RegEx{
			Pattern: regex.(string),
			Options: options,
		}

		return filter
	}

	for k, v := range data {
		vType := reflect.TypeOf(v)

		if vType.String() == "map[string]interface {}" {
			filter[k] = convertToBson(v.(map[string]interface{}))
		} else if vType.String() == "[]interface {}" {
			vSlice := v.([]interface{})
			filter[k] = make([]interface{}, len(vSlice))
			for i, j := range vSlice {
				filter[k].([]interface{})[i] = convertToBson(j.(map[string]interface{}))
			}
		} else {
			filter[k] = v
		}
	}

	return filter
}
