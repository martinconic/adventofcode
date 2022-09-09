use std::fs;

fn first(contents: String) {
    let mut floor = 0;
    for c in contents.chars() {
        if c == '(' {
            floor += 1;
        } else {
            floor -= 1;
        }
    }

    println!("{}", floor);
}

fn second(contents: String) {
    let mut floor = 0;
    let mut position = 0;
    for c in contents.chars() {
        if floor == -1 {
            println!("{}", position);
            break;
        }
        if c == '(' {
            floor += 1;
        } else {
            floor -= 1;
        }

        position += 1
    }

    println!("{}", floor);
}

fn main() {
    let contents = fs::read_to_string("./input.txt").expect("Error reading from file");
    //first(contents)
    second(contents)
}
