var socket = io.connect(`ws://${host}:9999`, { transports: ['websocket'] });

function SetEverything() {
    var tag = document.createElement('script');
    tag.src = "http://www.youtube.com/iframe_api";
    var firstScriptTag = document.getElementsByTagName('script')[0];
    firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);

    $ = jQuery;
    var ws;

    var applyChange = function(curHealth) {
        var a = curHealth * (100 / maxHealth);
        $(".health-bar-text").html("HP : " + Math.round(a) + "%");
        $(".health-bar").width(a + "%");
        // Prevents falling in negative numbers
        if (curHealth <= 0) {
            curHealth = 0
        }
        $('.total').html(curHealth + "/" + maxHealth);
    }

    if ("WebSocket" in window) {
        var ws = new WebSocket("wss://arena.rwbyadventures.com" + window.location.pathname + "ws" + document.location.search);
        ws.onopen = function() {};
        ws.onerror = function(err) {
            console.log(err);
        }

        ws.onmessage = function(evt) {
            var message = evt.data;
            var msg = JSON.parse(message);
            switch (msg["a"]) {
                case 'dmg':
                    applyChange(msg["h"])
                    curHealth = msg["h"]
                    document.getElementById("user").innerHTML = msg["n"] + " players"
                    if (curHealth <= 0) {
                        curHealth = 0
                        player.loadVideoById('-YCN-a0NsNk')
                    }
                    break;
            };
        };
    }

    applyChange(curHealth);

    $(".health-bar").css({
        "width": "100%"
    });
    $(".add-damage").click(function() {
        if (curHealth > 0) {
            ws.send('{"a":"dmg"}');
        }
    });
}

SetEverything()

function onYouTubeIframeAPIReady() {
    player = new YT.Player('player', {
        height: '0',
        width: '0',
        videoId: 'XyuY9sBE5f8',
        events: {
            'onReady': onPlayerReady,
            'onStateChange': onPlayerStateChange
        }
    });
}

function onPlayerReady(event) {
    event.target.playVideo();
    event.target.setVolume(2.5);
};

function onPlayerStateChange(event) {
    if (event.data != YT.PlayerState.PLAYING) {
        event.target.playVideo();
    }
}