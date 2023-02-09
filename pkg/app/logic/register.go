package logic

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/app/dal"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"github.com/sixwaaaay/sharing/pkg/common/errorx"
	"github.com/sixwaaaay/sharing/pkg/common/secu"
	"strings"
)

type RegisterLogic func(req *types.UserReq) (resp *types.UserResp, err error)

var NewRegisterLogic = newRegisterLogic

func newRegisterLogic(ctx context.Context, appCtx *service.AppContext) RegisterLogic {
	return func(req *types.UserReq) (resp *types.UserResp, err error) {
		if len(req.Username) > 32 || len(req.Password) > 32 || len(req.Username) < 1 || len(req.Password) < 6 {
			return nil, errorx.NewDefaultError("username or password is too long")
		}
		if strings.ContainsAny(req.Username, " ") || strings.ContainsAny(req.Password, " ") {
			return nil, errorx.NewDefaultError("username or password contains space")
		}
		user := dal.User{
			Username: req.Username,
			Password: secu.Hash(req.Password),
		}
		// 写入数据库
		ret, err := appCtx.UsersModel.Insert(ctx, &user)
		if err != nil {
			return nil, errorx.NewDefaultError("register failed, server error")
		}
		userId, err := ret.LastInsertId() // move to time increasingly unique id
		if err != nil {
			return nil, errorx.NewDefaultError("server busy, try later")
		}
		// 生成token
		token, err := appCtx.JWTSigner.GenerateToken(userId, 60*60*24*7)
		if err != nil {
			return nil, errorx.NewDefaultError("server busy, please login manual")
		}
		return &types.UserResp{
			StatusCode: 0,
			StatusMsg:  "",
			Token:      token,
			UserID:     userId,
		}, nil
	}
}
