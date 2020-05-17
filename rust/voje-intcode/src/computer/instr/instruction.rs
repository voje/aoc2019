use crate::computer::Computer;

pub trait Instruction {
	// Executes an instruction on given memory. Returns address of next instruction.   
    fn execute(&self, c: Computer) -> Result<usize, &str>;

    fn len(&self) -> usize;
}
