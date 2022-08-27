console.log("test");

class Rectangle {
 //   height = 0;
  //  width;
    constructor(height, width) {
      this.height = height;
      this.width = width;
    }
    // Getter
    get area() {
      return this.calcArea();
    }
    // Setter
    set _height(h) {
         this.height =  h * 3;
      }    
    // Method
    calcArea() {
      return this.height * this.width;
    }
  }
  
  const square = new Rectangle(10, 10);
  square._height = 5;
  
  console.log(`${square.height} * ${square.width} = ${square.area}`); // 100