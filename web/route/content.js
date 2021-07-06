var express = require('express')
var router = express.Router()

// middleware that is specific to this router
router.use(function timeLog (req, res, next) {
    // console.log(__filename)
    console.log('Time: ', Date.now())
    next()
})
// define the home page route
router.get('/', function (req, res) {
    res.render('content.ejs');
})
// define the about route
router.get('/payment', function (req, res) {
    res.render('payment.ejs');
})

router.post('/payment', function (req, res) {
    let id = req.body.id;
    res.render('ticketList.ejs');
})

module.exports = router