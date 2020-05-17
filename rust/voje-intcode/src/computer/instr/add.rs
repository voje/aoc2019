use crate::computer::Computer;
use crate::computer::instr::instruction::Instruction;

#[derive(Debug)]
pub struct Add {
	arg1: usize,
	arg2: usize,
	res: usize,
	len: usize,
}

impl Instruction for Add {
    fn execute(&self, c: &mut Computer) -> Result<usize, &str> {
    	c.mem[self.res] = c.mem[self.arg1] + c.mem[self.arg2];
    	Ok(self.len())
    }

    fn len(&self) -> usize {
    	self.len
    }
}

impl Add {
	pub fn new(ptr: usize, mem: &Vec<i32>) -> Add {
		Add{
			arg1: mem[ptr + 1] as usize,
			arg2: mem[ptr + 2] as usize,
			res: mem[ptr + 3] as usize,
			len: 4,
		}
	}
}
