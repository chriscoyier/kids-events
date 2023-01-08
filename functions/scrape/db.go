package main

import (
	"fmt"
	"os"

	supa "github.com/nedpals/supabase-go"
)

func saveToDB(events []KidsEvent) error {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	// For every event we're passed, we're going to loop over them and check the DB if they exist first.
	for _, event := range events {
		var results []KidsEvent
		// Can't get `url` to match??? Using `title` for now.
		err := supabase.DB.From("events").Select("id, date, title, url, venue, display").Single().Eq("title", event.Title).Execute(&results)
		if err != nil {
			fmt.Println("Error!", err)
		}

		// If there are no results, then we'll save the event to the DB. Otherwise, we'll update the existing row.
		if len(results) == 0 {
			// Save new event to DB
			var saveResults []KidsEvent
			err := supabase.DB.From("events").Insert(event).Execute(&saveResults)
			if err != nil {
				fmt.Println("Error while saving", err)
			}
		} else {
			// Update existing row in DB
			var updateResults []KidsEvent
			err := supabase.DB.From("events").Update(event).Eq("title", event.Title).Execute(&updateResults)
			if err != nil {
				fmt.Println("Error while updating", err)
			}
		}
	}

	return nil
}
