An attempt and implementing an Abstract Factory in Go. This was the first piece of Go code i've ever written so it's very
rough around the edges; and actually, the code doesn't work entirely as intended. If any product needed to have some sort of
internal state that it operated upon, issues arise. For example, if the product B of family 1 encapsulated some integral
value that would be the target of its `scale()` method, with some getter `get()` to grab the result, this code fails.

```go
b := factory.Construct_B()
b.Scale( 2 ) // Passes a COPY of b to scale and performs the arithmetic
b.Get() // Gets the DEFAULT, NON-SCALED value that b was constructed with
```
