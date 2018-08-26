package abstract_factory

import "fmt"

/*
インタフェースを定義。
抽象的な工場（AbstractFactory）
抽象的な製品（Product1、Product2）
*/
type Product1 interface {
	Print1()
}

type Product2 interface {
	Print2()
}

type AbstractFactory interface {
	CreateProduct1() Product1
	CreateProduct2() Product2
}

/*
implementsを定義。
抽象を実装したA
*/
type ConcreteProductA1 struct {
}

func (concreteProductA1 *ConcreteProductA1) Print1() {
	fmt.Println("ConcreteProductA1")
}

type ConcreteProductA2 struct {
}

func (concreteProductA1 *ConcreteProductA2) Print2() {
	fmt.Println("ConcreteProductA2")
}

type ConcreteFactoryA struct {
}

func (factory *ConcreteFactoryA) CreateProduct1() Product1 {
	return &ConcreteProductA1{}
}

func (factory *ConcreteFactoryA) CreateProduct2() Product2 {
	return &ConcreteProductA2{}
}

/*
implementsを定義。
抽象を実装したB
*/
type ConcreteProductB1 struct {
}

func (concreteProductB1 *ConcreteProductB1) Print1() {
	fmt.Println("ConcreteProductB1")
}

type ConcreteProductB2 struct {
}

func (concreteProductB1 *ConcreteProductB2) Print2() {
	fmt.Println("ConcreteProductB2")
}

type ConcreteFactoryB struct {
}

func (factory *ConcreteFactoryB) CreateProduct1() Product1 {
	return &ConcreteProductB1{}
}

func (factory *ConcreteFactoryB) CreateProduct2() Product2 {
	return &ConcreteProductB2{}
}

/*
生成するためのNewConcreteFactory。
factoryTypeに入る引数によって、生成する工場を切り替える事で、より汎用的に使える。
新しく生成方法を追加したい場合は、ここにパターンを追加する事で、使う側の修正は少なく済む。
*/
func NewConcreteFactory(factoryType string) AbstractFactory {
	if factoryType == "A" {
		return &ConcreteFactoryA{}
	}
	if factoryType == "B" {
		return &ConcreteFactoryB{}
	}
	return nil
}

/*
使う側の処理を実装。
工場を生成して、工場にどの製品を作るかを頼む流れ
*/
func AbstractFactoryExecute(factoryType string) {
	factory := NewConcreteFactory(factoryType)
	product1 := factory.CreateProduct1()
	product2 := factory.CreateProduct2()

	product1.Print1()
	product2.Print2()
}
