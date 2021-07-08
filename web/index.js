
const express = require('express');
const path = require('path');
const app = express();

const routePage = require('./route/routePage');

const PORT = 3000;
const HOST = '0.0.0.0';

app.use(express.static('views'));
app.set('views', './views')
app.set('view engine', 'ejs')

// parsing request
app.use(express.urlencoded({extended:true}));
app.use(express.json());

// route page
app.use('/',routePage);

// start server
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);