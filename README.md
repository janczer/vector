[![Build Status](https://travis-ci.org/janczer/vector.svg?branch=master)](https://travis-ci.org/janczer/vector)
[![codecov](https://codecov.io/gh/janczer/vector/branch/master/graph/badge.svg)](https://codecov.io/gh/janczer/vector)

# Vector

Library for work with vectors.

Methods you can use:

- [Eq](#eq)
- [Add](#add)
- [Sub](#sub)
- [Dot](#dot)
- [Scalar](#scalar)
- [Magnitude](#magnitude)
- [Normalize](#normalize)
- [Angle](#angle)
- [Paraller](#paraller)
- [Orthogonal](#orthogonal)

## Simple usage

First get the package:

```
$ go get github.com/janczer/vector
```

Now you can add this library to `import` and use it:

```Go
package main

import (
	"fmt"
	"github.com/janczer/vector"
)

func main() {
	v := vector.New([]float64{1, 2, 3})
	v2 := vector.New([]float64{1, 2, 4})
	fmt.Println(v.Print()) //Vector: 1, 2, 3
	fmt.Println(v.Eq(v2))  //false
}

```

## Eq

Method return `true` if vectors are equal:

```Go
	v := vector.New([]float64{1, 2, 3})
	v2 := vector.New([]float64{1, 2, 3})
	fmt.Println(v.Eq(v2))  //true
```

## Add

This method add coordinates of two vectors, and return new `Vector`:

```Go
	v := vector.New([]float64{1, 2, 4})
	v2 := vector.New([]float64{1, 1, 3})
 	v3 := v.Add(v2) //1+1, 2+1, 4+3
```

## Sub

Method subtract coordinates of two vectors, and return new `Vector`:

```Go
	v := vector.New([]float64{1, 2, 4})
	v2 := vector.New([]float64{1, 1, 3})
	v3 := v.Sub(v2) //1-1, 2-1, 4-3
```

## Dot

Method multiply coordinates of two vectors, and return new `Vector`:

```Go
	v := vector.New([]float64{1, 2, 4})
	v2 := vector.New([]float64{1, 1, 3})
	v3 := v.Dot(v2) //1, 2, 12
```

## Scalar 

Method multiply each coordinate by number:

```Go
	v := vector.New([]float64{1, 2, 4})
	v2 := v.Scalar(3) //3, 6, 12
```

## Magnitude 

Method return magnitude of vector:

```Go
	v := vector.New([]float64{1, 2, 4})
	m := v.Magnitude() //4,5825..
```

## Normalize 

Method return new unit vector:

```Go
	v := vector.New([]float64{1, 2, 4})
	m := v.Normalize() // 1/4,5825, 2/4.5825, 4/4.5825
```

## Angle 

Method return angle between two vectors:

```Go
	v := vector.New([]float64{3, 0})
	v2 := vector.New([]float64{0, 3})
	a := v.Angle(v2, true) // 45
```

## Paraller 

Method return `true` if two vectors are paraller:

```Go
	v := vector.New([]float64{3, 0})
	v2 := vector.New([]float64{0, 3})
	a := v.Paraller(v2) // false
```

## Orthogonal 

Method return `true` if two vectors are orthogonal:

```Go
	v := vector.New([]float64{3, 0})
	v2 := vector.New([]float64{0, 3})
	a := v.Orthogonal(v2) // true
```
