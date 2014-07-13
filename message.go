package main

import (
  "strings"
  "regexp"
)

type message struct {
  connection *connection
  body []byte
}

func (m *message) ToCommand() (bool, string, []string) {

  isCmd := regexp.MustCompile(`^(\/\w+){1}((\W(\w+))?)+`)

  if isCmd.MatchString(string(m.body)) {
    submatch := isCmd.FindStringSubmatch(string(m.body))
    matches := len(submatch)

    if matches > 0 {
      rep := strings.NewReplacer("/", "")
      name := rep.Replace(submatch[1])
      args := submatch[2:matches]

      return true, name, args

    } else {
      return false, "", nil
    }
  }
  return false, "", nil
}
