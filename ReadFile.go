package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "bytes"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]byte, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var blines []byte
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    blines = append(blines, []byte(scanner.Text())...)
  }
  return blines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}

func main() {
  blines, err := readLines("foo.in.txt")
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }

  var lines []string
  i := 0
  for  i < len(blines) {
    var buffer bytes.Buffer
    for j := 0; j < 50 && i < len(blines) ; j++ {
        buffer.WriteByte(blines[i])
        i++
    }
    lines = append(lines, buffer.String())
  }

  for i, line := range lines {
    fmt.Println(i, line)
  }

  if err := writeLines(lines, "foo.out.txt"); err != nil {
    log.Fatalf("writeLines: %s", err)
  }
}
