use crate::computer::instr::instruction::Instruction;
use std::fmt;

mod instr;

pub struct Computer {
    halt: bool,
    pc: usize,
    mem: Vec<i32>,
    instr: Vec<Box<dyn Instruction>>,
}

impl Computer {
    pub fn new(mem: &str) -> Computer {
        let instr: Vec<Box<dyn Instruction>> = Vec::new();
        let mut c = Computer {
            halt: true,
            pc: 0,
            mem: match Computer::read_mem(mem) {
                Ok(vec) => vec,
                Err(err) => panic!("{}", err),
            },
            instr: instr,
        };
        match c.parse_mem() {
            Ok(_) => {},
            Err(e) => println!("Failed parsing memory: {}", e),
        }
        c
    }

    fn read_mem(input: &str) -> Result<Vec<i32>, &str> {
        let mut mem: Vec<i32> = Vec::new();
        for el in input.split(',') {
            let trimmed = el.trim();
            if trimmed.len() == 0 {
                continue
            }
            match trimmed.parse::<i32>() {
                Ok(n) => mem.push(n),
                Err(msg) => println!("Failed parsing {} to i32: {}", el, msg)
            }
        }
        Ok(mem)
    }

    // Reads memory and generates vector of instructions.   
    fn parse_mem(&mut self) -> Result<(), &str> {
        self.instr = Vec::new();
        let mut ptr: usize = 0;
        while ptr < self.mem.len() {
            let binstr: Box<dyn Instruction> = match self.mem[ptr] {
                1 => Box::new(instr::add::Add::new(ptr, &self.mem)),
                2 => Box::new(instr::mul::Mul::new(ptr, &self.mem)),
                99 => Box::new(instr::halt::Halt::new()),
                _ => panic!("Unknown insruction opcode."),
            };
            ptr += (*binstr).len();
            self.instr.push(binstr); 
        }
        Ok(())
    }

    // Adds instruction to self.instructions, returns length of instruction.   
    /*
    fn parse_executable(&self, ptr: usize) -> Result<usize, &str> {
        Ok(42)
    }
    */
}

impl fmt::Display for Computer {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut res = String::new();
        let line_len = 16;
        for (i, el) in self.mem.iter().enumerate() {
            if i % line_len == 0 {
                res += "\n";
            }
            res += &format!("{:>5}", el);
        }
        write!(f, "{}", res)
    }
}
