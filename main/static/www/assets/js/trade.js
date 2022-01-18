$(document).ready(function() {
    var owninventory = document.getElementById("owninventory")
    Sortable.create(owninventory, {
        animation: 100,
        group: 'list-1',
        draggable: '.list-group-item',
        handle: '.list-group-item',
        filter: '.sortable-disabled',
        chosenClass: 'active',
        sort: false,
        onEnd: () => {
            check()
        }
    });

    var ownoffer = document.getElementById("ownoffer")
    Sortable.create(ownoffer, {
        animation: 100,
        group: 'list-1',
        draggable: '.list-group-item',
        handle: '.list-group-item',
        filter: '.sortable-disabled',
        chosenClass: 'active',
        sort: false,
        onEnd: () => {
            check()
        }
    });

    var theirinventory = document.getElementById("theirinventory")
    Sortable.create(theirinventory, {
        animation: 100,
        group: 'list-2',
        draggable: '.list-group-item',
        handle: '.list-group-item',
        filter: '.sortable-disabled',
        chosenClass: 'active',
        sort: false,
        onEnd: () => {
            check()
        }
    });

    var theiroffer = document.getElementById("theiroffer")
    Sortable.create(theiroffer, {
        animation: 100,
        group: 'list-2',
        draggable: '.list-group-item',
        handle: '.list-group-item',
        filter: '.sortable-disabled',
        chosenClass: 'active',
        sort: false,
        onEnd: () => {
            check()
        }
    });

    function check() {
        elem = document.getElementById("temp-own-inv")
        if (owninventory.children.length == 1) {
            elem.style.cssText = 'display: block !important;'
        } else {
            elem.style.cssText = 'display: none !important;'
        }

        elem = document.getElementById("temp-own-offer")
        if (ownoffer.children.length == 1) {
            elem.style.cssText = 'display: block !important;'
        } else {
            elem.style.cssText = 'display: none !important;'
        }

        elem = document.getElementById("temp-their-inv")
        if (theirinventory.children.length == 1) {
            elem.style.cssText = 'display: block !important;'
        } else {
            elem.style.cssText = 'display: none !important;'
        }

        elem = document.getElementById("temp-their-offer")
        if (theiroffer.children.length == 1) {
            elem.style.cssText = 'display: block !important;'
        } else {
            elem.style.cssText = 'display: none !important;'
        }
    }

    check()

    var socket = io.connect('ws://localhost:9999', { transports: ['websocket'] });

    socket.emit('echo', { text: 'Hello world.' }, function(response) {
        console.log(response);
    });
})

var menuSelf = 0
var menuTheir = 0

function selfChange(id) {
    menuSelf = id
    parent = document.getElementById("yourNav")
    for (var i = 0; i < parent.children.length; i++) {
        if (i == menuSelf) {
            parent.children[i].classList.add("primary-color")
        } else {
            parent.children[i].classList.remove("primary-color")
        }
    }
}

function otherChange(id) {
    menuTheir = id
    parent = document.getElementById("theirNav")
    for (var i = 0; i < parent.children.length; i++) {
        if (i == menuTheir) {
            parent.children[i].classList.add("primary-color")
        } else {
            parent.children[i].classList.remove("primary-color")
        }
    }
}