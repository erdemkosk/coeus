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

func getCachedConfigs(keys [] string) ([]interface{} , error) {
	configs, err := plugin.MGet(keys)
	return configs, err
}

func setCachedConfig(key string, config *models.Config) {
	plugin.SetValueWithTTL(key, config, 500)
}

func deleteCachedConfig(key string) {
	plugin.DelKey(key)
}
