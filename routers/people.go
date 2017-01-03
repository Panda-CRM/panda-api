package routers

import (
    "github.com/gin-gonic/gin"
    "panda-api/controllers"
)

func AddRoutesPeople(r *gin.RouterGroup) {
    r.GET("/people", controllers.GetPeople)
    r.GET("/people/:id", controllers.GetPerson)
    r.DELETE("/people/:id", controllers.DeletePerson)
    r.POST("/people", controllers.CreatePerson)
    r.PUT("/people/:id", controllers.UpdatePerson)
}