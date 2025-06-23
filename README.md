[![Go](https://github.com/rentifly/nilvalidator/actions/workflows/go.yml/badge.svg)](https://github.com/rentifly/nilvalidator/actions/workflows/go.yml)
# nilvalidator

`nilvalidator` is a minimalistic Go library for validating struct fields using the tag `nilvalidator:"required"`.

It is useful for checking that required dependencies are not nil, especially during initialization.

---

## 🔧 Installation

```bash
go get github.com/rentifly/nilvalidator
```

---

## 🚀 Usage

```go
package main

import (
	"log"
	"github.com/rentifly/nilvalidator"
)

type Deps struct {
	Logger *LoggerType `nilvalidator:"required"`
	Email  string       // not validated
	Store  StoreIface   `nilvalidator:"required"`
}

func main() {
	var deps Deps
	// populate deps

	if err := nilvalidator.ValidateStructNotNil(deps); err != nil {
		log.Fatalf("invalid dependencies: %v", err)
	}
}
```

---

## ✅ Behavior

Only fields tagged with `nilvalidator:"required"` are checked.

The following kinds are supported:

- pointers (`*T`)
- interfaces
- maps, slices, funcs, chans

If any of these fields are nil, the function returns an error like:
`field 'Logger' is nil`.

---

## 🧪 Testing

Run tests with:

```bash
go test ./...
```

---
