var socket = io.connect(`ws://${host}:${port}`, { transports: ['websocket'] });

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
        $('.total').html(curHealth + "/" + maxHealth +" HP");
    }


    socket.on('arenaLoop', function (data) {
        applyChange(data.body.h)
        curHealth = data.body.h
        document.getElementById("user-amount").innerHTML = data.body.n + " players with you."
        if (curHealth <= 0) {
            curHealth = 0
            player.loadVideoById('-YCN-a0NsNk')
        }
    });

    applyChange(curHealth);

    $(".health-bar").css({
        "width": "100%"
    });
    $(".add-damage").click(function() {
        if (curHealth > 0) {
            socket.emit('arenaHit', {
                body: {
                    token: token,
                },
            });
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

socket.emit('arenaConnect', {
    body: {
        id: ownID,
        token: token,
    },
});