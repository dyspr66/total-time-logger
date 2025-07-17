package main

import (
	"ttl/data"
	"ttl/logger"
)

func test() {
	acts := data.Activities{}
	acts = append(acts, &data.Activity{Name: "name1", Description: "desc1"})
	acts = append(acts, &data.Activity{Name: "name2", Description: "desc2"})
	acts = append(acts, &data.Activity{Name: "name3", Description: "desc3"})
	testTtl := logger.TotalTimeLogger{
		Activities: acts,
	}
	ttl = testTtl
}
