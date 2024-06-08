package models

type AdminProfile struct {
	Username string `gorm:"primaryKey"`
	Password string
}

func (obj *AdminProfile) CreateProfile(username string, password string) {
	obj.Username = username
	obj.Password = password
}
