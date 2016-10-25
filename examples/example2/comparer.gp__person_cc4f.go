///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Tue Oct 25 2016 23:13:45]
// Generate from:
//   [github.com/vipally/gogp/examples/example2/stack.gp_person.go]
//   [github.com/vipally/gogp/examples/example2/example2.gpg] [person]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct 24 2016 20:25:45]
// Version : 3.0.0.final
//
///////////////////////////////////////////////////////////////////

//this file is used to import by other gp files
//it cannot use independently
//simulation C++ stl functors
package example2

type ComparerPerson interface {
	F(left, right *Person) bool
}

type ComparerPersonCreator int

const (
	LESSER_Person ComparerPersonCreator = iota
	GREATER_Person
)

func (me ComparerPersonCreator) Create() (cmp ComparerPerson) {
	switch me {
	case LESSER_Person:
		cmp = LesserPerson(0)
	case GREATER_Person:
		cmp = GreaterPerson(0)
	}
	return
}

type LesserPerson byte

func (this LesserPerson) F(left, right *Person) (ok bool) {

	ok = left.Less(right)

	return
}

type GreaterPerson byte

func (this GreaterPerson) F(left, right *Person) (ok bool) {

	ok = left.Great(right)

	return
}
