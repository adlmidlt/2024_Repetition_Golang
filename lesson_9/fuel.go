package lesson_9

import "fmt"

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%.2f L", l)
}

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%.2f gal", g)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%.2f mL", m)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToLitters() Liters {
	return Liters(m / 1000)
}

func StartFuel() {
	soda := Liters(2)
	fmt.Printf("%.3f liters equals %.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%.3f milliliters equals %.3f gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%.3f gallons equals %.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%.3f gallons equals %.3f milliliters\n", milk, milk.ToMilliliters())

	l := Liters(3)
	fmt.Printf("%.1f liters is %.1f milliliters\n", l, l.ToMilliliters())
	m1 := Milliliters(500)
	fmt.Printf("%.1f milliliters is %.1f liters\n", m1, m1.ToLitters())

	fmt.Printf("%.2f gal\n", Gallons(12.09248342))
	fmt.Printf("%.2f L\n", Liters(12.09248342))
	fmt.Printf("%.2f mL\n", Milliliters(12.09248342))

	fmt.Println(Gallons(12.09248342).String())
	fmt.Println(Liters(12.09248342).String())
	fmt.Println(Milliliters(12.09248342).String())
}
