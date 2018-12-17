package models

type Category struct {
	Id           int      `orm:"column(id);pk";json:"id"`
	Manufacturer string   `orm:"column(manufacturer);size(50)";json:"manufacturer"`
	Phones       []*Phone `orm:"reverse(many)";json:"phones"`
}

func (this *Category) TableName() string {
	return "category"
}
