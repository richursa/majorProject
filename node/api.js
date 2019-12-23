const express = require('express');
const app = express();
const bodyParser = require('body-parser')
const api = require('./router/router.js')
api.use(bodyParser.json());
app.use('/api',api);
app.listen(8080,()=>{
    console.log('started listening on port 8080')
});