package repositories

import (
	"github.com/wilsontamarozzi/panda-api/logger"
	"github.com/wilsontamarozzi/panda-api/models"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/jinzhu/gorm"
	"time"
	"strconv"
	"net/url"
	"github.com/wilsontamarozzi/panda-api/database"
	"log"
)

type TaskRepositoryInterface interface{
	GetAll(q url.Values) models.Tasks
	Get(id string) models.Task
	Delete(id string) error
	Create(t *models.Task) error
	Update(t *models.Task) error
	CountRows() int
}

type taskRepository struct{}

func NewTaskRepository() *taskRepository {
	return new(taskRepository)
}

func (repository taskRepository) GetAll(q url.Values) models.Tasks {
	db := database.GetInstance()

	currentPage, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))
	pagination := helpers.MakePagination(repository.CountRows(), currentPage, itemPerPage)

	if q.Get("title") != "" {
		db = db.Where("title iLIKE ?", "%" + q.Get("title") + "%")
	}

	if q.Get("situation") == "open" {
		db = db.Where("completed_at IS NULL")
	}

	if q.Get("situation") == "done" {
		db = db.Where("completed_at IS NOT NULL")
	}

	userRequest := q.Get("user_request")

	switch q.Get("assigned") {
	case "author":	db = db.Where("registered_by_uuid = ?", userRequest)
	case "all": 	// all task
	default: 		db = db.Where("assignee_uuid = ?", userRequest)
	}

	var typeDateQuery string

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

	var tasks models.Tasks
	tasks.Meta.Pagination = pagination

	db.Preload("Category").
		Preload("RegisteredBy").
		Preload("Assignee").
		Preload("Person").
		Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("registered_at desc").
		Find(&tasks)

	return tasks
}

func (repository taskRepository) Get(id string) models.Task {
	db := database.GetInstance()

	var task models.Task
	db.Preload("Category").
		Preload("RegisteredBy").
		Preload("Assignee").
		Preload("TaskHistorics.RegisteredBy").
		Preload("Person").
		Where("uuid = ?", id).
		First(&task)

	return task
}

func (repository taskRepository) Delete(id string) error {
	db := database.GetInstance()

	err := db.Where("uuid = ?", id).Delete(&models.Task{}).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository taskRepository) Create(t *models.Task) error {
	db := database.GetInstance()

	record := models.Task{
		Title 				: t.Title,
		Due 				: t.Due,
		Visualized 			: false,
		CompletedAt 		: t.CompletedAt,
		RegisteredAt 		: time.Now(),
		RegisteredByUUID 	: t.RegisteredByUUID,
		CategoryUUID 		: t.Category.UUID,
		PersonUUID 			: t.Person.UUID,
		AssigneeUUID 		: t.Assignee.UUID,
	}

	err := db.Set("gorm:save_associations", false).
		Create(&record).Error

	if err != nil {
		logger.Fatal(err)
	} else {
		historics, err := CreateTaskComment(db, t.TaskHistorics, t.RegisteredByUUID, record.UUID)
		if err != nil {
			logger.Fatal(err)
		} else {
			record.TaskHistorics = historics
		}
	}

	*(t) = record
	return err
}

func (repository taskRepository) Update(t *models.Task) error {
	db := database.GetInstance()

	record := models.Task{
		Title 			: t.Title,
		Due 			: t.Due,
		CompletedAt 	: t.CompletedAt,
		CategoryUUID 	: t.Category.UUID,
		PersonUUID 		: t.Person.UUID,
		AssigneeUUID 	: t.Assignee.UUID,
	}

	err := db.Set("gorm:save_associations", false).
		Model(&models.Task{}).
		Where("uuid = ?", t.UUID).
		Updates(&record).Error

	if err != nil {
		logger.Fatal(err)
	} else {
		_, err := CreateTaskComment(db, t.TaskHistorics, t.RegisteredByUUID, t.UUID)
		if err != nil {
			logger.Fatal(err)
		}
	}

	return err
}

func (repository taskRepository) CountRows() int {
	db := database.GetInstance()

	var count int
	db.Model(&models.Task{}).Count(&count)

	return count
}

func CreateTaskComment(db *gorm.DB, historics models.TaskHistorics, registeredByUUID string, taskUUID string) (models.TaskHistorics, error) {
	for _, historic := range historics {
		historic.RegisteredByUUID 	= registeredByUUID
		historic.RegisteredAt 		= time.Now()
		historic.TaskUUID 			= taskUUID

		if err := db.Set("gorm:save_associations", false).Create(&historic).Error; err != nil {
			return historics, err
		}
	}

	return historics, nil
}
