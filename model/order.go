package model

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	ClickID         string
	AdsetID         uint
	EnterpriseID    uint
	Provider        string
	ServiceID       uint
	AccountID       string // 企业账号ID
	Service         *Service
	OrderID         string `gorm:"column:order_id;type:varchar(32);primary_key"`
	Phone           string
	OrderAt         *time.Time
	OrderStatus     OrderStatus
	RegisterStatus  RegisterStatus
	RegisterAt      *time.Time
	CertID          string
	CertName        string
	ContactPhone    string
	ProvinceCode    string
	CityCode        string
	DistrictCode    string
	Address         string
	RegisterMessage string `gorm:"column:register_message;type:varchar(255)"`
	OrderMessage    string `gorm:"column:order_message;type:varchar(255)"`
	Attributes      datatypes.JSON
}

type (
	OrderStatus    int
	RegisterStatus int
)

const (
	OrderStatusSuccess    OrderStatus = iota // 成功
	OrderStatusFailed                        // 失败
	OrderStatusPending                       // 等待
	OrderStatusAlready                       // 已经存在
	OrderStatusRecharge                      // 充值中
	OrderStatusExpired                       // 过期
	OrderStatusDeregister                    // 注销
)

const (
	RegisterStatusSuccess RegisterStatus = iota
	RegisterStatusFailed
	RegisterStatusPending
	RegisterStatusAlready
)

// StringMap attributes 转换为map[string]string
func (o *Order) AttrsStringMap() map[string]string {
	m := make(map[string]string)

	if err := json.Unmarshal(o.Attributes, &m); err != nil {
		return nil
	}

	return m
}
