## vvdf

`vvdf is a lightweight, zero-dependency Go library for parsing Valve Data Format (VDF / KeyValues) text into Go maps. It provides a robust, single-pass parser designed to handle complex nested structures, inline comments, and non-standard Valve configuration formats safely.

> **Note on Test Coverage / Scope:**  
> *This library is tailored for VPK metadata processing and has been tested specifically against internal VPK configuration files: `addoninfo.txt` and `mission/*.txt`.*

Features
- Single-Pass Parser: Memory-efficient stream parsing using a bottom-up assembly strategy.

- Comment Awareness: Strips // comments cleanly while preserving URLs and quoted text.

- Zero Dependencies: Uses only the Go standard library (bufio, strings, regexp).

- Source Engine Compatible: Robust against irregular whitespace, unquoted tokens, and deeply nested block scopes.


## Installation

go get [github.com/bahyqn/vvdf](https://github.com/bahyqn/vvdf)
```
go get https://github.com/bahyqn/vvdf
```


## Usage

```
import "github.com/bahyqn/vvdf"

func main(){
    text := ""

    tmap, err := vvdf.StringToMap(text)
}
```