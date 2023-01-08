package main

import (
	"os"

	supa "github.com/nedpals/supabase-go"
)

func getDataFromDB() ([]KidsEvent, error) {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	var results []KidsEvent

	err := supabase.DB.From("events").Select("*").Execute(&results)
	if err != nil {
		return []KidsEvent{}, err
	}

	// TODO: Filter out past events.

	return results, nil
}
