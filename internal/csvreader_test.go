package internal

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCsvReader(t *testing.T) {
	Convey("csv reader can read csv file", t, func() {
		csvContent := `1,"hello,world"
2,hi`
		reader := CsvReader{}
		rows, err := reader.Read(bytes.NewReader([]byte(csvContent)))

		So(err, ShouldBeNil)
		rec := next(rows).(Record)
		So(rec.String(), ShouldEqual, `1,"hello,world"`)
		So(rec.Parsed()["x"].([]string)[0], ShouldEqual, "1")
		So(rec.Parsed()["x"].([]string)[1], ShouldEqual, "hello,world")

		rec = next(rows).(Record)
		So(rec.String(), ShouldEqual, `2,hi`)
		So(rec.Parsed()["x"].([]string)[0], ShouldEqual, "2")
		So(rec.Parsed()["x"].([]string)[1], ShouldEqual, "hi")
	})
}
