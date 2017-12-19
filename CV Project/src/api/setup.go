package api

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"Dex"
	mid "api/middleware"
)

// Starts API
func StartAPI(dex Dex.Dex) {
	ConfigRuntime()
	StartGin(dex)
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// Sets up API files and routing functions
func StartGin(dex Dex.Dex) {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "resources/static")
	router.GET("/", mid.Index)
	router.GET("/basic", mid.Basic)
	router.GET("/drop", mid.Drop)

	router.POST("/upload", mid.UploadMultipleFiles(dex))
	router.GET("/getAll", mid.GetAll(dex))
	router.POST("/search", mid.SendResults(dex))
	router.POST("/receive", mid.ReceiveAjax)

	router.Run("localhost:8080")
}