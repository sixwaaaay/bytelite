package logic

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"github.com/sixwaaaay/sharing/pkg/common/errorx"
	"github.com/sixwaaaay/sharing/pkg/common/secu"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic func(req *types.UserReq) (resp *types.UserResp, err error)

var NewLoginLogic = newLoginLogic

func newLoginLogic(ctx context.Context, appCtx *service.AppContext) LoginLogic {
	logger := logx.WithContext(ctx)
	return func(req *types.UserReq) (resp *types.UserResp, err error) {
		if len(req.Username) > 32 || len(req.Password) > 32 || len(req.Username) < 1 || len(req.Password) < 6 {
			return nil, errorx.NewDefaultError("username or password is too long")
		}
		if strings.ContainsAny(req.Username, " ") || strings.ContainsAny(req.Password, " ") {
			return nil, errorx.NewDefaultError("username contains space")
		}
		// 查询用户
		user, err := appCtx.UsersModel.FindOneByUsername(ctx, req.Username)
		if err != nil {
			logger.Errorf("find user by user failed, err: %v", err)
			return nil, errorx.NewDefaultCodeErr("user not found")
		}
		// 校验密码
		if compare := secu.Compare(req.Password, user.Password); compare != true {
			return nil, errorx.NewDefaultError("password is not correct")
		}
		// 生成token
		token, err := appCtx.JWTSigner.GenerateToken(user.Id, 60*60*24*7) // 7 days 有效期
		if err != nil {
			logger.Errorf("generate token failed, err: %v", err)
			return nil, errorx.NewDefaultCodeErr("generate token failed")
		}
		return &types.UserResp{
			StatusCode: 0,
			StatusMsg:  "",
			Token:      token,
			UserID:     user.Id,
		}, nil
	}
}
