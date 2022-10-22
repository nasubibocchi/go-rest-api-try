package service

import (
	"xorm.io/xorm"
)

type Author struct {
	engine *xorm.Engine
}

func NewAuhor(engine *xorm.Engine) *Author {
	a := Author{
		engine: engine,
	}
	return &a
}

// // golangでは関数名の前に変数と構造体（型）を定義すると、構造体とその関数を紐づけてくれる
// func (a *Author) Create(params *model.AuthorFromJson) (*model.Author, error) {
// 	now := time.Now()
// 	author := model.Author{
// 		Name: params.Name,
// 		Books: ,
// 	}
// }

// func (a *Author) Delete(authorId int) error {

// }
