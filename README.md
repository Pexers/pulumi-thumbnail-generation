# Introduction
An experimental project using **Pulumi** and **Golang** to deploy a serverless function to Google Cloud Platform (GCP). [**Pulumi**](https://www.pulumi.com/) is a modern infrastructure as code (IaC) platform that allows developers to use familiar programming languages and tools to build, deploy, and manage cloud infrastructure.

The deployed function implements the use-case described below.

## _thumbnail-generation_ use-case

The use case starts with the upload of an image file to be persisted in a storage bucket (_Bucket1_). The upload action triggers the execution of a cloud function responsible for generating and storing a new image thumbnail in a second storage bucket (_Bucket2_). Prior to the thumbnail generation, the function makes a remote call to the provider's storage service to read the bytes of the uploaded image that triggered its execution. The thumbnail generation operation simply consists in cutting the image width in half using the [image package](https://pkg.go.dev/image) from Golang.

<p align="center">
  <img src="https://user-images.githubusercontent.com/47757441/200130281-8b086d3b-06b6-43c0-864d-bd512cc85f84.jpg" width="700">
</p>
