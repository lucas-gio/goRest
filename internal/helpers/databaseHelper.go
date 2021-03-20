package goRest

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

type DatabaseHelper struct{}

func (d *DatabaseHelper) ProcessMax(max string) int64 {
	var defaultMax int64 = int64(25)
	var totalMax int64 = int64(1000)
	var maxInt int64
	var err error

	maxInt, err = strconv.ParseInt(max, 10, 64)
	if err != nil {
		log.Info("Invalid max obtained. Replaced by default value.")
		maxInt = defaultMax
	}

	if maxInt < 1 {
		log.Info("Zero or negative max obtained. Replaced by default value.")
		maxInt = defaultMax
	}

	if maxInt > totalMax {
		log.Info("Max value is bigger than highest value. Replaced by default max value.")
		maxInt = defaultMax
	}

	return maxInt
}

func (d *DatabaseHelper) ProcessOffset(offset string) int64 {
	var defaultOffset int64 = int64(0)
	var offsetInt int64
	var err error

	offsetInt, err = strconv.ParseInt(offset, 10, 64)

	if err != nil {
		log.Info("Invalid offset obtained. Replaced by default value.")
		offsetInt = defaultOffset
	}

	if offsetInt < 0 {
		log.Info("Negative offset obtained. Replaced by default value.")
		offsetInt = defaultOffset
	}

	return offsetInt
}
