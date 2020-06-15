package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/pkg/errors"
)

var optMode = flag.String("mode", "e", "e: encryption, d: decryption")
var optRegion = flag.String("region", "ap-northeast-1", "region")

// encryption
var (
	// --key-id example
	// arn:aws:kms:ap-northeast-1:999999999:key/596b5b83-e50b-4b0f-9953-16228fe84b31
	// alias/some-alias
	optKeyID     = flag.String("key-id", "", "encryption: arn or alias of KMS key")
	optKeySpec   = flag.String("key-spec", "AES_256", "encryption: key spec for KMS")
	optPlainText = flag.String("plain-text", "", "encryption: plain text to encrypt. '-' for stdin.")
)

// decryption
var (
	optCipherText = flag.String("cipher-text", "", "decryption: base64 encoded cipher text to decrypt. '-' for stdin.")
	optDataKey    = flag.String("data-key", "", "decryption: base64 encoded encrypted data key")
)

type EncryptionResult struct {
	DataKey    []byte
	CipherText []byte
}

func Encrypt(kmsSvc *kms.KMS, keyID, keySpec string, plainText []byte) (*EncryptionResult, error) {

	out, err := kmsSvc.GenerateDataKey(&kms.GenerateDataKeyInput{
		KeyId:   aws.String(keyID),
		KeySpec: aws.String(keySpec),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "*** kms.GenerateDataKey: keyID=%s, keySpec=%s", keyID, keySpec)
	}

	block, err := aes.NewCipher(out.Plaintext)
	if err != nil {
		return nil, errors.Wrapf(err, "*** aes.NewCipher, blockSize=%d", len(out.Plaintext))
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.Wrapf(err, "*** rand.Reader for %d bytes", len(iv))
	}

	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	ret := &EncryptionResult{
		DataKey:    out.CiphertextBlob,
		CipherText: cipherText,
	}

	return ret, nil
}

func Decrypt(kmsSvc *kms.KMS, dataKey, cipherText []byte) ([]byte, error) {
	out, err := kmsSvc.Decrypt(&kms.DecryptInput{
		CiphertextBlob: dataKey,
	})
	if err != nil {
		return nil, errors.Wrap(err, "*** kms.Decrypt dataKey")
	}

	block, err := aes.NewCipher(out.Plaintext)
	if err != nil {
		return nil, errors.Wrapf(err, "*** aes.NewCipher, blockSize=%d", len(out.Plaintext))
	}

	decryptedText := make([]byte, len(cipherText[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(block, cipherText[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptedText, cipherText[aes.BlockSize:])

	return decryptedText, nil
}

func main() {

	flag.Parse()

	sess := session.Must(session.NewSession())
	kmsSvc := kms.New(sess, aws.NewConfig().WithRegion(*optRegion))

	switch *optMode {
	case "e":
		if len(*optKeyID) == 0 {
			log.Fatalf("--key-id must be specified")
		}
		if len(*optPlainText) == 0 {
			log.Fatalf("--plain must be specified")
		}
		if len(*optKeySpec) == 0 {
			log.Fatalf("--key-spec must be specified")
		}

		var plainText = []byte(*optPlainText)
		if *optPlainText == "-" {
			buf := new(bytes.Buffer)
			io.Copy(buf, os.Stdin)
			plainText = buf.Bytes()
		}

		encResult, err := Encrypt(kmsSvc, *optKeyID, *optKeySpec, plainText)
		if err != nil {
			log.Fatalf("err:=%v", err)
		}
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(&struct {
			DataKey    string `json:"dataKey"`
			CipherText string `json:"cipherText"`
		}{
			DataKey:    base64.StdEncoding.EncodeToString(encResult.DataKey),
			CipherText: base64.StdEncoding.EncodeToString(encResult.CipherText),
		})
		return
	case "d":
		key, err := base64.StdEncoding.DecodeString(*optDataKey)
		if err != nil {
			log.Fatalf("*** base64.DecodeString dataKey: %v", err)
		}

		var cipherText []byte
		if *optCipherText == "-" {
			buf := new(bytes.Buffer)
			io.Copy(buf, os.Stdin)

			dst := make([]byte, len(buf.Bytes()))
			n, err := base64.StdEncoding.Decode(dst, buf.Bytes())
			if err != nil {
				log.Fatalf("*** base64.DecodeString cipherText from stdin: %v", err)
			}
			cipherText = dst[:n]
		} else {
			c, err := base64.StdEncoding.DecodeString(*optCipherText)
			if err != nil {
				log.Fatalf("*** base64.DecodeString cipherText: %v", err)
			}
			cipherText = c
		}

		if len(cipherText) == 0 {
			log.Fatalf("*** cipherText was empty")
		}

		decResult, err := Decrypt(kmsSvc, key, cipherText)
		if err != nil {
			log.Fatalf("err=%v\n", err)
		} else {
			os.Stdout.Write(decResult)
		}
		return
	default:
		log.Fatalf("--mode must be 'e' or 'd': %s", *optMode)
	}
}
