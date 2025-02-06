import random

CONSTANTE_PESO = 1

class Grafo:
    def __init__(self, es_dirigido=False):
        self.es_dirigido = es_dirigido
        self.estructura = {}

    def __len__(self):
        return len(self.estructura)

    def __contains__(self, item):
        return item in self.estructura.keys()

    def __iter__(self):
        return iter(self.obtener_vertices())

    def agregar_vertice(self, vertice):
        '''
        Agrega un vertice al grafo.
        Pre: Recibe por parametro un clave como id para el vertice nuevo.
        '''
        if vertice in self:
            return
        self.estructura[vertice] = {}

    def borrar_vertice(self, vertice):
        '''
        Borra un vertice del grafo
        Pre: El vertice esta en el grafo.
        '''
        if vertice not in self:
            raise Exception("El vertice no se encuentra en el grafo")

        self.estructura.pop(vertice)

        for v in self.estructura:
            if vertice in v:
                v.pop[vertice]

    def agregar_arista(self, a, b, peso=CONSTANTE_PESO):
        '''
        Agregar arista al grafo.
        Pre: Ambos vertices recibidos por parametro pertenecen al grafo.
        Post: Se agrega la arista, y en caso de ser con peso, con peso.
        '''
        if a not in self:
            raise Exception(f"El vertice {a} no se encuentra en el grafo")
        if b not in self:
            raise Exception(f"El vertice {b} no se encuentra en el grafo")

        self.estructura[a][b] = peso
        if not self.es_dirigido:
            self.estructura[b][a] = peso

    def borrar_arista(self, a, b):
        '''
        Borrar arista del grafo.
        Pre: Ambos vertices recibidos por parametro pertenecen al grafo.
        Post: Borra la arista, no los vertices.
        '''
        if not a not in self:
            raise Exception(f"El vertice {a} no se encuentra en el grafo")

        if b not in self.estructura[a]:
            raise Exception(f"La arista no se encuentra en el grafo")

        self.estructura[a].pop(b)

        if not self.es_dirigido:
            self.estructura[b].pop(a)

    def estan_unidos(self, a, b):
        '''
        Devuelve True o False dependiendo si existe dicha arista que los una.
        Pre: Ambos vertices recibidos por parametro pertenecen al grafo.
        '''
        if a not in self:
            raise Exception(f"El vertice {a} no se encuentra en el grafo")

        if b not in self:
            raise Exception(f"El vertice {b} no se encuentra en el grafo")

        return b in self.estructura[a]

    def peso_arista(self, a, b):
        '''
        Devuelve el peso de la arista de los vertices por parametro.
        Pre: Ambos vertices recibidos por parametro pertenecen al grafo.
        '''
        if not self.estan_unidos(a, b):
            raise Exception(f"Los vertices {a} y {b} no estan unidos")
        return self.estructura[a][b]

    def obtener_vertices(self):
        '''
        Devuelve una lista con todos los vertices del grafo.
        '''
        return list(self.estructura.keys())

    def vertice_aleatorio(self):
        '''
        Devuelve un vertice aleatorio del grafo.
        '''
        return random.choice(self.obtener_vertices())

    def adyacentes(self, v):
        '''
        Devuelve una lista con todos los adyacentes en el grafo del vertice pasado por parametro.
        Pre: El vertice recibido por parametro debe estar en el grafo.
        '''
        if v not in self.estructura:
            raise Exception(f"El vertice {v} no se encuentra en el grafo")
        return list(self.estructura[v].keys())