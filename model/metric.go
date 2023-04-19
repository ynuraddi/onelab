package model

type Metric struct {
	UserID   int      `json:"user_id"        gorm:"column:user_id"`
	UserName string   `json:"user_name"      gorm:"column:name"`
	Books    []string `json:"books"          gorm:"column:books"`
}

func (Metric) TableName() string {
	return "metric"
}
