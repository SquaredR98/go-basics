### 1. How to run the code?

```
$ go run main.go
```
- Everything can be done by using the `go` command. To run the code we need to run the above command.

- Commands available with the `go-cli`:
```
$ go build    // Used to just compile and don't run it
$ go run      // Builds and runs the code
$ go fmt      // Formats the code in each file in the directory
$ go install  // Compiles and installs a package
$ go get      // Downloads the raw source code of someone else's package
$ go test     // Runs any test associated with a file
```

### 2. What does `package main` mean?
- A package is a collection of common source code files. Can also be considered as workspace or project.
- There are two different types of packages:
  
  1. **Executable Package**: Generates a runnable or executable file. `main` is required in order to make a package executable otherwise the compiler will skip compiling it.
  2. **Reusable Package**: Package which can used by other package similar to utils.


### 3. What does `import` means?

- Used to give access to some code written in another package. If from the C/C++ ecosystem then it is similar to `#include<iostream>` of if from NodeJS ecosystem it is similar to `import express from 'express'` or `const express = require('express')`

