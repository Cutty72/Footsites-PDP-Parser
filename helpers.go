package PDP


import (
	"encoding/json"
	"errors"
	//"fmt"
)

var (
	ReadBodyErr = errors.New("error reading response body")
	UnmarshalErr = errors.New("error unmarshalling response body")
)

func GetSizes(body []byte, sku, status string) (map[string]string, error) {
	var stock StockJSON

	err := json.Unmarshal(body, &stock)
	if err != nil {
		return nil, UnmarshalErr
	}

	a := ParseSizes(sku, status, stock)
	return a, nil
}

func ParseSizes(sku, status string, stock StockJSON) map[string]string {
	var productCode string

	for _, skuID := range stock.VariantAttributes {
		if skuID.Sku == sku {
			productCode = skuID.Code
		}
	}

	var StockInfo []Stock

	for _, value := range stock.SellableUnits {
		if value.Attributes[1].ID == productCode {
			StockInfo = append(StockInfo, Stock{
				StockLevel: value.StockLevelStatus,
				ID: value.Attributes[0].ID,
				Size: value.Attributes[0].Value,
			})
		}
	}

	switch status {
		case "inStock":
			// Returns only inStock PIDs
			return CreateInStockMap(StockInfo, "inStock")
		case "all":
			// Returns all PIDs, inStock or not.
			return CreateAllMap(StockInfo)
	}
	return nil
}

func CreateInStockMap(StockInfo []Stock, status string) map[string]string {
	sizeMap := map[string]string{}
	for _, item := range StockInfo {
		if item.StockLevel == status {
			sizeMap[item.Size] = item.ID
		}
	}
	return sizeMap
}

func CreateAllMap(StockInfo []Stock) map[string]string {
	sizeMap := map[string]string{}
	for _, item := range StockInfo {
		sizeMap[item.Size] = item.ID
	}
	return sizeMap
}