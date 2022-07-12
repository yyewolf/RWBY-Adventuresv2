var socket = io.connect(`ws://${host}:${port}`, { transports: ['websocket'] });

function SetEverything() {
    $ = jQuery;
    var ws;

}

SetEverything()

socket.emit('dungeonConnect', {
    body: {
        id: ownID,
        token: token,
    },
});