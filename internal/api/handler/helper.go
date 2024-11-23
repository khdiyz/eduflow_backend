package handler

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getUUIDParam(c *gin.Context, param string) (uuid.UUID, error) {
	paramValue := c.Param(param)
	if paramValue != "" {
		id, err := uuid.Parse(paramValue)
		if err != nil {
			return uuid.Nil, fmt.Errorf("invalid param %v", paramValue)
		}
		return id, nil
	}
	return uuid.Nil, errors.New("empty param value")
}
