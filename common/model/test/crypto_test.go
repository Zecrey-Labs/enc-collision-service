package test

import (
	"fmt"
	"testing"
	curve "zecrey-crypto/ecc/ztwistededwards/tebn254"

	"github.com/Zecrey-Labs/zecrey-collisions/common/model/crypto"
	"github.com/stretchr/testify/assert"
)

func TestCreateCryptoTable(t *testing.T) {
	// Get Service Context
	ctx := TestServiceContext(ConfigProvider("test.yaml"))
	err := ctx.Crypto.CreateCryptoTable()
	assert.Nil(t, err)
}

func TestDropCryptoTable(t *testing.T) {
	// Get Service Context
	ctx := TestServiceContext(ConfigProvider("test.yaml"))
	err := ctx.Crypto.DropCryptoTable()
	assert.Nil(t, err)
}

func TestCreateCollision(t *testing.T) {
	// Get Service Context
	ctx := TestServiceContext(ConfigProvider("test.yaml"))
	current := curve.H
	base := curve.H
	for i := int64(1); i <= 3; i++ {
		isSuccess, err := ctx.Crypto.CreateCollision(&crypto.Crypto{
			EncData:      curve.ToString(current),
			EncCollision: i,
		})
		assert.Nil(t, err)
		assert.NotEqual(t, false, isSuccess)
		ctx.Crypto.CreateCollision(&crypto.Crypto{
			EncData:      curve.ToString(curve.Neg(current)),
			EncCollision: -i,
		})
		assert.Nil(t, err)
		assert.NotEqual(t, false, isSuccess)
		current = curve.Add(current, base)
	}

}
func TestGetEncCollisionByEncData(t *testing.T) {
	// Get Service Context
	ctx := TestServiceContext(ConfigProvider("test.yaml"))
	collision, err := ctx.Crypto.GetEncCollisionByEncData("DIVSZPNo//N09abkmcZDPTVfi1KtuKFVFT4ElQyAJoM=")
	assert.Nil(t, err)
	assert.NotNil(t, collision)
}

func TestGetEncCollisionTotalCount(t *testing.T) {
	// Get Service Context
	ctx := TestServiceContext(ConfigProvider("test.yaml"))
	count, err := ctx.Crypto.GetEncCollisionTotalCount()
	assert.Nil(t, err)
	assert.Equal(t, count, int64(200000))
}

func TestGerRandomCollision(t *testing.T) {
	// Get Service Context
	ctx := TestServiceContext(ConfigProvider("test.yaml"))
	enc_data, err := ctx.Crypto.GerRandomCollision()
	assert.Nil(t, err)
	fmt.Println(enc_data)
	assert.NotNil(t, enc_data)
}
