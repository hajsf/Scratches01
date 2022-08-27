 // = true;
//console.log(`value: ${$('#menu').getAttribute('data-show')}`);
$('#menu').setAttribute('data-show', true)
//console.log(`value: ${$('#menu').getAttribute('data-show')}`);

/*
Object.observe(menu, function(changes) {
  if(menu.show == true)   $('fonix-menu').style.display = 'block';
  if(menu.show == false)  $('fonix-menu').style.display = 'none';
});
*/ /*
$onClick('#menuLines', function(){
//  menu.show = !menu.show;
  $('#fonix-footer').classList.remove('serverSUCCESS','serverERROR','socketMSG');
  $('#fonix-footer').innerHTML='Menu loaded';
});
*/
