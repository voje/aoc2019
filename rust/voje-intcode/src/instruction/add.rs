use crate::instruction::Instruction;

pub struct Add {
	arg1: usize,
	arg2: usize,
	res: usize,
	len: usize,
}

impl Instruction for Add {
    fn execute(&self, mem: &mut Vec<i32>) -> Result<usize, &str> {
    	mem[self.res] = mem[self.arg1] + mem[self.arg2];
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
