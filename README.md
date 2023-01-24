## Introduction
An experimental project using **Pulumi** and **Golang** to deploy a serverless use-case to Google Cloud Platform (GCP). [**Pulumi**](https://www.pulumi.com/) is a modern infrastructure as code (IaC) platform that allows developers to use familiar programming languages and tools to build, deploy, and manage cloud infrastructure.

<p align="center">
  <img src="https://user-images.githubusercontent.com/47757441/207642415-0530896a-9803-4360-afda-61a44a8ed439.png" width="170">
</p>

The deployed function implements the use-case described below.

## _thumbnail-generation_ use-case

The use case starts with the upload of an image file to be persisted in a storage bucket (_Bucket1_). The upload action triggers the execution of a cloud function responsible for generating and storing a new image thumbnail in a second storage bucket (_Bucket2_). Prior to the thumbnail generation, the function makes a remote call to the provider's storage service to read the bytes of the uploaded image that triggered its execution. The thumbnail generation operation simply consists in cutting the image width in half using the [image package](https://pkg.go.dev/image) from Golang.

<p align="center">
  <img src="https://user-images.githubusercontent.com/47757441/200130281-8b086d3b-06b6-43c0-864d-bd512cc85f84.jpg" width="700">
</p>

## Setting up the environment
- Before deploying the use-case to Google Cloud Platform, you will first need to install:
  - [Pulumi](https://www.pulumi.com/docs/get-started/install/)
  - [_gcloud_](https://cloud.google.com/sdk/docs/install) - required for Pulumi to work properly.
  - [Go](https://go.dev/dl/) - required for compilation.

- The use-case also requires the creation of two buckets (_Bucket1_ & _Bucket2_). _Bucket1_ detects changes and triggers function executions. _Bucket2_ simply stores new thumbnails.
  - _Bucket1_ was configured using the `Resource` property from `FunctionEventTriggerArgs`, defined in [_project/main.go_](https://github.com/Pexers/pulumi-thumbnail-generation/blob/main/project/main.go).
  - _Bucket2_ was specified using the `bucket2` variable, defined in [_project/app/main.go_](https://github.com/Pexers/pulumi-thumbnail-generation/blob/main/project/app/main.go)
> **Warning**  
> **Don't use the same bucket for detecting changes and storing thumbnails**. If you do, the cloud function will begin to loop executions.

### Pulumi deployment
1. Inside an empty directory, run the following command to download Pulumi's [_serverless-gcp-go_](https://github.com/pulumi/templates/tree/master/serverless-gcp-go) project template:
```
pulumi new serverless-gcp-go
```
2. Replace _main.go_ and _app/main.go_ with the provided source code.
3. Authenticate and obtain GCP credentials by executing the following command:
```
gcloud auth application-default login
```
4. Run the following command to specify the path to the generated JSON credentials file:
```
pulumi config set gcp:credentials JSON_FILE_PATH
```
5. Deploy the _thumbnail-generation_ use-case using the following _pulumi_ command:
```
pulumi up
```

## References
- [Pulumi - GCP Configuration](https://www.pulumi.com/registry/packages/gcp/installation-configuration/#configuration)
- [Pulumi - GCP Serverless Application](https://www.pulumi.com/templates/serverless-application/gcp/)
- [Pulumi - Function API](https://www.pulumi.com/registry/packages/gcp/api-docs/cloudfunctions/function/)
