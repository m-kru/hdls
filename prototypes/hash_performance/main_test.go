package main_test

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

var file_names []string
var N int

func init() {
	files_info, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("can't read working directory: %v", err)
	}

	for _, fi := range files_info {
		name := fi.Name()
		if strings.HasSuffix(name, ".vhd") {
			file_names = append(file_names, name)
		}
	}

	N := len(file_names)
	fmt.Println("Number of files: ", N)
}

func BenchmarkCRC32IEEEPoly(b *testing.B){
	for i := 0; i < b.N; i++ {
		for _, name := range file_names {
			content, err := ioutil.ReadFile(name)
			if err != nil {
				log.Fatalf("can't read file %s: %v", name, err)
			}

			crc32.ChecksumIEEE(content)
		}
	}
}

func BenchmarkCRC64ISOPoly(b *testing.B){
	for i := 0; i < b.N; i++ {
		for _, name := range file_names {
			content, err := ioutil.ReadFile(name)
			if err != nil {
				log.Fatalf("can't read file %s: %v", name, err)
			}

			tab := crc64.MakeTable(crc64.ISO)
			crc64.Checksum(content, tab)
		}
	}
}

func BenchmarkCRC64ECMAPoly(b *testing.B){
	for i := 0; i < b.N; i++ {
		for _, name := range file_names {
			content, err := ioutil.ReadFile(name)
			if err != nil {
				log.Fatalf("can't read file %s: %v", name, err)
			}

			tab := crc64.MakeTable(crc64.ECMA)
			crc64.Checksum(content, tab)
		}
	}
}

func BenchmarkSHA1(b *testing.B){
	for i := 0; i < b.N; i++ {
		for _, name := range file_names {
			content, err := ioutil.ReadFile(name)
			if err != nil {
				log.Fatalf("can't read file %s: %v", name, err)
			}

			sha1.Sum(content)
		}
	}
}

func BenchmarkSHA256(b *testing.B){
	for i := 0; i < b.N; i++ {
		for _, name := range file_names {
			content, err := ioutil.ReadFile(name)
			if err != nil {
				log.Fatalf("can't read file %s: %v", name, err)
			}

			sha256.Sum256(content)
		}
	}
}

func BenchmarkMD5(b *testing.B){
	for i := 0; i < b.N; i++ {
		for _, name := range file_names {
			content, err := ioutil.ReadFile(name)
			if err != nil {
				log.Fatalf("can't read file %s: %v", name, err)
			}

			md5.Sum(content)
		}
	}
}
