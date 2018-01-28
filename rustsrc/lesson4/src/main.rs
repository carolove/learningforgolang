fn main(){
    println!("hello world!" );
    exe(println);
    println!("x, y's sum is  {}", sum(3, 4));
    let sun_f : fn (i32, i32) -> i32 = sum_a;
    println!("{}", sun_f(12,23)); 

}

fn exe(a : fn ()){
    a();
}

fn println(){
    println!("fn in println!" );
}

fn  sum(x: i32, y :i32) -> i32 {
    return y + x;
}

fn sum_a(x: i32, y:i32) -> i32 {
    x + y
}