
package gojuliandate

import (
	"time"
	//"fmt"
	"math"
)

/**
 * Convert a Gregorian date to a Juliandate
 */
func GregorianToJulian(date time.Time) float64 {
	//Get values from date
	day := float64(date.Day())
	month := float64(date.Month())
	year := float64(date.Year())
	hour := float64(date.Hour())
	minute := float64(date.Minute())
	second := float64(date.Second())
	
	//Test the month
	s := float64(1)
	if(month < 9){
		s = float64(-1)
	}
	
	//
	b := math.Abs(month-9)
	
	//Calculate J1
	j1 := math.Trunc(year+s*math.Trunc(b/7))
	//Calculate J2
	j2 := -math.Trunc((math.Trunc(j1/100)+1)*0.75)
	//Calculate the julian day
	jj := -math.Trunc(7*(math.Trunc((month+9)/12)+year)/4)
	jj += math.Trunc(275*month/9)+day+(j2+2)+367*year
	jj += 1721027
	
	//Calculate the julian date
	jdn := jj + ((hour-12)/24)+(minute/1440)+(second/86400)
	
	return jdn
}

/**
 * Convert a Julian date into a Gregorian date
 * TODO: Calculate the hour
 */
func JulianToGregorian(jdn float64) time.Time {
	
	//Define the constant for the calcul
	y := float64(4716)
	m := float64(2)
	r := float64(4)
	v := float64(3)
	s := float64(153)
	b := float64(274277)
	j := float64(1401)
	n := float64(12)
	p := float64(1461)
	u := float64(5)
	w := float64(2)
	c := float64(-38)
	
	//Step One
	f := jdn+j+(((4*jdn+b)/146097)*3)/4+c
	//Step Two
	e := r*f+v
	//Step three
	g := math.Mod(e,p)/r
	//Step four
	h := u*g+w
	//We can now calculate the day
	day := (math.Mod(h,s))/u
	//The month
	month := math.Mod(h/s+m,n)+1
	//And the year
	year := (e/p)-y+(n+m-month)/n
	
	//Set the location
	location,_ := time.LoadLocation("")
	
	return time.Date(int(year),time.Month(month),int(day),0,0,0,0,location)
	
	
}
