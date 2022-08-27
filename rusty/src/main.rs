#![feature(proc_macro_hygiene)]
use inline_python::python;

fn main() {
    python! {
        import cv2
        print("Hello, Python")
        print(cv2.__version__)
        for x in range(1, 1000000, 1):
            print(x)
    }
}
