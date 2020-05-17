use crate::computer::Computer;
use crate::computer::instr::instruction::Instruction;

#[derive(Debug)]
pub struct Halt {
	len: usize,
}

impl Instruction for Halt{
    fn execute(&self, c: &mut Computer) -> Result<usize, &str> {
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
