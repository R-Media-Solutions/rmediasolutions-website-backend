package models

type AboutUs struct {
	Id               int64  `gorm:"primaryKey" json:"id"`
	Name             string `gorm:"type:varchar(300)" json:"name"`
	ShortDescription string `gorm:"type:Text" json:"short_description"`
	Description      string `gorm:"type:Text" json:"description"`
	Status           int32  `gorm:"type:int(5)" json:"status"`
	CreatedAt        string `gorm:"type:datetime" json:"created_at"`
	ModifiedAt       string `gorm:"type:datetime" json:"modified_at"`
	DeletedAt        string `gorm:"type:datetime" json:"deleted_at"`
}
