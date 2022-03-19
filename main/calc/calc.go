package cal

// las funciones que se exportan se escriben con PascalCase (1a mayuscula)

// Add recibe 2 parametros y los suma
func Add(a float64, b float64) float64 {
	return a + b
}

// Subtract recibe 2 parametros y los resta
func Subtract(a float64, b float64) float64 {
	return a - b
}

// Multiply recibe 2 parametros y los multiplica
func Multiply(a float64, b float64) float64 {
	return a * b
}

// Divide recibe 2 parametros y los divide
func Divide(a float64, b float64) float64 {
	return a / b
}
