package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/services"
	"github.com/wilsontamarozzi/panda-api/services/models"
	"github.com/wilsontamarozzi/panda-api/helpers"
)

func GetTaskCategories(c *gin.Context) {
	
	queryParams := c.Request.URL.Query()

	amountItems := services.CountRowsTaskCategory()

	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	itemPerPage, _ := strconv.Atoi(queryParams.Get("per_page"))

	pagination := helpers.MakePagination(amountItems, currentPage, itemPerPage)

	var content models.TaskCategories
	content = services.GetTaskCategories(pagination, queryParams)

	if len(content) <= 0 {
		c.JSON(200, gin.H{
			"errors": "Registros não encontrado.",
			"meta": gin.H{
				"pagination": pagination,
			},
		})
	} else {
		c.JSON(200, gin.H{
			"task_categories": content, 
			"meta": gin.H{
				"pagination": pagination,
			},
		})
	}
}

func GetTaskCategory(c *gin.Context) {

	taskCategoryId := c.Params.ByName("id")

	taskCategory := services.GetTaskCategory(taskCategoryId)

	if taskCategory == (models.TaskCategory{}) {
		c.JSON(404, gin.H{"errors": "Registros não encontrado."})
	} else {
		c.JSON(200, gin.H{"task_category": taskCategory})
	}	
}

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

func CreateTaskCategory(c *gin.Context) {

	var request models.TaskCategoryRequest
	err := c.BindJSON(&request)

	if err == nil {

		taskCategory := request.TaskCategory

		err := taskCategory.Validate()

		if err == nil {

			taskCategory, err := services.CreateTaskCategory(taskCategory)
			
			if err == nil {
				c.JSON(201, gin.H{"task_category": taskCategory})
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

func UpdateTaskCategory(c *gin.Context) {
	
	taskCategoryId := c.Params.ByName("id")

	taskCategory := services.GetTaskCategory(taskCategoryId)

	if taskCategory == (models.TaskCategory{}) {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
	} else {
		
		var request models.TaskCategoryRequest
		err := c.BindJSON(&request)

		if err == nil {

			taskCategory := request.TaskCategory

			err := taskCategory.Validate()

			if err == nil {
				taskCategory, err := services.UpdateTaskCategory(taskCategory)

				if err == nil {
					c.JSON(201, gin.H{"task_category": taskCategory})
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