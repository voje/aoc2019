use std::fmt::Debug;
use crate::computer::Computer;

pub trait Instruction: Debug {
	// Executes an instruction on given memory. Returns address of next instruction.   
    fn execute(&self, c: &mut Computer) -> Result<usize, &str>;

    fn len(&self) -> usize;
}
