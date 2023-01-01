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
  Mode modeE `env:"MODE" os:"MODE" required:"false" default:"true"`
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
* **OsFirst** If have os.env and env file will choose os.env.
* **ReadAll** Read all in file and put to os.env ( if PutToOs true )
* **PutToOs** Put variable from file to os.env variable.
* **OverRide** If already variable in same name in os.env variable will replace.
* **OsFirst** If have os.env and env file will choose os.env.
* **CustomType** Map for write custom default value and parser function.
```go
  opt := envstruct.Options{
    VarPtr:   &cfg,
    FileName: ".env.local", // can remove this for use default name
    IgnoreFile: false,
    ReadOS: true,
    ReadAll: true,
    OsFirst: false,
    PutToOs: true,
    OverRide: true,
    OsFirst: true,
    CustomType: map[reflect.Type]TypeDefaultBy{
      reflect.TypeOf(Pro): {
        ParserFunc: func(v string) (interface{}, error) {
          searchDefault, _ := envstruct.DefaultByType[reflect.TypeOf(false)]
          value, _ := searchDefault.ParserFunc(v)
          if value == true {
            return Pro, nil
          }
          return Dev, nil
        },
        ValueDefault: modeE(false),
      },
    },
  }

  type modeE bool

  const (
    Pro modeE = false
    Dev modeE = true
  )

  func (m modeE) String() string {
    if !m {
      return "production"
    }
    return "development"
  }

  func (m modeE) Bool() bool {
    return bool(m)
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
* [a6aa3b2](https://github.com/armada45-pixel/envstruct/commit/a6aa3b27764cdb3c0d810118b258251e60e38b9d#)

1.0.3
* Change set function [0e6505f](https://github.com/armada45-pixel/envstruct/commit/0e6505f9f6297885737fdd774372e0e3f1a0b1d8)

1.0.4
* Add Read OS Env [#4](https://github.com/armada45-pixel/envstruct/issues/4)

1.0.5
* Add Read OS First [#23](https://github.com/armada45-pixel/envstruct/issues/23)

1.0.6
* Add Read All in File [#8](https://github.com/armada45-pixel/envstruct/issues/8)
* Put to OS [#5](https://github.com/armada45-pixel/envstruct/issues/5)
* Over Ride OS [#6](https://github.com/armada45-pixel/envstruct/issues/6)

1.0.7
* Add Custom Parser function [#9](https://github.com/armada45-pixel/envstruct/issues/9)
* Fix bug : If open only "Put to os" And "Read all" will not working. [#26](https://github.com/armada45-pixel/envstruct/issues/26)
* Fix bug : If have any error in process prepare Variable Pointer Data won't set. [#27](https://github.com/armada45-pixel/envstruct/issues/27)