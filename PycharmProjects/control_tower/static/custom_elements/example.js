customElements.define('tag-name', class [ClassName] extends HTMLxxxElement {

} [, {extends: 'xxx'}]);  // use {extends: } block only if extend an element
// call it as <tag-name> or as <xxx is='tag-name' if extends an element

// “Autonomous” – new tags, extending HTMLElement.
class MyElement extends HTMLElement {
  constructor() { super(); /* ... */ }
  connectedCallback() { /* ... */ }
  disconnectedCallback() { /* ... */  }
  static get observedAttributes() { return [/* ... */]; }
  attributeChangedCallback(name, oldValue, newValue) { /* ... */ }
  adoptedCallback() { /* ... */ }
 }
customElements.define('my-element', MyElement);
/* <my-element> */

// Customized built-in elements” – extensions of existing elements.
class MyButton extends HTMLButtonElement { /*...*/ }
customElements.define('my-button', MyElement, {extends: 'button'});
/* <button is="my-button"> */

// The button that says "hello" on click
class HelloButton extends HTMLButtonElement {
  constructor() {
    super();
    this.addEventListener('click', () => alert("Hello!"));
  }
}

customElements.define('hello-button', HelloButton, {extends: 'button'});
</script>

<button is="hello-button">Click me</button>

<button is="hello-button" disabled>Disabled</button>
