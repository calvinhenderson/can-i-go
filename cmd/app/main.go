package main

import serviceapi "can-i-go/app/service-api"

func main() {
	api := serviceapi.New()
	api.Start()
}
