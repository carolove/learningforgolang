use std::sync::Arc;

fn  main(){
    let x = Arc::new(5);
    let y  = x.clone();
    println!("hello world!");
    println!("x is {}, y is {}",x, y )
}