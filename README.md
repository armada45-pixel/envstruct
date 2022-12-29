# ENV To Struct

Hi everyone I'am so new with golang,
I just learn golang about 1 day and start write this project,
And this is my first project with golang,
I hope this project will useful for you!

This Project use some code from
* [joho/godotenv](https://github.com/joho/godotenv)
* [caarlos0/env](https://github.com/caarlos0/env)

# installation
```shell
go get github.com/armada45-pixel/envstruct
```

# How to use
```go
import (
  "github.com/armada45-pixel/envstruct"
)

type env struct {
	Port uint16 `env:"PORT" os:"PORT" required:"true" default:"1234"`
	Mode string `env:"MODE" os:"MODE" required:"false" default:"testDefalut"`
}
```
## Tag
* **default** is defalut value if read file and not found some or parse value faild.
* **env** is name in env file.
* **os** is name in os.env variable.
* **required** is variable is required or not, not only true, false
```go
func main() {
```
## Secoundary Default Value
```go
	cfg := env{
		Port: 8080, // Default Value is 8080
		// Mode: "Development", // Disable Default Value
	}
```
## Options
For now you can use only 2 option
* **VarPtr** Send pointer of struct variable.
* **FileName** Path and File name to locate env file, *default is ".env"*.
* **IgnoreFile** Don't Read .env file.
* **ReadOS** Read environment variable in OS.
* **OsFirst** If have os.env and env file will choose os.env
```go
	opt := envstruct.Options{
		VarPtr:   &cfg,
		FileName: ".env.local", // can remove this for use default name
		IgnoreFile: false,
		ReadOS: true,
		OsFirst true,
	}
```
## Return
Return array of error.
```go
	if err := envstruct.Setup(opt); len(err) != 0 {
		fmt.Println(err)
	}
```
## Use Value
All value from your env file in here now.
```go
	fmt.Println(cfg)
  fmt.Println(cfg.Port)
  fmt.Println(cfg.Mode)
}

```

# Change Log

1.0.0
* Read File .env and parse value and put into variable(struct only and must pointer)

1.0.1
* Change name of repository to "envstruct"
* Change Mark down installation address

1.0.2
* Change function name ( export )
* Remove main function
* Change module name to envstruct

1.0.3
* Change set function

1.0.4
* Add Read OS Env

1.0.5
* Add Read OS First