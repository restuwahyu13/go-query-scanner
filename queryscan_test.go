package queryscan

import (
	"fmt"
	"net/url"
	"testing"
)

type (
	TestStruct struct {
		Field1 string         `query:"field1"`
		Field2 int            `query:"field2"`
		Field3 bool           `query:"field3"`
		Field4 map[string]any `query:"field4"`
		Field5 float64        `query:"field5"`
		Field6 float32        `query:"field6"`
		Field7 int32          `query:"field7"`
		Field8 int64          `query:"field8"`
	}

	TestStruct2 struct {
		Field4 map[string]string `query:"field4"`
	}
)

func TestScan(t *testing.T) {
	values := url.Values{}
	values.Add("field1", "value1")
	values.Add("field2", "2")
	values.Add("field3", "true")
	values.Add("field4", "{\"key\":\"value\"}")
	values.Add("field5", "3.14")
	values.Add("field6", "3.14")
	values.Add("field7", "3")
	values.Add("field8", "3")

	queryStr := values.Encode()
	queryDest := TestStruct{}

	t.Run("Shoud be a valid query string success", func(t *testing.T) {
		err := Scan(queryStr, &queryDest)
		t.Log(err)

		if err != nil {
			t.FailNow()
			t.Errorf("Scan error: %v", err)
		}
	})

	t.Run("Shoud be a invalid dest query string not pointer failed", func(t *testing.T) {
		err := Scan(queryStr, queryDest)
		t.Log(err)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("Shoud be a invalid dest query string not struct failed", func(t *testing.T) {
		dest := make(map[string]string)
		err := Scan(queryStr, &dest)
		t.Log(err)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("Shoud be a query string valid json format", func(t *testing.T) {
		queryStr := "field4={\"key\":\"value\"}"
		dest := TestStruct{}

		err := Scan(queryStr, &dest)
		t.Log(err)

		if err != nil {
			t.FailNow()
		}
	})

	t.Run("Shoud be a query string invalid json format", func(t *testing.T) {
		queryStr := "field4=123"
		dest := TestStruct{}

		err := Scan(queryStr, &dest)
		t.Log(err)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("Shoud be a query string json but type not map invalid", func(t *testing.T) {
		queryStr := "field4={\"key\":\"value\"}"
		dest := TestStruct2{}

		err := Scan(queryStr, &dest)
		t.Log(err)

		if err == nil {
			t.FailNow()
		}
	})
}

func BenchmarkScan(b *testing.B) {
	values := url.Values{}
	values.Add("field1", "value1")
	values.Add("field2", "2")
	values.Add("field3", "true")
	values.Add("field4", "{\"key\":\"value\"}")
	values.Add("field5", "3.14")
	values.Add("field6", "3.14")
	values.Add("field7", "3")
	values.Add("field8", "3")

	queryStr := values.Encode()
	queryDest := TestStruct{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Scan(queryStr, &queryDest)
	}
}

func BenchmarkScanLargePayload(b *testing.B) {
	values := url.Values{}
	for i := 0; i < 100; i++ {
		values.Add(fmt.Sprintf("field1_%d", i), "value1")
		values.Add(fmt.Sprintf("field2_%d", i), "2")
		values.Add(fmt.Sprintf("field3_%d", i), "true")
		values.Add(fmt.Sprintf("field4_%d", i), "{\"key\":\"value\"}")
	}

	queryStr := values.Encode()
	queryDest := make(map[string]interface{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Scan(queryStr, &queryDest)
	}
}

func BenchmarkScanWithComplexJSON(b *testing.B) {
	values := url.Values{}
	values.Add("field4", "{\"key1\":{\"nested\":\"value\"},\"key2\":[1,2,3],\"key3\":true}")

	queryStr := values.Encode()
	queryDest := TestStruct{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Scan(queryStr, &queryDest)
	}
}
