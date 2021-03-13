package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserGroup struct {
	gorm.Model
	Name             string
	DepartmentNumber uint
	Founder          string
}

type Department struct {
	gorm.Model
	Name           string
	PreDepartments string
	Subdivision    string
}

type User struct {
	gorm.Model
	Name      string
	NikeName  string
	phone     string
	Email     string
	Equipment string
	Source    string
	State     string
}

/*
 注意：
 1、创建数据库是结构体的名称对应数据库的表名 并且表名变为复数
 2、创建数据库的表的时候结构体中的属性对应字段名 结构体中的属性必须大写才能在数据库中创建相应字段
*/
func main() {
	//打开数据库  root:password@/gongsi?charset=utf8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/xiangmu?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return
	}
	//自动迁移模式
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Department{})
	db.AutoMigrate(&UserGroup{})
	//添加数据
	//user := User{
	//	Id:        1,
	//	Name:      "guocong",
	//	NikeName:  "郭聪",
	//	phone:     "134545",
	//	Email:     "134343@qq.com",
	//	Equipment: "研发",
	//	Source:    "手机",
	//	State:     "离线",
	//}
	user2 := User{
		//Id:        2,
		Name:      "join",
		NikeName:  "约翰",
		phone:     "56753",
		Email:     "223223@qq.com",
		Equipment: "测试",
		Source:    "电脑",
		State:     "上线",
	}
	department1 := Department{
		//Id:             1,
		Name:           "研发部门",
		PreDepartments: "总部",
		Subdivision:    "go开发",
	}
	usergroup1 := UserGroup{
		//Id:               1,
		Name: "表格导入",
		//UsersNumber:      4,
		DepartmentNumber: 4,
		//UpdataTime:       "2021-02-04",
		Founder: "admit",
	}
	//添加操作
	db.Create(&user2)
	//fmt.Println(user2)
	db.Create(&department1)
	//fmt.Println(department1)
	db.Create(&usergroup1)
	fmt.Println(usergroup1)
	//查询数据
	//db.First(&user2, 2)
	//fmt.Println(user2)
	//db.First(&department1, 1)
	//fmt.Println(department1)
	//db.First(&usergroup1, 1)
	//fmt.Println(usergroup1)
	//修改数据
	//修改用户的名字
	//user2.NikeName= "麦克"
	//修改用户的状态
	//user2.State="离线"
	//db.Model(&user2).Update(user2)
	//fmt.Println(user2)
	//修改部门的名字
	//department1.Name= "测试部门"
	//修改下级部门的名字
	//department1.Subdivision = "java开发"
	//db.Model(&department1).Update(department1)
	//fmt.Println(department1)
	//修改用户组的名字
	//usergroup1.Name= "本地导入"
	//修改用户组发起人
	//usergroup1.Founder = "root"
	//db.Model(&usergroup1).Update(usergroup1)
	//fmt.Println(usergroup1)
	//删除数据
	//db.Delete(&user2)
	//db.Delete(&department1)
	//db.Delete(&usergroup1)
	//关闭数据库
	//db.Close()
}
