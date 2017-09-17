package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	ModelBase
	Description string `json:"description" sql:"type:varchar(20);unique"`
	Permissions []RolePermission
}

type RolePermission struct {
	ModelBase
	RoleUUID string `json:"-" sql:"type:uuid; not null"`
	Route    string `json:"path"`
	Method   string `json:"method" sql:"type:varchar(6)"`
}

func (r Role) PopulateDefault(db *gorm.DB) {
	createdByUUID := "ce7405d8-3b78-4de7-8b58-6b32ac913701"
	administrator := []struct {
		method string
		path   string
	}{
		/* People */
		{"GET", "/api/v1/people"},
		{"GET", "/api/v1/people/*"},
		{"DELETE", "/api/v1/people/*"},
		{"POST", "/api/v1/people"},
		{"PUT", "/api/v1/people/*"},
		/* Tasks */
		{"GET", "/api/v1/tasks"},
		{"GET", "/api/v1/tasks/*"},
		{"DELETE", "/api/v1/tasks/*"},
		{"POST", "/api/v1/tasks"},
		{"PUT", "/api/v1/tasks/*"},
		/* Task Categories */
		{"GET", "/api/v1/task_categories"},
		{"GET", "/api/v1/task_categories/*"},
		{"DELETE", "/api/v1/task_categories/*"},
		{"POST", "/api/v1/task_categories"},
		{"PUT", "/api/v1/task_categories/*"},
		/* Sales */
		{"GET", "/api/v1/sales"},
		{"GET", "/api/v1/sales/*"},
		{"DELETE", "/api/v1/sales/*"},
		{"POST", "/api/v1/sales"},
		{"PUT", "/api/v1/sales/*"},
		/* Sales Product */
		{"GET", "/api/v1/sale_products"},
		/* Users */
		{"GET", "/api/v1/users/me"},
	}

	r.UUID = "899182db-4f57-4ec3-a263-3f83a4a66a6a"
	r.CreatedByUUID = createdByUUID
	r.Description = "Administrador"
	for _, permission := range administrator {
		r.Permissions = append(r.Permissions, RolePermission{
			RoleUUID: r.UUID,
			Method:   permission.method,
			Route:    permission.path,
			ModelBase: ModelBase{
				CreatedByUUID: createdByUUID,
			},
		})
	}
	db.Create(&r)
}
