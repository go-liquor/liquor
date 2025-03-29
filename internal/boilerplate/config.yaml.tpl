app:
  name: {{.Name}}
  debug: true
log:
  level: "debug"
  format: "json"
grpc:
 port: 50051
rest:
  # disabled: false # if you wants disable rest, you can put here as true
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
