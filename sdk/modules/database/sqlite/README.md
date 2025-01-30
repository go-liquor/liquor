# database/sqlite

## Enable

```
liquor app enable database/sqlite
# or
go get github.com/go-liquor/liquor/sdk/modules/database/sqlite
```

in `cmd/app/main.go` add module

```go
package main

import (
	"github.com/go-liquor/framework/internal/adapters/server/http"
	"github.com/go-liquor/framework/internal/app/services"
	"github.com/go-liquor/liquor/sdk/app"
    "github.com/go-liquor/liquor/sdk/modules/database/sqlite" // add this
)

func main() {
	app.NewApp(
        sqlite.DatabaseSqliteModule, // add this
		http.Server,
		app.RegisterService(
			services.NewInitialService,
		),
	)
}
```