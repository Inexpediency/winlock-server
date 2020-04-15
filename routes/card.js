var express = require("express")
var router = express.Router()

class Card {
    constructor(number, cvv, date, owner) {
        this.number = number;
        this.cvv = cvv;
        this.date = date;
        this.owner = owner;
    }
    
    is_valid() {
        if (this.number.length > 3) {
            return true
        } else {
            return false
        }
    }
}

router.post('/',(req, res, next) => {
    let {number, cvv, date, owner} = req.body;
    let user_card = new Card(number, cvv, date, owner);
    if(user_card.is_valid()){
        res.send({
            ok: true,
            message: "successful"
        })
    } else {
        res.send({
            ok: false,
            message: "incorrect"
        })
    }
})

module.exports = router;
