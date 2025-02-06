package ejercicios

import "golang.org/x/tools/go/analysis/passes/nilfunc"

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//Implementar un algoritmo que, dado un grafo no dirigido, nos devuelva un ciclo dentro del mismo, si es que los tiene. Indicar el 
// orden del algoritmo.

// EncontrarCiclo debe devolver una lista de vertices que conforman el ciclo
// (en un grafo no dirigido) en orden de recorrido que genera el ciclo. 
// En el segundo ejemplo, debería devolver [A, B, C] (o [B, C, A], etc...). 
// El segundo elemento a devolver es true si se encontró un ciclo, false en caso contrario
func EncontrarCiclo(g Grafo) ([]string, bool) {
	visitados := CrearHash[string,int]()
	padre := CrearHash[string,string]()
	for _,v := range g.ObtenerVertices(){
		if !visitados.Pertenece(v) {
			ciclo := dfs_ciclo(grafo, v, visitados, padre)
			if ciclo != nil {
				return ciclo,true
			}
		}
	}
    return nil, false
}

func dfs_ciclo(g Grafo, v string, visitados, padre *Diccionario[K,V]) []string {
	visitados.Guardar(v, 1)
	for _,w := range grafo.Adyacentes(v){
		if visitados.Pertenece(w) {
			// Si w fue visitado y es padre de v, entonces es la arista de donde
      		// vengo (no es ciclo).
      		// Si no es su padre, esta arista (v, w) cierra un ciclo que empieza
      		// en w.
			if w != padre.Obtener(v) {
				return reconstruir_ciclo(padre, w, v)
			}
		} else {
			padre.Guardar(w,v)
			ciclo := dfs_ciclo(grafo, w, visitados, padre)
			if ciclo != nil {
				return ciclo
			}
		}
	}
	return nil
}

func reconstruir_ciclo(padre *Diccionario[K,V], inicio, fin string) []string {
	v := fin
	camino := CrearListaEnlazada[string]()
	for v != inicio {
		camino.InsertarPrimero(v)
		v = padre.Obtener(v)
	}
	camino.InsertarPrimero(inicio)
	return camino
}

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

// Implementar una función que determine el:
//  a. El grado de todos los vértices de un grafo no dirigido.
//  b. El grado de salida de todos los vértices de un grafo dirigido.
//  c. El grado de entrada de todos los vértices de un grafo dirigido.

func grados(g Grafo) Diccionario[string, int] {
	gradosTotal := CrearHash[string, int]()
	vertices := g.ObtenerVertices()
	for _, vertice := range vertices {
		grados := len(g.Adyacentes(vertice))
		gradosTotal.Guardar(vertice, grados)
	}
	return gradosTotal
}

func gradosEntrada(g Grafo) Diccionario[string, int] {
	gradosEnt := CrearHash[string, int]()
	vertices := g.ObtenerVertices()

	for _, vertice := range vertices {
		adyacentes := g.Adyacentes(vertice)
		for _, ver := range adyacentes {
			if !gradosEnt.Pertenece(ver) {
				gradosEnt.Guardar(ver, 1)
			} else {
				grados := gradosEnt.Obtener(ver)
				gradosEnt.Guardar(ver, grados+1)
			}
		}
	}
	return gradosEnt
}

func gradosSalida(g Grafo) Diccionario[string, int] {
	gradosSal := CrearHash[string, int]()
	vertices := g.ObtenerVertices()

	for _, vertice := range vertices {
		grados := len(g.Adyacentes(vertice))
		gradosSal.Guardar(vertice, grados)
	}
	return gradosSal
}

//Por cada vertice paso por E aristas, entonces la complejidad es O(v + e)  si e != v

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//  a. El grado de v en un grafo no dirigido.
//  b. El grado de salida de v en un grafo dirigido.
//  c. El grado de entrada de v en un grafo dirigido.

func gradosDeV(g Grafo, v string) int {
	grados := len(g.Adyacentes(v))
	return grados
}

func gradosEntradaDeV(g Grafo, v string) int {
	vertices := g.ObtenerVertices()
	grado := 0
	for _, vertice := range vertices {
		adyacentes := g.Adyacentes(vertice)
		for _, ver := range adyacentes {
			if ver == v {
				grado++
			}
		}
	}
	return grado
}

func gradosSalidaDeV(g Grafo, v string) int {
	grados := len(g.Adyacentes(v))
	return grados
}

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//Obtener aristas:

func obtener_aristas(g Grafo) []string {
	//grafo no dirigido
	aristas := CrearListaEnlazada[(T,T)]()
	for _, v := range grafo {
		for _, w := range grafo.adyacentes(v) {
			aristas.InsertarUltimo((v, w))
		}
	}
	return aristas
}

func obtener_aristas2(g Grafo) []string {
	//grafo dirigido
	aristas := CrearListaEnlazada[(T,T)]()
	visitados := CrearHash[K,V]()
	for _, v := range grafo {
		for _, w := range grafo.adyacentes(v) {
			if !visitados.Pertenece(v) {
				aristas.InsertarUltimo((v, w))
			}
		}
		visitados.Guardar(v, 1)
	}
	return aristas
}
//La complejidad de ambos es O(v + e) porque por cada vertice vemos sus aristas.

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//Recorrido bfs

func bfs(grafo, origen){
	visitados := CrearHash[K,V]()
	padres := CrearHash[K,V]()
	orden := CrearHash[K,V]()
	padres.Guardar(origen, nil)
	orden.Guardar(origen,0)
	visitados.Guardar(origen, 1)   //el valor agregado no me interesa xq este dicc lo usare como un conjunto.
	q := CrearColaEnlazada[K]()

	for !q.EstaVacia(){
		v := q.Desencolar()
		for _,w := range grafo.adyacentes(v) {
			if !visitados.Pertenece(v){
				padres.Guardar(w,v)
				orden.Guardar(w, orden.Obtener(v) + 1)
				visitados.Guardar(w, 1)
				q.Encolar(w)
			}
		}
	}
	return padres, orden
}
//Vemos todos los vertices y cada uno de sus adyacencias una y solo una vez --> O(v + e)
//Este recorrido funciona sin importar si el grafo es dirigido o no dirigido, solo que el recorrido sera distinto.

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//Recorrido dfs

func dfs(grafo, origen){
	visitados := CrearHash[K,V]()
	padres := CrearHash[K,V]()
	orden := CrearHash[K,V]()
	padres.Guardar(origen, nil)
	orden.Guardar(origen,0)
	visitados.Guardar(origen, 1)

	_dfs(grafo, origen, visitados, padres, orden)
	return padres, orden
}

def _dfs(grado, v, visitados, padres, orden) {
	for _,w := range grafo.adyacentes(v) {
		if !visitados.Pertenece(v){
			padres.Guardar(w,v)
			orden.Guardar(w, orden.Obtener(v) + 1)
			visitados.Guardar(w, 1)

			_dfs(grado, w, visitados, padres, orden)
		}
	}
}

//Vemos todos los vertices y cada uno de sus adyacencias una y solo una vez --> O(v + e)
//Este recorrido funciona sin importar si el grafo es dirigido o no dirigido, solo que el recorrido sera distinto.

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//Que pasa si mi grado es dirigido, o que pasa si quiero recorrer todo el grafo, o que pasa si mi grafo es No dirigido pero no es 
// conexo, es decir, tiene mas de una componente conexa.

func recorrido_dfs_completo(grafo) {
	visitados := CrearHash[K,V]()
	padres := CrearHash[K,V]()
	orden := CrearHash[K,V]()
	for _,v := range grafo {
		if !visitados.Pertenece(v){
			padres.Guardar(v,nil)
			orden.Guardar(v, 0)
			visitados.Guardar(w, 1)

			_dfs(grado, v, visitados, padres, orden)
		}
	}
	return padres, orden
}
//Seria la misma idea para bfs

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//Reconstruir un camino

func reconstruir_camino(padres, destino) {
	recorrido := CrearListaEnlazada[T]()
	for destino != nil {
		recorrido.InsertarPrimero(destino)
		destino = padres.Obtener(destino)
	}
	return recorrido
}

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//Contar las componentes conexas de un grafo NO dirigido (componentes conexas son de un grafo NO dirigido).
// Grafo dirigido tiene componentes debilmente y fuertemente conexas, que las debilmente conexas son las mismas que las conexas del 
// grafo no digirido.

//lo hare con un dfs, pero en este ejercicio es lo mismo si uso bfs o dfs.

func contar_componentes_conexas(grado) {
	contador := 0
	visitados := CrearHash[K,V]()
	for _, v := range grafo {
		if !visitados.Pertenece(v) {
			contador++
			dfs_comps(grafo, v, visitados)
		}
	}
	return contador
}

func dfs_comps(grafo, vertice, visitados){
	for _, w := range grafo.adyacentes(vertice){
		if !visitados.Pertenece(w){
			visitados.Guardar(w, 1)
			dfs_comps(grafo, w, visitados)
		}
	}
}

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

// Un árbol es un grafo no dirigido que cumple con las siguientes propiedades:

// a. ∥E∥=∥V∥-1

// b. Es acíclico

// c. Es conexo

//Por teorema, si un grafo cumple dos de estas tres condiciones, será árbol (y por consiguiente, cumplirá la tercera). Haciendo uso de 
// ésto (y únicamente de ésto), se pide implementar una función que reciba un grafo no dirigido y determine si se trata de un árbol, o no. 
// Indicar el orden de la función implementada.

sol: La definicion de un arbol en grafos es, para todo v,w pertenceciente a los vertices, existe y es unico el camino de v a w.
     Esto tiene consecuencias:
	 1) tiene que ser aciclico, si tiene ciclos el camino no sera unico.
	 2) tiene que ser conexo. (si no, no habra camino)
	 3) la cantidad de aristas es igual a la cantidad de vertices - 1.
	 Si un grafo cumple con solo dos de estas, necesariamente es arbol y cumple la tercera.
	
func es_conexo(grafo) {
	contador := 0 
	visitados := CrearHash[K,V]()
	for _, v := range grafo {
		if !visitados.Pertenece(v){
			contador++
			if contador > 1{
				return false
			}
			fds_comps(grafo, v, visitados)
		}
	}
	return true
}

def es_arbol(grafo){
	return len(obtener_aristas(grafo)) == len(grafo) - 1 && es_conexo(grafo)
}

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

//Implementar un algoritmo que determine si un grafo no dirigido es conexo o no. Indicar la complejidad del algoritmo si el grafo está 
// implementado con una matriz de adyacencia.

func EsConexo(g Grafo) bool {
	contador := 0 
	visitados := CrearHash[string,int]()
	for _, v := range g.ObtenerVertices() {
		if !visitados.Pertenece(v){
			contador++
			if contador > 1{
				return false
			}
			dfs_comps(g, v, visitados)
		}
	}
	return true
}

func dfs_comps(g Grafo, vertice string, visitados *Diccionario[K,V]) {
	for _, w := range g.Adyacentes(vertice){
		if !visitados.Pertenece(w){
			visitados.Guardar(w, 1)
			dfs_comps(g, w, visitados)
		}
	}
}

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

ORDENAMIENTO TOPOLOGICO:

//bfs:

func gradosEntrada(g Grafo) Diccionario[string, int] {
	gradosEnt := CrearHash[string, int]()
	vertices := g.ObtenerVertices()

	for _, vertice := range vertices {
		adyacentes := g.Adyacentes(vertice)
		for _, ver := range adyacentes {
			if !gradosEnt.Pertenece(ver) {
				gradosEnt.Guardar(ver, 1)
			} else {
				grados := gradosEnt.Obtener(ver)
				gradosEnt.Guardar(ver, grados+1)
			}
		}
	}
	return gradosEnt
}

func topologico_grados(g Grafo) []string{
	g_ent := gradosEntrada(g) //Calculo los grados de entrada de todos los vertices.
	q := CrearColaEnlazada[string]()
	resultado := CrearListaEnlazada[string]()
	vertices := g.ObtenerVertices()

	for _, vertice := range vertices {
		if g_ent.Obtener(vertice) == 0{ // si el grado es 0, encolo
			q.Encolar(v)
		}
	}
	for !q.EstaVacia() {
		v = q.Desencolar()
		resultado.InsertarUltimo(v)
		for _,w := range g.Adyacentes(v){ //A cada uno de los adyacentes del que habiamos guardado, le restamos 1
			g_ent.Guardar(w, g_ent.Obtener(w) - 1)
			if g_ent.Obtener(w) == 0{ // Si ya esta apto para encolar, lo hago.
				q.encolar(w)
			}
		}
	}
	if len(resultado) < len(g) { //no se pudieron hacer todos los vertices por ser ciclo.
		return nil
	}
	return resultado
}

//Todo esto es O(v + e)
//Si usara una pila en vez de una cola, seria lo mismo, todas cumplirian con con los pedido pero tendrian un orden distinto los
// vertices del medio.


//dfs:

func _dfs(g Grafo, v string, visitados *Diccionario[string,int], pila *pila[string]) {
	for _, w := range g.Adyacentes(v){
		if !visitados.Pertenece(v){
			visitados.Guardar(w,1)
			_dfs(g, w, visitados, pila)
		}
	}
	pila.apilar(v)
}

func topologico_dfs(g Grafo) {
	visitados := CrearHash[string, int]()
	pila := CrearPilaDinamica[string]()

	for _,v := range g.ObtenerVertices() {
		if !visitados.Pertence(v){
			visitados.Guardar(v,1)
			_dfs(g, v, visitados, pila)
		}
	}
	return pila_a_lista(pila)
}

func pila_a_lista(pila *pila) []string {
	lista := CrearListaEnlzada[string]()
	for !pila.EstaVacia(){
		lista.InsertarUltimo(pila.Desapilar())
	}
	return lista
}

//Sigue siendo O(v+e) porque pasa por cada adyacente de cada vertice realizando operaciones de O(1).
// Devuelven resultados distintos pero ambos son correctos.

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================

EJ18) Dado un número inicial X se pueden realizar dos tipos de operaciones sobre el número:

Multiplicar por 2
Restarle 1.

Implementar un algoritmo que encuentra la menor cantidad de operaciones a realizar para convertir el número X en el número Y, con tan 
solo las operaciones mencionadas arriba (podemos aplicarlas la cantidad de veces que querramos).

//Menor cantidad de pasos suena a un recorrido bfs, puedo ir construyendo los vertices que vaya necesitando, (el grafo va a estar 
// implicito, no existe)

func camino(x, y int) []int {
	visitados := CrearHash[int,int]()
	visitados.Guardar(x,1)
	padres := CrearHash[int,int]()
	padres.Guardar(x, nil)
	q := CrearColaEnlazada[int]()
	q.Encolar(x)
	for !q.EstaVacia(){
		v := q.Desencolar()
		if v == y {
			return padres
		}
		if !visitados.Pertenece(2 * v) {
			visitados.Guardar(2 * v, 1)
			padres.Guardar(2 * v, v)
			q.Encolar(2 * v)
		}
		if !visitados.Pertenece(v - 1) {
			visitados.Guardar(v - 1, 1)
			padres.Guardar(v - 1, v)
			q.Encolar(v - 1)
		}
	}
	return nil
}

//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================


Caminos Minimos:

func camino_minimo_dijkstra(grafo, origen, destino){
	dist := CrearHash[K,V]()
	padre := CrearHash[K,V]()
	dist.Guardar(origen, 0)
	padre.Guardar(origen, nil)
	q := CrearHeap[]()
}




//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================






//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================










//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================











//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================









//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================







//=====================================================================================================================================
//|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//=====================================================================================================================================