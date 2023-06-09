// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"go.infratographer.com/example-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

type Tenant struct {
	ID   gidx.PrefixedID           `json:"id"`
	Todo *generated.TodoConnection `json:"todo"`
}

func (Tenant) IsEntity() {}

// Return response for createTodo mutation
type TodoCreatePayload struct {
	// Created todo
	Todo *generated.Todo `json:"todo"`
}

// Return response for deleteTodo mutation
type TodoDeletePayload struct {
	// Deleted todo
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response for updateTodo mutation
type TodoUpdatePayload struct {
	// Updated todo
	Todo *generated.Todo `json:"todo"`
}
