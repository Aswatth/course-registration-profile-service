package models

type StudentProfile struct {
	Email_id         string `gorm:"primaryKey"`
	Login            Login  `gorm:"foreignKey:Email_id;references:Email_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	First_name       string
	Last_name        string
	Program_enrolled string
}
