package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/panda8z/istone/src/pkg/models"
)

type Index interface {
	Generate() Index
	Bind(ctx *gin.Context) error
	GetPageIndex() int
	GetPageSize() int
	GetNeedSearch() interface{}
}

type Control interface {
	Generate() Control
	Bind(ctx *gin.Context) error
	GenerateM() (models.ActiveRecord, error)
	GetId() interface{}
}
