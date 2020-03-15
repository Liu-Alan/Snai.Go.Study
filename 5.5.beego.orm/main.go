package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Userinfo struct {
	Id         int `orm:"column(uid)";PK`
	Username   string
	Department string
	Created    time.Time
}

func init() {
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:810618@/test?charset=utf8&parseTime=True&loc=Local", 30)

	//注册定义的model
	orm.RegisterModel(new(Userinfo))

	//创建table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	/*
		//插入数据
		user := Userinfo{Username: "slene", Department: "研发支撑", Created: time.Now()}

		id, err := o.Insert(&user)
		if err == nil {
			fmt.Println(id)
		}
	*/
	/*
		//更新数据
		user := Userinfo{Id: 4}

		if o.Read(&user) == nil {
			user.Username = "Amos"
			user.Created = time.Now()
			if num, err := o.Update(&user); err == nil {
				fmt.Println(num)
			}
		}
	*/
	/*
		//查询数据
		user := Userinfo{Id: 4}

		err := o.Read(&user)

		if err == orm.ErrNoRows {
			fmt.Println("查询不到")
		} else if err == orm.ErrMissPK {
			fmt.Println("找不到主键")
		} else {
			fmt.Println(user.Id, user.Username)
		}
	*/
	/*
		//查询数据集合
		var users []*Userinfo

		num, err := o.QueryTable("Userinfo").Filter("Id", 4).All(&users)
		if err == nil && num > 0 {
			for _, u := range users {
				fmt.Println(u.Id, u.Username, u.Created)
			}
		}
	*/
	/*
		//删除数据
		if num, err := o.Delete(&Userinfo{Id: 4}); err == nil {
			fmt.Println(num)
		}
	*/
	//批量删除
	num, err := o.QueryTable("Userinfo").Filter("ID__in", 4, 5).Delete()
	fmt.Printf("Affected Num:%d,%s\n", num, err)
}
