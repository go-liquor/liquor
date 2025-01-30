# database/postgres

## Enable

```
liquor app enable database/postgres
# or
go get github.com/go-liquor/liquor/sdk/modules/database/postgres
```

in `cmd/app/main.go` add module

```go
package main

import (
	"github.com/go-liquor/framework/internal/adapters/server/http"
	"github.com/go-liquor/framework/internal/app/services"
	"github.com/go-liquor/liquor/sdk/app"
    "github.com/go-liquor/liquor/sdk/modules/database/postgres" // add this
)

func main() {
	app.NewApp(
        postgres.DatabasePostgresModule, // add this
		http.Server,
		app.RegisterService(
			services.NewInitialService,
		),
	)
}
```