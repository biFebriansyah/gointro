package module

import "fmt"

type IFCtrl_user interface {
	GetData() string
	SetData()
}

type ctrl_user struct {
	db string
}

func (c *ctrl_user) GetData() string {
	return c.db
}

func (c *ctrl_user) SetData() {
	fmt.Println(c.db)
}

func ExamInterface() {
	ctrl := ctrl_user{db: "hellow"}
	userss(&ctrl)

	Anytest(2, 3)
}

func userss(ifc IFCtrl_user) {
	ifc.GetData()
}

// type assertion
func Anytest(a1, a2 interface{}) {
	value, test := a1.(int)
	if !test {
		fmt.Println("data bukan int")
	}
	values := a1.(int) + a2.(int)

	fmt.Println(value)
	fmt.Println(values)
}
