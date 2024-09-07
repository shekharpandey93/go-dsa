package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/signintech/gopdf"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

type ResumeData struct {
	Name       string
	Email      string
	Phone      string
	Education  []string
	Experience []string
	Skills     []string
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_format>")
		fmt.Println("Output format should be 'pdf' or 'doc'")
		return
	}

	inputFile := os.Args[1]
	outputFormat := strings.ToLower(os.Args[2])

	content, err := extractContent(inputFile)
	if err != nil {
		log.Fatalf("Error extracting content: %v", err)
	}

	resumeData := parseResumeData(content)

	outputFile := filepath.Base(inputFile)
	outputFile = outputFile[:len(outputFile)-len(filepath.Ext(outputFile))] + "_custom_format." + outputFormat

	switch outputFormat {
	case "pdf":
		err = createCustomPDF(outputFile, resumeData)
	case "doc":
		err = createCustomDOC(outputFile, resumeData)
	default:
		log.Fatalf("Unsupported output format: %s", outputFormat)
	}

	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	fmt.Printf("Converted resume saved to %s\n", outputFile)
}

func extractContent(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".pdf":
		return extractPDFContent(filePath)
	case ".docx":
		return extractDocxContent(filePath)
	default:
		return "", fmt.Errorf("unsupported file format: %s", ext)
	}
}

func extractPDFContent(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return "", err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	var content string
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return "", err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return "", err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return "", err
		}

		content += text
	}

	return content, nil
}

func extractDocxContent(filePath string) (string, error) {
	doc, err := document.Open(filePath)
	if err != nil {
		return "", err
	}

	var content string
	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			content += run.Text()
		}
		content += "\n"
	}

	return content, nil
}

func parseResumeData(content string) ResumeData {
	lines := strings.Split(content, "\n")
	data := ResumeData{}

	for i, line := range lines {
		switch {
		case i == 0:
			data.Name = line
		case strings.Contains(line, "@"):
			data.Email = line
		case strings.Contains(line, "Phone"):
			data.Phone = line
		case strings.Contains(line, "Education"):
			data.Education = lines[i+1 : i+3]
		case strings.Contains(line, "Experience"):
			data.Experience = lines[i+1 : i+4]
		case strings.Contains(line, "Skills"):
			data.Skills = strings.Split(lines[i+1], ", ")
		}
	}

	return data
}

func createCustomPDF(outputFile string, data ResumeData) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("Arial", "arial.ttf")
	if err != nil {
		return err
	}

	pdf.SetFont("Arial", "", 16)
	pdf.Cell(nil, data.Name)
	pdf.Br(20)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(nil, data.Email)
	pdf.Br(15)
	pdf.Cell(nil, data.Phone)
	pdf.Br(25)

	sections := []struct {
		title   string
		content []string
	}{
		{"Education", data.Education},
		{"Experience", data.Experience},
		{"Skills", data.Skills},
	}

	for _, section := range sections {
		pdf.SetFont("Arial", "", 14)
		pdf.Cell(nil, section.title)
		pdf.Br(15)
		pdf.SetFont("Arial", "", 12)
		for _, item := range section.content {
			pdf.Cell(nil, item)
			pdf.Br(10)
		}
		pdf.Br(15)
	}

	return pdf.WritePdf(outputFile)
}

func createCustomDOC(outputFile string, data ResumeData) error {
	doc := document.New()

	// Add name
	para := doc.AddParagraph()
	run := para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetSize(16 * measurement.Point)
	run.AddText(data.Name)

	// Add email and phone
	para = doc.AddParagraph()
	run = para.AddRun()
	run.Properties().SetSize(12 * measurement.Point)
	run.AddText(data.Email + "\n" + data.Phone)

	sections := []struct {
		title   string
		content []string
	}{
		{"Education", data.Education},
		{"Experience", data.Experience},
		{"Skills", data.Skills},
	}

	for _, section := range sections {
		para = doc.AddParagraph()
		run = para.AddRun()
		run.AddBreak()
		run.Properties().SetBold(true)
		run.Properties().SetSize(14 * measurement.Point)
		run.AddText(section.title)

		for _, item := range section.content {
			para = doc.AddParagraph()
			run = para.AddRun()
			run.Properties().SetSize(12 * measurement.Point)
			run.AddText(item)
		}
	}

	return doc.SaveToFile(outputFile)
}
