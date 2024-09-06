use std::fs::File;
use std::io::{self, prelude::*, BufReader};


fn p1() -> io::Result<()> {

    let file = File::open("../../input/input.txt")?;
    let reader = BufReader::new(file);
    

    let mut sum = 0;


    for line in reader.lines() {
      let content = line.expect("Failed to read this line");
    
      let mut numbers: Vec<u32> = Vec::new();

      for c in content.chars() {
        if c.is_digit(10) {
          numbers.push(c.to_digit(10).unwrap());
        } 
      }

      sum += 10 * numbers[0] + numbers[numbers.len()-1]; 

    }

    println!("{}", sum);
        

    Ok(())
}


fn main() -> io::Result<()>{
  return p1();
}
