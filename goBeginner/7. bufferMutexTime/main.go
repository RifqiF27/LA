package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// buat kan func count down, di jalankan di goroutine, ditampilkan di main

func countDown(num int, ch chan bool) {
	for i := num; i > 0; i-- {
		fmt.Println("count :", i)
		time.Sleep(1 * time.Second)
	}
	ch <- true
	close(ch)
}

// buatkan 3 func, dimana func ini ada param dan message, 3 func ini dijalankan di goroutine berbeda,
// buatkan juga masing masing func punya channel sendiri,
// kemudian tampilkan pesan dari 3 func ini yang paling cepat memberikan pesannya

func func1(message string, ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- message
}

func func2(message string, ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- message
}

func func3(message string, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- message
}

// buatkan 1 func(sync waitGroup, num) dimana func itu menjalankan filter nilai yang ganjil untuk ditampilkan, kemudian jalankan func ganjil sebanyak 10 kali dimasing masing goroutine
// jalankan set 4 core

func ganjil(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= num; i++ {
		if i%2 != 0 {
			fmt.Printf("Ganjil : %d\n", i)
		}
	}
}

// Membuat struct dengan properti kuantiti dan nama produk.
// Membuat fungsi untuk mengurangi kuantiti dari struct tersebut.
// Menjalankan fungsi dalam goroutine sebanyak 10 kali.
// Mencetak kuantiti terakhir setelah semua goroutine selesai.

type Produk struct {
	nama     string
	kuantiti int
	mutex    sync.Mutex
}

func (p *Produk) kurangiKuantiti(jumlah int, wg *sync.WaitGroup) {
	defer wg.Done()
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.kuantiti >= jumlah {
		p.kuantiti -= jumlah
		fmt.Printf("Kuantiti %s dikurangi sebesar %d, sisa: %d\n", p.nama, jumlah, p.kuantiti)
	} else {
		fmt.Printf("Kuantiti %s tidak cukup untuk dikurangi sebesar %d, sisa: %d\n", p.nama, jumlah, p.kuantiti)
	}
}
func main() {
	ch := make(chan bool, 2)
	go countDown(3, ch)

	<-ch
	fmt.Println("Countdown selesai!")

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func1("Pesan dari func1", ch1)
	go func2("Pesan dari func2", ch2)
	go func3("Pesan dari func3", ch3)

	select {
	case msg1 := <-ch1:
		fmt.Println("Pesan tercepat dari func1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Pesan tercepat dari func2:", msg2)
	case msg3 := <-ch3:
		fmt.Println("Pesan tercepat dari func3:", msg3)
	}

	fmt.Println("Selesai menerima pesan tercepat.")

	numCPU := runtime.NumCPU()
	fmt.Printf("jumlah CPU: %d\n", numCPU)

	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go ganjil(10, &wg)
	}

	wg.Wait()
	fmt.Println("done")

	produk := &Produk{
		nama:     "ProdukA",
		kuantiti: 10,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go produk.kurangiKuantiti(1, &wg)
	}

	wg.Wait()

	fmt.Printf("Kuantiti akhir %s: %d\n", produk.nama, produk.kuantiti)
}
