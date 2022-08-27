function loadScript(src) {
  var head = document.getElementsByTagName("head")[0];
  var script = document.createElement("script");
  script.src = src;
  script.defer = true;
  var done = false;
  script.onload = script.onreadystatechange = function() {
        // attach to both events for cross browser finish detection:
      if ( !done && (!this.readyState ||
          this.readyState == "loaded" || this.readyState == "complete") ) {
          done = true;
          // cleans up a little memory:
          script.onload = script.onreadystatechange = null;
          // to avoid douple loading
          head.removeChild(script);
        }
      };
      head.appendChild(script);
}

var $view = document.body.children.view;
var $ = function (el) { return document.querySelector(el); }
//var $onClick = function (el, fn) { return document.querySelector(el).addEventListener("click", fn); }
var $action = function(shadow, el, fn){    // action function "fn" upon the click of element "el" in the shadowRoot "shadow"
var children = shadow.children;
for (var i = 0; i < children.length; i++) {
    var child = children[i];
       if (child.id === el) {
         return shadow.children[i].addEventListener('click', fn);
         break;
      }
  }
}

loadScript('static/scripts/load.js');
loadScript('static/custom_elements/timer.js')
loadScript('static/custom_events/tick_event.js')

loadScript('static/custom_elements/menu.js')
loadScript('static/custom_elements/side_menu.js')
