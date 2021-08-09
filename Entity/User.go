package Entity

type User struct {
	IsCompleted bool   `gorm:"is_completed"        json:"is_completed"`
	CreatedAt   string `gorm:"created_at"          json:"created_at"`
	Picture     string `gorm:"picture"             json:"picture"`
	NickName    string `gorm:"column:name"         json:"nickname"`
	RealName    string `gorm:"column:private_name" json:"private_name"`
	Mail        string `gorm:"column:mail"         json:"mail"`
	ID          string `gorm:"column:user_id"      json:"user_id"       gorm:"primary_key"`
}
