package po

import "time"

// 映射user表,存储用户信息
type User struct {
	Id           int64     `db:"id,pk"`
	UserId       string    `db:"user_id,uni"`
	UserName     string    `db:"user_name"`
	UserPassword string    `db:"user_password"`
	GmtCreated   time.Time `db:"gmt_created,omitempty"`
	GmtModified  time.Time `db:"gmt_modified,omitempty"`
}
