package set03

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type fakePrices map[string]int

func (p fakePrices) Price(sku string) (int, error) {
	price, ok := p[sku]
	if !ok {
		return 0, fmt.Errorf("price %q: %w", sku, ErrNotFound)
	}
	return price, nil
}

func TestCartTotal(t *testing.T) {
	prices := fakePrices{"book": 2500, "pen": 200}
	if got, err := CartTotal(prices, []string{"book", "pen"}); got != 2700 || err != nil {
		t.Errorf("CartTotal = (%d, %v); want (2700, nil)", got, err)
	}
	_, err := CartTotal(prices, []string{"book", "missing"})
	if !errors.Is(err, ErrNotFound) || !strings.Contains(err.Error(), "missing") {
		t.Errorf("error %v should preserve ErrNotFound and mention missing", err)
	}
}
