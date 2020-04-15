fn main() {
    let mut s = String::from("hola mundillo");
    idk(&mut s);
    println!("{}",s);
}

fn idk(s: &mut String) {
    s.push_str(" wow");
}