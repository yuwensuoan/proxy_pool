package main

import (
	"proxy_pool/app/schedule"
	"proxy_pool/boostrap"
)

func main() {
	schedule.StartServer()
	boostrap.StartServer()
}
