
pub fn main() {
    //Testing strings
    println!{"Testing code.."};
    let line = "Some line of text for example";
    let text = line.split(" ").nth(3).unwrap();
    println!("{:#?}", text);
    let text = line.split(" ").skip(3).next();
    println!("{:#?}", text);

}
