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
| atan         |      106 |    367 |   28% |
| atan2        |      167 |    782 |   21% |
| exp          |      501 |   2722 |   18% |
| fract        |      206 |    154 |  133% |
| hypot        |       94 |    203 |   46% |
| ln           |      196 |   4892 |    4% |
| powf         |      739 |   9167 |    8% |
| round        |      129 |    171 |   75% |
| sin          |      125 |   1237 |   10% |
| sqrt         |       82 |     57 |  143% |
| tan          |      138 |   1137 |   12% |
| trunc        |       57 |     57 |  100% |

To reproduce: `python3 size_bench.py`
