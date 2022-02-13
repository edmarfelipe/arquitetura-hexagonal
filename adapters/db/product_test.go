package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/edmarfelipe/go-hexagonal/adapters/db"
	"github.com/edmarfelipe/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	sql := `CREATE TABLE PRODUCTS (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	sql := `insert into products values ("abc", "Product", 0, "disabled")`
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
