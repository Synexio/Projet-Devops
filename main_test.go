// main_test.go
package main

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestFileCreation(t *testing.T) {

	// Date au format D_M_YYYY_Hh_Mm
	currentDate := strconv.Itoa(time.Now().Day()) + "_" + strconv.Itoa(int(time.Now().Month())) + "_" + strconv.Itoa(time.Now().Year()) + "_" + strconv.Itoa(time.Now().Hour()+2) + "h" + strconv.Itoa(time.Now().Minute()+3) + "m"

	filename := "Archive/" + currentDate + ".txt"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("File was not created in the Archive folder: %s", filename)
	}
}
