package app

import (
	"context"
	"fmt"
	"io"
	"log"
)

const nameTracer = "fib"

type App struct {
	r io.Reader
	l *log.Logger
}

func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r: r, l: l}
}

func (a *App) Run(ctx context.Context) error {
	for {
		n, err := a.Poll(ctx)
		if err != nil {
			return err
		}

		a.Write(ctx, n)
	}
}

func (a *App) Poll(ctx context.Context) (uint, error) {
	a.l.Println("What Fibonacci number would you like to know: ")
	var n uint

	_, err := fmt.Fscanf(a.r, "%d\n", &n)
	return n, err
}

func (a *App) Write(ctx context.Context, n uint) {
	f, err := Fibonacci(n)
	if err != nil {
		a.l.Printf("Fibonacci(%d): %v\n", n, err)
	} else {
		a.l.Printf("Fibonacci(%d): %v\n", n, f)
	}
}
