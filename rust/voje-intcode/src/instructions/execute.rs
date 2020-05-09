trait Instruction {
    fn execute(&self, mem: &mut Vec<i32>) -> Result<usize, &str>;
    fn parse()
}
