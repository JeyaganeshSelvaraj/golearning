package generics

import "testing"

func TestAddStudent(t *testing.T) {
	s1 := Student{FirstName: "Jeyaganesh", LastName: "Selvaraj", ID: 1}
	s2 := Student{FirstName: "Jeyamurugan", LastName: "Selvaraj", ID: 2}
	students := []Student{}
	students = addStudent[Student](s1, students)
	students = addStudent[Student](s2, students)
	t.Logf("%v", students)
}

func TestAddStudentString(t *testing.T) {
	s1 := "Jeyaganesh"
	s2 := "Selvaraj"
	students := []string{}
	students = addStudent[string](s1, students)
	students = addStudent[string](s2, students)
	t.Logf("%v", students)
}

func TestAddStudentInt(t *testing.T) {
	s1 := 1
	s2 := 2
	students := []int{}
	students = addStudent[int](s1, students)
	students = addStudent[int](s2, students)
	t.Logf("%v", students)
}
