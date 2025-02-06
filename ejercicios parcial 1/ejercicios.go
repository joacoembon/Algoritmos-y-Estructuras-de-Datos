package parcial1

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 3) Implementar una función que reciba un arreglo genérico e invierta su orden, utilizando los TDAs vistos. Indicar y justificar el orden de
//ejecución.

func InvertirArreglo[T any](arr []T) {
	pilaAux := CrearPilaDinamica[T]()
	cant := len(arr)
	for _, elem := range arr {
		pilaAux.Apilar(elem)
	}
	for i := 0; i < cant; i++ {
		arr[i] = pilaAux.Desapilar()
	}
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// Dada una lista enlazada implementada con las siguientes estructuras:

type nodoLista[T any] struct {
	prox *nodoLista[T]
	dato T
}

type ListaEnlazada[T any] struct {
	prim *nodoLista[T]
}

// 5) Escribir una primitiva que reciba una lista y devuelva el elemento que esté a k posiciones del final (el ante-k-último), recorriendo la
//lista una sola vez y sin usar estructuras auxiliares. Considerar que k es siempre menor al largo de la lista. Por ejemplo, si se recibe la
//lista [ 1, 5, 10, 3, 6, 8 ], y k = 4, debe devolver 10. Indicar el orden de complejidad de la primitiva.

func (lista *ListaEnlazada[T]) AnteKUltimo(k int) T {
	separador := lista.prim
	actual := lista.prim
	for i := 0; i < k; i++ {
		separador = separador.prox
	}
	for separador != nil {
		actual = actual.prox
		separador = separador.prox
	}
	return actual.dato
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 6) Dada una pila de punteros a enteros, escribir una función que determine si es piramidal. Una pila de enteros es piramidal si cada
// elemento es menor a su elemento inferior (en el sentido que va desde el tope de la pila hacia el otro extremo). La pila no debe ser modificada.

func EsPiramidal(pila Pila[int]) bool {
	if pila.EstaVacia() {
		return true
	}
	pilaAux := CrearPilaDinamica[int]()

	pilaAux.Apilar(pila.Desapilar())

	for pila.EstaVacia() == false {
		if pila.VerTope() <= pilaAux.VerTope() {
			for pilaAux.EstaVacia() == false {
				pila.Apilar(pilaAux.Desapilar())
			}
			return false
		}
		pilaAux.Apilar(pila.Desapilar())
	}
	for pilaAux.EstaVacia() == false {
		pila.Apilar(pilaAux.Desapilar())
	}
	return true
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 7) Implementar la primitiva func (cola *colaEnlazada[T]) Multiprimeros(k int) []T que dada una cola y un número k, devuelva los primeros k
// elementos de la cola, en el mismo orden en el que habrían salido de la cola. En caso que la cola tenga menos de k elementos. Si hay menos
//elementos que k en la cola, devolver un slice del tamaño de la cola. Indicar y justificar el orden de ejecución del algoritmo.

func (cola *colaEnlazada[T]) Multiprimeros(k int) []T {
	resultados := make([]T, k)
	actual := cola.primero
	if actual == nil {
		return nil
	}
	for i := 0; i < k; i++ {
		resultados[i] = actual.dato
		actual = actual.siguiente
		if actual == nil {
			resultadosMenork := make([]T, i+1)
			copy(resultadosMenork, resultados)
			return resultadosMenork
		}
	}
	return resultados
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 8) Implementar la función func Multiprimeros[T any](cola Cola[T], k int) []T con el mismo comportamiento de la primitiva anterior.

func Multiprimeros[T any](cola Cola[T], k int) []T {
	resultado := make([]T, k)
	colaAux := CrearColaEnlazada[T]()
	for i := 0; i < k; i++ {
		if cola.EstaVacia() {
			for !colaAux.EstaVacia() {
				cola.Encolar(colaAux.Desencolar())
			}
			menorK := make([]T, i)
			copy(menorK, resultado)
			return menorK
		}
		actual := cola.Desencolar()
		resultado[i] = actual
		colaAux.Encolar(actual)
	}
	for !cola.EstaVacia() {
		colaAux.Encolar(cola.Desencolar())
	}
	for !colaAux.EstaVacia() {
		cola.Encolar(colaAux.Desencolar())
	}
	return resultado
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 9) Implementar en Go una primitiva func (lista *listaEnlazada[T]) Invertir() que invierta la lista, sin utilizar estructuras auxiliares.
//Indicar y justificar el orden de la primitiva.

func (lista *ListaEnlazada[T]) Invertir() {
	var previo *nodoLista[T] = nil
	actual := lista.primero
	for actual != nil {
		siguienteTemp := actual.siguiente
		actual.siguiente = previo
		previo = actual
		actual = siguienteTemp
	}
	lista.primero, lista.ultimo = lista.ultimo, lista.primero
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 11) Implementar una función que ordene de manera ascendente una pila de enteros sin conocer su estructura interna y utilizando como
//estructura auxiliar sólo otra pila auxiliar. Por ejemplo, la pila [ 4, 1, 5, 2, 3 ] debe quedar como [ 1, 2, 3, 4, 5 ] (siendo el último
//elemento el tope de la pila, en ambos casos). Indicar y justificar el orden de la función.

func Ordenar(pila Pila[int]) {

	pilaAux := CrearPilaDinamica[int]()

	for pila.EstaVacia() == false {
		aux := pila.Desapilar()
		for pilaAux.EstaVacia() == false && aux > pilaAux.VerTope() {
			pila.Apilar(pilaAux.Desapilar())
		}
		pilaAux.Apilar(aux)
	}

	for pilaAux.EstaVacia() == false {
		pila.Apilar(pilaAux.Desapilar())
	}
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 12) Implementar una función func FiltrarCola[K any](cola Cola[K], filtro func(K) bool), que elimine los elementos encolados para los cuales
//la función filtro devuelve false. Aquellos elementos que no son eliminados deben permanecer en el mismo orden en el que estaban antes de
//invocar a la función. Se pueden utilizar las estructuras auxiliares que se consideren necesarias y no está permitido acceder a la estructura
//interna de la cola (es una función). ¿Cuál es el orden del algoritmo implementado?

func FiltrarCola[K any](cola Cola[K], filtro func(K) bool) {
	colaAux := CrearColaEnlazada[K]()
	for cola.EstaVacia() == false {
		a := cola.Desencolar()
		if filtro(a) {
			colaAux.Encolar(a)
		}
	}
	for colaAux.EstaVacia() == false {
		cola.Encolar(colaAux.Desencolar())
	}
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 13) Sabiendo que la firma del iterador interno de la lista enlazada es:

//Iterar(visitar func(K) bool)

// Se tiene una lista en donde todos los elementos son punteros a números enteros. Implementar una función SumaPares que reciba una lista y,
//utilizando el iterador interno (no el externo), calcule la suma de todos los números pares.

func SumaPares(lista Lista[*int]) int {
	suma := 0
	lista.Iterar(func(dato *int) bool {
		if *dato%2 == 0 {
			suma += *dato
		}
		return true // Continuar iterando
	})
	return suma
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 16) Dadas dos pilas de enteros positivos (con posibles valores repetidos) cuyos elementos fueron ingresados de menor a mayor, se pide
//implementar una función func MergePilas(pila1, pila2 Pila[int]) []int que devuelva un array ordenado de menor a mayor con todos los valores
//de ambas pilas sin repeticiones. Detallar y justificar la complejidad del algoritmo considerando que el tamaño de las pilas es N y M respectivamente.

func MergePilas(pila1, pila2 Pila[int]) []int {
	pilaAux := CrearPilaDinamica[int]()
	cont := 0
	for !pila1.EstaVacia() && !pila2.EstaVacia() {
		if pilaAux.EstaVacia() {
			if pila1.VerTope() > pila2.VerTope() {
				pilaAux.Apilar(pila1.Desapilar())
				cont++
			} else if pila2.VerTope() > pila1.VerTope() {
				pilaAux.Apilar(pila2.Desapilar())
				cont++
			} else {
				pilaAux.Apilar(pila1.Desapilar())
				cont++
			}
		} else if pila1.VerTope() == pilaAux.VerTope() {
			pila1.Desapilar()
		} else if pila2.VerTope() == pilaAux.VerTope() {
			pila2.Desapilar()
		} else if pila1.VerTope() > pila2.VerTope() {
			pilaAux.Apilar(pila1.Desapilar())
			cont++
		} else if pila2.VerTope() >= pila1.VerTope() {
			pilaAux.Apilar(pila2.Desapilar())
			cont++
		}
	}
	for !pila2.EstaVacia() {
		if pilaAux.EstaVacia() {
			pilaAux.Apilar(pila2.Desapilar())
			cont++
		} else if pila2.VerTope() != pilaAux.VerTope() || pilaAux.EstaVacia() {
			pilaAux.Apilar(pila2.Desapilar())
			cont++
		} else {
			pila2.Desapilar()
		}
	}
	for !pila1.EstaVacia() {
		if pilaAux.EstaVacia() {
			pilaAux.Apilar(pila1.Desapilar())
			cont++
		} else if pila1.VerTope() != pilaAux.VerTope() {
			pilaAux.Apilar(pila1.Desapilar())
			cont++
		} else {
			pila1.Desapilar()
		}
	}

	resultados := make([]int, cont)
	for i := 0; !pilaAux.EstaVacia(); i++ {
		resultados[i] = pilaAux.Desapilar()
	}
	return resultados
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 17) Escribir una primitiva para la pila (dinámica) cuya firma es func (pila pilaDinamica[T]) Transformar(aplicar func(T) T) Pila[T] que
//devuelva una nueva pila cuyos elementos sean los resultantes de aplicarle la función aplicar a cada elemento de la pila original. Los
//elementos en la nueva pila deben tener el orden que tenían en la pila original, y la pila original debe quedar en el mismo estado al inicial.
//Indicar y justificar la complejidad de la primitiva.

// Por ejemplo, para la pila de enteros [ 1, 2, 3, 6, 2 ] (tope es el número 2), y la función sumarUno (que devuelve la suma entre el número 1
//y el número recibido), la pila resultante debe ser [ 2, 3, 4, 7, 3 ] (el tope es el número 3).

func (pila *pilaDinamica[T]) Transformar(aplicar func(T) T) *pilaDinamica[T] {
	// para crear una nueva pila en este ejercicio usar, según la capacidad deseada:
	// nueva := &pilaDinamica[T]{make([]T, capacidadDeseada), capacidadDeseada}
	nueva := &pilaDinamica[T]{make([]T, pila.cantidad), pila.cantidad}

	for i := 0; i < pila.cantidad; i++ {
		elemento := pila.datos[i]
		nuevoElemento := aplicar(elemento)
		nueva.datos[i] = nuevoElemento
	}
	return nueva
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 18) Implementar una función recursiva que reciba una pila y devuelva, sin utilizar estructuras auxiliares, la cantidad de elementos de la
//misma. Al terminar la ejecución de la función la pila debe quedar en el mismo estado al original.

func Largo[T any](pila Pila[T]) int {
	if pila.EstaVacia() {
		return 0
	}
	elem := pila.Desapilar()
	cant := 1 + Largo(pila)
	pila.Apilar(elem)

	return cant
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 19) Implementar una función func Balanceado(texto string) bool, que retorne si texto esta balanceado o no. texto sólo puede contener los
//siguientes caracteres: [,],{,}(,). Indicar y justificar la complejidad de la función implementada. Un texto esta balanceado si cada agrupador
//abre y cierra en un orden correcto. Por ejemplo:

// balanceado("[{([])}]") => true
// balanceado("[{}") => false
// balanceado("[(])") => false
// balanceado("()[{}]") => true
// balanceado("()()(())") => true

func Balanceado(texto string) bool {
	pilaAux := CrearPilaDinamica[rune]()
	// Iterar sobre cada caracter en el texto
	// Mapa para mapear los caracteres de cierre con sus respectivos caracteres de apertura
	matching := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, char := range texto {
		// Si es un caracter de apertura, lo apilamos en la pila
		if char == '(' || char == '[' || char == '{' {
			pilaAux.Apilar(char)
		} else if char == ')' || char == ']' || char == '}' {
			// Si es un caracter de cierre
			// Verificamos si la pila está vacía o si el caracter de cierre coincide con el caracter de apertura en el tope de la pila
			if pilaAux.EstaVacia() || pilaAux.VerTope() != matching[char] {
				return false // No está balanceado
			}
			// Desapilamos el caracter de apertura correspondiente
			pilaAux.Desapilar()
		}
	}

	// Si la pila está vacía al finalizar la iteración, el texto está balanceado
	return pilaAux.EstaVacia()
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 0) Implementar un algoritmo en Go que reciba un arreglo de enteros de tamaño nn, ordenado ascendentemente y sin elementos repetidos, y
//determine en O(log n) si es mágico. Un arreglo es mágico si existe algún valor i tal que 0 <= i y arr[i] = i. Justificar el orden del algoritmo.

//Ejemplos:

//A = [ -3, 0, 1, 3, 7, 9 ] es mágico porque A[3] = 3.

//B = [ 1, 2, 4, 6, 7, 9 ] no es mágico porque B[i] != i para todo i.

func ArregloEsMagico(arr []int) bool {
	return busquedaMagica(arr, 0, len(arr)-1)
}

func busquedaMagica(arr []int, inicio, fin int) bool {
	// Caso base: si el inicio es mayor que el fin, no hay mas elementos para buscar
	if inicio > fin {
		return false
	}
	medio := (inicio + fin) / 2

	if arr[medio] == medio {
		return true
	}
	if arr[medio] > medio {
		return busquedaMagica(arr, inicio, medio-1)
	}
	//if arr[medio] < medio {
	return busquedaMagica(arr, medio+1, fin)
	//}
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 8) Implementar, por división y conquista, una función que determine el mínimo de un arreglo. Indicar y justificar el orden.

// BuscarMinimo devuelve el valor del minimo del arreglo, no su posicion
// Precondicion: el arreglo tiene al menos un elemento
func BuscarMinimo(arr []int) int {
	return minimoRecursivo(arr, 0, len(arr)-1)
}

func minimoRecursivo(arr []int, inicio, fin int) int {
	//Caso base: si hay solo un elemento en el sub arreglo
	if inicio == fin {
		return arr[inicio]
	}
	medio := (inicio + fin) / 2

	minIzq := minimoRecursivo(arr, inicio, medio)
	minDer := minimoRecursivo(arr, medio+1, fin)

	return min(minIzq, minDer)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 9) Implementar, por división y conquista, una función que dado un arreglo y su largo, determine si el mismo se encuentra ordenado. Indicar y
//justificar el orden.

func EstaOrdenado(arr []int) bool {
	return estaOrdenadoRecursivo(arr, 0, len(arr)-1)
}

func estaOrdenadoRecursivo(arr []int, inicio, fin int) bool {
	// Caso base: si el subarreglo tiene menos de dos elementos, está ordenado
	if fin-inicio < 1 {
		return true
	}

	// Calcular el índice medio
	medio := (inicio + fin) / 2

	// Verificar si la mitad izquierda está ordenada
	izquierdaOrdenada := estaOrdenadoRecursivo(arr, inicio, medio)

	// Verificar si la mitad derecha está ordenada
	derechaOrdenada := estaOrdenadoRecursivo(arr, medio+1, fin)

	// Combinar la información de ambas mitades
	return izquierdaOrdenada && derechaOrdenada && arr[medio] <= arr[medio+1]
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 10) Implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos y casi ordenado (todos los elementos se
//encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar. Indicar y justificar el orden.

func ElementoDesordenado(arr []int) int {
	return _elementoDesordenado(arr, 0, len(arr)-1)
}

func _elementoDesordenado(arr []int, inicio, fin int) int {
	if inicio == fin {
		return arr[inicio]
	}
	medio := (inicio + fin) / 2
	parIzq := _elementoDesordenado(arr, inicio, medio)
	parDer := _elementoDesordenado(arr, medio+1, fin)

	for i := inicio; i < fin; i++ {
		if arr[i] > arr[i+1] {
			if i > 0 && i < fin-1 && (arr[i+1] < arr[i+2] && arr[i+1] < arr[i-1]) {
				return arr[i+1]
			}
			return arr[i]
		}
	}
	if parIzq > parDer {
		return parIzq
	}
	return parDer
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 11) Se tiene un arreglo tal que [1, 1, 1, …, 0, 0, …] (es decir, unos seguidos de ceros). Se pide una función de orden O(log(n)) que
//encuentre el índice del primer 0. Si no hay ningún 0 (solo hay unos), debe devolver -1.

func IndicePrimeroCero(arr []int) int {
	if arr[0] == 0 {
		return 0
	}
	return _indicePrimeroCero(arr, 0, len(arr)-1)
}

func _indicePrimeroCero(arr []int, inicio, fin int) int {
	if inicio > fin {
		return -1
	}
	medio := (inicio + fin) / 2
	if arr[medio] == 0 && arr[medio-1] != 0 {
		return medio
	} else if arr[medio] == 0 {
		return _indicePrimeroCero(arr, inicio, medio)
	}
	return _indicePrimeroCero(arr, medio+1, fin)
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 12) mplementar un algoritmo que, por división y conquista, permita obtener la parte entera de la raíz cuadrada de un número n, en tiempo
//O(log n). Por ejemplo, para n = 10 debe devolver 3, y para n = 25 debe devolver 5.

func ParteEnteraRaiz(n int) int {
	return _parteEnteraRaiz(n, 0, n)
}

func _parteEnteraRaiz(n, desde, hasta int) int {
	if desde == hasta {
		return desde
	}
	medio := (desde + hasta) / 2
	sig := medio + 1
	if medio*medio <= n && sig*sig > n {
		return medio
	} else if medio*medio < n {
		return _parteEnteraRaiz(n, medio+1, hasta)
	}
	return _parteEnteraRaiz(n, desde, medio)
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 13) Se tiene un arreglo de N >= 3 elementos en forma de pico, esto es: estrictamente creciente hasta una determinada posición p, y
//estrictamente decreciente a partir de ella (con 0 < p < N - 1). Por ejemplo, en el arreglo [1, 2, 3, 1, 0, -2] la posición del pico es p = 2.
// Se pide:

// 1. Implementar un algoritmo de división y conquista de orden O(log n) que encuentre la posición p del pico:
//func PosicionPico(v []int, ini, fin int) int. La función será invocada inicialmente como: PosicionPico(v, 0, len(v)-1), y tiene como
//pre-condición que el arreglo tenga forma de pico.

// 2. Justificar el orden del algoritmo mediante el teorema maestro.

func PosicionPico(v []int, ini, fin int) int {
	medio := (ini + fin) / 2
	if v[medio] > v[medio-1] && v[medio] > v[medio+1] {
		return medio
	} else if v[medio] > v[medio+1] {
		return PosicionPico(v, ini, medio)
	}
	return PosicionPico(v, medio+1, fin)
}

//hay dos llamados recursivos pero se usa uno solo a la vez.
//los llamados recursivos llaman a la mitad
// lo demas que no es recursivo es O(1)
//T(n)= 2T(n/2)+O(1) --> O(log(n))

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 17) Tenemos un arreglo de tamaño 2n de la forma {C1, C2, C3, … Cn, D1, D2, D3, … Dn}, tal que la cantidad total de elementos del arreglo
//es potencia de 2 (por ende, n también lo es). Implementar un algoritmo de División y Conquista que modifique el arreglo de tal forma que
//quede con la forma {C1, D1, C2, D2, C3, D3, …, Cn, Dn}, sin utilizar espacio adicional (obviando el utilizado por la recursividad y
//variables de tipos simples). ¿Cual es la complejidad del algoritmo?

// Pista: Pensar primero cómo habría que hacer si el arreglo tuviera 4 elementos ({C1, C2, D1, D2}). Luego, pensar a partir de allí el caso
//de 8 elementos, etc… para encontrar el patrón.

func Alternar(arr []int) {
	AlternarRec(arr, 0)
}

func AlternarRec(arr []int, iteracion int) {
	largo := len(arr)
	if iteracion*2 == largo {
		return
	}
	pri := arr[iteracion]
	seg := arr[iteracion+largo/2]

	AlternarRec(arr, iteracion+1)

	arr[iteracion*2] = pri
	arr[iteracion*2+1] = seg
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 19) Se sabe, por el teorema de Bolzano, que si una función es continua en un intervalo [a, b], y que en el punto a es positiva y en el
//punto b es negativa (o viceversa), necesariamente debe haber (al menos) una raíz en dicho intervalo. Implementar una
//función func raiz(f func(int)int, a int, b int) int que reciba una función (univariable) y los extremos mencionados y devuelva una raíz
//dentro de dicho intervalo (si hay más de una, simplemente quedarse con una). La complejidad de dicha función debe ser logarítmica del largo
//del intervalo [a, b]. Asumir que por más que se esté trabajando con números enteros, hay raíz en dichos valores: Se puede trabajar con
//floats, y el algoritmo será equivalente, simplemente se plantea con ints para no generar confusiones con la complejidad. Justificar la
//complejidad de la función implementada.

func raiz(f func(int) int, a, b int) int {
	if f(a)*f(b) > 0 {
		panic("No se cumple el teorema de Bolzano en el intervalo dado")
	}
	return raizRecursivo(f, a, b)
}

func raizRecursivo(f func(int) int, a, b int) int {
	if b-a == 1 {
		return a
	}
	medio := (a + b) / 2
	if f(a)*f(medio) < 0 {
		return raizRecursivo(f, a, medio)
	} else if f(a)*f(medio) > 0 {
		return raizRecursivo(f, medio, b)
	}
	return medio
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 20) Es el año 1700, y la pirata Barba-ra Verde atacó un barco de la Royal British Shipping & Something, que transportaba una importante
//piedra preciosa de la corona británica. Al parecer, la escondieron en un cofre con muchas piedras preciosas falsas, en caso de un ataque.
//Barba-ra Verde sabe que los refuerzos británicos no tardarán en llegar, y deben uir lo más rápido posible. El problema es que no pueden
//llevarse el cofre completo por pesar demasiado. Necesita encontrar rápidamente la joya verdadera. La única forma de descubrir la joya
//verdadera es pesando. Se sabe que la joya verdadera va a pesar más que las imitaciones, y que las imitaciones pesan todas lo mismo. Cuenta
//con una balanza de platillos para poder pesarlas (es el 1700, no esperen una balanza digital).

// A) Escribir un algoritmo de división y conquista, para determinar cuál es la verdadera joya de la corona. Suponer que hay una función
//balanza(grupo_de_joyas1, grupo_de_joyas2) que devuelve 0 si ambos grupos pesan lo mismo, mayor a 0 si el grupo1 pesa más que el grupo2,
//y menor que 0 si pasa lo contrario, y realiza esto en tiempo constante.
// B) Indicar y justificar (adecuadamente) la complejidad de la función implementada.

func encontrarJoya(joyas []int) int {
	if len(joyas) == 1 {
		return 0
	}
	return buscar(joyas, 0, len(joyas))
}

func buscar(joyas []int, inicio, fin int) int {
	if fin == inicio+1 {
		return inicio
	}
	resto := (inicio + fin) % 2
	medio := (inicio + fin) / 2
	resultado := balanza(joyas[inicio:medio], joyas[medio+resto:fin])
	if resultado == 0 {
		return medio
	} else if resultado > 0 {
		return buscar(joyas, inicio, medio)
	}
	return buscar(joyas, medio+resto, fin)
}
//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// Encontrar el faltante en los dos arreglos casi iguales

A := []int{2, 4, 6, 8, 9, 10, 12}
B := []int{2, 4, 6, 8, 10, 12}

func encontrar_faltante(A []int, B []int) int {
	if izq < der {
		return A[izq]
	}
	medio := (izq + der)/2
	if A[medio] == B[medio] {
		return encontrar_faltante(A[medio+1:], B[medio+1:])
	}
	return encontrar_faltante(A[:medio],B[:medio])
}


//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// ORDENAMIENTOS NO COMPARATIVOS:

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 9) Se tiene una larga lista de números de tres cifras abc que representan números en notación científica de la forma: a,b * 10^c.
//Por ejemplo 123 representaría el número 1,2 * 10^3.

// Implemente un algoritmo para ordenar los números según su valor en notación científica. ¿De qué orden es?

// Se puede asumir que a nunca será 0 salvo que el número sea efectivamente 0. Es decir, la notación es correcta.

// --> puedo usar RadixSort porque son numeros todos de mismas cifras!
func OrdenarNotacionCientifica(arr []string) []string {
	//llamaremos countingSort una vez por cada digito, empezando por el menos significante
	// en este caso el mas significativo es c porque eleva al 10
	countingSortCien(arr, 10, 1)
	countingSortCien(arr, 10, 0)
	countingSortCien(arr, 10, 2)
	return arr
}

func countingSortCien(elementos []string, rango int, digito int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]string, len(elementos))

	for _, elem := range elementos {
		//el arreglo viene en strings por lo q lo voy a tener q convertir a numero
		valor := int(elem[digito] - '0') //si fueran letras en ves de "0" seria "A"
		//aumento la frecuencia de ese valor
		frecuencias[valor]++
	}
	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	for _, elem := range elementos {
		valor := int(elem[digito] - '0')
		pos := sumasAcumuladas[valor]
		resultado[pos] = elem
		sumasAcumuladas[valor]++
	}
	//Modifico la lista original
	for i := 0; i < len(resultado); i++ {
		elementos[i] = resultado[i]
	}
}

//COMPLEJIDAD: O(N + K) --> O(N + 10) --> O(3*(N + 10)) --> O(N)

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 14) Implementar en Go un algoritmo de RadixSort para ordenar un arreglo de Alumnos (estructuras) en función de sus notas en parcialitos, de
//menor a mayor. Los alumnos tienen su nombre y las notas (numéricas, 0-10) de los 3 parcialitos (tenemos las notas finales). El arreglo debe
//quedar ordenado primero por el promedio de notas. No importan los decimales, nada más si tuvo “promedio 9”, “promedio 8”, etc., es decir la
//parte entera del promedio. Luego, en caso de igualdad en este criterio, los alumnos deben quedar ordenados por la nota del parcialito 1, en
//caso de persistir la igualdad, la del parcialito 2, y finalmente por la del 3. En ningún caso se utiliza el nombre para desempatar. Suponer
//que cualquier algoritmo de ordenamiento auxiliar que se requiera ya se encuentra implementado. Sí justificar por qué se utiliza el o los
//algoritmos auxiliares utilizados, y no otros. Indicar y justificar la complejidad del algoritmo. El desarrollo de la complejidad debe tar
//completo, no se aceptan resultados parciales.

type Alumno struct {
	nombre     string
	p1, p2, p3 int
}

func OrdenarAlumnosRadixSort(arr []Alumno) []Alumno {
	countingSort(arr, 11, obtenerNota3)
	countingSort(arr, 11, obtenerNota2)
	countingSort(arr, 11, obtenerNota1)
	countingSort(arr, 11, obtenerPromedio)
	return arr
}

func obtenerNota3(alumno *Alumno) int {
	return alumno.p3
}

func obtenerNota2(alumno *Alumno) int {
	return alumno.p2
}

func obtenerNota1(alumno *Alumno) int {
	return alumno.p1
}

func obtenerPromedio(alumno *Alumno) int {
	return (alumno.p1 + alumno.p2 + alumno.p3) / 3
}

func countingSort(elementos []Alumno, rango int, nota func(*Alumno) int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultados := make([]Alumno, len(elementos))

	for _, elem := range elementos {
		valor := nota(&elem)
		frecuencias[valor]++
	}
	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	for _, elem := range elementos {
		valor := nota(&elem)
		pos := sumasAcumuladas[valor]
		resultados[pos] = elem
		sumasAcumuladas[valor]++
	}

	for i := 0; i < len(resultados); i++ {
		elementos[i] = resultados[i]
	}
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// 15) Implementar un algoritmo de ordenamiento, que sea lineal, que permita definir el orden en una fila de personas para comprar una Cajita
// CampeónFeliz en un establecimiento de comida rápida. Los datos (structs) a ordenar cuentan con edad (número), nombre (string) y nacionalidad
// (enumerativo, de 32 valores posibles). Primero deben ir los niños (todos con edad menor o igual a 12), y estos deben ordenarse por edad
// (de menor a mayor), independientemente de la nacionalidad. Luego deben ir los "no niños", que primero deben estar ordenados por nacionalidad
// (segundo Francia, por ejemplo) y en caso de igualdad de nacionalidad, por edad, también de menor a mayor.

// Justificar la complejidad de la función implementada. El desarrollo de la complejidad debe estar completo para el problema en cuestión, no se
// aceptarán resultados parciales genéricos.

type Persona struct {
	nombre       string
	nacionalidad Nacionalidad
	edad         int
}

//Ordeno primero por edad a todos y despues por nacionalidad pero que
// los menores de 12 tengan nacionalidad 0

func OrdenarFila(personas []Persona) []Persona {
	countingSort(personas, 101, tipoEdad)
	countingSort(personas, 33, tipoNacio)
	for _, persona := range personas {
		if persona.edad > 12 {
			persona.nacionalidad--
		}
	}
	return personas
}

func countingSort(personas []Persona, rango int, visitar func(persona Persona) int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]Persona, len(personas))

	for _, persona := range personas {
		valor := visitar(persona)
		frecuencias[valor]++
	}
	for i := 1; i < len(personas); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	for _, persona := range personas {
		valor := visitar(persona)
		pos := sumasAcumuladas[valor]
		resultado[pos] = persona
		sumasAcumuladas[valor]++
	}
	for i := 0; i < len(personas); i++ {
		if resultado[i].edad <= 12 {
			resultado[i].nacionalidad = 0
		} else {
			resultado[i].nacionalidad++
		}
		personas[i] = resultado[i]
	}
}

func tipoNacio(persona Persona) int {
	return int(persona.nacionalidad)
}
func tipoEdad(persona Persona) int {
	return persona.edad
}

