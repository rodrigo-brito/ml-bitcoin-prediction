package main

import (
	"encoding/json"
	"fmt"

	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/ddliu/go-httpclient"
	"github.com/sirupsen/logrus"
)

const CRIPTO_API = "https://api.coinmarketcap.com/v1/ticker/%s"

type Criptocoin struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

func (c *Criptocoin) ToArray() []string {
	timestamp := time.Now()
	return []string{
		strconv.FormatInt(timestamp.Unix(), 10),
		c.PriceUsd,
		c.PriceBtc,
		c.Two4HVolumeUsd,
		c.MarketCapUsd,
		c.AvailableSupply,
		c.TotalSupply,
		c.MaxSupply,
		c.PercentChange1H,
		c.PercentChange24H,
		c.PercentChange7D,
		c.LastUpdated,
	}
}

func GetCoin(name string) (*Criptocoin, error) {
	res, err := httpclient.Get(fmt.Sprintf(CRIPTO_API, name), nil)
	if err != nil {
		return nil, fmt.Errorf("Error on critpo request: %s", err)
	}

	coins := make([]*Criptocoin, 0)
	if err := json.NewDecoder(res.Body).Decode(&coins); err != nil {
		return nil, fmt.Errorf("Error on parse cripto response: %s", err)
	}

	if len(coins) > 0 {
		return coins[0], nil
	}
	return nil, nil
}

func saveCriptoCSV(cripto *Criptocoin, coin string) error {
	file, err := os.OpenFile(fmt.Sprintf("data/%s.csv", coin),
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write(cripto.ToArray())
}

func RunCripto() error {
	coins := []string{"bitcoin", "ripple", "ethereum", "iota", "bitcoin-cash"}
	for _, coin := range coins {
		res, err := GetCoin(coin)
		if err != nil {
			return err
		}
		err = saveCriptoCSV(res, coin)
		if err != nil {
			logrus.Errorf("Err on save cripto: %s", err)
		}
	}
	logrus.Infof("Cripto saved at %s", time.Now())
	return nil
}
