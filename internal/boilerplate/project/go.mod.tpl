module {{ .module }}

go {{ .goversion }}

require "github.com/go-liquor/liquor/v3" v3.0.0

replace (
    github.com/go-liquor/liquor/v3 => ../ 
)