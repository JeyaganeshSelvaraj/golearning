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

func TestFilterStudentString(t *testing.T) {
	s1 := "Jeyaganesh"
	s2 := "Selvaraj"
	students := []string{}
	students = addToSlice[string](s1, students)
	students = addToSlice[string](s2, students)
	t.Logf("%v", students)
	students = GenericFilter[string](students, func(item string) bool { return item == "Jeyaganesh" })
	t.Logf("%v", students)
}

func TestFilterStudentInt(t *testing.T) {
	s1 := 1
	s2 := 2
	students := []int{}
	students = addToSlice[int](s1, students)
	students = addToSlice[int](s2, students)
	t.Logf("%v", students)
	students = GenericFilter[int](students, func(item int) bool { return item == 1 })
	t.Logf("%v", students)
}

func TestFilterStudentStruct(t *testing.T) {
	s1 := types.Student{FirstName: "Jeyaganesh", LastName: "Selvaraj", ID: 1}
	s2 := types.Student{FirstName: "Jeyamurugan", LastName: "Selvaraj", ID: 2}
	students := []types.Student{}
	students = addToSlice[types.Student](s1, students)
	students = addToSlice[types.Student](s2, students)
	t.Logf("%v", students)
	students = GenericFilter[types.Student](students, func(item types.Student) bool { return item.FirstName == "Jeyaganesh" })
	t.Logf("%v", students)
}

func TestGenericMapInt(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := GenericMap[int](s1, func(item int) int { return item * item })
	t.Logf("%v", s2)
}
func TestGenericMapFloat(t *testing.T) {
	s1 := []float64{1.1, 2.1, 3.1, 4.1}
	s2 := GenericMap[float64](s1, func(item float64) float64 { return item * item })
	t.Logf("%v", s2)
}
