// Snap.svg http://snapsvg.io/ 

function configToggleAnimation() {
    var toggle = document.querySelector('.nav-toggle'),
      nav = document.querySelector('.nav'),
      snap = Snap(document.querySelector('.nav-morph svg')),
      nav_morph = document.querySelector('.nav-morph'),
      path = snap.select('path'),
      reset = path.attr('d'),
      open = nav_morph.getAttribute('data-open'),
      close = nav_morph.getAttribute('data-close'),
      speed = 250,
      speed_back = 800,
      easing = mina.easeinout,
      easing_back = mina.elastic,
      isOpen = false;

      content = document.querySelector("#content")
      header = document.querySelector("#header")
 
    toggle.addEventListener('click', function() {
      // si ouvert on ferme
      if (isOpen) {
        header.style.marginLeft="0rem"
        content.style.marginLeft="0rem"
        path.stop().animate({
          'path': close
        }, speed, easing, function() {
          path.animate({
            'path': reset
          }, speed_back, easing_back);
          isOpen = false;
        });
        nav.classList.remove('nav--open');
        
      } else {
        path.stop().animate({
          'path': open
        }, speed, easing, function() {
          path.animate({
            'path': reset
          }, speed_back, easing_back);
          isOpen = true;
        });
        nav.classList.add('nav--open');
        header.style.marginLeft="9rem"
        content.style.marginLeft="9rem"
      }
    });
 
  }
 
  function initialize() {
    configToggleAnimation();
  }
 
  document.addEventListener('DOMContentLoaded', initialize);