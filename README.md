# Assertions Repository

This repository exposes some useful methods for unit testing in Go.
Defines methods to do "assertions" on your tests to have more consice tests files.

# Assertions

## Assert

Asserts if the given condition is `true`.

```Go
func TestExampleOfUse(t *Testing.T) {
  Assertions.Assert(t, 3 == 3, "3 is equals to 3")
  Assertions.Assert(t, 3 == 0, "3 is equals to 0")
}
// Tests will fail with: "caller.go:10: Condition '3 is equals to 0' not met: expected 'true' given 'false'"
```

## AssertEquals

Asserts if the given values `expected` and `given` are equals.

```Go
func TestExampleOfUse(t *Testing.T) {
  Assertions.Assert(t, 3, 3, "3 is equals to 3")
  Assertions.Assert(t, 3, 0, "3 is equals to 0")
}
// Tests will fail with: "caller.go:10: Condition '3 is equals to 0' not met: expected '3' given '0'"
```
