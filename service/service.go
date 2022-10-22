package service

import (
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type Service struct {
	engine *xorm.Engine
}

func NewService(engine *xorm.Engine) Servicer {
	return &Service{engine}
}

type Servicer interface {
	NewBook() *Book
	NewAuthor() *Author
}

func (s *Service) NewBook() *Book {
	return NewBook(s.engine)
}

func (s *Service) NewAuthor() *Author {
	return NewAuhor(s.engine)
}

const (
	ServiceKey = "service_factory"
)

func ServiceFactoryModdleware(svcr Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ServiceKey, svcr)
		c.Next()
	}
}
