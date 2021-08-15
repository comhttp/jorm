package coingecko

import (
	"fmt"
	gecko "github.com/superoo7/go-gecko/v3"
	"log"
)

func main() {
	cg := gecko.NewClient(nil)
	list, err := cg.CoinsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available coins:", len(*list))
}
