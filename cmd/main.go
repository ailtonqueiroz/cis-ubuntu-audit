package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	
	"github.com/ailtonqueiroz/cis-ubuntu-audit/pkg/auditor"
)

type Benchmark struct {
	Categories []struct {
		Name     string
		Controls []auditor.Control
	}
}

func main() {
	// Carrega benchmark
	data, err := os.ReadFile("benchmarks/ubuntu-22.04.json")
	if err != nil {
		log.Fatal(err)
	}

	var benchmark Benchmark
	json.Unmarshal(data, &benchmark)

	// Executa verificações
	for _, category := range benchmark.Categories {
		fmt.Printf("\n=== %s ===\n", category.Name)
		for _, control := range category.Controls {
			compliant, _ := auditor.ScanControl(control)
			status := "❌"
			if compliant {
				status = "✅"
			}
			fmt.Printf("[%s] %s - %s\n", status, control.ID, control.Title)
		}
	}
}
