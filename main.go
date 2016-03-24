// Copyright 2016 Lennart Espe. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Gohash generates file hashes.

Usage:

	godoc [flag] file

The flags are:

		-a algorithm
			Available algorithms are crc32, crc64, adler32, fnv32, fnva32, fnv64, fnva64
		--version
			Shows the current verion of Gohash
		--help
			Displays the help page
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io/ioutil"
	"strings"
)

const versionInfo = "gohash 1.0"

// Returns the 32 bit hash of the file using the provided algorithm
func getHash32(filename string, hs hash.Hash32) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	hs.Write(bs)
	return hs.Sum32(), nil
}

// Returns the 64 bit hash of the file using the provided algorithm
func getHash64(filename string, hs hash.Hash64) (uint64, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	hs.Write(bs)
	return hs.Sum64(), nil
}

// Maps the name to an algorithm and hashes the file
func hashFile(filename string, algorithm string) (string, error) {
	if strings.HasSuffix(algorithm, "32") {
		var result uint32
		var err error
		switch algorithm {
		case "crc32":
			result, err = getHash32(filename, crc32.NewIEEE())
		case "adler32":
			result, err = getHash32(filename, adler32.New())
		case "fnv32":
			result, err = getHash32(filename, fnv.New32())
		case "fnva32":
			result, err = getHash32(filename, fnv.New32a())
		default:
			return "", errors.New("unknown algorithm")
		}

		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%x", result), nil
	} else if strings.HasSuffix(algorithm, "64") {
		var result uint64
		var err error
		switch algorithm {
		case "crc64":
			result, err = getHash64(filename, crc64.New(crc64.MakeTable(crc64.ISO)))
		case "fnv64":
			result, err = getHash64(filename, fnv.New64())
		case "fnva64":
			result, err = getHash64(filename, fnv.New64a())
		default:
			return "", errors.New("unknown algorithm")
		}
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%x", result), nil
	} else {
		return "", errors.New("unknown algorithm")
	}
}

func main() {
	version := flag.Bool("version", false, "display version information")
	help := flag.Bool("help", false, "show help information")
	algorithm := flag.String("a", "crc32", "set hash algorithm (`crc32`, crc64, adler32, fnv32, fnva32, fnv64, fnva64)")
	flag.Parse()

	// Show help information if requested
	if *help {
		fmt.Println("gohash [-a] [--help] [--version] [file]")
		flag.PrintDefaults()
		return
	}

	// Show version information if requested
	if *version {
		fmt.Println(versionInfo)
		return
	}

	filename := flag.Arg(0)

	result, err := hashFile(filename, *algorithm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(result)
}
