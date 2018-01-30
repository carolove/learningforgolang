enum Message {
    Quit,
    ChangeColor(i32, i32),
    Move{x:i32, y : i32},
    Write(String),
}

fn main(){
    let x = 1;
    match x {
        1 => println!("one"),
        5 => println!("five"),
        _ => println!("not understand!"),
    }
    let msg : Message = Message::Write("this is a string".to_string());
    process_message(msg);
    println!("hello world!");
}

fn process_message (msg :Message) {
    match msg {
        Message::Quit => println!("quit"),
        Message::ChangeColor(x, y) => println!("changeColor x is {}, y is {}" , x, y),
        Message::Move {x, y} => println!("Move x is {}, y is {}" , x, y),
        Message::Write(str) => println!("Write {}", str),
    }
}

