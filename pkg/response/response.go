package response

import (
	"eduflow/internal/model"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SuccessResponse(c *gin.Context, status Status, data interface{}, pagination *model.Pagination) {
	c.JSON(status.Code, model.BaseResponse{
		Success:     true,
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
		Pagination:  pagination,
	})
}

func ErrorResponse(c *gin.Context, status Status, err error) {
	c.JSON(status.Code, model.BaseResponse{
		Success:      false,
		Status:       status.Status,
		Description:  status.Description,
		ErrorMessage: err.Error(),
	})
}

func AbortResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(Aborted.Code, model.BaseResponse{
		Success:      false,
		Status:       Aborted.Status,
		Description:  Aborted.Description,
		ErrorMessage: message,
	})
}

// CONVERT SERVICE ERROR TO HANDLER ERROR

func FromError(c *gin.Context, serviceError error) {
	st, _ := status.FromError(serviceError)
	err := st.Message()

	switch st.Code() {
	case codes.Internal:
		ErrorResponse(c, Internal, errors.New(err))
	case codes.NotFound:
		ErrorResponse(c, NotFound, errors.New(err))
	case codes.InvalidArgument:
		ErrorResponse(c, BadRequest, errors.New(err))
	case codes.Unavailable:
		ErrorResponse(c, Unavailable, errors.New(err))
	case codes.AlreadyExists:
		ErrorResponse(c, AlreadyExists, errors.New(err))
	case codes.Unauthenticated:
		ErrorResponse(c, Unauthorized, errors.New(err))
	}
}

// SERVICE ERROR RESPONSE

var errorMapping = map[string]struct {
	code codes.Code
	msg  string
}{
	"no rows in result set":                          {codes.NotFound, "data is empty"},
	"duplicate key value violates unique constraint": {codes.AlreadyExists, "variable value is already exists"},
	"violates foreign key constraint":                {codes.InvalidArgument, "foreign key violation"},
	"no rows affected":                               {codes.NotFound, "variable value is not exists"},
}

func ServiceError(err error, code codes.Code) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()

	for substr, mapping := range errorMapping {
		if strings.Contains(errMsg, substr) {
			return status.Error(mapping.code, mapping.msg)
		}
	}

	if code != codes.OK {
		return status.Error(code, errMsg)
	}

	return status.Error(codes.Unknown, errMsg)
}
