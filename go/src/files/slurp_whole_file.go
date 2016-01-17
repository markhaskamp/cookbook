package files

import (
  "io/ioutil"
)


func slurp(f string) ([]byte, error) {

  slurped, err := ioutil.ReadFile(f)
  if err != nil {
    return nil, err
  }

  return slurped,nil
}

