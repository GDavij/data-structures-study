package models

import "data-structures/dst"

type User struct {
	Id     int
	Name   string
	Stacks dst.LinkedList[string]
}
