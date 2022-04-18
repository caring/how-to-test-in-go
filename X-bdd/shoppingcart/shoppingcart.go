package shoppingcart

type Item struct {
	ID    string
	Name  string
	Price int
}

type LineItem struct {
	Item     Item
	Quantity int
}

type ShoppingCart struct {
	lines map[string]LineItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{lines: map[string]LineItem{}}
}

func (s *ShoppingCart) IsEmpty() bool {
	return len(s.lines) == 0
}

func (s *ShoppingCart) AddItem(item Item, quantity uint) {
	if line, found := s.lines[item.ID]; found {
		//buggy code:
		line.Quantity += int(quantity)
		// correct code:
		// s.lines[item.ID] = LineItem{Item: item, Quantity: line.Quantity + int(quantity)}
		return
	}
	s.lines[item.ID] = LineItem{Item: item, Quantity: int(quantity)}
}

func (s *ShoppingCart) RemoveItem(itemID string, quantity uint) {
	if line, found := s.lines[itemID]; found {
		remainingQuantity := line.Quantity - int(quantity)
		if remainingQuantity > 0 {
			line.Quantity -= int(quantity)
			return
		}
		delete(s.lines, itemID)
	}
}

func (s *ShoppingCart) TotalCost() int {
	total := 0
	for _, line := range s.lines {
		total += line.Item.Price * line.Quantity
	}
	return total
}
