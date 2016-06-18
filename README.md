# go-teamcity
TeamCity API client for Golang

[![GoDoc](https://godoc.org/github.com/kapitanov/go-teamcity?status.svg)](https://godoc.org/github.com/kapitanov/go-teamcity)

## Features
* Supports both guest and basic authorization
* Provides read-only access to TeamCity REST API
* This library is totally *not* feature-complete so far

## API
Please have a look at the [GoDoc documentation](https://godoc.org/github.com/kapitanov/go-teamcity) for a detailed API description.

## Example

```go
package main

import (
	"fmt"
	"github.com/kapitanov/go-teamcity"
)

func main() {
  tcClient := teamcity.NewClient("https://teamcity.jetbrains.com", teamcity.GuestAuth())
  
  projects, err := tcClient.GetProjects()
  if err != nil {
    panic(err)
  }
  
  fmt.Printf("List of projects:\n")
  for _, project := range projects {
    fmt.Printf(" * %s\n", project.ID)
  }
}
```
