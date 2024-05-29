package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/joho/godotenv"

	"github.com/ibiscum/Go-for-DevOps/chapter15/pkg/helpers"
	"github.com/ibiscum/Go-for-DevOps/chapter15/pkg/mgmt"
)

func init() {
	err := godotenv.Load()
	helpers.HandleErr(err)
}

func main() {
	subscriptionID := helpers.MustGetenv("AZURE_SUBSCRIPTION_ID")
	factory := mgmt.NewStorageFactory(subscriptionID)
	fmt.Println("Starting to build Azure resources...")
	stack := factory.CreateStorageStack(context.Background(), "southcentralus")

	uploadBlobs(stack)
	printSASUris(stack)

	fmt.Println("Press enter to delete the infrastructure.")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	factory.DestroyStorageStack(context.Background(), stack)
}

func uploadBlobs(stack *mgmt.StorageStack) {
	serviceClient := stack.ServiceClient()
	containerClient := serviceClient.ServiceClient().NewContainerClient("jd-imgs")

	fmt.Printf("Creating a new container \"jd-imgs\" in the Storage Account...\n")
	_, err := containerClient.Create(context.Background(), nil)
	helpers.HandleErr(err)

	fmt.Printf("Reading all files ./blobs...\n")
	files, err := os.ReadDir("./blobs")
	helpers.HandleErr(err)
	for _, file := range files {
		fmt.Printf("Uploading file %q to container jd-imgs\n", file.Name())
		blobClient := helpers.HandleErrWithResult(containerClient.NewBlockBlobClient(file.Name()), nil)
		osFile := helpers.HandleErrWithResult(os.Open(path.Join("./blobs", file.Name())))
		_ = helpers.HandleErrWithResult(blobClient.UploadFile(context.Background(), osFile, &azblob.UploadFileOptions{}))
	}
}

func printSASUris(stack *mgmt.StorageStack) {
	serviceClient := stack.ServiceClient()
	containerClient := serviceClient.ServiceClient().NewContainerClient("jd-imgs")

	fmt.Printf("\nGenerating readonly links to blobs that expire in 2 hours...\n")
	files := helpers.HandleErrWithResult(os.ReadDir("./blobs"))
	for _, file := range files {
		blobClient := helpers.HandleErrWithResult(containerClient.NewBlockBlobClient(file.Name()), nil)
		//permissions := azblob.BlobSASPermissions{Read: true}
		permissions := sas.BlobPermissions{Read: true}
		now := time.Now().UTC()

		//sasQuery := helpers.HandleErrWithResult(blobClient.GetSASToken(permissions, now, now.Add(2*time.Hour)), nil)
		sasURL := helpers.HandleErrWithResult(blobClient.GetSASURL(permissions, now.Add(2*time.Hour), nil))
		//fmt.Println(blobClient.URL() + "?" + sasQuery.Encode())
		fmt.Println(sasURL)
	}
}
