module github.com/Eric-GreenComb/eth-server

go 1.15

replace github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e => golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e

require (
	github.com/Eric-GreenComb/go-bip39 v1.0.0
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/dgraph-io/ristretto v0.0.3
	github.com/eko/gocache v1.1.1
	github.com/ethereum/go-ethereum v1.10.10
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/freeport v0.0.0-20150612182905-d4adf43b75b9 // indirect
	github.com/facebookgo/grace v0.0.0-20180706040059-75cf19382434
	github.com/facebookgo/httpdown v0.0.0-20180706035922-5979d39b15c2 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/stats v0.0.0-20151006221625-1b76add642e4 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	github.com/google/uuid v1.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/pborman/uuid v1.2.1
	github.com/shopspring/decimal v1.2.0
	github.com/spf13/viper v1.7.1
)
