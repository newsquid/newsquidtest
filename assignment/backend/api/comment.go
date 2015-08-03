package api

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"strconv"
)

type Comment struct {
	Id      int
	TodoId  int
	Title   string
	Content string
}

func SetupCommentRoutes(r martini.Router) {
	r.Get("/comments", ListComments)
	r.Get("/todo/:id/comments", ListCommentsByTodo)
	r.Post("/comments", binding.Json(Comment{}), CreateComment)
}

func ListComments() (int, string) {
	commentSlice := make([]Comment, len(comments))
	i := 0
	for _, v := range comments {
		commentSlice[i] = v
		i++
	}
	jsonByte, err := json.Marshal(commentSlice)

	if nil != err {
		return 500, err.Error()
	}

	return 200, string(jsonByte)
}

func ListCommentsByTodo(params martini.Params) (int, string) {
	todoId, err := strconv.Atoi(params["id"])
	if nil != err {
		return 400, "id must be an integer"
	}

	commentSlice := make([]Comment, 0)
	for _, v := range comments {
		if todoId == v.TodoId {
			commentSlice = append(commentSlice, v)
		}
	}

	jsonByte, err := json.Marshal(commentSlice)
	if nil != err {
		return 500, err.Error()
	}
	return 200, string(jsonByte)
}

func CreateComment(comment Comment) (int, string) {
	comment.Id = commentsNextId()

	jsonByte, err := json.Marshal(comment)
	if nil != err {
		return 500, err.Error()
	}

	comments[comment.Id] = comment
	return 200, string(jsonByte)
}

var comments map[int]Comment = map[int]Comment{
	1: Comment{Id: 1, TodoId: 1, Title: "No, because I", Content: "Got milk!"},
	2: Comment{Id: 2, TodoId: 1, Title: "Cereal", Content: "Would you get some cereal too hon?"},
	3: Comment{Id: 3, TodoId: 3, Title: "Jen isn't comming", Content: "She's got the flu"},
}

func commentsNextId() int {
	maxId := 0
	for _, v := range comments {
		if v.Id > maxId {
			maxId = v.Id
		}
	}
	return maxId + 1
}
