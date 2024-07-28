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
