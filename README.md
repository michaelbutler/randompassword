# randompassword

Generate one or more random ASCII passwords securely in your terminal.

# Usage

```
NAME:
   randompassword - Generate secure random passwords

USAGE:
   randompassword [options]

OPTIONS:
   --length value  Length of the password(s) (default: 24)
   --count value   Number of passwords to generate (default: 10)
   --help, -h      show help
```

# Building

Requires Go 1.23+ to be installed.

```
go install
go build randompassword.go
```

Then run using `./randompassword`.