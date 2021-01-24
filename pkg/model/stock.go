package model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Stock struct {
	Name         string             `yaml:"name"`
	Code         string             `yaml:"code"`
	Open         float64            `yaml:"open"`
	Close        float64            `yaml:"close"`
	Current      float64            `yaml:"current"`
	Highest      float64            `yaml:"highest"`
	Lowest       float64            `yaml:"lowest"`
	Increase     float64            `yaml:"increase"`
	TradeVolume  float64            `yaml:"tradeVolume"`
	Weibi        float64            `yaml:"weibi"`
	Last3BuySum  float64            `yaml:"last3BuySum"`
	Last3SellSum float64            `yaml:"last3SellSum"`
	CurrentTime  string             `yaml:"current"`
	IsRecord     bool               `yaml:"isRecoprd"`
	History      map[string]float64 `yaml:"history"`
}

type StockInfo struct {
	Code          string `json:"code" xorm:"pk"`
	Name          string `json:"name" xorm:"notnull"`
	Type          string `json:"type" xorm:"notnull"`
	StockExchange string `json:"stockExchange" xorm:"notnull"`
	IsST          bool   `json:"isST" xorm:"is_st"`
	IsDelist      bool   `json:"isDelist" xorm:"is_delist"`
}

func (s Stock) String() string {
	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("Error Marshal for reason: %s\n", err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error Parse for reason: %s\n", err)
	}
	return out.String()
}
