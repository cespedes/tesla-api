package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/cespedes/tesla-api/api"
)

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func run(args []string) error {
	flag := flag.NewFlagSet("tesla-api", flag.ExitOnError)
	var flags struct {
		// General flags
		Debug bool
		NoOp  bool

		// How to get token:
		EncryptionKey          string // Used in TeslamateSqlURI and TeslamateToken
		TeslamateDockerCompose string
		TeslamateSqlURI        string
		TeslamateToken         string
		Token                  string
	}
	var err error

	flag.BoolVar(&flags.Debug, "debug", false, "debugging info")
	flag.BoolVar(&flags.NoOp, "n", false, "no operation")
	flag.StringVar(&flags.EncryptionKey, "encryption-key", "", "key to decrypt access token")
	flag.StringVar(&flags.TeslamateDockerCompose, "teslamate-docker-compose", "", "path to Teslamate's \"docker-compose.yml\"")
	flag.StringVar(&flags.TeslamateSqlURI, "sql-uri", "", "connection string to Teslamate's PostgreSQL")
	flag.StringVar(&flags.TeslamateToken, "teslamate-token", "", "path to Teslamate's crypted access token")
	flag.StringVar(&flags.Token, "token", "", "access token")
	flag.Parse(args[1:])

	if flags.TeslamateDockerCompose == "" && flags.TeslamateSqlURI == "" &&
		flags.TeslamateToken == "" && flags.Token == "" {
		flags.TeslamateDockerCompose = "docker-compose.yml"
	}

	if flags.TeslamateDockerCompose != "" {
		if flags.TeslamateSqlURI != "" || flags.TeslamateToken != "" || flags.Token != "" {
			return fmt.Errorf("incompatible flags")
		}
		if flags.EncryptionKey != "" {
			return fmt.Errorf("-encryption-key must not be used with -teslamate-docker-compose")
		}

		flags.TeslamateSqlURI, flags.EncryptionKey, err = teslamateParseDockerCompose(flags.TeslamateDockerCompose)
		if err != nil {
			return err
		}
		if flags.Debug {
			fmt.Printf("Encryption key = %q\n", flags.EncryptionKey)
			fmt.Printf("SQL URI = %q\n", flags.TeslamateSqlURI)
		}
	}

	if flags.TeslamateSqlURI != "" {
		if flags.TeslamateToken != "" || flags.Token != "" {
			return fmt.Errorf("incompatible flags")
		}
		if flags.EncryptionKey == "" {
			return fmt.Errorf("-encryption-key must be used with -teslamate-sql-uri")
		}

		flags.TeslamateToken, err = teslamateGetTokenFromSql(flags.TeslamateSqlURI)
		if err != nil {
			return err
		}
		if flags.Debug {
			fmt.Printf("TeslamateToken = %q\n", flags.TeslamateToken)
		}
	}

	if flags.TeslamateToken != "" {
		if flags.Token != "" {
			return fmt.Errorf("incompatible flags")
		}
		if flags.EncryptionKey == "" {
			return fmt.Errorf("-encryption-key must be used with -teslamate-token")
		}

		flags.Token, err = teslamateDecodeToken(flags.TeslamateToken, flags.EncryptionKey)
		if err != nil {
			return err
		}
		if flags.Debug {
			fmt.Printf("Token = %q\n", flags.Token)
		}
	}

	if flags.Token == "" {
		return fmt.Errorf("no token to use")
	}
	if flags.NoOp {
		fmt.Println("NoOp: will *not* connect to Tesla.")
		return nil
	}
	args = flag.Args()
	if len(args) < 2 || len(args) > 3 {
		fmt.Fprintln(os.Stderr, "Usage: tesla-api [options] <method> <url> [<data>]")
		os.Exit(1)
	}
	if flags.Debug {
		fmt.Println("Will connect to Tesla...")
	}

	api, err := api.New("https://owner-api.teslamotors.com", flags.Token)
	if err != nil {
		return fmt.Errorf("tesla: %w", err)
	}
	var dest any
	var data []byte
	if len(args) == 3 {
		data = []byte(args[2])
	}
	err = api.Request(args[0], args[1], data, &dest)
	if err != nil {
		return fmt.Errorf("api.Request: %w", err)
	}
	out, err := json.Marshal(dest)
	if err != nil {
		return fmt.Errorf("json response: %w", err)
	}
	fmt.Println(string(out))
	return nil
}
