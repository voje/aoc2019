use num_derive::FromPrimitive;
use num_traits::FromPrimitive;
use std::fs;
use std::fmt;

#[derive(FromPrimitive)]
enum Opcode {
    Add = 1,
    Mul = 2,
    Halt = 3,
}

trait Executable {
    fn execute(&self, mem: &mut Vec<i32>) -> Result<bool, &str>;
}

struct Instruction {
    opcode: Opcode,
}

struct Memory {
    fields: Vec<i32>
}

impl fmt::Display for Memory {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut res = String::new();
        let line_len = 16;
        for (i, el) in self.fields.iter().enumerate() {
            if i % line_len == 0 {
                res += "\n";
            }
            res += &format!("{:>5}", el);
        }
        write!(f, "{}", res)
    }
}

impl Memory {
    fn new(input: &str) -> Memory {
        let mut fields: Vec<i32> = Vec::new();
        for el in input.split(',') {
            let trimmed = el.trim();
            if trimmed.len() == 0 {
                continue
            }
            match trimmed.parse::<i32>() {
                Ok(n) => fields.push(n),
                Err(msg) => println!("Failed parsing {} to i32: {}", el, msg)
            }
        }
        Memory {
            fields: fields,
        }
    }
}

fn main() {
    // Read memory.
    let fdata = fs::read_to_string("./data.txt").expect("Failed to read file!");
    let mem = Memory::new(&fdata[..]);

    println!("{}", mem);

    // Parse memory as instructions.
    //  Throw error on incorrect parse.
}
