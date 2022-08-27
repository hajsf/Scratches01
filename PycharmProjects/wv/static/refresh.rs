/*
use web_view::*;
use std::thread;
use opencv::{core, highgui, prelude::*, videoio, core::Size2i, imgcodecs};
use opencv::videoio::VideoCapture;

use rocket_contrib::serve::StaticFiles;
use rust_embed::RustEmbed;
#[derive(RustEmbed)]
#[folder = "static"]
struct Asset;

fn main() {
    thread::spawn(move || {
        rocket::ignite()
            .mount("/static",
                   StaticFiles::from(concat!(env!("CARGO_MANIFEST_DIR"), "/static")))
            .launch();
    });

    let webview:WebView<()> = web_view::builder()
         .title("My Project")
         .content(Content::Url("http://localhost:8000/static/"))
         .size(800, 600)
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
            "Run openCV" => openCV(wv),
            "Stop openCV" =>     wv.eval(&format!(r#"
                            alert("not enabled now");
                        "#)).unwrap(),
            //videoio::VideoCapture::release(unsafe { &mut cam }).expect("Can not release Cam"),
            _ => (),
        }
    Ok(())
}

fn openCV(wv: &mut WebView<()>){
    run(wv).unwrap();
}

fn refresh(wv: &mut WebView<()>){
    wv.eval(&format!(r#"
        var dt = new Date();
        img = document.getElementById('img');
        img.src="/static/savedImage.jpg"+ "?" + dt.getTime();
    "#)).unwrap();
}

// static mut cam : VideoCapture = videoio::VideoCapture::new(0, videoio::CAP_ANY).unwrap();  // 0 is the default camera

fn run(wv: &mut WebView<()>) -> opencv::Result<()> {
    let window = "video capture";
    let filename = "static/savedImage.jpg";
    #[cfg(feature = "opencv-32")]
    let mut cam = videoio::VideoCapture::new_default(0)?;  // 0 is the default camera
    #[cfg(not(feature = "opencv-32"))]
    let mut cam = videoio::VideoCapture::new(0, videoio::CAP_ANY)?;  // 0 is the default camer

    let opened = videoio::VideoCapture::is_opened(unsafe { &cam })?;

    if !opened {
        panic!("Unable to open default camera!");
    }
    loop {
        let mut frame = core::Mat::default()?;
        unsafe { &mut cam }.read(&mut frame)?;
        if frame.size()?.width > 0 {

            imgcodecs::imwrite(filename, &frame, &opencv::types::VectorOfi32::new());

            refresh(wv);
        }
            wv.eval(&format!(r#"

            "#)).unwrap();

        let key = highgui::wait_key(10)?;
        if key > 0 && key != 255 {
            break;
        }
    }
    cam.release().expect("Can not release Cam");
    Ok(())
}
*/
