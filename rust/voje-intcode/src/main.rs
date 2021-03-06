mod computer;
use std::fs;

fn main() {
    // Read memory.
    let fdata = fs::read_to_string("./data.txt").expect("Failed to read file!");
    let c = computer::Computer::new(&fdata[..]);

    println!("{}", c);

    // Parse memory as instructions.
    //  Throw error on incorrect parse.
}
