package main

import "case-api/api"

func main() {
	api.Init()
	api.App.Start()
}
