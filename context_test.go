package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
}

func TestContextGetValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

// func CreateCounter() chan int {
// 	destination := make(chan int)
// 	go func() {
// 		defer close(destination)
// 		counter := 1
// 		for {
// 			destination <- counter
// 			counter++
// 		}
// 	}()

// 	return destination
// }

// func TestContextWithCancel(t *testing.T) {
// 	fmt.Println("Total goroutines:", runtime.NumGoroutine())

// 	destination := CreateCounter()
// 	for n := range destination {
// 		fmt.Println("Counter", n)
// 		if n == 10 {
// 			break
// 		}
// 	}

// 	fmt.Println("Total goroutines:", runtime.NumGoroutine())

// }

// func CreateCounter(ctx context.Context, wg *sync.WaitGroup) chan int {
// 	destination := make(chan int)
// 	wg.Add(1)
// 	go func() {
// 		defer close(destination)
// 		defer wg.Done()
// 		counter := 1
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("DONE")
// 				return
// 			default:
// 				select {
// 				case <-ctx.Done():
// 					fmt.Println("DONE")
// 					return
// 				case destination <- counter:
// 					counter++
// 				}
// 			}
// 		}
// 	}()
// 	return destination
// }

// func TestContextWithCancel(t *testing.T) {
// 	fmt.Println("Total goroutines:", runtime.NumGoroutine())
// 	parent := context.Background()
// 	ctx, cancel := context.WithCancel(parent)
// 	group := &sync.WaitGroup{}

// 	destination := CreateCounter(ctx, group)
// 	fmt.Println("Total goroutines:", runtime.NumGoroutine())
// 	for n := range destination {
// 		fmt.Println("Counter", n)
// 		if n == 10 {
// 			break
// 		}
// 	}

// 	cancel()

// 	group.Wait()

// 	fmt.Println("Total goroutines:", runtime.NumGoroutine())
// }

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
