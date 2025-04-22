package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type adsetCache struct {
	redisCli *redis.Client
}

type Adset struct {
	ID uint `json:"id,omitempty"`

	Name            string `json:"name,omitempty"`
	Prompt          string `json:"prompt,omitempty"`
	Provider        string `json:"provider,omitempty"`
	CreatorID       uint   `json:"creator_id,omitempty"`
	AccountID       uint   `json:"account_id,omitempty"`
	SourceAccountID string `json:"source_account_id,omitempty"`
	EnterpriseID    uint   `json:"enterpriseID,omitempty"`
	UserActionSetID uint   `json:"user_action_set_id,omitempty"`
	ProductSourceID string `json:"product_source_id,omitempty"`
}

type Cache interface {
	Get(key string) (*Adset, error)
}

const CachePrefixKey = "platform$$adset:"

// NewAdsetCache ...
func NewAdsetCache(redisCli *redis.Client) *adsetCache {
	return &adsetCache{
		redisCli: redisCli,
	}
}

// Get
func (c *adsetCache) Get(key string) (*Adset, error) {
	var ctx = context.Background()

	r, err := c.redisCli.Get(ctx, CachePrefixKey+key).Result()
	if err != nil {
		return nil, err
	}
	var adset Adset
	if err := json.Unmarshal([]byte(r), &adset); err != nil {
		return nil, err
	}

	return &adset, nil
}
