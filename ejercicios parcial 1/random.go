package parcial1

/*En la ciudad existe una biblioteca tan grande que todos sus visitantes pierden mucho tiempo buscando libros. Para
solucionar este problema, se adquirió un robot recepcionista que pueda darle a las personas la ubicación del libro que
desean leer. Se necesita implementar la función de búsqueda para que dicho robot funcione. Para ello, se dispone de
información con el siguiente formato:
*/
type libro struct {
	titulo    string
	ubicacion string
}

type autor struct {
	nombre string
	libros []libro //alfabeticamente
}

/*
En donde la lista de autores está ordenada alfabéticamente por autor, y cada lista de libros está ordenada alfabéticamente
por título del libro. La función de búsqueda debe tener la siguiente firma:
func buscarLibro(nombreAutor string, tituloLibro string, autores []Autor) string
La función recibe el nombre del autor, el título del libro y una lista de autores con sus respectivos libros. La función
debe devolver la ubicación del libro en la biblioteca. Además, se requiere una solución eficiente que minimice el tiempo
de espera de los usuarios. Indicar y justificar su orden de complejidad. Puede verse al dorso un ejemplo.
*/

func buscarLibro(nombreAutor string, tituloLibro string, autores []autor /*alfabeticamente*/) string /*ubicacion*/ {
	posicion_autor := autorRecursivo(nombreAutor, autores)
	libros := autores[posicion_autor].libros
	posicion_libro := libroRecursivo(tituloLibro, libros)
	return libros[posicion_libro].ubicacion
}

func autorRecursivo(nombreAutor string, autores []autor) int {
	if len(autores) <= 0 {
		return -1
	}
	medio := int(len(autores) / 2)
	if autores[medio].nombre == nombreAutor {
		return medio
	}
	if autores[medio].nombre < nombreAutor {
		autorRecursivo(nombreAutor, autores[medio+1])
	} else {
		autorRecursivo(nombreAutor, autores[medio-1])
	}
	return -1
}

func libroRecursivo(tituloLibro string, libros []libro) int {
	if len(libros) <= 0 {
		return -1
	}
	medio := int(len(libros) / 2)
	if libros[medio].titulo == tituloLibro {
		return medio
	}
	if libros[medio].titulo < tituloLibro {
		autorRecursivo(tituloLibro, libros[medio+1])
	} else {
		autorRecursivo(tituloLibro, libros[medio-1])
	}
	return -1
}

/*Dada una implementación de Pila, implementar una función func AgregarAlFondo[T any](p Pila[T], elem T)
recursiva que agregue el elemento al fondo de la pila sin usar TDAs o estructuras auxiliares. Justificar el orden de
complejidad de la función implementada. Nota: El “fondo” de la pila es tal que el elemento sea el último en ser
desapilado*/

func AgregarAlFondo[T any](p *Pila[T], elem T) {
	if p.EstaVacia() {
		p.Apilar(elem)
		return
	}

	desapilado := p.Desapilar()
	AgregarAlFondo(p, elem)
	p.Apilar(desapilado)
}

/*Implementar un algoritmo de ordenamiento lineal que permita definir el orden en una lista de sabores de helado para
presentar en una nueva heladería (de menor a mayor). Los datos (structs) a ordenar cuentan con el nombre del sabor
(string), popularidad (de 1 a 10, siendo 1 la popularidad más alta) (int), tipo de sabor (chocolate, ddl, etc) (string)
y si es apto para celíacos o no (bool). Primero deben ir los sabores sin TACC (aptos para celíacos), y estos deben
ordenarse por popularidad, independientemente del tipo de sabor. Luego deben ir los no-aptos para celíacos, que deben
estar ordenados por popularidad. Indicar y justificar el orden de complejidad.*/

type helado struct {
	sin_TACC    bool
	popularidad int
	sabor       string
}

func ordenarHelados(helados []helado) []helado {
	var helados_sin_TACC []helado
	var helados_con_TACC []helado
	for _, elem := range helados {
		if elem.sin_TACC {
			helados_sin_TACC = append(helados_sin_TACC, elem)
		} else {
			helados_con_TACC = append(helados_con_TACC, elem)
		}
	} //O(n)
	countingSort(helados_sin_TACC, func(h helado) int { return h.popularidad }, 10) //O(N)
	countingSort(helados_con_TACC, func(h helado) int { return h.popularidad }, 10) //O(n)
	return append(helados_sin_TACC, helados_con_TACC...)
}

func countingSort(helados []helado, popularidad func(helado) int, rango int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]helado, len(helados))
	for _, elem := range helados {
		valor := popularidad(elem)
		frecuencias[valor]++
	}
	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	for _, elem := range helados {
		valor := popularidad(elem)
		posicion := sumasAcumuladas[valor]
		resultado[posicion] = elem
		sumasAcumuladas[valor]++
	}
	copy(helados, resultado)
}
