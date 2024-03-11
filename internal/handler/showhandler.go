package handler

import (
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shortener/internal/logic"
	"shortener/internal/svc"
	"shortener/internal/types"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 在handler中添加基于validator的参数校验
		validate := validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			// 返回的是一个重定向响应
			http.Redirect(w, r, resp.LongUrl, http.StatusFound)
		}
	}
}
