package middleware


import (
	"time"

	"github.com/gin-gonic/gin"
	"models"
	"net/http"
	"Dex"
	"fmt"
	"os"
	"log"
	"io"
	"strconv"
	"math/rand"
	"path/filepath"
)

// Loads index upload page
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
	})
}

// Loads test index upload page
func Basic(c *gin.Context) {
	c.HTML(200, "index.tmpl.html", gin.H{
	})
}

// Loads dropzone upload page
func Drop(c *gin.Context) {
	c.HTML(200, "dropzone.html", gin.H{
	})
}

func roomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	nick := c.Query("nick")
	if len(nick) < 2 {
		nick = ""
	}
	if len(nick) > 13 {
		nick = nick[0:12] + "..."
	}
	c.HTML(200, "room_login.templ.html", gin.H{
		"roomid":    roomid,
		"nick":      nick,
		"timestamp": time.Now().Unix(),
	})

}
/*
func roomPOST(c *gin.Context) {
	roomid := c.Param("roomid")
	nick := c.Query("nick")
	message := c.PostForm("message")
	message = strings.TrimSpace(message)

	validMessage := len(message) > 1 && len(message) < 200
	validNick := len(nick) > 1 && len(nick) < 14
	if !validMessage || !validNick {
		c.JSON(400, gin.H{
			"status": "failed",
			"error":  "the message or nickname is too long",
		})
		return
	}

	post := gin.H{
		"nick":    html.EscapeString(nick),
		"message": html.EscapeString(message),
	}
	messages.Add("inbound", 1)
	room(roomid).Submit(post)
	c.JSON(200, post)
}
*/

// Test page to return changed values in index page
func GetAll(dex Dex.Dex) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "Main website",
			"owners": models.Owners(dex.DB).OneP(),
		})
	}
}

// Upload single file route function
func UploadFile(dex Dex.Dex) func (c *gin.Context) {
	return func (c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Before filename")
		filename := header.Filename
		fmt.Println(header.Filename)
		log.Println("After filename")

		out, err := os.Create("/users/david/goglandprojects/smarteye-cv-go-project/test_data/temp/" + filename + ".jpg")

		if err != nil {
			log.Fatal(err)
		}

		defer out.Close()
		_, err = io.Copy(out, file)

		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, "Uploaded...\n")
	}
}

// Receives the uploaded files and stores them in the server and database if they do
// not already exist
// The files are then uploaded to Google Cloud Vision
func UploadMultipleFiles(dex Dex.Dex) func (c *gin.Context) {
	return func (c * gin.Context) {
		err := c.Request.ParseMultipartForm(24000000)
		if err != nil {
			log.Println(err.Error(), http.StatusInternalServerError)
			return
		}

		form := c.Request.MultipartForm
		files := form.File["files[]"]
		log.Println("Number of files = ", len(files))
		folderNum := rand.Int63()
		directory := filepath.Join("/tmp/", strconv.Itoa(int(folderNum)))
		var initDoc * models.Document

		for _, file := range files {
			log.Println("File =", file.Filename)
			randNum := rand.Int63()
			randFileName := strconv.Itoa(int(randNum)) + filepath.Ext(file.Filename)
			filename := filepath.Join(directory, randFileName)

			os.Mkdir(directory, 0777)
			out, err := os.Create(filename)
			if err != nil {
				log.Println(err)
				log.Println(http.StatusConflict, "File " + file.Filename + " Already Exists...\n")
				continue
			}

			reader, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}

			fileHash, err := Sha256Hash(reader)
			if err != nil {
				log.Fatal(err)
			}

			hashCheck := CheckHashes(dex, fileHash)
			if hashCheck == true {
				log.Println(http.StatusConflict, "File " + file.Filename + " Already Exists In Database...\n")
				continue
			}

			initDoc, err = dex.InitOwnerAndDoc(fileHash)

			reader.Seek(0, 0)

			written, err := io.Copy(out, reader)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Uploaded file: " + file.Filename + "; length: " + strconv.Itoa(int(written)) + "\n")
			out.Close()
			reader.Close()
		}
		dex.ResizeAndSendFiles(directory, initDoc)
	}
}

type ResponseData struct {
	Message string `json:"message"`
}

func ReceiveAjax(c * gin.Context) {
	form := c.PostForm("ajax_post_data")
	log.Println(form)
	fmt.Println("Receive ajax post data string ", form)

	resp := ResponseData{
		Message: "This is a test response",
	}

	c.JSON(200, resp)
}

func SendResults(dex Dex.Dex) func (c *gin.Context) {
	return func (c * gin.Context) {
		form := c.PostForm("search_key")
		log.Println("In send results")
		log.Println(form)
		fmt.Println("Receive get query ", form)

		resp := ResponseData{
			Message: "This is a test get",
		}

		c.JSON(200, resp)
	}
}

/*func SendFiles(dex Dex.Dex) func (c *gin.Context) {
	return *gin.Context{}
}*/

/*
make own ginErr
form file give it  a request, parameter, extension, directory path
same with custom parameter values
*/

/*
Check sha256 of uploaded file to sha256 in document database
If not exists, then create
If exists, do not upload
Remember to set all database statuses for documents and pages
 */

/*
Use jQuery Ajax for html page templating functions
 */

/*
Store things to return to user html page via JSON structs
 */

/*
Do document type checks with % for the highest type return
 */

/*
jQuery stuff, have table updated with document and owner
if user clicks on document, they are taken to new page with document page data
 */