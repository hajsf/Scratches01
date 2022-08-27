// #![windows_subsystem = "windows"]
#![feature(proc_macro_hygiene, decl_macro)]

use rocket::*;
use rocket_contrib::*;
use serde_derive::*;

use std::str::FromStr;
use std::process::Command;

use web_view::*;
use std::thread;
use opencv::{core, highgui, prelude::*, videoio, core::Size2i, imgcodecs};
use opencv::videoio::VideoCapture;

use rocket_contrib::serve::StaticFiles;
use rocket_contrib::json::{Json, JsonValue};
use rust_embed::RustEmbed;
use std::sync::Mutex;
use std::collections::HashMap;
use rocket::State;

use regex::Regex;
use std::io::{Read, Cursor};
use rocket::response::Responder;
use opencv::types::VectorOfu8;

#[derive(RustEmbed)]
#[folder = "$CARGO_MANIFEST_DIR/static"]
struct Asset;

// The type to represent the ID of a message.
type ID = usize;

// We're going to store all of the messages here. No need for a DB.
type MessageMap = Mutex<HashMap<ID, String>>;

#[derive(Clone, Debug)]
struct MyRegex(Regex);

impl Eq for MyRegex {}

impl PartialEq for MyRegex {
    fn eq(&self, other: &MyRegex) -> bool {
        self.0.as_str() == other.0.as_str()
    }
}


impl std::hash::Hash for MyRegex {
    fn hash<H: std::hash::Hasher>(&self, state: &mut H) {
        self.0.as_str().hash(state);
    }
}

#[derive(Debug)]
#[derive(Serialize, Deserialize)]
struct Message {
    id: Option<ID>,
    contents: String
}

// TODO: This example can be improved by using `route` with multiple HTTP verbs.
#[post("/<id>", format = "json", data = "<message>")]
fn new(id: ID, message: Json<Message>, map: State<MessageMap>) -> JsonValue {
    let mut hashmap = map.lock().expect("map lock.");
    println!("msg received: {:#?}", message);

    if hashmap.contains_key(&id) {
        println!("error");
        json!({
            "status": "error",
            "reason": "ID exists. Try put."
        })
    } else {
        println!("ok");
        hashmap.insert(id, message.0.contents);
        json!({ "status": "ok" })
    }
}

#[derive(Debug)]
#[derive(Serialize, Deserialize)]
struct Function<'a> {
    command: &'a str,
}

#[post("/", format = "json", data = "<function>")]
fn invoke(function: Json<Function>) -> JsonValue {
    let command = function.command;
/*
     'Basic Mathematical Opertions': {
               // ?\s? can be used instead of space, also could use /i instead of $/,
                'regexp': /^(What is|What's|Calculate|How much is) ([\w.]+) (\+|and|plus|\-|less|minus|\*|\x|by|multiplied by|\/|over|divided by) ([\w.]+)$/,
                'callback': math,
              },
              var math = function(){
               //   'callback': function(par1,a,operation,b){  /// below is better
                    var operation = RegExp.$3;
                    var a = parseFloat(RegExp.$2);
                    var b = parseFloat(RegExp.$4);
                    switch(operation){
                      case '+':
                      case 'and':
                      case 'plus':
                            speak('The sum of: '+a+' and '+b+' is: '+(a+b));
                            alert('The sum of: '+a+' and '+b+' is: '+(a+b));
                     break;
                    }
*/ 
  //  let re = Regex::new(r"(?i)(hello|Hi) (\d) (How are you)").unwrap();
    println!("command {}", command);
    let re = Regex::new(r#"(?xi)
                (What\sis|What's|Calculate|How much is)
                \s(\d)
                \s(\+|and|plus|\-|less|minus|x|by|multiplied\sby|/|over|divided\sby)
                \s(\d)
            "#).unwrap();
    let matching = re.is_match(command);
    
    println!("re:{:#?}\n Command: {}\n matching: {}", re, command, matching);


    //re

    for caps in re.captures(command) {
        println!("groups: {} {} {} {}", 
            caps.get(1).unwrap().as_str(),
            caps.get(2).unwrap().as_str(),
            caps.get(3).unwrap().as_str(),
            caps.get(4).unwrap().as_str()
        );
    }

    json!({ "status": "ok" })
}

fn sum(x: i32, y: i32) -> i32 {
    x + y
}

fn main() {
 /*   let mut commands: HashMap<MyRegex, fn(x: i32, y: i32) -> i32> = HashMap::new();
    commands.insert(
        MyRegex(Regex::new(r#"(?xi)
                (What\sis|What's|Calculate|How much is)
                \s(\d)
                \s(\+|and|plus)
                \s(\d)
            "#).unwrap()), 
        sum
    ); */
    let mut commands: Vec<(Regex, fn(x: i32, y: i32) -> i32)> = Vec::new();
    commands.push(
        (Regex::new(r#"(?xi)(What\ is|What's|Calculate|How\ much\ is)\W*(\d+)\W*(\+|and|plus)\W*(\d+)"#).unwrap(), 
        sum)
    );

    let phrase = "what is 9 + 21";

    for (command, function) in &commands {
        match command.is_match(phrase) {
            true => {
                for caps in command.captures(phrase) {
                    let a = i32::from_str(caps.get(2).unwrap().as_str()).unwrap();
                    let b = i32::from_str(caps.get(4).unwrap().as_str()).unwrap();
                    let sum_num = function(a, b);
                    println!("Sum of {} and {} is {}", a, b, sum_num);
                }
            },
            false => {},
        }   
    }

    thread::spawn(move || {
            Command::new("chrome")
            .args(&["--chrome-frame", "--app=http://localhost:8000/",  "--fullscreen", // --kiosk
                "--window-size=2000,1200"])
            .output()
            .expect("failed to execute process");
    });

 //   thread::spawn(move || {
        rocket::ignite()
            .mount("/",
                   StaticFiles::from(concat!(env!("CARGO_MANIFEST_DIR"), "/static")))
            .mount("/message", routes![new])
            .mount("/invoke", routes![invoke])
            .manage(Mutex::new(HashMap::<ID, String>::new()))
            .launch();

   // openCV().unwrap();
 //   });

    let webview:WebView<()> = web_view::builder()
         .title("My Project")
         .content(Content::Url("http://localhost:8000/"))
         .size(1200, 800)
         .resizable(false)
         .debug(true)
         .user_data(())
         .invoke_handler(invoke_handler)
         .build()
         .unwrap();
    webview.run().unwrap();
}

fn invoke_handler(wv: &mut WebView<()>, arg: &str) -> WVResult {
        match arg {
            "Run openCV" => openCV().unwrap(),
            "Stop openCV" =>     wv.eval(&format!(r#"
                            alert("not enabled now");
                        "#)).unwrap(),
            //videoio::VideoCapture::release(unsafe { &mut cam }).expect("Can not release Cam"),
            _ => (),
        }
    Ok(())
}

fn refresh(wv: &mut WebView<()>){
    wv.eval(&format!(r#"
        var dt = new Date();
        img = document.getElementById('img');
        img.src="/static/savedImage.jpg"+ "?" + dt.getTime();
    "#)).unwrap();
}

// static mut cam : VideoCapture = videoio::VideoCapture::new(0, videoio::CAP_ANY).unwrap();  // 0 is the default camera

fn openCV() -> opencv::Result<()> {
    let window = "video capture";
    highgui::named_window(window, 1)?;
    // https://docs.rs/opencv/0.34.0/opencv/highgui/enum.WindowPropertyFlags.html
    // https://docs.rs/opencv/0.34.0/opencv/highgui/enum.WindowFlags.html
    highgui::set_window_property(window, highgui::WND_PROP_FULLSCREEN, highgui::WINDOW_NORMAL as f64)?;
    #[cfg(feature = "opencv-32")]
    let mut cam = videoio::VideoCapture::new_default(0)?;  // 0 is the default camera
    #[cfg(not(feature = "opencv-32"))]
    let mut cam = videoio::VideoCapture::new(0, videoio::CAP_ANY)?;  // 0 is the default camera

    let frame_width =  match cam.get(3){
        Ok(x) => x as i32,
        Err(_) => 0
    };
    let frame_height = match cam.get(4){
        Ok(x) => x as i32,
        Err(_) => 0
    };

    let opened = videoio::VideoCapture::is_opened(&cam)?;
    let fourcc = videoio::VideoWriter::fourcc(
        'h' as i8,
        '2' as i8,
        '5' as i8,
        '6' as i8
    )?;
    let mut out = videoio::VideoWriter::new(
        "output1.mp4",
        fourcc,
        30.0,
        Size2i::new(frame_width, frame_height),
        true
    ).expect("Can not open video writer");

    if !opened {
        panic!("Unable to open default camera!");
    }
    loop {
        let mut frame = core::Mat::default()?;
        cam.read(&mut frame)?;
        if frame.size()?.width > 0 {
            highgui::imshow(window, &mut frame)?;
            videoio::VideoWriter::write(&mut out, &frame).expect("Error incured");
        }
        let key = highgui::wait_key(10)?;
        if key > 0 && key != 255 {
            break;
        }
    }
    videoio::VideoCapture::release(&mut cam).expect("Can not release Cam");
    videoio::VideoWriter::release(&mut out).expect("Can not release video writer");
    highgui::destroy_all_windows();
    Ok(())
}


type VideoFrame = VectorOfu8;
struct StreamingVideo {
    source: Box<dyn Iterator<Item=VideoFrame>>,
    cursor: Option<Cursor<Vec<u8>>>,
}

impl Read for StreamingVideo {
    fn read(&mut self, buf: &mut [u8]) -> Result<usize, std::io::Error> {
        if self.cursor.is_none() {
            // Get the next frame
            let frame = match self.source.next() {
                Some(f) => f,
                None => return Ok(0),
            };

            // Add the header
            let header = b"--frame\r\nContent-Type: image/jpeg\r\n\r\n";
            let mut bytes = header.to_vec();
            bytes.extend(frame.to_vec());
            self.cursor = Some(Cursor::new(bytes));
        }

        // Read the current frame into the buffer
        let cursor = self.cursor.as_mut().unwrap();
        let n = cursor.read(buf);
        // If we completed the frame, reset `cursor` to `None` so that
        // the next frame will be taken the next time `read` is called
        if cursor.position() == cursor.get_ref().len() as u64 {
            self.cursor = None;
        }
        n
    }
}

impl<'r> Responder<'r> for StreamingVideo {
    fn respond_to(self, _: &Request) -> response::Result<'r> {
        Response::build()
            .streamed_body(self)
            .raw_header("Content-Type", "multipart/x-mixed-replace; boundary=frame")
            .ok()
    }
}