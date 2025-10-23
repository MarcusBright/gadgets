package handler

import (
	"directBTC/api/internal/logic"
	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func BindEvmAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BindEvmAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, response.ParamError(err.Error()))
			return
		}

		l := logic.NewBindEvmAddressLogic(r.Context(), svcCtx)
		resp, err := l.BindEvmAddress(&req)
		response.Response(w, resp, err)
	}
}
