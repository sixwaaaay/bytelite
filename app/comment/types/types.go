package types

type CommentReq struct {
	UserId      int64  `form:"user_id" json:"user_id"`           // 用户id
	Token       string `form:"token" json:"token"`               // 用户token
	VideoId     int64  `form:"video_id" json:"video_id"`         // 视频id
	ActionType  int8   `form:"action_type" json:"action_type"`   // 1- 发布评论 2- 删除
	CommentText string `form:"comment_text" json:"comment_text"` // 评论内容
	CommentID   string `form:"comment_id" json:"comment_id"`     // 评论id
}

type CommentResp struct {
	Comment    Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
}

type CommentListReq struct {
	Token   string `form:"token" json:"token"`       // 用户token
	VideoId int64  `form:"video_id" json:"video_id"` // 视频id
}

type CommentListResp struct {
	CommentList []Comment `json:"comment_list"` // 评论列表
	StatusCode  int64     `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`   // 返回状态描述
}

// Comment 评论
type Comment struct {
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户信息
}

// User 评论用户信息
type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}
