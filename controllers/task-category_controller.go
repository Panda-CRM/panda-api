package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/Panda-CRM/panda-api/models"
	"github.com/Panda-CRM/panda-api/repositories"
)

type TaskCategoryController struct {
	Repository repositories.TaskCategoryRepository
}

func (controller *TaskCategoryController) List(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	categories := controller.Repository.List(queryParams)

	c.JSON(200, categories)
}

func (controller *TaskCategoryController) Get(c *gin.Context) {
	categoryId := c.Param("id")
	category := controller.Repository.Get(categoryId)

	if category.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registros não encontrado."})
		return
	}

	c.JSON(200, gin.H{"task_category": category})
}

func (controller *TaskCategoryController) Delete(c *gin.Context) {
	categoryId := c.Param("id")
	category := controller.Repository.Get(categoryId)

	if category.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}

	if err := controller.Repository.Delete(categoryId); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		return
	}

	c.Status(204)
}

func (controller *TaskCategoryController) Create(c *gin.Context) {
	var category models.TaskCategory

	if err := c.BindJSON(&category); err != nil {
		c.JSON(400, gin.H{"errors: ": err.Error()})
		return
	}

	if err := category.Validate(); err != nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}

	if err := controller.Repository.Create(&category); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor"})
		return
	}

	c.JSON(201, gin.H{"task_category": category})
}

func (controller *TaskCategoryController) Update(c *gin.Context) {
	categoryId := c.Param("id")
	category := controller.Repository.Get(categoryId)

	if category.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}

	if err := c.BindJSON(&category); err != nil {
		c.JSON(400, gin.H{"errors: ": err.Error()})
		return
	}

	if err := category.Validate(); err != nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}

	if err := controller.Repository.Update(&category); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		return
	}

	c.JSON(201, gin.H{"task_category": category})
}
