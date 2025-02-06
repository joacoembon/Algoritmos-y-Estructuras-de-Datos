#!/usr/bin/python3
import sys
from grafo import Grafo
import biblioteca

FIN_LINEA = "\n"
SEPARADOR_ENTRADA = ' '
COMANDO = 0
PRIMER_PARAMETRO = 1
SEGUNDO_PARAMETRO = 2
ENTRADA_ARCHIVO = 0


def menu_entrada(grafo):
    delincuentes_mas_importantes = []
    lectura_entrada = sys.stdin
    for entrada in lectura_entrada:
        argumentos = entrada.rstrip(FIN_LINEA).split(SEPARADOR_ENTRADA)
        comando = argumentos[COMANDO]
        if not biblioteca.validar_parametros(comando, argumentos[PRIMER_PARAMETRO:]):
            continue
        if comando == "min_seguimientos":
            biblioteca.min_seguimientos(grafo, argumentos[PRIMER_PARAMETRO], argumentos[SEGUNDO_PARAMETRO])

        elif comando == "mas_imp":
            biblioteca.mas_imp(grafo, delincuentes_mas_importantes, argumentos[PRIMER_PARAMETRO])

        elif comando == "persecucion":
            biblioteca.persecucion(grafo, argumentos[PRIMER_PARAMETRO], argumentos[SEGUNDO_PARAMETRO], delincuentes_mas_importantes)

        elif comando == "comunidades":
            biblioteca.comunidades(grafo, argumentos[PRIMER_PARAMETRO])

        elif comando == "divulgar":
            biblioteca.divulgar(grafo, argumentos[PRIMER_PARAMETRO], argumentos[SEGUNDO_PARAMETRO])
            
        elif comando == "divulgar_ciclo":
            biblioteca.divulgar_ciclo(grafo, argumentos[PRIMER_PARAMETRO])

        elif comando == "cfc":
            biblioteca.cfc(grafo)


def main():
    ruta = sys.argv[1:]
    if len(ruta) != 1:
        print("Error al cargar archivo")
        return
    try:
        grafo_delincuentes = biblioteca.cargar_datos(ruta[ENTRADA_ARCHIVO])
        menu_entrada(grafo_delincuentes)
        return
    except FileNotFoundError:
        print("Error al acceder al archivo")
        return
    

main()