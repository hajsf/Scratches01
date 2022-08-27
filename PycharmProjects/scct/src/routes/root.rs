use crate::Chrome;
use rocket::State;
use rocket_contrib::templates::Template;
use std::collections::HashMap;

#[rocket::get("/")]
pub(crate) fn root(ui: State<Chrome>) -> Template {
    let mut context = HashMap::<&str, i32>::new();
    context.insert("value", 3);
    Template::render("index", &context)
}