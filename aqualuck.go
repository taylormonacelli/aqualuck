package aqualuck

import (
	"fmt"
	"log/slog"
)

func Main() int {
	slog.Debug("aqualuck", "test", true)
	main()

	return 0
}

func main() {
	tea := &Tea{}
	coffee := &Coffee{}

	fmt.Println("Making tea:")
	tea.PrepareBeverage()

	fmt.Println("\nMaking coffee:")
	coffee.PrepareBeverage()
}
