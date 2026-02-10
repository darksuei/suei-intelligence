package datasource

import (
	"gorm.io/gorm"
)

type Datasource struct {
	gorm.Model

	ProjectID		uint   `gorm:"not null;index"` // <- foreign key to Project
	CreatedBy       map[string]string 		 `gorm:"type:jsonb;serializer:json;default:'{}'"`

	// Connection Details - JSONB?? Since they would be different
}