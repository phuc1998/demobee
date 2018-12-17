package models

type Discount struct {
	Id      int      `orm:"column(id);pk"`
	Percent int      `orm:"column(percent)"`
	Phones  []*Phone `orm:"rel(m2m); rel_through(demobee/models.Statistic)"`
}
