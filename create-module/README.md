https://go.dev/doc/tutorial/create-module

https://go.dev/doc/modules/managing-dependencies#naming_module

*Referencing other Modules*
Refer to package with go.mod. Ex: `go mod edit -replace create-module/greetings=../greetings
Will make sure any import from create-module/greetings will look for this module in ../greetings

`go mod tidy` will make sure all modules are added.