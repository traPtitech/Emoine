package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/FujishigeTemma/Emoine/router"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbInitDirectory = "./mysql/init"
)

var (
	port = 80
)

func main() {
	// connect to db
	db := sqlx.MustConnect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		"root",
		"password",
		"mysql",
		"emoine",
	))
	// db connection for batch executing, allowing multi statements
	dbForBatch := sqlx.MustConnect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true",
		"root",
		"password",
		"mysql",
		"emoine",
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
	echo.Start(fmt.Sprintf(":%d", port))
}