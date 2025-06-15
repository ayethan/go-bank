package main

func main() {
	server := NewAPIServer(":3000")
	server.run()
	// fmt.Println("Hello, World!")
}
