package hello

import (
	"rsc.io/quote/v3"
)


// hello
func Hello(ss string ) string {
	return quote.HelloV3() + ss
}

func Proverb() string {
	return quote.Concurrency()
}