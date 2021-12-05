package cryptohandler

import (
	"fmt"
	curve "zecrey-crypto/ecc/ztwistededwards/tebn254"

	"github.com/Zecrey-Labs/zecrey-collisions/common/model/crypto"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

func InitCrypto(ctx *svc.ServiceContext) error {
	err := ctx.CryptoModel.CreateCryptoTable()
	if err != nil {
		errInfo := fmt.Sprintf("[cryptohandler.CreateCryptoTable] %s", err.Error())
		logx.Error(errInfo)
		return err
	}
	current := curve.H
	base := curve.H
	for i := int64(1); i <= 5000000; i++ {
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
		current = curve.Add(current, base)
	}
	return nil
}
