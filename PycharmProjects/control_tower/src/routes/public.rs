use rocket::*;
use rocket::http::*;
use std::path::PathBuf;
use rocket::response;
use std::ffi::OsStr;
use std::io::Cursor;

use rust_embed::RustEmbed;
use rocket_contrib::templates::Template;
use std::collections::HashMap;
use crate::Chrome;

#[derive(RustEmbed)]
#[folder = "static"]
pub(crate) struct StaticAsset;

#[derive(RustEmbed)]
#[folder = "templates"]
pub(crate) struct TemplateAsset;

#[get("/hi")]
pub(crate) fn public<'r>(ui: State<Chrome>) -> response::Result<'r> {
  StaticAsset::get("index.html").map_or_else(
    || Err(Status::NotFound),
    |d| response::Response::build().header(ContentType::HTML).sized_body(Cursor::new(d)).ok(),
  )
}

#[rocket::get("/")]
pub(crate) fn root(ui: State<Chrome>) -> Template {
    let mut context = HashMap::<&str, &str>::new();
    context.insert("org_short_name", "AIS");
    context.insert("org_long_name", "Aujan Industrial Solutions");
    Template::render("index", &context)
}

#[get("/static/<file..>")]
pub(crate) fn dist<'r>(file: PathBuf) -> response::Result<'r> {
  let filename = file.display().to_string();
  StaticAsset::get(&filename).map_or_else(
    || Err(Status::NotFound),
    |d| {
      let ext = file
          .as_path()
          .extension()
          .and_then(OsStr::to_str)
          .ok_or_else(|| Status::new(400, "Could not get file extension"))?;
      let content_type = ContentType::from_extension(ext).
          ok_or_else(|| Status::new(400, "Could not get file content type"))?;
      response::Response::build().header(content_type).sized_body(Cursor::new(d)).ok()
    },
  )
}


#[get("/templates/<file..>")]
pub(crate) fn templ<'r>(file: PathBuf) -> response::Result<'r> {
  println!("looking for template");
  let filename = file.display().to_string();
  //let x = TemplateAsset::g    //::get(&filename).unwrap();
  TemplateAsset::get(&filename).map_or_else(
    || Err(Status::NotFound),
    |d| {
      let ext = file
        .as_path()
        .extension()
        .and_then(OsStr::to_str)
        .ok_or_else(|| Status::new(400, "Could not get file extension"))?;
      let content_type = ContentType::from_extension(ext).
          ok_or_else(|| Status::new(400, "Could not get file content type"))?;
      response::Response::build().header(content_type).sized_body(Cursor::new(d)).ok()
    },
  )
}
