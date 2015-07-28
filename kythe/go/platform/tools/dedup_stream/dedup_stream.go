/*
 * Copyright 2015 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Binary dedup_stream reads a delimited stream from stdin and writes a delimited stream to stdout.
// Each record in the stream will be hashed, and if that hash value has already been seen, the
// record will not be emitted.
package main

import (
	"flag"
	"log"
	"os"

	"kythe.io/kythe/go/platform/delimited"
	"kythe.io/kythe/go/util/flagutil"
)

func init() {
	flag.Usage = flagutil.SimpleUsage("Remove duplicate records from a delimited stream")
}

var cacheSize = flag.Int("cache_size", 3*1024*1024*1024 /* 3 GB */, "Maximum size (in bytes) of the cache of known record hashes")

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		flagutil.UsageErrorf("unknown arguments: %v", flag.Args())
	}

	rd, err := delimited.NewUniqReader(delimited.NewReader(os.Stdin), *cacheSize)
	if err != nil {
		log.Fatalf("Error creating UniqReader: %v", err)
	}
	wr := delimited.NewWriter(os.Stdout)
	if err := delimited.Copy(wr, rd); err != nil {
		log.Fatal(err)
	}
	log.Printf("dedup_stream: skipped %d records", rd.Skipped())
}
