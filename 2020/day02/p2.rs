use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);

    let mut count = 0;
    for line in reader.lines() {
        let l = line?;
        let split = l.split(":");
        let vec: Vec<&str> = split.collect();

        let vec1: Vec<&str> = vec[0].split_whitespace().collect();
        let vec2: Vec<&str> = vec1[0].split("-").collect();

        let mut c = 0;
        let l1: usize = vec2[0].parse().unwrap();
        if vec[1].trim().chars().nth(l1 - 1).unwrap() == vec1[1].chars().next().unwrap() {
            c += 1
        }

        let v: usize = vec2[1].parse().unwrap();

        if vec[1].trim().chars().nth(v - 1).unwrap() == vec1[1].chars().next().unwrap() {
            c += 1
        }

        if c == 1 {
            count += 1
        }
    }

    println!("{}", count);
    Ok(())
}
