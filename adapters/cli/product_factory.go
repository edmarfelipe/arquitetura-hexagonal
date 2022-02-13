package cli

import (
	"database/sql"

	"github.com/edmarfelipe/go-hexagonal/adapters/db"
	"github.com/edmarfelipe/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func CreateProductService() application.ProductServiceInterface {
	sqliteDb, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db.NewProductDb(sqliteDb)
	productService := application.ProductService{
		Persistence: productDbAdapter,
	}

	return &productService
}
