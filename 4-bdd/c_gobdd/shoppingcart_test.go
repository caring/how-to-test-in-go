package c_gobdd

import (
	"errors"
	"fmt"
	"testing"

	"github.com/caring/test/4-bdd/shoppingcart"
	"github.com/go-bdd/gobdd"
)

// use custom structs to avoid collisions w/ keys as strings
type cartKey struct{}
type catalogKey struct{}

func initializeCart(t gobdd.StepTest, ctx gobdd.Context) {
	ctx.Set(cartKey{}, shoppingcart.NewShoppingCart())
}

func cartIsEmpty(t gobdd.StepTest, ctx gobdd.Context) {
	cart := getCart(ctx)
	if !cart.IsEmpty() {
		t.Error(errors.New("the cart was not empty"))
	}
}

func addItemToCatalog(t gobdd.StepTest, ctx gobdd.Context, name string, cost int) {
	catalog := getCatalog(ctx)
	catalog[name] = shoppingcart.Item{ID: name, Name: name, Price: cost}
	ctx.Set(catalogKey{}, catalog)
}

func addItemToCart(t gobdd.StepTest, ctx gobdd.Context, quantity int, name string) {
	catalog := getCatalog(ctx)
	cart := getCart(ctx)

	item, exists := catalog[name]
	if !exists {
		t.Error(fmt.Errorf("Invalid item %s", name))
	}
	cart.AddItem(item, uint(quantity))
}

func checkTotal(t gobdd.StepTest, ctx gobdd.Context, expectedTotal int) {
	cart := getCart(ctx)

	if cart.TotalCost() != expectedTotal {
		t.Error(fmt.Errorf("found total cost of %d but expected %d", cart.TotalCost(), expectedTotal))
	}
}

func getCart(ctx gobdd.Context) *shoppingcart.ShoppingCart {
	cart, err := ctx.Get(cartKey{})
	if err != nil {
		panic(err)
	}

	return cart.(*shoppingcart.ShoppingCart)
}

func getCatalog(ctx gobdd.Context) map[string]shoppingcart.Item {
	catalog, _ := ctx.Get(catalogKey{})
	if catalog == nil {
		catalog = map[string]shoppingcart.Item{}
		ctx.Set(catalogKey{}, catalog)
	}

	return catalog.(map[string]shoppingcart.Item)
}

func TestScenarios(t *testing.T) {
	suite := gobdd.NewSuite(t)
	suite.AddStep(`I have a new shopping cart`, initializeCart)
	suite.AddStep(`the cart is empty`, cartIsEmpty)
	suite.AddStep(`the item catalog has {word} for \$(\d+) each`, addItemToCatalog)
	suite.AddStep(`I add ([1-9]\d*) {word} to the cart`, addItemToCart)
	suite.AddStep(`the total quantity must be (\d+)`, checkTotal)
	suite.Run()
}
