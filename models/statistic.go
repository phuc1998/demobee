package models

type Statistic struct {
	Count    int       `orm:"column(count)"`
	Phone    *Phone    `orm:"rel(fk);column(phone_id)"`
	Discount *Discount `orm:"rel(fk);column(discount_id)"`
	Stt      int       `orm:"auto;column(stt);pk"`
}

func (this *Statistic) TableName() string {
	return "statistic"
}
