package test

import (
	"flag"

	"github.com/Zecrey-Labs/zecrey-collisions/common/model/crypto"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/rest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	rest.RestConf
	Postgres struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}

type ServiceContext struct {
	Config Config
	Crypto crypto.CryptoModel
}

func TestServiceContext(c Config) *ServiceContext {
	gormPointer, err := gorm.Open(postgres.Open(c.Postgres.DataSource))
	if err != nil {
		logx.Errorf("gorm connect db error, err = %s", err.Error())
	}
	conn := sqlx.NewSqlConn("postgres", c.Postgres.DataSource)
	return &ServiceContext{
		Config: c,
		Crypto: crypto.NewCryptoModel(conn, c.CacheRedis, gormPointer),
	}
}

func ConfigProvider(path string) Config {
	logx.Info("ConfigProvider")

	// Parse .yaml file
	var c Config
	var configFile = flag.String("f", path, "the config file")
	flag.Parse()
	conf.MustLoad(*configFile, &c)
	return c
}
