// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuzz

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func BenchmarkMutatorBytes(b *testing.B) {
	origEnv := os.Getenv("GODEBUG")
	defer func() { os.Setenv("GODEBUG", origEnv) }()
	os.Setenv("GODEBUG", fmt.Sprintf("%s,fuzzseed=123", origEnv))
	m := newMutator()

	for _, size := range []int{
		1,
		10,
		100,
		1000,
		10000,
		100000,
	} {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			buf := make([]byte, size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				// resize buffer to the correct shape and reset the PCG
				buf = buf[0:size]
				m.r = newPcgRand()
				m.mutate([]interface{}{buf}, workerSharedMemSize)
			}
		})
	}
}

func BenchmarkMutatorString(b *testing.B) {
	origEnv := os.Getenv("GODEBUG")
	defer func() { os.Setenv("GODEBUG", origEnv) }()
	os.Setenv("GODEBUG", fmt.Sprintf("%s,fuzzseed=123", origEnv))
	m := newMutator()

	for _, size := range []int{
		1,
		10,
		100,
		1000,
		10000,
		100000,
	} {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			buf := make([]byte, size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				// resize buffer to the correct shape and reset the PCG
				buf = buf[0:size]
				m.r = newPcgRand()
				m.mutate([]interface{}{string(buf)}, workerSharedMemSize)
			}
		})
	}
}

func BenchmarkMutatorAllBasicTypes(b *testing.B) {
	origEnv := os.Getenv("GODEBUG")
	defer func() { os.Setenv("GODEBUG", origEnv) }()
	os.Setenv("GODEBUG", fmt.Sprintf("%s,fuzzseed=123", origEnv))
	m := newMutator()

	types := []interface{}{
		[]byte(""),
		string(""),
		false,
		float32(0),
		float64(0),
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
	}

	for _, t := range types {
		b.Run(fmt.Sprintf("%T", t), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.r = newPcgRand()
				m.mutate([]interface{}{t}, workerSharedMemSize)
			}
		})
	}
}
