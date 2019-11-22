# Memorable

[![BuildStatus](https://api.travis-ci.org/syronz/memorable.svg?branch=master)](http://travis-ci.org/syronz/memorable) 
[![ReportCard](https://goreportcard.com/badge/github.com/syronz/memorable)](https://goreportcard.com/report/github.com/syronz/memorable) 
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

Memorable is a golang package for generate simple codes and sorting them from simplest to
complicated. It can also generate variation for specific range of characters.
Use Builder design pattern, inspired by https://gist.github.com/vaskoz/10073335

### Usage
Generate variation of listed chars
```
package main

import (
	"fmt"
	"github.com/syronz/memorable"
	"log"
)

func main() {
	mem, err := memorable.New().Chars("AB").Length(2).Build()
	if err != nil {
		log.Fatal(err.Error())
	}

	// variations type is []struct{ Code string; Mark uint64 }
	variations := mem.Variation()

	for _, v := range variations {
		fmt.Printf("%4d - %v\n", v.Mark, v.Code)
	}
}
```

#### Output will be
```
	1 - AA
	2 - AB
	3 - BA
	4 - BB
```

### Develop package with go module

in the destination app we can call local package by use replace command like below (go.mod):
```
require github.com/syronz/memorable v1.2.0
replace github.com/syronz/memorable => /path/to/the/package

```


License
----

MIT
