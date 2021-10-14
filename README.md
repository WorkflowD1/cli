# Workflow CLI

> This CLI was builted to help Workflow API users to make your work fast. Basically you will can create, update and delete Workflow entities like Products, Modalities, Documents and Attachments.

## Work in progres...

- [X] CLI command to configure instance url, email and password.
- [X] CLI command to create a product.
- [ ] CLI command to create a modality.
- [ ] CLI command to create a document.
- [ ] CLI command to create an attachment.

## How to run develop version

To run this project in your machine, you need to have installed Docker and docker-compose. Clone this repository to your machine and type the following commands:
```sh
docker-compose up
docker-compose exec goapp sh
```

You will open a shell terminal inside created Docker container. After that you are able to run CLI commands using
```sh
go run src/main.go <command> [params]
```