package auth_resource

import (
	"net/http"

	"time"

	"fmt"

	"github.com/MongoDBNavigator/go-backend/resource/auth-resource/representation"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
)

func (rcv *authResource) login(request *restful.Request, response *restful.Response) {
	postRequest := new(representation.PostCredentials)

	if err := request.ReadEntity(&postRequest); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	fmt.Println(postRequest, rcv.username, rcv.password)

	if postRequest.Username != rcv.username || postRequest.Password != rcv.password {
		response.WriteHeaderAndEntity(http.StatusForbidden, representation.Error{Message: "Invalid credentials."})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = rcv.username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(rcv.exp)).Unix()

	tokenString, err := token.SignedString([]byte(rcv.password))

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	result := new(representation.JwtToken)
	result.Token = tokenString

	response.WriteEntity(result)
}
