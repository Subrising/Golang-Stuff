package Dex

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"

	"models"
	"github.com/pkg/errors"
	"google.golang.org/api/vision/v1"
	"log"
//	"sync"
)

const (
	LabelDetection = "LABEL_DETECTION"
	TextDetection  = "TEXT_DETECTION"
	CropHints      = "CROP_HINTS"
)

// Takes a document and a list of file paths of files to upload to Google Cloud Vision
// The files are uploaded and stored into the SmartEye database
func (dex *Dex) ProcessWithCloudVision(document *models.Document, filePathChain... string) error {
	featureLabel := &vision.Feature{
		Type: LabelDetection,
	}

	featureText := &vision.Feature{
		Type: TextDetection,
	}

	featureCrop := &vision.Feature{
		Type: CropHints,
	}
	log.Println("File Path Chain Length =", len(filePathChain))
	var requests []*vision.AnnotateImageRequest
	for _, filePath := range filePathChain {
		data, err := ioutil.ReadFile(filePath)
		log.Println("File path =", filePath)
		if err != nil {
			log.Println("Errored in  ", "ProcessWithCloudVision")
			return errors.Wrap(err, fmt.Sprintf("Could not read file %s", filePath))
		}

		img := &vision.Image{
			Content: base64.StdEncoding.EncodeToString(data),
		}

		rqStore := &vision.AnnotateImageRequest{
			Image: img,
			Features: []*vision.Feature{
				featureLabel,
				featureText,
				featureCrop,
			},
		}

		requests = append(requests, rqStore)
	}

	log.Println("Requests Length = ", len(requests))

	batchReq := &vision.BatchAnnotateImagesRequest {
		Requests: requests,
	}

	err := dex.sendBatchRequest(batchReq, document)

	if err != nil {
		document.Status = "Failed"
	} else {
		document.Status = "Success"
	}

	document.Update(dex.DB)

	log.Println("Completed ", "ProcessWithCloudVision")

	return nil
}

func (dex *Dex) sendBatchRequest(batchReq *vision.BatchAnnotateImagesRequest, document *models.Document) error {
	i := 0
	for _, req := range batchReq.Requests {
		if req.Image != nil {
			log.Println(i)
			i++
		}
	}

	res, err := dex.CV.Images.Annotate(batchReq).Do()
	if err != nil {
		log.Println("Res Error = ", err)
		return errors.Wrap(err, "Errored performing the request")
	}

	log.Println("Res Responses = ", len(res.Responses))

	for _, imageAnnotation := range res.Responses {
		err := dex.addPageFromCloudResponse(document, imageAnnotation)
		if err != nil {
			log.Println("err = ", err)
			return err
		}
	}

	return nil
}

// A page row is created and inserted from the received Google Cloud Vision database results
// if there were no errors that occurred or were returned
// Label annotation, crop hints, and text annotation rows are added if there were no errors
// parsing the JSON results from Google Cloud Vision.
// If not errors occur, the page and documents STATUS are set to SUCCESS
// If errors occur, the page and documents STATUS are set to FAILED
func (dex *Dex) addPageFromCloudResponse(document *models.Document, imgAnnotation *vision.AnnotateImageResponse) error {
	// Creating page
	page := &models.Page{
		DocumentID: document.OwnerID,
	}
	document.AddPages(dex.TX, true, page)

	err := addLabelAnnotationsFromResponse(dex, page, imgAnnotation)
	if err != nil {
		page.Status = "Failed"
		page.Update(dex.DB)
		return errors.Wrap(err, "Error adding label annotations")
	}

	err = addTextAnnotationsFromResponse(dex, page, imgAnnotation)
	if err != nil {
		page.Status = "Failed"
		page.Update(dex.DB)
		return errors.Wrap(err, "Error adding text annotations")
	}

	err = addCropAnnotationsFromResponse(dex, page, imgAnnotation)
	if err != nil {
		page.Status = "Failed"
		page.Update(dex.DB)
		return errors.Wrap(err, "Error adding crop annotations")
	}

	page.Status = "Success"

	page.Update(dex.DB)
	document.Update(dex.DB)

	return nil
}

// Parses the Label Annotations from the returned JSON struct from Google Cloud Vision
// Parsed data is then stored in the SmartEye database
func addLabelAnnotationsFromResponse(dex *Dex, page *models.Page, imgAnnotation *vision.AnnotateImageResponse) error {
	// Add text label annotations for page
	labelAnnotations := []*models.LabelAnnotation{}
	for _, labelAnnotation := range imgAnnotation.LabelAnnotations {
		labelAnnotations = append(labelAnnotations, &models.LabelAnnotation{
			Description: labelAnnotation.Description,
			Score:       labelAnnotation.Score,
		})
	}

	err := page.AddLabelAnnotations(dex.TX, true, labelAnnotations...)
	if err != nil {
		return errors.Wrap(err, "Error adding label annotations")
	}

	return nil
}

// Parses the Text Annotations from the returned JSON struct from Google Cloud Vision
// Parsed data is then stored in the SmartEye database
func addTextAnnotationsFromResponse(dex *Dex, page *models.Page, imgAnnotation *vision.AnnotateImageResponse) error {
	textAnnotations := []*models.TextAnnotation{}
	for _, cvTextAnnotation := range imgAnnotation.TextAnnotations {
		textAnnotation := &models.TextAnnotation{
			Locale:      cvTextAnnotation.Locale,
			Description: cvTextAnnotation.Description,
		}
		textAnnotations = append(textAnnotations, textAnnotation)

		boundyPoly := NewBoundyBox(cvTextAnnotation.BoundingPoly)
		textAnnotation.Xcomax = float64(boundyPoly.MaxX)
		textAnnotation.Xcomin = float64(boundyPoly.MinX)
		textAnnotation.Ycomax = float64(boundyPoly.MaxY)
		textAnnotation.Ycomax = float64(boundyPoly.MinY)
	}

	err := page.AddTextAnnotations(dex.TX, true, textAnnotations...)
	if err != nil {
		return errors.Wrap(err, "Error adding text annotations")
	}

	return nil
}

// Parses the Crop Annotations from the returned JSON struct from Google Cloud Vision
// Parsed data is then stored in the SmartEye database
func addCropAnnotationsFromResponse(dex *Dex, page *models.Page, imgAnnotation *vision.AnnotateImageResponse) error {
	// Add crop annotation
	cropAnnotations := []*models.CropHint{}
	for _, cvCropAnnotation := range imgAnnotation.CropHintsAnnotation.CropHints {
		cropAnnotation := &models.CropHint{
			Confidence:         strconv.FormatFloat(cvCropAnnotation.Confidence, 'f', 12, 64),
			ImportanceFraction: float64(cvCropAnnotation.ImportanceFraction),
		}

		cropAnnotations = append(cropAnnotations, cropAnnotation)

		boundyPoly := NewBoundyBox(cvCropAnnotation.BoundingPoly)
		cropAnnotation.Xcomax = float64(boundyPoly.MaxX)
		cropAnnotation.Xcomin = float64(boundyPoly.MinX)
		cropAnnotation.Ycomax = float64(boundyPoly.MaxY)
		cropAnnotation.Ycomax = float64(boundyPoly.MinY)
	}

	err := page.AddCropHints(dex.TX, true, cropAnnotations...)
	if err != nil {
		return errors.Wrap(err, "Error adding crop annotations")
	}

	return nil
}