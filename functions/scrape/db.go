package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

func saveToDB(events []KidsEvent) error {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	// fmt.Println("Scrape data:", events)

	// For every event we're passed, we're going to loop over them and check the DB if they exist first.
	for _, event := range events {
		result := KidsEvent{
			Title: "No Match",
		}
		// TODO: `url` should be the key, but .Eq() on it doesn't seem to work.
		fmt.Println("----------")
		fmt.Println("Querying for event with title: ", event.Title)
		err := supabase.DB.From("events").Select("id, date, title, url, venue, display").Single().Eq("title", event.Title).Execute(&result)
		if err != nil {
			fmt.Println("Error!", err)
		}

		fmt.Println("Result after query: ", result)

		// If there are no results, then we'll save the event to the DB. Otherwise, we'll update the existing row.
		if result.Title == "No Match" {
			fmt.Println("Saving NEW row to DB", event)
			var saveResults []KidsEvent
			event.ID = uuid.New().String()
			err := supabase.DB.From("events").Insert(event).Execute(&saveResults)
			if err != nil {
				fmt.Println("Error while saving", err)
			}
		} else {
			var updateResults []KidsEvent
			err := supabase.DB.From("events").Update(event).Eq("title", event.Title).Execute(&updateResults)
			if err != nil {
				fmt.Println("Error while updating", err)
			}
			fmt.Println("Updated EXISTING row in DB", updateResults)
		}
	}

	return nil
}
