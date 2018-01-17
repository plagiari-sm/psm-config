###### psm-config

[psm-config](https://github.com/plagiari-sm/psm-config) is a centralized configuration package written in [Go](https://golang.org/) to store information for every service in [Plagiarism](https://github.com/plagiari-sm)'s Ecosystem.

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
# devevelopment.yaml
name: svc-listen
env: development
host: 0.0.0.0
port: 8000
mongo:
    port: 27017
    host: 0.0.0.0
    user: username
    pass: password
    path: collection
twitter:
    consumer-key: $TWITTER_CONSUMER_KEY
    consumer-secret: $TWITTER_CONSUMER_SECRET
    access-token: $TWITTER_ACCESS_TOKEN
    access-token-secret: $TWITTER_ACCESS_TOKEN_SECRET
```
