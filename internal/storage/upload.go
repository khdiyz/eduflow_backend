package storage

import (
	"context"
	"eduflow/config"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"
	"io"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type File struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type UploadMinio struct {
	minio  *minio.Client
	cfg    *config.Config
	logger *logger.Logger
}

func NewUploadMinio(minio *minio.Client, cfg *config.Config, logger *logger.Logger) *UploadMinio {
	return &UploadMinio{
		minio:  minio,
		cfg:    cfg,
		logger: logger,
	}
}

// UploadImage uploads an image file to MinIO and returns the generated file name or an error.
func (um *UploadMinio) UploadImage(image io.Reader, imageSize int64, contentType string) (file File, err error) {
	fileName := uuid.NewString()

	// Extract the file extension from the content type
	fileExtension := helper.GetFileExtension(contentType)

	// Construct the image name
	file.Name = fileName + "." + fileExtension

	// Upload the file to MinIO
	if err := um.uploadToMinio(file.Name, image, imageSize, contentType); err != nil {
		um.logger.Error("Internal Server Error: ", err)
		return file, err
	}

	// Generate the link for the uploaded image
	file.Link = helper.GenerateLink(*um.cfg, file.Name)

	return file, nil
}

func (um *UploadMinio) UploadDoc(doc io.Reader, docSize int64, contentType string) (file File, err error) {
	// Generate a new UUID for the file name
	fileName := uuid.NewString()

	// Determine the file extension based on the content type
	fileExtension := helper.GetFileExtensionForDoc(contentType)

	// Construct the document file name
	file.Name = fileName + "." + fileExtension

	// Upload the document file using the uploadToMinio method
	if err := um.uploadToMinio(file.Name, doc, docSize, contentType); err != nil {
		um.logger.Error("Internal Server Error: ", err)
		return file, err
	}

	// Generate the link for the uploaded document
	file.Link = helper.GenerateLink(*um.cfg, file.Name)

	return file, nil
}

func (um *UploadMinio) UploadExcel(excel io.Reader, excelSize int64, contentType string) (file File, err error) {
	// Generate a new UUID for the file name
	fileName := uuid.NewString()

	// Determine the file extension based on the content type
	fileExtension := helper.GetFileExtensionForExcel(contentType)

	// Construct the Excel file name
	file.Name = fileName + "." + fileExtension

	// Upload the Excel file using the uploadToMinio method
	if err := um.uploadToMinio(file.Name, excel, excelSize, contentType); err != nil {
		um.logger.Error("Internal Server Error: ", err)
		return file, err
	}

	// Generate the link for the uploaded Excel file
	file.Link = helper.GenerateLink(*um.cfg, file.Name)

	return file, nil
}

// uploadToMinio uploads the file to MinIO.
func (um *UploadMinio) uploadToMinio(fileName string, file io.Reader, fileSize int64, contentType string) error {
	_, err := um.minio.PutObject(
		context.Background(),
		um.cfg.MinioBucketName,
		fileName,
		file,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	return err
}
