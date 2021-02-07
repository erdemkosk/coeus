package models

import (
	constValues "github.com/erdemkosk/go-config-service/consts"
)

type Config struct {
	Type  constValues.ConfigValueType
	Key   string
	Value interface{}
}
