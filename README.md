# go-ingresso

A Go wrapper for the API of [Ingresso.com](https://www.ingresso.com/).

An **partnership code** is needed to use the API. Check if you can get one at [ingresso.com](https://suporte.ingresso.com/hc/pt-br/articles/115001166112-Integra%C3%A7%C3%A3o-com-a-API-de-Conte%C3%BAdo-).

Note: This product uses the Ingresso.com API but is not endorsed or certified by Ingresso.com.

## How to install

```shell
go get github.com/dsbezerra/go-ingresso
```

## How to use

Import the library

```go
import "github.com/dsbezerra/go-ingresso"
```

Initialize the library using your partnership code
```go
ing := ingresso.New(
  ingresso.Partnership("PARTNERSHIP_CODE"),
)
```

Use the API methods as you want, for example:

```go
theater, err := ing.GetTheater("372")
```

All functions return Go structs. You can get JSON using the ToJSON function such as:

```go
json, err := ing.ToJSON(theater)
```

## License

The MIT License (MIT)

Copyright (c) 2018 Diego Bezerra

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
