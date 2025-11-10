package handler

import (
	"directBTC/api/internal/logic"
	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetBtcAddressIsTrialHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBtcAddressIsTrialReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, &response.ErrorCode{Code: -2, Msg: err.Error()})
			return
		}

		l := logic.NewGetBtcAddressIsTrialLogic(r.Context(), svcCtx)
		resp, err := l.GetBtcAddressIsTrial(&req)
		response.Response(w, resp, err)
	}
}
