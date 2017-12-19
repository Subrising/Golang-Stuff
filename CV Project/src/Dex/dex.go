package Dex

import (
	"database/sql"
	"google.golang.org/api/vision/v1"
	"log"
)

// Dex object struct
type Dex struct {
	DB        *sql.DB
	TX        *sql.Tx
	CompanyID string
	CV        *vision.Service
}

// Creates a new Dex object using a given database service and Google Cloud Vision API service
func NewDex(db *sql.DB, cvService *vision.Service) (*Dex, error) {
	/*tx, err := db.Begin()
	if err != nil {
		log.Println("Error setting tx")
		return nil, err
	}*/
/*
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
*/
	return &Dex{
		DB: db,
		CV: cvService,
	}, nil
}

// Sets up a new transaction for the database
func (dex *Dex) NewTX() error {
	tx, err := dex.DB.Begin()
	if err != nil {
		log.Println("Error setting tx")
		return err
	}
	dex.TX = tx
	return nil
}

// Closes a current transaction for the database
func (dex *Dex) CloseTX() error {
	err := dex.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

/*
	Check initial owner in middleware
 */