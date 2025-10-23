package handler

import (
	"directBTC/api/internal/logic"
	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTaskListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, response.ParamError(err.Error()))
			return
		}

		l := logic.NewGetTaskListLogic(r.Context(), svcCtx)
		resp, err := l.GetTaskList(&req)
		response.Response(w, resp, err)
	}
}
