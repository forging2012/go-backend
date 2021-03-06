package document_writer

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Insert new document
// https://docs.mongodb.com/manual/tutorial/insert-documents/
func (rcv *documentWriter) Create(dbName value.DBName, collName value.CollName, doc interface{}) error {
	document, err := bson.NewDocumentEncoder().EncodeDocument(doc)

	if err != nil {
		return err
	}

	_, err = rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		InsertOne(context.Background(), document)

	return err
}
