use crate::computer::Computer;
use crate::computer::instr::instruction::Instruction;

pub struct Halt {
	len: usize,
}

impl Instruction for Halt{
    fn execute(&self, mut c: Computer) -> Result<usize, &str> {
    	c.halt = true;
    	Ok(self.len())
    }

    fn len(&self) -> usize {
    	self.len
    }
}

impl Halt {
	pub fn new() -> Halt {
		Halt {
			len: 1,
		}
	}
}
