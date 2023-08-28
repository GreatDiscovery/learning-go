package proometheus

import (
	"bufio"
	"fmt"
	"github.com/beorn7/perks/quantile"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestP99(t *testing.T) {
	ch := make(chan float64)
	go sendFloats(ch)

	// Compute the 50th, 90th, and 99th percentile.
	q := quantile.NewTargeted(map[float64]float64{
		0.50:  0.005,
		0.90:  0.001,
		0.99:  0.0001,
		0.999: 0.00001,
	})
	for v := range ch {
		q.Insert(v)
	}

	fmt.Println("perc50:", q.Query(0.50))
	fmt.Println("perc90:", q.Query(0.90))
	fmt.Println("perc99:", q.Query(0.99))
	fmt.Println("perc999:", q.Query(0.999))
	fmt.Println("count:", q.Count())
}

func sendFloats(ch chan<- float64) {
	f, err := os.Open("exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		b := sc.Bytes()
		v, err := strconv.ParseFloat(string(b), 64)
		if err != nil {
			log.Fatal(err)
		}
		ch <- v
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}
	close(ch)
}
