package main

import (
	"regexp"

	"strconv"
	"strings"

	"encoding/csv"
	"os"

	"time"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

var regexNumber = regexp.MustCompile(`[ a-zA-Z\$\.]`)

type Market struct {
	Timestamp time.Time
	Dolar     float64
	Euro      float64
	Nasdaq    float64
	Bovespa   float64
	Bitcoin   float64
}

func (m *Market) ToArray() []string {
	return []string{
		strconv.FormatInt(m.Timestamp.Unix(), 10),
		strconv.FormatFloat(m.Dolar, 'f', 2, 64),
		strconv.FormatFloat(m.Euro, 'f', 2, 64),
		strconv.FormatFloat(m.Nasdaq, 'f', 2, 64),
		strconv.FormatFloat(m.Bovespa, 'f', 2, 64),
		strconv.FormatFloat(m.Bitcoin, 'f', 2, 64),
	}
}

func toMoney(value string) float64 {
	value = regexNumber.ReplaceAllString(value, "")
	number, err := strconv.ParseFloat(strings.Replace(value, ",", ".", 1), 64)
	if err != nil {
		logrus.Errorf("Error on convert %s: %s", value, err)
	}
	return number
}

func saveMarketCSV(market Market) error {
	file, err := os.OpenFile("data/market.csv",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	logrus.Infof("Market saved at %s", time.Now())
	return writer.Write(market.ToArray())
}

func RunMarket() error {
	market := Market{}
	c := colly.NewCollector()
	c.OnHTML(".li-ibovespa .last", func(e *colly.HTMLElement) {
		market.Bovespa = toMoney(e.Text)
	})

	c.OnHTML(".li-dolar .last", func(e *colly.HTMLElement) {
		market.Dolar = toMoney(e.Text)
	})

	c.OnHTML(".li-euro .last", func(e *colly.HTMLElement) {
		market.Euro = toMoney(e.Text)
	})

	c.OnHTML(".li-nasdaq .last", func(e *colly.HTMLElement) {
		market.Nasdaq = toMoney(e.Text)
	})

	c.OnHTML(".last-child .last", func(e *colly.HTMLElement) {
		market.Bitcoin = toMoney(e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		logrus.WithField("status", r.StatusCode).
			WithField("response", string(r.Body)).
			Error(err)
	})

	c.Visit("http://www.infomoney.com.br/mercados/cambio")

	market.Timestamp = time.Now()

	return saveMarketCSV(market)
}
