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

        sum += 10 * numbers[0] + numbers[numbers.len() - 1];
    }

    println!("{}", sum);

    Ok(())
}

fn p2() -> io::Result<()> {
    let file = File::open("../../input/test.txt")?;
    let reader = BufReader::new(file);

    let digits = vec![
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    let numbers = vec!["1", "2", "3", "4", "5", "6", "7", "8", "9"];

    let mut sum = 0;

    for line in reader.lines() {
        let mut ifirst = 1000;
        let mut ilast = 0;
        let mut vf = 0;
        let mut vl = 0;

        let content = line.expect("Failed");

        for (i, d) in digits.iter().enumerate() {
            let first = content.find(d).unwrap_or(ifirst);
            let last = content.rfind(d).unwrap_or(ilast);
            print!("{} = {}\n", first, last);

            if first <= ifirst {
                ifirst = first;
                vf = i + 1;
            }

            if last >= ilast {
                ilast = last;
                vl = i + 1;
            }
        }

        for (i, n) in numbers.iter().enumerate() {
            let first = content.find(n).unwrap_or(ifirst);
            let last = content.rfind(n).unwrap_or(ilast);

            if first <= ifirst {
                ifirst = first;
                vf = i + 1;
            }
            if last >= ilast {
                ilast = last;
                vl = i + 1;
            }
        }
        print!("{} - {} \n", vf, vl);
        sum += 10 * vf + vl;
    }

    print!("{}", sum);

    Ok(())
}

fn main() -> io::Result<()> {
    return p1();
}
