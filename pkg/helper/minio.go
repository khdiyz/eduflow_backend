package helper

import (
	"eduflow/config"
	"strings"
)

const (
	docContentType  = "msword"
	docxContentType = "vnd.openxmlformats-officedocument.wordprocessingml.document"

	xlsxContentType = "vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	xlsContentType  = "vnd.ms-excel"
)

// getFileExtension extracts the file extension from the content type.
func GetFileExtension(contentType string) string {
	parts := strings.Split(contentType, "/")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

// getFileExtensionForDoc determines the file extension for the document based on the content type.
func GetFileExtensionForDoc(contentType string) string {
	switch {
	case strings.Contains(contentType, docContentType):
		return "doc"
	case strings.Contains(contentType, docxContentType):
		return "docx"
	default:
		return "docx" // Default to docx if content type is not recognized
	}
}

// getFileExtensionForExcel determines the file extension for the Excel file based on the content type.
func GetFileExtensionForExcel(contentType string) string {
	switch {
	case strings.Contains(contentType, xlsContentType):
		return "xls"
	case strings.Contains(contentType, xlsxContentType):
		return "xlsx"
	default:
		return "xlsx" // Default to xlsx if content type is not recognized
	}
}

func GenerateLink(cfg config.Config, fileName string) string {
	if fileName == "" {
		return ""
	}

	return cfg.MinioEndpoint + "/" + cfg.MinioBucketName + "/" + fileName
}
