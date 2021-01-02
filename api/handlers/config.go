package handlers

import (
	"net/url"
	"time"

	constValues "github.com/erdemkosk/go-config-service/api/consts"
	models "github.com/erdemkosk/go-config-service/api/models"
	services "github.com/erdemkosk/go-config-service/api/services"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

const (
	ERROR_CONNOT_FIND      = "Cannot find expected config."
	ERROR_INVALID_TYPE     = "Invalid config type."
	ERROR_ALREADY_CREADTED = "Config already created."
)

// GetConfig is a function to get a book by ID
// @Summary Get config with using key
// @Description Get config with using key
// @Tags config
// @Accept json
// @Produce json
// @Param key path string true "Config key value"
// @Param types query string true "type" Enums(Array, Boolean, Object, String, Number)
// @Success 200 {object} models.Config{}
// @Security Authorization
// @Router /api/config/{key} [get]
func GetConfig(c *fiber.Ctx) error {
	types := c.Query("types")
	key, _ := url.QueryUnescape(c.Params("key"))

	isValidTye := constValues.ConfigValueType(types).String()

	if isValidTye == "" {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   types,
		})
	}

	conf, err := services.GetConfig(types, key)

	if err != nil || conf.Key == "" {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   ERROR_CONNOT_FIND,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": false,
		"error":   conf,
	})
}

// CreateConfig is a function to get a book by ID
// @Summary Create config with using key
// @Description Create config with using key
// @Tags config
// @Accept json
// @Produce json
// @Param data body models.ConfigInput true "Config Value"
// @Success 200 {object} models.Config{}
// @Security Authorization
// @Router /api/config/ [post]
func CreateConfig(c *fiber.Ctx) error {
	userConfig := new(models.ConfigInput)

	if errParsing := c.BodyParser(userConfig); errParsing != nil {
		return errParsing
	}

	configType := constValues.ConfigValueType(userConfig.Type).String()

	if configType == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   ERROR_INVALID_TYPE,
		})
	}

	configAlreadyCreated, err := services.GetConfig(userConfig.Type, userConfig.Key)

	if configAlreadyCreated.Key != "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   ERROR_ALREADY_CREADTED,
		})
	}

	config := models.Config{Id: primitive.NewObjectID(), Key: userConfig.Key, Type: userConfig.Type, Value: userConfig.Value, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	result, err := services.CreateConfig(config)

	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    result,
	})
}

// UpdateConfig is a function to get a book by ID
// @Summary Update config with using key
// @Description Update config with using key
// @Tags config
// @Accept json
// @Produce json
// @Param key path string true "Config key value"
// @Param data body models.ConfigInput true "Config Value"
// @Success 200 {object} models.Config{}
// @Security Authorization
// @Router /api/config/{key} [put]
func UpdateConfig(c *fiber.Ctx) error {
	userConfig := new(models.ConfigInput)
	key, _ := url.QueryUnescape(c.Params("key"))

	if errParsing := c.BodyParser(userConfig); errParsing != nil {
		return errParsing
	}

	configType := constValues.ConfigValueType(userConfig.Type).String()

	if configType == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   ERROR_INVALID_TYPE,
		})
	}

	config := models.Config{Id: primitive.NewObjectID(), Key: userConfig.Key, Type: userConfig.Type, Value: userConfig.Value, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result, err := services.UpdateConfig(key, config)

	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    result,
	})
}

// DeleteConfig is a function to get a book by ID
// @Summary Delete config with using key
// @Description Delete config with using key
// @Tags config
// @Accept json
// @Produce json
// @Param key path string true "Config key value"
// @Success 200 {object} models.Config{}
// @Security Authorization
// @Router /api/config/{key} [delete]
func DeleteConfig(c *fiber.Ctx) error {

	key, _ := url.QueryUnescape(c.Params("key"))

	result, err := services.DeleteConfig(key)

	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": result,
	})
}
