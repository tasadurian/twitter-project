var scream = function() {
  var screamButton = document.querySelector('#screamButton');
  screamButton.addEventListener('click', function() {
    var message = window.prompt("SCREAM HERE!");
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/tweets");
    xhr.send(JSON.stringify({
      Message: message
    }));
    xhr.onreadystatechange = function(evt) {
      evt.preventDefault();
      if (xhr.status === 200 && xhr.readyState === 4) {
        location.reload();
      }
    };
  });
};

var createTweets = function(tweetContent, user, time) {
  var tweetList = document.querySelector('#list');
  var tweets = document.createElement('li');
  var userName = document.createElement('p');
  tweets.setAttribute("class", "list-group-item");
  var output = '<strong>' + user + '</strong>';
  output += '<em>' + time + '</em>';
  output += '<p>' + tweetContent + '</p>';
  tweets.innerHTML = output;
  tweetList.appendChild(tweets);
};

var loopJSON = function(json) {
  for (var i = 0; i < json.length; i++) {
    createTweets(json[i].Message, json[i].Username, json[i].Time);
  }
};

var pollNewTweets = function() {
  var originalLength = '';
  var httpRequest = new XMLHttpRequest();
  httpRequest.open('GET', '/api/tweets' + (username ? ("?username=" +
    username) : ""), true);
  httpRequest.setRequestHeader("Content-Type",
    "application/json;charset=UTF-8");
  httpRequest.send(null);
  http.onreadystatechange = function(evt) {
    evt.preventDefault();
    if ((httpRequest.readyState === 4) && (httpRequest.status === 200)) {
      var response = JSON.parse(httpRequest.responseText);
      originalLength = response.length;
    }
  };
  return originalLength;
};

var countPoll = function(poll) {
  var originalPoll = 0;
  var newPoll = '';
  window.setInterval(function() {
    newPoll = originalPoll + poll;
  }, 10000);
  return newPoll;
};

var displayNewPoll = function(newPoll) {
  var poller = document.querySelector('#poller');
  poller.textContent = newPoll;
};


var getTweets = function() {
  var username = location.pathname.split("/").pop();
  var httpRequest = new XMLHttpRequest();
  httpRequest.open('GET', '/api/tweets' + (username ? ("?username=" +
    username) : ""), true);
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

var followFunction = function() {
  var username = location.pathname.split("/").pop();
  var followButton = document.querySelector('#followButton');
  followButton.addEventListener('click', function() {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/api/follow');
    xhr.send(JSON.stringify(username));
    xhr.onreadystatechange = function(evt) {
      console.log('follow');
    };
  });
};

scream();
getTweets();
followFunction();
displayNewPoll();
