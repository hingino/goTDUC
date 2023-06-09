# Tiny Driver Update Checker

A tiny (and multi-platform!) GPU update checker written in Golang.


## Purpose

This package aims to create a quick and silent GPU driver upgrade
experience. Where other packages want to continuously add features,
we aim to only silently download and install the newest drivers for your
GPU with very little system overhead.

### Goals

- [ ] windows compatible
- [ ] linux compatible
- [ ] nvidia compatible
- [ ] amd compatible
- [ ] intel compatible

## Pipeline

To keep things actually tiny, we must adhear to a specific order 
of operations.

- [x] Get GPU information
- [x] Get locally installed driver version
- [ ] Get newest online driver version
- [ ] Compare driver versions, break if already on newest version
- [ ] Ask user if they would like to defer, or download/install
- [ ] If accept download, download newest driver to temp directory
- [ ] Silently run installer as a clean installation

## Known Issus

* Have to use custom version of ghw
    * current online version does not properly pull driver info
    * see [my pull request](https://github.com/jaypipes/ghw/pull/331)
* Only features checked off in the [Pipleine](#pipeline) section are complete

### Arguments (todo)

TDUC.go is designed to support the following arguments:
- [ ]  `--dryrun`: checks if drivers are to be updated, but does
not install updates
- [ ]  `--accept`: automatically grants permission to install the newest
driver (if there is one)
- [ ]  `--silent`: automatically grants permission to install the newest
driver (if there is one) silently in the background
