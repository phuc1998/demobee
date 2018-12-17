package models

type Phone struct {
	Id          int         `orm:"column(id);pk"`
	Name        string      `orm:"column(name);size(100)"`
	Screen_size int         `orm:"column(screen_size)"`
	Category    *Category   `orm:"rel(fk);column(id_category);null;on_delete(set_null)"` //colume(id_category);
	Discount    []*Discount `orm:"reverse(many)"`
}

func (this *Phone) TableName() string {
	return "phone"
}
