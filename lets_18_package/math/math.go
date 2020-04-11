package math

/*
	闭包高阶函数
*/
func SimpleCompute(Symbol string) func(n, m int) int {
	switch Symbol {
	case "+":
		return func(n, m int) int {
			return n + m
		}
	case "-":
		return func(n, m int) int {
			return n - m
		}
	case "*":
		return func(n, m int) int {
			return n * m
		}
	case "/":
		return func(n, m int) int {
			return n / m
		}
	default:
		panic("compute symbol invalid.")
	}
}
