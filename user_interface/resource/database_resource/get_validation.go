package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/transformer/request"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/transformer/response"
	"github.com/emicklei/go-restful"
)

// Method to get document validation
func (rcv *databaseResource) getValidation(req *restful.Request, res *restful.Response) {
	var dbName value.DBName
	var collName value.CollName

	if err := request.ExtractParametersFromRequest(req, &dbName, &collName, nil, nil); err != nil {
		res.WriteHeaderAndEntity(http.StatusNotFound, representation.Error{Message: err.Error()})
		return
	}

	validation, err := rcv.validationReader.Read(dbName, collName)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	res.WriteEntity(response.ValidationToView(validation))
}
