package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"panda-api/services"
	"panda-api/services/models"
	"panda-api/helpers"
)

func GetTasks(c *gin.Context) {
	
	q := c.Request.URL.Query()

	count := services.CountRowsTask()

	page, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))

	pag := helpers.MakePagination(count, page, itemPerPage)

	var content models.Tasks
	content = services.GetTasks(pag, q, c.MustGet("userRequest").(string))

	if len(content) <= 0 {
		c.JSON(200, gin.H{
			"errors": "Registros não encontrado.",
			"meta": gin.H{
				"pagination": pag,
			},
		})
	} else {
		c.JSON(200, gin.H{
			"tasks": content, 
			"meta": gin.H{
				"pagination": pag,
			},
		})
	}
}

func GetTask(c *gin.Context) {

	taskId := c.Params.ByName("id")

	task := services.GetTask(taskId)

	if task.UUID == "" {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
	} else {
		c.JSON(200, gin.H{"task": task})
	}	
}

func DeleteTask(c *gin.Context) {

	taskId := c.Params.ByName("id")

	task := services.GetTask(taskId)

	if task.UUID == "" {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
	} else {
		err := services.DeleteTask(taskId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		}
	}
}

func CreateTask(c *gin.Context) {

	var request models.TaskRequest
	err := c.BindJSON(&request)

	if err == nil {

		task := request.Task

		err := task.Validate()

		if err == nil {
			task.RegisteredByUUID = c.MustGet("userRequest").(string)
			
			err := services.CreateTask(task)
			
			if err == nil {
				c.JSON(201, task)
			} else {
				c.JSON(500, gin.H{"errors": "Houve um erro no servidor"})
			}
		} else {
			c.JSON(422, gin.H{"errors" : err})
		}
	} else {
		c.JSON(400, gin.H{"errors: " : err.Error()})
	}
}

func UpdateTask(c *gin.Context) {
	
	taskId := c.Params.ByName("id")

	task := services.GetTask(taskId)

	if task.UUID == "" {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
	} else {
		
		var request models.TaskRequest
		err := c.BindJSON(&request)

		if err == nil {

			task := request.Task

			err := task.Validate()

			if err == nil {
				task.RegisteredByUUID = c.MustGet("userRequest").(string)

				err := services.UpdateTask(task)

				if err == nil {
					c.JSON(201, task)
				} else {
					c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
				}
			} else {
				c.JSON(422, gin.H{"errors" : err})
			}
		} else {
			c.JSON(400, gin.H{"errors: " : err.Error()})
		}
	}
}