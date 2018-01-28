use std::sync::Arc;

struct Point {
    x: i32,
    y: i32,
}

fn  main(){
    let x = Arc::new(5);
    let y  = x.clone();
    println!("hello world!");
    println!("x is {}, y is {}",x, y );

    let mut p1 : Point = Point{x:5, y:6};

    p1.x  = 15;
    println!("p1.x is {}", p1.x);

    let p2 : Point = Point{x:5, y:8};

    println!("p2.y is {}", p2.y);

    p2.y = 45;

    println!("p2.y is {}", p2.y);

}