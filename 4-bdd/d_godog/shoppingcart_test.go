package d_godog

import (
	"context"
	"fmt"
	"testing"

	"github.com/caring/test/4-bdd/shoppingcart"
	"github.com/cucumber/godog"
)

// use custom structs to avoid collisions w/ keys as strings
type cartKey struct{}
type catalogKey struct{}

func iHaveANewShoppingCart(ctx context.Context) context.Context {
	return context.WithValue(ctx, cartKey{}, shoppingcart.NewShoppingCart())
}

func theCartIsEmpty(ctx context.Context) error {
	cart, err := getCart(ctx)
	if err != nil {
		return err
	}
	if !cart.IsEmpty() {
		return fmt.Errorf("cart was not empty")
	}
	return nil
}

func addItemToCatalog(ctx context.Context, name string, price int) context.Context {
	catalog := getCatalog(ctx)
	catalog[name] = shoppingcart.Item{ID: name, Name: name, Price: price}
	return context.WithValue(ctx, catalogKey{}, catalog)
}

func addItemToCart(ctx context.Context, quantity int, name string) context.Context {
	catalog := getCatalog(ctx)
	cart, err := getCart(ctx)
	if err != nil {
		panic("cart is required")
	}
	item, exists := catalog[name]
	if !exists {
		panic("invalid item")
	}
	cart.AddItem(item, uint(quantity))
	return ctx
}

func theTotalQuantityMustBe(ctx context.Context, expectedTotal int) error {
	cart, err := getCart(ctx)
	if err != nil {
		panic("cart is required")
	}
	if cart.TotalCost() != expectedTotal {
		return fmt.Errorf("Expected a total of %d but found %d", expectedTotal, cart.TotalCost())
	}
	return nil
}

func getCatalog(ctx context.Context) map[string]shoppingcart.Item {
	i := ctx.Value(catalogKey{})
	catalog, ok := i.(map[string]shoppingcart.Item)
	if !ok {
		catalog = map[string]shoppingcart.Item{}
	}
	return catalog
}

func getCart(ctx context.Context) (*shoppingcart.ShoppingCart, error) {
	i := ctx.Value(cartKey{})
	cart, ok := i.(*shoppingcart.ShoppingCart)
	if !ok {
		return nil, fmt.Errorf("invalid cart")
	}
	return cart, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	ctx.Step(`^the item catalog has (\w+) for \$(\d+) each$`, addItemToCatalog)
	ctx.Step(`^I have a new shopping cart$`, iHaveANewShoppingCart)
	ctx.Step(`^the cart is empty$`, theCartIsEmpty)
	ctx.Step(`^I add (\d+) (\w+) to the cart$`, addItemToCart)
	ctx.Step(`^the total quantity must be (\d+)$`, theTotalQuantityMustBe)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
