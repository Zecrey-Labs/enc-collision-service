package crypto

import (
	"net/http"

	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/logic/crypto"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/svc"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetEncCollisionHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqGetEncCollision
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := crypto.NewGetEncCollisionLogic(r.Context(), ctx)
		resp, err := l.GetEncCollision(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
