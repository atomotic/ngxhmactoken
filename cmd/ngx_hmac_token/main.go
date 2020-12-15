package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/atomotic/ngxhmactoken"
)

func main() {

	base := flag.String("base", "", "base url")
	secret := flag.String("secret", "", "secret")
	expires := flag.Int("expires", 3600, "expires")
	path := flag.String("path", "", "path")

	flag.Parse()

	if *base == "" || *secret == "" || *path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	now := time.Now().Unix()
	token := ngxhmactoken.GenerateToken(*path, *secret, *expires)
	fmt.Printf("%s%s?st=%s&ts=%d&e=%d\n", *base, *path, token, now, *expires)
}
