var http = require('http');
var querystring = require('querystring');
var postData = querystring.stringify({
    'key1': 'value1',
    'key2': 'value2'
});

var options = {
    hostname: '127.0.0.1', 
    port: '8001',
    path: '/', 
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded', 
      'Content-Length': Buffer.byteLength(postData)
    }
};

var req = http.request(options, (res) => {
  console.log(`STATUS: ${res.statusCode}`);
  console.log(`HEADERS: ${JSON.stringify(res.headers)}`);
  res.setEncoding('utf8');
  res.on('data', (chunk) => {
    console.log(`BODY: ${chunk}`);
  });
  res.on('end', () => {
    console.log('No more data in response.')
  });
});

req.on('error', (e) => {
  console.log(`problem with request:${e.message}`);
});

// write data to request body
req.write(postData);
req.end();
