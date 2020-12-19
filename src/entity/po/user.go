package po

import "time"

// 映射user表,存储用户信息
type User struct {
	/**
	id
	int(10) UN zerofill AI PK
	user_id
	varchar(32)
	user_name
	varchar(32)
	user_password
	varchar(32)
	gmt_created
	datetime(3)
	gmt_modified
	datetime(3)
	*/
	Id           int64     `db:"id,pk"`
	UserId       string    `db:"user_id,uni"`
	UserName     string    `db:"user_name"`
	UserPassword string    `db:"user_password"`
	GmtCreated   time.Time `db:"gmt_created,omitempty"`
	GmtModified  time.Time `db:"gmt_modified,omitempty"`
}
