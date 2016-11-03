#Setting environment for Go development: 

##Download Latest Go Lang

https://golang.org/dl/

##IDE

Get intelliJ IDEA (Community Edition)
https://www.jetbrains.com/idea/download/

Get Go Lang plugin for IntelliJ
https://github.com/go-lang-plugin-org/go-lang-idea-plugin

##Organizing Go code:

https://golang.org/doc/code.html

for example for a user home folder called "user" create a folder called "go"

```
go
    bin/
        hello                          # command executable
        outyet                         # command executable
    pkg/
        linux_amd64/
            github.com/golang/example/
                stringutil.a           # package object
    src/
        github.com/golang/example/ # example is your repo in Github
            .git/                  # Git repository metadata
	    hello/
	        hello.go               # command source
	    outyet/
	        main.go                # command source
	        main_test.go           # test source
	    stringutil/
	        reverse.go             # package source
	        reverse_test.go        # test source
        golang.org/x/image/
            .git/                      # Git repository metadata
	    bmp/
	        reader.go              # package source
	        writer.go              # package source
        ... (many more repositories and packages omitted) ...
```

####env variables

```
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
```

make sure $GOBIN is not conflicting with the $GOPATH/bin, it can be unset if not sure

## Testing framework

Ginkgo BDD Testing Framework
https://onsi.github.io/ginkgo/  

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
```