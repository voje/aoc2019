use std::fs;

fn calc_module_fuel(mass: f64) -> f64 {
	let r = (mass / 3.0).floor() - 2.0;
	if r < 0.0 {
		return 0.0;
	}
	r
}

fn r_calc_module_fuel_helper(mass: f64) -> f64 {
	if mass == 0.0 {
		return 0.0;
	}
	mass + r_calc_module_fuel_helper(calc_module_fuel(mass))
}

// Since the initial mass is needed for the recursive algorithm, we need to
// subtract it from the result.
fn r_calc_module_fuel(mass: f64) -> f64 {
	r_calc_module_fuel_helper(mass) - mass
}

fn main() {
	let file_data = fs::read_to_string("data.txt")
		.expect("Failed to read file.");

	let mut fuel = 0.0;
	let mut r_fuel = 0.0;
	for line in file_data.split("\n") {
		if line.len() == 0 {
			continue;
		}
		let current = line.parse::<f64>().unwrap();
		fuel += calc_module_fuel(current);
		r_fuel += r_calc_module_fuel(current);
	}
	println!("The mass of the fuel is: {}", fuel);
	println!("The mass of the r_fuel is: {}", r_fuel);
}
