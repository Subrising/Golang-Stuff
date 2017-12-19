package middleware

import (
	"encoding/base64"
	"io"
	"bytes"
	"crypto/sha256"
	"Dex"
	"log"
	"database/sql"
	"models"
	"github.com/vattle/sqlboiler/queries/qm"
)

// Sha256Hash creates a hash for a given file to use as a file database reference
func Sha256Hash(file io.Reader) (string, error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println("Failed to copy file to buffer in hashing function...")
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(buf.Bytes())

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha, nil
}

// Checks to see if a file exists in the database already by comparing file hash values
func CheckHashes(dex Dex.Dex, fileHash string) bool {
	log.Println("Check Hash fileHash = ", fileHash)
	exists, err := models.Documents(dex.DB, qm.Where("doc_hash=?", fileHash)).Exists()

	if err == sql.ErrNoRows {
		log.Println("NO ROWS")
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	if exists == true {
		return true
	}

	return false
}
