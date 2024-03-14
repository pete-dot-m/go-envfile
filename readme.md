## envfile
A silly package to support reading in an .envfile

**Usage**
```
go get github.com/pete-dot-m/go-envfile
```
Import as usual, and then call LoadEnv to populate environment variables, preferably before they're needed...
```golang
import (
    "os"

    "github.com/pete-dot-m/go-envfile"
)
...
func main() {
    // loads environment values from either an .env file (default)
    // or from a user-defined file (and path)
    err := envfile.LoadEnv()
    err := envfile.LoadEnv(".myenv")
    err := envfile.LoadEnv("~/environment/apikeys")
    // err will be nil on success or will have a relatively helpful
    // message to help diagnose what you did wrong
    ...
    apiKey, err := os.Getenv("MY_API_KEY")
}
```

**Coming Soon**
- Support for loading a slice of files for allowing overrides etc.
- Other things
- Refactoring to be more idiomatic
