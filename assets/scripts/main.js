var login = function() {
  var loginButton = document.querySelector('#loginButton');
  loginButton.addEventListener('click', function() {
    // var xhr = new XMLHttpRequest();
    // xhr.open("POST", "/");
    // xhr.send()
  });
};

var scream = function() {
  var screamButton = document.querySelector('#screamButton');
  screamButton.addEventListener('click', function() {
    var message = window.prompt("SCREAM HERE!");
    var date = new Date("2011-04-20 09:30:51.01");
    var time = String(date);
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/tweets");
    xhr.send(JSON.stringify({
      Message: message,
      Time: time
    }));
  });
};

var createTweets = function(tweetContent) {
  tweetList = document.querySelector('#list');
  tweets = document.createElement('li');
  tweets.setAttribute("class", "list-group-item");
  tweets.textContent = tweetContent;

  tweetList.appendChild(tweets);
};

var loopJSON = function(json) {
  for (var i = 0; i < json.length; i++) {
    createTweets(json[i].ID);
  }
};

var getTweets = function() {
  httpRequest = new XMLHttpRequest();
  httpRequest.open('GET', '/api/tweets', true);
  httpRequest.setRequestHeader("Content-Type",
    "application/json;charset=UTF-8");
  httpRequest.send(null);
  httpRequest.onreadystatechange = function(evt) {
    evt.preventDefault();
    if (httpRequest.readyState === 4) {
      if (httpRequest.status === 200) {
        var response = JSON.parse(httpRequest.responseText);
        loopJSON(response);
      } else {
        console.log('There was a problem with the request.');
      }
    }
  };
};

scream();
getTweets();
