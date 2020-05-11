use crate::instruction::Instruction;

pub struct Add {

}

impl Instruction for Add {
	fn new(_ptr: usize, _mem: &Vec<i32>) -> Add {
		Add{}
	}

    fn execute(&self, _mem: &mut Vec<i32>) -> Result<usize, &str> {
    	Ok(42)
    }

    fn len(&self) -> usize {
    	42
    }
}

