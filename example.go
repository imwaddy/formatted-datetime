package main

import (
	"fmt"
	"time"

	fd "github.com/imwaddy/formatted-datetime/v2/formatteddatetime"
)

func main() {
	testTime := []string{"-", "/", " ", "."}
	for _, sep := range testTime {
		v := fd.TimeStruct{CurrentTime: time.Now()}
		fmt.Println(v.GetFormattedDatetime(sep))
		time.Sleep(1000 * time.Millisecond)
	}

	v := fd.TimeStruct{CurrentTime: time.Now()}
	fmt.Println("-----Only Date---------------")
	fmt.Println(v.GetFormattedDate("/"))

	fmt.Println("-----Only Time---------------")
	fmt.Println(v.GetFormattedTime())
}
