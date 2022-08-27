let source = null;
let live = document.querySelector('.badge')

function connect() {
  source = new EventSource("http://localhost:1234/sse/signal");

  source.onopen = function(){
    live = document.querySelector('.badge')
    live.style.backgroundColor = '#ff334b';
    console.log("connected")
  }
  live.style.backgroundColor = '#ff334b';
  source.onerror = function() {
    live.style.backgroundColor = '#30000614';
    source.close();
  }

  source.addEventListener("ping", function(event) {
    console.log("Ping event received")
    console.log(event)
  });

  source.onmessage = function (event) {
    console.log(event)
    var message = event.data
    console.log(message)
  /*  document.querySelector('#qr').innerHTML = "";
    console.log(event)
    var message = event.data
    
   // const data = JSON.parse(event.data).data;
   // alert(data)
    //console.log(message)
    //const obj = JSON.parse(message);
    //console.log(obj)
    msg = document.querySelector('#message')
    msg.innerHTML = message.replace(/\"/g, "")+'<br>'+msg.innerHTML;
    if (new String(message).valueOf() == "success" || new String(message).valueOf() == "timeout/Refreshing"
        || new String(message).valueOf() == "Already logged") {
    } else {
     /  var qrcode = new QRCode("qr", {
            text: message,
            width: 128,
            height: 128,
            colorDark : "#000000",
            colorLight : "#ffffff",
            correctLevel : QRCode.CorrectLevel.M
        }); *
    } */
    const obj = JSON.parse(message)
    const Details = document.createElement("details");
    const Summary = document.createElement("summary");
    const Content = document.createElement("p");
    const Badge = document.createElement("span");
    const Attachment = document.createElement(null);

    Badge.className = "chip"
    incomingMsg = obj.MessageText.replace(/\"/g, "")
    // const result = 10 > 5 ? 'yes' : 'no';
    const from = obj.Sender === obj.Group ? obj.Sender : obj.Sender + " / " + obj.Group;
    const msg = obj.MessageCaption === "" ? 
    incomingMsg : 
    incomingMsg +'<br>'+ "Caption: " + obj.MessageCaption;

    const body = obj.Uri === "" ? msg : msg + '<br>' + "<a href='" + obj.Uri + "' target='_blank'>Open attachment</a>"

    Details.setAttribute("data-type", obj.MessageType)
    Details.setAttribute("data-url", obj.Uri)
    Details.setAttribute("data-id", obj.MessageID)
    Details.setAttribute("data-sender", obj.Sender)
    
    Content.innerHTML = "Text: " + body + "<br><label for='reply'>Reply: </label>" +
    "<input type='text' id='reply' name='reply' style='width:80%'>" +
   // "<input type='button' onclick='location.href=`https://google.com`;' value='WhatsApp Reply' />"
   "<input type='button' onclick='send(this);' value='WhatsApp Reply' />"

    if(obj.MessageType === "image") {
      Attachment.innerHTML = '<br><br>' + "<img src='"+obj.Uri+"' alt='Message attachment' width: 80%;' height='600'>"
    } else if(obj.MessageType === "audio"){
    //  console.log(`Audio source: ${obj.Uri}`)
      Attachment.innerHTML = '<br>' + "<audio controls> "+
        "<source src='"+obj.Uri+"' type='audio/ogg'>" +
        "<source src='"+obj.Uri+"' type='audio/mpeg'>" +
        "<source src='"+obj.Uri+"' type='audio/oga'>" +
     //   "<source src='test.oga' type='audio/ogg; codecs=`vorbis`'></source>" +
        "Your browser does not support the audio element." +
      "</audio>"
    } else if(obj.MessageType === "video"){
      Attachment.innerHTML = '<br>' + "<video width: 100%;' height='600' controls> " +
        "<source src='"+obj.Uri+"' type='video/ogg'>" +
        "<source src='"+obj.Uri+"' type='video/mp4'>" +
        "<source src='"+obj.Uri+"' type='video/m4v'>" +
        "Your browser does not support the video tag." +
      "</video>"
    } else if(obj.MessageType === "document" && obj.Uri.type === 'application/pdf'){
      Attachment.innerHTML = '<br>' +  "<embed id='pdf' type='application/pdf'" +
      "src='"+obj.Uri+"' style='width: 100%;' height='600'>"
    } else if (obj.MessageType === "document" && obj.Uri.type !== 'application/pdf'){
      var ext = obj.Uri.split('.').reverse()[0]
      Content.innerHTML += `<br> file recieved of of extension ${ext}`
    }

    console.log(body)
    let isMentioned = body.includes("@966506889946") || body.includes("966506889946@");
    console.log(`is mentioned: ${isMentioned}`)
    if (isMentioned){
      console.log(`Mentioned`)
      Badge.classList.add('danger')
      Badge.textContent = "Mentioned" // "<span class='chip danger'>Mentioned</span>"
      console.log(`Badge is: ${Badge.innerHTML}`)
    }

    Summary.innerHTML = "From: "+ '<code>' + obj.Name + '</code>' + " Number: " + from + " By: " + obj.Time + " "
    // "<span class='chip danger'>Mentioned</span>"
    
    Summary.appendChild(Badge);
    Details.appendChild(Summary);

    Content.appendChild(Attachment);
    Details.appendChild(Content);

    Details.addEventListener('click', function(){
      chip = this.querySelector('.chip')
      if (chip != null) {
        chip.remove();
      }
    })

    document.body.insertBefore(Details, document.body.children[1]);

   let ds= document.querySelectorAll('details');
   window.addEventListener('click', function(event){
      ds.forEach(function(d){
        const isClickInside = d.contains(event.target);
        if(!isClickInside){
          d.removeAttribute('open')
        }
      })
    }, true);

}

}
connect();

let reconnecting = false;
setInterval(() => {
    if (source.readyState == EventSource.CLOSED) {
        reconnecting = true;
        console.log("reconnecting...");
        connect();
    } else if (reconnecting) {
        reconnecting = false
        console.log("reconnected!");
    }
}, 3000);


function send(e){
  sender = e.parentNode.parentNode.dataset.sender; // .querySelector("details")

  reply = document.querySelector("#reply");
  response = encodeURIComponent(reply.value)
  console.log(e.parentNode.parentNode)
  console.log(sender)

  window.open(
    "https://wa.me/" + sender + "/?text=" + response,
    '_blank' // <- This is what makes it open in a new window.
  );

// Below are alternate options
 // location.href = r
//  let a= document.createElement('a');
//  a.target= '_blank';
//  a.href= "https://wa.me/" + sender + "/?text=" + response;
//  a.click();

  // <a id="anchorID" href="mynewurl" target="_blank"></a>
}

// To be checked
// will be called automatically whenever the server sends a message with the event field set to "ping"
// echo "event: ping\n"; then followed by echo 'data: {"time": "' . $curDate . '"}';
// fmt.Fprint(w, "event: ping\n\n", c) then followed by fmt.Fprintf(w, "data: %v\n\n", c)
source.addEventListener("ping", function(event) {
  console.log("Ping event received")
  console.log(event)
//  const newElement = document.createElement("li");
//  const eventList = document.getElementById("list");
//  const time = JSON.parse(event.data).time;
//  newElement.textContent = "ping at " + time;
//  eventList.appendChild(newElement);
});

