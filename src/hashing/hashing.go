/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package hashing

import (
	"encoding/hex"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

func Hash(input string, algorithm string) string {
	var hashed hash.Hash;
             
	switch (algorithm) {
		case ("sha1"):
			hashed = sha1.New();

		case ("sha224"):
			hashed = sha256.New224();

		case ("sha256"):
			hashed = sha256.New();

		case ("sha384"):
			hashed = sha512.New384();

		case ("sha512"):
			hashed = sha512.New();

		case ("md5"):
			hashed = md5.New();
	}

	hashed.Write([]byte(input));

	return hex.EncodeToString(hashed.Sum(nil));
}