package plaeve_auth

import (
	"github.com/jinzhu/gorm"
	uid "github.com/satori/go.uuid"
)

// BeforeCreate generates a uuid for the user before adding them to the db
func (model *Auth) BeforeCreate(scope *gorm.Scope) error {
	uuid := uid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}