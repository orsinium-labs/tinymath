# 🧮 tinymath

[ [📚 docs](https://pkg.go.dev/github.com/orsinium-labs/tinymath) ] [ [🐙 github](https://github.com/orsinium-labs/tinymath) ]

The fastest Go math library for constrained environments, like microcontrollers or WebAssembly.

* Optimizes for performance and small code size at the cost of precision.
* Uses float32 because most microcontrollers (like [ESP32](https://en.wikipedia.org/wiki/ESP32)) have much faster computation for float32 than for float64.
* Designed and tested to work with both Go and [TinyGo](https://tinygo.org/), hence the name.
* Most algorithms are ported from [micromath](https://github.com/tarcieri/micromath) Rust library.
* Zero dependency.

## 📦 Installation

```bash
go get github.com/orsinium-labs/tinymath
```

## 🔧 Usage

```go
fmt.Println(tinymath.Sin(tinymath.Pi))
```

## 🔬 Size

Here is a comparison of WebAssembly binary size (built with TinyGo) when using tinymath vs stdlib math:

| function     | tinymath | stdlib | ratio |
| ------------ | --------:| ------:| -----:|
| atan2        |      167 |    782 |   21% |
| exp          |      535 |   2722 |   19% |
| fract        |      206 |    154 |  133% |
| hypot        |       94 |    203 |   46% |
| ln           |      196 |   4892 |    4% |
| powf         |      859 |   9167 |    9% |
| round        |      129 |    171 |   75% |
| sin          |      198 |   1237 |   16% |
| trunc        |      136 |     57 |  238% |

To reproduce: `python3 size_bench.py`

The two functions that are bigger in tinymath (but still small!) are the ones that have an optimized wasm-specific assembly in the standard library implementation. We're working on adding such assembly code on our side as well.
