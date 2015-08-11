
package ephemtools

import (
	"time"
	"math"
	//"errors"
)

/**
 * Convert a Gregorian date to a Juliandate (JD)
 */
func GregorianToJulian(inTime time.Time) (float64, error) {

	//Get values from date
	day := inTime.Day()
	month := inTime.Month()
	year := inTime.Year()
	hour := inTime.Hour()
	minute := inTime.Minute()
	second := inTime.Second()

	//Test the month
	if month <= 2 {
		year -= 1
		month -= 12
	}

	// Calculate the Julian date without hour, minutes, seconds
	julianDate := float64(365.25)*float64(year+4716)
	julianDate += float64(30.6001)*float64(month+1)
	julianDate += float64(day)
	julianDate += (2-(float64(year)/100)+((float64((float64(year)/100)/4))))
	julianDate -= float64(1524.5)

	//Add the hour, minutes, seconds informations in the julian date
	jdHour := float64(hour-12)/24
  jdMinute := float64(minute)/1440
	jdSecond := float64(second)/86400

	completeJulianDate := float64(int(julianDate)) + jdHour + jdMinute + jdSecond

	return completeJulianDate, nil
}

/**
 * Convert a Grogorian date to a modified Julian date (MJD)
 */
func GregorianToModifJulian(inTime time.Time) (float64, error) {
	jd, err := GregorianToJulian(inTime)
 	return (jd-float64(2400000.5)), err
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
