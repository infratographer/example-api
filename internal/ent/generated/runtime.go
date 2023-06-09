// Copyright 2023 The Infratographer Authors
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
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"time"

	"go.infratographer.com/example-api/internal/ent/generated/todo"
	"go.infratographer.com/example-api/internal/ent/schema"
	"go.infratographer.com/x/gidx"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoMixin := schema.Todo{}.Mixin()
	todoMixinFields0 := todoMixin[0].Fields()
	_ = todoMixinFields0
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoMixinFields0[0].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoMixinFields0[1].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(func() time.Time)
	// todo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	todo.UpdateDefaultUpdatedAt = todoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// todoDescName is the schema descriptor for name field.
	todoDescName := todoFields[1].Descriptor()
	// todo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	todo.NameValidator = todoDescName.Validators[0].(func(string) error)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoFields[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() gidx.PrefixedID)
}
