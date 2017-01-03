package routers

import (
    "github.com/gin-gonic/gin"
    "panda-api/controllers"
)

func AddRoutesTaskCategories(r *gin.RouterGroup) {
    r.GET("/task_categories", controllers.GetTaskCategories)
    r.GET("/task_categories/:id", controllers.GetTaskCategory)
    r.DELETE("/task_categories/:id", controllers.DeleteTaskCategory)
    r.POST("/task_categories", controllers.CreateTaskCategory)
    r.PUT("/task_categories/:id", controllers.UpdateTaskCategory)
}