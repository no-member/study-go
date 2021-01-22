package geo

import "errors"

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
	if !(-90 <= latitude && latitude <= 90) {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}

func (c *Coordinates) SetLongitude(longitude float64) error {
	if !(-180 <= longitude && longitude <= 180) {
		return errors.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}

