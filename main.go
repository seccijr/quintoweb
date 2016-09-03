package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/facebookgo/inject"
	"github.com/seccijr/quintoweb/environment"
	"github.com/seccijr/quintoweb/handler"
	"github.com/seccijr/quintoweb/repository"
	"github.com/seccijr/quintoweb/service"
	"golang.org/x/text/language"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

func main() {
	var g inject.Graph
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	root := environment.Root()
	i18n := service.NewJsonI18n()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Printf("Could connect to database: %+v\n", err)
		return
	}
	err = i18n.ParseTranslationRoot(filepath.Join(root, "resource/translation"))
	if err != nil {
		fmt.Printf("Could not install translations: %+v\n", err)
		return
	}
	g.Provide(
		&inject.Object{Value: &db},
		&inject.Object{Value: &i18n},
	)
	r := handler.Router()
	http.ListenAndServe(":"+port, r)
}
