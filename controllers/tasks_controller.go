package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/models"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

type TaskController struct{
	Repository repositories.TaskRepositoryInterface
}

func (controller TaskController) GetAll(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	queryParams.Add("user_request", c.MustGet("userRequest").(string))
	tasks := controller.Repository.GetAll(queryParams)

	c.JSON(200, tasks)
}

func (controller TaskController) Get(c *gin.Context) {
	taskId := c.Param("id")
	task := controller.Repository.Get(taskId)

	if task.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}

	c.JSON(200, gin.H{"task": task})
}

func (controller TaskController) Delete(c *gin.Context) {
	taskId := c.Param("id")
	task := controller.Repository.Get(taskId)

	if task.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}

	if err := controller.Repository.Delete(taskId); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		return
	}

	c.Status(204)
}

func (controller TaskController) Create(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"errors: ": err.Error()})
		return
	}

	if err := task.Validate(); err != nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}

	task.RegisteredByUUID = c.MustGet("userRequest").(string)
	if err := controller.Repository.Create(&task); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor"})
		return
	}

	c.JSON(201, gin.H{"task": task})
}

func (controller TaskController) Update(c *gin.Context) {
	taskId := c.Param("id")
	task := controller.Repository.Get(taskId)

	if task.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}

	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"errors: ": err.Error()})
		return
	}

	if err := task.Validate(); err != nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}

	task.RegisteredByUUID = c.MustGet("userRequest").(string)
	if err := controller.Repository.Update(&task); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		return
	}

	c.JSON(201, gin.H{"task": task})
}
