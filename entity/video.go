package entity

type Video struct {
	Title       string `gorm: "varchar(255)" json: "title"`
	Description string `gorm: "varchar(255)" json: "description"`
	URL         string `gorm: "varchar(255)" json: "url"`
}
