package helpers

import (
	"bufio"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
)

func ReadFileLines(path string) ([]string, error) {
	log.Printf("Lendo arquivo: %s", path)
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Erro ao ler arquivo: %s", err)
		return nil, err
	}

	defer file.Close()
	r := transform.NewReader(file, charmap.Windows1252.NewDecoder())

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	if scanner.Err() != nil {
		log.Printf("Erro ao ler arquivo: %s", scanner.Err())
		return nil, scanner.Err()
	}

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ListFileFromDiretory(diretoryPath string) []string {
	files, _ := ioutil.ReadDir(diretoryPath)
	var filesPath []string
	for _, f := range files {
		filesPath = append(filesPath, f.Name())
	}
	return filesPath
}
