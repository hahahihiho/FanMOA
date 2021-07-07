var express = require('express')
var router = express.Router()

const data_util = require("../data/data")

// middleware that is specific to this router
router.use(function timeLog (req, res, next) {
    // console.log(__filename)
    console.log('Time: ', Date.now())
    next()
})
// define the home page route
router.get('/', function (req, res) {
    const ei = parseInt(req.query.ei);
    
    const data = data_util.event
    const i = data.ids.indexOf(ei);
    const keys = Object.keys(data);

    let result = {};
    keys.forEach((k) => {
        result[k] = data[k][i];
    })
    const entry = data_util.makeEntry(true,result,null);
    res.render('content.ejs',entry);
})
// define the about route
router.get('/payment', function (req, res) {
    const ei = parseInt(req.query.ei);
    
    const data = data_util.event
    const i = data.ids.indexOf(ei);
    const keys = Object.keys(data);

    let result = {};
    keys.forEach((k) => {
        result[k] = data[k][i];
    })
    const entry = data_util.makeEntry(true,result,null);
    res.render('payment.ejs',entry);
})

router.post('/payment', function (req, res) {
    let id = req.body.id;
    res.render('ticketList.ejs');
})

module.exports = router