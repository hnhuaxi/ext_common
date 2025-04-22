package model

import (
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Service 服务商
type Service struct {
	gorm.Model

	Name        string
	Description string
	Provider    string

	Orders     []*Order
	Attributes datatypes.JSON
}

// StringMap attributes 转换为map[string]string
func (s *Service) AttrsStringMap() map[string]string {
	m := make(map[string]string)

	if err := json.Unmarshal(s.Attributes, &m); err != nil {
		return nil
	}

	return m
}

type ServiceAdset struct {
	SerivceID uint
	AdsetID   uint
}
