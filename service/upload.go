package service

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/diegoclair/best-route-travel/domain/contract"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo"
)

type uploadService struct {
	svc *Service
}

func newUploadService(svc *Service) contract.UploadService {
	return &uploadService{
		svc: svc,
	}
}

func (s *uploadService) SaveFileForUser(c echo.Context, file *multipart.FileHeader, userUUID, fileType string) (restErr resterrors.RestErr) {

	var path string

	// Upload the file to specific destination.
	if fileType == "prescriptions" {
		path = "upload/" + userUUID + "/prescriptions"
	}

	//create folder
	_, err := os.Stat("/" + path)
	if os.IsNotExist(err) {

		errDir := os.MkdirAll(path, os.FileMode(0770))
		if errDir != nil {
			logger.Error("SaveFileForUser: Error to create folder", errDir)
			return resterrors.NewInternalServerError("Error to save file")
		}
	}

	src, err := file.Open()
	if err != nil {
		logger.Error("SaveFileForUser: Error to open the file", err)
		return resterrors.NewInternalServerError("Error to save file")
	}
	defer src.Close()

	// Destination
	pathWithName := path + "/" + file.Filename

	dst, err := os.Create(pathWithName)
	if err != nil {
		logger.Error("SaveFileForUser: Error to create folder", err)
		return resterrors.NewInternalServerError("Error to save file")
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		logger.Error("SaveFileForUser: Error to copy folder", err)
		return resterrors.NewInternalServerError("Error to save file")
	}

	return nil
}
