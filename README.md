# round-robin
generic round-robin is balancing algorithm written in golang

## Installation

```shell
go get github.com/pyropy/round-robin
```

## Example

```go
rr, _ := roundrobin.New(
    &url.URL{Host: "192.168.33.10"},
    &url.URL{Host: "192.168.33.11"},
    &url.URL{Host: "192.168.33.12"},
    &url.URL{Host: "192.168.33.13"},
)

rr.Next() // {Host: "192.168.33.10"}
rr.Next() // {Host: "192.168.33.11"}
rr.Next() // {Host: "192.168.33.12"}
rr.Next() // {Host: "192.168.33.13"}
rr.Next() // {Host: "192.168.33.10"}
rr.Next() // {Host: "192.168.33.11"}
```
## Author
[hlts2](https://github.com/hlts2)

## LICENSE
round-robin released under MIT license, refer [LICENSE](https://github.com/pyropy/round-robin/blob/master/LICENSE) file.
