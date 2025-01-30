# database/mongodb

## Enable
```
liquor app enable database/mongodb
# or
go get github.com/go-liquor/liquor/sdk/modules/database/mongodb
```

in `cmd/app/main.go` add module

```go
package main

import (
	"github.com/go-liquor/framework/internal/adapters/server/http"
	"github.com/go-liquor/framework/internal/app/services"
	"github.com/go-liquor/liquor/sdk/app"
    "github.com/go-liquor/liquor/sdk/modules/database/mongodb" // add this
)

func main() {
	app.NewApp(
        mongodb.DatabaseMongoDBModule, // add this
		http.Server,
		app.RegisterService(
			services.NewInitialService,
		),
	)
}
```