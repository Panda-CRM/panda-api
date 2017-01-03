package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"panda-api/services"
	"panda-api/services/models"
	"panda-api/helpers"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Pessoas

	Method: GET
	Rota: /people
*/
func GetPeople(c *gin.Context) {
	
	page		:= c.Query("page")
	itemPerPage	:= c.Query("per_page")
	code		:= c.Query("code")
	name 		:= c.Query("name")

	count := services.CountRowsPerson()

	pageConv, _ := strconv.Atoi(page)
	itemPerPageConv, _ := strconv.Atoi(itemPerPage)

	pag := helpers.MakePagination(count, pageConv, itemPerPageConv)

	var content models.People
	content = services.GetPeople(pag, code, name)

	if len(content) <= 0 {
		c.JSON(200, gin.H{"errors": "Registros não encontrado."})
	} else {
		c.JSON(200, gin.H{
			"people": content, 
			"meta": gin.H{
				"pagination": pag,
			},
		})
	}
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa especifica pelo ID

	Method: GET
	Rota: /people/{id:[0-9]+}
*/
func GetPerson(c *gin.Context) {

	personId := c.Params.ByName("id")

	person := services.GetPerson(personId)

	if person == (models.Person{}) {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
	} else {
		c.JSON(200, gin.H{"person": person})
	}	
}

/*	@autor: Wilson T.J.

	Método responsável por deletar uma Pessoa especifica pelo ID

	Method: DELETE
	Rota: /people/{id:[0-9]+}
*/
func DeletePerson(c *gin.Context) {

	personId := c.Params.ByName("id")

	person := services.GetPerson(personId)

	if person == (models.Person{}) {
		c.JSON(404, gin.H{"errors": "Registro não encontrado."})
	} else {
		err := services.DeletePerson(personId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"errors": "Houve um erro no servidor."})
		}
	}
}

/*	@autor: Wilson T.J.

	Método responsável por cadastrar uma Pessoa

	Method: POST
	Rota: /people
*/
func CreatePerson(c *gin.Context) {

	var request models.PersonRequest
	err := c.BindJSON(&request)

	if err == nil {

		person := request.Person

		err := person.Validate()

		if err == nil {
			person.RegisteredByUUID = c.MustGet("userRequest").(string)

			err := services.CreatePerson(person)
			
			if err == nil {
				c.JSON(201, person)
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

	Método responsável por alterar uma Pessoa

	Method: PUT
	Rota: /people/{id:[0-9]+}
*/
func UpdatePerson(c *gin.Context) {
	
	personId := c.Params.ByName("id")

	person := services.GetPerson(personId)

	if person == (models.Person{}) {
		c.JSON(404, gin.H{"errors": "Registros não encontrado."})
	} else {
		
		var request models.PersonRequest
		err := c.BindJSON(&request)

		if err == nil {

			person := request.Person

			err := person.Validate()

			if err == nil {
				err := services.UpdatePerson(person)

				if err == nil {
					c.JSON(201, person)
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