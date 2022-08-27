#![windows_subsystem = "windows"]

#![feature(decl_macro, proc_macro_hygiene)]
// To minimize executable size > start
use std::alloc::System;

#[global_allocator]
static A: System = System;
// To minimize executable size < end

mod routes;

use rocket_contrib::serve::StaticFiles;
use rocket_contrib::templates::Template;
use std::sync::Arc;
use alcro::{UI, UIBuilder, Content, JSObject};
use std::thread::{spawn, sleep};
use std::process::exit;
use std::env;
use std::time::Duration;

use rust_embed::RustEmbed;

#[derive(RustEmbed)]
#[folder = "static"]
struct Static;

#[derive(RustEmbed)]
#[folder = "templates"]
struct Templates;

struct Chrome(Arc<UI>);
fn main() {
    println!("Hello, world!");

    let ui = Arc::new(
        UIBuilder::new()
            .content(Content::Url("http://localhost:8000"))
            .size(1200, 800)
            .run()
    );

  //  let ui_clone = ui.clone();
    let ui_clone = Arc::clone(&ui);
    ui.bind("exit", move |args| {
    match args[0].as_i64() {
            Some(0i64) => { println!("good bye: normal exit");
                ui_clone.close(); exit(0);  },
            Some(1i64) => { println!("good bye: user closed browser");
                exit(0); },
            _ => {}
    };
        Ok(JSObject::Null)
    }).unwrap();

    let ui_clone2 = Arc::clone(&ui);
    spawn(move||{
        loop {
            if ui_clone2.done() { println!("good bye: exit before load");
                exit(0); }
        }
    });
    let executable = env::current_exe().unwrap();
    let exe_dir = match executable.parent() {
        Some(parent) => parent,
        _ => panic!()
    };
    let static_path = exe_dir.join("static");
    rocket::ignite()
        .attach(Template::fairing())
        .mount("/static",
                StaticFiles::from("static"))
       //        StaticFiles::from(concat!(env!("CARGO_MANIFEST_DIR", "/static"))))
      //             StaticFiles::from(static_path))
        .mount("/", rocket::routes![routes::root::root])
        .manage(Chrome(ui))   // ui value used here after move
        .launch();
}
