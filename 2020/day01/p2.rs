use std::vec::Vec;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
	let mut entries = Vec::new();

	let file = File::open("input.txt")?;
	let reader = BufReader::new(file);

	for line in reader.lines() {
		entries.push(line.unwrap().parse::<i32>().unwrap());	
	}

	let mut stop = false;

	for i in &entries {
		for j in &entries {
			for k in &entries {
				if i + j + k == 2020 {
					println!("{} {} {} {}", i, j, k, i*j*k);
					stop = true;
					break;
				}

				if stop {
					break;
				}
			}
			if stop {
				break;
			}
		}
	}	

	

	Ok(())
}