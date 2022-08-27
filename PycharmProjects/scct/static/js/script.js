    alert("hi");
     window.addEventListener('beforeunload', function (e) {
           e.preventDefault();
           e.returnValue = '';
           exit(1);
     });

          function context() {
          alert("ready!");
          window.event.returnValue = '';
     }
// document.addEventListener("DOMContentLoaded", function(event) {}