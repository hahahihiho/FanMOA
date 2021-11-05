var express = require('express')
const data_util = require('../data/data')
const utils = require("../util/utils")
var router = express.Router()

// middleware that is specific to this router
router.use(function timeLog (req, res, next) {
    console.log(__filename)
    console.log('Time: ', Date.now())
    next()
})
// define the home page route
router.get('/', function (req, res) {
    const data = data_util.mypage

    let result = {
        mypage_list : data.mypage_list,
        mypage_link : data.mypage_link
    };
    const entry = utils.makeEntry(true,result,null);
    res.render('myPage.ejs',entry);
})
// define the about route
router.get('/ticketlist', function (req, res) {
    const ei = req.query.ei;
    if(ei==undefined){
        const data = data_util.event;
        const ids = data_util.mytickets;
        const result = [];
        ids.forEach(id=>{
            i = data.ids.indexOf(id);
            result.push(data.names[i]);
        })
        const entry = utils.makeEntry(true,result,null);
        res.render('ticketList.ejs',entry);
    }else{
        // generate qrCode
        const i = data.ids.indexOf(ei);
        const keys = Object.keys(data);
    
        let result = {};
        keys.forEach((k) => {
            result[k] = data[k][i];
        })
        const entry = utils.makeEntry(true,result,null);
        res.render("qrCode.ejs");
    }
})
router.get('/registerevent', function (req, res) {
    res.render('registerEvent.ejs');
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