package database_writer

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Removes the current database, deleting the associated data files.
// https://docs.mongodb.com/manual/reference/method/db.dropDatabase/
func (rcv *databaseWriter) Delete(name value.DBName) error {
	return rcv.db.DB(string(name)).DropDatabase()
}
