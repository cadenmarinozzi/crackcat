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
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
	"hash"
)

// I really want to add more hashes like WPA and bcrypt but but my priority should be to optimzie these first
func Hash(input string, algorithm string) string {
	var hashed hash.Hash;
             
	switch (algorithm) {
		case ("ripemd160"):
			hashed = ripemd160.New();

			break;

		case ("sha1"):
			hashed = sha1.New();

			break;

		case ("sha224"):
			hashed = sha256.New224();

			break;

		case ("sha3-224"):
			hashed = sha3.New224();

			break;

		case ("sha256"):
			hashed = sha256.New();

			break;

		case ("sha3-256"):
			hashed = sha3.New256();

			break;

		case ("keccak256"):
			hashed = sha3.NewLegacyKeccak256();
			
			break;

		case ("sha384"):
			hashed = sha512.New384();

			break;

		case ("sha3-384"):
			hashed = sha3.New384();

			break;

		case ("sha512"):
			hashed = sha512.New();

			break;

		case ("sha3-512"):
			hashed = sha3.New512();

			break;

		case ("keccak512"):
			hashed = sha3.NewLegacyKeccak512();
			
			break;

		case ("md5"):
			hashed = md5.New();

			break;

		case ("md4"):
			hashed = md4.New();

			break;
	}

	hashed.Write([]byte(input));

	return hex.EncodeToString(hashed.Sum(nil));
}