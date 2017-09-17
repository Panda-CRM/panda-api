package models

import "time"

type ModelBase struct {
	UUID          string     `json:"id,omitempty" sql:"type:uuid; primary_key; default:uuid_generate_v4(); unique"`
	CreatedByUUID string     `json:"created_by,omitempty" sql:"type:uuid"`
	CreatedAt     time.Time  `json:"created_at,omitempty" sql:"default:NOW()"`
	UpdatedAt     time.Time  `json:"update_at,omitempty" sql:"default:NOW()"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}
