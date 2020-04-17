let express = require("express");
let router = express.Router();
let data_worker = require("../data/data_worker")

router.post('/', (req, res, next) => {
    /* If request token is valid -> return cards list, else -> return error */
    let {token} = req.body;
    data_worker.get_cards(token).then(response => {
        res.send(response)
    });
})

module.exports = router;
