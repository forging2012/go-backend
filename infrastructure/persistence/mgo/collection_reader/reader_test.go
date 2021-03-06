package collection_reader

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
)

func TestImplements(t *testing.T) {
	assert.Implements(t, (*repository.CollectionReader)(nil), new(collectionReader))
}
