class MyMenu extends HTMLElement {
     menu = document.createElement('div');
  //   status = true;
    constructor() {
      super();
      let self = this;
      const shadow = this.attachShadow({mode: 'open'});
      if (self.hasAttribute('data-show')) {
          self.status = `${self.getAttribute('data-show')}`;
      } else {
          self.status = true;
      }

      self.menu.innerHTML = 'Main Menu'+
  					   '<label id="items" class="menulabel ion-android-promotion"> Items</label>'+
  					   '<label id="Partners" class="menulabel ion-android-contacts"> Partners</label>'+
  					   '<label id="Orders" class="menulabel ion-android-archive"> Orders</label>'+
  					   '<label id="Transactions" class="menulabel ion-ios-pulse-strong"> Transactions</label>'+
  					   '<label id="Planning" class="menulabel ion-android-promotion"> Planning</label>'+
  					   '<label id="Dashboard" class="menulabel ion-android-promotion"> Dashboard</label>'+
                       '<p></p>'+
                     '';
      shadow.appendChild(self.menu);
  }

  get value() {
    console.log("question came");
    return this.getAttribute('data-show');
  }

  set value(newValue) {
    console.log("data set");
    this.setAttribute('data-show', newValue);
  }
}
customElements.define('fonix-menu', MyMenu);
