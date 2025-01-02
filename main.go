package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/russross/blackfriday/v2"
)

var (
	resumeSourceFile = "RESUME.md"
	headshotFile     = "headshot.jpeg"
)

func main() {
	// Read markdown resume file
	markdownBytes, err := os.ReadFile(resumeSourceFile)
	if err != nil {
		log.Fatal("Error reading resume file:", err)
	}

	// Read headshot image and convert to base64
	//headshotBytes, err := ioutil.ReadFile(headshotFile)
	//if err != nil {
	//	log.Fatal("Error reading headshot file:", err)
	//}
	//headshotBase64 := base64.StdEncoding.EncodeToString(headshotBytes)
	//imgSrc := fmt.Sprintf("data:image/jpeg;base64,%s", headshotBase64)

	// Convert Markdown to HTML
	htmlBytes := blackfriday.Run(markdownBytes)

	// Enhanced HTML template for a programmer's resume.
	htmlTemplate := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Resume</title>
    <link href="https://fonts.googleapis.com/css2?family=Fira+Code&family=Roboto:ital,wght@0,400;0,700;1,400&display=swap" rel="stylesheet">
    <style>
        body { 
            background-color: #FFF; 
            font-family: 'Roboto', sans-serif; 
            margin: 20px;
            color: #333;
        }
        h1, h2, h3 {
            color: #38454a;
        }
        p, ul, ol {
            line-height: 1.6;
        }
		img {
			float: right;
			margin: 5px;
			height: 256px;
			width: 256px;
		}
    </style>
</head>
<body>%s</body>
</html>`

	// Insert the HTML content into the body of the template.
	fullHtml := fmt.Sprintf(htmlTemplate, string(htmlBytes))
	//fullHtml := fmt.Sprintf(htmlTemplate, imgSrc, string(htmlBytes))

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set the input HTML file
	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(fullHtml))))

	// Generate the PDF document
	err = pdfg.Create()
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return
	}

	err = pdfg.WriteFile("resume.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
