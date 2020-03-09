package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {
	currentTime := time.Now()

	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal("Can not receive time with error: ", err)
	}

	const layout = "2006-01-02 15:04:05.000000000"
	fmt.Println("current time: ", currentTime.Format(layout))
	fmt.Println("exact time: ", time.local().Format(layout))
}
