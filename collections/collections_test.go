package generics

import (
	"golearning/types"
	"testing"
)

func TestAddStudent(t *testing.T) {
	s1 := types.Student{FirstName: "Jeyaganesh", LastName: "Selvaraj", ID: 1}
	s2 := types.Student{FirstName: "Jeyamurugan", LastName: "Selvaraj", ID: 2}
	students := []types.Student{}
	students = addToSlice[types.Student](s1, students)
	students = addToSlice[types.Student](s2, students)
	t.Logf("%v", students)
}

func TestAddStudentString(t *testing.T) {
	s1 := "Jeyaganesh"
	s2 := "Selvaraj"
	students := []string{}
	students = addToSlice[string](s1, students)
	students = addToSlice[string](s2, students)
	t.Logf("%v", students)
}

func TestAddStudentInt(t *testing.T) {
	s1 := 1
	s2 := 2
	students := []int{}
	students = addToSlice[int](s1, students)
	students = addToSlice[int](s2, students)
	t.Logf("%v", students)
}
