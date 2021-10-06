module github.com/figment-networks/filecoin-indexer

go 1.15

require (
	github.com/figment-networks/indexing-engine v0.7.4
	github.com/filecoin-project/go-address v0.0.4
	github.com/filecoin-project/go-bitfield v0.2.1
	github.com/filecoin-project/go-jsonrpc v0.1.2-0.20201008195726-68c6a2704e49
	github.com/filecoin-project/go-state-types v0.0.0-20201013222834-41ea465f274f
	github.com/filecoin-project/lotus v1.1.2
	github.com/filecoin-project/specs-actors/v2 v2.2.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.8.2
	github.com/ipfs/go-cid v0.0.7
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pressly/goose v2.6.0+incompatible
	github.com/rollbar/rollbar-go v1.2.0
	github.com/shopspring/decimal v0.0.0-20200227202807-02e2044944cc
	github.com/stretchr/stew v0.0.0-20130812190256-80ef0842b48b
	go.uber.org/multierr v1.6.0
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.12
)
