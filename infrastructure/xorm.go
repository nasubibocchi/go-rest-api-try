package infra

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nasubibocchi/go_rest_api_try/model"
	"xorm.io/xorm"
)

func DBInit() *xorm.Engine {
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_PORT")
	localhost := os.Getenv("DOCKER_LOCAL_HOST")

	connection := user + ":root@tcp(" + localhost + ":" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=true"
	// fmt.Println(connection)

	engine, err := xorm.NewEngine("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	engine.ShowSQL(true)

	exist, err := engine.IsTableExist("book")
	if err != nil {
		log.Fatal(err)
	}

	if !exist {
		engine.CreateTables(&model.Book{}, &model.Author{})
	}

	return engine
}
