package logprint

import (
	"fmt"
	"time"
)

func Debug(msg interface{}) {
	t := time.Now()
	fmt.Printf("[debug] %s: %s\n", t.Format("2006-01-02 15:04:05.00"), msg)

}

func Info(msg interface{}) {
	t := time.Now()
	fmt.Printf("[Info] %s: %s\n", t.Format("2006-01-02 15:04:05.00"), msg)

}
