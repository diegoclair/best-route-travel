package contract

import (
	"mime/multipart"

	"github.com/diegoclair/best-route-travel/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo"
)

// PingService holds a ping service operations
type PingService interface {
}

// UserService holds a user service operations
type UserService interface {
	SignIn(user entity.User) (entity.User, resterrors.RestErr)
}

// UploadService holds a upload service operations
type UploadService interface {
	SaveFileForUser(c echo.Context, file *multipart.FileHeader, userUUID, fileType string) (restErr resterrors.RestErr)
}

type CommandLineService interface {
	RunCLI()
}
