use crate::computer::Computer;
use crate::computer::instr::instruction::Instruction;

pub struct Mul {
	arg1: usize,
	arg2: usize,
	res: usize,
	len: usize,
}

impl Instruction for Mul {
    fn execute(&self, mut c: Computer) -> Result<usize, &str> {
    	c.mem[self.res] = c.mem[self.arg1] * c.mem[self.arg2];
    	Ok(self.len())
    }

    fn len(&self) -> usize {
    	self.len
    }
}

impl Mul {
	pub fn new(ptr: usize, mem: &Vec<i32>) -> Mul {
		Mul {
			arg1: mem[ptr + 1] as usize,
			arg2: mem[ptr + 2] as usize,
			res: mem[ptr + 3] as usize,
			len: 4,
		}
	}
}