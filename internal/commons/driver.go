package commons

func SelectDatabaseDriver() string {
	return NewList("Select the module", []string{"mongodb", "mysql", "postgres", "sqlite"})
}
