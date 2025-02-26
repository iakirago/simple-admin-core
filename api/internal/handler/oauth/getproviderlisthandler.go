package oauth

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/oauth"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route POST /oauth/provider/list oauth getProviderList
// Get provider list | 获取提供商列表
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PageInfo
// Responses:
//   200: ProviderListResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetProviderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewGetProviderListLogic(r.Context(), svcCtx)
		resp, err := l.GetProviderList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
