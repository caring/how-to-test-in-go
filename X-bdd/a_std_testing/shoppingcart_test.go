package a_std_testing

import (
	"fmt"
	"testing"

	"github.com/caring/test/X-bdd/shoppingcart"
	"github.com/matryer/is"
)

type testCase struct {
	lines             []shoppingcart.LineItem
	expectedTotalCost int
}

func newCase(expectedTotal int, lines ...shoppingcart.LineItem) testCase {
	return testCase{lines: lines, expectedTotalCost: expectedTotal}
}

func (c testCase) Name(num int) string {
	return fmt.Sprintf("Example %d: %d lines", num, len(c.lines))
}

func (c testCase) Run(t *testing.T) {
	is := is.New(t)
	t.Logf("Given I have a new shopping cart\n")
	cart := shoppingcart.NewShoppingCart()
	if len(c.lines) == 0 {
		t.Logf(" When it is empty\n")
	} else {
		for i, line := range c.lines {
			keyword := "  And"
			if i == 0 {
				keyword = " When"
			}
			t.Logf("%s %d %s which costs %d each is added to the cart\n", keyword, line.Quantity, line.Item.Name, line.Item.Price)
			cart.AddItem(line.Item, uint(line.Quantity))
		}
		t.Logf(" Then the total cost must be %d", c.expectedTotalCost)
		is.Equal(cart.TotalCost(), c.expectedTotalCost) // should be equal
	}
}

func TestTotalCost(t *testing.T) {
	soap := shoppingcart.Item{ID: "1", Name: "Soap", Price: 2}
	shampoo := shoppingcart.Item{ID: "2", Name: "Shampoo", Price: 3}
	for i, testCase := range []testCase{
		newCase(0),
		newCase(2, shoppingcart.LineItem{Item: soap, Quantity: 1}),
		newCase(4, shoppingcart.LineItem{Item: soap, Quantity: 2}),
		newCase(3, shoppingcart.LineItem{Item: shampoo, Quantity: 1}),
		newCase(6, shoppingcart.LineItem{Item: shampoo, Quantity: 2}),
		newCase(5, shoppingcart.LineItem{Item: soap, Quantity: 1}, shoppingcart.LineItem{Item: shampoo, Quantity: 1}),
		newCase(6, shoppingcart.LineItem{Item: shampoo, Quantity: 1}, shoppingcart.LineItem{Item: shampoo, Quantity: 1}),
	} {
		t.Run(testCase.Name(i), testCase.Run)
	}
}
