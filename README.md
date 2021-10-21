# Workflow CLI

> This CLI was builted to help Workflow API users to make your work fast. Basically you will can create, update and delete Workflow entities like Products, Modalities, Documents and Attachments.

## Work in progress...

- [X] CLI command to configure instance url, email and password.
- [X] CLI command to create a product.
- [X] CLI command to create a document.
- [ ] CLI command to create a modality.
- [ ] CLI command to create an attachment.
- [ ] Write unit tests for all above.

## Build & Install

To build the binary, you need to go to project root and type the following command:
```sh
make build
```

> Disclaimer: If you don't have Golang installed on your machine, you should build the project inside the docker container. Remember to return to project root on your machine for the following steps.

After build was done, you can run:
```sh
sudo make install
```

> You can also run the installation inside the docker container, but you will only be able to use the CLI there.

## Usage

After built and install, you are ready to use CLI. Just type:
```sh
workflow <command> [flags]
```

## How to run develop version

To run this project in your machine, you need to have installed Docker and docker-compose. Clone this repository to your machine and type the following commands:
```sh
make up
```

to turn docker image on and:

```sh
make sh
```

to open a shell terminal inside created Docker container. After that you are able to run CLI commands using:
```sh
go run src/main.go <command> [flags]
```