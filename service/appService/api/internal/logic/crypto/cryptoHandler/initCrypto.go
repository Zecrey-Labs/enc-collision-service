package cryptohandler

import (
	"fmt"
	curve "zecrey-crypto/ecc/ztwistededwards/tebn254"

	"github.com/Zecrey-Labs/zecrey-collisions/common/model/crypto"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

const (
	BASICCOUNT = 20000000
)

func InitCrypto(ctx *svc.ServiceContext) error {
	err := ctx.CryptoModel.CreateCryptoTable()
	if err != nil {
		errInfo := fmt.Sprintf("[cryptohandler.CreateCryptoTable] %s", err.Error())
		logx.Error(errInfo)
		return err
	}
	count, err := ctx.CryptoModel.GetEncCollisionTotalCount()
	if err != nil {
		errInfo := fmt.Sprintf("[cryptohandler.GetEncCollisionTotalCount] %s", err.Error())
		logx.Error(errInfo)
		return err
	}
	if count < BASICCOUNT {
		current := curve.H
		base := curve.H
		for i := int64(1); i <= count/2; i++ {
			current = curve.Add(current, base)
		}
		for i := int64(count/2) + 1; i <= BASICCOUNT/2; i++ {
			isSuccess, err := ctx.CryptoModel.CreateCollision(&crypto.Crypto{
				EncData:      curve.ToString(current),
				EncCollision: i,
			})
			if err != nil {
				errInfo := fmt.Sprintf("[cryptohandler.CreateCollision] %s", err.Error())
				logx.Error(errInfo)
				return err
			} else if !isSuccess {
				errInfo := fmt.Sprintf("[cryptohandler.CreateCollision] %s", ErrInvalidCryptoInput)
				logx.Error(errInfo)
				return ErrInvalidCryptoInput
			}
			isSuccess, err = ctx.CryptoModel.CreateCollision(&crypto.Crypto{
				EncData:      curve.ToString(curve.Neg(current)),
				EncCollision: -i,
			})
			if err != nil {
				errInfo := fmt.Sprintf("[cryptohandler.CreateCollision] %s", err.Error())
				logx.Error(errInfo)
				return err
			} else if !isSuccess {
				errInfo := fmt.Sprintf("[cryptohandler.CreateCollision] %s", ErrInvalidCryptoInput)
				logx.Error(errInfo)
				return ErrInvalidCryptoInput
			}
			if i%10000 == 0 {
				logx.Info("Insert 10000 records")
			}
			current = curve.Add(current, base)
		}
	}
	return nil
}
