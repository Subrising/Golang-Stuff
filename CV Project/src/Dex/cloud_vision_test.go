package Dex

import (
	"testing"
	"google.golang.org/api/vision/v1"
	"log"
	"github.com/bouk/monkey"
	"net/http"
	"fmt"
	"google.golang.org/api/googleapi"
	"reflect"
	"encoding/json"
	"database/sql"
	"google.golang.org/api/googleapi/transport"
	_ "github.com/lib/pq"
	"models"
	"gopkg.in/nullbio/null.v5"
)

const (
	ERR_LABEL_ADD = "Error adding label annotations"
	ERR_TEXT_ADD = "Error adding text annotations"
	ERR_CROP_ADD = "Error adding CROP annotations"
	ERR_BATCH_REQ = "Error performing the request"
)

var Input = []byte(`
{
  "labelAnnotations": [
    {
      "mid": "/m/0h8ls87",
      "description": "Automotive Exterior",
      "score": 73
    },
    {
      "mid": "/m/01jwgf",
      "description": "Product",
      "score": 71
    },
    {
      "mid": "/m/02csf",
      "description": "Drawing",
      "score": 65
    },
    {
      "mid": "/m/02mnkq",
      "description": "Bumper",
      "score": 64
    },
    {
      "mid": "/m/0919rx",
      "description": "Line Art",
      "score": 56
    },
    {
      "mid": "/m/07yv9",
      "description": "Vehicle",
      "score": 52
    },
    {
      "mid": "/m/05h7rm",
      "description": "Coloring Book",
      "score": 51
    },
    {
      "mid": "/m/03scnj",
      "description": "Line",
      "score": 51
    }
  ],
  "textAnnotations": [
    {
      "locale": "en",
      "description": "For more step by step drawing tutorials visit us at www.drawingtutorialslol.com\n",
      "boundingPoly": {
        "vertices": [
          {
            "x": 128,
            "y": 541
          },
          {
            "x": 667,
            "y": 541
          },
          {
            "x": 667,
            "y": 559
          },
          {
            "x": 128,
            "y": 559
          }
        ]
      }
    }
  ],
  "cropHintsAnnotation": {
    "cropHints": [
      {
        "boundingPoly": {
          "vertices": [
            {},
            {
              "x": 799
            },
            {
              "x": 799,
              "y": 562
            },
            {
              "y": 562
            }
          ]
        },
        "confidence": 0.79999995,
        "importanceFraction": 1
      }
    ]
    }
}`)

func GetCVMockFn() func(client *http.Client) (*vision.Service, error) {
	return func (client *http.Client) (*vision.Service, error) {
		return &vision.Service{
			Images: &vision.ImagesService{},
		}, nil
	}
}

func GetImagesAnnotateMock() func(service *vision.ImagesService,_ *vision.BatchAnnotateImagesRequest) *vision.ImagesAnnotateCall{
	return func (service *vision.ImagesService, _ *vision.BatchAnnotateImagesRequest) *vision.ImagesAnnotateCall {
		return &vision.ImagesAnnotateCall{}
	}
}

func DoImagesMock() func(call *vision.ImagesAnnotateCall, opts ...googleapi.CallOption)  (*vision.BatchAnnotateImagesResponse, error){
	return func (call *vision.ImagesAnnotateCall, opts ...googleapi.CallOption)  (*vision.BatchAnnotateImagesResponse, error){
		imgResp := &vision.AnnotateImageResponse{}


		//buf := bytes.NewBuffer([]byte(Input))
		err := json.Unmarshal(Input, imgResp)
		if err != nil {
			panic(err.Error())
		}

		var responseList []*vision.AnnotateImageResponse
		responseList = append(responseList, imgResp)

		//check := append(check, Input)
		return &vision.BatchAnnotateImagesResponse{
			Responses: responseList,
		}, nil
	}
}

type CloudVisionConfig struct {
	Key string
}

type DatabaseConfig struct {
	Host       string
	DBName     string
	DBUser     string
	DBPassword string
	NoSSL      bool
}

var (
	cloudVisionConfig CloudVisionConfig
	dbConfig          DatabaseConfig
)

func initDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@192.168.99.100/postgres?sslmode=disable")

	if err != nil {
		log.Println("failed to open database")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return db
}

func initCloudVisionVision() (*vision.Service, error) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: cloudVisionConfig.Key},
	}

	srv, err := vision.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve vision Client %v", err)
	}

	return srv, nil
}

func (d DatabaseConfig) SSLMode() string {
	// Enable by default
	if d.NoSSL == true {
		return "disable"
	}

	return "enable"
}


func TestAddPageFromCloudResponse(t *testing.T) {
	//db, _, err := sqlmock.New()
	//assert.NoError(t, err)

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

	//csv := GetCVMockFn()
	//cloudService, _ := vision.New(&http.Client{})

	resp, _ := cv.Images.Annotate(&vision.BatchAnnotateImagesRequest{}).Do()
	fmt.Println(resp)
	newDex := Dex{
		DB: db,
		TX: tx,
		CompanyID: "DZG",
		CV: cv,
	}
	var doco = &models.Document{Status: "Hello", DocType: "sucker", OwnerID: null.StringFrom("6943fd9a-6da1-4064-9263-35157b4cb4b4"),}
	err = doco.Insert(db)
	if err != nil {
		panic(err)
	}
	// doco.Reload(db)

	chec := resp.Responses[0].LabelAnnotations
	fmt.Println(chec)
	err = newDex.addPageFromCloudResponse(doco, resp.Responses[0])
	if err != nil{
		panic(err)
	}

	/*fpc := []string{"/users/david/goglandprojects/smarteye-cv-go-project/images/test/jesus000.jpg", "/users/david/goglandprojects/smarteye-cv-go-project/images/test/jesus001.jpg"}

	cherr := newDex.ProcessWithCloudVision(doco, fpc...)

	select {
	case cerr := <- cherr:
		if cerr.Error() == ERR_LABEL_ADD {
			log.Println(ERR_LABEL_ADD)
			t.Fail()
		}
		if cerr.Error() == ERR_TEXT_ADD {
			log.Println(ERR_TEXT_ADD)
			t.Fail()
		}
		if cerr.Error() == ERR_CROP_ADD {
			log.Println(ERR_CROP_ADD)
			t.Fail()
		}
		if cerr.Error() == ERR_BATCH_REQ {
			log.Println(ERR_BATCH_REQ)
			t.Fail()
		}
	}
	*/

	//assert.NoError(t, )

	//jsResp := &http.Response{}

}



/* create bytes buffer, buf.write

write

someresp http.response

responsebodu = response
unmarshal the response

 */

