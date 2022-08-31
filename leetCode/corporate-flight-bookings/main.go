package main

func corpFlightBookings(bookings [][]int, n int) []int {
	temp := make([]int, n+1)
	for i := 0; i < len(bookings); i++ {
		first, last, sets := bookings[i][0], bookings[i][1], bookings[i][2]
		temp[first-1] += sets
		temp[last-1+1] -= sets
	}
	for i := 1; i < len(temp); i++ {
		temp[i] = temp[i-1] + temp[i]
	}
	return temp[:len(temp)-1]
}
func main() {

}
