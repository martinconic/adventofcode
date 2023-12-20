use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

fn p1() -> std::io::Result<()> {
    let file = File::open("src/puzzle")?;
    let mut reader = BufReader::new(file);

    let mut line1 = String::new();
    let _ = reader.read_line(&mut line1);
    let vec1: Vec<&str> = line1.split(":").collect();
    println!("{}", vec1[1]);
    let times: Vec<i32> = vec1[1]
        .split_whitespace()
        .map(|s| s.parse::<i32>().unwrap())
        .collect();

    println!("{} - {} - {} - {}", vec1[0], times[0], times[1], times[2]);

    let mut line2 = String::new();
    let _ = reader.read_line(&mut line2);
    let vec2: Vec<&str> = line2.split(":").collect();
    println!("{:?}", vec2);
    let dist: Vec<i32> = vec2[1]
        .split_whitespace()
        .map(|s| s.parse::<i32>().unwrap())
        .collect();

    //println!("{} - {} - {} - {}", vec2[0], dist[0], dist[1], dist[2]);
    let mut totalw = 1;

    for (i, val) in times.iter().enumerate() {
        println!("{} {}", i, val);

        let mut win = 0;

        for k in 1..*val {
            let score = k * (val - k);
            if score > dist[i] {
                win += 1;
            }
            if score < dist[i] && win > 0 {
                break;
            }
        }

        if win > 0 {
            totalw *= win;
        }
    }

    println!("{}", totalw);
    Ok(())
}

fn p2() -> std::io::Result<()> {
    let file = File::open("src/puzzle")?;
    let mut reader = BufReader::new(file);

    let mut line1 = String::new();
    let _ = reader.read_line(&mut line1);
    let vec1: Vec<&str> = line1.split(":").collect();
    let times: i64 = vec1[1]
        .split_whitespace()
        .collect::<String>()
        .parse::<i64>()
        .unwrap();

    let mut line2 = String::new();
    let _ = reader.read_line(&mut line2);
    let vec2: Vec<&str> = line2.split(":").collect();
    let dist: i64 = vec2[1]
        .split_whitespace()
        .collect::<String>()
        .parse::<i64>()
        .unwrap();

    let mut win = 0;

    for k in 1..times {
        let score = k * (times - k);
        if score > dist {
            win += 1;
        }
        if score < dist && win > 0 {
            break;
        }
    }

    println!("{}", win);
    Ok(())
}

fn main() -> std::io::Result<()> {
    p2();
    Ok(())
}
