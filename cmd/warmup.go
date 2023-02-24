package cmd

import (
	"fmt"
	"github.com/Unkn0wnCat/calapi/internal/database"
	"github.com/Unkn0wnCat/calapi/internal/db_model"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// warmupCmd represents the warmup command
var warmupCmd = &cobra.Command{
	Use:   "warmup",
	Short: "Fills the database with random trash data - used for testing.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := database.Initialize()
		if err != nil {
			return err
		}
		defer database.Shutdown()

		now := time.Now()

		eventBox := db_model.BoxForEvent(database.ObjectBox)
		calendarBox := db_model.BoxForCalendar(database.ObjectBox)

		calendar1 := db_model.Calendar{
			Name:        "Caledar 1",
			Description: "This is a calendar used for testing database functionality. It contains daily events for a year starting " + now.String(),
			DateCreated: now,
		}
		calendar2 := db_model.Calendar{
			Name:        "Caledar 2",
			Description: "This is a calendar used for testing database functionality.",
			DateCreated: now,
		}

		day := time.Date(now.Year(), now.Month(), now.Day(), 12, 30, 00, 00, now.Location())

		for i := 0; i <= 365; i++ {
			event := db_model.Event{
				Title:        fmt.Sprintf("Daily Test %03d", i),
				Description:  fmt.Sprintf("This is a test event. %03d/365", i),
				Calendar:     &calendar1,
				LocationLat:  0,
				LocationLon:  0,
				LocationName: "",
				LocationAddr: "",
				Start:        day,
				End:          day.Add(time.Hour),
				DateCreated:  now,
			}

			id, err := eventBox.Put(&event)
			if err != nil {
				return err
			}
			log.Printf("successfully inserted event %d", id)

			day = day.Add(time.Hour * 24)
		}

		id, err := calendarBox.Put(&calendar2)
		log.Printf("successfully inserted calendar %d", id)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(warmupCmd)
}
