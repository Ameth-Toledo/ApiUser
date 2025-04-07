package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Conn_MySQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_MySQL {
	if err := godotenv.Load(); err != nil {
		log.Printf("Advertencia: No se pudo cargar el archivo .env: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatalf("Error: Faltan variables de entorno. Verifica tu .env")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	db.SetMaxOpenConns(10)

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error al verificar la conexión a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos exitosa.")
	return &Conn_MySQL{DB: db}
}

func (conn *Conn_MySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_MySQL) FetchRows(query string, values ...interface{}) (*sql.Rows, error) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}

	return rows, nil
}
