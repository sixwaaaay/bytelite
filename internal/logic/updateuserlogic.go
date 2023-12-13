/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/data"
	"github.com/sixwaaaay/shauser/user"
)

type UpdateUserLogic struct {
	conf        *config.Config
	userCommand *data.UserCommand
}

func NewUpdateUserLogic(conf *config.Config, userCommand *data.UserCommand) *UpdateUserLogic {
	return &UpdateUserLogic{conf: conf, userCommand: userCommand}
}

func (l *UpdateUserLogic) UpdateUser(ctx context.Context, in *user.UpdateUserRequest) (*user.UpdateUserReply, error) {
	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user id is empty")
	}
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "user name is empty")
	}
	var u data.User
	u.ID = in.UserId
	u.Username = in.Name
	u.Bio = in.Bio
	u.AvatarURL = in.AvatarUrl
	u.BgURL = in.BgUrl
	if err := l.userCommand.UpdateUser(ctx, &u); err != nil {
		return nil, err
	}
	return &user.UpdateUserReply{
		Profile: &user.User{
			Id:        u.ID,
			Name:      u.Username,
			Bio:       u.Bio,
			AvatarUrl: u.AvatarURL,
			BgUrl:     u.BgURL,
		},
	}, nil
}
