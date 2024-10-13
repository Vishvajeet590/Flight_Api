package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var DB *sql.DB

func Connect() error {
	db, err := sql.Open("postgres", dsn())
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return err
	}
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(20)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err = db.PingContext(ctx)

	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return err
	}

	DB = db
	log.Printf("Connected to DB")

	return nil
}

func dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"root",
		"root",
		"0.0.0.0",
		"5432",
		"flightDB")
	//return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", viper.GetString("database.username"), viper.GetString("database.password"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.schema"))
	//return  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "dev_test", "Dev1$4%ZStg","ls-a7721fc5b9485221abf59be59cc5487dc4b0a9d3.crtfvoumhxz2.ap-south-1.rds.amazonaws.com", "3306", "z_b")
}
