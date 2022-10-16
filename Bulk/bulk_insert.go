package main

import (
    "Bulk/models"
    "os"

    "github.com/gocarina/gocsv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/google/uuid"
    "fmt"
)

type Promotion struct {
	ID             uuid.UUID `csv:"id"`
	Price          float64   `csv:"price"`
	ExpirationDate string    `csv:"expiration_date"`
}

const (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "Valencia1#"
	DBNAME   = "postgres"
	PORT     = "5432"
	SSL_MODE = "disable"
)

func main() {

    // Open the CSV file for reading
    file, err := os.Open("promotions.csv") /* Use bulk.csv for huge amount of records */
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var promotions []Promotion
    err = gocsv.Unmarshal(file, &promotions)
    if err != nil {
        panic(err)
    }

    // Open a postgres database connection using GORM
    URI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", HOST, USER, PASSWORD, DBNAME, PORT, SSL_MODE)
    db, err := gorm.Open(postgres.Open(URI), &gorm.Config{ SkipDefaultTransaction: true, })
    if err != nil {
        panic(err)
    }

    // Drop and Create `promotions` table | Keeping Immutability
    db.Migrator().DropTable(&models.Promotion{})

    err = db.AutoMigrate(&models.Promotion{})

    if err != nil {
        panic(err)
    }

    // Save the records at once in the database by max chunk
    const max_chunk_size = 21845

    var sliceLen int
    for i := 0; i < len(promotions); i += max_chunk_size{
	    sliceLen += max_chunk_size

	    if sliceLen > len(promotions) {
		    sliceLen = len(promotions)
	    }

	    result := db.Create(promotions[i:sliceLen])
	    if result.Error != nil {
	        panic(result.Error)
	    }
    }
    fmt.Printf("%v Raws Inserted Successfully!\n", len(promotions))
}
