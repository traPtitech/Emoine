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

	user, ok := os.LookupEnv("MYSQL_USERNAME")
	if !ok {
		user = "root"
	}
	pass, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		pass = "password"
	}
	host, ok := os.LookupEnv("MYSQL_HOSTNAME")
	if !ok {
		host = "mysql"
	}
	dbPort, ok := os.LookupEnv("MYSQL_PORT")
	if !ok {
		dbPort = "3306"
	}

	dbname, ok := os.LookupEnv("MYSQL_DATABASE")
	if !ok {
		dbname = "emoine"
	}

	// connect to db
	db := sqlx.MustConnect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		pass,
		host,
		dbPort,
		dbname,
	))
	// db connection for batch executing, allowing multi statements
	dbForBatch := sqlx.MustConnect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true",
		user,
		pass,
		host,
		dbPort,
		dbname,
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
