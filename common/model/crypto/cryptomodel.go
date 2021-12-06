package crypto

import (
	"fmt"

	"github.com/tal-tech/go-zero/core/bloom"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var (
	ZeroPoint          = "AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
	cacheEncDataPrefix = "cache::crypto:endData:"
)

type (
	CryptoModel interface {
		CreateCryptoTable() error
		DropCryptoTable() error
		CreateCollision(Crypto *Crypto) (bool, error)
		CreateCryptoInBatches(Crypto []*Crypto) (rowsAffected int64, err error)
		GetEncCollisionByEncData(encData string) (collisions *Crypto, err error)
		GetEncCollisionTotalCount() (count int64, err error)
	}

	defaultCryptoModel struct {
		sqlc.CachedConn
		table string
		DB    *gorm.DB
		bloom *bloom.Filter
	}

	Crypto struct {
		gorm.Model
		EncCollision int64
		EncData      string
	}
)

func NewCryptoModel(conn sqlx.SqlConn, c cache.CacheConf, db *gorm.DB) CryptoModel {
	return &defaultCryptoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `crypto`,
		DB:         db,
		bloom:      bloom.New(redis.New(c[0].Host), "cryptomodelbloom", 200000),
	}
}

func (*Crypto) TableName() string {
	return `crypto`
}

/*
	Func: cryptoQueryRow
	Params: v interface{}, key string, query func() error
	Return: error
	Description: set and get by cache
*/
func (m *defaultCryptoModel) cryptoQueryRow(v interface{}, key string, query func() error) error {
	isExists, err := m.bloom.Exists([]byte(key))
	if err != nil {
		return err
	}
	if isExists {
		err = m.GetCache(key, v)
		logx.Info("[crypto.cache] key exists")
		return err
	} else {
		err = query()
		if err != nil {
			if err.Error() == ErrNotFound.Error() {
				err := m.SetCache(key, v)
				if err != nil {
					errInfo := fmt.Sprintf("[crypto.cache] %s", err)
					logx.Error(errInfo)
					return err
				}

				err = m.bloom.Add([]byte(key))
				if err != nil {
					errInfo := fmt.Sprintf("[crypto.bloom] %s", err)
					logx.Error(errInfo)
					return err
				}
				return ErrNotFound
			}
			return err
		}

		err := m.SetCache(key, v)
		if err != nil {
			errInfo := fmt.Sprintf("[crypto.cache] %s", err)
			logx.Error(errInfo)
			return err
		}
		err = m.bloom.Add([]byte(key))
		if err != nil {
			errInfo := fmt.Sprintf("[crypto.bloom] %s", err)
			logx.Error(errInfo)
			return err
		}
		return nil
	}
}

/*
	Func: CreateCryptoTable
	Params:
	Return: err error
	Description: create crypto table
*/
func (m *defaultCryptoModel) CreateCryptoTable() error {
	return m.DB.AutoMigrate(Crypto{})
}

/*
	Func: DropCryptoTable
	Params:
	Return: err error
	Description: drop crypto table
*/
func (m *defaultCryptoModel) DropCryptoTable() error {
	return m.DB.Migrator().DropTable(m.table)
}

/*
	Func: GetEncCollisionByEncData
	Params: encData string
	Return: collision *Crypto, err error
	Description: get collision by enc data
*/
func (m *defaultCryptoModel) GetEncCollisionByEncData(encData string) (collision *Crypto, err error) {
	encDataKey := fmt.Sprintf("%s%s", cacheEncDataPrefix, encData)
	err = m.cryptoQueryRow(&collision, encDataKey, func() error {
		if encData == ZeroPoint {
			collision = &Crypto{
				EncCollision: 0,
				EncData:      encData,
			}
			return nil
		}
		dbTx := m.DB.Table(m.table).Where("enc_data=?", encData).Find(&collision)
		if dbTx.Error != nil {
			logx.Error("[crypto.GetEncCollisionByEncData] %s", dbTx.Error)
			return dbTx.Error
		} else if dbTx.RowsAffected == 0 {
			logx.Error("[crypto.GetEncCollisionByEncData] %s", ErrNotFound)
			return ErrNotFound
		}
		return nil
	})
	if err != nil {
		errInfo := fmt.Sprintf("[crypto.GetEncCollisionByEncData] %s", err)
		logx.Error(errInfo)
		return nil, err
	}
	if collision.EncCollision == 0 && collision.EncData != ZeroPoint {
		errInfo := fmt.Sprintf("[crypto.GetEncCollisionByEncData] %s", ErrNotFound)
		logx.Error(errInfo)
		return nil, ErrNotFound
	}
	return collision, nil
}

/*
	Func: GetEncCollisionByEncData
	Params: encData string
	Return: collision *Crypto, err error
	Description: get collision by enc data
*/
func (m *defaultCryptoModel) GetEncCollisionTotalCount() (count int64, err error) {
	dbTx := m.DB.Table(m.table).Count(&count)
	if dbTx.Error != nil {
		logx.Error("[crypto.GetEncCollisionByEncData] %s", dbTx.Error)
		return -1, dbTx.Error
	} else if dbTx.RowsAffected == 0 {
		logx.Error("[crypto.GetEncCollisionByEncData] %s", ErrNotFound)
		return -1, ErrNotFound
	}
	return count, nil
}

/*
	Func: CreateCollision
	Params: Crypto *Crypto
	Return: bool, error
	Description: create collision
*/
func (m *defaultCryptoModel) CreateCollision(Crypto *Crypto) (bool, error) {
	dbTx := m.DB.Table(m.table).Create(Crypto)
	if dbTx.Error != nil {
		err := fmt.Sprintf("[crypto.CreateCollision] %s", dbTx.Error)
		logx.Error(err)
		return false, dbTx.Error
	}
	if dbTx.RowsAffected == 0 {
		err := fmt.Sprintf("[crypto.CreateCollision] %s", ErrInvalidCryptoInput)
		logx.Error(err)
		return false, ErrInvalidCryptoInput
	}
	return true, nil
}

/*
	Func: CreateCryptoInBatches
	Params: []*Crypto
	Return: rowsAffected int64, err error
	Description: create Crypto batches
*/
func (m *defaultCryptoModel) CreateCryptoInBatches(Crypto []*Crypto) (rowsAffected int64, err error) {
	dbTx := m.DB.Table(m.table).CreateInBatches(Crypto, len(Crypto))
	if dbTx.Error != nil {
		err := fmt.Sprintf("[crypto.CreateCryptoInBatches] %s", dbTx.Error)
		logx.Error(err)
		return 0, dbTx.Error
	}
	if dbTx.RowsAffected == 0 {
		return 0, nil
	}
	return dbTx.RowsAffected, nil
}
