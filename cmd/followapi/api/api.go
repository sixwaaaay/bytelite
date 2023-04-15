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

package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"github.com/sixwaaaay/sharing/pkg/sign"
	"strconv"
)

type FollowApi struct {
	uc     pb.UserServiceClient
	secret []byte
}

func NewFollowApi(uc pb.UserServiceClient, secret string) *FollowApi {
	return &FollowApi{
		uc:     uc,
		secret: []byte(secret),
	}
}

func (u *FollowApi) Update(e *echo.Echo) {
	e.POST("/follow/following", u.Following, echo.WrapMiddleware(sign.Middleware(u.secret, false)))
	e.POST("/follow/followers", u.Followers, echo.WrapMiddleware(sign.Middleware(u.secret, false)))
}

func (h *FollowApi) subjectId(ctx echo.Context) (int64, error) {
	subjectID, _ := ctx.Request().Context().Value("x-id").(string)
	id, err := strconv.ParseInt(subjectID, 10, 64)
	return id, err
}
