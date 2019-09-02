# Array Flattening
## if running code locally:

####* Environment Requirements*

Golang 9.0 or above
At least 1.2 MB of free space

#### Install golang on mac

```bash
brew update
brew install golang
```

###### Setup:

```bash
export GOPATH=/your/directory/of/flatten/array/clone
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin

```

#### Install golang on ubuntu

```bash
sudo apt-get update
sudo apt-get -y upgrade
sudo apt-get wget 'if you dont have it already'
wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
sudo tar -xvf go1.12.7.linux-amd64.tar.gz
sudo mv go /usr/local
```

###### Setup:

```bash
export GOROOT=/usr/local/go
export GOPATH=/your/directory/of/flatten/array/clone
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

#### Install golang on windows

[Download Windows Installer here](https://golang.org/dl/ "Download Windows Installer")

### Running the project

```bash
cd /your/directory/of/flatten/array/clone/src
go run main.go
```

### Using the unit tests

After you have followed the above steps and have your environment ready, you can run some local testing utilizing golangs standard testing library "testing".  Use the following steps to run testing in your terminal.

**For mac, linux and windows type the following in the terminal or cmd**:

```bash
cd /your/directory/of/flatten/array/clone/src/testing
go test
```

### API Endpoints

The Flatten Array project offers 2 endpoints, both utilizing json as the data input.  These endpoints are designed to flatten 2 types of arrays, 3d arrays and also standard arrays at variable lengths.

#### Endpoint /v1/flatten_standard

Request Data:

```bash
{
  "array_of_arrays" : [
    [1,2,3,4],
    [5,6,7,8,9],
    [10,11,12,13]
  ]
}
```

Response Data Success:

```bash
{"message" : [1,2,3,4,5,6,7,8,9,10,11,12,13], "status" : 200}
```
Response Data Failure:

```bash
 {"message" : "Bad Request, json improperly formatted", "status" : 400}
```







