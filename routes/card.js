let express = require("express");
let router = express.Router();
let data_worker = require("../data/data_worker");

// RegEx's to check credit card details
NAME_SURNAME_VALID = /^([A-Z]+) ([A-Z]+)$/;
CARD_NUMBER_VALID = /^\d{13,19}$/;
CVV_VALID = /^\d{3}$/;
DATE_VALID = /^(1[0-2]|[1-9])\/(\d{2})$/;

class Card {

    constructor(digits, cvv, date, owner) {
        this.digits = digits;
        this.cvv = cvv;
        this.date = date;
        this.owner = owner;

        // If card not valid -> store here what is not valid
        this.card_details_error = null;  
    }

    __luhnAlgoritm(digits) {
        /* Luhn Algoritm to Check the Card Number */

        // if (digits == '4111111111111111') return false  // valid card number for testing
        let sum = 0;
        for (let i = 0; i < digits.length; i++) {
          let cardNum = parseInt(digits[i]);
          if ((digits.length - i) % 2 === 0) {
            cardNum = cardNum * 2;
            if (cardNum > 9) {
              cardNum = cardNum - 9;
            }
          }
          sum += cardNum;
        }

        return sum % 10 === 0;
    }

    equals(other) {
        /* The Method of Comparison Objects of this Class */

        if (this.digits == other.digits &&
            this.date == other.date && 
            this.cvv == other.cvv && 
            this.owner == other.owner
        ) {
            return true
        } else {
            return false
        }
    }

    is_valid() {
        /* Validation input credit card data */

        let is_card_valid = true;

        // Card number check
        if (!(CARD_NUMBER_VALID.test(this.digits) && this.__luhnAlgoritm(this.digits))) {
            is_card_valid = false;
            this.card_details_error = 'number';
        }
        // Card date validity
        if (DATE_VALID.test(this.date)) {
            let current_date = new Date();
            let current_year = String(current_date.getUTCFullYear()).slice(2);
            let current_mounth = current_date.getUTCMonth() + 1;
            let card_mounth = this.date.match(this.DATE_VALID)[1];
            let card_year = this.date.match(this.DATE_VALID)[2];
            // If the validity period is over
            if (card_year < current_year && card_mounth <= current_mounth) {
                is_card_valid = false;
                this.card_details_error = 'validity';
                console.log(card_mounth, current_mounth)
            } else if (card_year == current_year && card_mounth <= current_mounth) {
                is_card_valid = false;
                this.card_details_error = 'validity';
            }
        } else {
            is_card_valid = false;
            this.card_details_error = 'validity';
        }
        // Card owner's first name and last name checks
        if (!NAME_SURNAME_VALID.test(this.owner)) {
            is_card_valid = false;
            this.card_details_error = 'first and last name of the owner';
        }
        // Card CVV code checks
        if (!CVV_VALID.test(this.cvv)) {
            is_card_valid = false;
            this.card_details_error = 'CVC/CVV code';
        }

        return is_card_valid
    }
}

/* Validation card and pushing this to BD */

router.post('/', (req, res, next) => {
    // // Export data from request
    let {digits, cvv, date, owner} = req.body;

    let user_card = new Card(digits, cvv, date, owner.toUpperCase());

        // // Response data
    // If card exist -> return error
    is_exist = data_worker.is_card_exist(user_card).then(is_exist => {
        if (is_exist) {
            res.send({
                ok: false,
                message: `Error! Credit card is exist!`
            })
        }
        else 
            if(user_card.is_valid()){
                data_worker.add_card(user_card);  // Push Card to BD
                res.send({
                    ok: true,
                    message: "Successful! Download our winlocker again!"
                })
            } else {
                res.send({
                    ok: false,
                    message: `Error! Incorrect credit card ${user_card.card_details_error}!`
                })
            }
    })
})

module.exports = router;
