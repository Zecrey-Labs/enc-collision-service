package crypto

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Zecrey-Labs/zecrey-collisions/common/model/utils"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/svc"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetEncCollisionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEncCollisionLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetEncCollisionLogic {
	return GetEncCollisionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func packGetGetEncCollision(
	status int,
	msg string,
	err string,
	result types.ResultGetEncCollision,
) (res *types.RespGetEncCollision) {
	return &types.RespGetEncCollision{
		Status: status,
		Msg:    msg,
		Err:    err,
		Result: result,
	}
}

func (l *GetEncCollisionLogic) GetEncCollision(req types.ReqGetEncCollision) (resp *types.RespGetEncCollision, err error) {
	var respResult types.ResultGetEncCollision

	err = utils.CheckRequestParam(utils.TypeEncData, reflect.ValueOf(req.EncData))
	if err != nil {
		errInfo := fmt.Sprintf("[appService.crypto.GetEncCollision] %s", err)
		logx.Error(errInfo)
		return packGetGetEncCollision(types.FailStatus, types.FailMsg, errInfo, respResult), nil
	}
	encData := utils.CleanEncData(req.EncData)
	err = utils.CheckRequestParam(utils.TypeEncDataOmitSpace, reflect.ValueOf(encData))
	if err != nil {
		errInfo := fmt.Sprintf("[appService.crypto.GetEncCollision] %s", err)
		logx.Error(errInfo)
		return packGetGetEncCollision(types.FailStatus, types.FailMsg, errInfo, respResult), nil
	}
	collision, err := l.svcCtx.CryptoModel.GetEncCollisionByEncData(encData)
	if err != nil {
		errInfo := fmt.Sprintf("[appService.crypto.GetEncCollision]<=>[CryptoModel.GetEncCollision] %s", err.Error())
		logx.Errorf(errInfo)
		return packGetGetEncCollision(types.FailStatus, types.FailMsg, errInfo, respResult), nil
	}
	return packGetGetEncCollision(types.SuccessStatus, types.SuccessMsg, "", types.ResultGetEncCollision{
		CollisionResult: collision.EncCollision,
	}), nil
}
