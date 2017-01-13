package services

import (
	"net/url"
	"github.com/wilsontamarozzi/panda-api/services/models"
	"github.com/wilsontamarozzi/panda-api/helpers"
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
	return Con.Set("gorm:save_associations", false).
		Create(&models.TaskCategory{
			Description : taskCategory.Description,
		}).Error
}

func UpdateTaskCategory(taskCategory models.TaskCategory) error {
	return Con.Set("gorm:save_associations", false).
		Model(&models.TaskCategory{}).
		Where("uuid = ?", taskCategory.UUID).
		Updates(&models.TaskCategory{
			Description : taskCategory.Description,
		}).Error
}

func CountRowsTaskCategory() int {
	var count int
	Con.Model(&models.TaskCategory{}).Count(&count)

	return count
}