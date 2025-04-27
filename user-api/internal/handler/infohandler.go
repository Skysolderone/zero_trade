package handler

import (
	"fmt"
	"net/http"

	"trade/user-api/internal/svc"
	"trade/user-api/util"
	"trade/user-rpc/pb/user"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		id := util.GetId(r.Context())
		res, err := svcCtx.UserRpc.GetUserInfo(r.Context(), &user.InfoReq{
			Id: id,
		})

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, user.UserInfo{
				Id:       res.Id,
				Username: res.Username,
				Password: res.Password,
			})
		}
	}
}
