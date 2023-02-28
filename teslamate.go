package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

func teslamateParseDockerCompose(dockerComposeFile string) (string, string, error) {
	var sqlURI, encryptionKey string
	var dockerCompose struct {
		Services struct {
			Teslamate struct {
				Environment []string
			}
			Database struct {
				Environment []string
				Ports       []string
			}
		}
	}

	data, err := os.ReadFile(dockerComposeFile)
	if err != nil {
		return "", "", err
	}
	err = yaml.Unmarshal(data, &dockerCompose)
	if err != nil {
		return "", "", err
	}
	for _, v := range dockerCompose.Services.Teslamate.Environment {
		p := strings.Split(v, "=")
		if len(p) != 2 {
			continue
		}
		if p[0] == "ENCRYPTION_KEY" {
			encryptionKey = p[1]
			break
		}
	}
	if encryptionKey == "" {
		return "", "", fmt.Errorf("%s: no ENCRYPTION_KEY found", dockerComposeFile)
	}
	// POSTGRES_USER=teslamate
	// POSTGRES_PASSWORD=faeTh9ei
	// POSTGRES_DB=teslamate
	// --
	// 127.0.0.1:5181:5432

	var pgUser, pgPass, pgDb string
	for _, v := range dockerCompose.Services.Database.Environment {
		p := strings.Split(v, "=")
		if len(p) != 2 {
			continue
		}
		switch p[0] {
		case "POSTGRES_USER":
			pgUser = p[1]
		case "POSTGRES_PASSWORD":
			pgPass = p[1]
		case "POSTGRES_DB":
			pgDb = p[1]
		}
	}
	var pgHost, pgPort string
	for _, v := range dockerCompose.Services.Database.Ports {
		p := strings.Split(v, ":")
		if len(p) == 2 && p[1] == "5432" {
			pgHost = "127.0.0.1"
			pgPort = p[0]
			break
		}
		if len(p) == 3 && p[2] == "5432" {
			pgHost = p[0]
			pgPort = p[1]
			break
		}
	}
	sqlURI = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pgUser, pgPass, pgHost, pgPort, pgDb)
	return sqlURI, encryptionKey, nil
}

func teslamateGetTokenFromSql(sqlURI string) (string, error) {
	db, err := sql.Open("postgres", sqlURI)
	if err != nil {
		return "", err
	}
	var token string
	err = db.QueryRow(`SELECT ENCODE(access,'hex') AS token FROM tokens ORDER BY updated_at DESC LIMIT 1`).
		Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func teslamateDecodeToken(hexKey string, encryptionKey string) (string, error) {
	const nonceSize = 12
	const debug = false
	key := sha256.Sum256([]byte(encryptionKey))

	crypted, err := hex.DecodeString(hexKey)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	if debug {
		fmt.Printf("Key: %d bytes: %x\n", len(key[:]), key[:])
		fmt.Printf("Type: %d\n", crypted[0])
		fmt.Printf("Length: %d\n", crypted[1])
		fmt.Printf("Key Tag: %q\n", string(crypted[2:2+crypted[1]]))
	}
	crypted = crypted[2+crypted[1]:]
	nonce := crypted[:nonceSize]
	if debug {
		fmt.Printf("nonce: %d bytes: %x\n", len(nonce), nonce)
	}
	crypted = crypted[nonceSize:]
	cipherTag := crypted[:16]
	if debug {
		fmt.Printf("cipherTag: %d bytes: %x\n", len(cipherTag), cipherTag)
	}
	crypted = crypted[16:]
	if debug {
		fmt.Printf("Rest: %d bytes\n", len(crypted))
	}
	crypted = append(crypted, cipherTag...)
	if debug {
		fmt.Printf("Total: %d bytes (%x)\n", len(crypted), crypted)
	}

	// aesgcm, err := cipher.NewGCMWithTagSize(block, 16)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plain, err := aesgcm.Open(nil, nonce, crypted, []byte("AES256GCM"))
	if err != nil {
		return "", err
	}

	return string(plain), nil
}
