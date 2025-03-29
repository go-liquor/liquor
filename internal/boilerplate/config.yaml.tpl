app:
  name: {{.Name}}
  debug: true
log:
  level: "debug"
  format: "json"
rest:
  port: 8080
  cors:
    default: true
database:
  driver: sqlite
  sqlite:
    dns: "file::memory:?cache=shared"
  # origin:
  #   - "*"
  # methods:
  #   - "GET"
  #   - "POST"
  #   - "PUT"
  #   - "DELETE"
  #   - "OPTIONS"
  # headers:
  #   - "Origin"
  #   - "Authorization"
  #   - "Content-Type"
  # credentials: true