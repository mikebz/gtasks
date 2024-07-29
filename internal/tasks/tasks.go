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
	"strings"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

// List return a list of tasks
// in a text format.
func Lists() ([]string, error) {
	tasksService, err := createTaskService()
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
func Tasks(hidden bool, completed bool, assigned bool) ([]*tasks.Task, error) {
	tasksService, err := createTaskService()
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
	tlCall.ShowHidden(hidden)
	tlCall.ShowCompleted(completed)
	tlCall.ShowAssigned(assigned)
	tasksList, err := tlCall.Do()
	if err != nil {
		return nil, err
	}
	return tasksList.Items, nil
}

func TaskToLine(task *tasks.Task) string {
	return `- ` + task.Title
}

func TaskVerbose(task *tasks.Task) string {

	var r strings.Builder
	r.WriteString("- " + task.Title + "\n")
	if task.Due != "" {
		r.WriteString("  Due: " + taskDueString(task) + "\n")
	}
	r.WriteString("  " + task.Notes + "\n")

	return r.String()
}

func taskDueString(task *tasks.Task) string {
	parsed, err := time.Parse(time.RFC3339, task.Due)
	if err == nil {
		formatted := parsed.Format("2006-01-02")
		return formatted
	}
	return ""
}

// Single function to create a task service
// created as one function because we are anticipating
// that the authentication methods are going to make this
// much more complex going forward.
func createTaskService() (*tasks.Service, error) {
	ctx := context.Background()
	tasksService, err := tasks.NewService(
		ctx,
		option.WithScopes(tasks.TasksReadonlyScope),
	)
	return tasksService, err
}
