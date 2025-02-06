from grafo import Grafo
import funciones

FIN_LINEA = "\n"
SEPARADOR = "\t"
VERTICE_1 = 0
VERTICE_2 = 1

def cargar_datos(ruta):
    grafo = Grafo(es_dirigido=True)
    vertices = set()
    with open(ruta) as archivo:
        for linea in archivo:
            ids = linea.rstrip(FIN_LINEA)
            ids = ids.split(SEPARADOR)
            if not ids[VERTICE_1] in vertices:
                grafo.agregar_vertice(ids[VERTICE_1])
                vertices.add(ids[VERTICE_1])
            if not ids[VERTICE_2] in vertices:
                grafo.agregar_vertice(ids[VERTICE_2])
                vertices.add(ids[VERTICE_2])
            if not grafo.estan_unidos(ids[VERTICE_1], ids[VERTICE_2]):
                grafo.agregar_arista(ids[VERTICE_1], ids[VERTICE_2])
    return grafo


def validar_parametros(comando, parametros):
    if (comando == 'min_seguimientos'):
        if len(parametros) == 2: 
            return True
    elif (comando == 'mas_imp'):
        if len(parametros) == 1 and parametros[0].isdigit(): 
            return True
    elif (comando == 'persecucion'):
        if len(parametros) == 2 and parametros[-1].isdigit(): 
            return True
    elif (comando == 'comunidades'):
        if len(parametros) == 1 and parametros[0].isdigit(): 
            return True
    elif (comando == 'divulgar'):
        if len(parametros) == 2 and parametros[-1].isdigit(): 
            return True
    elif (comando == 'divulgar_ciclo'):
        if len(parametros) == 1 and parametros[-1].isdigit(): 
            return True
    elif (comando == 'cfc'):
        if len(parametros) == 0: 
            return True
    return False
    

def min_seguimientos(grafo, origen, destino):
    resultado = []
    recorrido = funciones.bfs_destinos(grafo, origen, [destino])
    if recorrido == None:
        print("Seguimiento imposible")
        return
    padres, _, _ = recorrido
    resultado = funciones.reconstruir_camino(padres, origen, destino)
    print(" -> ".join(resultado))


def mas_imp(grafo, arreglo, cantidad):
    if len(arreglo) == 0:
        arreglo = funciones.pagerank(grafo)
    delincuentes_este_caso = arreglo[:int(cantidad)]
    resultado = ", ".join(delincuentes_este_caso)
    print(resultado)
    

def persecucion(grafo, delincuentes, kMasImportantes, delincuentes_mas_importantes):
    if len(delincuentes_mas_importantes) == 0:
        delincuentes_mas_importantes = funciones.pagerank(grafo)
    importantes = {}
    for i in range(int(kMasImportantes)): 
        importantes[delincuentes_mas_importantes[i]] = i
    encubiertos = delincuentes.split(",")
    camino_final = {}
    distancia_final = None
    encontrado_actual = None
    inicio = None
    for agente in encubiertos:
        info = funciones.bfs_destinos(grafo,agente,importantes)
        if not info:
            continue
        camino, distancia, encontrado = info
        if not distancia_final or distancia < distancia_final:
            camino_final = camino
            distancia_final = distancia
            encontrado_actual = encontrado
            inicio = agente
        if distancia == distancia_final:
            if importantes[encontrado] < importantes[encontrado_actual]:
                camino_final = camino
                distancia_final = distancia
                encontrado_actual = encontrado
                inicio = agente
    if inicio and encontrado_actual:
        recorrido = funciones.reconstruir_camino(camino_final, inicio, encontrado_actual)
        print(" -> ".join(recorrido))


def comunidades(grafo, integrantes):
    conj_comunidades = funciones.label_propagation(grafo)
    i = 0
    for comunidad in conj_comunidades:
        if len(conj_comunidades[comunidad]) >= int(integrantes):
            integrantes_comunidad = ", ".join(conj_comunidades[comunidad])
            print(f"Comunidad {i}: {integrantes_comunidad}")
            i+=1


def divulgar(grafo, delincuente, saltos):
    divulgacion = funciones.bfs_hasta_n(grafo, delincuente, saltos)
    print(", ".join(divulgacion))


def divulgar_ciclo(grafo, delincuente):
    padres = funciones.bfs_reconstruir_camino_hacia_si_mismo(grafo, delincuente)
    if padres is not None:
        recorrido = funciones.reconstruir_camino(padres, padres[delincuente], delincuente)
        print(delincuente, "->", " -> ".join(recorrido))
        return
    print("No se encontro recorrido")


def cfc(grafo):
    resultado = funciones.cfcs_grafo(grafo)
    for i in range(len(resultado)):
        cfc_final = ", ".join(resultado[i])
        print(f"CFC {i + 1}: {cfc_final}")
