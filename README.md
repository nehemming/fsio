# fsio

File system io support functions that support "~" HOME paths

 * [Installation](#install) 
 * [Features](#features)
 * [Contributing](#contrib)
 * [License](#license)

## Status

![Status](https://img.shields.io/badge/Status-COMPLETE-yellow?style=for-the-badge)
[![Build Status](https://img.shields.io/circleci/build/gh/nehemming/fsio/master?style=for-the-badge)](https://github.com/nehemming/fsio) 
[![Release](https://img.shields.io/github/v/release/nehemming/fsio.svg?style=for-the-badge)](https://github.com/nehemming/fsio/releases/latest)
[![Coveralls](https://img.shields.io/coveralls/github/nehemming/fsio?style=for-the-badge)](https://coveralls.io/github/nehemming/fsio)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](#license)
[![GoReportCard](https://goreportcard.com/badge/github.com/nehemming/fsio?test=0&style=for-the-badge)](https://goreportcard.com/report/github.com/nehemming/fsio)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/nehemming/cirocket)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)
[![Uses: cirocket](https://img.shields.io/badge/Uses-cirocket-orange?style=for-the-badge)](https://github.com/nehemming/cirocket)
[![Uses: GoReleaser](https://img.shields.io/badge/uses-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)

## <a name="install">Installation

```bash
go get -u https://github.com/nehemming/fsio
```

or clone this repo to your local machine using

```bash
git clone https://github.com/nehemming/fsio
```

This project requires Go 1.15 or newer and supports modules.

## <a name="features">Key features

 * Support functions to help interact with file systems
 * Supports expansion of paths starting with ~
 * Create nested directories with ~ support
 * Read and write files to / from buffers with ~ support
 * Support expanding file paths relative to another file

## Acknowledgments

  * [Mitchell Hashimoto](https://github.com/mitchellh) for [go-homedir](https://github.com/mitchellh/go-homedir).

## <a name="contrib">Contributing
We would welcome contributions to this project.  Please read our [CONTRIBUTION](https://github.com/nehemming/fsio/blob/master/CONTRIBUTING.md) file for further details on how you can participate or report any issues.

## <a name="license">License

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B26823%2Fgit%40github.com%3Anehemming%2Ffsio.git.svg?type=small)](https://app.fossa.com/projects/custom%2B26823%2Fgit%40github.com%3Anehemming%2Ffsio.git?ref=badge_small)


[MIT](https://choosealicense.com/licenses/mit/)

