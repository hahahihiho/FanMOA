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
    res.render('user.ejs');
})
// define the about route
router.get('/signup', function (req, res) {
    res.render('signUp.ejs');
})
router.get('/login', function (req, res) {
    res.render('login.ejs');
})
router.get('/logout', function (req, res) {
    res.render('index.ejs');
})

module.exports = router