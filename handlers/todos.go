package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/gocondor/core"
	"github.com/harranali/gocondor-todo-app/models"
)

func ListTodos(c *core.Context) *core.Response {
	var todos []models.Todo
	result := c.GetGorm().Find(&todos)
	if result.Error != nil {
		return c.Response.SetStatusCode(500).Json(c.MapToJson(map[string]string{"message": result.Error.Error()}))
	}
	todosJson, err := json.Marshal(todos)
	if err != nil {
		return c.Response.Json(c.MapToJson(map[string]string{"message": err.Error()}))
	}

	return c.Response.Json(string(todosJson))
}

func CreateTodos(c *core.Context) *core.Response {
	title := c.CastToString(c.GetRequestParam("title"))
	body := c.CastToString(c.GetRequestParam("body"))
	isDone := c.CastToString(c.GetRequestParam("isDone"))
	v := c.GetValidator().Validate(map[string]interface{}{
		"title":  title,
		"body":   body,
		"isDone": isDone,
	}, map[string]interface{}{
		"title": "required",
		"body":  "required",
	})
	if v.Failed() {
		return c.Response.Json(v.GetErrorMessagesJson())
	}
	result := c.GetGorm().Create(&models.Todo{
		Title:  title,
		Body:   body,
		IsDone: false,
	})
	if result.Error != nil {
		return c.Response.SetStatusCode(500).Json(c.MapToJson(map[string]string{
			"message": result.Error.Error(),
		}))
	}

	return c.Response.Json(c.MapToJson(map[string]string{
		"message": "created successfully",
	}))
}

func ShowTodo(c *core.Context) *core.Response {
	todoID := c.CastToString(c.GetPathParam("id"))
	var todo models.Todo
	result := c.GetGorm().First(&todo, todoID)
	if result.Error != nil {
		return c.Response.SetStatusCode(500).Json(c.MapToJson(map[string]string{"message": result.Error.Error()}))
	}
	todoJson, err := json.Marshal(todo)
	if err != nil {
		return c.Response.Json(c.MapToJson(map[string]string{"message": err.Error()}))
	}

	return c.Response.Json(string(todoJson))
}

func DeleteTodo(c *core.Context) *core.Response {
	todoID := c.CastToString(c.GetPathParam("id"))
	var todo models.Todo
	result := c.GetGorm().Delete(&todo, todoID)
	if result.Error != nil {
		return c.Response.SetStatusCode(500).Json(c.MapToJson(map[string]string{"message": result.Error.Error()}))
	}

	return c.Response.Json(c.MapToJson(map[string]string{"message": "record deleted successfully"}))
}

func UpdateTodo(c *core.Context) *core.Response {
	var title string = ""
	var body string = ""
	var data map[string]interface{} = map[string]interface{}{}
	var rules map[string]interface{} = map[string]interface{}{}
	todoID := c.GetPathParam("id")
	var todo models.Todo
	result := c.GetGorm().First(&todo, todoID)
	if result.Error != nil {
		return c.Response.Json(c.MapToJson(map[string]string{"message": result.Error.Error()}))
	}
	if c.RequestParamExists("title") {
		title = c.CastToString(c.GetRequestParam("title"))
		data["title"] = title
	}
	if c.RequestParamExists("body") {
		body = c.CastToString(c.GetRequestParam("body"))
		data["body"] = body
	}
	if c.RequestParamExists("isDone") {
		isDoneStr := c.CastToString(c.GetRequestParam("isDone"))
		data["isDone"] = isDoneStr
		rules["isDone"] = "in:true,false"
	}
	v := c.GetValidator().Validate(data, rules)
	if v.Failed() {
		return c.Response.Json(v.GetErrorMessagesJson())
	}
	if c.RequestParamExists("title") {
		todo.Title = title
	}
	if c.RequestParamExists("body") {
		todo.Body = body
	}
	if c.RequestParamExists("isDone") {
		isDoneStr := c.CastToString(c.GetRequestParam("isDone"))
		isDone, err := strconv.ParseBool(isDoneStr)
		if err != nil {
			return c.Response.Json(c.MapToJson(map[string]string{"message": err.Error()}))
		}
		todo.IsDone = isDone
	}
	c.GetGorm().Save(&todo)
	todoJson, err := json.Marshal(todo)
	if err != nil {
		return c.Response.Json(c.MapToJson(map[string]string{"message": err.Error()}))
	}

	return c.Response.Json(string(todoJson))
}
