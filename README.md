# openurl

[![Build Status](https://travis-ci.org/darkowlzz/openurl.svg?branch=master)](https://travis-ci.org/darkowlzz/openurl)
[![codecov](https://codecov.io/gh/darkowlzz/openurl/branch/master/graph/badge.svg)](https://codecov.io/gh/darkowlzz/openurl)

golang package for opening URLs in default web browser.


## Usage

```
import (
  ...

  "github.com/stevenzack/openurl"
)

if err := openurl.Open("example.com"); err != nil {
  log.Fatal(err)
}
```  
