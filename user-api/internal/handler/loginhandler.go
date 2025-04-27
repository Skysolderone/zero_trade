package handler

import (
	"net/http"

	"trade/user-api/internal/svc"
	"trade/user-api/internal/types"
	"trade/user-rpc/pb/user"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		res, err := svcCtx.UserRpc.Login(r.Context(), &user.LoginReq{
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, res)
	}
}
