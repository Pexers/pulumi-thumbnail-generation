## Introduction
An experimental project using **Pulumi** and **Golang** to deploy a serverless use-case to Google Cloud Platform (GCP). [**Pulumi**](https://www.pulumi.com/) is a modern infrastructure as code (IaC) platform that allows developers to use familiar programming languages and tools to build, deploy, and manage cloud infrastructure.

The deployed function implements the use-case described below.

## _thumbnail-generation_ use-case

The use case starts with the upload of an image file to be persisted in a storage bucket (_Bucket1_). The upload action triggers the execution of a cloud function responsible for generating and storing a new image thumbnail in a second storage bucket (_Bucket2_). Prior to the thumbnail generation, the function makes a remote call to the provider's storage service to read the bytes of the uploaded image that triggered its execution. The thumbnail generation operation simply consists in cutting the image width in half using the [image package](https://pkg.go.dev/image) from Golang.

<p align="center">
  <img src="https://user-images.githubusercontent.com/47757441/200130281-8b086d3b-06b6-43c0-864d-bd512cc85f84.jpg" width="700">
</p>

## Setting up your environment
In order to deploy this use-case to Google Cloud Platform, you will first need to install:
- [Pulumi](https://www.pulumi.com/docs/get-started/install/)
- [_gcloud_](https://cloud.google.com/sdk/docs/install)
- [Go](https://go.dev/dl/)

### Deploy a new template project using Pulumi
1. Inside an empty directory, run the following command to download Pulumi's [_serverless-gcp-go_](https://github.com/pulumi/templates/tree/master/serverless-gcp-go) project template:
```
pulumi new serverless-gcp-go
```
2. Authenticate and obtain GCP credentials by executing the following command:
```
gcloud auth application-default login
```
3. Run the following command to specify the path to the generated JSON credentials file:
```
pulumi config set gcp:credentials <path_to_json_file>
```
4. Once the new project is created and configured, you can deploy it using the following _pulumi_ command:
```
pulumi up
```

## References
- [Pulumi - GCP Configuration](https://www.pulumi.com/registry/packages/gcp/)
- [Pulumi - GCP Serverless Application](https://www.pulumi.com/templates/serverless-application/gcp/)
- [Pulumi - Function API](https://www.pulumi.com/registry/packages/gcp/api-docs/cloudfunctions/function/)
