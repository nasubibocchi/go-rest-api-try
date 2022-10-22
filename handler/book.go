package handler

// routesで定義され呼び出されるCRUD処理
import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nasubibocchi/go_rest_api_try/model"
	"github.com/nasubibocchi/go_rest_api_try/service"
)

func CreateBook(c *gin.Context) {
	params := model.BookFromJson{}

	// 複数回呼び出せない : ShouldBindWith
	// Bind : JSONをparams構造体に格納
	err := c.ShouldBindWith(&params, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// service層のインスタンスを作成
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	book := service.NewBook()

	// service層のCreateメソッドを呼び出し
	payload, err := book.Create(&params)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func GetAll(c *gin.Context) {
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	book := service.NewBook()

	payload, err := book.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func GetById(c *gin.Context) {
	bookId, err := strconv.Atoi((c.Param("id")))

	if err != nil {
		log.Fatal(err)
	}

	service := c.MustGet(service.ServiceKey).(service.Servicer)
	book := service.NewBook()

	payload, err := book.GetOne(bookId)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Update(c *gin.Context) {
	bookId, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		log.Fatal(err)
	}

	params := model.BookFromJson{}
	err = c.ShouldBindBodyWith(&params, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	service := c.MustGet(service.ServiceKey).(service.Servicer)
	book := service.NewBook()

	payload, err := book.Update(&params, bookId)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Delete(c *gin.Context) {
	bookId, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		log.Fatal(err)
	}

	service := c.MustGet(service.ServiceKey).(service.Servicer)
	book := service.NewBook()

	err = book.Delete(bookId)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, "削除完了")
}
