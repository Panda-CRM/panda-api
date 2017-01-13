package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/wilsontamarozzi/panda-api/controllers"
)

func AddRoutesTasks(r *gin.RouterGroup) {
    r.GET("/tasks", controllers.GetTasks)
    r.GET("/tasks/:id", controllers.GetTask)
    r.DELETE("/tasks/:id", controllers.DeleteTask)
    r.POST("/tasks", controllers.CreateTask)
    r.PUT("/tasks/:id", controllers.UpdateTask)
}