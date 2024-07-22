function getInput():string{
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

const treeChar = '#'

const things = getInput()
  .split("\n")

const collen = things[0].length
let treeCount = 0

things.forEach((thing,i) =>{
  if(thing[(i * 3) % collen] === treeChar){
    treeCount++
  }
})

console.log('treeCount', treeCount)