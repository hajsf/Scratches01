// <ul is="side-menu"></ul>
customElements.define('side-menu', class SideMenu extends HTMLElement {
  ul = document.createElement('ul');
  constructor() {
    self = super();
    const style = document.createElement('style');
    const shadow = this.attachShadow({mode: 'open'});
    if (self.hasAttribute('data-show')) {
        self.status = `${self.getAttribute('data-show')}`;
    } else {
        self.status = true;
    }
    self.ul.innerHTML= `
      <li><a href="#home">Home</a></li>
      <li>Orders
        <ul>
          <li>Creat</li>
          <li>Update</li>
        </ul>
      </li>
      <li><a href="#news">News</a></li>
    `;

  self.ul.style.width = '210px';
  self.ul.style.height = '500px';
  self.ul.style.background= 'white';
  self.ul.style.zIndex = "1";
  style.textContent = `
    a {  text-decoration: none; color: #fff; }
    ul { margin: 0; padding-left: 0; }
    ul, li { background: coral; list-style: none;}
    li:hover { background: red; cursor: pointer; }
    ul li ul { position: absolute; transition: all 0.5s ease;
            margin-top: -2rem; left: 211px; }
    ul li > ul { display: none; }
    li { display: block; color: #fff;  border: 1px solid #ddd;
          padding: 1rem; text-decoration: none; transition-duration: 0.5s; }

    ul li:hover > ul,
    ul li:focus-within > ul,
    ul li ul:hover {
      visibility: visible;
      display: block;
    }
   `;
     // focus-within: focusing on the link (a) within the li under the ul
    // Attach the created element (and style) to the shadow dom
    shadow.appendChild(style);
    shadow.appendChild(self.ul);
  }
});
