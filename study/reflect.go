package study

import (
	"fmt"
	"reflect"
)

func Myreflect(){
	user:=User{1,"cb",25}
	getvalue:=reflect.ValueOf(user)
	method:=getvalue.MethodByName("ReflectCallFuncHasArgs")
	args:=[]reflect.Value{reflect.ValueOf("cb2"),reflect.ValueOf(26)}
	method.Call(args)

	method=getvalue.MethodByName("ReflectCallFunc")
	args=make([]reflect.Value,0)
	method.Call(args)
	//a:=1
	////type_ :=reflect.TypeOf(a)
	//pointer_:=reflect.ValueOf(&a)
	//newvalue:=pointer_.Elem()
	//
	//fmt.Println(pointer_)
	//fmt.Println(newvalue)
	//fmt.Println(newvalue.Type())
	//fmt.Println(newvalue.CanSet())
	//
	//newvalue.SetInt(77)
	//fmt.Println(a)
	//convtype:=type_.Interface().(*int)
	//convvalue:=value_.Interface().(int)
	//fmt.Println(convtype)
	//fmt.Println(convvalue)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	u:=User{1,"cb",25}
	DoFieldAndMethod(u)
}

type User struct {
	Id int
	Name string
	Age int
}
func (u User) ReflectCallFunc(){
	fmt.Println("ReflectCallFunc")
}
func (u User) ReflectCallFuncHasArgs(name string,age int){
	fmt.Println("name ",name,"     age",age)
}
func DoFieldAndMethod(input interface{}){
	gettype:=reflect.TypeOf(input)
	fmt.Println("gettype is ",gettype.Name())
	getvalue:=reflect.ValueOf(input)
	fmt.Println("gettype is ",getvalue)
	for i:=0;i<gettype.NumField();i++{
		field:=gettype.Field(i)
		value:=getvalue.Field(i).Interface()
		fmt.Printf("%s   %s    %s\n",field.Name,field.Type,value)
	}
	for i:=0;i<gettype.NumMethod();i++{
		m:=gettype.Method(i)
		fmt.Printf("%s  %v",m.Type,m.Name)
	}
}