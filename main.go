package main

import (
	"fmt"
	"os"

	"github.com/BaritoLog/barito-router/router"
)

const (
	EnvRouterAddress   = "BARITO_ROUTER_ADDRESS"
	EnvBaritoMarketUrl = "BARITO_ROUTER_MARKET_URL"
	Version            = "0.0.1"
)

func main() {
	routerAddress := os.Getenv(EnvRouterAddress)
	if routerAddress == "" {
		routerAddress = ":8081"
	}

	baritoMarketUrl := os.Getenv(EnvBaritoMarketUrl)
	if baritoMarketUrl == "" {
		baritoMarketUrl = "http://localhost:3000/api/apps"
	}

	fmt.Printf(".: Barito Router v%s :.\n\n", Version)
	fmt.Printf("%s=%s\n", EnvRouterAddress, routerAddress)
	fmt.Printf("%s=%s\n", EnvBaritoMarketUrl, baritoMarketUrl)
	fmt.Printf("\n")

	trader := router.NewTrader(baritoMarketUrl)
	consul := router.NewConsulHandler()

	r := router.NewRouter(routerAddress, trader, consul)
	r.Server().ListenAndServe()
}
