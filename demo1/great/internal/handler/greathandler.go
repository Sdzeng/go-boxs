package handler

import (
	"net/http"

	"demo1/great/internal/logic"
	"demo1/great/internal/svc"
	"demo1/great/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GreatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGreatLogic(r.Context(), svcCtx)
		resp, err := l.Great(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
