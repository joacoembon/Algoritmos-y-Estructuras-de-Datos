from cola import Cola
from pila import Pila
import heapq
"""
SACADOS DE CLASE:
"""

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
Implementar una función que determine el:

a. El grado de todos los vértices de un grafo no dirigido.

b. El grado de salida de todos los vértices de un grafo dirigido.

c. El grado de entrada de todos los vértices de un grafo dirigido.

Nota: Las funciones deberan devolver un diccionario con clave vertice y valor grado.
"""
def grados(g):
    # devolver un diccionario string -> int
    return grados_salida(g)

def grados_entrada(g):
    # devolver un diccionario string -> int
    d = {}
    for v in g.obtener_vertices():
        d[v]=0
    for v in g.obtener_vertices():
        for w in g.adyacentes(v):
            d[w] = d[w] + 1
    return d

def grados_salida(g):
    # devolver un diccionario string -> int
    d = {}
    for v in g.obtener_vertices():
        d[v] = len(g.adyacentes(v))
    return d

# Por cada vertice paso por E aristas, entonces la complejidad es O(v + e)  si e != v

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
Mismo idem. anterior pero para un vertice en particular
"""

def grados(g,vertice):
    return len(g.adyacentes(vertice))

def grados_entrada(g, vertice):
    # devolver un diccionario string -> int
    vertices = g.obtener_vertices()
    grado = 0
    for v in vertices:
        if vertice in g.adyacentes(v):
            grado +=1
    return grado

def grados_salida(g,vertice):
    return len(g.adyacentes(vertice))

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""Obtener las aristas de un grafo"""

def obtener_aristas(grafo):
    #grafo no dirigido
    aristas = []
    visitados = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v,w))
        visitados.add(v)
    return aristas

def obtener_aristas2(grafo):
    #grafo dirigido
    aristas = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            aristas.append((v,w))
    return aristas

#La complejidad de ambos es O(v + e) porque por cada vertice vemos sus aristas.

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
""""
RECORRIDOS BFS Y DFS
"""
#BFS:
def bfs(grafo, origen):
    visitados = set()
    padres = {}
    orden = {}
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    q = Cola()
    q.encolar(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                q.encolar(w)
    return padres, orden

#DFS:
def dfs(grafo, origen):
    padres = {}
    orden = {}
    visitados = set()
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    _dfs(grafo, origen, visitados, padres, orden)
    return padres, orden

def _dfs(grafo, v, visitados, padres, orden):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            orden[w] = orden[v] + 1
            _dfs(grafo, w, visitados, padres, orden)

#para varias componentes conexas:
def recorrido_dfs_completo(grafo):
    visitados = set()
    padres = {}
    orden = {}
    for v in grafo:
        if v not in visitados:
            visitados.add(v)
            padres[v] = None
            orden[v] = 0
            _dfs(grafo, v, visitados, padres, orden)
    return padres, orden
#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================

"""Implementar un algoritmo que, dado un grafo no dirigido, nos devuelva un ciclo dentro del mismo, si es que los tiene. Indicar 
el orden del algoritmo."""

#dfs
def encontrar_ciclo(g):
    visitados = {}
    padre = {}
    for v in g:
        if v not in visitados:
            ciclo = dfs_ciclo(g, v, visitados, padre)
            if ciclo is not None:
                return ciclo
    return None

def dfs_ciclo(g, v, visitados, padre):
    visitados[v] = True
    for w in g.adyacentes(v):
        if w in visitados:
      # Si w fue visitado y es padre de v, entonces es la arista de donde
      # vengo (no es ciclo).
      # Si no es su padre, esta arista (v, w) cierra un ciclo que empieza
      # en w.
            if w != padre[v]:
                return reconstruir_ciclo(padre, w, v)
        else:
            padre[w] = v
            ciclo = dfs_ciclo(g, w, visitados, padre)
            if ciclo is not None:
                return ciclo

  # Si llegamos hasta acá es porque no encontramos ningún ciclo.
    return None

def reconstruir_ciclo(padre, inicio, fin):
  v = fin
  camino = []
  while v != inicio:
    camino.append(v)
    v = padre[v]
  camino.append(inicio)
  return camino[::-1]

#bfs
def obtener_ciclo_bfs(grafo):
    visitados = set()
    for v in grafo:
        if v not in visitados:
            ciclo = bfs_ciclo(grafo, v, visitados)
            if ciclo is not None:
                return ciclo
    return None

def bfs_ciclo(grafo, v, visitados):
    q = Cola()
    q.encolar(v)
    visitados.add(v)
    padre = {}  # Para poder reconstruir el ciclo
    orden = {}
    padre[v] = None
    orden[v] = 0

    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w in visitados:
        # Si w fue visitado y es padre de v, entonces es la arista
        # de donde vengo (no es ciclo).
        # Si no es su padre, esta arista (v, w) cierra un ciclo que
        # empieza en w.
                if w != padre[v]:
                    return reconstruir_ciclo(padre, orden, w, v)
            else:
                q.encolar(w)
                visitados.add(w)
                padre[w] = v
                orden[w] = orden[v] + 1

  # Si llegamos hasta acá es porque no encontramos ningún ciclo.
    return None
  
def reconstruir_camino(padre, orden, v1, v2):
    ciclo = []
    if orden[v1] != orden[v2]: # no puede haber más que 1 de diferencia
        if orden[v1] > orden[v2]:
            ciclo.append(v1)
            v1 = padre[v1]
        else:
            ciclo.append(v2)
            v2 = padre[v2]
    while v1 != v2:
        ciclo.append(v1)
        ciclo.append(v2)
        v1 = padre[v1]
        v2 = padre[v2]
    ciclo.append(v1)
    return ciclo

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================

"Para un grafo NO dirigido contar componentes conexas"

def contar_componenetes_conexas(grafo):
    contador = 0
    visitados = set()
    for v in grafo:
        if v not in visitados:
            contador += 1
            dfs_comps(grafo, v, visitados)

    return contador

def dfs_comps(grafo, vertice, visitados):
    for w in grafo.adyacentes(vertice):
        if w not in visitados:
            visitados.add(w)
            dfs_comps(grafo, w, visitados)


#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
ORDEN TOPOLOGICO
"""

def topologico_grados(grafo):
    g_ent = grados_entrada(grafo)
    q = Cola()                       #si usara una pila me seguiria dando un orden topologico valido solo q distinto.
    resultado = []
    for v in grafo:
        if g_ent[v] == 0:
            q.encolar(v)
    while not q.esta_vacia():
        v = q.desencolar()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.encolar(w)

    if len(resultado) < len(grafo): # detecta ciclo (no se puede acceder a todos los vertices por los grados de entrada)
        return None
    return resultado

def grados_entrada(g):
    # devolver un diccionario string -> int
    d = {}
    for v in g.obtener_vertices():
        d[v]=0
    for v in g.obtener_vertices():
        for w in g.adyacentes(v):
            d[w] = d[w] + 1
    return d

# CON DFS:

def topologico_dfs(grafo):
    visitados = set()
    pila = Pila()
    for v in grafo:
        if v not in visitados:
            visitados.add(v)
            _dfs(grafo, v, visitados, pila)
    return pila_a_lista(pila)

def _dfs(grafo, v, visitados, pila):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados(w)
            _dfs(grafo, w, visitados, pila)
    pila.apilar(v)

def pila_a_lista(pila):
    lista = []
    while not pila.esta_vacia():
        lista.append(pila.desapilar())
    return lista
#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
CAMINO MINIMO
"""
#Algoritmo de Dijkstra:
def camino_minimo_dijkstra(grafo, origen, destino):
    dist = {}
    padre = {}
    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = heapq()     # de minimos
    q.encolar((0,origen))
    while not q.esta_vacia():
        _, v = q.desencolar()
        if v == destino:
            return padre, dist
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso(v,w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                q.encolar((dist[w],w))
    return padre, dist
#Funciona para grafos dirigidos y no dirigidos pero no con pesos negativos.
#La complejidad es de O(e . log(v)) porque encolo en el heap todas las aristas que en el peor de los casos seria log(v^2) pero por
#propiedad de log el 2 queda despreciable.

#Algoritmo de Bellman-Ford:
def camino_minimo_bf(grafo,origen):
    distancia = {}
    padre = {}
    for v in grafo:
        distancia[v] = float("inf")
    distancia[origen] = 0
    padre[origen] = None
    aristas = obtener_aristas(grafo)
    for i in range(len(grafo)):
        cambio = False
        for origen, destino, peso in aristas:
            if distancia[origen] + peso < distancia[destino]:
                cambio = True
                padre[destino] = origen
                distancia[destino] = distancia[origen] + peso
        if not cambio:
            return padre, distancia
    
    for v, w, peso in aristas:
        if distancia[v] + peso < distancia[w]:
            return None  #Hay ciclo negativo
        
    return padre, distancia

def obtener_aristas(grafo):
    #grafo dirigido
    aristas = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            aristas.append((v,w))
    return aristas
#Complejidad es O(v . e)
# lo uso solamente si el grafo es dirigido y tiene pesos negativos.

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
ARBOLES DE TENDIDO MINIMO MST:
"""
#Algoritmo PRIM:
def mst_prim(grafo):
    v = random.choice(grafo.keys()) #o si el grafo tiene una primitiva de vertice aleatorio
    visitados = set()
    visitados.add(v)
    q = heapq()
    for w in grafo.adyacentes(v):
        q.encolar((v, w), grafo.peso(v, w))
    arbol = Grafo(es_dirigido= True, lista_vertices = grafo.obtener_vertices())
    while not q.esta_vacio():
        (v, w), peso = q.desencolar()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, peso)
        visitados.add(w)
        for x in grafo.adyacentes(w):
            if not x in visitados:
                q.encolar((w, x), grafo.peso(w, x))
    return arbol
#COMPLEJIDAD: O(e. log(v))

#Algoritmo kruskal:
def mst_kruskal(grafo):
    conjuntos = UnionFind(grafo.obtener_vertices())
    aristas = sorted(obtener_aristas(grafo))
    arbol = Grafo(False, grafo.obtener_vertices())
    for a in aristas: # O(e.log(v))
        v, w, peso = a
        if conjuntos.find(v) == conjuntos.find(w):
            continue
        arbol.arista(v, w, peso)
        conjuntos.union(v, w)
    return arbol

class UnionFind:
    def __init__(self, elems):
        #self.groups = {e: e for e in elems}
        self.groups = {}
        for e in elems:
            self.groups[e] = e
    
    def find(self, v):
        if self.groups[v] == v:
            return v

        real_group = self.find(self.groups[v])
        #plancho la estructura
        self.groups[v] = real_group
        return real_group
    
    def union(self, u, v):
        new_group = self.find(u)
        other = self.find(v)
        self.groups[other] = new_group

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
ALGORITMOS DE TARJAN
"""
#PARA PUNTOS DE ARTICULACION (GRAFO NO DIRIGIDO)
def dfs_puntos_articulacion(grafo, v, visitados, padre, orden, mas_bajo, ptos, es_raiz):
    hijos = 0
    mas_bajo[v] = orden[v]
    for w in grafo.adyacentes(v):
        if w not in visitados:
            hijos += 1
            orden[w] = orden[v] + 1
            padre[w] = v
            visitados.add(w)
            dfs_puntos_articulacion(grafo, w, visitados, padre, orden, mas_bajo, ptos, es_raiz = False)
            
            # Lo siguiente se ejecuta una vez ya aplicado a W, i recursivamente a sus hijos
            if mas_bajo[w] >= orden[v] and not es_raiz:
                # No hubo forma de pasar por arriba a este vertice, es punto de articulacion.
                ptos.add(v)
            # Al volver me quedo con que puedo ir tan arriba como mi hijo, si es que me supera.
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])
        elif padre[v] != w: # evitamos considerar a la arista con el padre como una de retorno.
            # Si es uno ya visitado, significa que puedo subir (si es que no podia ya ir mas arriba).
            mas_bajo[v] = min(mas_bajo[v], orden[w])
    if es_raiz and hijos > 1:
        ptos.add(v)
#ejem:
def puntos_articulacion(grafo):
    origen = grafo.random_vertice()
    puntos_articulacion = set()
    dfs_puntos_articulacion(grafo, origen, {origen}, {origen:None}, {origen:0}, {}, puntos_articulacion, True)
    return puntos_articulacion

# Modificacion del algoritmo de Tarjan para detectar CFC's: (Componentes Fuertemente Conexas) (GRAFO DIRIGIDO)
def dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global):
    orden[v] = mas_bajo[v] = contador_global[0]
    contador_global[0] += 1
    visitados.add(v)
    pila.apilar(v)
    apilados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            # Llamamos recursivamente
            dfs_cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])

    if orden[v] == mas_bajo[v]:
        # Se cumple condicion de cierre de un CFC, armamos
        nueva_cfc = []
        while True:
            w = pila.desapilar()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)

def cfcs_grafo(grafo):
    resultados = []
    visitados = set()
    for v in grafo:
        if v not in visitados:
            dfs_cfc(grafo, v, visitados, {}, {}, Pila(), set(), resultados, [0])
    return resultados


#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================


#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
""""
GUIA DE GRAFOS:
"""
#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
Detectar si un grafo es lista enlazada
"""
def es_lista_enlazada(grafo):
    gr_ent = grados_entrada(grafo)
    gr_salida = grados_salida(grafo)
    inicio = fin = None
    for v in grafo:
        if gr_ent[v] > 1:
            return False
        if gr_salida[v] > 1:
            return False
        if gr_ent[v] == 0:
            if inicio is not None:
                return False
            inicio = v
        if gr_salida[v] == 0:
            if fin is not None:
                return False
            fin = v
    if inicio is None or fin is None:
        return False
# Hasta aca esta todo cubierto pero puede ser que tenga asi y ademas tenga otra componente la cual sea un ciclo:
    visitados = set()
    v = inicio
    while True:
        visitados.add(v)
        ady = grafo.adyacentes(v)
        if len(ady) == 0:
            break
        v = ady[0]
    return len(visitados) == len(grafo)

#O(v + e)

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
""""
4)
Implementar un algoritmo que determine si un grafo no dirigido es conexo o no. Indicar la complejidad del algoritmo si el grafo 
está implementado con una matriz de adyacencia.
"""
def es_conexo(grafo):
    contador = 0
    visitados = set()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            contador += 1
            dfs_comps(grafo, v, visitados)
    return contador <= 1 

def dfs_comps(grafo, vertice, visitados):
    for w in grafo.adyacentes(vertice):
        if w not in visitados:
            visitados.add(w)
            dfs_comps(grafo, w, visitados)

# UN GRADO NO DIRIGO ES CONEXO CUANDO SOLAMENTE TIENE UNA COMPONENTE CONEXA ( O ESTA VACIO).

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================




#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
6)
Un árbol es un grafo no dirigido que cumple con las siguientes propiedades:

a. E = V - 1

b. Es acíclico

c. Es conexo

Por teorema, si un grafo cumple dos de estas tres condiciones, será árbol (y por consiguiente, cumplirá la tercera). Haciendo uso de 
ésto (y únicamente de ésto), se pide implementar una función que reciba un grafo no dirigido y determine si se trata de un árbol, o no. 
Indicar el orden de la función implementada.
"""

def es_arbol(grafo):
    
    return len(obtener_aristas(grafo)) == len(grafo) - 1 and es_conexo(grafo)

def obtener_aristas(grafo):
    #grafo no dirigido
    aristas = []
    visitados = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v,w))
        visitados.add(v)
    return aristas

def es_conexo(grafo):
    contador = 0
    visitados = set()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            contador += 1
            dfs_comps(grafo, v, visitados)
    return contador <= 1 

def dfs_comps(grafo, vertice, visitados):
    for w in grafo.adyacentes(vertice):
        if w not in visitados:
            visitados.add(w)
            dfs_comps(grafo, w, visitados)

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
7)
Proponer una función para calcular el grafo traspuesto G^T de un grafo dirigido G. El grafo traspuesto G^T posee los mismos 
vértices que G, pero con todas sus aristas invertidas (por cada arista (v, w) en G, existe una arista (w, v) en G^T). Indicar la 
complejidad para un grafo implementado con:

a. lista de adyancencias

b. matriz de adyacencias"""

def grafo_traspuesto(grafo):
    grafo_nuevo = Grafo(es_dirigido= True)
    for v in grafo.obtener_vertices():
        grafo_nuevo.agregar_vertice(v)
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            grafo_nuevo.agregar_arista(w,v)
    return grafo_nuevo

"""Si fuera lista de adyacencias:"""
# Agregar un vertice: voy a tener q buscarlo --> O(v)
# Agregar una arista: tengo que buscar el vertice y despues preguntar si ya existe --> O(2v)= O(v)
# Obtener todos los adyacentes a un vertice: O(v)

"""Si fuera matriz de adyacencias:"""
# Agregar un vertice: Como minimo O(e)
# Agregar una arista: Como minimo O(v)
# ----> Si tuviera que hacer otra matriz, ambas serian O(v.e) (redimensionar)
# Obtener todos los adyacentes a un vertice: O(v^2)

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
8)
La teoría de los 6 grados de separación dice que cualquiera en la Tierra puede estar conectado a cualquier otra persona del planeta a 
través de una cadena de conocidos que no tiene más de cinco intermediarios (conectando a ambas personas con solo seis enlaces). 
Suponiendo que se tiene un grafo G en el que cada vértice es una persona y cada arista conecta gente que se conoce (el grafo es no 
dirigido):

a. Implementar un algoritmo para comprobar si se cumple tal teoría para todo el conjunto de personas representadas en el grafo G. 
Indicar el orden del algoritmo."""

def seis_grados(grafo):
    visitados = set()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            verdad = bfs(grafo, v, visitados)
            if not verdad:
                return False
    return True

def bfs(grafo, origen, visitados):
    distancia = {}
    distancia[origen] = 0
    visitados.add(origen)
    q = Cola()
    q.encolar(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                distancia[w] = distancia[v] + 1
                visitados.add(w)
                q.encolar(w)
                if distancia[w] > 5:
                    return False
    return True

# COMPLEJIDAD: Utilizo un bfs el cual pasa por cada vertice y cada arista haciendo operaciones O(1) de tiempo constante, por lo que 
#queda con una complejidad de O(v + e).

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================





#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
10)
Escribir una función bool es_bipartito(grafo) que dado un grafo no dirigido devuelva true o false de acuerdo a si es bipartito o no. 
Indicar y justificar el orden del algoritmo. ¿Qué tipo de recorrido utiliza?
"""
#CON BFS:
#seria la funcion generalizada
def es_bipartito(grafo):
    colores= {}
    for vertice in grafo.obtener_vertices():
        if vertice not in colores:
            if not es_bipartito_con_vertice_inicial(grafo, vertice):
                return False
    return True

# esta funcion tambien podria recibir el diccionario colores de la funcion general pero es lo mismo
#arrancando por un vertice
def es_bipartito_con_vertice_inicial(grafo, vertice_inicial):
    colores = {}
    cola = Cola()
    cola.encolar(vertice_inicial)
    colores[vertice_inicial] = 0  # defino colores 0 y 1

    while not cola.esta_vacia():
        v = cola.desencolar()
        for w in grafo.adyacentes(v):
            if w in colores:
                if colores[w] == colores[v]:
                    return False
            else:
                colores[w] = 1 - colores[v]
                cola.encolar(w)
    return True

# CON DFS:
def es_bipartito(grafo):
    colores= {}
    for vertice in grafo.obtener_vertices():
        if vertice not in colores:
            colores[vertice] = 0
            if not es_bipartito_dfs(grafo, vertice, colores):
                return False
    return True

def es_bipartito_dfs(grafo, vertice, colores):
    for w in grafo.adyacentes(vertice):
        if w in colores:
            if colores[w] == colores[vertice]:
                return False
        else:
            colores[w] = 1 - colores[vertice]
            if not es_bipartito_dfs(grafo, w, colores):
                return False
    return True

# LA COMPLEJIDAD DE AMBOS TERMINA SIENDO O(V + E) PORQUE ESTOY USANDO UN BFS Y DFS QUE AMBOS PASAN POR CADA VERTICE POR TODAS SUS
#ARISTAS.

"""
TRUE O FALSE

SI UN GRAFO ES CONEXO ENTONCES ES BIPARTITO --> FALSE
SI UN GRAFO ES BIPARTITO ENTONCES ES CONEXO --> FALSE

SI UN GRAFO TIENE CICLOS DE LONGITUD IMPAR ETONCES NO ES BIPARTITO --> TRUE
SI UN GRAFO NO TIENE CICLOS ENTONCES ES BIPARTITO --> TRUE
SI UN GRAFO TIENE CICLOS Y TODOS SON DE LONGITUD PAR, ENTONCES DEBE SER BIPARTITO --> TRUE
"""

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
11)
Implementar un algoritmo que reciba un grafo dirigido, un vértice V y un número N, y devuelva una lista con todos los vértices que se 
encuentren a exactamente N aristas de distancia del vértice V. Indicar el tipo de recorrido utilizado y el orden del algoritmo. 
Justificar."""

def a_n_aristas(grafo, v, n):
    'Devolver una lista con los vértices que cumplen la propiedad'
    if n < 0:
        return []
    distancia_n = []
    visitados = set()
    visitados.add(v)
    distancia = {v:0}
    cola = Cola()
    cola.encolar(v)
    while not cola.esta_vacia():
        v = cola.desencolar()
        if distancia[v] == n:
            distancia_n.append(v)
        elif distancia[v] < n:
            for w in grafo.adyacentes(v):
                if w not in visitados:
                    distancia[w] = distancia[v] + 1
                    visitados.add(w)
                    cola.encolar(w)
    return distancia_n

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
12)
Implementar una función que permita determinar si un grafo puede ser no dirigido. Determinar el orden del algoritmo implementado.
"""

def puede_ser_no_dirigido(grafo):
    visitados = set()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            visitados.add(v)
            la_verdad = bfs(grafo,v,visitados)
            if not la_verdad:
                return False
    return True

def bfs(grafo, v, visitados):
    q = Cola()
    q.encolar(v)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                q.encolar(w)
            if not grafo.estan_unidos(v,w) or not grafo.estan_unidos(w,v):
                return False
    return True

# COMPLEJIDAD: Utilizo un bfs que trata de ir por cada vertice a cada una de las aristas realizando operacion O(1) de tiempo
#constante por lo que es --> O(v + e).

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================



#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================



#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
15)
Un autor decidió escribir un libro con varias tramas que se puede leer de forma no lineal. Es decir, por ejemplo, después del 
capítulo 1 puede leer el 2 o el 73; pero la historia no tiene sentido si se abordan estos últimos antes que el 1."""

def obtener_orden(grafo):
    grados_ent = grados_entrada(grafo)
    q = Cola()
    resultado = []
    for v in grafo.obtener_vertices():
        if grados_ent[v] == 0:
            q.encolar(v)
    while not q.esta_vacia():
        v = q.desencolar()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            grados_ent[w] -= 1
            if grados_ent[w] == 0:
                q.encolar(w)

    if len(resultado) < len(grafo):
        return "Error" #Ciclo
    return resultado

def grados_entrada(grafo):
    d = {}
    for v in grafo.obtener_vertices():
        d[v] = 0
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            d[w] = d[w] + 1
    return d

#COMPLEJIDAD: 
# - La funcion grados_entrada es O(v + e) por que pasa por cada vertice sus adyacentes.
# - El primer for de la funcion principal es O(v).
# - A partir del while es O(v + e).
# --> O(v + e) + O(v) + O(v + e) ---> O(v + e).

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
16)
Implementar una función que reciba un grafo no dirigido y no pesado implementado con listas de adyacencia (diccionario de diccionarios)
y devuelva una matriz que sea equivalente a la representación de matriz de adyacencia del mismo grafo. Indicar y justificar el orden
del algoritmo implementado."""

def convertir_a_matriz(grafo):
    'Devolver la Matriz construida'
    matriz = []
    vertices = grafo.obtener_vertices()
    for v in vertices:
        l = []
        append_vertice(v, l, vertices, grafo)
        matriz.append(l)
    return matriz, vertices


def append_vertice(v, l, vertices, grafo):
    for w in vertices:
        if grafo.estan_unidos(v, w):
            l.append(grafo.peso_arista(v, w))
        else:
            l.append(0)

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
17)
Implementar una función que reciba un grafo no dirigido, y que compruebe la siguiente afirmación: “La cantidad de vértices de grado 
IMPAR es PAR”. Indicar y justificar el orden del algoritmo si el grafo está implementado como matriz de adyacencia."""

def comprobar_teorema(grafo):
    grados_totales = grados(grafo)
    grados_impares = 0
    for v in grados_totales:
        if grados_totales[v]%2 != 0:
            grados_impares+=1
    return grados_impares%2 == 0


def grados(g):
    d = {}
    for v in g.obtener_vertices():
        d[v] = len(g.adyacentes(v))
    return d
6
#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
18)
Dado un número inicial X se pueden realizar dos tipos de operaciones sobre el número:

    Multiplicar por 2
    Restarle 1.

Implementar un algoritmo que encuentra la menor cantidad de operaciones a realizar para convertir el número X en el número Y, 
con tan solo las operaciones mencionadas arriba (podemos aplicarlas la cantidad de veces que querramos).
"""

#El vertice no existe pero puedo ir creandolo a medida que avanzo

def camino(x, y):
    visitados = set()
    visitados.add(x)
    padres = {x: None}
    q = Cola()
    q.encolar(x)
    while not q.esta_vacia():
        v = q.desencolar()
        if v == y:
            return padres
        if (2 * v) not in visitados:
            visitados.add(2*v)
            padres[2*v] = v
            q.encolar(2*v)
        if (v-1) not in visitados:
            visitados.add(v-1)
            padres[v-1] = v
            q.encolar(v-1)
    return None

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
19)
Se tiene un arreglo de palabras de un lenguaje alienigena. Dicho arreglo se encuentra ordenado para dicho idioma (no conocemos el 
orden de su abecedario). Implementar un algoritmo que reciba dicho arreglo y determine un orden posible para las letras del abecedario 
en dicho idioma. Por ejemplo:

{"caa", "acbd", "acba", "bac", "bad"} --> ['c', 'd', 'a', 'b']
"""
# Guadare las relaciones de las letras con un grafo dirigido, en el cual podre indicar que una letra es mayor a otra, por ejem, A es 
#mayor a B entonces existe la arista A --> B.

# Y si tenemos un grafo y nos piden un orden, ese orden debe ser TOPOLOGICO! lo puedo usar porque el grafo es dirigido, y como se 
#trata de un alfabeto sabemos que es aciclico.

def idioma_alien(palabras):
    grafo = grafo_desde_palabras(palabras)
    grados = {}
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] = grados.get(w,0) + 1
    cola = Cola()
    for v in grafo:
        if v not in grados:
            cola.encolar(v)

    result = []
    while len(cola) > 0:
        v = cola.desencolar()
        result.append(v)
        for ady in grafo.adyacentes(v):
            grados[ady] = grados[ady] - 1
            if grados[ady] == 0:
                cola.encolar(ady)

    return result

def grafo_desde_palabras(palabras):
    grafo = Grafo(es_dirigido = True)
    for i in range(len(palabras)-1):
        p1 = palabras[i]
        p2 = palabras[i+1]

        for letra in p1:
            grafo.agregar_vertice(letra)

        for j in range(len(p1)):
            if p1[j] != p2[j]:
                grafo.agregar_vertice(p2[j])
                grafo.agregar_arista(p1[j], p2[j], 1)
                break
    return grafo

#COMPLEJIDAD O(N + (V+E))

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
20)
Implementar un algoritmo que reciba un grafo dirigido y nos devuelva la cantidad de componentes débilmente
conexas de este. Indicar y justificar la complejidad del algoritmo implementado. """

def cantidad_componentes_debiles(grafo):
    visitados = set() 
    componentes = 0 
    
    for v in grafo:
        if v not in visitados:  
            visitados.add(v)
            componentes += 1  
            dfs(grafo, v, visitados)  
    
    return componentes

def dfs(grafo, v, visitados):
    for w in grafo.adyacentes(v): 
        if w not in visitados:  
            visitados.add(w)
            dfs(grafo, w, visitados)

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
21)
Contamos con un grafo dirigido que modela un ecosistema. En dicho grafo, cada vértice es una especie, y cada arista (v, w) indica que 
v es depredador natural de w. Considerando la horrible tendencia del ser humano por llevar a la extinción especies, algo que nos puede
interesar es saber si existe alguna especie que, si llegara a desaparecer, rompería todo el ecosistema: quienes la depredan no tienen
un sustituto (y, por ende, pueden desaparecer también) y/o quienes eran depredados por esta ya no tienen amenazas, por lo que crecerán 
descontroladamente. Implementar un algoritmo que reciba un grafo de dichas características y devuelva una lista de todas las especies 
que cumplan lo antes mencionado. Indicar y justificar la complejidad del algoritmo implementado."""

def amenazados(grafo):
    'Devolver una lista con los vértices que cumplen la condición'
    grados_ent, grados_sal = grados_entrada(grafo), grados_salida(grafo)
    especies = []
    agregadas = set()
    for v in grafo.obtener_vertices():
        if grados_ent[v] == 0 and v not in agregadas:
            especies.append(v)
            agregadas.add(v)
        for w in grafo.adyacentes(v):
            if grados_sal[v] == 1 and w not in agregadas:
                especies.append(w)
                agregadas.add(w)
            elif grados_ent[w] == 1 and v not in agregadas:
                especies.append(v)
                agregadas.add(v)
    return especies

def grados_entrada(g):
    # devolver un diccionario string -> int
    d = {}
    for v in g.obtener_vertices():
        d[v]=0
    for v in g.obtener_vertices():
        for w in g.adyacentes(v):
            d[w] = d[w] + 1
    return d

def grados_salida(g):
    # devolver un diccionario string -> int
    d = {}
    for v in g.obtener_vertices():
        d[v] = len(g.adyacentes(v))
    return d

#COMPLEJIDAD: O(v + e)

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
23)
El diámetro de una red es el máximo de las distancias mínimas entre todos los vértices de la misma. Implementar un algoritmo que 
permita obtener el diámetro de una red, para el caso de un grafo no dirigido y no pesado. Indicar el orden del algoritmo propuesto."""

def diametro(grafo):
    if len(grafo) == 0:
        return 0
    distancias = camino_minimo_dijkstra(grafo,grafo.vertice_aleatorio())
    distancia_max = 0
    for dist in distancias:
        if distancias[dist] > distancia_max:
            distancia_max = distancias[dist]
    return distancia_max

def camino_minimo_dijkstra(grafo, origen):
    dist = {}
    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
    dist[origen] = 0
    q = []
    heapq.heappush(q, (0, origen))
    while q:
        _, v = heapq.heappop(q)
        for w in grafo.adyacentes(v):
            distancia_act = dist[v] + grafo.peso_arista(v,w)
            if distancia_act < dist[w]:
                dist[w] = distancia_act
                heapq.heappush(q, (dist[w],w))
    return dist

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
31)
Implementar un algoritmo que, dado un grafo dirigido, un vértice s y otro t determine la cantidad mínima de aristas que deberían 
cambiar de sentido en el grafo para que exista un camino de s a t."""
def minimas_inversiones(grafo, s, t):
    copia = Grafo(es_dirigido = True)
    for v in grafo:
        if not copia.existe_vertice(v):
            copia.agregar_vertice(v)
        for w in grafo.adyacentes(v):
            if not copia.existe_vertice(w):
                copia.agregar_vertice(w)
            copia.agregar_arista(v, w, 0)
            if not copia.existe_arista(w, v):
                copia.agregar_arista(w, v, 1)
    padre, dist = camino_minimo_dijkstra(grafo, s, t)
    return dist[t]

def camino_minimo_dijkstra(grafo, origen, destino):
    dist = {}
    padre = {}
    for v in grafo:
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = heapq()
    q.encolar((0,origen))
    while not q.esta_vacia():
        _, v = q.desencolar()
        if v == destino:
            return padre, dist
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso(v,w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                q.encolar((dist[w],w))
    return padre, dist

#ESTE ESTA MEJOR:
def minimas_inversiones(grafo, s, t):
    pesado = Grafo(es_dirigido = True)
    for v in grafo:
        if not v in pesado:
            pesado.agregar_vertice(v)
        for w in grafo.adyacentes(v):
            if not w in pesado:
                pesado.agregar_vertice(w)
                pesado.agregar_arista(v, w, 0)
            if not pesado.estan_unidos(w, v):
                pesado.agregar_arista(w, v, 1)
    return dijkstra(pesado, s, t)

def dijkstra(grafo, origen, destino):
    distancia = {}
    for v in grafo:
        distancia[v] = float("inf")
    distancia[origen] = 0
    q = []
    heapq.heappush(q, (0, origen))
    while q:
        _, v = heapq.heappop(q)
        for w in grafo.adyacentes(v):
            distancia_por_aca = distancia[v] + grafo.peso_arista(v,w)
            if distancia_por_aca < distancia[w]:
                distancia[w] = distancia_por_aca
                heapq.heappush(q, (distancia[w],w))
    return distancia[destino]

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================

"""
TWITTER
Se quiere disenar e implementar un algoritmo de recomendaciones para twitter para un usuario en particular. Se considera que un
usuario V puede estar interesado en seguir a X si:
 - V no sigue a X.
 - V sigue a alguien (W) que sigue a X.
Proponer un algoritmo. Detallas y justificar la complejidad algoritmica
"""
# tenemos un grafo dirigido
# Utilizo un BFS porque me interesa la distancia.

def recomendaciones(grafo, origen, k):
    recomendados = set()
    cola = Cola()
    cola.encolar(origen)
    visitados = set()
    visitados.add(origen)
    distancia = {origen: 0}
    while not cola.esta_vacia():
        v = cola.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                distancia[w] = distancia + 1
                visitados.add(w)
                if distancia[w] == k:
                    recomendados.add(w)
                else:
                    cola.encolar(w)
    return recomendados

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
Aerolineas Buchwald quiere hacerle competencia a Aerolineas Argentinas. Como Martin suele vacacionar en Las Toninas, queria poner la
sede centrar alli. Sin embargo, le pedimos que se encargue de la comida mientras nosotros nos ocupamis de elegir la sede central para
los aeropuertos. Despues de mucho debate, nos queda definir entre dos opciones, cordoba o Mar del plata, decidimos elegir aquella que
sea mas central.
"""
def centralidad(grafo):
     #usaremos la de betweeness que es la frecuencia con la que aparece en cada camino minimo.
    freq = {}
    for v in grafo:
        freq[v] = 0
    
    for v in grafo:
        #dist es un dict con la distancia a cada nodo desde v.
        # padres es un dict con el padre de cada nodo para cada uno de sus caminos.
        dist, padres = camino_minimo_dijkstra(grafo,v) #tambien podria ser bfs(grafo, v)
        for nodo, padre in padres:
            if padre != v:
                freq[padre] += 1
    return freq

def mas_central(grafo):
    freq = centralidad(grafo)
    return ordenar(freq)[0]

#Complejidad depende si usamos el algoritmo de dijkstra si el grafo es pesado o bfs solo si el grafo es NO pesado.
#Complejidad con dijkstra --> O(v^2 (e.logv+v))
#Complejidad con bfs -->O(v^2(e + v))

def camino_minimo_dijkstra(grafo, origen):
    dist = {}
    padre = {}
    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = heapq()     # de minimos
    q.encolar((0,origen))
    while not q.esta_vacia():
        _, v = q.desencolar()
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso(v,w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                q.encolar((dist[w],w))
    return padre, dist

def bfs(grafo, origen):
    visitados = set()
    padres = {}
    orden = {}
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    q = Cola()
    q.encolar(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                q.encolar()
    return padres, orden

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
Implementar un algoritmo que dado un grafo dirigido, un vertice S y otro T determine la minima cantidad de aristas que deberian
cambiar de sentido en el grafo para que exista un camino de S a T.
"""
#La idea va a ser agregar aristas inversas a las orgininales, le ponemos peso a todas las aristas, 0 a las originales y 1 a las
#invertidas.
#Usamos Dijkstra:

def cantidad_minima_inversiones(grafo, s, t):
    copia = Grafo(pesado = True)
    for v in grafo:
        if not copia.existe_vertice(v):
            copia.agregar_vertice(v)
        for w in grafo.adyacentes(v):
            if not copia.existe_vertice(w):
                copia.agregar_vertice(w)
            copia.agregar_arista(v, w, 0)
            if not copia.existe_arista(w, v):
                copia.agregar_arista(w, v, 1)
    dist, padres = camino_minimo_dijkstra(grafo, s, t)
    return dist
#COMPLEJIDAD: O(e logv) + O(v + e) --> O(e logv)

def camino_minimo_dijkstra(grafo, origen, destino):
    dist = {}
    padre = {}
    for v in grafo:
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = heapq()
    q.encolar((0,origen))
    while not q.esta_vacia():
        _, v = q.desencolar()
        if v == destino:
            return padre, dist
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso(v,w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                q.encolar((dist[w],w))
    return padre, dist

#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================
"""
Dado un grafo con valores de peso 1, 2 o 3, dar un algoritmo que, en orden de tiempo lineal, encuentra el camino minimo desde un
origen a todos los demas vertices del grafo.
"""
#Cuando me piden orden lineal en un grafo es O(E + V)
def camino_minimo_lineal(grafo, origen):
    copia = Grafo()
    originales = {}
    for v in grafo:
        if not v in copia:
            copia.agregar_vertice(v)
        for w, peso in grafo.adyacentes(v):
            if not w in copia:
                copia.agregar_vertice(w)
            
            actual = v
            if peso != 1:
                for i in range(peso-1):
                    nuevo = f"{v}{w}{i}"
                    originales[nuevo] = v
                    copia.agregar_vertice(nuevo)
                    copia.agregar_arista(actual, nuevo)
                    actual = nuevo
            copia.agregar_arista(actual, w)
    dist , padres = bfs(copia, origen)
#   limpio los nodos que agregue:
    for nodo in padres:
        if not grafo.contiene_vertice(nodo):
            del(padres[nodo])
        if not grafo.contiene(padres[nodo]): #cambio los padres del diccionario
            padres[nodo] = originales[padres[nodo]]
    return padres
                    



#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================




#=====================================================================================================================================
#|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
#=====================================================================================================================================

"""
1) Se tiene un grafo dirigido G que representa la jerarquía de personal dentro de una empresa (es un organigrama). En este grafo los 
vértices modelan a personas, y una arista (v, w) representa la relación v es jefe directo de w. Implementar una función que nos 
devuelva una secuencia para comunicar cierta noticia a todos los empleados, siendo que una persona no puede enterarse hasta que 
su jefe no se haya enterado. Tener en cuenta que hay personas que tienen varios jefes, y que no pueden enterarse de la noticia antes 
de que todos ellos lo hagan. Indicar y justificar el orden del algoritmo implementado. 
"""
def topologico_grados(grafo):
    g_ent = grados_entrada(grafo)
    q = Cola()                       #si usara una pila me seguiria dando un orden topologico valido solo q distinto.
    resultado = []
    for v in grafo:
        if g_ent[v] == 0:
            q.encolar()
    while not q.esta_vacia():
        v = q.desencolar()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.encolar(w)

    if len(resultado) < len(grafo): # detecta ciclo (no se puede acceder a todos los vertices por los grados de entrada)
        return None
    return resultado

def grados_entrada(g):
    # devolver un diccionario string -> int
    d = {}
    for v in g.obtener_vertices():
        d[v]=0
    for v in g.obtener_vertices():
        for w in g.adyacentes(v):
            d[w] = d[w] + 1
    return d

"""
2) Implementar un algoritmo que, dado un grafo dirigido, nos devuelva un ciclo dentro del mismo, si es que lo tiene. Indicar el orden 
del algoritmo.
"""
#dfs
def encontrar_ciclo(g):
    visitados = {}
    padre = {}
    for v in g:
        if v not in visitados:
            ciclo = dfs_ciclo(g, v, visitados, padre)
            if ciclo is not None:
                return ciclo
    return None

def dfs_ciclo(g, v, visitados, padre):
    visitados[v] = True
    for w in g.adyacentes(v):
        if w in visitados:
      # Si w fue visitado y es padre de v, entonces es la arista de donde
      # vengo (no es ciclo).
      # Si no es su padre, esta arista (v, w) cierra un ciclo que empieza
      # en w.
            if w != padre[v]:
                return reconstruir_ciclo(padre, w, v)
        else:
            padre[w] = v
            ciclo = dfs_ciclo(g, w, visitados, padre)
            if ciclo is not None:
                return ciclo

  # Si llegamos hasta acá es porque no encontramos ningún ciclo.
    return None

def reconstruir_ciclo(padre, inicio, fin):
  v = fin
  camino = []
  while v != inicio:
    camino.append(v)
    v = padre[v]
  camino.append(inicio)
  return camino[::-1]

"""
3) Escribir el pseudocódigo del algoritmo de Dijkstra para obtener los caminos mínimos desde un vértice hacia todos los demás vértices 
del grafo. ¿Cómo podemos modificar el algoritmo de Dijkstra para que en caso de tener dos caminos mínimos se elija al de menor cantidad 
de aristas? Indicar y justificar el orden del algoritmo."""
def camino_minimo_dijkstra(grafo, origen):
    dist = {}
    padre = {}
    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = heapq()
    q.encolar((0,origen))
    while not q.esta_vacia():
        _, v = q.desencolar()
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso(v, w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                q.encolar((dist[w],w))
    return padre, dist

def camino_minimo_modificado_dijkstra(grafo, origen):
    dist = {}
    padre = {}



"""
4) Dado un grafo dirigido, obtener todos los vértices que sean vértices Madre. Un vértice es madre si desde dicho vértice se puede 
llegar a todos los demás.


 Esto es básicamente obtener las CFC y construir un grafo reducido con las CFC, y ver los vértices de grado 
de entrada 0. Si hay más de 1, no hay vértices madre, si hay 1, todos los vértices del original que son de esa CFC, son vértices madre. """
#CFC: componente fuertemente conexa

def vertices_madre(grafo):
    visitados = set()
