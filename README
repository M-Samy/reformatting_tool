# Formatting Tool
A tool that converts the data from one format to other formats.

## About Tool
* A tool that converts data from one format to other formats, note that this tool till now supports converting from "CSV"
to "JSON and XML" formats.
* Tool designed to be scalable , Developed over factory design pattern trying to be opened to support other multiple formats.
* In case we need to support another format we only need to add another write-extension as "XmlExtension and JsonExtension"
to implement the main writer-interface file "ExtensionsWriterInterface".
* In case we need to support another reading format with "CSV" we only need to add another read-extension as "CsvExtension"
to implement the main read-interface "ExtensionReaderInterface" and another formatter file as "CsvFormatter".
* Tool support service container to be simple as possible.
* Tool extensible to new output formats.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

```
Go 1.13.7
Docker 19.03.5
Docker Compose 1.23.1
```

### Installing
Steps that tells you how to get a build and run this tool.

* Go into reformatting_tool directory.
```
cd reformatting_tool
```

* Go to .env file and write which extensions you need to convert to. "This tool now supports only json and xml formats"

* Run docker-compose.yml file.
```
docker-compose up --build
```

## Built With
* [Go](https://golang.org/)
* [Docker](https://docs.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)

## Features
* Service Container.
* Factory Design Pattern.
* Validations.
* Scalable design.

## Missing Features
* Apply more test cases.
* Add multiple read formats as CSV.
* Apply sort/group data functionality.
