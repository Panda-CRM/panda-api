package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/services"
	"github.com/wilsontamarozzi/panda-api/services/models"
	"github.com/wilsontamarozzi/panda-api/helpers"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Pessoas

	Method: GET
	Rota: /people
*/
func GetPeople(c *gin.Context) {

	q := c.Request.URL.Query()

	count := services.CountRowsPerson()

	page, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))

	pag := helpers.MakePagination(count, page, itemPerPage)

	var content models.People
	content = services.GetPeople(pag, q)

	if len(content) <= 0 {
		c.JSON(200, gin.H{
			"errors": "Registros não encontrado.",
			"meta": gin.H{
				"pagination": pag,
			},
		})
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

			person, err := services.CreatePerson(person)
			
			if err == nil {
				c.JSON(201, gin.H{"person": person})
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
				person, err := services.UpdatePerson(person)

				if err == nil {
					c.JSON(201, gin.H{"person": person})
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