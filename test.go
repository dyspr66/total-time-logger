package main

func test() {
	acts := activities{}
	acts = append(acts, &activity{Name: "name1", Description: "desc1"})
	acts = append(acts, &activity{Name: "name2", Description: "desc2"})
	acts = append(acts, &activity{Name: "name3", Description: "desc3"})
	testTtl := totalTimeLogger{
		activities: acts,
	}
	ttl = testTtl
}
