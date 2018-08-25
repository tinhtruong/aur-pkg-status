package main

import (
	"context"
	"fmt"

	"bitbucket.org/tinhtruong/aur-ood/aur"
)

func main() {
	ctx := context.Background()
	update, err := aur.GetPackageUpdates(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("update: %+v", update)
}
