Feature: Calculating Total Cost

Background: Item catalog
Given the item catalog has soap for $2 each
  And the item catalog has shampoo for $3 each

Scenario: Empty cart
Given I have a new shopping cart
 When the cart is empty
 Then the total quantity must be 0

Scenario Outline: Single items added
Given I have a new shopping cart
 When I add <quantity> <item> to the cart
 Then the total quantity must be <expectedQuantity>
Examples:
|item   |quantity|expectedQuantity|
|soap   |2       |4               |
|shampoo|3       |9               |
|shampoo|0       |0               |

Scenario: Multiple items added
Given I have a new shopping cart
 When I add 1 soap to the cart
  And I add 1 shampoo to the cart
 Then the total quantity must be 5

Scenario: Same item added multiple times
Given I have a new shopping cart
 When I add 1 shampoo to the cart
  And I add 1 shampoo to the cart
 Then the total quantity must be 6