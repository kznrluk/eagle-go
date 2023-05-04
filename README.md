# Eagle API in Go

Eagle-Go is a work-in-progress Go client for the [Eagle](https://en.eagle.cool/) application's API. With this library, you can manage items, folders, and libraries with Go.

Please note that this project is under development and may contain some bugs. We welcome any bug reports, feature requests, and pull requests to help us improve the package.

## Usage

To use Eagle-Go, first, import the package into your project:

```go
import "github.com/kznrluk/eagle-go"
```

Then, create a `Client` instance with the base URL of your Eagle application:

```go
cli := eagle.Client{BaseURL: "http://localhost:41595"} // default url
```

Now you can call various endpoints to interact with your Eagle library.

Here's a simple example of adding an item from a local file path:

```go
package main

import (
        "fmt"
        "github.com/kznrluk/eagle-go"
)

func main() {
        cli := eagle.Client{BaseURL: "http://localhost:41595"}

        resp, err := cli.AddItemFromPath(eagle.ItemAddFromPathRequest{
                Path:       "C:\\Users\\kznrl\\Desktop\\test.png",
                Name:       "test.img",
                Website:    "https://example.com",
                Annotation: "here is annotation",
                Tags:       []string{},
                FolderID:   "",
        })

        if err != nil {
                panic(err)
        }

        fmt.Printf("%+v\n", resp)
}
```

## Contributing

We're actively looking for contributors who can help improve and expand the package. If you find any issues or want to add a new feature, feel free to create an issue or submit a pull request. Your contributions are greatly appreciated!

## License

MIT