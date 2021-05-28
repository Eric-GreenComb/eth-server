package main

import (
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Eric-GreenComb/eth-server/cached"
	"github.com/Eric-GreenComb/eth-server/config"
	"github.com/Eric-GreenComb/eth-server/ethereum"
	"github.com/Eric-GreenComb/eth-server/handler"
	myrouter "github.com/Eric-GreenComb/eth-server/router"
)

var (
	g errgroup.Group
)

func main() {
	if config.ServerConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	cached.Init()

	ethereum.Init()

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api base */
	myrouter.SetupBaseRouter(router)

	// api
	rethereum := router.Group("/ethereum")
	{
		rethereum.GET("/chainid", handler.GetChainID)
		rethereum.GET("/nonce", handler.PendingNonce)
		rethereum.POST("/send", handler.SendEthCoin)
		rethereum.GET("/balance/:addr", handler.GetBalance)

		rethereum.GET("/wei/string/:val/:decimals", handler.StringToWei)
	}

	r101 := router.Group("/account")
	{
		r101.POST("/create/:passphrase", handler.CreateAccount)
		r101.POST("/bip39/create", handler.CreateBIP39)
		r101.POST("/bip39/keystore/create", handler.CreateBIP39Keysore)

		r101.GET("/bip39/keystore/:address", handler.GetBIP39Keysore)
		r101.POST("/checkpwd", handler.CheckPassphrase)
	}

	for _, _port := range config.ServerConfig.Port {
		server := &http.Server{
			Addr:         ":" + _port,
			Handler:      router,
			ReadTimeout:  300 * time.Second,
			WriteTimeout: 300 * time.Second,
		}

		g.Go(func() error {
			return gracehttp.Serve(server)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
