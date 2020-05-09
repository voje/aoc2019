pub mod add

struct Add {

}

impl Execute for Add {
    fn execute(&self, mem: &mut Vec<i32>) -> Result<usize, &str> {
    	42
    }
}