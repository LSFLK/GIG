package parsers

/*
 * PDF to text: Extract all text for each page of a pdf file.
 *
 * N.B. Only outputs character codes as seen in the content stream.  Need to account for text encoding to get readable
 * text in many cases.
 *
 * Run as: go run pdf_extract_text.go input.pdf
 */

import (
	"fmt"
	pdfcontent "github.com/unidoc/unidoc/pdf/contentstream"
	pdf "github.com/unidoc/unidoc/pdf/model"
	"os"
)

const NewPageMarker= "\n*******************\n"

/**
return the string content of a given PDF file
 */
func ParsePdf(source string) string {

	text, err := listContentStreams(source)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return text
}

func listContentStreams(inputPath string) (string, error) {
	f, err := os.Open(inputPath)
	text := ""
	if err != nil {
		return text, err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return text, err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return text, err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return text, err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return text, err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return text, err
		}

		contentStreams, err := page.GetContentStreams()
		if err != nil {
			return text, err
		}

		// If the value is an array, the effect shall be as if all of the streams in the array were concatenated,
		// in order, to form a single stream.
		pageContentStr := ""
		for _, cstream := range contentStreams {
			pageContentStr += cstream
		}

		cstreamParser := pdfcontent.NewContentStreamParser(pageContentStr)
		txt, err := cstreamParser.ExtractText()
		text = text + NewPageMarker + txt
	}

	return text, err
}
