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

        let c = vec[1].trim().matches(vec1[1]).count();
        if c >= vec2[0].parse().unwrap() && c <= vec2[1].parse().unwrap() {
            count += 1
        }
    }

    println!("{}", count);
    Ok(())
}
