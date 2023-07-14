window._gnative = {
    conn: null,
    callbacks: {},
}
function _gnativeCallGo(id, callback, args) {
    var s = '';
    if (args && args.length > 0) {
        s = ' ';
        for (var i = 0; i < args.length; i++) {
            s += encodeURIComponent(args[i])
            if (i < args.length - 1) {
                s += ' ';
            }
        }
    }
    if (callback) {
        window._gnative.callbacks[id] = callback;
    }
    window._gnative.conn.send('call ' + id + s)
}

window._gnative.conn = new WebSocket(location.origin.replace('http://', 'ws://').replace('https://', 'wss://') + '{{.}}')
_gnative.conn.onerror = function (e) {
    console.error(e);
}
_gnative.conn.onclose = function () {
    console.error('_gnative.conn closed');
}
_gnative.conn.onopen = function () {
    console.log('_gnative.conn open');
}
_gnative.conn.onmessage = function (e) {
    var ss = e.data.split(' ');
    if (ss.length < 2) {
        console.error('invalid message:' + e.data);
    }
    if (ss[0] === 'return') {
        // return <fid> data
        var id = parseInt(ss[1]);
        var c = window._gnative.callbacks[id];
        if (!c) {
            console.error('callback for return id:' + id + ' not found');
            return;
        }
        c(ss[2] ? decodeURIComponent(ss[2]) : null);
        return;
    }
    if (ss[0] === 'call') {
        // call <fid> jscode
        var id = parseInt(ss[1]);
        if (ss[2]) {
            var result = eval(decodeURIComponent(ss[2]));
            window._gnative.conn.send('return ' + id + ' ' + encodeURIComponent(JSON.stringify(result)));
        } else {
            window._gnative.conn.send('return ' + id);
        }
        return;
    }
    console.error('unsupported operation:' + ss[0]);
}