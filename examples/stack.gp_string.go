///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sun Oct 23 2016 15:24:25]
// Generate from:
//   [github.com/vipally/gogp/examples/stack.gp]
//   [github.com/vipally/gogp/examples/example.gpg] [stack_string]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct  6 2016 14:25:07]
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

package examples



//stack object
type StringStack []string

//new object
func NewStringStack() *StringStack {
	return &StringStack{}
}

//push
func (this *StringStack) Push(v string) {
	*this = append(*this, v)
}

//pop
func (this *StringStack) Pop() (top string, ok bool) {
	if top, ok = this.Top(); ok {
		*this = (*this)[:this.Depth()-1]
	}
	return
}

//top
func (this *StringStack) Top() (top string, ok bool) {
	if this.Depth() > 0 {
		top = (*this)[this.Depth()-1]
		ok = true
	}
	return

}

//depth
func (this *StringStack) Depth() int {
	return len(*this)
}



