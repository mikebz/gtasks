// Copyright 2024 Mike Borozdin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	tlCall.ShowHidden(true)
	tlCall.ShowCompleted(false)
	tlCall.ShowAssigned(true)
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
