var tick_event = setInterval(() => {
    window.dispatchEvent(new CustomEvent('tick', {
      detail: {
        time: new Date()
      }
    }));
 }
, 1000);

window.addEventListener('tick', event =>
  $("#timer").timerElem.setAttribute('datetime', event.detail.time)
);

/*
var tick_event = new CustomEvent("tick", {
  detail: {
    time: new Date()
  }
});
*/
/*
var tick_event = setInterval(() =>
new CustomEvent("tick", {
  detail: {
      time: new Date()
    }
}), 1000);
*/
