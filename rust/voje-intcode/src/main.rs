use std::fs;
use std::fmt;
use std::ops::{Index, IndexMut};


mod instructions;

struct Computer {
    mem: Memory,
    instr: Vec<Box<dyn Execute>>,
}

impl Computer {
    fn load_mem(&mut self, mem: String) {
        self.mem = Memory::new(&mem[..]);
        self.parse_mem();
    }

    // Reads memory and generates vector of instructions.   
    fn parse_mem(&mut self) -> Result<(), &str> {
        self.instr = Vec::new();
        let mut ptr: usize = 0;
        while ptr < self.mem.len() {
            ptr = self.parse_executable(ptr)?;
        }
        Ok(())
    }

    // Adds instruction to self.instructions, returns pointer to next 
    // instruction address.
    fn parse_executable(&self, ptr: usize) -> Result<usize, &str> {
        let instr: Execue
        Ok(42)
    }
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

impl Index<usize> for Memory {
    type Output = i32;
    fn index(&self, i: usize) -> &Self::Output {
        &self.fields[i]
    }
}

impl IndexMut<usize> for Memory {
    fn index_mut(&mut self, i: usize) -> &mut Self::Output {
        &mut self.fields[i]
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

    fn len(&self) -> usize {
        self.fields.len()        
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
