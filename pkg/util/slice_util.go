package util

import "github.com/linclaus/stock-exportor/pkg/model"

func SliceDelete(slice []string, item string) []string {
	newSlice := make([]string, 0)
	for _, v := range slice {
		if item != v {
			newSlice = append(newSlice, v)
		}
	}
	slice = newSlice
	return newSlice
}

func SplitIntSlice(slice []int, max int) [][]int {
	var sliceLists [][]int
	start := 0
	i := 1
	end := max
	for start < len(slice) {
		i = i + 1
		end = If(end < len(slice), end, len(slice)).(int)
		sliceList := slice[start:end]
		sliceLists = append(sliceLists, sliceList)
		start = end
		end = i * max
	}
	return sliceLists
}

func SplitStringSlice(slice []string, max int) [][]string {
	var sliceLists [][]string
	start := 0
	i := 1
	end := max
	for start < len(slice) {
		i = i + 1
		end = If(end < len(slice), end, len(slice)).(int)
		sliceList := slice[start:end]
		sliceLists = append(sliceLists, sliceList)
		start = end
		end = i * max
	}
	return sliceLists
}

func SplitStockSlice(slice []model.Stock, max int) [][]model.Stock {
	var sliceLists [][]model.Stock
	start := 0
	i := 1
	end := max
	for start < len(slice) {
		i = i + 1
		end = If(end < len(slice), end, len(slice)).(int)
		sliceList := slice[start:end]
		sliceLists = append(sliceLists, sliceList)
		start = end
		end = i * max
	}
	return sliceLists
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
