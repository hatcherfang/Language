var http = require('http');
var querystring = require('querystring');
http.get('http://127.0.0.1:8001', function(response){
   var body = [];
   console.log(response.statusCode);
   console.log(response.headers);
   //console.log(response);
   response.on('data', function(chunk){
      body.push(chunk);
   });
   response.on('end', function(){
      body = Buffer.concat(body);
      console.log(body.toString());
   });
});
