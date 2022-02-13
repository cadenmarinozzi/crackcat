/*
	author....: nekumelon
	License...: MIT (Check LICENSE)
*/

package hashing

func DetectAlgorithm(hash string) string {
	switch len(hash) {
		case (32):
			return "md5";

		case (40):
			return "sha1";

		case (56):
			return "sha224";

		case (64):
			return "sha256";

		case (96):
			return "sha384";

		case (128):
			return "sha512";
	}

	return "";
}