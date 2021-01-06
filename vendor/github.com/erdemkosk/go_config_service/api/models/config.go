package models

import (
	constValues "github.com/erdemkosk/go_config_service/consts"
)

type Config struct {
	Type  constValues.ConfigValueType
	Key   string
	Value interface{}
}
