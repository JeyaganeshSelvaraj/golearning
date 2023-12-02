package sorting

import (
	"golearning/types"
	"testing"
)

func TestAddStudent(t *testing.T) {
	s1 := types.Student{FirstName: "Jeyaganesh", LastName: "Selvaraj", ID: 1}
	s2 := types.Student{FirstName: "Jeyamurugan", LastName: "Selvaraj", ID: 2}
	students := []types.Student{s1, s2}
	PerformSort[types.Student](students, func(s1, s2 types.Student) bool { return s1.ID < s2.ID })
	t.Logf("%v", students)
}

func TestAddStudentString(t *testing.T) {
	s1 := "Jeyaganesh"
	s2 := "Selvaraj"
	students := []string{s1, s2}
	PerformSort[string](students, func(s1, s2 string) bool { return s1 < s2 })
	t.Logf("%v", students)
}

func TestAddStudentInt(t *testing.T) {
	s1 := 1
	s2 := 2
	students := []int{s1, s2}
	PerformSort[int](students, func(s1, s2 int) bool { return s1 < s2 })
	t.Logf("%v", students)
}
