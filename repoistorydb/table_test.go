package repoistorydb

import (
	"fmt"
	"testing"
)

type Course struct {
	name   string
	points float32
	grade  float32
}
type Student struct {
	id      int `constraint: primary`
	name    string
	grades  []float32
	courses []Course
}

func (Student) newTableObject() interface{} {
	return Student{1, "elad", []float32{50.0, 62.0}, []Course{}}
}
func (Course) newTableObject() interface{} {
	return Course{"math", 2.5, 100.0}
}

func TestCreate(t *testing.T) {

	var rep Repoistory
	rep.Init()
	s := Student{1, "elad", []float32{50.0, 62.0}, []Course{}}
	c := Course{"math", 2.5, 100.0}
	rep.CreateTable(c)
	rep.CreateTable(Student{})
	for _, v := range generateSqlCreateTables(&rep) {
		fmt.Println(v)
	}
	for _, v := range generateSqlCreateRelations(&rep) {
		fmt.Println(v)
	}
	fmt.Println(rep.tabels["Course"].InsertCommand(c))
	fmt.Println(rep.tabels["Student"].InsertCommand(s))

}
