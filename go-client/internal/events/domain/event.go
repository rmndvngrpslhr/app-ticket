package domain

import (
	"time"

	"github.com/google/uuid"
)

// Validate checks if the event data is valid.
func (e Event) Validate() error {
	if len(e.Name) == 0 {
		return ErrEventNameRequired
	}
	if e.Date.Before(time.Now()) {
		return ErrEventFutureDateRequired
	}
	if e.Capacity <= 0 {
		return ErrEventCapacityInvalid
	}
	if e.Price <= 0 {
		return ErrEventPriceInvalid
	}

	return nil
}

// AddSpot adds a spot to the event.
func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}

	e.Spots = append(e.Spots, *spot)
	return spot, nil
}


// NewEvent creates a new event with the given parameters.
func NewEvent(
	name, location, organization, imageURL string,
	rating Rating,
	date time.Time,
	capacity, partnerID int,
	price float64,
) (*Event, error) {
	event := &Event{
		ID:           uuid.New().String(),
		Name:         name,
		Location:     location,
		Organization: organization,
		Rating:       rating,
		Date:         date,
		Capacity:     capacity,
		Price:        price,
		ImageURL:     imageURL,
		PartnerID:    partnerID,
		Spots:        make([]Spot, 0),
	}
	if err := event.Validate(); err != nil {
		return nil, err
	}
	return event, nil
}
