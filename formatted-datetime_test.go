package formatteddatetime

import (
	"testing"
	"time"
)

var (
	time1 time.Time
)

func init() {
	str := "2020-04-11T11:45:26.371Z"
	time1, _ = time.Parse(time.RFC3339, str)
}

func Test_GetFormattedDate(t *testing.T) {
	type fields struct {
		CurrentTime time.Time
	}
	type args struct {
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Test for /", fields{CurrentTime: time1}, args{format: "/"}, "11/Apr/2020"},
		{"Test for -", fields{CurrentTime: time1}, args{format: "-"}, "11-Apr-2020"},
		{"Test for ` `", fields{CurrentTime: time1}, args{format: " "}, "11 Apr 2020"},
		{"Test for .", fields{CurrentTime: time1}, args{format: "."}, "11.Apr.2020"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tst := &TimeStruct{
				CurrentTime: tt.fields.CurrentTime,
			}
			if got := tst.GetFormattedDate(tt.args.format); got != tt.want {
				t.Errorf("TimeStruct.GetFormattedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetFormattedTime(t *testing.T) {
	type fields struct {
		CurrentTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Time", fields{CurrentTime: time1}, "05:15:26 PM"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TimeStruct{
				CurrentTime: tt.fields.CurrentTime,
			}
			if got := ts.GetFormattedTime(); got != tt.want {
				t.Errorf("TimeStruct.GetFormattedTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetFormattedDateTime(t *testing.T) {
	type fields struct {
		CurrentTime time.Time
	}
	type args struct {
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Test for /", fields{CurrentTime: time1}, args{format: "/"}, "11/Apr/2020 05:15:26 PM"},
		{"Test for -", fields{CurrentTime: time1}, args{format: "-"}, "11-Apr-2020 05:15:26 PM"},
		{"Test for ` `", fields{CurrentTime: time1}, args{format: " "}, "11 Apr 2020 05:15:26 PM"},
		{"Test for .", fields{CurrentTime: time1}, args{format: "."}, "11.Apr.2020 05:15:26 PM"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tst := &TimeStruct{
				CurrentTime: tt.fields.CurrentTime,
			}
			if got := tst.GetFormattedDatetime(tt.args.format); got != tt.want {
				t.Errorf("TimeStruct.GetFormattedDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_GetFormattedDateTime(b *testing.B) {
	// run the GetFormattedDate function b.N times
	v := TimeStruct{CurrentTime: time.Now()}
	for n := 0; n < b.N; n++ {
		v.GetFormattedDate("/")
	}
}

func Benchmark_GetFormattedTime(b *testing.B) {
	// run the GetFormattedTime function b.N times
	v := TimeStruct{CurrentTime: time.Now()}
	for n := 0; n < b.N; n++ {
		v.GetFormattedTime()
	}
}

func Benchmark_GetFormattedDate(b *testing.B) {
	// run the GetFormattedDate function b.N times
	v := TimeStruct{CurrentTime: time.Now()}
	for n := 0; n < b.N; n++ {
		v.GetFormattedDate("/")
	}
}
