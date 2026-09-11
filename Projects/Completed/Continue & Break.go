func printPrimes(max int) {
	for n := 2; n <= max; i++ {
		if n == 2 {
			fmt.Println(i)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}