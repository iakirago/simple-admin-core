package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateDictionaryLogic {
	return &CreateOrUpdateDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateDictionaryLogic) CreateOrUpdateDictionary(req *types.CreateOrUpdateDictionaryReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateDictionary(context.Background(), &core.DictionaryInfo{
		Id:     req.Id,
		Title:  req.Title,
		Name:   req.Name,
		Status: req.Status,
		Desc:   req.Description,
	})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
