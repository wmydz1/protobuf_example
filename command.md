the compiler will read the files src/foo.proto and src/bar/baz.proto. It produces two output files: build/gen/foo.pb.go and build/gen/bar/baz.pb.go.

The compiler automatically creates the directory build/gen/bar if necessary, but it will not create build or build/gen; they must already exist.

```
protoc --proto_path=src --go_out=build/gen src/foo.proto src/bar/baz.proto

```

```
protoc --java_out=. -I .  example.proto
```

```
protoc --go_out=plugins=grpc:. *.proto

```

```
protoc --go_out=. *.proto

```
