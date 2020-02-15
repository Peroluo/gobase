package main

import (
	"fmt"
	"os"
)

type Student struct {
	Name string
	Id   uint64
}

type Manager struct {
	Age        int
	AllStudent map[uint64]Student
}

// 查看学生
func (m Manager) ShowStudends() {
	for _, v := range m.AllStudent {
		fmt.Printf("学号：%d,姓名：%s\n", v.Id, v.Name)
	}
}

// 新增学生
func (m Manager) AddManger() {
	var (
		id   uint64
		name string
	)
	fmt.Println("请输入学生ID")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名")
	fmt.Scanln(&name)
	m.AllStudent[id] = Student{
		Name: name,
		Id:   id,
	}
	fmt.Println("新增成功！")
}

// 修改学生
func (m Manager) EditManger() {
	fmt.Println("输入修改的学生id")
	var (
		id uint64
	)
	fmt.Scanln(&id)
	_, ok := m.AllStudent[id]
	if !ok {
		fmt.Println("没有找到id")
	} else {
		var name string
		fmt.Println("输入修改的学生姓名")
		fmt.Scanln(&name)
		m.AllStudent[id] = Student{
			Name: name,
			Id:   id,
		}
		fmt.Println("已修改")
	}
}

// 删除学生
func (m Manager) DeleteManger() {
	fmt.Println("输入删除的学生id")
	var (
		id uint64
	)
	fmt.Scanln(&id)
	_, ok := m.AllStudent[id]
	if !ok {
		fmt.Println("没有找到id")
	} else {
		delete(m.AllStudent, id)
		fmt.Println("已修改")
	}
}

func showMenu() {
	fmt.Println("welcome sms")
	fmt.Println(`
		1.查看所有学生
		2.添加学生
		3.修改学生
		4.删除学生
		5.退出
	`)
}
func main() {
	showMenu()
	smr := Manager{
		AllStudent: make(map[uint64]Student),
		Age:        10,
	}
	for {
		showMenu()
		var todo int
		fmt.Println("请输入用户选项！")
		fmt.Scanln(&todo)
		switch todo {
		case 1:
			smr.ShowStudends()
		case 2:
			smr.AddManger()
		case 3:
			smr.EditManger()
		case 4:
			smr.DeleteManger()
		case 5:
			os.Exit(1)
		}
		fmt.Println(smr)
	}
}
