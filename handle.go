package main

import (
	"fmt"
	"net/http"
)

func handleAddActivity(w http.ResponseWriter, r *http.Request) {
	var newActivity activity

	// TODO - validate
	name := r.FormValue("name")
	description := r.FormValue("description")

	newActivity.name = name
	newActivity.description = description

	ttl.activities[len(ttl.activities)+1] = newActivity

	// TODO - update data storage
	fmt.Fprint(w, "Success!")
}
