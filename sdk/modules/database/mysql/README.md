# database/mysql

## Enable

```
liquor app enable database/mysql
# or
go get github.com/go-liquor/liquor/sdk/modules/database/mysql
```

in `cmd/app/main.go` add module

```go
package main

import (
	"github.com/go-liquor/framework/internal/adapters/server/http"
	"github.com/go-liquor/framework/internal/app/services"
	"github.com/go-liquor/liquor/sdk/app"
    "github.com/go-liquor/liquor/sdk/modules/database/mysql" // add this
)

func main() {
	app.NewApp(
        mysql.DatabaseMysqlModule, // add this
		http.Server,
		app.RegisterService(
			services.NewInitialService,
		),
	)
}
```