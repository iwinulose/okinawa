# okinawa
A simple Golang pretty printer.

Usage:
```go

type MyStruct struct {
  Foo string
}

func (m MyStruct) String() string {
  r := okinawa.Ramen{}
  r.Name = "MyStruct"
  r.Append("Foo", m.Foo)
  return r.String()
}

```
