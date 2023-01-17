#include <openssl/engine.h>
#include <openssl/evp.h>
#include <openssl/rsa.h>

int main() {
  EVP_PKEY_CTX *ctx;
  ENGINE *eng;
  unsigned char *out, *in;
  size_t outlen, inlen;
  EVP_PKEY *key;

  /*
   * NB: assumes eng, key, in, inlen are already set up,
   * and that key is an RSA public key
   */
  ctx = EVP_PKEY_CTX_new(key, eng);
  if (!ctx) return 1;
  /* Error occurred */
  if (EVP_PKEY_encrypt_init(ctx) <= 0) return 1;
  /* Error */
  if (EVP_PKEY_CTX_set_rsa_padding(ctx, RSA_PKCS1_OAEP_PADDING) <= 0) return 1;
  /* Error */

  /* Determine buffer length */
  if (EVP_PKEY_encrypt(ctx, NULL, &outlen, in, inlen) <= 0) return 1;
  /* Error */

  out = OPENSSL_malloc(outlen);

  if (!out) return 1;
  /* malloc failure */

  if (EVP_PKEY_encrypt(ctx, out, &outlen, in, inlen) <= 0) return 1;
  /* Error */

  /* Encrypted data is outlen bytes written to buffer out */
}
