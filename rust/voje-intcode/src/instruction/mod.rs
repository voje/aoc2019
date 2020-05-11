pub mod add;

pub trait Instruction {
	fn new(ptr: usize, mem: &Vec<i32>) -> Self;

	// Executes an instruction on given memory. Returns address of next instruction.   
    fn execute(&self, mem: &mut Vec<i32>) -> Result<usize, &str>;

    fn len(&self) -> usize;
}
