package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func (t *totalTimeLogger) readFromJson() error {
	b, err := os.ReadFile("record.json")
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	if len(b) == 0 {
		return nil
	}

	var acts []activity
	err = json.Unmarshal(b, &acts)
	if err != nil {
		return fmt.Errorf("unmarshaling from json: %w", err)
	}

	var a activities
	for _, v := range acts {
		a = append(a, &v)
	}

	ttl.activities = a

	return nil
}

func (t *totalTimeLogger) saveToJson() error {
	var acts []activity
	for _, v := range t.activities {
		acts = append(acts, *v)
	}

	// Marshal
	jsonData, err := json.Marshal(acts)
	if err != nil {
		return fmt.Errorf("marshaling to json: %w", err)
	}

	// Save
	err = os.WriteFile("record.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("writing data to json file: %w", err)
	}

	return nil
}
