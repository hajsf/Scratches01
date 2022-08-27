macro_rules! yo{
    ($num:expr) => {
        $num as f64 *2f64
    }
}

macro_rules! hi{
    ($name:expr) => {
        format!("Hi {}", $name);
    }
}

macro_rules! hi2{
    ($($element:expr), *) => {
        $(
           println!("{:#?}", $element);
        )*
    }
}

macro_rules! sum{
    ($($element:expr), *) => {
        {
            let mut sum = 0f64;
            $(sum += $element as f64;)*
            sum
        }
    };
}

fn main() {
    println!("{} {}", hi!("Hasan"), yo!(2.1));
    hi2!("Hasan", "Ali");
    println!("sum: {}", sum!(1,2.3,3));

    let x= hi2!();
    if true {hi2!("Hasan")} else { println!("Sorry"); }
}
