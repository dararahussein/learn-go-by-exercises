package application

import "fmt"

type PriceLookup interface {
	Price(sku string) (int, error)
}

// CartTotal is the end-of-set application: it consumes an interface, handles
// errors from another implementation, and wraps the failing SKU as context.
func CartTotal(prices PriceLookup, skus []string) (int, error) {
	return 0, fmt.Errorf("TODO: CartTotal for %d items", len(skus))
}
