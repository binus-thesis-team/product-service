package utils

import "time"

func SetTimeLocation(timeLocation string) (*time.Location, error) {
	return time.LoadLocation(timeLocation)
}
