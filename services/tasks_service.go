package services

import (
	"time"
	"net/url"
	"panda-api/services/models"
	"panda-api/helpers"
)

func GetTasks(pag helpers.Pagination, q url.Values, userRequest string) models.Tasks {

	var tasks models.Tasks
	var typeDateQuery string

	db := Con

	if q.Get("title") != "" {
		db = db.Where("title iLIKE ?", "%" + q.Get("title") + "%")	
	}

	if q.Get("situation") == "open" {
		db = db.Where("completed_at IS NULL")
	}

	if q.Get("situation") == "done" {
		db = db.Where("completed_at IS NOT NULL")
	}

	switch q.Get("assigned") {
		case "author":	db = db.Where("registered_by_uuid = ?", userRequest)
		case "all": 	// all task
		default: 		db = db.Where("assignee_uuid = ?", userRequest)
	}

	switch q.Get("type_date") {
		case "registered":	typeDateQuery = "registered_at"
		case "due": 		typeDateQuery = "due"
		case "completed": 	typeDateQuery = "completed_at"
		default: 			typeDateQuery = "registered_at"
	}

	startDate := q.Get("start_date")
	endDate := q.Get("end_date")

	if startDate != "" && endDate != "" {
		db = db.Where(typeDateQuery + "::DATE BETWEEN ? AND ?", startDate, endDate)
	} else if startDate != "" {
		db = db.Where(typeDateQuery + "::DATE >= ?", startDate)
	} else if endDate != "" {
		db = db.Where(typeDateQuery + "::DATE <= ?", endDate)
	}
	
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