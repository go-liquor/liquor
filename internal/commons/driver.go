package commons

// SelectDatabaseDriver presents an interactive list of supported database drivers
// and returns the user's selection. The available options include MongoDB, MySQL,
// PostgreSQL, and SQLite.
//
// Returns:
//   - string: The selected database driver name
func SelectDatabaseDriver() string {
	return NewList("Select the module", []string{"mongodb", "mysql", "postgres", "sqlite"})
}
