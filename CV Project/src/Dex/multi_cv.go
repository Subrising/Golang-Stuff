package Dex

import (
	"os"
	"path/filepath"
	"gopkg.in/gographics/imagick.v2/imagick"
	"io/ioutil"
	"models"
	"gopkg.in/nullbio/null.v5"
	"log"
	"sync"
	"fmt"
)

const Mb = 1000000
const maxNumFilesPerReq int = 16
const maxSizePerFile int64 = 4 * Mb
const maxSizePerBatch int64 = 8 * Mb

type Batch struct {
	currentSize 	int64
	files       	[]*os.File
	filePaths	[]string
}

func (b Batch) Size() int64 {
	return b.currentSize
}

func (b Batch) Files() []*os.File {
	return b.files
}

func (b *Batch) AddFile(file *os.File) {
	// TODO: Check if we can actually add this file here (eg. check size, batch size, etc.)
	b.files = append(b.files, file)
}

func (b *Batch) AddFilePath(file string) {
	// TODO: Check if we can actually add this file here (eg. check size, batch size, etc.)
	b.filePaths = append(b.filePaths, file)
}

// RunMulti Runs multiple files
// Owner is the owner of the file AKA site
// fp is the filePath
// ResizeAndSendFiles takes a directory and goes through all the files within the directory,
// creating a set of file batches that will be uploaded to Google Cloud Vision
func (dex *Dex) ResizeAndSendFiles(dirPath string, document *models.Document) error {
	log.Println(document.DocumentID, document.OwnerID)

	log.Println(dirPath)
	var fileTotal int64

	imagick.Initialize()

	// Schedule cleanup
	defer imagick.Terminate()

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Println("Error reading file path ", err)
		return err
	}

	numberOfFiles := len(files)

	// batch of paths going to be sent
	s := []string{}

	// TODO: CHECK THIS
	batches := []Batch{}

	i := 0
	bI := 0
	for _, file := range files {
		log.Println(numberOfFiles)
		filePath := filepath.Join(dirPath, file.Name())
		log.Println(filePath)

		i++

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Println("Could not get file stat ", err)
			return err
		}

		// get the size
		size := fileInfo.Size()

		err = ResizeImage(filePath, size)
		if err != nil {
			log.Println("Resizing Issue")
			continue
		}

		if size >= maxSizePerFile {
			// TODO: fix this
			// Try to resize it again
			log.Println(filePath, " is over 4MB. Cannot upload to GCV servers.")
			ResizeImage(filePath, fileInfo.Size())
			log.Println("In here")
			// TODO: Check this
		} else if ((size+fileTotal) > maxSizePerBatch) || (i % maxNumFilesPerReq == 0) || (i == numberOfFiles) {
			// increase batch number and and reset values
			s = append(s, filePath)
			newBatch := Batch{
				filePaths: s,
			}
			log.Println("Files in Batch =", newBatch.filePaths)
			batches = append(batches, newBatch)
			bI++
			fileTotal = 0
			i = 0
			fileTotal += size
			s = nil
		} else {
			log.Println("In else")
			fileTotal += size
			s = append(s, filePath)
		}
	}

	log.Println("Total Batches = ", len(batches))

	// send batches to cloud vision
	SendBatches(dex, document, batches...)

	return nil
}

// SendBatches sends a list of files in each batch to Google Cloud Vision and then receives
// a batch response for those files which is a set of metadata that needs to be parsed
func SendBatches(dex * Dex, document *models.Document, batchChain... Batch) error {
	// send batches to cloud vision
	var err error

	log.Println("Send Batches = ", len(batchChain))
	var wg sync.WaitGroup
	txErr := dex.NewTX()
	if txErr != nil {
		log.Println("txErr = ", txErr)
	}

	for _, batch := range batchChain {
		wg.Add(1)
		go func() {
			log.Println("File Path Chain =", len(batch.filePaths))
			err = dex.ProcessWithCloudVision(document, batch.filePaths...)
			log.Println("Cherr multi_csv passed ", err)

			if err != nil {
				fmt.Println("Autobots rollout!")
				dex.TX.Rollback()
			}

			wg.Done()
		}()

		if err != nil {
			log.Println("Batch Err = ", err)
			return err
		}
		wg.Wait()
	}

	fmt.Println("Time to commit!")
	dex.TX.Commit()

	return nil
}

// ResizeImage takes a file and resizes the files inside (if required)
// before sending the files to Google Cloud Vision in order to make sure the files are within
// Google Cloud Vision's file restrictions
func ResizeImage(filePath string, fileSize int64) error {
	mw := imagick.NewMagickWand()

	err := mw.ReadImage(filePath)
	if err != nil {
		log.Println("Received an image reading error ", err)
		return err
	}

	// TODO: do ratio instead
	// Add to its own function
	// Get original size
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()
	if width <= 640 && height <= 480 && fileSize <= 4000000 {
		return nil
	}

	var heightRatio float64 = float64(height) / float64(width) * 480

	// Calculate half the size
	hWidth := uint(640)
	hHeight := uint(heightRatio)

	// Resize the image using the Lanczos filter
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		log.Println("Resizing image error ", err)
		return err
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		log.Println("Compression error ", err)
		return err
	}

	err = mw.WriteImage(filePath)
	if err != nil {
		log.Println("Image writing error ", err)
		return err
	}

	return nil
}

// Initialises the document and owner rows within the database to be used for uploaded files
func (dex *Dex) InitOwnerAndDoc(fileHash string) (*models.Document, error){
	initOwner := &models.Owner{
		OwnerName: dex.CompanyID,
	}
	err := initOwner.Insert(dex.DB)
	if err != nil{
		log.Println("Owner Insert error = ", err)
		panic(err)
	}

	initDoc := &models.Document {
		DocType: "Document",
		OwnerID: null.StringFrom(initOwner.OwnerID),
		DocHash: fileHash,
	}

	initDoc.Insert(dex.DB)
	if err != nil {
		log.Println("Document Insert error = ", err)
		panic(err)
	}

	return initDoc, nil
}

/*
LOGTABLE

logid
pageid
type
timestamp
*/

