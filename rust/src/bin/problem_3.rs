
fn get_input() -> &'static str{
  return "..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#";
}

fn main(){
  let result = get_input()
  .lines()
  .enumerate()
  .flat_map(|(index,line)| line.chars().nth(index * 3 % line.len()))
  .filter(|&x| x == '#')
  .count();
  print!("treecount: {:?}\n",result)
}
