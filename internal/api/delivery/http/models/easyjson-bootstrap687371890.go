// +build ignore

// TEMPORARY AUTOGENERATED FILE: easyjson bootstapping code to launch
// the actual generator.

package main

import (
  "fmt"
  "os"

  "github.com/mailru/easyjson/gen"

  pkg "go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
)

func main() {
  g := gen.NewGenerator("authsignup_easyjson.go")
  g.SetPkg("models", "go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models")
  g.DisallowUnknownFields()
  g.Add(pkg.EasyJSON_exporter_UserSignupRequest(nil))
  g.Add(pkg.EasyJSON_exporter_UserSignupResponse(nil))
  if err := g.Run(os.Stdout); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}