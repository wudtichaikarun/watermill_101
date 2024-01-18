# Watermill 101

## install protobuf

```
$ brew install protobuf

// check version
$ protoc --version
```

## generate protobuf

```
$ protoc --go_out=. inputs/events.proto

```
