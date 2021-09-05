# formatted-datetime

formatted-datetime is special service written for accessing some handy dates and time.
This service will return time date with some formatting.
This service has a single parameter i.e. Format string which is any of 4  dash(-), slash(/), space( ), and dot(.).


*INSTALLATION*

```
go get github.com/imwaddy/formatted-datetime/formatteddatetime

````

*USAGE*

<b> v := fd.TimeStruct{CurrentTime: time.Now()} </b>

following will return specified output (Use any).

<b> v.GetFormattedDatetime("/") </b> <br>
<b> v.GetFormattedDatetime("-") </b> <br>
<b> v.GetFormattedDatetime(" ") </b> <br>
<b> v.GetFormattedDatetime(".") </b> <br>

` Sample Program `
```
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

```

* <b> OUTPUT </b>
```
29-Mar-2021 11:24:19 PM
29/Mar/2021 11:24:20 PM
29 Mar 2021 11:24:21 PM
29.Mar.2021 11:24:22 PM
-----Only Date---------------
29/Mar/2021
-----Only Time---------------
11:24:23 PM
```

` Benchmarking `

```
goos: windows
goarch: amd64
pkg: github.com/imwaddy/formatted-datetime/formatteddatetime
cpu: Intel(R) Core(TM) i7-10610U CPU @ 1.80GHz
Benchmark_Package/Run_bGetFormattedDateTime-8         	 4784924	       290.8 ns/op	      16 B/op	       1 allocs/op
Benchmark_Package/Run_bGetFormattedTime-8             	 3307118	       330.1 ns/op	      16 B/op	       1 allocs/op
Benchmark_Package/Run_bGetFormattedDate-8             	 3366686	       346.9 ns/op	      16 B/op	       1 allocs/op

```

## License

formatted-datetime source code is available under the MIT [License](/LICENSE).