from grafo import Grafo
from cola import Cola
from pila import Pila
import heapq, random

def bfs_destinos(grafo, origen, destinos):
    padres = {}
    dist = {}
    visitados = set()
    padres[origen] = None
    dist[origen] = 0
    visitados.add(origen)
    q = Cola()
    q.encolar(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        if v in destinos:
            return padres, dist[v], v
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                dist[w] = dist[v]+1
                visitados.add(w)
                q.encolar(w)
    return None


def bfs_reconstruir_camino_hacia_si_mismo(grafo, origen):
    visitados = set()
    padres = {}
    padres[origen] = None 
    visitados.add(origen)
    q = Cola()
    q.encolar(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                visitados.add(w)
                q.encolar(w)
            if w in visitados and w == origen:
                padres[w] = v
                return padres
    return None


def bfs_hasta_n(grafo, origen, n):
    vertices_a_n = []
    visitados = set()
    padres = {}
    dist = {}
    visitados.add(origen)
    dist[origen] = 0
    q = Cola()
    q.encolar(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                dist_hasta_aca = dist[v] + 1
                if dist_hasta_aca <= int(n):
                    visitados.add(w)
                    padres[w] = v
                    dist[w] = dist_hasta_aca
                    q.encolar(w)
                    vertices_a_n.append(w)
    return vertices_a_n


def pagerank(grafo):
    pagerank_elementos = {}
    d = 0.85
    largo = len(grafo.obtener_vertices())
    constante = (1 - d) / largo
    for v in grafo.obtener_vertices():
        pagerank_elementos[v] = constante
    vertices_ent = vertices_entrada(grafo)
    grados_sal = grados_salida(grafo)
    for i in range(50):
        for v in grafo.obtener_vertices():
            suma = 0
            for w in vertices_ent[v]:
                suma += pagerank_elementos[w] / grados_sal[w]
            pagerank_elementos[v] = constante + suma * d
    vertice_valor = []
    for vertice, valor in pagerank_elementos.items():
        vertice_valor.append((-valor, vertice))
    heapq.heapify(vertice_valor)
    sorted_list = [heapq.heappop(vertice_valor)[1] for _ in range(len(vertice_valor))]

    return sorted_list


def vertices_entrada(grafo):
    vertices = {}
    for v in grafo.obtener_vertices():
        vertices[v] = []
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            vertices[w].append(v)
    return vertices


def grados_salida(grafo):
    d = {}
    for v in grafo.obtener_vertices():
        d[v] = len(grafo.adyacentes(v))
    return d


def cfcs_grafo(grafo):
    resultados = []
    visitados = set()
    for v in grafo:
        if v not in visitados:
            dfs_cfc(grafo, v, visitados, {}, {}, Pila(), set(), resultados, [0])
    return resultados


def dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global):
    orden[v] = mas_bajo[v] = contador_global[0]
    contador_global[0] += 1
    visitados.add(v)
    pila.apilar(v)
    apilados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs_cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])

    if orden[v] == mas_bajo[v]:
        nueva_cfc = []
        while True:
            w = pila.desapilar()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)


def reconstruir_camino(padres, origen, destino):
    recorrido = []
    while destino is not origen:
        recorrido.append(destino)
        destino = padres[destino]
    recorrido.append(origen)
    return recorrido[::-1]


def aristas_de_entrada(grafo):
    entradas = {}
    for v in grafo:
        entradas[v] = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            entradas[w].append(v)
    return entradas


def label_propagation(grafo):
    label = {}
    lista = []
    resultado = {}
    i = 0
    entradas = aristas_de_entrada(grafo)
    for v in grafo:
        label[v] = i
        lista.append(v)
        i += 1
    random.shuffle(lista)
    for _ in range (100):    
        for v in lista:
            label[v] = max_freq(entradas[v], label)
    for v in label:
        if label[v] not in resultado:
            resultado[label[v]] = []
        resultado[label[v]].append(v)
    return resultado


def max_freq(lista, label):
    frecuencias = {}
    for v in lista:
        if label[v] not in frecuencias:
            frecuencias[label[v]] = 0
        frecuencias[label[v]] += 1
    maxima_label = None
    for numero in frecuencias:
        if maxima_label is None:
            maxima_label = numero
        if frecuencias[numero] > frecuencias[maxima_label]:
            maxima_label = numero
    return maxima_label