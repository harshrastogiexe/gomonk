# Go-Monk Validator

Go-Monk is schema validator

### Usage

```go
type UserInfo struct {
	Name string
	Age  int
}

rb := gomonk.NewRuleBook[UserInfo]()

rb.AddRuleFor("name", func(v *UserInfo, pe gomonk.PropertyError) {
    if v.Name == "" {
        pe["required"] = "name is required"
    }

    if len(v.Name) < 3 {
        pe["len"] = "length should be more than or equal to 3"
    }
})

v := rb.Build()

result := v.Validate(&UserInfo{
    Name: "Harsh Rastogi",
    Age : 10
})
```
