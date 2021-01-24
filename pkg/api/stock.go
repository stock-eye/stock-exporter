package api

import (
	"fmt"
	"net/http"

	"github.com/linclaus/stock-exportor/pkg/cache"

	"github.com/linclaus/stock-exportor/pkg/util"

	"github.com/gorilla/mux"
)

func (s *Server) GetStocks(w http.ResponseWriter, r *http.Request) {
	StoreStocks(cache.StockMap)
	sm := fmt.Sprint(cache.StockMap.List())
	w.Write([]byte(sm))
}

func (s Server) GetStockByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
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
		w.Write([]byte(stock.String()))
	} else {
		fmt.Printf("Code:%s doesn't exist\n", code)
		http.Error(w, fmt.Sprintf("Code:%s doesn't exist\n", code), http.StatusBadRequest)
	}

}

func (s *Server) AddStockByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
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
		stock, _ := cache.StockMap.Get(code)
		fmt.Printf("Create Dashboard:%s\n", stock.Name)
		s.Codes.Add(code)
		StoreCodes(s.Codes)
	} else {
		fmt.Printf("Code:%s doesn't exist\n", code)
		http.Error(w, fmt.Sprintf("Code:%s doesn't exist\n", code), http.StatusBadRequest)
	}
}

func (s *Server) DeleteStockByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
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
	} else {
		fmt.Printf("Code:%s doesn't exist\n", code)
		http.Error(w, fmt.Sprintf("Code:%s doesn't exist\n", code), http.StatusBadRequest)
	}

}
