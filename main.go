package main

import "os"

func main() {
	a := App{}
	a.Initialize(
        "postgres",
        "issecure",
        "postgres")

	println("init done")
	a.Run(":8010")
}