package model

import (
	"github.com/Unkn0wnCat/calapi/internal/db_model"
	"strconv"
)

func FromEvent(event db_model.Event) Event {
	modelEvent := Event{
		ID:          strconv.FormatUint(event.Id, 16),
		Title:       event.Title,
		Location:    nil,
		Description: event.Description,
		Start:       event.Start,
		End:         event.End,
		DbCalendar:  event.Calendar,
	}

	if (event.LocationLat != 0 && event.LocationLon != 0) || event.LocationName != "" || event.LocationAddr != "" {
		modelEvent.Location = &Location{}

		if event.LocationLat != 0 && event.LocationLon != 0 {
			modelEvent.Location.Lat = &event.LocationLat
			modelEvent.Location.Lon = &event.LocationLon
		}

		if event.LocationName != "" {
			modelEvent.Location.Name = &event.LocationName
		}

		if event.LocationAddr != "" {
			modelEvent.Location.Address = &event.LocationAddr
		}
	}

	return modelEvent
}

func FromCalendar(calendar db_model.Calendar) Calendar {
	return Calendar{
		ID:          strconv.FormatUint(calendar.Id, 16),
		DbID:        calendar.Id,
		Name:        calendar.Name,
		Description: calendar.Description,
	}
}
