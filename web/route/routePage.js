const express = require('express')
const utils = require("../util/utils")

const app = express();
const router = express.Router();

const content = require('./content');
const myPage = require('./myPage');
const user = require('./user');

const data_util = require("../data/data");

// middleware that is specific to this router
router.use(function timeLog (req, res, next) {
    console.log(__filename);
    console.log('Time: ', Date.now())
    next()
})

// define the home page route
router.get('/', function (req, res) {
    const tx = utils.submitTx("getAllEvents");
    console.log(tx);
    if(tx.success) {
        console.log(tx.result)
    }
    const ids = data_util.event.ids;
    const success = true;
    const error = null;
    const result = {
        ids : ids,
        links : ids.map((id)=>"/content?ei="+id),
        titles : ids.map((id)=>"Fan meeting "+id)
        // need to add thumbn ail
    }
    let entry = data_util.makeEntry(success,result,error);
    res.render('index',entry);
})

router.use('/content', content);
router.use('/mypage', myPage);
router.use('/user', user);


module.exports = router