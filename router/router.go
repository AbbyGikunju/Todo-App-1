package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/github.com/AbbyGikunju/todo-app-1/db"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodo)
	r.POST("/todos", postTodo)
	r.PUT("/todos/:id", putTodo)
	r.DELETE("todos/:id", deleteTodo)

	return r

}

// posts a todo to the db
func postTodo(c *gin.Context) {
	var body struct {
		Description string
		Title       string
		Status      bool
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	todo := db.Todo{
		Title:       body.Title,
		Description: body.Description,
		Status:      body.Status,
	}

	res, err := db.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"todo" :res,
	})
}

//gets a todo from the db
func getTodos(c *gin.Context){
	res, err := db.ReadTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"todos": res,
	})
}

//get a todo by id
func getTodo(c *gin.Context) {
	id := c.Param("id")
	res, err := db.ReadTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"todo": res,
	})

}

//update a todo in the db
func putTodo(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		Status      bool
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	id := c.Param("id")
	dbTodo, err := db.ReadTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbTodo.Title = body.Title
	dbTodo.Description = body.Description
	dbTodo.Status = body.Status

	res, err := db.UpdateTodo(id, dbTodo)
	if err != nil {
		c.JSON(http.StatusBadRequest,  gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task":res,
	})
}


//delete a task from the db
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := db.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}