//#![windows_subsystem = "windows"]
#![feature(decl_macro, proc_macro_hygiene)]
// To minimize executable size > start
use std::alloc::System;
use rocket_contrib::templates::Template;

#[global_allocator]
static A: System = System;
// To minimize executable size < end
mod constants;
mod routes;
use routes::public::*;

use rocket::*;
use alcro::{UI, UIBuilder, Content, JSObject};
use std::sync::Arc;
use std::process::exit;
use std::thread::spawn;

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

    ignite()
        .attach(Template::fairing())
        .mount("/", routes![public, dist, templ, root])
        .manage(Chrome(ui))
        .launch();
}
