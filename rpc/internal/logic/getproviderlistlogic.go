package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProviderListLogic {
	return &GetProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProviderListLogic) GetProviderList(in *core.PageInfoReq) (*core.ProviderListResp, error) {
	var providers []model.OauthProvider
	result := l.svcCtx.DB.Limit(int(in.PageSize)).Offset(int((in.Page - 1) * in.PageSize)).Find(&providers)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	resp := &core.ProviderListResp{}
	resp.Total = uint64(result.RowsAffected)
	for _, v := range providers {
		resp.Data = append(resp.Data, &core.ProviderInfo{
			Id:           uint64(v.ID),
			Name:         v.Name,
			ClientId:     v.ClientID,
			ClientSecret: v.ClientSecret,
			RedirectUrl:  v.RedirectURL,
			Scopes:       v.Scopes,
			AuthUrl:      v.AuthURL,
			TokenUrl:     v.TokenURL,
			AuthStyle:    uint64(v.AuthStyle),
			InfoUrl:      v.InfoURL,
			CreateAt:     v.CreatedAt.UnixMilli(),
		})
	}
	return resp, nil
}
