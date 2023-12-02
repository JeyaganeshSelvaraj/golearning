package types

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	ID        int
}

func (s Student) String() string {
	return fmt.Sprintf("ID: %d, FirstName: %s, LastName: %s", s.ID, s.FirstName, s.LastName)
}
