package database_resource

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/MongoDBNavigator/go-backend/tests/mock"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
	"github.com/emicklei/go-restful"
	"github.com/golang/mock/gomock"
)

var (
	databaseReader    *mock.MockDatabaseReader
	databaseWriter    *mock.MockDatabaseWriter
	collectionsReader *mock.MockCollectionReader
	collectionsWriter *mock.MockCollectionWriter
	documentReader    *mock.MockDocumentReader
	documentWriter    *mock.MockDocumentWriter
	indexReader       *mock.MockIndexReader
	indexWriter       *mock.MockIndexWriter
)

func initResource(t *testing.T) *restful.Container {
	ctrl := gomock.NewController(t)

	databaseReader = mock.NewMockDatabaseReader(ctrl)
	databaseWriter = mock.NewMockDatabaseWriter(ctrl)
	collectionsReader = mock.NewMockCollectionReader(ctrl)
	collectionsWriter = mock.NewMockCollectionWriter(ctrl)
	documentReader = mock.NewMockDocumentReader(ctrl)
	documentWriter = mock.NewMockDocumentWriter(ctrl)
	indexReader = mock.NewMockIndexReader(ctrl)
	indexWriter = mock.NewMockIndexWriter(ctrl)

	wsContainer := restful.NewContainer()

	NewDatabaseResource(
		databaseReader,
		databaseWriter,
		collectionsReader,
		collectionsWriter,
		documentReader,
		documentWriter,
		indexReader,
		indexWriter,
		middleware.NewJwtMiddleware(helper.PASSWORD),
	).Register(wsContainer)

	return wsContainer
}
