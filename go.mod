module github.com/Zecrey-Labs/zecrey-collisions

go 1.15

require (
	github.com/stretchr/testify v1.7.0
	github.com/tal-tech/go-zero v1.2.4
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.4
	zecrey-crypto v0.0.1
)

replace zecrey-crypto => github.com/Zecrey-Labs/zecrey-crypto v0.0.1
