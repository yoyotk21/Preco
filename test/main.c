#include <stdio.h> 
#include <openssl/evp.h>
#include <string.h>

int main()
{
    // creating new evp context
    EVP_CIPHER_CTX *ctx = EVP_CIPHER_CTX_new();

    if (1 != EVP_EncryptInit_ex(ctx, EVP_aes_256_cbc(), NULL, (unsigned char *)"Key", (unsigned char *)"IV"))
    {
        printf("Errors initializing evp context\n");
    }
 
    unsigned char data[50];
    int len;

    if (1 != EVP_EncryptUpdate(ctx, data, &len, (unsigned char *)"Hello World", 12))
    {
        printf("Errors encrypting!\n");
    }

    if (1 != EVP_EncryptFinal_ex(ctx, data, &len))
    {
        printf("Errors finalizing!\n");
    }

    printf("len = %d\n", len);
}