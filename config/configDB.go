package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func criarArquivoEnv() {
	fmt.Println("Tentando criar arquivo .env")
	_, err := os.Create(".env")
	if err == nil {
		fmt.Println(".env criado com sucesso!")
	} else {
		log.Fatal("erro encontrando .env")
	}
}

func SetupDB() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		criarArquivoEnv()
	} else {
		fmt.Println("Env carregado com sucesso!")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	dbConn, err := sql.Open("mysql", connectionStr)

	if err != nil {
		log.Fatal("Could not open database")
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal("Erro ao conectar com o banco", err)
	}

	fmt.Println("Banco conectado com sucesso!")

	return dbConn
}
