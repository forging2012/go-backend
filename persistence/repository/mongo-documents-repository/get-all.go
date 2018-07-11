package mongo_documents_repository

import (
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"gopkg.in/mgo.v2"
)

func (rcv *documentsRepository) GetAll(conditions *repository.GetListConditions) ([]interface{}, error) {
	collection := rcv.db.DB(conditions.DatabaseName()).C(conditions.CollectionName())
	var result []interface{}
	var query *mgo.Query

	if len(conditions.Filter()) != 0 {
		query = collection.Find(conditions.Filter())
	} else {
		query = collection.Find(nil)
	}

	query.Limit(conditions.Limit()).Skip(conditions.Skip())

	if len(conditions.Sort()) != 0 {
		query.Sort(conditions.Sort()...)
	}

	if err := query.All(&result); err != nil {
		return nil, err
	}

	return result, nil
}
