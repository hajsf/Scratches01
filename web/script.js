/**
 * Moves the map to display over Berlin
 *
 * @param  {H.Map} map      A HERE Map instance within the application
 */
 function moveMapToBerlin(map){
    map.setCenter({lat:52.5159, lng:13.3777});
    map.setZoom(14);
  }
  
  /**
   * Boilerplate map initialization code starts below:
   */
  
  //Step 1: initialize communication with the platform
  // In your own code, replace variable window.apikey with your own apikey
  var platform = new H.service.Platform({
    apikey: "pjOMuy9ZVb2HJPpPY2_f9njk-YRG_hF1_BkQVuVWCDE" // window.apikey
  });
  var defaultLayers = platform.createDefaultLayers();
  
  //Step 2: initialize a map - this map is centered over Europe
  var map = new H.Map(document.getElementById('map'),
    defaultLayers.vector.normal.map,{
    center: {lat:50, lng:5},
    zoom: 4,
    pixelRatio: window.devicePixelRatio || 1
  });
  // add a resize listener to make sure that the map occupies the whole container
  window.addEventListener('resize', () => map.getViewPort().resize());
  
  //Step 3: make the map interactive
  // MapEvents enables the event system
  // Behavior implements default interactions for pan/zoom (also on mobile touch environments)
  var behavior = new H.mapevents.Behavior(new H.mapevents.MapEvents(map));
  
  // Create the default UI components
  var ui = H.ui.UI.createDefault(map, defaultLayers);
  
  // Now use the map as required...
  window.onload = function () {
    moveMapToBerlin(map);
  }

  /*        var platform = new H.service.Platform({ 
            apikey: "pjOMuy9ZVb2HJPpPY2_f9njk-YRG_hF1_BkQVuVWCDE" //"HERE_API_KEY"    
        }); 

		const lat = 52.5; 
		const lng = 13.4; 

        // Obtain the default map types from the platform object: 
        var defaultLayers = platform.createDefaultLayers(); 

        // Your current position 
        var myPosition = {lat: lat, lng: lng}; 

        // Instantiate (and display) a map object: 
        var map = new H.Map( 
            document.getElementById('mapContainer'), 
            defaultLayers.vector.normal.map, 
            { 
                zoom: 11, 
                center: myPosition 
            }); 

        var ui = H.ui.UI.createDefault(map, defaultLayers, 'en-US'); 
        var mapEvents = new H.mapevents.MapEvents(map); 
        var behavior = new H.mapevents.Behavior(mapEvents); 

        const marker = new H.map.Marker({lat: lat, lng: lng}); 

        map.addObject(marker); 

        marker.addEventListener('tap', function(evt) { 
        
        // Create an info bubble object at a specific geographic location: 
        var bubble = new H.ui.InfoBubble({ lng: lng, lat: lat }, { 
                content: '<p>Golang</p>' 
             }); 
        // Add info bubble to the UI: 
        ui.addBubble(bubble); 

    const marker = new H.map.Marker({lat: lat, lng: lng});
      map.addObject(marker);
  }); 

*/