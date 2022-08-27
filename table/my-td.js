class MyInput extends HTMLTableDataCellElement {
    constructor() {
      // Always call super first in constructor
      super();
  
      // Create a shadow root
      const shadow = this.attachShadow({mode: 'open'});

      const master = document.createElement('td');
      master.setAttribute('class', 'master');
      const slave = document.createElement('input');
      slave.style.display = 'none';
      slave.setAttribute('class', 'slave');
    //  const info = document.createElement('span');

      const text = this.getAttribute('data-text');
    //  master.setAttribute('placeholder', text);
    //  master.value = text;
      console.log(text);

      master.addEventListener('click', function(e) {
        slave.style.display = '';
        slave.value = master.value;
        master.style.display = 'none';
      });

      slave.addEventListener("focusout", function(event) {
        slave.style.display = 'none';
        master.style.display = '';
        master.value = slave.value;

      })
      
      // Apply external styles to the shadow dom
      const linkElem = document.createElement('link');
      linkElem.setAttribute('rel', 'stylesheet');
      linkElem.setAttribute('href', 'my-input.css');
      // Create some CSS to apply to the shadow dom
      const style = document.createElement('style');
      console.log(style.isConnected);
  
      style.textContent = `
      .slave {
        display: 'none';
        }
     `;

      // Attach the created element to the shadow dom
      shadow.appendChild(linkElem);
    //  shadow.appendChild(style);
      // Attach the created elements to the shadow dom
      shadow.appendChild(master);
    //  shadow.appendChild(slave);
    }

    get value(){ this.master.value; }
    set value() { this.matches.value = this.value || this.text}
}
  
// Define the new element
//customElements.define('my-cell', MyCell);
customElements.define('my-cell', MyCell, { extends: "td" });