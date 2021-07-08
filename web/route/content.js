var express = require('express')
var router = express.Router()

const utils = require("../util/utils")

// middleware that is specific to this router
router.use(function timeLog (req, res, next) {
    // console.log(__filename)
    console.log('Time: ', Date.now())
    next()
})
// define the home page route
router.get('/', async function (req, res) {
    const ei = req.query.ei;

    const tx = await utils.submitTx("getEvent",ei);
    const entry = utils.makeEntry(tx.success,tx.result,null);

    res.render('content.ejs',entry);
})
// define the about route
router.get('/payment', async function (req, res) {
    const ei = req.query.ei;
    
    const tx = await utils.submitTx("getEvent",ei);
    const entry = utils.makeEntry(tx.success,tx.result,null);

    res.render('payment.ejs',entry);
})

router.post('/payment', function (req, res) {
    let id = req.body.id;
    res.render('ticketList.ejs');
})

module.exports = router