package services

import (
	"net/url"
	"github.com/wilsontamarozzi/panda-api/services/models"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/wilsontamarozzi/panda-api/logger"
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
	err := Con.Where("uuid = ?", taskCategoryId).Delete(&models.TaskCategory{}).Error

	if err != nil {
		logger.Fatal(err)
	}

	return err;
}

func CreateTaskCategory(taskCategory models.TaskCategory) (models.TaskCategory, error) {
	
	record := models.TaskCategory{
		Description : taskCategory.Description,
	}

	err := Con.Set("gorm:save_associations", false).
		Create(&record).Error

	if err != nil {
		logger.Fatal(err)
	}

	return record, err
}

func UpdateTaskCategory(taskCategory models.TaskCategory) (models.TaskCategory, error) {
	
	record := models.TaskCategory{
		Description : taskCategory.Description,
	}

	err := Con.Set("gorm:save_associations", false).
		Model(&models.TaskCategory{}).
		Where("uuid = ?", taskCategory.UUID).
		Updates(&record).Error

	if err != nil {
		logger.Fatal(err)
	}

	return record, err
}

func CountRowsTaskCategory() int {
	var count int
	Con.Model(&models.TaskCategory{}).Count(&count)

	return count
}