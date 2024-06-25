package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	apiHandlers "github.com/felipedavid/ihml/api/handlers"

	_ "modernc.org/sqlite"
)

func main() {
	addr := os.Getenv("ADDR")
	dbPath := os.Getenv("DB_PATH")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := http.Server{
		Addr:    addr,
		Handler: apiHandlers.DefineRoutes(),
	}

	slog.Info("Starting server", "addr", addr)
	err = s.ListenAndServe()
	panic(err)
}
