enum Message {
    Quit,
    ChangeColor(i32, i32, i32),
}

fn main(){
    println!("hello world!" );
    let msg : Message = Message::Quit;
    let msg1 : Message = Message::ChangeColor(12,34,76);
}