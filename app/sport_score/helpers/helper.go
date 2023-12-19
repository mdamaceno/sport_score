package helpers

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func DecodeRawJson(c echo.Context) (map[string]interface{}, error) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		return nil, err
	}

	return json_map, nil
}
