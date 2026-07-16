package application

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	errorexercises "learn-go-by-exercises/set03/01_errors"
)

type fakePrices map[string]int

func (p fakePrices) Price(sku string) (int, error) {
	price, ok := p[sku]
	if !ok {
		return 0, fmt.Errorf("price %q: %w", sku, errorexercises.ErrNotFound)
	}
	return price, nil
}

func TestCartTotal(t *testing.T) {
	prices := fakePrices{"book": 2500, "pen": 200}
	if got, err := CartTotal(prices, []string{"book", "pen"}); got != 2700 || err != nil {
		t.Fatalf("CartTotal: got (%d, %v), want (2700, nil)", got, err)
	}
	_, err := CartTotal(prices, []string{"book", "missing"})
	if !errors.Is(err, errorexercises.ErrNotFound) || !strings.Contains(err.Error(), "missing") {
		t.Fatalf("CartTotal error: got %v, want wrapped ErrNotFound mentioning missing", err)
	}
}
