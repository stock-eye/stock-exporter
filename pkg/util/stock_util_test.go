package util

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetStockFromTushare(t *testing.T) {
	codes := GetStocksMetaDataFromTuShare()
	logrus.Print(codes)
}
