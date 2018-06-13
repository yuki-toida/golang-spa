package todo

import (
	"sort"
	"time"

	"github.com/jinzhu/gorm"
)

// Todo struct
type Todo struct {
	ID        int `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// All func
func All() []Todo {
	todos := []Todo{}
	db := connectDB()
	db.Find(&todos)
	return todos
}

// Get func
func Get(id int) Todo {
	todo := Todo{}
	db := connectDB()
	db.First(&todo, id)
	return todo
}

// Create func
func Create(name string) []Todo {
	all := All()
	nextID := nextID(all)
	now := time.Now()
	todo := Todo{ID: nextID, Name: name, CreatedAt: now, UpdatedAt: now}
	db := connectDB()
	db.Create(&todo)
	result := append(all, todo)
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func nextID(all []Todo) int {
	for i, v := range all {
		if i+1 != v.ID {
			return i + 1
		}
	}
	return len(all) + 1
}

// Update func
func Update(id int, name string) Todo {
	todo := Get(id)
	if todo == (Todo{}) {
		return Todo{}
	}
	db := connectDB()
	db.Model(&todo).Update("name", name)
	return todo
}

// Delete func
func Delete(id int) bool {
	todo := Get(id)
	if todo == (Todo{}) {
		return false
	}
	db := connectDB()
	db.Delete(&todo)
	return true
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:zaqroot@tcp(127.0.0.1:9306)/todo_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	return db
}
