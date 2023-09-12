package models

import "time"

type Product struct {
	ID                   string           `json:"id,omitempty" bson:"_id,omitempty"`
	Name                 string           `json:"name,omitempty" bson:"name,omitempty"`
	Type                 string           `json:"type,omitempty" bson:"type,omitempty"`
	Code                 string           `json:"code,omitempty" bson:"code,omitempty"`
	BasePrice            uint64           `json:"base_price,omitempty" bson:"base_price,omitempty"`
	SellPrice            uint64           `json:"sell_price,omitempty" bson:"sell_price,omitempty"`
	MinimumStockLimit    uint64           `json:"minimum_stock_limit,omitempty" bson:"minimum_stock_limit,omitempty"`
	Weight               float64          `json:"weight,omitempty" bson:"weight,omitempty"`
	Unit                 string           `json:"unit,omitempty" bson:"unit,omitempty"`
	Discount             float32          `json:"discount,omitempty" bson:"discount,omitempty"`
	RackPosition         string           `json:"rack_position,omitempty" bson:"rack_position,omitempty"`
	Description          string           `json:"description,omitempty" bson:"description,omitempty"`
	IsUseStock           bool             `json:"is_use_stock,omitempty" bson:"is_use_stock,omitempty"`
	IsShownInTransaction bool             `json:"is_shown_in_transaction,omitempty" bson:"is_shown_in_transaction,omitempty"`
	Variants             []ProductVariant `json:"variants,omitempty" bson:"variants,omitempty"`
	Units                []ProductUnit    `json:"units,omitempty" bson:"units,omitempty"`
	CreatedAt            time.Time        `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt            time.Time        `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
