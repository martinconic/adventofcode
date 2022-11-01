use std::fs;
use std::str::Split;

fn first() {}

fn main() {
    let contents = fs::read_to_string("./input.txt").expect("Can not read the file!");
    let carr: Split<&str> = contents.split(", ");
    for c in carr {
        print!("{} ", c);
    }
}
