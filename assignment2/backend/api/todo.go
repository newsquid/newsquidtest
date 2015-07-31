package api

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

/*
Setup of routing / endpoints
*/

type Todo struct {
	Id      int
	Title   string
	Content string
}

func SetupTodoRoutes(r martini.Router) {
	r.Get("/todos", ListTodos)
	r.Post("/todos", binding.Json(Todo{}), CreateTodo)
	r.Put("/todos", binding.Json(Todo{}), UpdateTodo)
}

func ListTodos() (int, string) {
	todoSlice := make([]Todo, len(todos))
	i := 0
	for _, v := range todos {
		todoSlice[i] = v
		i++
	}
	jsonByte, err := json.Marshal(todoSlice)

	if nil != err {
		return 500, err.Error()
	}

	return 200, string(jsonByte)
}

func UpdateTodo(todo Todo) (int, string) {
	_, exists := todos[todo.Id]
	if !exists {
		return 404, "Not found"
	}

	jsonByte, err := json.Marshal(todo)
	if nil != err {
		return 500, err.Error()
	}

	todos[todo.Id] = todo
	return 200, string(jsonByte)
}

func CreateTodo(todo Todo) (int, string) {
	todo.Id = todosNextId()

	jsonByte, err := json.Marshal(todo)
	if nil != err {
		return 500, err.Error()
	}

	todos[todo.Id] = todo
	return 200, string(jsonByte)
}

var todos map[int]Todo = map[int]Todo{
	1: Todo{Id: 1, Title: "groceries", Content: "get milk!"},
	2: Todo{Id: 2, Title: "Game Night", Content: "call jen"},
	3: Todo{Id: 3, Title: "Walk the dog", Content: "... before 9 AM"},
}

func todosNextId() int {
	maxId := 0
	for _, v := range todos {
		if v.Id > maxId {
			maxId = v.Id
		}
	}
	return maxId + 1
}
