## Example Todo App using GoCondor framework
This repository contains the code for an example todo app using [GoCondor framework](https://github.com/gocondor/gocondor)

### How to run locally?
1- Clone the repository  
2- Next add your database credentials (mysql) to the `.env` file
3- `cd` into the project dir and run `go mod tidy` to install any missing dependency 
4- Run the app using GoCondor's cli tool [Gaffer](https://gocondor.github.io/docs/gaffer)
```bash
 gaffer run:dev
```
if [Gaffer](https://gocondor.github.io/docs/gaffer) is not installed you can install it by executing the following command
```bash
go install github.com/gocondor/gaffer@latest

```

All routers are defined in the file `routes.go`

All request handlers are defined in the directory `handlers/`