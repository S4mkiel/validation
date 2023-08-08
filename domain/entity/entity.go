package entity

type User struct {
	ID    int64  `json:"id" gorm:"primary_key;auto_increment"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Email string `json:"email" gorm:"type:varchar(255)"`
}
