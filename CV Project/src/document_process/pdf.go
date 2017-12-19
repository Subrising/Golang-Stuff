package document_process

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/gographics/imagick.v2/imagick"
)

// SplitPDF splits a PDF to jpeg images to a temp directory.
// It returns the tmp directory on success and the number of files exported.
func SplitPDF(pdfPath string) (string, error) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	mw.ReadImage(pdfPath)
	mw.SetImageFormat("jpeg")

	tmpDir := os.TempDir()
	outputDir := filepath.Join(tmpDir, filepath.Base(pdfPath))

	// Files will be exported to /path/to/tmp/dir/filename.pdf/filename_0001.jpeg
	outputFile := filepath.Join(outputDir, filepath.Base(pdfPath))
	err := os.MkdirAll(outputDir, 0664)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Could not create tmp directory for pdf split %s", outputFile))
	}

	mw.WriteImages(outputFile, true)
	mw.Clear()

	return outputDir, nil
}
