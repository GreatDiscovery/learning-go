package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	vaultclient "learning-go/test/framework/vault"
)

func main() {
	operation := flag.String("op", "get", "operation: put, get, or delete")
	secretPath := flag.String("path", "learning/demo", "KV v2 secret path")
	message := flag.String("message", "hello vault", "message stored by put")
	mount := flag.String("mount", envOrDefault("VAULT_MOUNT", "secret"), "KV v2 mount path")
	flag.Parse()

	client, err := vaultclient.New(vaultclient.Config{
		Address: os.Getenv("VAULT_ADDR"),
		Token:   os.Getenv("VAULT_TOKEN"),
		Mount:   *mount,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch *operation {
	case "put":
		err = client.Put(ctx, *secretPath, map[string]interface{}{
			"message":  *message,
			"saved_at": time.Now().UTC().Format(time.RFC3339),
		})
		if err == nil {
			fmt.Printf("secret %q written\n", *secretPath)
		}
	case "get":
		var data map[string]interface{}
		data, err = client.Get(ctx, *secretPath)
		if err == nil {
			fmt.Printf("secret %q: %#v\n", *secretPath, data)
		}
	case "delete":
		err = client.Delete(ctx, *secretPath)
		if err == nil {
			fmt.Printf("latest version of %q soft-deleted\n", *secretPath)
		}
	default:
		log.Fatalf("unsupported operation %q; use put, get, or delete", *operation)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func envOrDefault(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
