# fsio

File system io support functions

[![Build Status](https://circleci.com/gh/nehemming/fsio.svg?style=svg&circle-token=276f5c1da2e59f527fb5e323c3cc5ae69382d50d)](https://github.com/nehemming/fsio)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nehemming/fsio)](https://pkg.go.dev/github.com/nehemming/fsio)
[![GoReportCard](https://goreportcard.com/badge/github.com/nehemming/fsio)](https://goreportcard.com/report/github.com/nehemming/fsio)

## Installation

Install the project to your GOPATH using 

```bash
go get -u https://github.com/nehemming/fsio
```

or clone this repo to your local machine using

```bash
git clone https://github.com/nehemming/fsio
```

This project requires Go 1.15 or newer and supports modules.

## Key features

 * Support functions to help interact with file systems
 * Supports expansion of paths starting with ~
 * Create nested directories with ~ support
 * Read and write files to / from buffers with ~ support
 * Support expanding file paths relative to another file

## Acknowledgments

  * [Mitchell Hashimoto](https://github.com/mitchellh) for [go-homedir](https://github.com/mitchellh/go-homedir).

## Contributing
We would welcome contributions to this project.  Please read our [CONTRIBUTION](https://github.com/nehemming/fsio/blob/master/CONTRIBUTING.md) file for further details on how you can participate or report any issues.

## License
[MIT](https://choosealicense.com/licenses/mit/)