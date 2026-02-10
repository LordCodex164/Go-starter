package main

import (
	"fmt"
	"strings"
)

var salesItems = map[string]float64{
	"cloth": 79.09,
	"toothpaste": 87.0,
	"lotion": 100.3,
}

func calculateSalesPrice (itemCode string) (float64, bool) {
	itemPrice, found := salesItems[itemCode]; if !found {
		if strings.HasSuffix(itemCode, "inventory") {
			formattedCode := strings.TrimSuffix(itemCode, "inventory")
			itemPrice, found := salesItems[formattedCode]; if found {
				newPrice := itemPrice * 0.21
				return newPrice, true
			}
			return 0.0, false
		}
		return 0.0, false
	}
	return itemPrice, true
}

func main() {
	orderItems := map[string]float64{
		" ": 79.09,
		"toothpaste": 87.0,
		"lotioninventory": 100.3,
	}

	subTotal := 0

	for item, _ := range(orderItems) {
		price, _ := calculateSalesPrice(item)
		fmt.Println("newprice", price)
		subTotal += int(price)
	}

	fmt.Println("subtotal", subTotal)

	fmt.Println("orderItems", orderItems)
}
