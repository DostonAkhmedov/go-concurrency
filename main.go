package main

import "time"

func main() {
	go sayHello()
	time.Sleep(time.Millisecond)
	println("world")
}

func sayHello()  {
	print("Hello ")
}
