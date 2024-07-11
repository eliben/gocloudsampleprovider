// This sample shows how to use the custom out-of-tree blob provider.
package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/eliben/gocloudsampleprovider"
	"gocloud.dev/blob"
)

func main() {
	ctx := context.Background()

	// gocloudsampleprovider registers the gcspfile:// route for blob handling,
	// so we can pass that to blob.OpenBucket directly.
	b, err := blob.OpenBucket(ctx, "gcspfile://tempfile1")
	if err != nil {
		log.Fatal(err)
	}
	err = b.WriteAll(ctx, "my-key", []byte("hello world"), nil)
	if err != nil {
		log.Fatal(err)
	}
	data, err := b.ReadAll(ctx, "my-key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
