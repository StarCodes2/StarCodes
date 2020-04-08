package main

import (
	"fmt"
	"math"
)

type Converter struct {
	Feet       float64
	Centimeter float64
	Seconds    float64
	Minutes    float64
	Celsius    float64
	Fahrenheit float64
	Radian     float64
	Degree     float64
	Kilogram   float64
	Pounds     float64
}

func main() {
	convert := Converter{}
	fmt.Println("Conversions")
	fmt.Printf("45.5cm in Feet is %fft\n", CentimeterToFeet(convert, 45.5))
	fmt.Printf("50t in Feet is %fcm\n", FeetToCentimeter(convert, 50.0))

	fmt.Printf("30Mins in Seconds is %fs\n", MinutesToSeconds(convert, 30.0))
	fmt.Printf("60s in Miliseconds is %fms\n", SecondsToMiliseconds(convert, 60.0))

	fmt.Printf("30Celsius in Fahrenheit is %f\n", CelsiusToFahrenheit(convert, 30.0))
	fmt.Printf("60Fahrenheit in Celsius is %f\n", FahrenheitToCelsius(convert, 60.0))

	fmt.Printf("70Rad in Degree is %f\n", RadianToDegree(convert, 70.0))
	fmt.Printf("90Degree in Radian is %f\n", DegreeToRadian(convert, 90.0))

	fmt.Printf("50KG in Pounds is %fPounds\n", KilogramToPounds(convert, 50.0))
	fmt.Printf("60Pounds in Kilogram is %fKG\n", PoundsToKilogram(convert, 60.0))
}

func FeetToCentimeter(cvr Converter, c float64) float64 {
	cvr.Feet = c
	return cvr.Feet * 30.48
}

func CentimeterToFeet(cvr Converter, c float64) float64 {
	cvr.Centimeter = c
	return cvr.Centimeter / 30.48
}

func MinutesToSeconds(cvr Converter, c float64) float64 {
	cvr.Minutes = c
	return cvr.Minutes * 60
}

func SecondsToMiliseconds(cvr Converter, c float64) float64 {
	cvr.Seconds = c
	return cvr.Seconds * 1000
}

func CelsiusToFahrenheit(cvr Converter, c float64) float64 {
	cvr.Celsius = c
	return (cvr.Celsius * 1.8) + 32
}

func FahrenheitToCelsius(cvr Converter, c float64) float64 {
	cvr.Fahrenheit = c
	return (cvr.Fahrenheit - 32) * 0.555556
}

func RadianToDegree(cvr Converter, c float64) float64 {
	cvr.Radian = c
	return cvr.Radian * (180 / math.Pi)
}

func DegreeToRadian(cvr Converter, c float64) float64 {
	cvr.Degree = c
	return cvr.Degree * (math.Pi / 180)
}

func KilogramToPounds(cvr Converter, c float64) float64 {
	cvr.Kilogram = c
	return cvr.Kilogram * 2.205
}

func PoundsToKilogram(cvr Converter, c float64) float64 {
	cvr.Pounds = c
	return cvr.Pounds / 2.205
}
