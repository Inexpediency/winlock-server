const {promisify} = require("util");
const fs = require("fs");
const tokens = require("./tokens");

/* Worker with BD file */

class DataWorker {

    static async get_cards(token) {
        /* Get List of Cards from BD */
        if (tokens.indexOf(token) != -1) {
            const readFile = promisify(fs.readFile)
            let cards_list = readFile('./data/data.json', "utf8");
            return cards_list;
        }
        return 'Invalid token'
    }

    static async is_card_exist(card) {
        /* Check the Same Card in BD */
        let take_cards_list = async () => {
            let cards_list;
            try {
                cards_list = await DataWorker.get_cards('is_exist_check')
                cards_list = JSON.parse(cards_list);
            }
            catch(err) {
                console.log(Error(err))
            }
            return cards_list
        }
        let cards_list = await take_cards_list();

        let is_exist = false;
        // if cards list is empty check
        if (cards_list[0] != null) { 
            for (let id in cards_list) {
                if (card.equals(cards_list[id]))
                    is_exist = true
            }
        } else {
            return false
        }

        return is_exist
    }

    static async add_card(card) {
        /* Add Card to BD */
        let cards_list;
        try {
            cards_list = await DataWorker.get_cards('adding_task');
            cards_list = JSON.parse(cards_list);
        } 
        catch(err) {
            console.log(Error(err))
        }
        // if cards list is empty check
        if (cards_list[0] != null)
            cards_list.push(card)
        else 
            cards_list[0] = card
        cards_list = JSON.stringify(cards_list, null, '   ');

        fs.writeFile('./data/data.json', cards_list, () => {})
    }

}

module.exports = DataWorker
