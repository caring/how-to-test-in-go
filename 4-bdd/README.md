# BDD/TDD style tests

The tests in this folder are intended to demostrate using various approaches and libraries that are commonly used when taking a TDD or BDD approach.

Normally, the code and the tests live in the same folders.  Just to try each approach against the same code, I created subfolders with each individual test
style. If/when chosing one of these test styles, I would recommend continuing with your current approach to organizing code and tests. Utilizing the separate 
folders here is just to separate the examples and is not intended to be a recommendation for changing how any of your code or tests are structured.

The approaches illustrated here all try to incorporate [gherkin](https://cucumber.io/docs/gherkin/reference/) or gherkin-like syntax, which is expressed as given-when-then statements.
A common pattern for writing good tests is [arrange-act-assert](https://automationpanda.com/2020/07/07/arrange-act-assert-a-pattern-for-writing-good-tests/).
The given-when-then's of gherkin is roughly analagous to arrange-act-assert.
- given/arrange: sets up test inputs and targets
- when/act: performs the target behavior
- then/assert: validates expected outcomes

The approached illustrated are:
- using the standard test library
- using the goconvey library
- using the gobdd library
- using the godog library

(note: If you wanted to take an approach where you did something like combine using a library like goconvey with a library that uses external feature files like godog, that would be a valid approach.  Using goconvey could be useful in making the step implementations for godog more expressive, for example.)

The first two examples are intended to illustrate how to write more expressive tests. Utilizing either of these approached would involve initially writing
acceptance criteria as something that would translate into given-when-then's, and then incorporating those given-when-then's into your test code to
decorate the test code and make it more expressive.  By making the tests more expressive, what you're trying to accomplish is minimizing the mental overhead
required in context switching between business/operational processes, acceptance criteria, and test code.

The next two examples utilize libraries that allow you to capture business/operational processes and/or acceptance criteria as a feature file that is readable
by non-technical stake holders. Ideally, this feature file is created collaboratively with those people in a way that is less ambiguous and more easily captures
non-happy path cases then can typically be done in paragraph format. The libraries in these examples will then parse the exact same feature files that have
been collaborativley created and turns them into test code. The developer would then be responsible to creating the code that backs the steps in the feature file.

Note that for both gobdd and godog, a scenario outline in gherkin is like a template for scenarios. Each example executes as if it was its own scenario, and like
other scenarios should have its own clean context. Any background section will be executed before each example in a scenario outline, just like it is executed
before each ordinary scenario.

To illustrate what the output for a failed test looks like with each approach, I have intentionally introduced a bug in the AddItem function.

## To see the failures using each example:
In shoppingcart/shoppingcart.go, you can uncomment line 29 (under "buggy code:") and comment line 31 (under "correct code:") to see what the output would
look like for yourself when tests fail.

```go
func (s *ShoppingCart) AddItem(item Item, quantity uint) {
	if line, found := s.lines[item.ID]; found {
		//buggy code:
		// line.Quantity += int(quantity)
		// correct code:
		s.lines[item.ID] = LineItem{Item: item, Quantity: line.Quantity + int(quantity)}
		return
	}
	s.lines[item.ID] = LineItem{Item: item, Quantity: int(quantity)}
}
```
