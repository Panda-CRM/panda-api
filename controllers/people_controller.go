package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/models"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

type PersonController struct {
	Repository repositories.PersonRepository
}

func (p PersonController) List(gc *gin.Context) {
	params := gc.Request.URL.Query()
	people := p.Repository.List(params)
	gc.JSON(200, people)
}

func (p PersonController) Get(gc *gin.Context) {
	personId := gc.Param("id")
	person := p.Repository.Get(personId)
	// Valida se existe a pessoa (404)
	if person.IsEmpty() {
		gc.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}

	gc.JSON(200, gin.H{"person": person})
}

func (p PersonController) Delete(gc *gin.Context) {
	personId := gc.Param("id")
	person := p.Repository.Get(personId)
	// Valida se existe a pessoa que será excluida (404)
	if person.IsEmpty() {
		gc.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}
	// Valida se deu erro ao tentar excluir (500)
	if err := p.Repository.Delete(personId); err != nil {
		gc.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		return
	}

	gc.Status(204)
}

func (p PersonController) Create(gc *gin.Context) {
	var person models.Person
	// Valida BadRequest (400)
	if err := gc.BindJSON(&person); err != nil {
		gc.JSON(400, gin.H{"errors: ": err.Error()})
		return
	}
	// Valida Invalid Entity (422)
	if err := person.Validate(); err != nil {
		gc.JSON(422, gin.H{"errors": err})
		return
	}
	// Seta usuário que está logado e fazendo cadastro
	person.RegisteredByUUID = gc.MustGet("userRequest").(string)
	// Valida se deu erro ao inserir (500)
	if err := p.Repository.Create(&person); err != nil {
		gc.JSON(500, gin.H{"errors": "Houve um erro no servidor"})
		return
	}

	gc.JSON(201, gin.H{"person": person})
}

func (p PersonController) Update(gc *gin.Context) {
	personId := gc.Param("id")
	person := p.Repository.Get(personId)
	// Valida se existe a pessoa que será editada (404)
	if person.IsEmpty() {
		gc.JSON(404, gin.H{"errors": "Registros não encontrado."})
		return
	}
	// Valida BadRequest (400)
	if err := gc.BindJSON(&person); err != nil {
		gc.JSON(400, gin.H{"errors: ": err.Error()})
		return
	}
	// Valida Invalid Entity (422)
	if err := person.Validate(); err != nil {
		gc.JSON(422, gin.H{"errors": err})
		return
	}
	// Valida se deu erro ao inserir (500)
	if err := p.Repository.Update(&person); err != nil {
		gc.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		return
	}

	gc.JSON(201, gin.H{"person": person})
}
