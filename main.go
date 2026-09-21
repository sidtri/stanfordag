package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

type Cell struct {
	X int
	Y int
	C rune
}

func PrintSecret(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	tagRe := regexp.MustCompile(`<[^>]*>`)
	text := tagRe.ReplaceAllString(string(body), "")

	re := regexp.MustCompile(`(\d+)([^\d\s])(\d+)`)
	matches := re.FindAllStringSubmatch(text, -1)

	var cells []Cell
	maxX, maxY := 0, 0

	for _, m := range matches {
		x, _ := strconv.Atoi(m[1])
		ch := []rune(m[2])[0]
		y, _ := strconv.Atoi(m[3])

		cells = append(cells, Cell{X: x, Y: y, C: ch})

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	grid := make([][]rune, maxY+1)
	for y := range grid {
		grid[y] = make([]rune, maxX+1)
		for x := range grid[y] {
			grid[y][x] = ' '
		}
	}

	for _, c := range cells {
		grid[c.Y][c.X] = c.C
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	return nil
}

func main() {
	err := PrintSecret("https://docs.google.com/document/d/e/2PACX-1vSvM5gDlNvt7npYHhp_XfsJvuntUhq184By5xO_pA4b_gCWeXb6dM6ZxwN8rE6S4ghUsCj2VKR21oEP/pub")
	if err != nil {
		fmt.Println(err)
	}
}
