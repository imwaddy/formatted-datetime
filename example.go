package main

import (
	"fmt"
	"time"

	fd "github.com/imwaddy/formatted-datetime/formatteddatetime"
)

func main() {
	testTime := []string{"-", "/", " ", "."}
	for i := 0; i < len(testTime); i++ {
		v := fd.TimeStruct{CurrentTime: time.Now()}
		fmt.Println(v.GetFormattedDatetime(testTime[i]))
		time.Sleep(1000 * time.Millisecond)
	}

	v := fd.TimeStruct{CurrentTime: time.Now()}
	fmt.Println("-----Only Date---------------")
	fmt.Println(v.GetFormattedDate("/"))

	fmt.Println("-----Only Time---------------")
	fmt.Println(v.GetFormattedTime())
}
