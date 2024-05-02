package main

import "fmt"

type GridItem struct {
    Shape string
    Level int
}

func getShape(item GridItem) string {
    return item.Shape
}

func getLevel(item GridItem) int {
    return item.Level
}

func main() {
    square := GridItem {
        Shape: "square",
        Level: 100,
    }

    fmt.Printf("Grid item (%s) (%d)", getShape(square), getLevel(square))
}
