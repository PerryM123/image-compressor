## Notes
- [ ] `go install 〇〇` installs a binary to be used globally (?)
- [ ] `go get 〇〇` installs a dependency and adds it to `go.mod`
### Things to learn:

#### defer
runs code even if there is a sudden return so this is great for ensuring cleanup even if an exception or error occurs

#### mutex (mu.Lock() or mu.Unlock())
TODO

#### channels
TODO


#### const PascalCase vs const pascalCase
- PascalCase means, it is a constant variable that can be exported to another file
- pascalCase means it is a constant variable that cannot be exported to another file

#### Declaring a variable
```go
f := "apple" // shorthand for declares NEW variables
```

```go
var f string = "apple" // longhand for declares NEW variables
```