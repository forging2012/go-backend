package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIndex(t *testing.T) {
	name := "MyIndex"
	fields := []string{"name"}

	index := NewIndex(name, true, true, true, fields)

	assert.Equal(t, name, index.Name())
	assert.Equal(t, fields, index.Fields())
	assert.True(t, index.Unique())
	assert.True(t, index.Background())
	assert.True(t, index.Sparse())
}

func TestNewIndex2(t *testing.T) {
	name := "MyIndex2"
	fields := []string{"name", "gender"}

	index := NewIndex(name, false, false, false, fields)

	assert.Equal(t, name, index.Name())
	assert.Equal(t, fields, index.Fields())
	assert.False(t, index.Unique())
	assert.False(t, index.Background())
	assert.False(t, index.Sparse())
}
