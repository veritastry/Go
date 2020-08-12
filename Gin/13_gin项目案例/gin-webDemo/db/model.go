package db

type TBL_USERS struct {
	ID    uint   `gorm:"primary_key; AUTO_INCREMENT"  json:"id,omitempty" example:"1"`
	Name  string `gorm:"unique_index; size:15"  json:"name,omitempty" example:"may"`
	Pwd   string `gorm:"default:'123456'; size:20"  json:"pwd,omitempty" example:"123"`
	Phone string `gorm:"size:11"  json:"phone,omitempty"`
	Email string `gorm:"size:25"  json:"email,omitempty"`
}

func (tbl_user *TBL_USERS) Insert() {
	gormConn := ConnGormMysql()
	gormConn.Create(&tbl_user)

}
