package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeNowUTC(t *testing.T) {
	fmt.Println(time.Now().UTC())
}
func TestTimeNow(t *testing.T) {
	fmt.Println(time.Now())
}
