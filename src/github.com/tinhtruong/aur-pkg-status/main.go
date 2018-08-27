package main

import (
	"context"
	"flag"
	"fmt"

	"bitbucket.org/tinhtruong/aur-pkg-status/aur"
)

var status string

func init() {
	flag.StringVar(&status, "status", "updated", "What kind of status to show. The value can be 'updated', 'removed', 'all'")
}

func main() {
	flag.Parse()

	ctx := context.Background()
	updates, err := aur.GetPackageStatus(ctx)
	if err != nil {
		panic(err)
	}
	updates = filterByStatus(updates, status)
	fmt.Printf("┌───────────────────────────────┬─────────────────────┬──────────────────┐\n")
	fmt.Printf("│  Package Name                 │  Installed Version  │  Latest Version  │\n")
	fmt.Printf("├───────────────────────────────┼─────────────────────┼──────────────────┤\n")
	for _, update := range updates {
		fmt.Printf("│  %-29s│  %-19s│  %-16s│\n", update.PkgName, update.InstalledVersion, update.LatestVersion)
	}
	fmt.Printf("└───────────────────────────────┴─────────────────────┴──────────────────┘\n")
}

func filterByStatus(updates []aur.PackageStatus, opt string) []aur.PackageStatus {
	switch opt {
	case "updated":
		return filter(updates, func(update aur.PackageStatus) bool {
			return update.InstalledVersion != update.LatestVersion && update.LatestVersion != ""
		})
	case "removed":
		return filter(updates, func(update aur.PackageStatus) bool {
			return update.LatestVersion == ""
		})
	case "all":
		return updates
	default:
		return nil
	}
}

func filter(updates []aur.PackageStatus, f func(aur.PackageStatus) bool) []aur.PackageStatus {
	result := make([]aur.PackageStatus, 0)
	for _, update := range updates {
		if f(update) {
			result = append(result, update)
		}
	}
	return result
}
