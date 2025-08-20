app:
  name: {{.name}}
  debug: true
log:
  level: "debug"
  format: "json"
grpc:
 port: 50051
rest:
  # disabled: false
  port: 8080
  cors:
    origins:
      - "*"
    methods:
      - "GET"
      - "POST"
      - "PUT"
      - "DELETE"
      - "OPTIONS"
      - "PATCH"
    headers:
      - "Origin"
      - "Authorization"
      - "Content-Type"
    credentials: true
database: {}
