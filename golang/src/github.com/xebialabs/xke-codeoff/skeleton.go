package main

import "os"
import "fmt"
import "bufio"
import "strconv"
import "strings"

func solveOne(scanner *bufio.Scanner) string {
  fmt.Printf(".")
  // TODO implement
  return "."
}

func runCases(scanner *bufio.Scanner, writer *bufio.Writer) {
  nrCases := nextInt(scanner)

  for i := 1; i <= nrCases; i++ {
    writer.WriteString(fmt.Sprintf("Case #%d: ", i))
    writer.WriteString(solveOne(scanner))
    writer.WriteString("\n")
  }
}

func main() {
  args := os.Args[1:]

  inFile, _ := os.Open(args[0])
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines) 

  outFile, _ := os.Create(args[1])
  defer outFile.Close()
  writer := bufio.NewWriter(outFile)

  runCases(scanner, writer)
  writer.Flush()
}

func nextString(scanner *bufio.Scanner) string {
  scanner.Scan()
  return scanner.Text()
}

func nextInt(scanner *bufio.Scanner) int {
  i, _ := strconv.Atoi(nextString(scanner))
  return i
}

func nextFloat32(scanner *bufio.Scanner) float32 {
  f, _ := strconv.ParseFloat(nextString(scanner), 32)
  return f
}

func nextStringArray(scanner *bufio.Scanner) []string {
  return strings.Split(nextString(scanner), " ")
}

func nextIntArray(scanner *bufio.Scanner) []int {
  stringArr := nextStringArray(scanner)
  intArr := make([]int, len(stringArr))
  for i, v := range stringArr {
    in, _ := strconv.Atoi(v)
    intArr[i] = in
  }
  return intArr
}

func nextFloat32Array(scanner *bufio.Scanner) []float32 {
  stringArr := nextStringArray(scanner)
  floatArr := make([]float32, len(stringArr))
  for i, v := range stringArr {
    in, _ := strconv.ParseFloat(v, 32)
    floatArr[i] = in
  }
  return floatArr
}

func skipLines(scanner *bufio.Scanner, lines int) {
  for i := 0; i < lines; i++ {
    scanner.Scan()
  }
}

func skipLine(scanner *bufio.Scanner) {
  skipLines(scanner, 1)
}