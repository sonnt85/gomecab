# gomecab

CGo bindings for MeCab — a Japanese morphological analyzer.

## Installation

```bash
go get github.com/sonnt85/gomecab
```

> **Requirement:** MeCab and its development headers must be installed on the system (`libmecab-dev` on Debian/Ubuntu).

## Features

- Create a MeCab model with optional option string
- Create lattice objects for sentence parsing
- Create tagger objects for morphological analysis
- Parse a sentence and get the result as a string
- Walk parsed nodes to access surface form, feature, POS id, cost, probabilities, and more

## Usage

```go
package main

import (
    "fmt"

    "github.com/sonnt85/gomecab"
)

func main() {
    m, err := gomecab.New()
    if err != nil {
        panic(err)
    }
    defer m.Destroy()

    tagger, err := m.NewTagger()
    if err != nil {
        panic(err)
    }
    defer tagger.Destroy()

    lattice, err := m.NewLattice("日本語を解析する")
    if err != nil {
        panic(err)
    }
    defer lattice.Destroy()

    // Get parsed string
    fmt.Println(tagger.Parse(lattice))

    // Walk nodes
    node := tagger.ParseToNode(lattice)
    for {
        fmt.Printf("surface=%q feature=%q\n", node.Surface(), node.Feature())
        if err := node.Next(); err != nil {
            break
        }
    }
}
```

## API

### `MeCab`

- `New(option ...string) (*MeCab, error)` — creates a MeCab model
- `(*MeCab).Destroy()` — frees the model
- `(*MeCab).NewLattice(input string) (*Lattice, error)` — creates a lattice for the given input sentence
- `(*MeCab).NewTagger() (*Tagger, error)` — creates a tagger

### `Tagger`

- `(*Tagger).Destroy()` — frees the tagger
- `(*Tagger).Parse(lt *Lattice) string` — parses a lattice and returns the result string
- `(*Tagger).ParseToNode(lt *Lattice) *Node` — parses and returns the BOS node for iteration

### `Lattice`

- `(*Lattice).Destroy()` — frees the lattice and its sentence buffer

### `Node`

- `(*Node).Next() error` — advances to the next node; returns `StopIteration` at end
- `(*Node).Surface() string` — surface form of the token
- `(*Node).Feature() string` — feature string (POS, reading, etc.)
- `(*Node).Id() int` — unique node ID
- `(*Node).Length() int` — surface form length in bytes
- `(*Node).Rlength() int` — length including leading whitespace
- `(*Node).Posid() int` — part-of-speech ID
- `(*Node).Wcost() int` — word cost
- `(*Node).Cost() int` — cumulative best cost
- `(*Node).Prob() float32` — marginal probability
- `(*Node).Alpha() float32` — forward log summation
- `(*Node).Beta() float32` — backward log summation
- `(*Node).StartPos() int` — rune start position in the original text

## Author

**sonnt85** — [thanhson.rf@gmail.com](mailto:thanhson.rf@gmail.com)

## License

MIT License - see [LICENSE](LICENSE) for details.
