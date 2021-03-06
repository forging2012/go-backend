package value

import (
	"strings"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
)

// Custom string type for database name
// https://docs.mongodb.com/manual/reference/limits/#Database-Name-Case-Sensitivity
// https://docs.mongodb.com/manual/reference/limits/#Restrictions-on-Database-Names-for-Unix-and-Linux-Systems
type DBName string

// Validation method
func (rcv DBName) Valid() error {
	// Database names cannot be empty
	if len(strings.TrimSpace(string(rcv))) == 0 {
		return err.EmptyDBName
	}

	// For MongoDB deployments running on Unix and Linux systems, database names cannot contain any of the following characters
	if strings.ContainsAny(string(rcv), `/\. "$`) {
		return err.NotValidDBName
	}

	// Database names cannot have fewer than 64 characters.
	if len(rcv) > 64 {
		return err.InvalidDBNameMaxLength
	}

	return nil
}
