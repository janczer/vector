[![Build Status](https://travis-ci.org/janczer/vector.svg?branch=master)](https://travis-ci.org/janczer/vector)
[![codecov](https://codecov.io/gh/janczer/vector/branch/master/graph/badge.svg)](https://codecov.io/gh/janczer/vector)

# Vector

Library for work with vectors.

Methods you can use:

- [Eq](#eq)
- [Add](#add)
- [Sub](#sub)
- Dot
- Scalar
- Magnitude
- Normalize
- Angle
- Paraller
- Orthogonal

## Simple usage

First get the package:

```
$ go get github.com/janczer/vector
```

Now you can add this library to `import` and use it:

```
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

```
	v := vector.New([]float64{1, 2, 3})
	v2 := vector.New([]float64{1, 2, 3})
	fmt.Println(v.Eq(v2))  //true
```

## Add

This method add coordinates of two vectors, and return new `Vector`:

```
	v := vector.New([]float64{1, 2, 4})
	v2 := vector.New([]float64{1, 1, 3})
 	v3 := v.Add(v2) //1+1, 2+1, 4+3
```

## Sub

Method subtract coordinates of two vectors, and return new `Vector`:


```
	v := vector.New([]float64{1, 2, 4})
	v2 := vector.New([]float64{1, 1, 3})
	v3 := v.Sub(v2) //1-1, 2-1, 4-3
```
