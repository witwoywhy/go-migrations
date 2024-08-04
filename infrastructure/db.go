package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/viper"
	"github.com/witwoywhy/go-cores/dbs"
)

var (
	SchemaMysql database.Driver
	SchemaPg    database.Driver
	DataMysql   database.Driver
	DataPg      database.Driver
)

func InitDb() {
	initMySql()
	initPg()
}

func initMySql() {
	var config dbs.DbConfig
	if err := viper.UnmarshalKey("db.mysql", &config); err != nil {
		panic(fmt.Errorf("failed to load config db.mysql: %v", err))
	}

	db, err := sql.Open("mysql", config.ToDsn())
	if err != nil {
		panic(fmt.Errorf("failed to open mysql db: %v", err))
	}

	schema, err := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: TableSchema,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get schema driver mysql: %v", err))
	}

	data, err := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: TableData,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get data driver mysql: %v", err))
	}

	SchemaMysql = schema
	DataMysql = data
}

func initPg() {
	var config dbs.DbConfig
	if err := viper.UnmarshalKey("db.pg", &config); err != nil {
		panic(fmt.Errorf("failed to load config db.mysql: %v", err))
	}

	db, err := sql.Open("postgres", config.Dsn)
	if err != nil {
		panic(fmt.Errorf("failed to open pg db: %v", err))
	}

	schema, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: TableSchema,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get schema driver pg: %v", err))
	}

	data, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: TableData,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get data driver pg: %v", err))
	}

	SchemaPg = schema
	DataPg = data
}
