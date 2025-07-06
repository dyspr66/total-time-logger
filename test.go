package main

func test() {
	acts := make(map[int]activity)
	acts[1] = activity{name: "name1", description: "desc1"}
	acts[2] = activity{name: "name2", description: "desc2"}
	acts[3] = activity{name: "name3", description: "desc3"}
	testTtl := totalTimeLogger{
		activities: acts,
	}
	ttl = testTtl
}
