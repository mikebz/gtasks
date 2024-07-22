package tasks

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

// List return a list of tasks
// in a text format.
func List() ([]string, error) {
	ctx := context.Background()
	tasksService, err := tasks.NewService(
		ctx,
		option.WithScopes(tasks.TasksReadonlyScope),
	)
	if err != nil {
		return nil, err
	}

	tlCall := tasksService.Tasks.List(`default`)
	tasksList, err := tlCall.Do()
	if err != nil {
		return nil, err
	}

	result := make([]string, len(tasksList.Items))

	for i, task := range tasksList.Items {
		result[i] = task.Title
	}

	return result, nil
}
