package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sensor(name string, ch chan int, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			data := rand.Intn(100)
			formattedTime := t.Format("2006-01-02 15:04:05")
			fmt.Printf("\033[32m%s mengirim data: %d pada tanggal %s\033[0m\n", name, data, formattedTime)
			ch <- data
		case <-time.After(5 * time.Second):
			fmt.Printf("%s timeout!\n", name)
			return
		case <-done:
			fmt.Printf("%s berhenti atas permintaan\n", name)
			return

		}
	}
}

func main() {

	var wg sync.WaitGroup

	chSuhu := make(chan int, 1)
	chKelembaban := make(chan int, 1)
	chTekanan := make(chan int, 1)
	done := make(chan struct{})

	wg.Add(3)
	go sensor("Sensor Suhu", chSuhu, &wg, done)
	go sensor("Sensor Kelembaban", chKelembaban, &wg, done)
	go sensor("Sensor Tekanan", chTekanan, &wg, done)

	go func() {
		for data := range chSuhu {
			fmt.Printf("Menerima data dari Sensor Suhu: %d\n", data)
		}
	}()

	go func() {
		for data := range chKelembaban {
			fmt.Printf("Menerima data dari Sensor Kelembaban: %d\n", data)
		}
	}()

	go func() {
		for data := range chTekanan {
			fmt.Printf("Menerima data dari Sensor Tekanan: %d\n", data)
		}
	}()

	wg.Wait()

	close(chSuhu)
	close(chKelembaban)
	close(chTekanan)
	close(done)

	time.Sleep(1 * time.Second)

	fmt.Println("Semua sensor selesai.")
}
