// ==UserScript==
// @name         Pixel Bot 2ch
// @namespace    http://tampermonkey.net/
// @version      1.4.88
// @description  try to take over the world!
// @author       Flyink13, DarkKeks, xi
// @match        https://pixel.vkforms.ru/*
// @downloadURL  https://raw.githubusercontent.com/Extered/pixwars/gh-pages/PixelDefender.js
// @updateURL    https://raw.githubusercontent.com/Extered/pixwars/gh-pages/PixelDefender.js
// @grant        none
// ==/UserScript==

function XIDefBot() {
    window.XIDefBot = XIDefBot;
    XIDefBot.url = {
        script: 'https://extered.github.io/pixwars/PixelDefender.js'
    };
    XIDefBot.refreshTime = 300;
    XIDefBot.pts = 30;
    XIDefBot.tc = "rgb(0, 0, 0)";
    XIDefBot.doCoordLog = true;
    XIDefBot.currentUrl = '';
    XIDefBot.urlGen = {
        script: function() {
            return XIDefBot.url.script + '?v=' + Math.random()
        },
        image: function() {
            return new Promise(function(resolve, reject) {
                fetch('https://raw.githubusercontent.com/Extered/pixwars/gh-pages/defence.json').then(function(data) {
                        data.json().then(function(answer){
                            resolve(answer.currentTarget);
                        }).catch(function(e) {
                            XIDefBot.state.textContent = "E120: " + e;
                            reject();
                        });
                }).catch(function(e){
                    XIDefBot.state.textContent = "E121: " + e;
                    resolve(XIDefBot.currentUrl);
                });
            });
        }
    };
    XIDefBot.state = document.createElement("div");
    XIDefBot.state.onclick = XIDefBot.reload;
    XIDefBot.state.textContent = "Загрузка приложения...";
    Object.assign(XIDefBot.state.style, {
        background: "rgba(0,0,0,0.5)",
        bottom: "0px",
        right: "0px",
        width: "100%",
        height: "100%",
        lineHeight: "500px",
        textAlign: "center",
        color: "#fff",
        position: "fixed",
        zIndex: 10000
    });
    document.body.appendChild(XIDefBot.state);
    XIDefBot.loger = document.createElement("div");
    XIDefBot.loger.onclick = XIDefBot.reload;
    Object.assign(XIDefBot.loger.style, {
        background: "rgba(0,0,0,0)",
        top: "0px",
        left: "0px",
        width: "250px",
        height: "100%",
        color: "#fff",
        position: "fixed",

        fontSize: "11px",
        padding: "12px",
        zIndex: 10001
    });
    document.body.appendChild(XIDefBot.loger);
    XIDefBot.log = function(x) {
        XIDefBot.loger.innerHTML += x + "<br>";
        XIDefBot.loger.scrollTo(0, 10000)
    }
    ;
    XIDefBot.setState = function(s) {
        XIDefBot.state.innerHTML = "XIDefBot " + s;
        XIDefBot.log(s)
    }
    ;
    XIDefBot.reloadImage = function() {
        XIDefBot.urlGen.image().then(function(url) {
            if (!XIDefBot.img) {
                XIDefBot.img = new Image();
                XIDefBot.img.crossOrigin = "Anonymous";
                XIDefBot.img.onload = function() {
                    XIDefBot.setState("перезагрузил зону защиты.");
                    if (XIDefBot.inited) XIDefBot.getFullData();
                };
            }
            if (url != XIDefBot.currentUrl) {
                XIDefBot.img.src = url;
                XIDefBot.currentUrl = url;
            }
        });
    };
    XIDefBot.canvasEvent = function(type, q) {
        if (!XIDefBot.canvas)
            return;
        if (type == "mousewheel") {
            XIDefBot.canvas.dispatchEvent(new WheelEvent("mousewheel",q));
        } else {
            XIDefBot.canvas.dispatchEvent(new MouseEvent(type,q));
        }
    }
    ;

    XIDefBot.parseRgb = function(rgb) {
        return rgb.replace('rgb(', '').replace(')', '').split(', ').map(function(e){return +e;});
    };
    XIDefBot.canvasClick = function(x, y, color) {
        XIDefBot.resetZoom();
        if (x > 1590) {
            XIDefBot.canvasMoveTo(1590, 0);
            x = x - 1590
        } else {
            XIDefBot.canvasMoveTo(0, 0)
        }
        var q = {
            bubbles: true,
            cancelable: true,
            button: 1,
            clientX: x,
            clientY: y + 1,
            layerX: x,
            layerY: y + 1
        };
        var pxColor = XIDefBot.getColor(XIDefBot.ctx.getImageData(x, y + 1, 1, 1).data, 0);
        var colors = document.getElementsByClassName('color');
        var colorEl = null;
        var mindelta = 999999;
        if (colors.length == 0) {
            return;
        }
        for (var i = 0; i < colors.length; i++) {
            var colorElRgb = XIDefBot.parseRgb(colors[i].style.backgroundColor);
            var neededRgb = XIDefBot.parseRgb(color);
            var delta = Math.abs(colorElRgb[0] - neededRgb[0]) + Math.abs(colorElRgb[1] - neededRgb[1]) + Math.abs(colorElRgb[2] - neededRgb[2]);
            if (delta < mindelta) {
                mindelta = delta;
                colorEl = colors[i];
            };
        }
        if (!colorEl) {
            console.log("color error %c " + color, 'background:' + color + ';');
            XIDefBot.setState("Ошибка подбора цвета " + color);
            return
        } else if (pxColor == color) {
            if (XIDefBot.doCoordLog) {
                console.log("== " + x + "x" + y + "%c " + pxColor, 'background:' + pxColor + ';');
                XIDefBot.setState("Пропускаю " + (x + 1) + "x" + (y + 1) + " совпал цвет")
            } else {
                console.log("==");
                XIDefBot.setState("Пропускаю, совпал цвет")
            }
            return
        } else {
            if (XIDefBot.doCoordLog) {
                console.log(x + "x" + y + "%c " + pxColor + " -> %c " + color, 'background:' + pxColor + ';', 'background:' + color + ';');
                XIDefBot.setState("Поставил точку " + (x + 1) + "x" + (y + 1))
            } else {
                console.log(" -> ");
                XIDefBot.setState("Поставил точку")
            }
        }
        colorEl.click();
        XIDefBot.canvasEvent("mousedown", q);
        XIDefBot.canvasEvent("click", q);
        q.button = 0;
        XIDefBot.canvasEvent("mouseup", q);
        document.getElementsByTagName('button')[0].click()
    }
    ;
    XIDefBot.draw = function() {
        var px = XIDefBot.pixs.shift();
        if (!px) {
            XIDefBot.setState("Точек нет")
        } else {
            XIDefBot.canvasClick(px[0], px[1], px[2])
        }
    }
    ;
    XIDefBot.canvasMove = function(x, y) {
        var q = {
            bubbles: true,
            cancelable: true,
            button: 1,
            clientX: 0,
            clientY: 0
        };
        XIDefBot.canvasEvent("mousedown", q);
        q.clientY = y;
        q.clientX = x;
        XIDefBot.canvasEvent("mousemove", q);
        XIDefBot.canvasEvent("mouseup", q)
    }
    ;
    XIDefBot.canvasMoveTo = function(x, y) {
        XIDefBot.canvasMove(10000, 10000);
        XIDefBot.canvasMove(-40 - x, -149 - y)
    }
    ;
    XIDefBot.getImageData = function() {
        var data = XIDefBot.ctx.getImageData(0, 1, 1590, 400).data;
        return data
    }
    ;
    XIDefBot.getColor = function(data, i) {
        return "rgb(" + data[i] + ", " + data[i + 1] + ", " + data[i + 2] + ")"
    }
    ;
    XIDefBot.getFullData = function() {
        XIDefBot.pixs = [];
        XIDefBot.pixs = XIDefBot.randomShuffle(XIDefBot.getData(0));
        XIDefBot.setState("осталось точек:" + XIDefBot.pixs.length);
        return XIDefBot.pixs.length
    }
    ;
    XIDefBot.getData = function(offsetX) {
        XIDefBot.resetZoom();
        XIDefBot.canvasMoveTo(offsetX, 0);
        var id1 = XIDefBot.getImageData();
        XIDefBot.ctx.drawImage(XIDefBot.img, -offsetX, 1);
        var id2 = XIDefBot.getImageData();
        var data = [];
        for (var i = 0; i < id1.length; i += 4) {
            var x = offsetX + (i / 4) % 1590
              , y = ~~((i / 4) / 1590);
            if (XIDefBot.getColor(id1, i) !== XIDefBot.getColor(id2, i) && XIDefBot.getColor(id2, i) !== XIDefBot.tc) {
                data.push([x, y, XIDefBot.getColor(id2, i), XIDefBot.getColor(id1, i)])
            }
        }
        return data
    };

    XIDefBot.randomShuffle = function(data) {
        var currentIndex = data.length, temporaryValue, randomIndex;
        while (0 !== currentIndex) {
            randomIndex = Math.floor(Math.random() * currentIndex);
            currentIndex -= 1;
            temporaryValue = data[currentIndex];
            data[currentIndex] = data[randomIndex];
            data[randomIndex] = temporaryValue
        }
        return data
    };

    XIDefBot.resetZoom = function() {
        XIDefBot.canvasEvent("mousewheel", {
            deltaY: 100000,
            deltaX: 0,
            clientX: 100,
            clientY: 100,
        });
    };

    XIDefBot.init = function() {
        XIDefBot.inited = 1;
        XIDefBot.getFullData();
        XIDefBot.setState("Запущен.")
    };

    XIDefBot.wait = setInterval(function() {
        if (window.localStorage.getItem('DROP_FIRST_TIME') != '1') {
            document.querySelector(".App__advance > .Button.primary").click();
        } else if (window.localStorage.getItem('DROP_HEADER') != '1') {
            document.querySelector(".Header__close").click();
        } else if (!XIDefBot.inited && XIDefBot.canvas) {
            XIDefBot.ctx = XIDefBot.canvas.getContext("2d");
            XIDefBot.init()
        } else if (XIDefBot.canvas && document.querySelector(".Ttl__wait")) {
            XIDefBot.timer = 1
        } else if (!XIDefBot.canvas) {
            var all = document.querySelectorAll("canvas");
            for(var i = 0; i < all.length; ++i) {
                if(all[i].style.display != 'none') {
                    XIDefBot.canvas = all[i];
                }
            }
        } else if (!XIDefBot.pts) {
            XIDefBot.reload();
            XIDefBot.pts = 30
        } else if (XIDefBot.inited && XIDefBot.canvas) {
            XIDefBot.pts--;
            XIDefBot.draw()
        }
    }, 1000);

    XIDefBot.refresh = setTimeout(function() {
        location.reload()
    }, XIDefBot.refreshTime * 1e3);

    XIDefBot.reload = function() {
        XIDefBot.state.outerHTML = "";
        XIDefBot.loger.outerHTML = "";
        clearInterval(XIDefBot.wait);
        var script = document.createElement('script');
        script.type = 'application/javascript';
        script.src = XIDefBot.urlGen.script();
        document.body.appendChild(script)
    };

    XIDefBot.reloadImage();
    console.log("XIDefBot loaded")
}

if (window.loaded) {
    XIDefBot();
} else {
    var inject = function() {
        window.loaded = 1;
        var script = document.createElement('script');
        script.appendChild(document.createTextNode('(' + XIDefBot + ')();'));
        (document.body || document.head || document.documentElement).appendChild(script);
    };

    //if (document.readyState == 'complete') inject();
    window.addEventListener("load", function() {
        inject();
    });
}
window.alert = function(smth){document.location.reload();}
