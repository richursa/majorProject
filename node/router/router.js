const express = require('express')
const router = express.Router()
router.get('/getCount',(request,response)=>{
    return 4
})