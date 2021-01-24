package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linclaus/stock-exportor/pkg/cache"

	"github.com/linclaus/stock-exportor/pkg/util"
)

func (s *Server) GetStocks(c *gin.Context) {
	StoreStocks(cache.StockMap)
	sm := fmt.Sprint(cache.StockMap.List())
	c.String(http.StatusOK, sm)
}

func (s Server) GetStockByCode(c *gin.Context) {
	code := c.Param("code")
	exists := false
	switch {
	case util.StockExixts(code):
		exists = true
	case util.StockExixts("sh" + code):
		code = "sh" + code
		exists = true
	case util.StockExixts("sz" + code):
		code = "sz" + code
		exists = true
	}
	if exists {
		stock, _ := cache.StockMap.Get(code)
		c.String(http.StatusOK, stock.String())
	} else {
		fmt.Printf("Code:%s doesn't exist\n", code)
		c.JSON(http.StatusNotFound, fmt.Sprintf("Code:%s doesn't exist\n", code))
	}

}

func (s *Server) AddStockByCode(c *gin.Context) {
	code := c.Param("code")
	exists := false
	switch {
	case util.StockExixts(code):
		exists = true
	case util.StockExixts("sh" + code):
		code = "sh" + code
		exists = true
	case util.StockExixts("sz" + code):
		code = "sz" + code
		exists = true
	}
	if exists {
		util.SetByCodes([]string{code}, cache.StockMap)
		s.Codes.Add(code)
		StoreCodes(s.Codes)
		c.JSON(http.StatusOK, "")
	} else {
		fmt.Printf("Code:%s doesn't exist\n", code)
		c.JSON(http.StatusNotFound, fmt.Sprintf("Code:%s doesn't exist\n", code))
	}
}

func (s *Server) DeleteStockByCode(c *gin.Context) {
	code := c.Param("code")
	exists := false
	switch {
	case util.StockExixts(code):
		exists = true
	case util.StockExixts("sh" + code):
		code = "sh" + code
		exists = true
	case util.StockExixts("sz" + code):
		code = "sz" + code
		exists = true
	}
	if exists {
		stock, _ := cache.StockMap.Get(code)
		fmt.Printf("Remove Dashboard:%s\n", stock.Name)
		s.Codes.Remove(code)
		cache.StockMap.Remove(code)
		StoreCodes(s.Codes)
		c.JSON(http.StatusOK, "")
	} else {
		fmt.Printf("Code:%s doesn't exist\n", code)
		c.JSON(http.StatusNotFound, fmt.Sprintf("Code:%s doesn't exist\n", code))
	}

}
