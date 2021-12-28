use std::collections::HashMap;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
	let mut entries = HashMap::new();

	let file = File::open("input.txt")?;
	let reader = BufReader::new(file);

	for line in reader.lines() {
		entries.insert(line.unwrap().parse::<i32>().unwrap(), 1);	
	}

	for (key, value) in &entries {
		let skey = 2020 - key; 
		if entries.contains_key(&skey) {
			println!("{}", key*skey);
			break;
		}
	}

	Ok(())
}