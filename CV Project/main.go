package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"Dex"
	"github.com/spf13/viper"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/vision/v1"
	cli "gopkg.in/urfave/cli.v1"
	"os"
	_ "github.com/lib/pq"
	"log"
	"runtime"
	"path/filepath"
	"api"
)

type CloudVisionConfig struct {
	APIKey string
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

func main() {
	extract()
	//initAppCli()
}

func extract() {
	initConfig()
	db := initDB()

	cv, err := initCloudVisionVision()
	if err != nil {
		panic(err)
	}

	dex, err := Dex.NewDex(db, cv)
	if err != nil {
		panic(err)
	}

	api.StartAPI(*dex)

	_, currTestPath, _, _ := runtime.Caller(0)
	testDataPath := filepath.Join(filepath.Dir(currTestPath), "test_data/images/test")
	log.Println("Test Data Path = ", testDataPath)
	//dex.ResizeAndSendFiles(testDataPath)
	//dex.PageDisplayResults()
	dex.PrintOwnerTable()
}

func initDB() *sql.DB {
	dbConfig.DBUser = "user"
	dbConfig.DBPassword ="password"
	dbConfig.Host = "192.168.99.100"
	dbConfig.DBName = "postgres"
	dbConfig.NoSSL = true
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.Host, dbConfig.DBName, dbConfig.SSLMode())
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Println("failed to open database", "err", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return db
}

func initCloudVisionVision() (*vision.Service, error) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: cloudVisionConfig.APIKey},
	}

	srv, err := vision.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve vision Client %v", err)
	}

	return srv, nil
}

func initConfig() {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config") // name of config file (without extension)

	v.AddConfigPath("/etc/smarteye") // path to look for the config file in
	v.AddConfigPath("config/")

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Println("Config file err = ", err)
		panic(err)
	}

	v.UnmarshalKey("cloud_vision", &cloudVisionConfig)
}

func (d DatabaseConfig) SSLMode() string {
	// Enable by default
	if d.NoSSL == true {
		return "disable"
	}

	return "require"
}

func initAppCli() {
	app := cli.NewApp()
	app.Name = "SmartEye Documents Extractor"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2017 RedEye Apps"
	app.HelpName = "dex - SmartEye Document Extractor Help"
	app.Usage = "dex"
	app.UsageText = "dex command [options] arguments..."

	app.Commands = []cli.Command{
		cli.Command{
			Name:      "start",
			Category:  "Start Servers",
			Usage:     "api|grpc",
			UsageText: "dex start [service] - Starts the API server",
			ArgsUsage: "dex start api",
			Subcommands: cli.Commands{
				cli.Command{
					Name:   "api",
					Action: func(c *cli.Context) { fmt.Println("start server") },
					Before: func(c *cli.Context) error {
						fmt.Fprintf(c.App.Writer, "Stuff to do before the server start\n")
						return nil
					},
					After: func(c *cli.Context) error {
						fmt.Fprintf(c.App.Writer, "Stuff to do after the server has started")
						return nil
					},
					OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
						fmt.Fprintf(c.App.Writer, "for shame\n")
						return err
					},
				},
			},
		},
		cli.Command{
			Name:      "extract",
			Category:  "Extract from file path or dir path",
			Usage:     "path/to/file/or/dir",
			UsageText: "dex extract filepath",
			ArgsUsage: "dex extract /tmp/path/to/dir",
			Action:    func(c *cli.Context) {
				fmt.Println("start server")
				//c.
			},
			Before: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "Stuff to do before extraction\n")
				return nil
			},
			After: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "Stuff to do after extraction")
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "for shame\n")
				return err
			},
		},
	}
	app.Run(os.Args)
}
