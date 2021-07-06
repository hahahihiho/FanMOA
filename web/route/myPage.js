var express = require('express')
var router = express.Router()

// middleware that is specific to this router
router.use(function timeLog (req, res, next) {
    console.log(__filename)
    console.log('Time: ', Date.now())
    next()
})
// define the home page route
router.get('/', function (req, res) {
    res.render('myPage.ejs');
})
// define the about route
router.get('/ticketlist', function (req, res) {
    const mi = req.query.mi;
    if(mi==undefined){
        res.render('ticketList.ejs');
    }else{
        // generate qrCode
        res.render("qrCode.ejs");
    }
})
router.get('/registermeeting', function (req, res) {
    res.render('registerMeeting.ejs');
})

router.get('/refund', function (req, res) {
    res.render('refund.ejs');
})
router.post('/refund', function (req, res) {
    res.render('refund.ejs');
})
router.get('/history', function (req, res) {
    const mi = req.query.mi;
    if(mi==undefined){
        res.render('history.ejs');
    }else{
        res.render("historyContent.ejs");
    }
})

router.get('/cancellist', function (req, res) {
    res.render('cancelList.ejs');
})
router.post('/cancellist', function (req, res) {
    res.render('cancellist.ejs');
})


module.exports = router