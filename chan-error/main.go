package main

import (
	"fmt"
	"time"
)

func main() {
	const max = 10_000_000

	now := time.Now()
	for i := 0; i < max; i++ {
		_ = f10(1)
	}
	fmt.Println(time.Since(now))

	ch := make(chan error, 1)
	now = time.Now()
	for i := 0; i < max; i++ {
		e10(1, ch)
		_ = <-ch
	}
	fmt.Println(time.Since(now))

	now = time.Now()
	for i := 0; i < max; i++ {
		g10(1, ch)
		_ = <-ch
	}
	fmt.Println(time.Since(now))

	now = time.Now()
	for i := 0; i < max; i++ {
		_ = c10(1)
	}
	fmt.Println(time.Since(now))

}

func f10(id int) error {
	id++
	return f9(id)
}
func f9(id int) error {
	id++
	return f8(id)
}
func f8(id int) error {
	id++
	return f7(id)
}
func f7(id int) error {
	id++
	return f6(id)
}
func f6(id int) error {
	id++
	return f5(id)
}
func f5(id int) error {
	id++
	return f4(id)
}
func f4(id int) error {
	id++
	return f3(id)
}
func f3(id int) error {
	id++
	return f2(id)
}
func f2(id int) error {
	id++
	return f1(id)
}
func f1(id int) error {
	id++
	return fmt.Errorf("parachute failed: %d", id)
}

func e10(id int, ch chan error) {
	id++
	e9(id, ch)
}
func e9(id int, ch chan error) {
	id++
	e8(id, ch)
}
func e8(id int, ch chan error) {
	id++
	e7(id, ch)
}
func e7(id int, ch chan error) {
	id++
	e6(id, ch)
}
func e6(id int, ch chan error) {
	id++
	e5(id, ch)
}
func e5(id int, ch chan error) {
	id++
	e4(id, ch)
}
func e4(id int, ch chan error) {
	id++
	e3(id, ch)
}
func e3(id int, ch chan error) {
	id++
	e2(id, ch)
}
func e2(id int, ch chan error) {
	id++
	e1(id, ch)
}
func e1(id int, ch chan error) {
	id++
	ch <- fmt.Errorf("parachute failed: %d", id)
}

func g10(id int, ch chan error) {
	id++
	go g9(id, ch)
}
func g9(id int, ch chan error) {
	id++
	go g8(id, ch)
}
func g8(id int, ch chan error) {
	id++
	go g7(id, ch)
}
func g7(id int, ch chan error) {
	id++
	go g6(id, ch)
}
func g6(id int, ch chan error) {
	id++
	go g5(id, ch)
}
func g5(id int, ch chan error) {
	id++
	go g4(id, ch)
}
func g4(id int, ch chan error) {
	id++
	go g3(id, ch)
}
func g3(id int, ch chan error) {
	id++
	go g2(id, ch)
}
func g2(id int, ch chan error) {
	id++
	go g1(id, ch)
}
func g1(id int, ch chan error) {
	id++
	ch <- fmt.Errorf("parachute failed: %d", id)
}

func c10(id int) string {
	id++
	return c9(id)
}
func c9(id int) string {
	id++
	return c8(id)
}
func c8(id int) string {
	id++
	return c7(id)
}
func c7(id int) string {
	id++
	return c6(id)
}
func c6(id int) string {
	id++
	return c5(id)
}
func c5(id int) string {
	id++
	return c4(id)
}
func c4(id int) string {
	id++
	return c3(id)
}
func c3(id int) string {
	id++
	return c2(id)
}
func c2(id int) string {
	id++
	return c1(id)
}
func c1(id int) string {
	id++
	// return "parachute failed: " + strconv.Itoa(id)
	return fmt.Errorf("parachute failed: %d", id).Error()
}
