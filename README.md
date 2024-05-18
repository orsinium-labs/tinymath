# tinymath

The fastest Go math library for constrained environments, like microcontrollers or WebAssembly.

* Optimizes for performance and small code size at the cost of precision.
* Uses float32 because most microcontrollers (like [ESP32](https://en.wikipedia.org/wiki/ESP32)) have much faster computation for float32 than for float64.
* Designed and tested to work with both Go and [TinyGo](https://tinygo.org/), hence the name.
* Most algorithms are ported from [micromath](https://github.com/tarcieri/micromath) Rust library.
* Zero dependency.

## Installation

```bash
go get github.com/orsinium-labs/tinymath
```

## Usage

```go
fmt.Println(tinymath.Sin(tinymath.Pi))
```

## Size

Here is a comparison of WebAssembly binary size (built with TinyGo) when using tinymath vs stdlib math:

| function     | tinymath | stdlib | ratio |
| ------------ | --------:| ------:| ----- |
| atan2        |      167 |    782 |   21% |
| exp          |      489 |   2722 |   17% |
| fract        |      166 |    154 |  107% |
| hypot        |       94 |    203 |   46% |
| ln           |      196 |   4892 |    4% |
| powf         |      813 |   9167 |    8% |
| round        |      129 |    171 |   75% |
| sin          |      198 |   1237 |   16% |
| trunc        |      136 |     57 |  238% |

To reproduce: `python3 size_bench.py`
