package main

import m "github.com/yundream/simpleweb/lib/database"

func main() {
	db := new(m.Database)
	db.Open("../data")
	db.FindVcard("yundream.json")
}
