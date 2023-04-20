package models

import "bubble/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo这个Model增删改查操作都放在这里

*/

// CreateATodo  创建todo
func CreateATodo(todo *Todo) (err error) {

	err = dao.DB.Create(&todo).Error
	return
}

/*
在Go语言中，slice 类型是引用类型，即它存储的是底层数组的地址。
所以当你声明一个 slice 类型的变量时，不需要为其分配存储空间，因为在使用前 slice 已经被初始化为 nil。
在这个函数中，todoList 变量被声明为一个 slice 类型的指针变量，这意味着该变量指向一个 slice 对象。
在函数中，通过调用 dao.DB.Find(&todoList) 语句，
将数据库中查询出的所有Todo记录填充到 todoList 所指向的 slice 对象中。
由于 slice 对象本身只存储了底层数组的地址，而不是数组本身，因此即使 slice 没有被显式地初始化，当向其添加元素时，它会自动为其分配足够的存储空间。
所以在 dao.DB.Find(&todoList) 执行之后，todoList 将包含查询到的所有 Todo记录
*/
func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

// 返回值定义的变量不会设置默认值，尤其是指针变量会出现野指针情况，所以要先初始化
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
