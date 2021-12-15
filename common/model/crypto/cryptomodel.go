package crypto

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
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
		GerRandomCollision() (collision *Crypto, err error)
	}

	defaultCryptoModel struct {
		sqlc.CachedConn
		table string
		DB    *gorm.DB
	}

	Crypto struct {
		gorm.Model
		EncCollision int64
		EncData      string `gorm:"uniqueIndex"`
	}
)

func NewCryptoModel(conn sqlx.SqlConn, c cache.CacheConf, db *gorm.DB) CryptoModel {
	return &defaultCryptoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `crypto`,
		DB:         db,
	}
}

func (*Crypto) TableName() string {
	return `crypto`
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
	err = m.QueryRow(&collision, encDataKey, func(conn sqlx.SqlConn, v interface{}) error {
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
	Func: GerRandomCollision
	Params:
	Return: collision *Crypto, err error
	Description: get random collision
*/
func (m *defaultCryptoModel) GerRandomCollision() (collision *Crypto, err error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := r.Intn(50000)
	randFlag := r.Intn(2)
	if randFlag == 1 {
		randNum = -randNum
	}
	dbTx := m.DB.Table(m.table).Where("enc_collision=?", randNum).Find(&collision)
	if dbTx.Error != nil {
		logx.Error("[crypto.GetEncCollisionByEncData] %s", dbTx.Error)
		return nil, dbTx.Error
	} else if dbTx.RowsAffected == 0 {
		logx.Error("[crypto.GetEncCollisionByEncData] %s", ErrNotFound)
		return nil, ErrNotFound
	}
	return collision, nil
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
