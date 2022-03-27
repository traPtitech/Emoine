package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/traPtitech/Emoine/repository"
	"github.com/traPtitech/Emoine/router"
)

const (
	dbInitDirectory = "./mysql"
)

var (
	port = 80
)

func main() {
	log.SetFlags(log.Llongfile)

	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbDatabase := os.Getenv("MYSQL_DATABASE")

	// connect to db
	db := sqlx.MustConnect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		dbUser,
		dbPass,
		dbHost+":"+dbPort,
		dbDatabase,
	))
	// db connection for batch executing, allowing multi statements
	dbForBatch := sqlx.MustConnect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true",
		dbUser,
		dbPass,
		dbHost+":"+dbPort,
		dbDatabase,
	))

	// create schema
	var paths []string
	err := filepath.Walk(dbInitDirectory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, path := range paths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		dbForBatch.MustExec(string(data))
	}

	repo, err := repository.NewSqlxRepository(db)
	if err != nil {
		panic(err)
	}
	echo := router.Setup(repo)
	if err = echo.Start(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
