package book

// CADA COSA QUE SE VAYA A EXPORTAR, YA SEA METODO O INCLUSO LOS ATRIBUTOS DE UN STRUCT, DEBEN EMPEZAR POR MAYUSCULAS
// UN STRUCT DENTRO DE UNA MISMA CLASE, PUEDE EMPEZAR EN MINUSCULAS, PERO UNO DE FUERA, EN MAYUSCULAS

type FantasyBook struct {
	Title      string
	Author     string
	CopiesSold int
}

func SayHi() string {
	return "hi"
}
