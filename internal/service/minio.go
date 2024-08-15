package service

import (
	"eduflow/internal/storage"
	"eduflow/pkg/logger"
	"eduflow/pkg/response"
	"io"

	"google.golang.org/grpc/codes"
)

type MinioService struct {
	strg   storage.Storage
	logger logger.Logger
}

func NewMinioService(strg storage.Storage, logger logger.Logger) *MinioService {
	return &MinioService{strg: strg, logger: logger}
}

func (m *MinioService) UploadImage(image io.Reader, imageSize int64, contextType string) (storage.File, error) {
	file, err := m.strg.UploadStorage.UploadImage(image, imageSize, contextType)
	if err != nil {
		return storage.File{}, response.ServiceError(err, codes.Internal)
	}

	return file, nil
}

func (m *MinioService) UploadDoc(doc io.Reader, docSize int64, contextType string) (storage.File, error) {
	file, err := m.strg.UploadStorage.UploadDoc(doc, docSize, contextType)
	if err != nil {
		return storage.File{}, response.ServiceError(err, codes.Internal)
	}

	return file, nil
}

func (m *MinioService) UploadExcel(doc io.Reader, docSize int64, contextType string) (storage.File, error) {
	file, err := m.strg.UploadStorage.UploadExcel(doc, docSize, contextType)
	if err != nil {
		return storage.File{}, response.ServiceError(err, codes.Internal)
	}

	return file, nil
}
