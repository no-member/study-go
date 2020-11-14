package main

var num1 complex64 = 1 + 2i
var num2 complex64 = 4.2342 + 2.3527i

var num3 complex64 = 1.e+3i
var num4 complex64 = 7.27151e-10i

var num5 complex128 = 1 + 2i
var num6 complex128 = 5.32521e-10 + .12345e+2i

var r1 float32 = real(num1)
var i1 float32 = imag(num1)

var r2 float64 = real(num5)
var i2 float64 = imag(num5)
