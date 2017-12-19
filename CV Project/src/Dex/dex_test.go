package Dex

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/vision/v1"
)

// TODO: Check tx
func TestNewDex(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)

	var csv *vision.Service
	_, err = NewDex(db, csv)
	assert.NoError(t, err)
	//assert.NotNil(t, myDex.TX)
}
