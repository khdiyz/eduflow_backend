package helper

import (
	"eduflow/internal/constants"
	"eduflow/internal/model"
	"eduflow/pkg/logger"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPagination(c *gin.Context) (pagination model.Pagination, err error) {
	page, err := getPageQuery(c)
	if err != nil {
		logger.GetLogger().Error(err)
		return pagination, err
	}
	pageSize, err := getPageSizeQuery(c)
	if err != nil {
		logger.GetLogger().Error(err)
		return pagination, err
	}
	offset, limit := calculatePagination(page, pageSize)
	pagination.Limit = limit
	pagination.Offset = offset
	pagination.Page = page
	pagination.PageSize = pageSize
	return pagination, nil
}

func getPageQuery(c *gin.Context) (offset int64, err error) {
	offsetStr := c.DefaultQuery("page", constants.DefaultPage)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error while parsing query: %v", err.Error())
	}
	if offset < 0 {
		return 0, fmt.Errorf("page should be unsigned")
	}
	return offset, nil
}

func getPageSizeQuery(c *gin.Context) (limit int64, err error) {
	limitStr := c.DefaultQuery("page_size", constants.DefaultPageSize)
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error while parsing query: %v", err.Error())
	}
	if limit < 0 {
		return 0, fmt.Errorf("page_size should be unsigned")
	}
	return limit, nil
}

func calculatePagination(page, pageSize int64) (offset, limit int64) {
	if page < 0 {
		page = 1
	}
	offset = (page - 1) * pageSize
	limit = pageSize
	return offset, limit
}

// Get Query Functions

func GetInt64Query(c *gin.Context, queryName string) (int64, error) {
	queryData := c.DefaultQuery(queryName, "0")

	paramValue, err := strconv.ParseInt(queryData, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid query: %s", queryData)
	}

	return paramValue, nil
}

func GetInt64Param(c *gin.Context, paramName string) (int64, error) {
	paramData := c.Param(paramName)

	if paramData == "" {
		return 0, fmt.Errorf("param %s is required", paramName)
	}

	paramValue, err := strconv.ParseInt(paramData, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid param: %s", paramData)
	}

	return paramValue, nil
}

func GetUserId(c *gin.Context) (int64, error) {
	id, ok := c.Get(constants.UserCtx)
	if !ok {
		return 0, constants.ErrInvalidUserId
	}
	userId, ok := id.(int64)
	if !ok {
		return 0, constants.ErrInvalidUserId
	}

	return userId, nil
}

func GetUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get(constants.RoleCtx)
	if !ok {
		return "", constants.ErrInvalidUserRole
	}
	userRole, ok := role.(string)
	if !ok {
		return "", constants.ErrInvalidUserId
	}

	return userRole, nil
}
