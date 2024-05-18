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
