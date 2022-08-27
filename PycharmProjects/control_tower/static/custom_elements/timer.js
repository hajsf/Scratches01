customElements.define("time-formatted", class TimeFormatted extends HTMLElement {

  render() {
    let date = new Date(this.getAttribute('datetime') || Date.now());

    this.innerHTML = new Intl.DateTimeFormat("default", {
      year: this.getAttribute('year') || undefined,
      month: this.getAttribute('month') || undefined,
      day: this.getAttribute('day') || undefined,
      hour: this.getAttribute('hour') || undefined,
      minute: this.getAttribute('minute') || undefined,
      second: this.getAttribute('second') || undefined,
      timeZoneName: this.getAttribute('time-zone-name') || undefined,
    }).format(date);
  }

  connectedCallback() {
    if (!this.rendered) {
      this.render();
      this.rendered = true;
    }
  }

  static get observedAttributes() {
    return ['datetime', 'year', 'month', 'day', 'hour', 'minute', 'second', 'time-zone-name'];
  }

  attributeChangedCallback(name, oldValue, newValue) {
    this.render();
  }
});

customElements.define("live-timer", class LiveTimer extends HTMLElement {
  constructor() {
    super();
      this.innerHTML = `
        <time-formatted hour="numeric" minute="numeric" second="numeric">
        </time-formatted>
        `;

      this.timerElem = this.firstElementChild;

//      this.addEventListener('tick', event =>
//        this.timerElem.setAttribute('datetime', event.detail.time)
//      );
  }

/*  connectedCallback() {
      this.timer = setInterval(() =>
          this.dispatchEvent(new CustomEvent('tick', {
            detail: {
              time: new Date()
            }
          }))
      , 1000);
  } */

  disconnectedCallback() {
    clearInterval(this.timer); // important to let the element be garbage-collected
  }
});
