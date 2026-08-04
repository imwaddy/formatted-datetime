# formatted-datetime

formatted-datetime is special service written for accessing some handy dates and time.
This service will return time date with some formatting.
This service has a single parameter i.e. Format string which is any of 4  dash(-), slash(/), space( ), and dot(.).


*INSTALLATION*

```
go get github.com/imwaddy/formatted-datetime/v2/formatteddatetime

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
goos: darwin
goarch: arm64
pkg: github.com/imwaddy/formatted-datetime/v2/formatteddatetime
cpu: Apple M2 Pro
Benchmark_GetFormattedDateTime-12    	55128481	        21.48 ns/op	      16 B/op	       1 allocs/op
Benchmark_GetFormattedTime-12        	54073033	        22.68 ns/op	      16 B/op	       1 allocs/op
Benchmark_GetFormattedDate-12        	55455535	        21.53 ns/op	      16 B/op	       1 allocs/op
```

## License

formatted-datetime source code is available under the MIT [License](/LICENSE).