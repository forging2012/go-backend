package transformer

import (
	"errors"
	"strings"

	"github.com/emicklei/go-restful"
)

func RequestToDocumentId(request *restful.Request) (string, error) {
	recordId := request.PathParameter("documentId")

	if len(strings.TrimSpace(recordId)) == 0 {
		return "", errors.New("Document ID should not be blank.")
	}

	return recordId, nil
}