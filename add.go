package mypackage

import "rsc.io/quote"

func Add(a, b int) string {
	_ = a + b
	return quote.Hello()
}

func sub() {
	return
}
