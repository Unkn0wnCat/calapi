package main

import "github.com/Unkn0wnCat/calapi/cmd"

//go:generate go run github.com/99designs/gqlgen generate

func main() {
	cmd.Execute()
}
