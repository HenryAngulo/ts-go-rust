package main

import (
	"log"
	"strconv"
	"strings"
)

func getInput() string{
  return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

type Point struct {
  x int
  y int
}

type Line struct {
  p1 *Point
  p2 *Point
}

func isHorizontalOrVertical(line *Line) bool{
  return line.p1.x == line.p2.x || line.p1.y == line.p2.y
}

func parsePoint(point string) (*Point,error){
	parsedPoint	:= strings.Split(point,",")
	x,xerr := strconv.Atoi(parsedPoint[0])
	if xerr != nil{
		return nil, xerr
	}
	y,yerr := strconv.Atoi(parsedPoint[1])
	if yerr != nil{
		return nil, yerr
	}
	return &Point{x:x,y:y},nil
}

 func parseLine(line string) (*Line,error){
	parsedLine	:= strings.Split(line," -> ")
	pointOne, pointOneErr := parsePoint(parsedLine[0])
	if pointOneErr != nil {
		return nil, pointOneErr
	}
	pointTwo, pointTwoErr := parsePoint(parsedLine[1])
	if pointTwoErr != nil {
		return nil, pointTwoErr
	}
  return &Line{p1:pointOne, p2:pointTwo}, nil
}


func main(){
	lines := []Line{}
	for _, l := range strings.Split(getInput(), "\n") {
		line,err := parseLine(l)
		if err !=nil{
			log.Fatal("There was an error parsing the line")
		}
		if(isHorizontalOrVertical(line)){
			lines = append(lines, *line)
		}
	}

	log.Printf("lines: %+v",lines)
}