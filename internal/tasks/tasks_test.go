package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTasks(t *testing.T) {

	// Call the List function with the mock service
	tasks, err := Tasks()
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
}

func TestList(t *testing.T) {

	// Call the List function with the mock service
	tasks, err := Lists()
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
}
