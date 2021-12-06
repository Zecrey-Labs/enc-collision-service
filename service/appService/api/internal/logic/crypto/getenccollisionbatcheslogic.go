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

type GetEncCollisionBatchesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEncCollisionBatchesLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetEncCollisionBatchesLogic {
	return GetEncCollisionBatchesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func packGetEncCollisionBatches(
	status int,
	msg string,
	err string,
	result []*types.ResultGetEncCollisionBatches,
) (res *types.RespGetEncCollisionBatches) {
	return &types.RespGetEncCollisionBatches{
		Status: status,
		Msg:    msg,
		Err:    err,
		Result: result,
	}
}

func (l *GetEncCollisionBatchesLogic) GetEncCollisionBatches(req types.ReqGetEncCollisionBatches) (resp *types.RespGetEncCollisionBatches, err error) {
	// todo: add your logic here and delete this line
	respResult := make([]*types.ResultGetEncCollisionBatches, 0)
	encDataBatches := make([]string, 0)
	for _, encData := range req.EncDataBatches {
		err = utils.CheckRequestParam(utils.TypeEncData, reflect.ValueOf(encData))
		if err != nil {
			errInfo := fmt.Sprintf("[appService.crypto.GetEncCollisionBatches] %s:%s", encData, err)
			logx.Error(errInfo)
			return packGetEncCollisionBatches(types.FailStatus, types.FailMsg, errInfo, respResult), nil
		}
		encData := utils.CleanEncData(encData)
		err = utils.CheckRequestParam(utils.TypeEncDataOmitSpace, reflect.ValueOf(encData))
		if err != nil {
			errInfo := fmt.Sprintf("[appService.crypto.GetEncCollisionBatches] %s:%s", encData, err)
			logx.Error(errInfo)
			return packGetEncCollisionBatches(types.FailStatus, types.FailMsg, errInfo, respResult), nil
		}
		encDataBatches = append(encDataBatches, encData)
	}
	for _, encData := range encDataBatches {
		collision, err := l.svcCtx.CryptoModel.GetEncCollisionByEncData(encData)
		if err != nil {
			errInfo := fmt.Sprintf("[appService.crypto.GetEncCollisionBatches]<=>[CryptoModel.GetEncCollision] %s:%s", encData, err)
			logx.Errorf(errInfo)
			return packGetEncCollisionBatches(types.FailStatus, types.FailMsg, errInfo, make([]*types.ResultGetEncCollisionBatches, 0)), nil
		}
		respResult = append(respResult, &types.ResultGetEncCollisionBatches{
			EncData:         encData,
			CollisionResult: collision.EncCollision,
		})
	}
	return packGetEncCollisionBatches(types.SuccessStatus, types.SuccessMsg, "", respResult), nil
}
