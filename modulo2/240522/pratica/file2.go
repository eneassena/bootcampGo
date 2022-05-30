package pratica

import "fmt"

func PlayPraticaCanal() {
  grades := []int{10, 5, 6, 3, 8, 9}

  sumCh := make(chan int)
  avgCh := make(chan int)

  go sum(grades, sumCh)
  go avg(sumCh, avgCh, len(grades))
  printAvg(avgCh)
}

func sum(grades []int, out chan<- int) {
  var total int

  for _, grade := range grades {
    total += grade
  }

  out <- total

  close(out)
}

func avg(in <-chan int, out chan<- int, length int) {
  var avgGrades int

  for v := range in {
    avgGrades = v / length
  }

  out <- avgGrades

  close(out)
}

func printAvg(in <-chan int) {
  for v := range in {
    fmt.Println(v)
  }
}
 