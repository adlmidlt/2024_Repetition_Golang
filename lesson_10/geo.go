package lesson_10

import (
	"errors"
	"fmt"
	"log"
)

type Coordinates struct {
	latitude  float64
	longitude float64
}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude

	return nil
}

func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid longitude")
	}
	c.longitude = longitude

	return nil
}

func StartGeo() {
	coordinates := Coordinates{}
	err := coordinates.SetLongitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-1122.08)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
