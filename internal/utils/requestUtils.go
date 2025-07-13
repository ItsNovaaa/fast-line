package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUUIDFromPath extracts a URL path parameter and parses it into a UUID.
// It returns an error if the parameter is missing or is not a valid UUID.
func GetUUIDFromPath(c *gin.Context, paramName string) (uuid.UUID, error) {
	// 1. Get the string value of the parameter from the URL
	idStr := c.Param(paramName)
	if idStr == "" {
		return uuid.Nil, errors.New("missing path parameter: " + paramName)
	}

	// 2. Parse the string into a UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		// The string was not a valid UUID
		return uuid.Nil, errors.New("invalid UUID format for parameter: " + paramName)
	}

	// 3. Return the valid UUID
	return id, nil
}