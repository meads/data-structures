# Data Structures

A repo for common data structures implemented in Go.

[![Build Status](https://travis-ci.com/meads/datastructures.svg?branch=master)](https://travis-ci.com/meads/datastructures)

[![codecov](https://codecov.io/gh/meads/datastructures/branch/master/graph/badge.svg?sanitize=true)](https://codecov.io/gh/meads/datastructures)


## Usage
Use the source code or explore the examples like below.
```bash
go run main.go -example trie

```
## Installation

This package repo is using go modules. https://github.com/golang/go/wiki/Modules
It's recommended to use go version 1.11 or greater. If you have not done so already, you may need to export this environment variable in your ~/.profile. e.g. 
```bash 
export GO111MODULE=on
```
don't forget to `source ~/.profile` afterwards to pick up the change

NOTE: removing module support should be as easy as removing the above export from the ~/.profile, save and sourcing the file again. 


## Examples
- trie - Launches a browser tab with a search box that searches a backing Trie data structure that has been loaded with some baseline English dictionary (appears to be European btw).


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
- [MIT](https://choosealicense.com/licenses/mit/)

visit me at:[ mikeads.com](https://mikeads.com/)