package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  entrada := os.Args[1:]
  numeros := make([]int, len(entrada)) // Cria um slice que irá conter apenas números inteiros e tem o tamanho da entrada

  for i, n := range entrada {
    numero, err := strconv.Atoi(n)

    if err != nil {
      fmt.Printf("%s não é um número válido\n", n)
      os.Exit(1)
    }
    numeros[i] = numero

    fmt.Println(quicksort(numeros))
  }
}

func quicksort(numeros []int) []int{
  if len(numeros) <= 1 {
    return numeros
  }
  n := make([]int, len(numeros))
  copy(n, numeros)

  indicePivo := len(n)/2
  pivo := n[indicePivo]
  n = append(n[:indicePivo], n[indicePivo+1:]...) // Pega o slice até o pivo (sem incluir ele) e depois pega tds dps do pivo e adiciona

  menores, maiores := particionar(n, pivo)

  return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int){
  for _, n := range numeros {
    if n <= pivo {
      menores = append(menores, n)
    } else {
      maiores = append(maiores, n)
    }
  }

  return menores, maiores
  
}