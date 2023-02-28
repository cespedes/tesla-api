package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func run(args []string) error {
	flag := flag.NewFlagSet("tesla-cli", flag.ExitOnError)
	var flags struct {
		// General flags
		Debug bool

		// How to get token:
		EncryptionKey          string // Used in TeslamateSqlURI and TeslamateToken
		TeslamateDockerCompose string
		TeslamateSqlURI        string
		TeslamateToken         string
		Token                  string
	}
	var err error

	flag.BoolVar(&flags.Debug, "debug", false, "debugging info")
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
	fmt.Println("Will connect to Tesla...")
	err = teslaCli(flags.Token, flag.Args())
	if err != nil {
		return fmt.Errorf("tesla: %w", err)
	}
	return nil
}
