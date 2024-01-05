function Header2JSON(xhr){
    var headersObj = {};
    xhr.getAllResponseHeaders().trim().split('\r\n').forEach(function(header) {
        var parts = header.split(': ');
        var headerName = parts.shift();
        var headerValue = parts.join(': ');
        headersObj[headerName] = headerValue;
    });

    // return JSON.stringify(headersObj, null, 2)
    return JSON.parse(JSON.stringify(headersObj, null, 2))
}

function Cookies2JSON(cookie){
    var cookies = document.cookie.split(';');
    var cookieObject = {};

    for (var i = 0; i < cookies.length; i++) {
        var cookie = cookies[i].trim();
        var cookieParts = cookie.split('=');
        var cookieName = decodeURIComponent(cookieParts[0]);
        var cookieValue = decodeURIComponent(cookieParts.slice(1).join('='));

        cookieObject[cookieName] = cookieValue;
    }

    // return JSON.stringify(cookieObject, null, 2);
    return JSON.parse(JSON.stringify(cookieObject, null, 2))

}