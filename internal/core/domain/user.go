package domain

type User struct {
	ID    uint   `gorm:"primaryKey;column:id" json:"id"`
	Name  string `gorm:"column:nombre" json:"nombre"`
	Email string `gorm:"column:correo" json:"correo"`
}
