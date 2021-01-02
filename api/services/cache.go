package services

import (
	models "github.com/erdemkosk/go-config-service/api/models"
	plugin "github.com/erdemkosk/go-config-service/api/plugin"
)

func init() {
	plugin.InitializeRedis()
}

func getCachedConfig(key string) (models.Config, error) {
	config, err := plugin.GetValue(key)
	return config, err
}

func setCachedConfig(key string, config *models.Config) {
	plugin.SetValueWithTTL(key, config, 1000)
}

func deleteCachedConfig(key string) {
	plugin.DelKey(key)
}
