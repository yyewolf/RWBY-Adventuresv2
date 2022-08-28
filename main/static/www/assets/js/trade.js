var socket = io.connect(host, { transports: ['websocket'] });

var data = {
    elems: {
        ownInventory: document.getElementById("owninventory"),
        ownOffer: document.getElementById("ownoffer"),
        theirInventory: document.getElementById("theirinventory"),
        theirOffer: document.getElementById("theiroffer"),

        tempOwnInv: document.getElementById("temp-own-inv"),
        tempOwnOffer: document.getElementById("temp-own-offer"),
        tempTheirInv: document.getElementById("temp-their-inv"),
        tempTheirOffer: document.getElementById("temp-their-offer"),
    },

    cache: {
        ownCharacters: [],
        ownGrimms: [],
        ownMisc: [],

        theirCharacters: [],
        theirGrimms: [],
        theirMisc: [],
    }
}

$(document).ready(function() {

    Sortable.create(data.elems.ownInventory, {
        animation: 100,
        group: 'list-1',
        draggable: '.list-group-item',
        handle: '.list-group-item',
        filter: '.sortable-disabled',
        chosenClass: 'active',
        sort: false,
        onEnd: (evt) => {
            check();
            cache(evt);
        }
    });

    Sortable.create(data.elems.ownOffer, {
        animation: 100,
        group: 'list-1',
        draggable: '.list-group-item',
        handle: '.list-group-item',
        filter: '.sortable-disabled',
        chosenClass: 'active',
        sort: false,
        onEnd: (evt) => {
            check();
            cache(evt);
        }
    });

    Sortable.create(data.elems.theirInventory, {
        animation: 100,
        group: 'list-2',
        draggable: '.list-group-item',
        handle: '.list-group-item',
        filter: '.sortable-disabled',
        chosenClass: 'active',
        sort: false,
        onEnd: (evt) => {
            check();
            cache(evt);
        }
    });

    Sortable.create(data.elems.theirOffer, {
        animation: 100,
        group: 'list-2',
        draggable: '.list-group-item',
        handle: '.list-group-item',
        filter: '.sortable-disabled',
        chosenClass: 'active',
        sort: false,
        onEnd: (evt) => {
            check();
            cache(evt);
        }
    });

    function remFromArray(array, value) {
        var index = array.indexOf(value);
        if (index !== -1) {
            array.splice(index, 1);
        }
        return array
    }

    function cache(evt) {
        switch (evt.to.id) {
            case 'owninventory':
                switch (evt.item.type) {
                    case 'char':
                        data.cache.ownCharacters = remFromArray(data.cache.ownCharacters, evt.item.id);
                        break
                    case 'grimm':
                        data.cache.ownGrimms = remFromArray(data.cache.ownGrimms, evt.item.id);
                        break
                    case 'misc':
                        data.cache.ownMisc = remFromArray(data.cache.ownMisc, evt.item.id);
                        break
                }
                break
            case 'theirinventory':
                switch (evt.item.type) {
                    case 'char':
                        data.cache.theirCharacters = remFromArray(data.cache.theirCharacters, evt.item.id);
                        break
                    case 'grimm':
                        data.cache.theirGrimms = remFromArray(data.cache.theirGrimms, evt.item.id);
                        break
                    case 'misc':
                        data.cache.theirMisc = remFromArray(data.cache.theirMisc, evt.item.id);
                        break
                }
                break
            case 'ownoffer':
                switch (evt.item.type) {
                    case 'char':
                        data.cache.ownCharacters.push(evt.item.id)
                        break
                    case 'grimm':
                        data.cache.ownGrimms.push(evt.item.id)
                        break
                    case 'misc':
                        data.cache.ownMisc.push(evt.item.id)
                        break
                }
                break
            case 'theiroffer':
                switch (evt.item.type) {
                    case 'char':
                        data.cache.theirCharacters.push(evt.item.id)
                        break
                    case 'grimm':
                        data.cache.theirGrimms.push(evt.item.id)
                        break
                    case 'misc':
                        data.cache.theirMisc.push(evt.item.id)
                        break
                }
                break
        }

    }

    check()

    requestSomething("chars", true)
    requestSomething("chars", false)
})

var menuSelf = 0
var menuTheir = 0

function requestSomething(what, own) {
    id = otherID
    if (own) {
        id = ownID
    }
    socket.emit('tradeInfos', {
        body: {
            data: {
                action: what,
                target: id,
            },
            token: token,
        },
    }, (resp) => {
        handleInfo(resp, own)
        check();
    });
}

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
    switch (id) {
        case 0:
            requestSomething("chars", true)
            break
        case 1:
            requestSomething("grimms", true)
            break
        case 2:
            requestSomething("misc", true)
            break
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
    switch (id) {
        case 0:
            requestSomething("chars", false)
            break
        case 1:
            requestSomething("grimms", false)
            break
        case 2:
            requestSomething("misc", false)
            break
    }
}

function createPersona(perso) {
    /* <li class="d-flex list-group-item col-xs-4 bg-dark text-light m-2 align-items-center rounded-lg">
    <img height="50px" class="float-left rounded-circle mr-2" src="https://img.rwbyadventures.com/Penny_Polendina/Default_Icon.webp" />
    <span class="align-middle">Penny Polendina (Rare 50%)</span>
    </li> 
    */
    li = document.createElement('li');
    li.setAttribute('type', perso.type);
    li.setAttribute('id', perso.id);
    li.setAttribute('class', 'd-flex list-group-item col-xs-4 bg-dark text-light m-2 align-items-center rounded-lg');

    img = document.createElement('img');
    img.setAttribute('height', '50px');
    img.setAttribute('class', 'float-left rounded-circle mr-2');
    img.setAttribute('src', perso.imageurl);

    span = document.createElement('span');
    span.setAttribute('class', 'align-middle');
    span.innerText = `${perso.rarity} ${perso.name} (${perso.value}%)`;

    li.appendChild(img);
    li.appendChild(span);
    return li
}

function createMisc(misc) {
    /* <li class="d-flex list-group-item col-xs-4 bg-dark text-light m-2 align-items-center rounded-lg">
    <img height="50px" class="float-left rounded-circle mr-2" src="https://img.rwbyadventures.com/Penny_Polendina/Default_Icon.webp" />
    <input type="number" min="0" max="44" class="align-middle"/>
    </li> 
    */
    li = document.createElement('li');
    li.setAttribute('type', 'misc');
    li.setAttribute('id', misc.id);
    li.setAttribute('class', 'd-flex list-group-item col-xs-4 bg-dark text-light m-2 align-items-center rounded-lg');

    img = document.createElement('img');
    img.setAttribute('height', '50px');
    img.setAttribute('class', 'float-left rounded-circle mr-2');
    img.setAttribute('src', misc.imageurl);

    span = document.createElement('span');
    span.setAttribute('class', 'mr-3');
    span.innerText = `${misc.name} :`;

    input = document.createElement('input');
    input.setAttribute('type', 'number');
    input.setAttribute('min', '0');
    input.setAttribute('default', '0');
    input.setAttribute('max', misc.max);
    input.setAttribute('class', 'align-middle');

    li.appendChild(img)
    li.appendChild(span)
    li.appendChild(input)
    return li
}

function handleInfo(info, own) {
    console.log(info)
    info = info.body
    if (own) {
        data.elems.ownInventory.innerHTML = `
        <li id="temp-own-inv" class="sortable-disabled d-flex list-group-item col-xs-4 bg-dark text-light m-2 align-items-center rounded-lg">
            <div style="height:100px; width:100px; background-color: rgb(56, 56, 56); border-radius: 20px;" class="d-flex align-items-center text-center">
                <span class="vw-100 align-middle ">
                    Empty
                 </span>
            </div>
        </li>`
        info.characters.forEach(element => {
            e = createPersona(element);
            if (data.cache.ownCharacters.includes(e.id)) {
                return
            }
            data.elems.ownInventory.appendChild(e)
                //data.cache.ownCharacters.push(element.id)
        });
        info.grimms.forEach(element => {
            e = createPersona(element);
            if (data.cache.ownGrimms.includes(e.id)) {
                return
            }
            data.elems.ownInventory.appendChild(e)
                //data.cache.ownGrimms.push(element.id)
        });
        info.misc.forEach(element => {
            element.id = `${element.id}-own`
            e = createMisc(element);
            if (data.cache.ownMisc.includes(e.id)) {
                return
            }
            data.elems.ownInventory.appendChild(e)
                //data.cache.ownMisc.push(element.id)
        });
    } else {
        data.elems.theirInventory.innerHTML = `
        <li id="temp-their-inv" class="sortable-disabled d-flex list-group-item col-xs-4 bg-dark text-light m-2 align-items-center rounded-lg">
            <div style="height:100px; width:100px; background-color: rgb(56, 56, 56); border-radius: 20px;" class="d-flex align-items-center text-center">
                <span class="vw-100 align-middle ">
                    Empty
                 </span>
            </div>
        </li>`
        info.characters.forEach(element => {
            e = createPersona(element);
            if (data.cache.theirCharacters.includes(e.id)) {
                return
            }
            data.elems.theirInventory.appendChild(e)
                //data.cache.theirCharacters.push(element.id)
        });
        info.grimms.forEach(element => {
            e = createPersona(element);
            if (data.cache.theirGrimms.includes(e.id)) {
                return
            }
            data.elems.theirInventory.appendChild(e)
                //data.cache.theirGrimms.push(element.id)
        });
        info.misc.forEach(element => {
            element.id = `${element.id}-their`
            e = createMisc(element);
            if (data.cache.theirMisc.includes(e.id)) {
                return
            }
            data.elems.theirInventory.appendChild(e)
                //data.cache.theirMisc.push(element.id)
        });
    };
}

function check() {

    data.elems.tempOwnInv = document.getElementById("temp-own-inv");
    data.elems.tempTheirInv = document.getElementById("temp-their-inv");

    elem = data.elems.tempOwnInv
    if (data.elems.ownInventory.children.length == 1) {
        elem.style.cssText = 'display: block !important;'
    } else {
        elem.style.cssText = 'display: none !important;'
    }

    elem = data.elems.tempOwnOffer
    if (data.elems.ownOffer.children.length == 1) {
        elem.style.cssText = 'display: block !important;'
    } else {
        elem.style.cssText = 'display: none !important;'
    }

    elem = data.elems.tempTheirInv
    if (data.elems.theirInventory.children.length == 1) {
        elem.style.cssText = 'display: block !important;'
    } else {
        elem.style.cssText = 'display: none !important;'
    }

    elem = data.elems.tempTheirOffer
    if (data.elems.theirOffer.children.length == 1) {
        elem.style.cssText = 'display: block !important;'
    } else {
        elem.style.cssText = 'display: none !important;'
    }
}

function submit() {
    response = {
        own: {
            chars: data.cache.ownCharacters,
            grimms: data.cache.ownGrimms,
        },
        other: {
            chars: data.cache.theirCharacters,
            grimms: data.cache.theirGrimms,
        },
    }

    data.cache.ownMisc.forEach((elem) => {
        try {
            e = document.getElementById(elem)
            inputs = e.getElementsByTagName('input');
            input = inputs[0]
            id = e.id.replace("-own", "")
            response.own[id] = parseInt(input.value, 10)

        } catch (e) {}
    })

    data.cache.theirMisc.forEach((elem) => {
        try {
            e = document.getElementById(elem)
            inputs = e.getElementsByTagName('input');
            input = inputs[0]
            id = e.id.replace("-their", "")
            response.other[id] = parseInt(input.value, 10)

        } catch (e) {}
    })
    console.log(response)
    socket.emit('tradeValidate', {
        body: {
            data: response,
            token: token,
        },
    }, (resp) => {
        if (resp.success) {
            window.location = "/success"
        } else {
            alert(resp.text)
        }
    });
}