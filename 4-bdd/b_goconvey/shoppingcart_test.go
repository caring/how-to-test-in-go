package bgoconvey

import (
	"fmt"
	"testing"

	"github.com/caring/test/4-bdd/shoppingcart"
	. "github.com/smartystreets/goconvey/convey"
)

type testCase struct {
	lines             []shoppingcart.LineItem
	expectedTotalCost int
}

func newCase(expectedTotal int, lines ...shoppingcart.LineItem) testCase {
	return testCase{lines: lines, expectedTotalCost: expectedTotal}
}

func (c testCase) Name() string {
	if len(c.lines) == 0 {
		return "is empty"
	}
	contents := "contains "
	for i, line := range c.lines {
		if i > 0 {
			contents = fmt.Sprintf("%s, ", contents)
		}
		contents = fmt.Sprintf("%s%d %s (%d/ea)", contents, line.Quantity, line.Item.Name, line.Item.Price)
	}
	return contents
}

func TestTotalCost(t *testing.T) {
	soap := shoppingcart.Item{ID: "1", Name: "Soap", Price: 2}
	shampoo := shoppingcart.Item{ID: "2", Name: "Shampoo", Price: 3}
	for _, testCase := range []testCase{
		newCase(0),
		newCase(2, shoppingcart.LineItem{Item: soap, Quantity: 1}),
		newCase(4, shoppingcart.LineItem{Item: soap, Quantity: 2}),
		newCase(3, shoppingcart.LineItem{Item: shampoo, Quantity: 1}),
		newCase(6, shoppingcart.LineItem{Item: shampoo, Quantity: 2}),
		newCase(5, shoppingcart.LineItem{Item: soap, Quantity: 1}, shoppingcart.LineItem{Item: shampoo, Quantity: 1}),
		newCase(6, shoppingcart.LineItem{Item: shampoo, Quantity: 1}, shoppingcart.LineItem{Item: shampoo, Quantity: 1}),
	} {
		Convey("Given I have a new shopping cart", t, func() {
			cart := shoppingcart.NewShoppingCart()

			Convey(fmt.Sprintf("When the cart %s", testCase.Name()), func() {
				for _, line := range testCase.lines {
					cart.AddItem(line.Item, uint(line.Quantity))
				}

				Convey(fmt.Sprintf("Then the total quantity must be %d", testCase.expectedTotalCost), func() {
					So(cart.TotalCost(), ShouldEqual, testCase.expectedTotalCost)
				})
			})
		})
	}
}
