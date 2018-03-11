package projectory

import (
	"os"
	"path"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func RunProjectoryGenerator() {
	wd, wdError := os.Getwd()
	if wdError != nil {
		os.Stderr.WriteString("cannot resolve current working directory")
		os.Exit(1)
	}

	var incidents [][]string = ReadOrWriteIncidents(wd)
	var projects []Project = GenerateProjects()

	var pdf *gofpdf.Fpdf
	pdf = gofpdf.New("P", "mm", "A4", "")

	PrintProjects(projects, pdf)
	PrintIncidents(incidents, pdf)

	LOG.Print("writing the output file...")
	// now := time.Now()
	// timestamp := fmt.Sprintf("%d0%d0%dT0%d0%d0%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	timestamp := time.Now().Format("20060102T150405")
	filename := "Projectory_" + timestamp + ".pdf"
	fullPath := path.Join(wd, filename)
	pdfErr := pdf.OutputFileAndClose(fullPath)
	if pdfErr != nil {
		LOG.Fatalln("Error writing PDF file: ", pdfErr)
	}
	LOG.Println()
	LOG.Println("The output is generated:")
	LOG.Println(fullPath)
}
