# Gohash
Gohash is a fast and easy tool to generate file hashes.

## Installation
```bash
$ export PATH=$PATH:$GOPATH/bin
$ go get github.com/mooxmirror/gohash
```

## Usage
```
gohash [flag] file
```

### Available flags
```
-a algorithm
    Available algorithms are crc32, crc64, adler32, fnv32, fnva32, fnv64, fnva64
--version
    Shows the current verion of Gohash
--help
    Displays the help page
```
