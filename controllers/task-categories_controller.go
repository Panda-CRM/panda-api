package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"panda-api/services"
	"panda-api/services/models"
	"panda-api/helpers"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Categorias de Tarefa

	Method: GET
	Rota: /task-categories
*/
func GetTaskCategories(c *gin.Context) {
	
	page		:= c.Query("page")
	itemPerPage	:= c.Query("per_page")
	number		:= c.Query("number")
	title 		:= c.Query("title")

	count := services.CountRowsTaskCategory()

	pageConv, _ := strconv.Atoi(page)
	itemPerPageConv, _ := strconv.Atoi(itemPerPage)

	pag := helpers.MakePagination(count, pageConv, itemPerPageConv)

	var content models.TaskCategories
	content = services.GetTaskCategories(pag, number, title)

	if len(content) <= 0 {
		c.JSON(200, gin.H{"errors": "Registros não encontrado."})
	} else {
		c.JSON(200, gin.H{
			"task_categories": content, 
			"meta": gin.H{
				"pagination": pag,
			},
		})
	}
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Categoria de Tarefa especifica pelo ID

	Method: GET
	Rota: /task-categories/{id:[0-9]+}
*/
func GetTaskCategory(c *gin.Context) {

	taskCategoryId := c.Params.ByName("id")

	taskCategory := services.GetTaskCategory(taskCategoryId)

	if taskCategory == (models.TaskCategory{}) {
		c.JSON(404, gin.H{"errors": "Registros não encontrado."})
	} else {
		c.JSON(200, gin.H{"taskCategory": taskCategory})
	}	
}

/*	@autor: Wilson T.J.

	Método responsável por deletar uma Categoria de Tarefa especifica pelo ID

	Method: DELETE
	Rota: /task-categories/{id:[0-9]+}
*/
func DeleteTaskCategory(c *gin.Context) {

	taskCategoryId := c.Params.ByName("id")

	taskCategory := services.GetTaskCategory(taskCategoryId)

	if taskCategory == (models.TaskCategory{}) {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
	} else {
		err := services.DeleteTaskCategory(taskCategoryId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		}
	}
}

/*	@autor: Wilson T.J.

	Método responsável por cadastrar uma Categoria de Tarefa

	Method: POST
	Rota: /task-categories
*/
func CreateTaskCategory(c *gin.Context) {

	var request models.TaskCategoryRequest
	c.BindJSON(&request)

	taskCategory := request.TaskCategory

	err := taskCategory.Validate()

	if err != nil {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		// Cria o UUID
		taskCategory.UUID = uuid.NewV4().String()

		err := services.CreateTaskCategory(taskCategory)
		
		if err != nil {
			c.JSON(500, gin.H{"errors": "Houve um erro no servidor"})
		} else {
			c.JSON(201, taskCategory)
		}
	}
}

/*	@autor: Wilson T.J.

	Método responsável por alterar uma Categoria de Tarefa

	Method: PUT
	Rota: /task-categories/{id:[0-9]+}
*/
func UpdateTaskCategory(c *gin.Context) {
	
	taskCategoryId := c.Params.ByName("id")

	taskCategory := services.GetTaskCategory(taskCategoryId)

	if taskCategory == (models.TaskCategory{}) {
		c.JSON(404, gin.H{"errors": "Registros não encontrado."})
	} else {
		
		var request models.TaskCategoryRequest
		c.BindJSON(&request)

		taskCategory := request.TaskCategory

		err := taskCategory.Validate()

		if err != nil {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateTaskCategory(taskCategory)

			if err != nil {
				c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
			} else {
				c.JSON(201, taskCategory)
			}
		}
	}
}