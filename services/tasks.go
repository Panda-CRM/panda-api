package services

import (
	"time"
	"panda-api/services/models"
	"panda-api/helpers"
)

func GetTasks(pag helpers.Pagination) models.Tasks {

	var tasks models.Tasks

	db := Con
	
	db.Preload("Category").
		Preload("RegisteredBy").
		Preload("Assignee").
		Preload("Person").
		Limit(pag.ItemPerPage).
		Offset(pag.StartIndex).
		Order("registered_at desc").
		Find(&tasks)

    return tasks
}

func GetTask(taskId string) models.Task {

	var task models.Task

	Con.Preload("Category").
		Preload("RegisteredBy").
		Preload("Assignee").
		Preload("TaskHistorics.RegisteredBy").
		Preload("Person").
		Where("uuid = ?", taskId).
		First(&task)

	return task
}

func DeleteTask(taskId string) error {
	return Con.Where("uuid = ?", taskId).Delete(&models.Task{}).Error
}

func CreateTask(task models.Task) error {
	
	record := models.Task{
		Title 				: task.Title,
		Due 				: task.Due,
		Visualized 			: false,
		CompletedAt 		: task.CompletedAt,
		RegisteredAt 		: time.Now(),
		RegisteredByUUID 	: task.RegisteredByUUID,
		CategoryUUID 		: task.Category.UUID,
		PersonUUID 			: task.Person.UUID,
		AssigneeUUID 		: task.Assignee.UUID,
	}

	err := Con.Set("gorm:save_associations", false).
		Table("tasks").
		Create(&record).Error

	if err != nil {
		return err
	}

	return CreateTaskComment(task.TaskHistorics, task.RegisteredByUUID, record.UUID)
}

func UpdateTask(task models.Task) error {
	
	err := Con.Set("gorm:save_associations", false).
		Table("tasks").
		Where("uuid = ?", task.UUID).
		Updates(models.Task{
			Title 			: task.Title,
			Due 			: task.Due,
			CompletedAt 	: task.CompletedAt,
			CategoryUUID 	: task.Category.UUID,
			PersonUUID 		: task.Person.UUID,
			AssigneeUUID 	: task.Assignee.UUID,
		}).Error

	if err != nil {
		return err
	}

	return CreateTaskComment(task.TaskHistorics, task.RegisteredByUUID, task.UUID)
}

func CountRowsTask() int {
	var count int
	Con.Table("tasks").Count(&count)

	return count
}

func CreateTaskComment(historics models.TaskHistorics, registeredByUUID string, taskUUID string) error {
	
	for _, historic := range historics {
		historic.RegisteredByUUID 	= registeredByUUID
		historic.RegisteredAt 		= time.Now()
		historic.TaskUUID 			= taskUUID

		if err := Con.Set("gorm:save_associations", false).Create(&historic).Error; err != nil {
			return err
		}
	}

	return nil
}