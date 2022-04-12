package test

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("BULK INSERTION TEST")
	bulkInsertionTest()

	time.Sleep(15 * time.Second)

	log.Println("HASH INSERTION TEST")
	hashInsertionTest()

	time.Sleep(15 * time.Second)

	log.Println("SEARCH TEST")
	searchInsertionTest()

	time.Sleep(15 * time.Second)
}
