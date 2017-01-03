package services

import (
	"panda-api/services/models"
	"panda-api/helpers"
)

func GetTaskCategories(pag helpers.Pagination, number string, title string) models.TaskCategories {

	var taskCategories models.TaskCategories

	db := Con

	if number != "" {
		db = db.Where("number = ?", number)
	}

	if title != "" {
		db = db.Where("title iLIKE ?", "%" + title + "%")	
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
	return Con.Set("gorm:save_associations", false).Create(&taskCategory).Error
}

func UpdateTaskCategory(taskCategory models.TaskCategory) error {
	return Con.Set("gorm:save_associations", false).Save(&taskCategory).Error
}

func CountRowsTaskCategory() int {
	var count int
	Con.Table("task_categories").Count(&count)

	return count
}