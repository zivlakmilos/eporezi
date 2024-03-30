package scard

import "fmt"

type SubjectInfo struct {
	Name       string
	Surname    string
	PersonalId string
}

func (i SubjectInfo) String(string) {
	fmt.Sprintf("{ name: %s, surname: %s, personalId: %s}", i.Name, i.Surname, i.PersonalId)
}
