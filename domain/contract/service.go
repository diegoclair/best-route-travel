package contract

import (
	"mime/multipart"

	"github.com/diegoclair/best-route-travel/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo"
)

// TravelService holds a travel service operations
type TravelService interface {
	GetBestRoute(whereFrom, whereTo string) (bestRoute entity.BestRoute, err resterrors.RestErr)
}

// UploadService holds a upload service operations
type UploadService interface {
	SaveFileForUser(c echo.Context, file *multipart.FileHeader, userUUID, fileType string) (restErr resterrors.RestErr)
}

// CommandLineService holds a cli service operations
type CommandLineService interface {
	RunCLI()
	InputNewFile(fileName string)
}
