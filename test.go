package main

import (
	"fmt"
)

type user struct {
	name   string
	age    int
	height float32
	weight float32
}

type userMethods interface {
	strength(height, weight int) int
	hello(name string) string
}

func (u user) strength(height, weight int) int {
	strength := height * weight
	return strength
}

func (u user) hello(name string) string {
	userGreeting := fmt.Sprintf("Hello from vim %s", u.name)
	return userGreeting
}

func doSomething(s string, c chan string) {
	value := fmt.Sprintf("Hello from vim %s", s)
	c <- value
}
func main() {
	//me := user {
	//	name: "David",
	//	age: 21,
	//	height: 185.3,
	//	weight: 80.4,
	//}
	//fmt.Printf("You're %v joules strong", me.strength(int(me.height), int(me.weight)))
	//fmt.Printf("You're %v years old", me.age)
	//fmt.Printf("Hello %s", me.name)
	//fmt.Println(me.hello(me.name))
	//fmt.Println("Hello from my new configs which might be slowing down my vim tbh")

	myChannel := make(chan string)
	go doSomething("thatguy", myChannel)
	go doSomething("nnaemeka", myChannel)
	go doSomething("david", myChannel)

	x := <-myChannel
	fmt.Println(x)
	y := <-myChannel
	fmt.Println(y)
	z := <-myChannel
	fmt.Println(z)
	fmt.Println("go program ended")
}
