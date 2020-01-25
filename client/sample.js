const request = require('request');
var val
function callApi() {
  request('http://localhost:8080/api/getBlock/1', { json: true }, (err, res, body) => {
  if (err) { return console.log(err); }
  //body il value und 
  console.log(body)
});
}
callApi()