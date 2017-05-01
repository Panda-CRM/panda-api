package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/models"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

type PersonController struct {
	Repository repositories.PersonRepositoryInterface
}

func (controller PersonController) GetAll(c *gin.Context) {
	params := c.Request.URL.Query()
	people := controller.Repository.GetAll(params)

	c.JSON(200, people)
}

func (controller PersonController) Get(c *gin.Context) {
	personId := c.Param("id")
	person := controller.Repository.Get(personId)
	// Valida se existe a pessoa (404)
	if person.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}

	c.JSON(200, gin.H{"person": person})
}

func (controller PersonController) Delete(c *gin.Context) {
	personId := c.Param("id")
	person := controller.Repository.Get(personId)
	// Valida se existe a pessoa que será excluida (404)
	if person.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}
	// Valida se deu erro ao tentar excluir (500)
	if err := controller.Repository.Delete(personId); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		return
	}

	c.Status(204)
}

func (controller PersonController) Create(c *gin.Context) {
	var person models.Person
	// Valida BadRequest (400)
	if err := c.BindJSON(&person); err != nil {
		c.JSON(400, gin.H{"errors: ": err.Error()})
		return
	}
	// Valida Invalid Entity (422)
	if err := person.Validate(); err != nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}
	// Seta usuário que está logado e fazendo cadastro
	person.RegisteredByUUID = c.MustGet("userRequest").(string)
	// Valida se deu erro ao inserir (500)
	if err := controller.Repository.Create(&person); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor"})
		return
	}

	c.JSON(201, gin.H{"person": person})
}

func (controller PersonController) Update(c *gin.Context) {
	personId := c.Param("id")
	person := controller.Repository.Get(personId)
	// Valida se existe a pessoa que será editada (404)
	if person.IsEmpty() {
		c.JSON(404, gin.H{"errors": "Registros não encontrado."})
		return
	}
	// Valida BadRequest (400)
	if err := c.BindJSON(&person); err != nil {
		c.JSON(400, gin.H{"errors: ": err.Error()})
		return
	}
	// Valida Invalid Entity (422)
	if err := person.Validate(); err == nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}
	// Valida se deu erro ao inserir (500)
	if err := controller.Repository.Update(&person); err != nil {
		c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		return
	}

	c.JSON(201, gin.H{"person": person})
}
