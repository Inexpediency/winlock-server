var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
    res.render('index', { title: 'Winlock API Say BB ur PC bruh xd' });
});

module.exports = router;
