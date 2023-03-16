package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Controleer of er een argument is meegegeven
	if len(os.Args) == 1 {
		fmt.Println("Geef mij nummer a sabbi.")
		return
	}

	// Zet het argument om naar een integer
	n, err := strconv.Atoi(os.Args[1])

	// Controleer of er een error is opgetreden tijdens de conversie of als n <= 0 is
	if err != nil || n <= 0 {
		fmt.Println("Wallah jij probeert fraude te plegen.")
	}

	// Herhaal de tekst "Alarm x !" n keer, waarbij x ophoogt bij elke iteratie
	for i := 1; i <= n; i++ {
		fmt.Printf("Politie %d !\n", i)
		if i >= 10 {
			fmt.Println("Wallah jij probeert fraude te plegen.")
			break
		}
	}
}
