# Go + Docker = <3 Workshop

> Daugavpils, 2019/03/02

## How to run and build binaries

```bash
# Run using `go run`
cd cmd/chatserver && go run main.go
cd cmd/client && go run main.go

# Build binaries
cd cmd/chatserver && go build
cd cmd/client && go build
```


## How to build and run Docker image

```bash
docker build -t chatserver .
docker run --rm --publish 9999:9999 chatserver
```

## Worth to watch

- [**Go Proverbs** by Rob Pike](https://www.youtube.com/watch?v=PAAkCSZUG1c)
- [**7 common mistakes in Go and when to avoid them** by Steve Francia](https://www.youtube.com/watch?v=29LLRKIL_TI)
- [**Best Practices for Industrial Programming** by Peter Bourgon](https://www.youtube.com/watch?v=PTE4VJIdHPg)
- [**Advanced Testing with Go** by Mitchell Hashimoto](https://www.youtube.com/watch?v=8hQG7QlcLBk)
- [**Things in Go I Never Use** by Mat Ryer](https://www.youtube.com/watch?v=5DVV36uqQ4E)
- [**JustForFunc: Programming in Go** by Francesc Campoy](https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw)
- [**Go Lift** by John Cinnamond](https://www.youtube.com/watch?v=1B71SL6Y0kA)
- [**Twelve Go Best Practices** by Francesc Campoy](https://www.youtube.com/watch?v=8D3Vmm1BGoY)
