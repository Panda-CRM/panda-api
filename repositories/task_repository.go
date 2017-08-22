package repositories

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/wilsontamarozzi/panda-api/models"
	"log"
	"net/url"
	"strconv"
)

type TaskRepository interface {
	List(q url.Values) models.TaskList
	Get(id string) models.Task
	Delete(id string) error
	Create(t *models.Task) error
	Update(t *models.Task) error
	CountRows() int
	CreateComments(historics *models.TaskHistorics, registeredByUUID string, taskUUID string) error
	ReportGeneral() models.ReportGeneral
	ReportByAssignees() models.ReportAssignees
	ReportByAssigneesAndCategory() models.ReportAssigneesAndCategory
	ReportByCategories() models.ReportCategories
}

type taskRepository struct{}

func NewTaskRepository() *taskRepository {
	return new(taskRepository)
}

func (repository taskRepository) List(q url.Values) models.TaskList {
	db := database.GetInstance()

	currentPage, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))
	pagination := helpers.MakePagination(repository.CountRows(), currentPage, itemPerPage)

	if q.Get("title") != "" {
		db = db.Where("title iLIKE ?", "%"+q.Get("title")+"%")
	}

	if q.Get("situation") == "open" {
		db = db.Where("completed_at IS NULL")
	}

	if q.Get("situation") == "done" {
		db = db.Where("completed_at IS NOT NULL")
	}

	userRequest := q.Get("user_request")

	switch q.Get("assigned") {
	case "author":
		db = db.Where("registered_by_uuid = ?", userRequest)
	case "all": // all task
	default:
		db = db.Where("assignee_uuid = ?", userRequest)
	}

	var typeDateQuery string

	switch q.Get("type_date") {
	case "registered":
		typeDateQuery = "registered_at"
	case "due":
		typeDateQuery = "due"
	case "completed":
		typeDateQuery = "completed_at"
	default:
		typeDateQuery = "registered_at"
	}

	startDate := q.Get("start_date")
	endDate := q.Get("end_date")

	if startDate != "" && endDate != "" {
		db = db.Where(typeDateQuery+"::DATE BETWEEN ? AND ?", startDate, endDate)
	} else if startDate != "" {
		db = db.Where(typeDateQuery+"::DATE >= ?", startDate)
	} else if endDate != "" {
		db = db.Where(typeDateQuery+"::DATE <= ?", endDate)
	}

	var tasks models.TaskList
	tasks.Meta.Pagination = pagination

	db.Preload("Category").
		Preload("RegisteredBy").
		Preload("Assignee").
		Preload("Person").
		Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("registered_at desc").
		Find(&tasks.Tasks)

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

	err := db.Set("gorm:save_associations", false).
		Create(&t).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository taskRepository) Update(t *models.Task) error {
	db := database.GetInstance()

	err := db.Set("gorm:save_associations", false).
		Model(&t).
		Omit("uuid", "code", "registered_at", "registered_by_uuid").
		Updates(&t).Error

	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository taskRepository) CountRows() int {
	db := database.GetInstance()

	var count int
	db.Model(&models.Task{}).Count(&count)

	return count
}

func (repository taskRepository) CreateComments(historics *models.TaskHistorics, taskUUID string, registeredByUUID string) error {
	db := database.GetInstance()

	for _, historic := range *historics {
		historic.TaskUUID = taskUUID
		historic.RegisteredByUUID = registeredByUUID

		err := db.Set("gorm:save_associations", false).
			Create(&historic).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository taskRepository) ReportGeneral() models.ReportGeneral {
	db := database.GetInstance()

	var result models.ReportGeneral

	db.Raw(`
		SELECT
			COUNT(*) AS total,
			COUNT(t.completed_at) AS completed,
			SUM(CASE WHEN (t.completed_at IS NULL) THEN 1 ELSE 0 END) AS not_completed,
			SUM(CASE WHEN (t.due < NOW()) AND (t.completed_at IS NULL) THEN 1 ELSE 0 END) AS overdue
		FROM
			tasks AS t;`).Scan(&result)

	return result
}

func (repository taskRepository) ReportByAssignees() models.ReportAssignees {
	db := database.GetInstance()

	var results models.ReportAssignees

	db.Raw(`
		SELECT
			t.assignee_uuid,
			a.name,
			COUNT(*) AS total,
			COUNT(t.completed_at) AS completed,
			SUM(CASE WHEN (t.completed_at IS NULL) THEN 1 ELSE 0 END) AS not_completed,
			SUM(CASE WHEN (t.due < NOW()) AND (t.completed_at IS NULL) THEN 1 ELSE 0 END) AS overdue
		FROM
			tasks AS t
		INNER JOIN
			people AS a
			ON (a.uuid = t.assignee_uuid)
		GROUP BY
			t.assignee_uuid,
			a.name
		ORDER BY
			a.name;`).Scan(&results)

	return results
}

func (repository taskRepository) ReportByAssigneesAndCategory() models.ReportAssigneesAndCategory {
	db := database.GetInstance()

	type Result struct {
		AssigneeUUID string
		Name         string
		CategoryUUID string
		Description  string
		Total        int
		Completed    int
		NotCompleted int
		Overdue      int
	}

	var results []Result

	db.Raw(`
		SELECT
			t.assignee_uuid,
			a.name,
			t.category_uuid,
			c.description,
			COUNT(*) AS total,
			COUNT(t.completed_at) AS completed,
			SUM(CASE WHEN (t.completed_at IS NULL) THEN 1 ELSE 0 END) AS not_completed,
			SUM(CASE WHEN (t.due < NOW()) AND (t.completed_at IS NULL) THEN 1 ELSE 0 END) AS overdue
		FROM
			tasks AS t
		INNER JOIN
			task_categories AS c
			ON (c.uuid = t.category_uuid)
		INNER JOIN
			people AS a
			ON (a.uuid = t.assignee_uuid)
		GROUP BY
			t.assignee_uuid,
			a.name,
			t.category_uuid,
			c.uuid
		ORDER BY
			a.name,
			c.description;`).Scan(&results)

	var assignees models.ReportAssigneesAndCategory
	var name string
	for _, assignee := range results {
		if name != assignee.Name {
			name = assignee.Name

			a := models.ReportAssigneeAndCategory{
				Name:         assignee.Name,
				AssigneeUUID: assignee.AssigneeUUID,
			}

			for _, category := range results {
				if category.Name == a.Name {
					a.Categories = append(a.Categories, models.ReportCategory{
						CategoryUUID: category.CategoryUUID,
						Description:  category.Description,
						ReportGeneral: models.ReportGeneral{
							Total:        category.Total,
							Completed:    category.Completed,
							NotCompleted: category.NotCompleted,
							Overdue:      category.Overdue,
						},
					})
				}
			}

			assignees = append(assignees, a)
		}
	}

	return assignees
}

func (repository taskRepository) ReportByCategories() models.ReportCategories {
	db := database.GetInstance()

	var results models.ReportCategories

	db.Raw(`
		SELECT
			t.category_uuid,
			c.description,
			COUNT(*) AS total,
			COUNT(t.completed_at) AS completed,
			SUM(CASE WHEN (t.completed_at IS NULL) THEN 1 ELSE 0 END) AS not_completed,
			SUM(CASE WHEN (t.due < NOW()) AND (t.completed_at IS NULL) THEN 1 ELSE 0 END) AS overdue
		FROM
			tasks AS t
		INNER JOIN
			task_categories AS c
			ON (c.uuid = t.category_uuid)
		GROUP BY
			t.category_uuid,
			c.uuid
		ORDER BY
			c.description;`).Scan(&results)

	return results
}
