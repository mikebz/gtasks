package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {

	// Call the List function with the mock service
	tasks, err := List()
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
}
