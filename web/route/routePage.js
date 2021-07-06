const express = require('express')
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
router.get('/', function (req, res) {
    res.render('index');
})

router.use('/content', content);
router.use('/myPage', myPage);
router.use('/user', user);


module.exports = router