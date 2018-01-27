fn main(){
    println!("hello world!" );
    exe(println);
}

fn exe(a : fn ()){
    a();
}

fn println(){
    println!("fn in println!" );
}