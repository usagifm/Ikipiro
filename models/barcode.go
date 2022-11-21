package models

type Barcode struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Barcode string `gorm:"type:varchar(20)" json:"barcode"`
	Nama    string `gorm:"type:varchar(255)" json:"nama"`
}
