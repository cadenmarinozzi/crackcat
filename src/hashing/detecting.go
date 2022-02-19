/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package hashing

func DetectAlgorithm(hash string) string {
	switch (len(hash)) {
		case (32):
			return "md5"; // Or md4

		case (40):
			return "sha1";

		case (56):
			return "sha224"; // Or sha3-224

		case (64):
			return "sha256"; // Or sha3-256

		case (96):
			return "sha384"; // Or sha3-384

		case (128):
			return "sha512"; // Or sha3-512
	}

	return "";
}