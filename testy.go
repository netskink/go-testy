package main

import "fmt"

type Project struct {
	Id    int64  `json:"project_id"`
	Title string `json:"title"`
	Name  string `json:"name"`
	//Data    Data    `json:"data"`
	//Commits Commits `json:"commits"`
}

var (
	Testy Project
)

func dump_project(foo Project) {
	fmt.Println("== Dump Project Struct ====")
	fmt.Printf("Id: %d\n", foo.Id)
	fmt.Println("Title: ", foo.Title)
	fmt.Printf("Name: %v\n", foo.Name)
}

func main() {
	fmt.Println("hello from go")
	Testy.Id = 3
	Testy.Title = "yo"
	Testy.Name = "my name"
	fmt.Println(Testy)
	dump_project(Testy)
}
