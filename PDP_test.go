package PDP


import (
	"fmt"
	"testing"
)

func TestParseInStock(t *testing.T) {
	sizes, err := Parse("footlocker.com", "W2288111", "inStock")
	if err != nil {
		t.Fatal(err)
	}
	// to grab a specific size we would do
	// fmt.Println(sizes["08.0"] etc
	fmt.Println(sizes)
}

func TestParseAll(t *testing.T) {
	sizes, err := Parse("footlocker.com", "W2288111", "all")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sizes)
}

func TestParseRandomSize(t *testing.T) {
	sizes, err := Parse("footlocker.com", "W2288111", "inStock")
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range sizes {
		fmt.Printf("Size: %s, PID: %s", k, v)
		break
	}
}

func BenchMarkParse(b *testing.B) {
	// don't time while creating client.
	b.StopTimer()
	// start benchmark.
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Parse("footlocker.com", "W2288111", "inStock"); err != nil {
			panic(err)
		}
	}
}