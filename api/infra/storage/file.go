package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func Upload(file io.Reader, fileName string) (string, bool) {
	account := os.Getenv("sub_az_account")
	key := os.Getenv("sub_az_key")
	url := fmt.Sprintf("https://%s.blob.core.windows.net/", account)
	ctx := context.Background()

	// Create a default request pipeline using your storage account name and account key.
	credential, err := azblob.NewSharedKeyCredential(account, key)

	if err != nil {
		log.Fatal("AZBlob: Invalid credentials with error: " + err.Error())
		return "", false
	}

	serviceClient, err := azblob.NewServiceClientWithSharedKey(url, credential, nil)
	if err != nil {
		log.Fatal("AZBlob: Invalid credentials with error: " + err.Error())
		return "", false
	}

	containerClient, _ := serviceClient.NewContainerClient("files")

	blobClient, _ := containerClient.NewBlockBlobClient(fileName)
	_, err2 := blobClient.UploadStream(ctx, file, azblob.UploadStreamOptions{})
	if err2 != nil {
		fmt.Println("error")
	}
	return blobClient.URL(), true
}
