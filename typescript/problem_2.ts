function getInput(): string{
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

type Point = {
  x: number,
  y: number,
}

type Line = {
  p1:Point,
  p2: Point
}

function isHorizontalOrVertical(line: Line){
  return line.p1.x === line.p2.x || line.p1.y === line.p2.y
}

function parsePoint(point:string):Point{
  const [x,y] = point.trim().split(',')
  return {
    x: parseInt(x, 10),
    y: parseInt(y, 10)
  }
}

function parseLine(line:string):Line{
  const [pointOne, pointTwo] = line.trim().split(" -> ")
  return {
    p1:parsePoint(pointOne),
    p2:parsePoint(pointTwo)
  }
}

const out = getInput()
  .split('\n')
  .map(parseLine)
  .filter(isHorizontalOrVertical)

  console.log(out)