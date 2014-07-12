package main

import (
  "strings"
  "regexp"
  "log"
)

type message struct {
  connection *connection
  body []byte
}

func (m *message) ToCommand() (bool, string, []string) {

  isCmd := regexp.MustCompile(`(\/\w+){1}((\W(\w+))?)+`)

  if isCmd.MatchString(string(m.body)) {
    submatch := isCmd.FindStringSubmatch(string(m.body))
    matches := len(submatch)

    if matches > 0 {

      rep := strings.NewReplacer("/", "")
      log.Print("Submatch", len(submatch))

      name := rep.Replace(submatch[1])
      log.Println("Command name here", name)

      args := submatch[2:matches]


      return true, name, args

    } else {
      return false, "", nil
    }
  }
  return false, "", nil
}
