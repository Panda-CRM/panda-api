package services

import (
	"net/url"
	"panda-api/services/models"
	"panda-api/helpers"
)

func GetTaskCategories(pag helpers.Pagination, q url.Values) models.TaskCategories {

	var taskCategories models.TaskCategories

	db := Con

	if q.Get("description") != "" {
		db = db.Where("description iLIKE ?", "%" + q.Get("description") + "%")	
	}
	
	db.Limit(pag.ItemPerPage).
		Offset(pag.StartIndex).
		Order("description desc").
		Find(&taskCategories)

    return taskCategories
}

func GetTaskCategory(taskCategoryId string) models.TaskCategory {

	var taskCategory models.TaskCategory

	Con.Where("uuid = ?", taskCategoryId).
		First(&taskCategory)

	return taskCategory
}

func DeleteTaskCategory(taskCategoryId string) error {
	return Con.Where("uuid = ?", taskCategoryId).Delete(&models.TaskCategory{}).Error
}

func CreateTaskCategory(taskCategory models.TaskCategory) error {
	
	record := models.TaskCategory{
		Description : taskCategory.Description,
	}

	return Con.Set("gorm:save_associations", false).
		Table("task_categories").
		Create(&record).Error
}

func UpdateTaskCategory(taskCategory models.TaskCategory) error {
	return Con.Set("gorm:save_associations", false).
		Table("task_categories").
		Where("uuid = ?", taskCategory.UUID).
		Updates(models.TaskCategory{
			Description : taskCategory.Description,
		}).Error
}

func CountRowsTaskCategory() int {
	var count int
	Con.Table("task_categories").Count(&count)

	return count
}