package global

import (
	"code_memory/config"

	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG *config.Server
)
