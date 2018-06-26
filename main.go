package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Eric-GreenComb/eth-server/config"
	"github.com/Eric-GreenComb/eth-server/ethereum"
	"github.com/Eric-GreenComb/eth-server/handler"
)

var (
	g errgroup.Group
)

func main() {
	if config.ServerConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	ethereum.Init()

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api base */
	r0 := router.Group("/")
	{
		r0.GET("", handler.Index)
		r0.GET("health", handler.Health)
	}

	// api
	r1 := router.Group("/ethereum")
	{
		r1.GET("/nonce", handler.PendingNonce)

		r1.POST("/send", handler.SendEthCoin)
		r1.GET("/balance/:addr", handler.GetBalance)

		r1.GET("/wei/string/:val/:decimals", handler.StringToWei)
	}

	r2 := router.Group("/erc20")
	{
		r2.POST("/deploy", handler.DeployErc20)
		r2.POST("/transfer", handler.TransferErc20)
		r2.GET("/balance/:conaddr/:addr", handler.GetErc20Balance)

		r2.POST("/rawtransfer", handler.RawTransferErc20)
	}

	r3 := router.Group("/badger")
	{
		r3.POST("/set", handler.SetBadgerKey)
		r3.POST("/setwithttl", handler.SetBadgerKeyTTL)
		r3.GET("/get/:key", handler.GetBadgerKey)
	}

	for _, _port := range config.ServerConfig.Port {
		server := &http.Server{
			Addr:         ":" + _port,
			Handler:      router,
			ReadTimeout:  300 * time.Second,
			WriteTimeout: 300 * time.Second,
		}

		g.Go(func() error {
			return server.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
