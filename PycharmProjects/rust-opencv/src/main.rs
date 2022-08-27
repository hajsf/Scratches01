#![feature(proc_macro_hygiene, decl_macro)]
use rocket::*;
use rocket_contrib::*;
use rocket_contrib::serve::StaticFiles;

use opencv::{
    core,
    highgui,
    prelude::*,
    videoio,
    core::Size2i,
};
use std::thread;
use std::sync::mpsc::channel;
use web_view::*;
use opencv::imgcodecs::{imencode, imwrite};
use std::io::{Cursor, Read};
use rocket::response::{Responder, Stream};
use rocket::{Request, Response, response};
use std::fs::File;
use opencv::types::{VectorOfi32, VectorOfu8};
use opencv::imgcodecs::ImwriteFlags::IMWRITE_JPEG_QUALITY;

fn run() -> opencv::Result<()> {
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
            ////
            frame.resize(400);
          //  let mut encodedImage = opencv::core::Vec2i::default(); //    ::Vector::<u8>::new();

            let mut encodedImage = VectorOfu8::new();
            let mut params = VectorOfi32::new();
         //   params.push(1);
            match opencv::imgcodecs::imencode(".JPG", &frame,
                                              &mut encodedImage, &params).unwrap()  {
                true => { println!("done");

                    let filename = "static/savedImageX.jpg"; //' # set filename
                    imwrite(&filename, &frame,  &params);

                    // Generator

                },
                false => println!("not done")
            }

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

#[get("/v")]
fn m() -> Stream<File> {
    let file = File::open("static/savedImageX.jpg").unwrap();
    let response = Stream::chunked(file, 10);
    response
}
/*
#[get("/v")]
fn stream() -> StreamingVideo {
    StreamingVideo {
        source: Box::new(()),
        cursor: None
    }
}
*/
#[get("/")]
fn index() -> &'static str {
    let (tx, rx) = channel();
    // https://doc.rust-lang.org/std/thread/fn.spawn.html
    // http://squidarth.com/rc/rust/2018/06/04/rust-concurrency.html
    let sender = thread::spawn(move || {
        run().unwrap();
        tx.send("Hello, thread".to_owned())
            .expect("Unable to send on channel");
    });

    let receiver = thread::spawn(move || {
        let value = rx.recv().expect("Unable to receive from channel");
        println!("{}", value);
    });

    sender.join().expect("The sender thread has panicked");
    receiver.join().expect("The receiver thread has panicked");

    "Hello, world!"
}

fn main() {

    /*    web_view::builder()
        .title("Graceful Exit Example")
        .content(Content::Html(include_str!("index.html")))
        .size(800, 600)
        .resizable(true)
        .debug(true)
        .user_data(0)
        .invoke_handler(invoke_handler)
        .run()
        .unwrap();  
    run().unwrap(); */
    thread::spawn(move || {
        run().unwrap();
    });

    rocket::ignite()
        .mount("/", routes![index, m])
        .mount("/static",
                   StaticFiles::from(concat!(env!("CARGO_MANIFEST_DIR"), "/static")))

        .launch();

}

fn invoke_handler(_webview: &mut WebView<usize>, arg: &str) -> WVResult {
    match arg {
        "open" => {
        },
        "close" => {
        },
        _ => (),
    }
    Ok(())
}
/*
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

            let x = frame;
            // Add the header
            let header = b"--frame\r\nContent-Type: image/jpeg\r\n\r\n";
            let mut bytes = header.to_vec();
            bytes.extend(frame.as_bytes());
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
*/

