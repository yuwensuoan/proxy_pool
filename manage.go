package main

import (
	"proxy_pool/app/schedule"
)

func main() {

	job := &schedule.Job{}
	job.Run()
	//schedule.StartServer()
	//boostrap.Server.Run(":8080")
}
