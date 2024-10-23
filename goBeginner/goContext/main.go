package main

import (
	"context"
	"fmt"
	"time"
)

// buatkan 3 func yang funcnya func1 cetak sebuah text setiap 2 detik, func2 mencetak text setiap 1 detik, func3 mencetak teks setiap 3 detik, buatkan context untuk membatalkan semua func yg berjalan didetik ke 5

func func1(ctx context.Context, name string) {
	for {

		select {
		case <-ctx.Done():
			fmt.Printf("%s dibatalkan\n", name)
		default:
			time.Sleep(2 * time.Second)
			fmt.Printf("%s working\n", name)
		}
	}
}

func func2(ctx context.Context, name string) {
	for {

		select {
		case <-ctx.Done():
			fmt.Printf("%s dibatalkan\n", name)
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("%s working\n", name)
		}
	}
}

func func3(ctx context.Context, name string) {
	for {

		select {
		case <-ctx.Done():
			fmt.Printf("%s dibatalkan\n", name)
		default:
			time.Sleep(3 * time.Second)
			fmt.Printf("%s working\n", name)
		}
	}
}

func main() {
	parentCtx := context.Background()

	ctx, cancel := context.WithCancel(parentCtx)

	defer cancel()

	go func1(ctx, "working1...")
	go func2(ctx, "working2...")
	go func3(ctx, "working3...")

	time.Sleep(4 * time.Second)
	fmt.Println("cancel context..")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("done")
}
