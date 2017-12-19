package Dex

import (
	"testing"

	//"github.com/DATA-DOG/go-sqlmock"
	//"github.com/stretchr/testify/assert"
	"google.golang.org/api/vision/v1"
	"github.com/bouk/monkey"
	"reflect"
	"fmt"
	//"csv_sql/models"
	//"gopkg.in/nullbio/null.v5"
//	"log"
	"path/filepath"
	"log"
)

func TestOtherStuff(t *testing.T) {
	p := '.'

	b := filepath.Join(string(p), "SomePath")

	fmt.Println(b)
}

func TestMultithreaded(t *testing.T) {
	db := initDB()

	cv, err := initCloudVisionVision()
	if err != nil {
		panic(err)
	}

	tx, _ := db.Begin()

	monkey.Patch(vision.New, GetCVMockFn())

	var imgServiceType *vision.ImagesService
	monkey.PatchInstanceMethod(reflect.TypeOf(imgServiceType), "Annotate", GetImagesAnnotateMock())

	var imgAnnotateCall *vision.ImagesAnnotateCall
	monkey.PatchInstanceMethod(reflect.TypeOf(imgAnnotateCall), "Do", DoImagesMock())

	resp, _ := cv.Images.Annotate(&vision.BatchAnnotateImagesRequest{}).Do()
	fmt.Println(resp)
	newDex := Dex{
		DB:        db,
		TX:        tx,
		CompanyID: "DZG",
		CV:        cv,
	}

	bErr := newDex.ResizeAndSendFiles("/users/david/goglandprojects/smarteye-cv-go-project/test_data/images/test/")

	log.Println("Error check: ", bErr)
	if bErr != nil {
		panic(err)
	}
}