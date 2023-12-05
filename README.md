# Blockchain

## initialize a module
```sh
$ go mod init github.com/KTra-sssc/Blockchain
```
## add module requirements and sums
```sh
$ go mod tidy
```
## compile the code into an executable
```sh
$ go build
```
## run executable file
```sh
$ .\blockchain.exe {{command}}
```
E.g: Print all blocks of the blockchain named chain1
```sh
$ .\blockchain.exe -print chain1
```