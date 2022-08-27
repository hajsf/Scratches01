#![feature(decl_macro, proc_macro_hygiene)]
use rocket_contrib::templates::{Template};

fn main() {
    println!("Hello, world!");

    rocket::ignite()
    .attach(Template::custom(|engine|{
        engine.tera.add_template_file("")
       // engine.tera.add_template_files("",);
    }))
    .launch();
//    rocket::ignite().attach(Template::custom(|e|{
//        e.tera.
//    }))
}
