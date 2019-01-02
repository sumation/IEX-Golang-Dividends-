package main

import (
	"fmt"
	"net/http"

	"github.com/timpalpant/go-iex"
)

func main() {
	client := iex.NewClient(&http.Client{})
	dividends, err := client.GetDividends("AAPL")
	quotes, err := client.GetTOPS([]string{"AAPL", "SPY"})

	if err != nil {
		panic(err)
	}

	for _, quote := range quotes {
		fmt.Printf("%v: bid $%.02f (%v shares), ask $%.02f (%v shares) [as of %v]\n",
			quote.Symbol, quote.BidPrice, quote.BidSize,
			quote.AskPrice, quote.AskSize, quote.LastUpdated)

		for _, d := range dividends {
			fmt.Printf("Ex date: %v, Payment date: %v, Record date: %v, Declared date: %v, Amount: %v\n",
				d.ExDate, d.PaymentDate, d.RecordDate, d.DeclaredDate, d.Amount)
		}
	}
}
