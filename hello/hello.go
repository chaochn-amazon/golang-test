package hello

// add a dependency and the go.mod will automatically update due to the change
import (
	"rsc.io/quote/v3"
)

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
