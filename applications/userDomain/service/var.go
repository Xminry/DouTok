package service

import (
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/query"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Do query.IUserDo
var User = query.UserStruct
