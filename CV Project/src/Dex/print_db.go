package Dex

import (
	"fmt"
	"models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func (dex *Dex) PrintDocTable(){
	docPrint, _ := models.Documents(dex.DB).All()
	for i := 0; i < len(docPrint); i++ {
		fmt.Println(docPrint[i])
	}
}

func (dex *Dex) PrintOwnerTable(){
	ownerPrint, _ := models.Owners(dex.DB).All()
	log.Println("Owners Length: ",len(ownerPrint))
	for i := 0; i < len(ownerPrint); i++ {
		fmt.Println(ownerPrint[i])
	}
}

func (dex *Dex) PrintCropTable(){
	cropPrint, _ := models.CropHints(dex.DB).All()
	for i := 0; i < len(cropPrint); i++ {
		fmt.Println(cropPrint[i])
	}
}

func (dex *Dex) PrintTextTable(){
	textPrint, _ := models.TextAnnotations(dex.DB).All()
	for i := 0; i < len(textPrint); i++ {
		fmt.Println(textPrint[i])
	}
}

func (dex *Dex) PageDisplayResults(){
	router := gin.Default()
	router.LoadHTMLFiles("/users/david/goglandprojects/smarteye-cv-go-project/web/templates/index.tmpl.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "Main website",
			"owners": models.Owners(dex.DB).OneP(),
		})
	})
	router.Run("localhost:8080")
}