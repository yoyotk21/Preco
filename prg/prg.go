package prg

// I could not really figure out where to put the C code so I stuck it here

/*

#cgo LDFLAGS:  -lcrypto -lssl -lm

#include <openssl/evp.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef unsigned char * ustring;
typedef int * intPtr;

int initContext(EVP_CIPHER_CTX *ctx,  const char *key, const char *iv)
{

    if (1 != EVP_EncryptInit_ex(ctx, EVP_aes_256_cbc(), NULL, (unsigned char *)key, (unsigned char *) iv))
    {
        printf("Errors initializing context!\n");
        EVP_CIPHER_CTX_free(ctx);
        return 0;
    }

    return 1;
}

int encrypt(EVP_CIPHER_CTX *ctx, unsigned char *msg, int msgLen, unsigned char *ans, int *ansLen)
{
    if (1 != EVP_EncryptUpdate(ctx, ans, ansLen, msg, msgLen))
    {
        return 0;
    }

    if (1 != EVP_EncryptFinal_ex(ctx, ans, ansLen))
    {
        return 0;
    }

    return 1;
}

unsigned char* ustr(const char *str)
{
	unsigned char *ans = malloc(sizeof(char) * strlen(str));

	for (int i = 0; i < strlen(str); i++)
	{
		ans[i] = str[i];
	}

	return ans;
}

int* newIntPtr()
{
	return malloc(sizeof(int));
}

void printIntPtr(int *iPtr)
{
	printf("%d\n", *iPtr);
}

unsigned char uStringIndex(unsigned char *us, int index)
{
	return us[index];
}

void freeIntPtr(int *i)
{
	free(i);
}

void freeUString(unsigned char *us)
{
	free(us);
}
*/
import "C"
import "fmt"

func ustr(str string) C.ustring {
	return C.ustr(C.CString(str))
}

func ustrFromLength(length int) C.ustring {
	runeList := make([]rune, length)
	return C.ustr(C.CString(string(runeList)))
}

// uses C evp stuff to generate 16 C chars and convert them to go runes
func get16RandomRunes(key string) []rune {
	ctx := C.EVP_CIPHER_CTX_new()
	defer C.EVP_CIPHER_CTX_free(ctx) 
	// I free all the C stuff using defers so they destruct once the function is finished

	C.initContext(ctx, C.CString(key), C.CString("Initializing Vector"))

	ans := ustrFromLength(50)
	defer C.freeUString(ans)

	i := C.newIntPtr()
	defer C.freeIntPtr(i)

	C.encrypt(ctx, ustr("Hello"), 5, ans, i)
	randomNumbers := make([]rune, int(*i))

	for j := 0; j < int(*i); j++ {
		randomNumbers[j] = rune(C.uStringIndex(ans, C.int(j)))
	}
	return randomNumbers
}

// ok! done with all the C stuff! everything that follow can just be pure go related


// A prg function built on C openssl/evp
// Basically it encrypts initString enough to generate the required
// amound of randomn runes.
// This function needs a string to encrypt however to make it easier I will design 
// another function that takes []uint64 and converts it to a string
func PrgWithString(amountOfRunes int, initString string) []rune {
	amountOfTimes := amountOfRunes / 16
	extraTimes := amountOfRunes % 16

	runeList := make([]rune, 0)

	keyString := initString

	for i := 0; i < amountOfTimes; i++ {
		randomRunes := get16RandomRunes(keyString)
		keyString = string(randomRunes)

		runeList = append(runeList, randomRunes...)
	}

	extraRandomRunes := get16RandomRunes(keyString)

	for i := 0; i < extraTimes; i++ {
		runeList = append(runeList, extraRandomRunes[i])
	}

	return runeList
}

func PrgList(data []uint64) []uint64 {
	dataLength := len(data)

	runeList := make([]rune, 0)

	for _, val := range data {
		runeList = append(runeList, rune(val))
	}

	runeList = PrgWithString(dataLength, string(runeList))

	uInt64List := make([]uint64, dataLength)

	for ind, val := range runeList {
		uInt64List[ind] = uint64(val)
	}

	return uInt64List
}

func Prg(num uint64) []uint64 {
	randomRunes := get16RandomRunes(fmt.Sprint(num))
	uInt64List := make([]uint64, len(randomRunes))

	for ind, val := range randomRunes {
		uInt64List[ind] = uint64(val)
	}

	return uInt64List
}