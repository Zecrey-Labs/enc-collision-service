package svc

import (
	"github.com/Zecrey-Labs/zecrey-collisions/common/model/crypto"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/config"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	CryptoModel crypto.CryptoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormPointer, err := gorm.Open(postgres.Open(c.Postgres.DataSource))
	if err != nil {
		logx.Errorf("gorm connect db error, err = %s", err.Error())
	}
	conn := sqlx.NewSqlConn("postgres", c.Postgres.DataSource)
	return &ServiceContext{
		Config:      c,
		CryptoModel: crypto.NewCryptoModel(conn, c.CacheRedis, gormPointer),
	}
}
