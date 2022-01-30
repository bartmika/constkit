# Constants Kit
[![GoDoc](https://godoc.org/github.com/gomarkdown/markdown?status.svg)](https://pkg.go.dev/github.com/bartmika/go-constkit)
[![Go Report Card](https://goreportcard.com/badge/github.com/bartmika/go-constkit)](https://goreportcard.com/report/github.com/bartmika/go-constkit)

A curated list of constants and some additional functions you can use in your Golang code.

## Installation

In your Golang project, please run:

```
go get github.com/bartmika/go-constkit
```

## Usage

```go
import (
    "fmt"

    "github.com/bartmika/go-constkit"
)

allergies := constkit.ListAllAlergies()
fmt.Println(allergies)
```

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [ISC License](LICENSE).

Resource used:

* [adamawolf/Apple_mobile_device_types.txt](https://gist.github.com/adamawolf/3048717) is used for the data source that powers this library.
* [List of Allergies](https://en.wikipedia.org/wiki/List_of_allergens) was used for listing common allergies.
* [List of Diseases](https://en.wikipedia.org/wiki/Lists_of_diseases) was used for listing common diseases.
