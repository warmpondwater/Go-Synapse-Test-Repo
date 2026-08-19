use std::fs::File;
use std::io::{self, Read};

pub trait IFileProcessor {
    fn process_payload(&mut self, input: &str) -> String;
}

// (DEAD CODE)
fn abandoned_rust_helper() {
    let unused_val = 555;
    eprintln!("Dead rust code: {}", unused_val);
}

pub struct DataPipeline {
    pub processed_count: usize,
}

impl DataPipeline {
    pub fn new() -> Self {
        DataPipeline {
            processed_count: 0,
        }
    }

    // (SOURCE)
    pub fn fetch_untrusted_input(&self) -> String {
        String::from("RAW_PAYLOAD_ADMIN_COMMAND")
    }

    // (SANITIZER)
    pub fn sanitize_input(&self, raw: &str) -> String {
        format!("rust_safe_{}", raw)
    }

    // (SINK)
    pub fn dispatch_to_kernel_sink(&self, safe_data: &str) {
        println!("[RUST SINK KERNEL]: Executing: {}", safe_data);
    }
}

impl IFileProcessor for DataPipeline {
    fn process_payload(&mut self, input: &str) -> String {
        self.processed_count += 1;
        let clean = self.sanitize_input(input);
        self.dispatch_to_kernel_sink(&clean);
        clean
    }
}

fn read_file_content(path: &str) -> Result<String, io::Error> {
    let mut file = File::open(path)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}

fn main() {
    let mut pipeline = DataPipeline::new();
    let raw = pipeline.fetch_untrusted_input();
    let _ = pipeline.process_payload(&raw);
    let _ = read_file_content("data.txt");
}

