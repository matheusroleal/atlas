/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:57:42
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-01 22:59:51
 */
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
