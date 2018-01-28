fn main(){
    println!("hello world!");

    let mut v1:Vec<i32>;
    v1 = vec![0;10];

    v1[0] = 12;
    println!("{}",v1[0] );
}