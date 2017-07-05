package sort2mbmem

import (
	"bufio"
	"math/rand"
	"os"

	"github.com/riccardomc/golab/skiena/sortingandsearching/priorityqueue"
)

func readBytes(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	r := bufio.NewReader(f)
	for {
		data := make([]byte, 4)
		n, err := r.Read(data)
		if err != nil || n <= 0 {
			break
		}
	}
}

//
func Sort(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		return
	}

	q := priorityqueue.NewPriorityQueue(500000)
	r := bufio.NewReader(f)
	for {
		data := make([]byte, 4)
		n, err := r.Read(data)
		if err != nil || n <= 0 {
			break
		}
		q.Insert(int(bytes2int(data)))
	}
}

func generateInput(filename string, megabytes int) {
	integers := megabytes * 1000000
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()
	for integers > 0 {
		integer := rand.Int31()
		written, err := w.Write(int2bytes(integer))
		if err != nil {
			return
		}
		integers -= written
	}

}

func int2bytes(integer int32) []byte {
	return []byte{
		byte(integer >> 24),
		byte(integer >> 16),
		byte(integer >> 8),
		byte(integer),
	}
}

func bytes2int(bytes []byte) int32 {
	return int32(uint32(bytes[3]) | uint32(bytes[2])<<8 | uint32(bytes[1])<<16 | uint32(bytes[0])<<24)
}
