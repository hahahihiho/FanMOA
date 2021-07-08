const express = require('express')
const utils = require("../util/utils")

const app = express();
const router = express.Router();

const content = require('./content');
const myPage = require('./myPage');
const user = require('./user');

// middleware that is specific to this router
router.use(function timeLog (req, res, next) {
    console.log(__filename);
    console.log('Time: ', Date.now())
    next()
})

// define the home page route
router.get('/', async function (req, res) {
    const tx = await utils.submitTx("getAllEvents");
    let result;
    if(tx.success) {
        console.log(tx.result)
        const event_list = tx.result.map(e => e.Key);
        const prices = tx.result.map(e => e.Record.Fee)
        const ids = event_list;
        result = {
            ids : ids,
            links : ids.map((id)=>"/content?ei="+id),
            titles : ids.map((id)=>"Fan meeting "+id),
            prices : prices 
            // need to add thumbn ail
        }
    } else {
        result = {};
    }
    const entry = utils.makeEntry(tx.success,result,tx.error);
    res.render('index',entry);
})

router.use('/content', content);
router.use('/mypage', myPage);
router.use('/user', user);


module.exports = router
