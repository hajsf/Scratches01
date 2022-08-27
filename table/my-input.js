class MyInput extends HTMLElement {
    master = document.createElement('input');
    slave = document.createElement('input');

    constructor() {
      // Always call super first in constructor
      super();
      let self = this;   // Using `self` to avoid confusion among many `this` used
  
      // Create a shadow root
      const shadow = this.attachShadow({mode: 'open'});

      
      self.master.setAttribute('class', 'master');
      
      self.slave.style.display = 'none';
      self.slave.setAttribute('class', 'slave');
    //  const info = document.createElement('span');

    if (self.hasAttribute('data-text')) {
        self.value = `${self.getAttribute('data-text')}`;
    } else {
        self.value = '';
    }

    self.master.value = self.value;

    self.addEventListener('click', function(e) {
        // Here `this` is for `s elf` itself 
        this.slave.style.display = '';
        this.slave.value = this.value;
        this.master.style.display = 'none';
        this.slave.focus();
      });

    ['focusout','keydown'].forEach( function(evt) {
        self.slave.addEventListener(evt, function(event) {
            // Here `this` is for the slave, i.e. `self.slave`
            if ((event.type === 'keydown' && event.which === 27) || event.type === 'focusout') {
                this.style.display = 'none';
                this.parentNode.querySelector('.master').style.display = '';
                this.parentNode.querySelector('.master').value = this.value;
                console.log('out');
            }
        }, false);
    });


 /*   self.slave.addEventListener("focusout", function(event) { })
      self.slave.addEventListener("keydown", function(event) {
         if (event.which === 27) {  // Esc
            }
      })
*/
      
      // Apply external styles to the shadow dom
      const linkElem = document.createElement('link');
      linkElem.setAttribute('rel', 'stylesheet');
      linkElem.setAttribute('href', 'my-input.css');
      // Create some CSS to apply to the shadow dom
      const style = document.createElement('style');
      console.log(style.isConnected);
  
    self.style.textContent = `
      .slave {
        display: 'none';
        }
     `;

      // Attach the created element to the shadow dom
      shadow.appendChild(linkElem);
    //  shadow.appendChild(style);
      // Attach the created elements to the shadow dom
      shadow.appendChild(this.master);
      shadow.appendChild(this.slave);
    }

   get value(){ return this.entry; }
   set value(val) { console.log(`got = ${val}`); this.entry = val}
}
  
// Define the new element
customElements.define('my-input', MyInput);
//customElements.define('my-input', MyInput, { extends: "input" });