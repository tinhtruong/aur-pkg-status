package main

import (
	"context"
	"fmt"

	"bitbucket.org/tinhtruong/aur-ood/aur"
)

func main() {
	ctx := context.Background()
	updates, err := aur.GetPackageUpdates(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("┌───────────────────────────────┬─────────────────────┬──────────────────┐\n")
	fmt.Printf("│  Package Name                 │  Installed Version  │  Latest Version  │\n")
	fmt.Printf("├───────────────────────────────┼─────────────────────┼──────────────────┤\n")
	for _, update := range updates {
		fmt.Printf("│  %-29s│  %-19s│  %-16s│\n", update.Name, update.InstalledVersion, update.LatestVersion)
	}
	fmt.Printf("└───────────────────────────────┴─────────────────────┴──────────────────┘\n")
}
