package tasks

import (
	"context"
	"errors"

	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

// List return a list of tasks
// in a text format.
func Lists() ([]string, error) {
	ctx := context.Background()
	tasksService, err := tasks.NewService(
		ctx,
		option.WithScopes(tasks.TasksReadonlyScope),
	)
	if err != nil {
		return nil, err
	}

	listCall := tasksService.Tasklists.List()
	listResult, err := listCall.Do()
	if err != nil {
		return nil, err
	}

	result := make([]string, len(listResult.Items))

	for i, item := range listResult.Items {
		result[i] = item.Title
	}

	return result, nil
}

// Tasks return a list of tasks
// in a text format.
func Tasks() ([]string, error) {
	ctx := context.Background()
	tasksService, err := tasks.NewService(
		ctx,
		option.WithScopes(tasks.TasksReadonlyScope),
	)
	if err != nil {
		return nil, err
	}

	listCall := tasksService.Tasklists.List()
	listResult, err := listCall.Do()
	if err != nil {
		return nil, err
	}

	if len(listResult.Items) == 0 {
		return nil, errors.New("no tasklists found")
	}

	tasklist := listResult.Items[0]

	tlCall := tasksService.Tasks.List(tasklist.Id)
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
