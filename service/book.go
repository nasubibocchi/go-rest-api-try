package service

import (
	"time"

	"github.com/nasubibocchi/go_rest_api_try/model"
	"xorm.io/xorm"
)

type Book struct {
	engine *xorm.Engine
}

func NewBook(engine *xorm.Engine) *Book {
	b := Book{
		engine: engine,
	}
	return &b
}

// golangでは関数名の前に変数と構造体（型）を定義すると、構造体とその関数を紐づけてくれる
func (b *Book) Create(params *model.BookFromJson) (*model.Book, error) {
	now := time.Now()
	book := model.Book{
		Name: params.Name,
		// Author:    params.Author,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := b.engine.Table("book").Insert(&book)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *Book) GetOne(bookId int) (*model.Book, error) {
	book := model.Book{}

	_, err := b.engine.Table("book").Where("id = ?", bookId).Get()

	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *Book) GetAll() ([]*model.Book, error) {
	books := []*model.Book{}

	err := b.engine.Table("books").Find(&books)

	if err != nil {
		return nil, err
	}
	return books, err
}

func (b *Book) Update(params *model.BookFromJson, bookId int) (*model.Book, error) {
	now := time.Now()
	book := model.Book{
		Name: params.Name,
		// Author:    params.Author,
		UpdatedAt: now,
	}

	_, err := b.engine.Table("book").Where("id = ?", bookId).Update(&book)

	if err != nil {
		return nil, err
	}
	return b.GetOne(bookId)
}

func (b *Book) Delete(bookId int) error {
	book := model.Book{}

	_, err := b.engine.Table("book").Where("id = ?", bookId).Delete(&book)

	if err != nil {
		return err
	}
	return nil
}
