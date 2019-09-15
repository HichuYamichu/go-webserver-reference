package user

import "github.com/jinzhu/gorm"

// Controller : controller with user related methods and user specyfic DAO
type Controller struct {
	dao *dao
}

// NewController : creates new Controller
func NewController(db *gorm.DB) *Controller {
	return &Controller{dao: &dao{db: db}}
}
