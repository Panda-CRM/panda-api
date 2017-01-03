package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"panda-api/services"
	"panda-api/services/models"
	"panda-api/helpers"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Tarefas

	Method: GET
	Rota: /tasks
*/
func GetTasks(c *gin.Context) {
	
	page		:= c.Query("page")
	itemPerPage	:= c.Query("per_page")

	count := services.CountRowsTask()

	pageConv, _ := strconv.Atoi(page)
	itemPerPageConv, _ := strconv.Atoi(itemPerPage)

	pag := helpers.MakePagination(count, pageConv, itemPerPageConv)

	var content models.Tasks
	content = services.GetTasks(pag)

	if len(content) <= 0 {
		c.JSON(200, gin.H{"errors": "Registros não encontrado."})
	} else {
		c.JSON(200, gin.H{
			"tasks": content, 
			"meta": gin.H{
				"pagination": pag,
			},
		})
	}
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Tarefa especifica pelo ID

	Method: GET
	Rota: /tasks/{id:[0-9]+}
*/
func GetTask(c *gin.Context) {

	taskId := c.Params.ByName("id")

	task := services.GetTask(taskId)

	if task.UUID == "" {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
	} else {
		c.JSON(200, gin.H{"task": task})
	}	
}

/*	@autor: Wilson T.J.

	Método responsável por deletar uma Tarefa especifica pelo ID

	Method: DELETE
	Rota: /tasks/{id:[0-9]+}
*/
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

/*	@autor: Wilson T.J.

	Método responsável por cadastrar uma Tarefa

	Method: POST
	Rota: /tasks
*/
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

/*	@autor: Wilson T.J.

	Método responsável por alterar uma Tarefa

	Method: PUT
	Rota: /tasks/{id:[0-9]+}
*/
func UpdateTask(c *gin.Context) {
	
	taskId := c.Params.ByName("id")

	task := services.GetTask(taskId)

	if task.UUID == "" {
		c.JSON(404, gin.H{"errors": "Registros não encontrado."})
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