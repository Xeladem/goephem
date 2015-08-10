package main

import (
	"fmt"
	"time"
	"git.piichigo.fr/gojuliandate"
)


func main() {
	//Test the Gregorian to Julian date
	jd := gojuliandate.GregorianToJulian(time.Now())
	fmt.Println(jd)
	
	//Test the Julian date to Gregorian Date	
	fmt.Println(gojuliandate.JulianToGregorian(jd))
}
