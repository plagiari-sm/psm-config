###### psm-config

[psm-config](https://github.com/plagiari-sm/psm-config) is a centralized configuration package written in [Go](https://golang.org/) to store information for every service in the [Plagiarism Ecosystem](https://github.com/plagiari-sm).

#### How to use
```go
package main

import (
	"fmt"

	// import psm-config package
	psmconfig "github.com/plagiari-sm/psm-config"
)

func init() {
	// during init or in main read `--config ./conf/default.yaml` flag
	// and parse configuration file
	psmconfig.NewConfig()
}

func main() {
	// access global configuration variable
	config := psmconfig.Config
	fmt.Println(config)
}
```
#### Configuration file example
```yaml
# dev.yaml
name: svc-listen
env: development
host: 0.0.0.0
port: 8000
db:
    mongo:
        host: 0.0.0.0
        port: 27017
        path: streams
    redis:
        host: 0.0.0.0
        port: 6379
        path: streams
log:
    level: DEBUG
    output: stdout
api:
    api-feeds: https://0.0.0.0:8080/
```
