package formatteddatetime

import (
	"reflect"
	"testing"
	"time"

	"github.com/araddon/dateparse"
)

func GetDate() time.Time {
	str := "2020-04-11T11:45:26.371Z"
	t, _ := time.Parse(time.RFC3339, str)
	return t
}

func TestTimeStruct_GetFormattedDate(t *testing.T) {
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
		{"Test for /", fields{CurrentTime: GetDate()}, args{format: "/"}, "11/Apr/2020"},
		{"Test for -", fields{CurrentTime: GetDate()}, args{format: "-"}, "11-Apr-2020"},
		{"Test for ` `", fields{CurrentTime: GetDate()}, args{format: " "}, "11 Apr 2020"},
		{"Test for .", fields{CurrentTime: GetDate()}, args{format: "."}, "11.Apr.2020"},
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

func TestTimeStruct_GetFormattedTime(t *testing.T) {
	type fields struct {
		CurrentTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Time", fields{CurrentTime: GetDate()}, "05:15:26 PM"},
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

func TestConvertUTCFormatDate(t *testing.T) {
	res, _ := dateparse.ParseLocal("2020-04-11 11:45:26.371Z")
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"Convert UTC", args{date: "2020-04-11 11:45:26.371Z"}, res, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertUTCFormatDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertUTCFormatDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertUTCFormatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// from ConvertUTCFormatDate
func BenchmarkConvertUTCFormatDate(b *testing.B) {
	// run the ConvertUTCFormatDate function b.N times
	for n := 0; n < b.N; n++ {
		ConvertUTCFormatDate("2020-04-11 11:45:26.371Z")
	}
}

func BenchmarkGetFormattedDateTime(b *testing.B) {
	// run the GetFormattedDate function b.N times
	v := TimeStruct{CurrentTime: time.Now()}
	for n := 0; n < b.N; n++ {
		v.GetFormattedDate("/")
	}
}
func BenchmarkGetFormattedTime(b *testing.B) {
	// run the GetFormattedTime function b.N times
	v := TimeStruct{CurrentTime: time.Now()}
	for n := 0; n < b.N; n++ {
		v.GetFormattedTime()
	}
}

func BenchmarkGetFormattedDate(b *testing.B) {
	// run the GetFormattedDate function b.N times
	v := TimeStruct{CurrentTime: time.Now()}
	for n := 0; n < b.N; n++ {
		v.GetFormattedDate("/")
	}
}
