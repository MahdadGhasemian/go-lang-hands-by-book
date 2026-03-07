# Go Lang Hands by Book

## Introduction
Journey to mastering Go, through focused study and hands-on coding exercises from Books.

## Learning Plan
- **Book 1**: [Learning Go - An Idiomatic Approach to Real-World Go -- Jon Bodner --](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/)
- **Book 2**: [gRPC Microservices in Go -- Huseyin Babal --](https://www.oreilly.com/library/view/grpc-microservices-in/9781633439207/)
- **Book 3**: [Domain-Driven Design with Golang -- Matthew Boyle --](https://www.oreilly.com/library/view/domain-driven-design-with/9781804613450/)


## First Go Program

```bash
mkdir ch1
cd ch1
go mod init hello_world
go build
go build -o hello
go fmt ./...
go vet ./...
make
```

## Types

```bash
mkdir -p ch2/complex
cd ch2/complex
go mod init complex_numbers
touch complex.go
# write codes
go run complex.go
```

```bash
mkdir -p ch2/type_conversions
cd ch2/type_conversions
go mod init type_conversions
touch conversions.go
# write codes
go run conversions.go
```

## Len Cap

```bash
mkdir -p ch3/len_cap
cd ch3/len_cap
go mod init len_cap
touch len_cap
# write codes
go run len_cap
```