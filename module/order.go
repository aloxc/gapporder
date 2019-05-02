package module

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Order struct {
	Id          int32
	GoodsId     int32
	UserId      int32
	UserName    string    `orm"size(20)"`
	Address     string    `orm"size(100)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Version     int32
}

func InsertTestOrder() {
	order := &Order{
		Id:       1,
		GoodsId:  1,
		UserId:   1,
		UserName: "aloxc",
		Address:  "海南海口",
		Version:  1,
	}
	orm := orm.NewOrm()
	orm.Insert(order)
}
func GetOrder(order *Order) error {
	orm := orm.NewOrm()
	err := orm.Read(order)
	if err != nil {
		return err
	}
	return nil
}
