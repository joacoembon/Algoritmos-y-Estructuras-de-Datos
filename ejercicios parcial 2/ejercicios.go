package main

import (
	"math"
	"tdas/diccionario"
	"tdas/lista"
)

const HASH string 

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Para un hash cerrado, implementar una primitiva func (hash *hashCerrado[K, V]) Claves() Lista[K] que devuelva una lista con sus claves, 
// sin utilizar el iterador interno.

func (hash *hashCerrado[K, V]) Claves() Lista[K] {

	listaClaves := CrearListaEnlazada[K]()

	for _, celda := range hash.tabla {
		
		if celda.estado == OCUPADO {
			listaClaves.InsertarUltimo(celda.clave)
		}
	}

    return listaClaves
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Para un hash abierto, implementar una primitiva func (hash *hashCerrado[K, V]) Claves() Lista[K] que devuelva una lista con sus claves, 
// sin utilizar el iterador interno.

func (hash *hashAbierto[K, V]) Claves() Lista[K] {
	
	listaClaves := CrearListaEnlazada[K]()
	// Iterar sobre cada lista en la tabla
	for _, lista := range hash.tabla {

		// Iterar sobre cada elemento en la lista actual
		lista.Iterar(func(par parClaveValor[K, V]) bool {

			// Agregar la clave del elemento actual a la lista de claves
			listaClaves.InsertarUltimo(par.clave)
			return true // Continuar iterando
			
		})
	}

	return listaClaves
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Implementar una función func Claves[K comparable, V any](Diccionario[K, V]) Lista[K] que reciba un diccionario y devuelva una lista con sus claves.

func Claves[K comparable, V any](dic Diccionario[K, V]) Lista[K] {
	listaClaves := CrearListaEnlazada[K]()

	dic.Iterar(func(clave K, dato V) bool {
		listaClaves.InsertarUltimo(clave)
		return true
	})

    return listaClaves
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Suponer que se tiene un hash cerrado que se redimensiona cuando el factor de carga llega a 0.75, y que no se tienen en cuenta los elementos 
// borrados a la hora de calcular el factor de carga.

  // A) Describir, en términos generales, el peor escenario posible para esta implementación.

  // B) Dado un hash de estas características, con capacidad inicial 100, calcular el número máximo de casillas que podría llegar a visitar
  //    hash_obtener() si la cantidad actual de elementos en el hash es 1, y no se realizó ningúna redimensión, pero sí se insertaron y 
  //    borraron elementos. (En otras palabras, poner una cota superior al caso peor de este hash.)


  //RESPUESTA: El peor escenario es que se llene el hash de borrados y no se redimensione, esto colapsaria a la hora de agregar nuevos elementos.
  // Para el B) el peor caso seria que visite 99 casillas hasta encontrar el elemento.


//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// ¿Para qué casos la función hash_obtener() tiene una complejidad peor que O(1)? Explicar tanto para el hash abierto, como el cerrado.

//RESPUESTA: Cuando la funcion de hashing es mala.


//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//  Justificar si la siguiente función de hashing es correcta o no:

func calcularHash(string clave) int {
    // rand.Intn(x) devuelve un numero entero entre 0 y x
    return rand.Intn(10000)
}

//RESPUESTA: NO, porque no te va a devolver el mismo indice para la misma clave siempre que la invoques, ya que no depende de ella.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// Implementar una función de orden O(n) que dado un arreglo de n números enteros devuelva true o false según si existe algún elemento
// que aparezca más de la mitad de las veces. Justificar el orden de la solución. Ejemplos:

[1, 2, 1, 2, 3] -> false
[1, 1, 2, 3] -> false
[1, 2, 3, 1, 1, 1] -> true
[1] -> true

func MasDeLaMitad(arr []int) bool {
    contador := CrearHash[int, int] ()

    for _, elem := range arr {
        veces := 1

        if contador.Pertenece(elem) {
            veces = contador.Obtener(elem) + 1
        }
        contador.Guardar(elem, veces)

        if veces > len(arr)/2 {
            return true
        }

    }

    return false
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Asumiendo que se tiene disponible una implementación completa del TDA Hash, se desea implementar una función que decida si dos Hash dados 
// representan o no el mismo Diccionario. Considere para la solución que es de interés la mejor eficiencia temporal posible. Indique, para su 
// solución, eficiencia en tiempo y espacio. Nota: Dos tablas de hash representan el mismo diccionario si tienen la misma cantidad de elementos; 
// todas las claves del primero están en el segundo; todas las del segundo, en el primero; y los datos asociados a cada una de esas claves son 
// iguales (se pueden comparar los valores con “==”).

func SonIguales[K comparable, V comparable](d1, d2 Diccionario[K, V]) bool {
    if d1.Cantidad() != d2.Cantidad() {
        return false
    }
    d1.Iterar(func(clave K, dato V) bool {
		if d2.Pertenece(clave){
            if dato != d2.Obtener(clave){
                return false
            }
            return true
        }
		return false
	})
    return true
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Implementar el TDA MultiConjunto. Este es un Conjunto que permite más de una aparición de un elemento, por lo que eliminando una aparición,
// el elemento puede seguir perteneciendo. Dicho TDA debe tener como primitivas:

   // CrearMulticonj[K](): crea un multiconjunto.
   // Guardar(elem K): guarda un elemento en el multiconjunto.
   // Pertence(elem K) bool: devuelve true si el elemento aparece al menos una vez en el conjunto.
   // Borrar(elem K) bool: elimina una aparición del elemento dentro del conjunto. Devuelve true si se eliminó una aparición del elemento.

//Dar la estructura del TDA y la implementación de las 4 primitivas marcadas, de forma tal que todas sean O(1).

type multiConj[K comparable] struct {
    dicc Diccionario[K, int]
}

func CrearMulticonjunto[K comparable]() MultiConjunto[K] {
    conjunto := new(multiConj[K])
    conjunto.dicc = CrearHash[K,int]()
    return conjunto
}

func (conj multiConj[K]) Guardar(elem K) {
    if conj.dicc.Pertenece(elem) {
        cont := conj.dicc.Obtener(elem)
        conj.dicc.Guardar(elem, cont + 1)
    } else {
        conj.dicc.Guardar(elem, 1)
    }
}

func (conj multiConj[K]) Pertenece(elem K) bool {
    return conj.dicc.Pertenece(elem)
}

func (conj multiConj[K]) Borrar(elem K) {
    if !conj.dicc.Pertenece(elem) {
        panic("Elemento no esta en el multiconjunto")
    }
    cont := conj.dicc.Obtener(elem)
    if cont == 1 {
        conj.dicc.Borrar(elem)
    } else {
        conj.dicc.Guardar(elem, cont - 1)
    }
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Se tiene un hash que cuenta con una función de hashing que, recibida una clave, devuelve la posición de su inicial en el abecedario. La 
// capacidad inicial del hash es 26. Para los puntos B, C y D indicar y justificar si las afirmaciones son verdaderas o falsas. Se puede 
// considerar que todas las claves serán palabras (sólo se usan letras para las claves).

  // A) Mostrar cómo quedan un hash abierto y un hash cerrado (sólo el resultado final) tras guardar las siguientes claves: Ambulancia (0), 
  //    Gato (6), Manzana (12), Ananá (0), Girasol (6), Zapato (25), Zapallo (25), Manzana (12), Bolso (1). Aclaración: No es necesario 
  //    hacer una tabla de 26 posiciones, lo importante es que quede claro en cuál posición está cada elemento.

  // B) En un hash abierto con dicha función de hashing, se decide redimensionar cuando la cantidad alcanza la capacidad (factor de carga = 1). 
  //    El rendimiento de hash_obtener() es mejor en este caso respecto a si se redimensionara al alcanzar un factor de carga 2.

  // C) En un hash cerrado con dicha función de hashing, si se insertan n + 1 claves diferentes (considerar que se haya redimensionado acordemente),
  //    n con la misma letra inicial, y 1 con otra distinta, en el primer caso Obtener() es O(n)O(n) y en el segundo siempre O(1)O(1).

  // D) En un hash abierto con dicha función de hashing, si se insertan n + 1 claves diferentes (considerar que se haya redimensionado acordemente), 
  //    n con la misma letra inicial, y 1 con otra distinta, en el primer caso Obtener() es O(n)O(n) y en el segundo siempre O(1)O(1).

  
  //RESPUESTA: Se puede considerar que la función de hashing siempre va a devolver las mismas posiciones y que la redimensión no surge efecto en la 
  //           distribucion de claves. En ese caso:
  //           B) Sería falso, al redimensionar, las claves siguen en el mismo lugar. Si tenía capacidad 26 y ahora tengo 52 porque redimensione, 
  //              las claves van a seguir en esos 26 debido a la funcion de hashing.
  //           C) Falso, suponiendo que n + 1 >> 26, esas n claves van a ser constantemente desplazadas hasta encontrar una celda vacia debido a 
  //              que la funcion de hashing devuelve siempre la misma posicion para las n claves. Si uso Obtener(), tengo que potencialmente 
  //              recorrer n claves O(n). Y en el caso de la clave distinta, voy a tener que partir de la posicion que me devuelva la funcion de 
  //              hashing, donde de ahi en adelante hay potencialmente n claves que fueron desplazadas de su posicion original. Por ende tambien es O(n)
  //           D) Verdadero, en el caso de que quiera buscar una de las n claves iguales, voy a tener que acceder a la posición en la tabla del hash 
  //              abierto e iterar hasta encontrarla O(n). En cambio, para la clave distina, voy a acceder a su posición y ya la primera posición de 
  //              la lista contiene a la clave O(1).

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//El Ing. Musumeci quiere implementar un hash abierto, pero en el que las listas de cada posición se encuentren ordenadas por clave (se le 
// pasa por parmámetro la función de comparación, por ejemplo strcmp). Explicar cómo mejora o empeora respecto a la versión que vemos en clase 
// para el caso de inserciones, borrados, búsquedas con éxito (el elemento se encuentra en el hash) y sin éxito (no se encuentra).

//RESPUESTA:Las inserciones empeorarian ya que en el peor de los casos deberia recorrerse todo el arreglo para insertar al final, mientras que 
//          en nuestra implementacion siempre se inserta en el ultimo lugar (o siempre en el primero).

//          No modifica el proceso de busquedas con exito ya que no podemos utilizar esa informacion de que la lista esta ordenada para buscar 
//          mas eficientemente.

//          Tampoco modifica el proceso de borrar, ya que de por si se recorria la lista para eliminar el elemento.

//          Lo que si mejoraria serian las busquedas sin exito, porque ya no seria necesario recorrer todo el arreglo para darse cuenta que un 
//          elemento no esta.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//En un diccionario todas las claves tienen que ser diferentes, no así sus valores. Escribir en Go una primitiva para el hash cerrado 
// func (dicc *hashCerrado[K, V]) CantidadValoresDistintos() int que, sin usar el iterador interno, dado un hash devuelva la cantidad de 
// valores diferentes que almacena. Se puede suponer que el tipo V (el genérico de los valores) en este caso también es comparable, como 
// lo son las claves. Indicar y justificar el orden del algoritmo.

func (dicc *hashCerrado[K, V]) CantidadValoresDistintos() int {
	cantidad := 0
	contador := CrearHash[V,int]()
	for i:= 0; i<dicc.tam; i++ {
		if dicc.tabla[i].estado == OCUPADO && !contador.Pertenece(dicc.tabla[i].dato){
			contador.Guardar(dicc.tabla[i].dato,1)
			cantidad++
		}
	}
    return cantidad
}
//Pasa por todos los espacios del hash cerrado por lo que hasta ese momento es O(n) pero se matiene en O(n) porque las demas operaciones son
// O(1).

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//La diferencia simétrica entre dos conjuntos A y B es un conjunto que contiene todos los elementos que se encuentran en A y no en B, y viceversa.

Implementar una función DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V] que devuelva un nuevo 
Diccionario (puede ser el hash que hayas implementado) con la diferencia simétrica entre los dos recibidos por parámetro. La diferencia 
tiene que ser calculada teniendo en cuenta las claves, y los datos asociados a las claves deben ser los mismos que estaban en cada uno de 
los hashes originales. Ejemplo:

d1 = { "perro": "guau", "gato": "miau", "vaca": "mu" }
d2 = { "perro": "woof", "zorro": "ding-ding" }
DiferenciaSimetricaDict(d1, d2) => { "gato": "miau", "vaca": "mu", "zorro": "ding-ding" }

//Indicar y justificar el orden de la función implementada.

func DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V] {
	diferencias := CrearHash[K,V]()
	d1.Iterar(func(clave K, dato V) bool {
		if !d2.Pertenece(clave) {
			diferencias.Guardar(clave,dato)
		}
		return true
	})
	d2.Iterar(func(clave K, dato V) bool {
		if !d1.Pertenece(clave) {
			diferencias.Guardar(clave,dato)
		}
		return true
	})
    return diferencias
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Una fábrica de pastas de Lanús le pide a alumnos de Algoritmos y Estructuras de Datos que le solucionen un problema: sus dos distribuidoras 
// de materias primas le enviaron un hash cada una, dónde sus claves son los nombres de los productos, y sus valores asociados, sus precios. 
// La fábrica de pastas le pide a los alumnos que le implementen una función que le devuelva un nuevo hash con la unión de todos esos productos, 
// y en caso de que una misma materia prima se encuentre en ambos hashes, elegir la que tenga el precio más barato. Indicar y justificar el orden
// del algoritmo.

func MergeProveedores(prov1, prov2 Diccionario[string, int]) Diccionario[string, int] {
	baratos := CrearHash[string,int]()
	prov1.Iterar(func(clave string, dato int) bool {
		if prov2.Pertenece(clave) {
			if prov2.Obtener(clave) < dato {
				baratos.Guardar(clave,prov2.Obtener(clave))
			} else {
				baratos.Guardar(clave,dato)
			}
		} else {
			baratos.Guardar(clave,dato)
		}
		return true
	})
	prov2.Iterar(func(clave string, dato int) bool {
		if prov1.Pertenece(clave) {
			if prov1.Obtener(clave) < dato {
				baratos.Guardar(clave,prov1.Obtener(clave))
			} else {
				baratos.Guardar(clave,dato)
			}
		} else {
			baratos.Guardar(clave,dato)
		}
		return true
	})
    return baratos
}

//Se recorren dos veces (una vez cada hash) por lo que seria 2 O(n) a lo cual el 2 es depreciable y las demas operaciones son O(1), por lo
// que se matiene en O(n)

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Implementar un algoritmo que reciba un arreglo desordenado de enteros, su largo (n) y un número K y determinar en O(n) si existe un
// par de elementos en el arreglo que sumen exactamente K.

func ParElementosSumanK(arr []int, k int) bool {
    almacenados := CrearHash[int,int]()
    for _,num := range arr {
        if !almacenados.Pertenece(k-num) {
            almacenados.Guardar(num,num)
        } else {
            return true
        }
    }
    return false
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

const ARBOLES string = 

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Se tiene un árbol binario de búsqueda con cadenas como claves y función de comparación strcmp. 
//Implementar una primitiva func (abb *abb[K, V]) Mayores(cadena K) Lista[K] que, dados un ABB y una clave, 
//devuelva una lista ordenada con las claves del árbol estrictamente mayores a la recibida por parámetro 
//(que no necesariamente está en el árbol). Implementar sin utilizar el iterador Interno del ABB.

func (abb *abb[K, V]) Mayores(clave K) Lista[K] {
	listaMayores := CrearListaEnlazada[K]()
    abb.raiz.mayoresRecu(abb.cmp, listaMayores,clave)
	return listaMayores
}

func (nodo *nodoAbb[K,V]) mayoresRecu(funcion_cmp func(K,K) int,lista Lista[K],clave K) {
	if nodo == nil {
		return
	}
	comparacionDesde := funcion_cmp(nodo.clave, clave)

	if comparacionDesde > 0 {
		nodo.izquierdo.mayoresRecu(funcion_cmp, lista, clave)
		lista.InsertarUltimo(nodo.clave)
	
	} 
	
	nodo.derecho.mayoresRecu(funcion_cmp, lista, clave)

}
//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Dado un árbol binario, escribir una primitiva recursiva que determine la altura del mismo. Indicar y justificar el orden de la primitiva.

type ab struct {
    izq *ab
    der *ab
    dato int
}

func maximo(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func (arbol *ab) Altura() int { 
    if arbol == nil {
        return 0
    }
    h_izq := arbol.izq.Altura()
    h_der := arbol.der.Altura()

    return maximo(h_izq, h_der) + 1
}

//Estoy pasando por todos los nodos haciendo cosas te tiempo constante,
//por lo que es O(N)

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Implementar una primitiva que devuelva la suma de todos los datos (números) de un árbol binario.
//Indicar y justificar el orden de la primitiva.

type ab struct {
    izq  *ab
    der  *ab
    dato int
}

func (arbol *ab) Suma() int {
    if arbol == nil {
        return 0
    }
    NodoIzq := arbol.izq.Suma()
    NodoDer := arbol.der.Suma()
    
    return arbol.dato + NodoIzq + NodoDer
}

//es pre order, primero yo, despues mi hijo izq y despues mi hijo der.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Se tiene un AB con números enteros como datos, y se quiere reemplazar cada dato por el resultado de multiplicarlo con los datos de los hijos. 
//Hacer un seguimiento de hacer dicho procesamiento con un preorder, inorder o postorder. A continuación se deja la implementación mediante 
//cada recorrido:

func datoONeutro(ab *arbol[int]) int {
    if ab == nil {
        return 1
    } else {
        return ab.dato
    }
}
    
func MultiplicarConHijosPre(arbol *Arbol[int]) {
    if arbol == nil {
        return  
    } 
    valor_izq := datoONeutro(arbol.izq)
    valor_der := datoONeutro(arbol.der)
    arbol.dato *= valor_izq * valor_der
    MultiplicarConHijosPre(arbol.izq)
    MultiplicarConHijosPre(arbol.der)
}

func MultiplicarConHijosIn(arbol *Arbol[int]) {
    if arbol == nil {
        return  
    } 
    MultiplicarConHijosIn(arbol.izq)
    valor_izq := datoONeutro(arbol.izq)
    valor_der := datoONeutro(arbol.der)
    arbol.dato *= valor_izq * valor_der
    MultiplicarConHijosIn(arbol.der)
}

func MultiplicarConHijosPos(arbol *Arbol[int]) {
    if arbol == nil {
        return
    } 
    MultiplicarConHijosPos(arbol.izq)
    MultiplicarConHijosPos(arbol.der)
    valor_izq := datoONeutro(arbol.izq)
    valor_der := datoONeutro(arbol.der)
    arbol.dato *= valor_izq * valor_der
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Dado un árbol binario, escriba una primitiva recursiva que cuente la cantidad de nodos que tienen exactamente dos hijos directos. 
//¿Qué orden de complejidad tiene la función implementada?

type ab struct {
    izq  *ab
    der  *ab
    dato int
}

func (arbol *ab) DosHijos() int {
    if arbol == nil {
        return 0
    }
    NodoIzq := arbol.izq.DosHijos()
    NodoDer := arbol.der.DosHijos()
    if arbol.izq != nil && arbol.der != nil{
        return NodoIzq + NodoDer + 1
    } else {
        return NodoIzq + NodoDer
    }
}

//La complejidad, al pasar por cada nodo es O(n) pero a su vez realizo por cada nodo operaciones de tiempo constante entonces se mantiene en O(n).

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Escribir una primitiva con la firma func (arbol *Arbol) Invertir() que invierta el árbol binario pasado por parámetro, de manera tal que 
//los hijos izquierdos de cada nodo se conviertan en hijos derechos.

//La estructura Arbol respeta la siguiente definición:

type ab struct {
    izq  *ab
    der  *ab
    dato int
}

func (arbol *ab) Invertir() {
    if arbol == nil {
        return
    }
    arbol.izq.Invertir()
    arbol.der.Invertir()

    arbol.der, arbol.izq = arbol.izq, arbol.der

}

// El orden es Post order, llamo a mi hijo izq, llamo a mi hijo derecho y despues me visito a mi
// La complejidad, al pasar por cada nodo es O(n) pero a su vez tambien realizo operaciones de tiempo constante O(1) por lo que sigue siendo O(n).
//Otra forma seria poner asi:

//arbol.der, arbol.izq = arbol.izq, arbol.der
//arbol.izq.Invertir()
//arbol.der.Invertir()

// tambien funciona y es Pre order

// hacerlo In order :

//arbol.izq.Invertir()
//arbol.der, arbol.izq = arbol.izq, arbol.der
//arbol.der.Invertir()

// No funcionaria porque quedaria ordenado como estaba al principio.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Suponer que se tiene un ABB A con una función de comparación cmp1 con nn claves. También, se tiene 
//otro ABB vacío B con función de comparación cmp2 (con cmp1 y cmp2 diferentes). ¿Es posible insertar en algún orden todas 
//las claves de A en B de tal forma que ambos tengan exactamente la misma estructura? Si es posible, describir un algoritmo 
//que permita lograr esto; si no lo es, razonar por qué. (Considerar que la lógica a emplear debe funcionar para cualquier 
//valor de n y cualquier estructura que tenga el ABB A.)


//RESPUESTA: Si dos árboles tienen misma estructura, en particular tienen mismo inorder. Eso quiere decir que los elementos se 
//comparan de la misma forma para una función y la otra, por lo que, o casualmente justo las funciones comparan a esos elementos 
//de la misma forma a pesar de ser diferentes funciones, o esto no es posible. Como depende de la función y no hay forma de asegurar 
//que pueda funcionar, entonces la respuesta es que no se puede hacer.Como mucho, podrías recorrer inorder el árbol A con la cmp2 y 
//chequear que se cumpla el mismo orden, pero eso es lo mejor que se puede hacer, pero tendría poco sentido a los fines del ejercicio.



//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Definimos como quiebre en un árbol binario cuando ocurre que:

    //un hijo derecho tiene un solo hijo, y es el izquierdo
    //un hijo izquierdo tiene un solo hijo, y es el derecho

//Implementar una primitiva para el árbol binario func (arbol Arbol) Quiebres() int que, dado un árbol binario, nos devuelva la cantidad 
//de quiebres que tiene. La primitiva no debe modificar el árbol. La estructura del tipo Arbol es:

type ab struct {
    izq *ab
    der *ab
    dato int
}

//Indicar y justificar el orden de la primitiva, e indicar el tipo de recorrido implementado.

func (arbol *ab) Quiebres() int {
    if arbol == nil {
        return 0
    }
    NodoIzq := arbol.izq.Quiebres()
    NodoDer := arbol.der.Quiebres()

    suma := NodoIzq + NodoDer

    if arbol.izq != nil && arbol.izq.izq == nil && arbol.izq.der != nil {
        suma += 1
    }
    if arbol.der != nil && arbol.der.der == nil && arbol.der.izq != nil {
        suma += 1
    }
    return suma
}

//Complejidad, recorremos todos los nodos por lo que es O(n) y a su vez estamos realizando operaciones O(1) por lo que sigue siendo O(n).
//Orden, es Post orden porque llamo primero izquierda, despues derecha y dsp yo, osea la suma esta despues de los llamados recursivos.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// Indicar si las siguientes afirmaciones son verdaderas o falsas. En caso de ser verdaderas, justificar, en caso de ser falsas 
//poner un contraejemplo:

 // A) Si dos árboles binarios tienen el mismo recorrido inorder, entonces tienen la misma estructura.

 // B) Si dos árboles binarios tienen el mismo recorrido preorder, entonces tienen la misma estructura.

 // C) Si dos árboles binarios de búsqueda (e idéntica función de comparación) tienen el mismo recorrido preorder, entonces tienen la misma estructura.


 //RESPUESTA: Lo que sucede es que ningún recorrido por sí sólo nos puede permitir reconstruir porque siempre podemos encontrar más de un 
 // árbol que lo cumple.En particular, si tenemos el preorder + inorder, podemos recontruir univocamente. También sucedería con el inorder 
 // y post order, pero no es ni amigable ni directo hacerlo.En el caso del ABB, lo que sucede es que tenemos implicido el inorder, por lo 
 // que con sólo el preorder, ya tendríamos los dos para poder reconstruir.Por ende, es efectivamente F, F, V, pero por eso que digo.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Implementar una primitiva para el ABB, que reciba el ABB y devuelva una lista con las claves del mismo, ordenadas tal que si insertáramos 
// las claves en un ABB vacío (con la misma función de comparación), dicho ABB tendría la misma estructura que el árbol original. ¿Qué tipo 
// de recorrido utilizaste? Indicar y justificar el orden de la primitiva.

func (abb *abb[K, V]) Claves() Lista[K] {
	//el recorrido debe ser preorder:
	reconstruccion := CrearListaEnlazada[K]()
	abb.raiz.clavesPreOrder(reconstruccion)
    return reconstruccion
}

func (abb *nodoAbb[K,V]) clavesPreOrder(lista Lista[K]){
	//preorder: act,izq y der.
	if abb == nil {
		return
	}
	lista.InsertarUltimo(abb.clave)
	abb.izquierdo.clavesPreOrder(lista)
	abb.derecho.clavesPreOrder(lista)

}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Implementar una primitiva para el AB que reciba dos arreglos de claves. El primer arreglo corresponde al preorder de un árbol binario. 
// El segundo al inorder del mismo árbol (ambos arreglos tienen los mismos elementos, sin repetidos). La función debe devolver un árbol 
// binario que tenga dicho preorder e inorder. Indicar y justificar el orden de la primitiva (tener cuidado con este punto). Considerar 
// que la estructura del árbol binario es:

type Arbol struct {
    izq *Arbol
    der *Arbol
    clave int
}

//El preorder te los da ordenados para volver a ponerlos en un arbol

func Reconstruir(preorder, inorder []int) *Arbol {
   if len(preorder) == 0 {
		return nil
	}

	// La primera clave en el preorder es la raíz del árbol
	raiz := &Arbol{clave: preorder[0]}

	// Encontrar la posición de la raíz en el inorder
	var raizIn int
	for i, elem := range inorder {
		if elem == preorder[0] {
			raizIn = i
			break
		}
	}

	// Recursivamente construir los subárboles izquierdo y derecho
	raiz.izq = Reconstruir(preorder[1:raizIn+1], inorder[:raizIn])
	raiz.der = Reconstruir(preorder[raizIn+1:], inorder[raizIn+1:])

	return raiz
}

//La complejidad sera O(n) porque se pasa solo una vez por cada nodo realizando operaciones de tiempo constante.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Implementar una función que reciba un arreglo ordenado y devuelva una lista con los elementos en orden para ser insertados en un ABB, de 
// tal forma que al insertarlos en dicho orden se asegure que el ABB quede balanceado. ¿Cómo cambiarías tu resolución si en vez de querer
// guardarlos en un ABB se fueran a insertar en un AVL?

func balancear[K comparable](claves []K) Lista[K] {
    listaBal := CrearListaEnlazada[K]()
	return balancearRecu(claves, listaBal)
}

func balancearRecu[K comparable](claves []K, lista Lista[K]) Lista[K] {
	if len(claves) == 0 {
		return lista
	}
	raiz := len(claves) / 2
	lista.InsertarUltimo(claves[raiz])

	balancearRecu(claves[:raiz],lista)
	balancearRecu(claves[raiz + 1:], lista)
	return lista
}

//Para un avl lo puedo ordenar de cualquier forma que se reeordena solo con las rotaciones.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Determinar cómo es el Árbol cuyo pre order es EURMAONDVSZT, e in order es MRAUOZSVDNET.

type ab struct {
    izq *ab
    der *ab
    clave string
}

func ReconstruirParticular() *ab {
    preorder := "EURMAONDVSZT"
    inorder := "MRAUOZSVDNET"
    return reconstruirRecu(preorder, inorder)
}

func reconstruirRecu(preorder, inorder string) *ab {
   if len(preorder) == 0 {
		return nil
	}

	// La primera clave en el preorder es la raíz del árbol
	raiz := &ab{clave: string(preorder[0])}

	// Encontrar la posición de la raíz en el inorder
	var raizIn int
	for i:=0 ; i<len(inorder) ; i++ {
		if string(inorder[i]) == string(preorder[0]) {
			raizIn = i
			break
		}
	}

	// Recursivamente construir los subárboles izquierdo y derecho
	raiz.izq = reconstruirRecu(preorder[1:raizIn+1], inorder[:raizIn])
	raiz.der = reconstruirRecu(preorder[raizIn+1:], inorder[raizIn+1:])

	return raiz
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// Implementar una primitiva del ABB que dado un valor entero M, una clave inicial inicio y una clave final fin, se devuelva una lista con 
// todos los datos cuyas claves estén entre inicio y fin, que estén dentro de los primeros M niveles del árbol (considerando a la raíz en nivel 1). 
// Indicar y justificar la complejidad temporal.

// Por ejemplo, si tenemos el siguiente ABB con M = 3, inicio = 5 y fin = 15:

//      10
//    /    \
//   5      15                    Un resultado final serían los datos de las 
//  / \    /  \                   claves 10, 5, 8, 15, 12 (en cualquier orden).
// 3   8  12   20
//    /     \
//   7       14

func (abb *DiccionarioOrdenado[K, V]) Ej17(M, inicio, fin int) Lista[K] {
	lista := CrearListaEnlazada[K]()
	abb.ej17_rec(abb.raiz, M, inicio, fin, 1, lista)
	return lista
}

func (abb *DiccionarioOrdenado[K, V]) ej17_rec(nodo *nodoAbb[K, V], M int, inicio int, fin int, nivel int, lista *Lista[K]) {
	if nodo == nil {
		return
	}
	if nivel > M {
		return
	}
	if abb.cmp(nodo.clave, inicio) >= 0 && abb.cmp(nodo.clave, fin) <= 0 {
		lista.InsertarUltimo(nodo.clave)
	}
	if abb.cmp(nodo.clave, inicio) > 0 {
		abb.ej17_rec(nodo.izquierdo, M, inicio, fin, nivel+1, lista)
	}
	if abb.cmp(nodo.clave, fin) < 0 {
		abb.ej17_rec(nodo.derecho, M, inicio, fin, nivel+1, lista)
	}
}

//complejidad: hay dos casos en el que m no es tan grande, ahi seria O(2^m) y si es muy grande seria O(n) porque paso por todos los elementos.

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Es AVL:

func (arbol *ab[K])EsAVL(funcion_cmp func(K,K) int) (bool, int) {
	//arrancamos comprobando que este balanceado
	if arbol == nil {
		return true, 0
	}

	esAVL_izq, alturaizq := arbol.izq.EsAVL(cmp)
	esAVL_der, alturader := arbol.der.EsAVL(cmp)
	
	altura := mayor(alturaizq, alturader) + 1

	if !esAVL_izq || !esAVL_der || math.Abs(alturaizq-alturader) > 1 {
		return false, altura
	}
	//Aca coroboramos que sea abb, es decir el de izq es mas chico que el padre y el derecho es mas grande que el padre
	if ab.izq != nil && cmp(ab.dato, ab.izq.dato) < 0 {
		return false, altura
	}
	if ab.der != nil && cmp(ab.dato, ab.der.dato) > 0 {
		return false, altura
	}
	return true, altura
}

func mayor(a,b int) int{
	if a >= b {
		return a
	} else {
		return b
	}
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Contar nodos internos de un arbol binario

func (ab *arbol) NodosInternos() int {
	//Si el nodo tiene al menos un hijo, lo cuenta y sino no.
	if ab == nil {
		return 0
	}
	InternosIzq := ab.izq.NodosInternos()
	InternosDer := ab.der.NodosInternos()

	if ab.izq != nil || ab.der != nil {
		return InternosIzq + InternosDer + 1
	} else {
		return InternosIzq + InternosDer
	}
	//otra forma:
	//if ab == nil {
	//	return 0
	//} else if ab.izq == nil && ab.der == nil {
	//	return 0
	//}
	//InternosIzq := ab.izq.NodosInternos()
	//InternosDer := ab.der.NodosInternos()
	//return InternosIzq + InternosDer + 1
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Cantidad de hojas de un arbol binario

func cant_hojas(ab *arbol) int {
	return hojas(ar.raiz)
}
func hojas(nodo *nodoArbol) int {
	if nodo == nil {
		return 0
	}
	if nodo.izq == nil && nodo.der == nil {
		return 1
	}
	return hojas(nodo.izq) + hojas(nodo.der)
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

//Cantidad de hojas de un arbol N-ario (en general)

func Hojas(nodo *nodoArbol) int {
	if nodo == nil {
		return 0
	}
	if nodo.hijos.EstaVacia() {
		return 1
	}
	suma := 0
	iter := nodo.hijos.Iterador()
	for !iter.HaySiguiente() {
		suma += hojas(iter.VerActual())
		iter.Siguiente()
	}
	return suma
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

// ejercicio 17 guiar arboles:
func (abb *DiccionarioOrdenado[K, V]) Ej17(M, inicio, fin int) Lista[K] {
	lista := CrearListaEnlazada[K]()
	abb.ej17_rec(abb.raiz, M, inicio, fin, 1, lista)
	return lista
}

func (abb *DiccionarioOrdenado[K, V]) ej17_rec(nodo *nodoAbb[K, V], M int, inicio int, fin int, nivel int, lista *Lista[K]) {
	if nodo == nil {
		return
	}
	if nivel > M {
		return
	}
	if abb.cmp(nodo.clave, inicio) >= 0 && abb.cmp(nodo.clave, fin) <= 0 {
		lista.InsertarUltimo(nodo.clave)
	}
	if abb.cmp(nodo.clave, inicio) > 0 {
		abb.ej17_rec(nodo.izquierdo, M, inicio, fin, nivel+1, lista)
	}
	if abb.cmp(nodo.clave, fin) < 0 {
		abb.ej17_rec(nodo.derecho, M, inicio, fin, nivel+1, lista)
	}
}

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

const HEAP 

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

func TopK(promedios []int, k int) []int {
	//heap de maximos
	heap := cola_prioridad.CrearHeapArr(promedios, func(a,b int) int { return a-b})

	resultado := make([]int, k)
	for i := range k {          //esto es k log(n)
		resultado[i] = heap.Desencolar()
	}
	return resultado
}

//Complejidad: O(n + k Log(n))


//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================
 
func TopKStream(promedios []int, k int) []int {
	//heap de minimos
	heap := cola_prioridad.CrearHeapArr(promedios[:k], menor(a,b) int)

	for i, promedio := range promedios[k:] {
		if promedio < heap.VerTope() {
			continue
		}
		heap.Encolar(promedio)
		heap.Desencolar()
	}
	resultado := make([]int, k)
	for i := range k {
		resultado[i] = heap.Desencolar()
	}
	return resultado
}

func menor(a,b int) int {
	return b-a
}

//otra forma para tener la funcion de comparacion para minimos:
func(a, b int) int { return cmp(b, a)}

//Esta idea es encolar los k elementos en un heap e ir coparando el elemento mas chico del heap con el elemento siguiente de la lista,
// lo cual esto nos da una complejidad: O(K + N log K) K<<N --> O(N log K)


//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

Se tienen k arreglos de enteros previamente ordenados y se quiere obtener
un arreglo ordenado que contenga a todos los elementos de los k arreglos. Sabiendo que cada arreglo
tiene tamaño h, definimos como n a la sumatoria de la cantidad de elementos de todos los arreglos,
es decir, n = k \times h.

Escribir en Go una función func KMerge(arr [][]int) que reciba los k arreglos y
devuelva uno nuevo con los n elementos ordenados entre sí. La función debe ser de orden
O(n log k). Justificar el orden del algoritmo propuesto.


func KMerge(arr [][]int) []int {
	//este heap va a tener los k arreglos, entonces siempre que desencolemos sera log(k)
    heap := CrearHeapArr[[]int](arr, func (arr1, arr2 []int) int {
        return arr2[0] - arr1[0]
    })

    len_total := len(arr) * len(arr[0])
    merged := make([]int, len_total)
	// este for lo haremos n veces
    for i:=0; i < len_total; i++ {
        primer_arr := heap.Desencolar()
        merged[i] = primer_arr[0]
		// encolar sera log(k)
        if len(primer_arr) > 1 {
            heap.Encolar(primer_arr[1:])
        }
    }
    return merged
}

//complejidad: O(n log(k)) :)

//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================

Dado un arreglo de enteros y un número K, se desea que todos los elementos del arreglo sean mayores a K. Aquellos números que sean menores o 
iguales a K deberían combinarse de la siguiente forma: buscar los dos números más chicos del vector, sacarlos y generar uno nuevo de la forma 
Nuevo número = número-más-chico + 2 x segundo-más-chico. Por ejemplo, si K = 10 y los números más chicos del arreglo son 3 y 4: 3 + 2 * 4 = 11. 
Los números combinados pueden volver a ser combinados con otros en caso de ser necesario (en el ejemplo no lo es), aplicando la misma lógica 
hasta que el número resultante sea mayor a K.

Implementar una función que reciba un arreglo de enteros, su largo y un número K, y devuelva una lista con los elementos que quedarían luego 
de aplicar las modificaciones. El arreglo original debe quedar en el estado original. El orden de la lista no es importante. En caso de no 
poderse combinar todos los elementos para que todos los elementos sea mayores a K, devolver nil. Determinar y justificar la complejidad del 
algoritmo implementado.

Ejemplo: Entrada: [11, 14, 8, 19, 42, 3, 1, 9]; K = 13:

[11, 14, 8, 19, 42, 3, 1, 9] - (1,3)  -> 1 + 2*3 = 7
[11, 14, 8, 19, 42, 9, 7]    - (7,8)  -> 7 + 8*2 = 23
[11, 14, 19, 42, 9, 23]      - (9,11) -> 9 + 11*2 = 31
[14, 19, 42, 23, 31]

Notar que si el 9 no estuviera en el arreglo, no se podría resolver el problema (debemos devolver nil), ya que el 11 no podría combinarse con 
ningún otro número.

func Combinar(arr []int, k int) []int {
    //creo heap de minimos
    heap := CrearHeapArr[int](arr, func(a, b int) int {
        return b - a
    })

    for heap.Cantidad() >= 2 && heap.VerMax() <= k {
        a := heap.Desencolar() //sera el mas chico
        b := heap.Desencolar() // sera el segundo mas chico

        if b > k {
            return nil
        }
        heap.Encolar(a + 2*b)
    }
    if heap.VerMax() <= k {
        return nil
    }
    combinado := make([]int, heap.Cantidad())
    for i:= 0; i < len(combinado); i++{
        combinado[i] = heap.Desencolar()
    }

    return combinado
}

//desencolar y encolar es log(n) 



//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================




//============================================================================================================================================
//--------------------------------------------------------------------------------------------------------------------------------------------
//============================================================================================================================================
