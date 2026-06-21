use std::fs::File;
use std::io::{self, Read};

fn read_file_content(path: &str) -> Result<String, io::Error> {
    let mut file = File::open(path)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}

fn process_file(path: &str) {
    match read_file_content(path) {
        Ok(data) => println!("File content: {}", data),
        Err(err) => eprintln!("Error reading file: {}", err),
    }
}

fn main() {
    process_file("data.txt");
}
