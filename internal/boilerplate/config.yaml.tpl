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
  # origins:
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
database:
  driver: sqlite
  sqlite:
    dns: "file::memory:?cache=shared"
