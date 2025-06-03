## Go Query Scanner

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/queryscan?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/restuwahyu13/queryscan)](https://goreportcard.com/report/github.com/restuwahyu13/queryscan) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/restuwahyu13/queryscan/blob/master/CONTRIBUTING.md)

**queryscan** is a lightweight, pure Go scanner implemented without external dependencies. This package is designed to facilitate the transformation of query string requests into Go struct formats. It was developed as a solution to challenges encountered while i using the `gorilla schema` package for similar conversions. Although the schema structures were correctly defined when using `gorilla schema`, persistent errors indicated that the schema was invalid. These issues motivated the creation of **queryscan** as a reliable alternative.

- [Go Query Scanner](#queryscan)
  - [Installation](#installation)
  - [Example Usage](#example-usage)
- [Testing](#testing)
  - [Bugs](#bugs)
  - [Contributing](#contributing)
  - [License](#license)

### Installation

```sh
get github.com/restuwahyu13/queryscan
```

### Example Usage

```go
  package main

  import (
    "fmt"
	"net/url"

	"github.com/restuwahyu13/queryscan"
  )

type TestStructx struct {
	Field1 string         `query:"field1"`
	Field2 int            `query:"field2"`
	Field3 bool           `query:"field3"`
	Field4 map[string]any `query:"field4"`
	Field5 float64        `query:"field5"`
	Field6 float32        `query:"field6"`
	Field7 int32          `query:"field7"`
	Field8 int64          `query:"field8"`
}

func main() {
	//  http://localhost:3000/?field1=value1&field2=123&field3=true&field4={%22key1%22:%22value1%22,%22key2%22:%22value2%22}&field5=123.456&field6=123.456&field7=123&field8=123

	server := http.ServeMux{}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		dest := new(TestStruct)

		w.Header().Set("Content-Type", "application/json")

		if err := queryscan.Scan(query.Encode(), dest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid query"))
			return
		}

		if err := json.NewEncoder(w).Encode(dest); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		return
	})

	http.ListenAndServe(":3000", &server)
	// {
	// 	"Field1": "value1",
	// 	"Field2": 123,
	// 	"Field3": true,
	// 	"Field4": {
	// 	"key1": "value1",
	// 	"key2": "value2"
	// },
	// 	"Field5": 123.456,
	// 	"Field6": 123.456,
	// 	"Field7": 123,
	// 	"Field8": 123
	// }
}
```

## Testing

- Testing Via Local

  ```sh
   go test --race -v --failfast | make test
  ```

## Bechmark

- Testing Via Local

  ```sh
   go test --race -v --failfast -bench=. -benchmem | make btest
  ```

  ```sh
	goos: linux
	goarch: amd64
	pkg: restuwahyu13/queryscan
	cpu: AMD Ryzen 3 3200G with Radeon Vega Graphics
	BenchmarkQueryToStruct
	BenchmarkQueryToStruct-4                           35204             34304 ns/op           11471 B/op        112 allocs/op
	BenchmarkQueryToStructLargePayload
	BenchmarkQueryToStructLargePayload-4                1831            588038 ns/op           64871 B/op        610 allocs/op
	BenchmarkQueryToStructWithComplexJSON
	BenchmarkQueryToStructWithComplexJSON-4            35511             37386 ns/op           11955 B/op        109 allocs/op
	PASS
	ok      restuwahyu13/queryscan    4.380s
  ```

### Bugs

For information on bugs related to package libraries, please visit
[here](https://github.com/restuwahyu13/queryscan/issues)

### Contributing

Want to make **Go Query Scanner** more perfect ? Let's contribute and follow the
[contribution guide.](https://github.com/restuwahyu13/queryscan/blob/master/CONTRIBUTING.md)

### License

- [MIT License](https://github.com/restuwahyu13/queryscan/blob/master/LICENSE.md)

<p align="right" style="padding: 5px; border-radius: 100%; font-size: 2rem;">
  <b><a href="#queryscan">BACK TO TOP</a></b>
</p>