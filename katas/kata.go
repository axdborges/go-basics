package kata

import (
  "strings"
  "strconv"
)

// Retornar o maior das funções matemáticas
func ExpressionMatter(a int, b int, c int) int {
  var total int = 0
  case1 := a * (b + c)
  case2 := a * b * c
  case3 := a + b * c
  case4 := (a + b) * c
  case5 := a + b + c
 
    if total < case1 {
      total = case1
    }
    if total < case2 {
      total = case2
    }
    if total < case3 {
      total = case3
    }
    if total < case4 {
      total = case4 
    }
    if total < case5{
      total = case5
    }
  
  return total
}

// Contagem de pontos de qual time foi vencedor
func Points(games []string) int {
  var total int = 0
  for i := 0; i < len(games); i++ {
    points := strings.Split(games[i], ":")
    we, _ := strconv.Atoi(points[0])
    they, _ := strconv.Atoi(points[1])
    if we > they { 
        total += 3
    } else if  we == they {
        total += 1
    } else {
        total += 0
    }
  }
  return total
}

// Qual o quarter do ano baseado no mês
func QuarterOf(month int) int {
  switch {
    case month <= 3:
      return 1
    case month <= 6:
      return 2
    case month <= 9:
      return 3
    case month <= 12:
      return 4
    default:
      return 0
  }
}

// Trocar todas as letras T de uma sigla DNA para U de RNA
func DNAtoRNA(dna string) string {
  replacedStr := strings.ReplaceAll(dna, "T", "U")
  return replacedStr
}

// Retorna Ímpar ou Par baseado no resto da divisão de 2
func EvenOrOdd(number int) string {
  if number % 2 == 0 {
    return "Even"
  } else {
    return "Odd"
  }
}