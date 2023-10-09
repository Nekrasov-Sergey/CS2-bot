package ds

type Session struct {
	ID          int    `db:"id"`
	UserVkID    string `db:"user_vk_id"`
	SessionDate string `db:"sessions_date"`
	Solutions   string `db:"solutions"`
}
