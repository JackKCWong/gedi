package internal

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJsonLReader(t *testing.T) {
	Convey("jsonl reader can read jsonl content", t, func() {
		jsonlContent := `{"id": 1, "value": "hello"}
						 {"id": 2, "value": "world"}
`
		reader := JsonLReader{}
		rows, err := reader.Read(bytes.NewReader([]byte(jsonlContent)))

		So(err, ShouldBeNil)
		rec := next(rows).(Record)
		obj := rec.Parsed().(map[string]any)
		So(obj["id"], ShouldEqual, 1)
		So(obj["value"], ShouldEqual, "hello")

		rec = next(rows).(Record)
		obj = rec.Parsed().(map[string]any)
		So(obj["id"], ShouldEqual, 2)
		So(obj["value"], ShouldEqual, "world")
	})
}

func TestJsonReader(t *testing.T) {
	Convey("json reader can read json array content", t, func() {
		jsonArrayContent := `[
					{"id": 1, "value": "hello"},
					{"id": 2, "value": "world"}
				]`
		reader := JsonReader{}
		rows, err := reader.Read(bytes.NewReader([]byte(jsonArrayContent)))

		So(err, ShouldBeNil)
		rec := next(rows).(Record)
		obj := rec.Parsed().(map[string]any)
		So(obj["id"], ShouldEqual, 1)
		So(obj["value"], ShouldEqual, "hello")

		rec = next(rows).(Record)
		obj = rec.Parsed().(map[string]any)
		So(obj["id"], ShouldEqual, 2)
		So(obj["value"], ShouldEqual, "world")
	})
}
