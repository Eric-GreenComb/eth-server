package main

import (
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
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

	r3 := router.Group("/obj")
	{
		r3.POST("/deploy", handler.DeployObj)
		r3.POST("/set", handler.ModifyObjValue)
		r3.GET("/get/:address", handler.GetObjValue)

		r3.POST("/rawset", handler.RawModifyObjValue)
	}

	r4 := router.Group("/tipjar")
	{
		r4.POST("/deploy", handler.DeployTipJar)
		r4.POST("/deposit", handler.DepositTipJar)
		r4.GET("/balance/:address", handler.GetTipJarBalance)
	}

	r5 := router.Group("/savings")
	{
		r5.POST("/deploy", handler.DeploySavings)
		r5.POST("/deposit", handler.DepositSavings)
		r5.POST("/withdraw", handler.WithdrawSavings)
	}

	r6 := router.Group("/bank")
	{
		r6.POST("/deploy", handler.DeployBank)
		r6.POST("/deposit", handler.DepositBank)
		r6.POST("/withdraw", handler.WithdrawBank)
		r6.GET("/balance/:conaddr/:address", handler.GetBankBalance)
	}

	r7 := router.Group("/callobj")
	{
		r7.POST("/deploy", handler.DeployCallObj)
		r7.POST("/set", handler.ModifyCallObjValue)
	}

	r8 := router.Group("/info")
	{
		r8.POST("/deploy", handler.DeployInfo)
		r8.POST("/set_channel", handler.SetInfoChannel)
		r8.GET("/get_channel/:address", handler.GetInfoChannel)
		r8.POST("/set", handler.SetInfo)
		r8.GET("/get/:address", handler.GetInfo)
	}

	rfomo3d := router.Group("/fomo3d")
	{
		rfomo3d.POST("/deploy/team", handler.DeployTeam)
		rfomo3d.POST("/deploy/playerbook", handler.DeployPlayerBook)
		rfomo3d.POST("/deploy/f3d", handler.DeployF3d)
	}

	r100 := router.Group("/badger")
	{
		r100.POST("/set", handler.SetBadgerKey)
		r100.POST("/setwithttl", handler.SetBadgerKeyTTL)
		r100.GET("/get/:key", handler.GetBadgerKey)
	}

	r101 := router.Group("/account")
	{
		r101.POST("/create/:passphrase", handler.CreateAccount)
		r101.POST("/bip39/create", handler.CreateBIP39)
		r101.POST("/bip39/keystore/create", handler.CreateBIP39Keysore)

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
