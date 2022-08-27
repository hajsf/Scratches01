use std::fs;
use std::io::prelude::*;
use std::net::TcpListener;
use std::net::TcpStream;

fn main() {
    let listener = TcpListener::bind("127.0.0.1:7878").unwrap();

    for stream in listener.incoming() {
        let stream = stream.unwrap();
        handle_connection(stream);
    }
}

fn handle_connection(mut stream: TcpStream) {
    let mut buffer = [0; 512];

    let contents = fs::read_to_string("hello.html").unwrap();

    match stream.read(&mut buffer) {
        Ok(s) => {
            println!("{}", contents);
            let response = format!("HTTP/1.1 200 OK\r\n\r\n{}", contents);
            stream.write(response.as_bytes()).unwrap(); // Not OK
            stream.flush().unwrap();
        },
        Err(e) => {
            println!("error: {}", e);
            let response = format!("HTTP/1.1 200 OK\r\n\r\n{}", e);
            stream.write(response.as_bytes()).unwrap(); // Not OK
            stream.flush().unwrap();
        }
    }


  //  println!("{}", contents); // OK

  //  let response = format!("HTTP/1.1 200 OK\r\n\r\n{}", contents);



}
