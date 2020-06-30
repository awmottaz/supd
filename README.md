![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/awmottaz/supd?style=flat-square&logo=Go)
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/awmottaz/supd?include_prereleases&logo=GitHub&style=flat-square)
![License MIT](https://img.shields.io/badge/license-MIT-informational?style=flat-square)

`supd` (**s**crum **upd**ate) is a CLI app I made for myself to track my daily updates for work. Every day I like to track my plan for the day and the things I actually did that day.

This app is under development. Progress to a v1 release is being tracked in [this issue](https://github.com/awmottaz/supd/issues/4).

- [Installation](#installation)
  - [Download](#download)
  - [Build](#build)
- [Usage](#usage)

## Installation

### Download

You can download a pre-built binary from the [releases](https://github.com/awmottaz/supd/releases). All pre-release binaries are built for Linux with amd64 architecture.

### Build

To build yourself, you must have Go 1.11+ installed with modules enabled for this to work. [See installation instructions for Go.](http://golang.org/doc/install.html)

Clone this repo and run

- `make` to build a binary in the current directory
- `make install` to install the binary on your system

## Usage

Run `supd help` for current usage instructions. The design for the eventual API can be found in [the API design document](./API-design.md).
