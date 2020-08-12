package bean

type User struct {
	Id       int    `json:"id",gorm:"auto-increment"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}