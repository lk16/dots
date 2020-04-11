'use strict';
(function () {
  /**
   * @param {string} p
   * @param {string} c
   * @return {undefined}
   */
  function analyzeAll(p, c) {
    if (c) {
      var obj = _key;
      p = p.split(".");
      /** @type {number} */
      var i = 0;
      for (; i < p.length - 1; i++) {
        var prop = p[i];
        if (!(prop in obj)) {
          obj[prop] = {};
        }
        obj = obj[prop];
      }
      p = p[p.length - 1];
      i = obj[p];
      c = c(i);
      if (c != i && null != c) {
        defineProperty(obj, p, {
          configurable: true,
          writable: true,
          value: c
        });
      }
    }
  }
  /**
   * @param {!Object} observer
   * @return {undefined}
   */
  function init(observer) {
    var options = this;
    /** @type {string} */
    this.N = "";
    var data = this;
    /** @type {!Object} */
    this.o = observer;
    this.o.ve = this.o.ve || false;
    this.o.hd = this.o.hd || false;
    this.o.ob = this.o.ob || false;
    this.o.Ce = this.o.Ce || false;
    this.Mb = this.o.alt || null;
    this.hc = this.o.nf;
    /** @type {number} */
    this.Ve = Date.now();
    /** @type {null} */
    this.Cd = this.La = this.$b = this.Yb = this.yb = this.B = this.H = this.C = this.a = null;
    /** @type {boolean} */
    this.Qb = true;
    /** @type {null} */
    this.Bd = this.va = null;
    /** @type {number} */
    this.wa = 1064;
    /** @type {number} */
    this.hf = 600 / this.wa;
    /** @type {boolean} */
    this.Pd = this.Pb = this.Ub = false;
    /** @type {number} */
    this.xa = -1;
    /** @type {boolean} */
    this.Qa = this.Ab = this.S = this.P = this.oa = this.ua = this.Ka = false;
    /** @type {string} */
    this.Nd = this.zd = this.Bb = this.Jd = this.Bc = this.zc = "";
    /** @type {number} */
    this.Ga = this.fb = 0;
    /** @type {string} */
    this.zb = this.sd = "";
    /** @type {boolean} */
    this.wc = false;
    this.Ld = {};
    /** @type {boolean} */
    this.Ec = this.bc = this.cc = this.ec = this.qa = false;
    /** @type {string} */
    this.nc = "";
    /** @type {null} */
    this.Cc = this.W = null;
    this.text = {};
    this.Fc = {};
    this.ac = {};
    /** @type {boolean} */
    var d = false;
    /** @type {!Array} */
    var allConditionals = [];
    /**
     * @param {string} m
     * @param {!Object} b
     * @param {number} a
     * @return {?}
     */
    window.onerror = function (m, b, a) {
      if (!b || !a || d || -1 == b.indexOf("/j/")) {
        return false;
      }
      /** @type {boolean} */
      d = true;
      /** @type {string} */
      m = "ERR " + window.k2url.game + " " + b + ":" + a + " " + m;
      if (-1 == allConditionals.indexOf(m)) {
        allConditionals.push(m);
        /** @type {number} */
        b = Math.floor((Date.now() - data.Ve) / 1E3);
        /** @type {number} */
        a = Math.floor(b % 60);
        split("(" + Math.floor(b / 60) + ":" + Math.floor(a / 10) % 10 + "" + a % 10 + "s,v:" + window.location.hash.substring(1) + ") " + m);
      }
      return d = false;
    };
    this.Qd = window.k2lang || (window.navigator.language || "en").split("-")[0];
    /** @type {string} */
    this.Kd = "http:" == document.location.protocol ? "http:" : "https:";
    /** @type {string} */
    this.zc = document.domain ? document.domain : "pl" == this.Qd ? "www.kurnik.pl" : "www.playok.com";
    this.Bc = window.k2hconnect || this.zc.replace("www.", "x.");
    this.Jd = window.k2hassets || this.Bc;
    /** @type {string} */
    this.Nd = document.title || "";
    /** @type {string} */
    var ua = window.navigator.userAgent || "";
    /** @type {number} */
    var opener = Math.max(screen.width, screen.height);
    /** @type {number} */
    var m = Math.min(screen.width, screen.height);
    /** @type {boolean} */
    opener = /\(Macintosh;/.test(ua) && 4 * m == 3 * opener;
    /** @type {boolean} */
    this.G = /\(iPhone|\(iPod|\(iPad/.test(ua) || opener;
    /** @type {boolean} */
    this.Fd = -1 != ua.indexOf("Android");
    /** @type {boolean} */
    this.Nb = this.G || this.Fd;
    /** @type {boolean} */
    this.Gd = -1 != ua.indexOf("Mobile") && -1 == ua.indexOf("(iPad");
    /** @type {boolean} */
    this.kf = -1 != ua.indexOf("(Macintosh") && -1 != ua.indexOf("Safari/") && -1 == ua.indexOf("Chrome");
    /** @type {boolean} */
    this.$e = /\(iPhone/.test(ua);
    /** @type {boolean} */
    this.lf = -1 != ua.indexOf("Trident/");
    this.S = window.k2app || 0;
    if (window.k2beta && !opener) {
      /** @type {boolean} */
      this.P = true;
      if (-1 == window.k2beta.indexOf("/1")) {
        /** @type {boolean} */
        this.Pb = true;
      }
    }
    if (!(this.P && 0 != this.hc)) {
      /** @type {boolean} */
      this.Pd = true;
    }
    if (document.domain) {
      if (-1 != document.cookie.indexOf("kguest=1")) {
        /** @type {boolean} */
        this.Ub = true;
        /** @type {string} */
        document.cookie = "kguest=0;path=/";
      }
      if (-1 != (opener = document.cookie.indexOf("kroom="))) {
        /** @type {string} */
        this.Bb = document.cookie.substring(opener + 6);
        if (-1 != (opener = this.Bb.indexOf(";"))) {
          /** @type {string} */
          this.Bb = this.Bb.substring(0, opener);
        }
        /** @type {string} */
        document.cookie = "kroom=;path=/";
      }
    }
    if (0 != this.hc) {
      /** @type {boolean} */
      this.Qa = true;
    }
    if (this.S) {
      /** @type {boolean} */
      this.Ab = this.Qa = true;
    }
    $(document.getElementsByTagName("head")[0], "style", {
      type: "text/css"
    }, window.k2style || "");
    this.a = document.getElementById("appcont") || $(document.body, "div", {});
    callback(this.a, "k2base");
    if (this.P) {
      callback(this.a, "aleg");
      callback(this.a, "hvok");
    }
    if (this.G && !/ OS [5-9]_/.test(ua)) {
      callback(this.a, "iosfix");
    }
    callback(this.a, "devios", this.G);
    callback(this.a, "devmob", this.Gd);
    callback(this.a, "devtch", this.Nb);
    callback(this.a, "h100vh", observer.Df || true);
    /**
     * @return {?}
     */
    this.a.onselectstart = function () {
      return false;
    };
    if (this.G) {
      /**
       * @return {undefined}
       */
      this.a.ontouchstart = function () {
      };
    }
    this.H = document.getElementById("prehead") || $(this.a, "div");
    if (this.H != this.a.firstChild) {
      this.a.insertBefore(this.H, this.a.firstChild);
    }
    callback(this.H, "anav");
    callback(this.H, "usno");
    extend(this.H, {
      zIndex: cluezIndex
    });
    /**
     * @return {?}
     */
    this.H.ontouchmove = function () {
      return false;
    };
    this.B = document.getElementById("precont") || $(this.a, "div");
    callback(this.B, "acon");
    callback(this.B, "bsbb");
    if (this.Pb) {
      extend(this.a.parentNode, {
        overflow: "hidden"
      });
      callback(this.a, "asizing");
      /** @type {boolean} */
      var dark = false;
      this.Rd = $(this.a, "div", {
        className: "szpan"
      }, [$("button", {
        className: "bmax",
        onclick: function () {
          if (880 <= data.wa - 100) {
            data.wa -= 100;
          }
          data.na();
          this.blur();
          return false;
        }
      }, "-"), $("button", {
        className: "bmax",
        onclick: function () {
          data.wa += 100;
          data.na();
          this.blur();
          return false;
        }
      }, "+"), $("button", {
        className: "bmax",
        style: {
          background: "#444",
          color: "#888"
        },
        onclick: function () {
          extend(document.body, {
            background: dark ? "#fff" : "#222"
          });
          /** @type {boolean} */
          dark = !dark;
        }
      }, "\u00b7")]);
    }
    if ("#_=_" == window.location.hash) {
      /** @type {string} */
      window.location.hash = "";
    }
    /**
     * @return {undefined}
     */
    window.onhashchange = function () {
      _init(data);
    };
    /**
     * @param {!Event} event
     * @return {undefined}
     */
    window.onkeydown = function (event) {
      if (27 == event.keyCode && data.va && data.Qb) {
        success(data);
      }
    };
    if (window.cordova) {
      var StatusBar;
      if (this.G && (StatusBar = window.StatusBar)) {
        if (StatusBar.overlaysWebView) {
          StatusBar.overlaysWebView(false);
        }
        if (StatusBar.styleBlackOpaque) {
          StatusBar.styleBlackOpaque();
        }
        if (StatusBar.backgroundColorByHexString) {
          StatusBar.backgroundColorByHexString("#000");
        }
      }
      document.addEventListener("resume", function () {
        /** @type {number} */
        data.We = Date.now();
        if (data.Se) {
          find(data, true);
        }
      }, false);
      document.addEventListener("backbutton", function () {
        if (data.va && data.Qb) {
          success(data);
        } else {
          if (navigator.app) {
            navigator.app.backHistory();
          }
        }
      }, false);
    }
    /**
     * @return {undefined}
     */
    window.onresize = function () {
      if (options.fb) {
        /** @type {boolean} */
        options.Td = true;
      } else {
        /** @type {number} */
        var window_w = window.innerWidth;
        /** @type {number} */
        var inner_height = window.innerHeight;
        options.na();
        /** @type {boolean} */
        options.Td = false;
        /** @type {number} */
        options.fb = setTimeout(function () {
          /** @type {number} */
          options.fb = 0;
          if (options.Td || window_w != window.innerWidth || inner_height != window.innerHeight || options.G) {
            options.na();
          }
        }, 150);
      }
    };
    if (!this.P && "hidden" in document) {
      document.addEventListener("visibilitychange", function () {
        /** @type {number} */
        options.Id = document.hidden ? Date.now() : 0;
      }, false);
    }
    if (!this.o.ve) {
      /**
       * @param {string} params
       * @return {undefined}
       */
      var _initialize = function (params) {
        if (data.Sb) {
          clearTimeout(data.Sb);
        }
        /** @type {number} */
        data.Sb = setTimeout(function () {
          /** @type {number} */
          data.Sb = 0;
          if (data.ua) {
            data.send(params ? [ia] : [below_centered], null); // lk16:64 send signal lose focus s:absent
            // lk16:63 send signal gain focus s:absent
          }
        }, 200);
      };
      /** @type {number} */
      this.Sb = 0;
      if (this.$e) {
        document.addEventListener("visibilitychange", function () {
          if (options.ua) {
            options.send(document.hidden ? [below_centered] : [ia], null);
          }
        });
      } else {
        /**
         * @return {?}
         */
        window.onblur = function () {
          return _initialize(false);
        };
        /**
         * @return {?}
         */
        window.onfocus = function () {
          return _initialize(true);
        };
      }
    }
    /**
     * @return {undefined}
     */
    window.onbeforeunload = function () {
      /** @type {boolean} */
      options.Ad = true;
    };
    if (this.Ub && this.o.hd) {
      createHPipe(this);
    } else {
      this.Bd = new ready({
        host: this.Bc,
        ports: window.k2hcons || ["wss:17003", "wss:443", "https:443"],
        Xc: function (value) {
          if (!options.Ad) {
            if (value == encoding || value == or) {
              if (!options.Ka && options.o.hd) {
                createHPipe(options);
              } else {
                /** @type {boolean} */
                options.ua = false;
                if (options.Ka) {
                  options.se();
                }
                /** @type {number} */
                options.Se = Date.now();
                if (value == or && options.We > Date.now() - 2E3) {
                  find(options, true);
                } else {
                  generate(options, {
                    connect: true
                  });
                }
              }
            } else {
              if (value == runlist) {
                generate(options, {
                  sf: true
                });
              }
            }
          }
        },
        ee: function () {
          return !(options.Id && 6E4 < Date.now() - options.Id);
        },
        de: function (m, z) { // lk16:1713 send i:[1713] s:[ua data]
          /** @type {string} */
          var host = (window.ap ? "|" + window.ap : "") + (window.ge ? "|" + window.ge : "");
          /** @type {number} */
          var maxW = screen.width;
          /** @type {number} */
          var maxH = screen.height;
          return {
            J: [options.o.zf],
            O: [
              wrapped(options) + host,
              options.Qd,
              window.k2beta ? "b" : options.Gd ? "m" : "",
              options.Bb,
              window.navigator.userAgent || "", "/" + z + "/" + m,
              options.S ? "" : "w",
              (options.Fd && maxH <= maxW ? maxH + "x" + maxW : maxW + "x" + maxH) + " " + Math.round(100 * getUnderlineBackgroundPositionY()) / 100,
              "ref:" + window.location.href,
              "ver:" + window.k2ver + (options.S ? "/app" : "")
            ]
          };
        },
        Ye: function (a, b) {
          return options.xc(a, b);
        }
      });
    }
  }
  /**
   * @param {!Array} s
   * @param {string} name
   * @param {!Function} value
   * @return {undefined}
   */
  function expect(s, name, value) {
    if (!s.ac[name]) {
      /** @type {!Array} */
      s.ac[name] = [];
    }
    s.ac[name].push(value);
  }
  /**
   * @param {!Array} item
   * @param {string} name
   * @param {?} v
   * @return {undefined}
   */
  function cb(item, name, v) {
    (item.ac[name] || []).forEach(function (flatten) {
      if ("function" === typeof flatten) {
        flatten(v);
      } else {
        flatten.Bf(v);
      }
    }, item);
  }
  /**
   * @param {!Object} self
   * @return {undefined}
   */
  function display(self) {
    self.La = $(self.a, "div", {
      className: "noth",
      style: {
        display: "none",
        zIndex: default_zIndex,
        position: "fixed",
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        width: "100%",
        height: "100%",
        background: "rgba(0,0,0,0.4)"
      },
      onclick: function (event) {
        if (!self.Qb) {
          return false;
        }
        if (event.target != self.La && event.target != self.Cd) {
          return true;
        }
        success(self);
        return false;
      }
    }, self.Cd = $("div", {
      style: {
        display: "table-cell",
        verticalAlign: "middle",
        textAlign: "center"
      }
    }));
  }
  /**
   * @param {!Object} value
   * @param {string} md
   * @param {!Object} data
   * @param {boolean} opts
   * @return {?}
   */
  function create(value, md, data, opts) {
    if (!value.La) {
      display(value);
    }
    var self = {};
    self.D = $(value.Cd, "div", {
      className: "bsbb bs",
      style: {
        display: "none",
        background: "#fff",
        textAlign: "left",
        minWidth: "260px",
        padding: ".3em 0",
        borderRadius: "4px"
      }
    });
    if (data) {
      extend(self.D, data);
    }
    if (!(opts && opts.noclose)) {
      self.Qf = $(self.D, "button", {
        onclick: function () {
          return success(value);
        },
        style: {
          cssFloat: "right",
          width: "3.4em",
          height: "3.4em",
          margin: 0,
          padding: 0,
          color: "#ccc",
          background: "transparent",
          border: "none",
          fontWeight: "normal",
          cursor: "pointer"
        }
      }, "X");
    }
    self.title = $(self.D, "p", {
      className: "fb",
      style: {
        padding: "0 15px"
      }
    }, [md]);
    self.Ma = $(self.D, "div", opts && opts.nopad ? {} : {
      style: {
        padding: "0 15px"
      }
    });
    return self;
  }
  /**
   * @param {!Object} o
   * @param {!Object} options
   * @param {!Object} name
   * @param {?} d
   * @return {undefined}
   */
  function filter(o, options, name, d) {
    if (!o.va) {
      /** @type {boolean} */
      o.Qb = !d || !d.nocancel;
      if ("undefined" != typeof name && null !== name && options.title) {
        resolve(options.title, [name]);
      }
      extend(options.D, {
        display: "inline-block"
      });
      callback(o.La, "usno", !d || !d.okselect);
      extend(o.La, {
        display: "table"
      });
      /** @type {!Object} */
      o.va = options;
    }
  }
  /**
   * @param {!Object} o
   * @return {undefined}
   */
  function success(o) {
    if (o.va) {
      extend(o.La, {
        display: "none"
      });
      extend(o.va.D, {
        display: "none"
      });
      /** @type {null} */
      o.va = null;
    }
  }
  /**
   * @param {?} result
   * @return {?}
   */
  function wrapped(result) {
    if (result.Ub && !result.Qa) {
      return "guest";
    }
    /** @type {null} */
    var message = null;
    /** @type {null} */
    var default_favicon = null;
    if (result.Qa) {
      try {
        /** @type {(null|string)} */
        default_favicon = window.localStorage.getItem("autoid");
        /** @type {(null|string)} */
        message = window.localStorage.getItem("ksession");
      } catch (k) {
      }
    }
    if (!message && document.domain) {
      /** @type {!Array<string>} */
      var results = (document.cookie || "").split(";");
      /** @type {number} */
      var i = 0;
      for (; i < results.length; i++) {
        /** @type {string} */
        var row = results[i];
        if (" " == row[0]) {
          /** @type {string} */
          row = row.substring(1);
        }
        if ("ksession=" == row.substring(0, 9)) {
          /** @type {string} */
          message = row.substring(9).split(":")[0];
          break;
        }
      }
    }
    return (message || "") + (result.Qa ? default_favicon && "+" == default_favicon[0] ? default_favicon.toString() : "+" : "");
  }
  /**
   * @param {!Object} event
   * @return {undefined}
   */
  function logout(event) {
    /** @type {null} */
    var hotkey = null;
    try {
      /** @type {(null|string)} */
      hotkey = window.localStorage.getItem("ksession");
    } catch (c) {
    }
    if (event.Qa && hotkey) {
      try {
        window.localStorage.removeItem("ksession");
      } catch (c) {
      }
      find(event, true);
    } else {
      /** @type {string} */
      document.body.innerHTML = "";
      /** @type {string} */
      window.location = get(event, "logout") + "?t=" + window.k2url.game;
    }
  }
  /**
   * @param {!Object} value
   * @param {boolean} variation
   * @return {undefined}
   */
  function find(value, variation) {
    if (variation) {
      /** @type {null} */
      window.onhashchange = null;
      window.location.hash = value.N;
    }
    /** @type {string} */
    document.body.innerHTML = "";
    if (value.Lf && window.AdMob.destroyBannerView) {
      window.AdMob.destroyBannerView();
    }
    window.location.reload();
  }
  /**
   * @param {!Object} data
   * @param {!Object} opts
   * @return {undefined}
   */
  function generate(data, opts) {
    if (!data.sc) {
      text(data, "status");
    }
    resolve(data.sc.f, $("table", $("td", [opts.sf ? $("p", $("div", {
      className: "loader"
    })) : null, opts.gd ? $("p", {
      className: "fb"
    }, opts.gd) : null, opts.connect ? $("p", $("button", {
      className: "minwd ttup",
      onclick: function () {
        find(data);
      }
    }, data.j("t_recn", "connect"))) : null, opts.link ? $("p", {
      className: "ttup"
    }, $("a", opts.link, opts.mf || "-")) : null])));
    _init(data);
  }
  /**
   * @param {!Object} options
   * @param {string} type
   * @return {?}
   */
  function get(options, type) {
    /** @type {boolean} */
    var isFriend = "pl" == options.lang;
    /** @type {string} */
    var QUERY_PREFIX = options.Kd + "//" + options.zc;
    /** @type {string} */
    var href = "/" + options.lang;
    /** @type {string} */
    var g = "&r=" + Date.now();
    switch (type) {
      case "login":
        return QUERY_PREFIX + (isFriend ? "/login.phtml" : href + "/login.phtml") + "?js=1" + g;
      case "register":
        return QUERY_PREFIX + (isFriend ? "/rejestracja.phtml" : href + "/register.phtml") + "?js=1" + g;
      case "profile":
        return QUERY_PREFIX + (isFriend ? "/prof.phtml" : href + "/prof.phtml") + "?js" + g;
      case "tourns":
        return QUERY_PREFIX + (isFriend ? "/turnieje/" : href + "/tournaments/") + "?js&on=" + options.zb + g;
      case "newtourn":
        return QUERY_PREFIX + (isFriend ? "/turnieje/nowy.phtml" : href + "/tournaments/new.phtml") + "?js&gid=" + options.zb + g;
      case "feedback":
        return QUERY_PREFIX + (isFriend ? "/uwagi/" : href + "/feedback/");
      case "passwd":
        return QUERY_PREFIX + (isFriend ? "/haslo.phtml" : href + "/pass.phtml");
      case "logout":
        return QUERY_PREFIX + (isFriend ? "/logout.phtml" : href + "/logout.phtml");
      case "stat":
        return QUERY_PREFIX + (isFriend ? "/stat.phtml?u=%s&g=" + options.zb : href + "/stat.phtml?u=%s&g=" + options.zb);
    }
    return "";
  }
  /**
   * @param {!Object} params
   * @return {undefined}
   */
  function clear(params) {
    params.send([id], ["noi=" + (params.ec ? "1" : "0") + "&nop=" + (params.cc ? "1" : "0") + "&prb=" + (params.bc ? "1" : "0") + "&snd=" + (params.qa ? "1" : "0") + (params.Mb ? "&" + params.Mb + "=" + (params.Ec ? "1" : "0") : "") + params.nc]);
  }
  /**
   * @param {!Object} args
   * @return {undefined}
   */
  function createHPipe(args) {
    if (!args.Ka) {
      args.Ic();
      /** @type {boolean} */
      args.Ka = true;
    }
    _init(args);
  }
  /**
   * @param {number} v
   * @return {?}
   */
  function val(v) {
    return 0 != v ? "#" + v : "";
  }
  /**
   * @param {!Object} a
   * @param {number} v
   * @return {?}
   */
  function fetch(a, v) {
    if (a.o.Nf) {
      return v.toString();
    }
    if (0 > v) {
      return "";
    }
    /** @type {number} */
    a = Math.floor(v / 1E3);
    return 0 == a && 0 == v % 1E3 ? "?" : Math.floor(5 * a / 10) + "." + 5 * a % 10;
  }
  /**
   * @param {number} b
   * @param {number} a
   * @return {?}
   */
  function exec(b, a) {
    if (0 == a) {
      return "-";
    }
    if (!b.o.Ce) {
      return a + "";
    }
    /** @type {number} */
    b = a - 15E3;
    /** @type {number} */
    a = 0 > b ? -b : b;
    return (0 > b ? "\u2013" : "+") + Math.floor(a / 100) + "." + Math.floor(a / 10) % 10 + a % 10;
  }
  /**
   * @param {!Object} a
   * @param {number} b
   * @return {?}
   */
  function join(a, b) {
    return a.W ? 1 + a.W[b] : 0;
  }
  /**
   * @param {!Object} result
   * @param {!Object} end
   * @return {?}
   */
  function merge(result, end) {
    if (!result.W) {
      return 0;
    }
    /** @type {number} */
    var j = 0;
    for (; j + 1 < result.W.length && !(end <= result.W[j]);) {
      j++;
    }
    return j;
  }
  /**
   * @param {!Object} data
   * @param {string} s
   * @param {boolean} str
   * @return {undefined}
   */
  function debug(data, s, str) {
    /** @type {string} */
    var diffUnnormalized = window.location.hash.toString();
    if ("-1" == s) {
      window.history.back();
    } else {
      if (s != diffUnnormalized.substring(1)) {
        /** @type {string} */
        window.location.hash = s;
      }
    }
    if (str) {
      _init(data);
    }
  }
  /**
   * @return {?}
   */
  function proceed() {
    /** @type {string} */
    var result = window.location.hash.substring(1);
    var b;
    if (-1 != (b = result.indexOf("/"))) {
      /** @type {string} */
      result = result.substring(0, b);
    }
    return result;
  }
  /**
   * @return {?}
   */
  function y() {
    /** @type {string} */
    var str = window.location.hash.substring(1);
    var pt;
    return -1 != (pt = str.indexOf("/")) ? str.substring(pt + 1) : null;
  }
  /**
   * @param {!Object} params
   * @param {string} name
   * @return {?}
   */
  function text(params, name) {
    if (!params.Fc[name]) {
      params.Fc[name] = params.cd(name);
    }
    return params.Fc[name];
  }
  /**
   * @param {!Object} item
   * @return {undefined}
   */
  function _init(item) {
    if (-1 == item.xa) {
      setTimeout(function () {
        return item.na();
      }, 0);
    } else {
      var total = item.re(proceed());
      var opts = text(item, item.ua ? total : "status");
      if (opts) {
        var state = opts.f;
        /** @type {boolean} */
        var _contextIsSetGet = item.C == state;
        if ("function" === typeof opts.Db) {
          opts.Db(!_contextIsSetGet);
        }
        if (item.C && item.C != state) {
          extend(item.C, {
            display: "none"
          });
        }
        if (item.Pb) {
          var data = item.Yd;
          extend(item.a, {
            maxWidth: data.Af + "px",
            height: (opts.yd ? data.gf : data.ff) + "px",
            minHeight: 0,
            top: data.top + "px"
          });
          callback(item.a, "dosize", opts.yf && 0 < data.top || false);
        }
        callback(item.a, "donav", !opts.yd);
        if ("string" == typeof item.o.ef) {
          callback(item.a, item.o.ef, !!opts.xf);
        }
        if (!opts.xd) {
          callback(item.a, "doddmenu", "undefined" === typeof opts.xd ? !item.P || 2 > item.xa : false);
        }
        data = 2 <= item.xa && item.P ? item.$b : "function" === typeof opts.Le ? opts.Le() : null;
        /** @type {boolean} */
        var vert = false;
        if (!data && item.Yb && opts.Ja) {
          data = item.Yb;
          item.Yb.innerHTML = opts.Ja;
          /** @type {boolean} */
          vert = true;
        }
        if (item.yb && item.yb != data) {
          extend(item.yb, {
            display: "none"
          });
        }
        if (item.yb != data && (item.yb = data)) {
          extend(data, {
            display: vert ? "block" : "inline-block"
          });
        }
        item.C = state;
        cb(item, "nav", total);
        if (!_contextIsSetGet && (extend(state, {
          display: "block"
        }), item.oa || ((document.scrollingElement || document.documentElement).scrollTop = 0), "function" === typeof opts.onshow)) {
          opts.onshow();
        }
        item.yc(!_contextIsSetGet);
      }
    }
  }
  /**
   * @param {?} callback
   * @return {undefined}
   */
  function move(callback) {
    init.call(this, callback);
    /** @type {string} */
    this.Pe = "f";
    /** @type {string} */
    this.Ne = "p";
    /** @type {string} */
    this.Re = "i";
    /** @type {string} */
    this.wd = "t";
    /** @type {string} */
    this.Qe = "n";
    /** @type {null} */
    this.M = this.ba = this.g = this.tab = null;
    this.A = {};
    /** @type {!Array} */
    this.vb = [];
    this.b = {};
    this.I = {};
    this.pa = {};
    /** @type {!Array} */
    this.eb = [];
    /** @type {!Array} */
    this.Od = [];
    /** @type {null} */
    this.Ed = this.Hd = this.Md = null;
    /** @type {number} */
    this.Sd = this.L = this.lc = this.mc = this.Fa = this.gc = 0;
    /** @type {null} */
    this.Lb = null;
    /** @type {number} */
    this.Kb = 0;
  }
  /**
   * @param {!Object} self
   * @param {?} name
   * @return {undefined}
   */
  function config(self, name) {
    if (!self.Gc) {
      self.Gc = create(self, "-");
    }
    resolve(self.Gc.Ma, $("p", [$("button", {
      className: "minw",
      onclick: function () {
        success(self);
        self.send([event, 0], [name]);
        return false;
      }
    }, self.j("tu_bcan")), $("button", {
      className: "minw",
      onclick: function () {
        success(self);
        self.send([workerId], ["/boot " + name]);
        return false;
      }
    }, self.j("bl_boot"))]));
    filter(self, self.Gc, name);
  }
  /**
   * @param {!Object} self
   * @param {?} name
   * @param {number} fmt
   * @return {undefined}
   */
  function compile(self, name, fmt) {
    if (!self.Dc) {
      self.Dc = create(self, "-");
    }
    resolve(self.Dc.Ma, [$("p", [$("button", {
      className: "minw",
      onclick: function () {
        success(self);
        self.send([action], [name]);
        return false;
      }
    }, "info"), $("button", {
      className: "minw",
      disabled: 0 != fmt,
      onclick: function () {
        success(self);
        self.send([event, 1], [name]);
        return false;
      }
    }, self.j("tu_bacc"))]), $("p", [$("button", {
      className: "minw",
      disabled: !(0 <= fmt),
      onclick: function () {
        success(self);
        self.send([event, 0], [name]);
        return false;
      }
    }, self.j("tu_bcan")), $("button", {
      className: "minw",
      onclick: function () {
        success(self);
        self.send([workerId], ["/boot " + name]);
        return false;
      }
    }, self.j("bl_boot"))])]);
    filter(self, self.Dc, name);
  }
  /**
   * @param {!Object} a
   * @return {undefined}
   */
  function emit(a) {
    /** @type {string} */
    document.title = a.Nd + (a.Lb ? " " + a.Lb : 0 < a.Kb ? " (" + a.Kb + ")" : "");
  }
  /**
   * @param {!Object} a
   * @param {number} b
   * @return {undefined}
   */
  function reduce(a, b) {
    if (a.Kb != b) {
      /** @type {number} */
      a.Kb = b;
      emit(a);
    }
  }
  /**
   * @param {!Object} msg
   * @return {undefined}
   */
  function then(msg) {
    debug(msg, msg.tab && msg.tab.ue ? "-1" : msg.N);
  }
  /**
   * @param {!Object} res
   * @param {number} name
   * @param {string} data
   * @return {undefined}
   */
  function server(res, name, data) {
    res.send(data ? [overwrite, name, 1] : [overwrite, name], null); // lk16:72 send join a room
  }
  /**
   * @param {?} query
   * @param {?} index
   * @return {?}
   */
  function process(query, index) {
    return query.A.hasOwnProperty(index) ? query.A[index] : null;
  }
  /**
   * @param {?} opts
   * @param {!Object} tx
   * @return {undefined}
   */
  function logger(opts, tx) {
    func(tx, opts.A);
    opts.vb.push(tx);
  }
  /**
   * @param {!Array} self
   * @param {string} id
   * @param {string} text
   * @return {undefined}
   */
  function unwrap(self, id, text) {
    if ("set_tournaments" == id) {
      /** @type {number} */
      id = parseInt(text, 10);
      if (id != self.lc) {
        /** @type {string} */
        self.lc = id;
        cb(self, "tumode", self.lc);
      }
    } else {
      if ("set_cols" == id) {
        /** @type {number} */
        self.mc = parseInt(text, 10);
      } else {
        if ("set_chat" == id) {
          cb(self, "chmode", "0" != text);
        }
      }
    }
  }
  /**
   * @param {!Object} self
   * @param {?} name
   * @param {!Array} list
   * @param {?} result
   * @param {!Object} input
   * @param {?} text
   * @return {undefined}
   */
  function doSearch(self, name, list, result, input, text) {
    var messages = self.g.M;
    var selector = input && 0 < input.length;
    if (messages) {
      resolve(messages, $("div", {
        className: "alrt dcpd"
      }, [$("div", {
        className: "mbsp"
      }, name), $("button", {
        className: "minw",
        onclick: function () {
          if (list) {
            self.send(list, result);
          }
          do_export(self);
        }
      }, self.j(selector ? "bl_yes" : "bl_ok")), " ", selector ? $("button", {
        className: "minw",
        onclick: function () {
          self.send(input || [], text);
          do_export(self);
        }
      }, self.j("bl_no")) : null]));
    }
  }
  /**
   * @param {!Object} item
   * @return {undefined}
   */
  function do_export(item) {
    if (item = item.g.M) {
      resolve(item, []);
    }
  }
  /**
   * @param {?} component
   * @return {undefined}
   */
  function drawTextCss(component) {
    Object.keys(component.A).forEach(function (i) {
      attr(component.A[i]);
      delete component.A[i];
    }, component);
  }
  /**
   * @param {!Arguments} keys
   * @return {undefined}
   */
  function rebuildModelFromFields(keys) {
    /** @type {number} */
    var results = 0;
    Object.keys(keys.I).forEach(function (i) {
      if (0 != keys.I[i].Oa) {
        results++;
      }
    }, keys);
    cb(keys, "nmessages", results);
    cb(keys, "chatalert", 0 < results);
    reduce(keys, results);
  }
  /**
   * @param {!Object} options
   * @param {string} index
   * @return {undefined}
   */
  function each(options, index) {
    if (!options.M) {
      text(options, "m");
    }
    request(options.M, index);
    debug(options, "m" + (options.oa ? "" : "/" + index));
  }
  /**
   * @param {!Object} options
   * @param {string} id
   * @param {!Array} s
   * @param {!Object} result
   * @return {undefined}
   */
  function initialize(options, id, s, result) {
    if (!(2 > s.length)) {
      if (!options.qc) {
        options.qc = create(options, "-");
        /**
         * @param {!Event} event
         * @return {?}
         */
        options.qc.D.onselectstart = function (event) {
          event.stopPropagation();
          return true;
        };
      }
      /** @type {number} */
      var j = 5 + 2 * s[4];
      var messageKey = 1 + s[4];
      var t;
      /** @type {!Array} */
      var newNodeLists = [];
      /** @type {number} */
      var index = 0;
      for (; index < Math.min(s[4], 2); index++) {
        var b = (0 > s[5 + 2 * index] ? 65536 + s[5 + 2 * index] : s[5 + 2 * index]) + 65536 * s[2 * index + 6];
        newNodeLists.push(result[1 + index] + " " + (0 < index ? b.toString() : exec(options, b)));
        if (0 == index) {
          t = b;
        }
      }
      /** @type {string} */
      var a = b = "";
      /** @type {!Array} */
      var data = ["-", "", "", ""];
      /** @type {number} */
      index = 0;
      for (; index < s.length - j; index++) {
        var value = result[messageKey + index];
        switch (s[j + index]) {
          case 0:
            data[0] = value;
            break;
          case 2:
            data[1] = value;
            break;
          case 3:
            data[2] = value;
            break;
          case 4:
            data[3] = value;
            break;
          case 5:
            b = value;
            break;
          case 6:
            a = value;
        }
      }
      result = a || "photos/none.jpg";
      /** @type {string} */
      j = data[0] + ("" != data[1] && "-" != data[1] ? " (" + data[1] + ")" : "");
      /** @type {string} */
      data = data[2] + ("" != data[3] ? ("" != data[2] ? ", " : "") + data[3] : "");
      /** @type {boolean} */
      var checked = 0 != (s[1] & DIRECTION_HORIZONTAL);
      /** @type {boolean} */
      var status = 0 != (s[1] & arg);
      var v = s[2];
      var inputEl;
      var entry;
      resolve(options.qc.Ma, ["undefined" != typeof t ? $("p", {
        className: "fb mbh",
        style: {
          marginTop: "-0.25em"
        }
      }, [newNodeLists[0], " ", $("div", {
        className: "r" + merge(options, t)
      }), $("br"), newNodeLists[1]]) : null, s = $("p", {
        style: {
          margin: ".5em 0",
          width: "220px",
          padding: ".2em 0",
          wordWrap: "break-word",
          overflowY: "auto"
        }
      }, [$("div", {
        style: {
          cssFloat: "left",
          marginTop: "3px",
          marginRight: ".5em",
          width: "52px",
          height: "52px",
          border: "solid 1px #aaa",
          overflow: "hidden"
        }
      }, $("img", {
        src: options.Kd + "//" + options.Jd + (result && "/" != result[0] ? "/" : "") + result
      })), $("div", j), $("div", data), $("div", b)]), options.S ? null : $("p", {
        className: "mbh"
      }, $("a", {
        className: "lbut minwd",
        target: "_blank",
        href: toString(get(options, "stat"), encodeURIComponent(id)),
        onclick: function () {
          success(options);
        }
      }, options.j("ui_stats"))), $("p", {
        className: "mbh"
      }, [$("button", {
        className: "minw",
        disabled: 0 == v,
        onclick: function () {
          success(options);
          server(options, v);
        }
      }, 0 < v ? "#" + v : "#000"), 0 == options.Ga ? $("a", {
        href: "",
        title: options.j("bl_whisper"),
        className: "spbb",
        onclick: function () {
          success(options);
          each(options, id);
          return false;
        },
        style: {
          margin: "0 1em"
        }
      }) : null]), $("div", {
        className: "dtline"
      }), 0 == options.Ga ? $("p", [inputEl = $("input", {
        type: "checkbox",
        checked: checked
      }), options.j("bl_buds"), " ", entry = $("input", {
        type: "checkbox",
        checked: status
      }), options.j("ui_block")]) : null, $("p", $("button", {
        className: "minw",
        onclick: function () {
          success(options);
          if (inputEl && inputEl.checked != checked) {
            options.X((inputEl.checked ? "/buddy" : "/unbuddy") + " " + id);
          }
          if (entry && entry.checked != status) {
            options.X((entry.checked ? "/ignore" : "/unignore") + " " + id);
          }
          return false;
        }
      }, options.j("bl_ok")))]);
      write(s);
      filter(options, options.qc, id, {
        okselect: true
      });
    }
  }
  /**
   * @return {undefined}
   */
  function addRoomButton() {
    var options = {
      uf: e
    };
    var b = document.getElementById("appcont") || $(document.body, "div", {
      id: "appcont"
    });
    options = new (options.uf || Object.constructor)(options.Pf || {});
    resolve(b, options.f);
    options.start();
  }
  /**
   * @param {!Object} options
   * @return {undefined}
   */
  function show(options) {
    var c = this;
    /** @type {!Object} */
    this.Sc = options;
    this.app = {
      af: function () {
        return 0;
      },
      Yd: {
        top: 1
      }
    };
    /** @type {string} */
    var u = window.navigator.userAgent || "";
    /** @type {boolean} */
    var t = /\(iPhone|\(iPod|\(iPad/.test(u);
    /** @type {boolean} */
    u = -1 != u.indexOf("Android");
    /** @type {boolean} */
    this.app.Nb = t || u;
    /** @type {boolean} */
    this.app.P = !this.app.Nb;
    /** @type {!Array} */
    this.a = [];
    /** @type {number} */
    this.Y = 0;
    /** @type {number} */
    this.A = -1;
    /** @type {string} */
    this.b = "";
    /** @type {boolean} */
    this.I = this.C = false;
    /** @type {!Array} */
    this.B = [];
    $(document.getElementsByTagName("head")[0], "style", {
      type: "text/css"
    }, "button { font: inherit; background: none; color: inherit; border-width: 1px; padding: .3em; outline: 0; } button { touch-action: manipulation; } a { text-decoration: none; color: inherit; } .usno { user-select: none; -webkit-user-select: none; }.noth { -webkit-tap-highlight-color: rgba(0,0,0,0); }.bsbb { box-sizing: border-box; } .bcont { position: absolute; left: 0; top: 0; width: 100%; height: 100%; } .rbcont { position: relative; }\t.rbratio { padding-top: 95%; } .rmcont { min-height: 5em; } .rmcont { word-wrap: break-word; } .rmcont span { cursor: pointer; } @media (min-width: 570px) { \t.rbcont { margin-right: 32%; }\t.rbratio { padding-top: 82%; } \t.rmcont { padding: 0 10px; overflow-y: auto; position: absolute; width: 32%; right: 0; top: 0; bottom: 0; box-sizing: border-box; }   .rmcbg { background: rgba(0,0,0,.02); }}");
    var values = this;
    this.f = $("div", {
      tabIndex: "0",
      style: {
        position: "relative",
        outline: 0
      }
    }, [t = $("div", {
      className: "rbcont"
    }, [$("div", {
      className: "rbratio"
    })]), $("div", {
      className: "rmcont rmcbg"
    }, [$("p", [$("button", {
      onclick: function () {
        if (0 < values.Y) {
          require(values, values.Y - 1);
        }
      },
      style: {
        minWidth: "4em"
      }
    }, "<"), " ", $("button", {
      onclick: function () {
        if (values.Y < values.a.length - 1) {
          require(values, values.Y + 1);
        }
      },
      style: {
        minWidth: "4em"
      }
    }, ">"), " ", this.M = $("span", {
      style: {
        whiteSpace: "nowrap"
      }
    })]), this.L = $("p", ["-"]), this.N = $("p", {}), this.H = $("p", {})])]);
    (window != window.top ? this.f : window).addEventListener("keydown", function (event) {
      switch (event.keyCode) {
        case 37:
        case 38:
          if (0 < c.Y) {
            require(c, c.Y - 1);
          }
          break;
        case 39:
        case 40:
          if (c.Y < c.a.length - 1) {
            require(c, c.Y + 1);
          }
          break;
        case 35:
          require(c, c.a.length - 1);
          break;
        case 36:
          require(c, 0);
          break;
        default:
          return;
      }
      event.preventDefault();
    }, false);
    this.u = new (options.Te || Object.constructor);
    this.u.ha(this);
    t.appendChild(this.u.f);
    if (this.Sc.tf) {
      this.u.setActive(true);
    }
    if ("undefined" != typeof this.u.jf) {
      $(t, "a", {
        onclick: function () {
          /** @type {number} */
          values.u.qe = values.u.jf ? 0 : 1;
        },
        style: {
          cursor: "pointer",
          color: "rgba(0,0,0,0.25)",
          position: "absolute",
          left: 0,
          top: 0,
          padding: "2px 8px",
          zIndex: 1E3
        }
      }, "\u25be");
    }
    /**
     * @return {?}
     */
    window.onresize = function () {
      return c.na();
    };
    /**
     * @return {undefined}
     */
    window.onhashchange = function () {
      var selector = search();
      if (selector.Xb != c.A) {
        update(c, selector.Xb, selector.Cb);
      } else {
        if (selector.Cb != c.Y) {
          require(c, selector.Cb);
        }
      }
    };
  }
  /**
   * @param {!Object} a
   * @param {number} val
   * @return {undefined}
   */
  function require(a, val) {
    a.u.qb(val);
    a.ce();
  }
  /**
   * @return {?}
   */
  function search() {
    /** @type {!Array<string>} */
    var a = window.location.hash.substring(1).split("/");
    /** @type {number} */
    var val = parseInt(a[0], 10);
    if (isNaN(val) || 0 > val) {
      /** @type {number} */
      val = 0;
    }
    /** @type {number} */
    a = 1 < a.length ? parseInt(a[1], 10) : 1;
    if (isNaN(a) || 1 > a) {
      /** @type {number} */
      a = 1;
    }
    return {
      Xb: a,
      Cb: val
    };
  }
  /**
   * @param {!Window} self
   * @param {boolean} c
   * @return {undefined}
   */
  function createElement(self, c) {
    if (self.C = c) {
      /** @type {string} */
      self.G = window.location.href;
      self.S = self.Y;
      resolve(self.H, [$("a", {
        href: "",
        onclick: function () {
          if (self.G) {
            if (window.history.replaceState) {
              try {
                window.history.replaceState({}, "", self.G);
              } catch (d) {
              }
            }
            var selector = search();
            update(self, selector.Xb, selector.Cb);
          } else {
            update(self, self.A, self.S);
          }
          createElement(self, false);
          return false;
        },
        style: {
          color: "#c22"
        }
      }, "RESET"), " | ", $("a", {
        href: "",
        onclick: function () {
          window.prompt("LINK", window.location.href);
          return false;
        }
      }, "LINK")]);
    } else {
      resolve(self.H, window.k2pback.gdata ? $("a", {
        target: "_blank",
        href: window.k2pback.gdata
      }, "TXT") : []);
    }
  }
  /**
   * @param {!Object} config
   * @param {number} i
   * @param {number} notSetValue
   * @return {undefined}
   */
  function update(config, i, notSetValue) {
    /** @type {number} */
    config.A = i;
    var source = 1 < i ? i.toString() : "";
    if (config.B[i]) {
      /** @type {boolean} */
      config.B[i].checked = true;
    }
    /** @type {string} */
    config.b = "";
    if ("undefined" !== typeof window.k2pback.w) {
      /** @type {boolean} */
      i = 0 == config.mb(0);
      config.b = window.k2pback[(i ? "w" : "b") + source] + " - " + window.k2pback[(i ? "b" : "w") + source];
    }
    if (i = config.fe()) {
      config.b += (config.b ? " " : "") + "(" + i + ")";
    }
    resolve(config.L, $("b", [config.b]));
    source = config.xe(window.k2pback["m" + source]);
    /** @type {number} */
    config.ga = 0;
    config.u.history(source.J, source.O);
    config.g = source.J;
    config.a = source.O;
    require(config, notSetValue || 0);
  }
  /**
   * @return {undefined}
   */
  function Tool() {
  }
  /**
   * @param {!Object} self
   * @return {undefined}
   */
  function link(self) {
    if (!self.Rb) {
      /**
       * @param {!Event} event
       * @return {?}
       */
      self.f.onclick = function (event) {
        if (!self.ia) {
          return false;
        }
        var el = event.target;
        if ("BUTTON" == el.tagName || "INPUT" == el.tagName || "IMG" == el.tagName && self.K && self.K.Z && self.K.Z.a == el) {
          return true;
        }
        el = getMousePosition(self.f, event.clientX, event.clientY);
        self.Vb(el.x, el.y, false, 3 == event.which || event.altKey || event.ctrlKey);
        return false;
      };
    }
    if (self.V || self.kb) {
      if (self.V) {
        extend(self.f, {
          touchAction: "none"
        });
      }
      self.f.addEventListener(self.V ? "pointerdown" : "mousedown", function (event) {
        if (!self.ia || "BUTTON" == event.target.tagName) {
          return true;
        }
        var val = "undefined" == typeof event.pointerId ? 0 : event.pointerId;
        if (self.Rb) {
          return val = getMousePosition(self.f, event.clientX, event.clientY), self.Vb(val.x, val.y, false, 3 == event.which || event.altKey || event.ctrlKey), event.stopPropagation(), false;
        }
        if (2 == self.Ba) {
          if (self.za) {
            return true;
          }
          self.a = val;
          trigger(self);
          getMousePosition(self.f, event.clientX, event.clientY);
          event.stopPropagation();
          return false;
        }
        return 0 == self.Ba || "IMG" != event.target.tagName ? true : log(self, event.target, event.clientX, event.clientY) ? (self.V && (self.a = val, self.f.setPointerCapture(val)), event.stopPropagation(), event.preventDefault(), false) : true;
      }, false);
      self.f.addEventListener(self.V ? "pointermove" : "mousemove", function (event) {
        if (!self.ia) {
          return true;
        }
        var order_ID = "undefined" == typeof event.pointerId ? 0 : event.pointerId;
        if (2 == self.Ba) {
          if (!self.za || !self.V || self.a != order_ID) {
            return true;
          }
          getMousePosition(self.f, event.clientX, event.clientY);
          return false;
        }
        if (self.ca) {
          if (!self.K || !self.V || self.a != order_ID) {
            return true;
          }
          startDrag(self, event.clientX, event.clientY);
          return false;
        }
        if (self.V && "mouse" != event.pointerType) {
          return true;
        }
        event = getMousePosition(self.f, event.clientX, event.clientY);
        self.Zc(event.x, event.y, !self.ia || self.ca);
        return false;
      }, false);
      self.f.addEventListener(self.V ? "pointerup" : "mouseup", function (event) {
        if (!self.ia) {
          return true;
        }
        var order_ID = "undefined" == typeof event.pointerId ? 0 : event.pointerId;
        if (2 == self.Ba) {
          if (!self.za || !self.V || self.a != order_ID) {
            return true;
          }
          getMousePosition(self.f, event.clientX, event.clientY);
          replace(self);
          return false;
        }
        if (self.ca) {
          if (!self.K || !self.V || self.a != order_ID) {
            return true;
          }
          notify(self, event.clientX, event.clientY, 3 == event.which || event.altKey || event.ctrlKey);
          return false;
        }
        return true;
      }, false);
      self.f.addEventListener(self.V ? "pointerout" : "mouseout", function (event) {
        if (!self.ia) {
          return false;
        }
        if (2 == self.Ba || self.ca || self.V && "mouse" != event.pointerType) {
          return true;
        }
        event = getMousePosition(self.f, event.clientX, event.clientY);
        self.Zc(event.x, event.y, !self.ia || self.ca);
        return false;
      }, false);
    }
    if (self.be) {
      closeLightbox(self);
    }
  }
  /**
   * @param {!Object} e
   * @return {undefined}
   */
  function closeLightbox(e) {
    /**
     * @param {!Object} event
     * @return {?}
     */
    e.f.ontouchstart = function (event) {
      if (!e.ia || !event.touches || 1 != event.touches.length) {
        return true;
      }
      event = event.touches[0];
      if ("BUTTON" == event.target.tagName) {
        return true;
      }
      if (2 == e.Ba) {
        if (e.za) {
          return true;
        }
        e.Hb = event.identifier;
        trigger(e);
        getMousePosition(e.f, event.clientX, event.clientY);
        return false;
      }
      return "IMG" != event.target.tagName ? true : log(e, event.target, event.clientX, event.clientY) ? (e.Hb = event.identifier, false) : true;
    };
    /**
     * @param {!Event} evt
     * @return {?}
     */
    e.f.ontouchmove = function (evt) {
      if (!e.ia || !evt.changedTouches) {
        return true;
      }
      var i;
      /** @type {null} */
      var event = null;
      /** @type {number} */
      i = evt.changedTouches.length - 1;
      for (; 0 <= i; i--) {
        if (evt.changedTouches[i].identifier == e.Hb) {
          event = evt.changedTouches[i];
          break;
        }
      }
      if (!event) {
        return true;
      }
      if (2 == e.Ba) {
        if (!e.za) {
          return true;
        }
        getMousePosition(e.f, event.clientX, event.clientY);
        return false;
      }
      if (!e.ca || !e.K) {
        return true;
      }
      startDrag(e, event.clientX, event.clientY);
      return false;
    };
    /**
     * @param {!Event} evt
     * @return {?}
     */
    e.f.ontouchend = function (evt) {
      if (!e.ia || !evt.changedTouches) {
        return true;
      }
      var i;
      /** @type {null} */
      var event = null;
      /** @type {number} */
      i = evt.changedTouches.length - 1;
      for (; 0 <= i; i--) {
        if (evt.changedTouches[i].identifier == e.Hb) {
          event = evt.changedTouches[i];
          break;
        }
      }
      if (!event) {
        return true;
      }
      if (2 == e.Ba) {
        if (!e.za) {
          return true;
        }
        getMousePosition(e.f, event.clientX, event.clientY);
        replace(e);
        return true;
      }
      if (!e.ca || !e.K) {
        return true;
      }
      notify(e, event.clientX, event.clientY, false);
      /** @type {number} */
      e.Hb = -1;
      return true;
    };
  }
  /**
   * @param {!Object} o
   * @param {!Object} element
   * @param {string} value
   * @param {number} callback
   * @param {number} options
   * @param {boolean} extra
   * @return {undefined}
   */
  function animate(o, element, value, callback, options, extra) {
    value = "undefined" != typeof value ? value : false;
    callback = "undefined" != typeof callback ? callback : 0;
    options = "undefined" != typeof options ? options : 0;
    extra = "undefined" != typeof extra ? extra : false;
    if (o.K != element && o.K) {
      if (!o.K.Z) {
        split("ERR " + window.k2url.game + " PI/NULL/0 " + (element ? element.x + "," + element.y : "p=null") + ", s.x,y=" + o.K.x + "," + o.K.y + ", act:" + o.wb);
      }
      extend(o.K.Z.a, {
        zIndex: 25
      });
    }
    if (o.K = element) {
      extend(o.K.Z.a, {
        zIndex: 50
      });
      /** @type {string} */
      element.a = value;
      if (extra) {
        construct(o, element.Z.a, callback, options);
      }
    } else {
      if (o.ca) {
        addEventListener(o);
      }
    }
  }
  /**
   * @param {!Object} o
   * @param {string} k
   * @param {undefined} options
   * @param {undefined} callback
   * @return {?}
   */
  function log(o, k, options, callback) {
    /** @type {null} */
    var element = null;
    /** @type {number} */
    var i = o.Ea.length - 1;
    for (; 0 <= i; i--) {
      if (o.Ea[i].Z && k == o.Ea[i].Z.a) {
        element = o.Ea[i];
        break;
      }
    }
    if (!element) {
      return false;
    }
    if (o.K && o.K != element) {
      return animate(o, element, false, options, callback, true), true;
    }
    animate(o, element, o.K == element, options, callback, true);
    return true;
  }
  /**
   * @param {!Object} a
   * @param {?} b
   * @param {?} c
   * @return {undefined}
   */
  function startDrag(a, b, c) {
    if (a.ca && a.K) {
      var d = a.K.Z;
      if (d) {
        /** @type {string} */
        d.a.style.left = b + a.$d + "px";
        /** @type {string} */
        d.a.style.top = c + a.ae + "px";
      }
      if (!a.jb && (5 < Math.abs(a.Tc - b) || 5 < Math.abs(a.Uc - c))) {
        /** @type {boolean} */
        a.jb = true;
      }
    }
  }
  /**
   * @param {!Object} options
   * @param {number} value
   * @param {!Date} width
   * @param {boolean} percent
   * @return {undefined}
   */
  function notify(options, value, width, percent) {
    if (options.ca && options.K) {
      addEventListener(options);
      if (!options.jb && (5 < Math.abs(options.Tc - value) || 5 < Math.abs(options.Uc - width))) {
        /** @type {boolean} */
        options.jb = true;
      }
      if (!options.jb) {
        if (options.K.a) {
          animate(options, null);
        }
        value = getMousePosition(options.f, value, width);
        options.Vb(value.x, value.y, true, percent);
      }
    }
  }
  /**
   * @param {!Object} data
   * @param {!Element} target
   * @param {number} value
   * @param {number} current
   * @return {undefined}
   */
  function construct(data, target, value, current) {
    /** @type {boolean} */
    data.ca = true;
    if (target) {
      /** @type {number} */
      data.$d = parseInt(target.style.left, 10) - value;
      /** @type {number} */
      data.ae = parseInt(target.style.top, 10) - current;
      /** @type {number} */
      data.Tc = value;
      /** @type {number} */
      data.Uc = current;
      /** @type {boolean} */
      data.jb = false;
    }
    if (data.kb) {
      /**
       * @param {!Event} event
       * @return {?}
       */
      data.f.ownerDocument.onmousemove = function (event) {
        startDrag(data, event.clientX, event.clientY);
        return false;
      };
      /**
       * @param {!Event} event
       * @return {?}
       */
      data.f.ownerDocument.onmouseup = function (event) {
        notify(data, event.clientX, event.clientY, false);
        return false;
      };
    }
  }
  /**
   * @param {!Object} options
   * @return {undefined}
   */
  function addEventListener(options) {
    if (options.ca) {
      /** @type {boolean} */
      options.ca = false;
      if (options.V && 0 <= options.a) {
        try {
          options.f.releasePointerCapture(options.a);
        } catch (b) {
        }
        /** @type {number} */
        options.a = -1;
      }
      if (options.kb) {
        /** @type {null} */
        options.f.ownerDocument.onmousemove = null;
        /** @type {null} */
        options.f.ownerDocument.onmouseup = null;
      }
    }
  }
  /**
   * @param {!Object} e
   * @return {undefined}
   */
  function trigger(e) {
    /** @type {boolean} */
    e.za = true;
    if (e.V && 0 <= e.a) {
      e.f.setPointerCapture(e.a);
    }
    if (e.kb) {
      /**
       * @param {!Event} event
       * @return {?}
       */
      e.f.ownerDocument.onmousemove = function (event) {
        getMousePosition(e.f, event.clientX, event.clientY);
        return false;
      };
      /**
       * @param {!Event} event
       * @return {?}
       */
      e.f.ownerDocument.onmouseup = function (event) {
        getMousePosition(e.f, event.clientX, event.clientY);
        replace(e);
        return false;
      };
    }
  }
  /**
   * @param {!Object} e
   * @return {undefined}
   */
  function replace(e) {
    /** @type {boolean} */
    e.za = false;
    if (e.V && 0 <= e.a) {
      e.f.releasePointerCapture(e.a);
      /** @type {number} */
      e.a = -1;
    }
    if (e.kb) {
      /** @type {null} */
      e.f.ownerDocument.onmousemove = null;
      /** @type {null} */
      e.f.ownerDocument.onmouseup = null;
    }
  }
  /**
   * @param {!Object} plugin
   * @return {?}
   */
  function loadPlugin(plugin) {
    /** @type {boolean} */
    var viewElement = 0 == plugin.Zb;
    if (!viewElement) {
      /** @type {boolean} */
      plugin.we = true;
    }
    return viewElement;
  }
  /**
   * @param {?} o
   * @return {undefined}
   */
  function isArray(o) {
    if (!o.vd && 0 < o.hb && 0 < o.gb) {
      /** @type {number} */
      o.vd = setTimeout(function () {
        /** @type {number} */
        o.vd = 0;
        o.ic();
      }, 0);
    }
  }
  /**
   * @param {?} item
   * @param {number} c
   * @param {number} v
   * @return {?}
   */
  function f(item, c, v) {
    /** @type {number} */
    var i = 0;
    var responsiveLayoutsCount = item.Da.length;
    for (; i < responsiveLayoutsCount; i++) {
      var cell = item.Da[i];
      if (0 > cell.x && 0 > cell.y && -1 == cell.F) {
        return cell.x = c, cell.y = v, cell;
      }
    }
    c = new Particle(item.Ac[-1].cloneNode(true), c, v);
    item.f.appendChild(c.a);
    item.Da.push(c);
    extend(c.a, {
      position: "absolute",
      zIndex: 25,
      display: "none"
    });
    return c;
  }
  /**
   * @param {number} radius
   * @param {number} x
   * @param {number} h
   * @param {number} i
   * @param {string} color
   * @return {undefined}
   */
  function paint(radius, x, h, i, color) {
    var context = ctx;
    /** @type {number} */
    radius = 10 * radius.ja;
    /** @type {number} */
    var len = .6 * radius;
    /** @type {number} */
    i = i ? -1 : 1;
    context.beginPath();
    /** @type {string} */
    context.fillStyle = color;
    context.moveTo(x, h + len / 2 * i);
    context.lineTo(x - radius / 2, h - len / 2 * i);
    context.lineTo(x + radius / 2, h - len / 2 * i);
    context.fill();
  }
  /**
   * @param {number} img
   * @param {number} posx
   * @param {number} posy
   * @return {undefined}
   */
  function Particle(img, posx, posy) {
    /** @type {number} */
    this.a = img;
    /** @type {number} */
    this.F = -1;
    /** @type {number} */
    this.x = posx;
    /** @type {number} */
    this.y = posy;
  }
  /**
   * @param {!Object} el
   * @param {?} res
   * @param {!Object} options
   * @return {undefined}
   */
  function Controller(el, res, options) {
    var state = this;
    /** @type {!Object} */
    this.app = el;
    /** @type {!Object} */
    this.o = options;
    this.f = $(res, "div", options.He);
    this.a = $(this.f, "div", options.Wd || {
      style: {
        position: "absolute",
        padding: "3px 4px 3px 8px",
        top: 0,
        left: 0,
        right: 0,
        bottom: "2.5em",
        overflowY: "scroll",
        WebkitOverflowScrolling: "touch"
      }
    });
    extend(this.a, {
      wordWrap: "break-word",
      background: "#fff"
    });
    /**
     * @param {!Event} event
     * @return {?}
     */
    this.a.onselectstart = function (event) {
      event.stopPropagation();
      return true;
    };
    el = $(this.f, "form", options.Xe || {});
    res = $(el, "div", {
      style: {
        display: "table",
        width: "100%"
      }
    });
    if (!options.Wd) {
      extend(res, {
        position: "absolute",
        bottom: 0
      });
    }
    res = $(res, "div", {
      style: {
        display: "table-cell"
      }
    });
    this.ab = $(res, "input", options.Kf || {
      className: "bsbb",
      style: {
        width: "100%",
        margin: 0
      }
    });
    Object.assign(this.ab, {
      name: "somename",
      type: "text",
      autocomplete: "off",
      autocapitalize: "off"
    });
    /**
     * @return {?}
     */
    el.onsubmit = function () {
      var a = state.ab.value.trim();
      /** @type {string} */
      state.ab.value = "";
      if (0 < a.length) {
        state.o.Zd(a);
      }
      return false;
    };
  }
  /**
   * @param {number} marker
   * @param {!Object} transform
   * @param {(Object|string)} options
   * @return {undefined}
   */
  function draw(marker, transform, options) {
    /** @type {number} */
    this.tab = marker;
    /** @type {(Object|string)} */
    this.o = options;
    this.o.Ta = this.o.Ta || 0;
    /** @type {number} */
    this.g = 0;
    /** @type {number} */
    this.b = -1;
    this.A = this.o.fc + (this.o.Rc ? 1 : 0);
    var self = this;
    this.f = $(transform, "div", {
      style: {
        position: "absolute",
        width: "100%",
        height: "100%"
      }
    });
    this.sa = $(this.f, "div", {
      style: {
        position: "absolute",
        left: 0,
        right: 0,
        top: 0,
        bottom: 0,
        overflowY: "scroll",
        WebkitOverflowScrolling: "touch",
        background: "#fff"
      }
    });
    this.a = $(this.sa, "table", {
      className: "br",
      style: {
        width: "100%",
        borderCollapse: "collapse"
      }
    });
    if (this.o.Va || this.o.ad) {
      callback(this.sa, "mb1s");
      if (this.o.Va) {
        /**
         * @param {string} e
         * @return {?}
         */
        this.a.onclick = function (e) {
          e = e.target;
          if ("TD" == e.tagName && e.cellIndex >= self.o.Ta) {
            /** @type {number} */
            e = e.parentNode.rowIndex * self.A + e.cellIndex - self.o.Ta;
            if (e != self.b && e < self.g) {
              close(self.tab, e);
            }
          }
          return false;
        };
      }
      marker = $(this.f, "div", {
        className: "lh1s",
        style: {
          position: "absolute",
          left: 0,
          right: 0,
          bottom: 0
        }
      });
      if (this.o.ad) {
        $(marker, [$("button", {
          className: "minw",
          onclick: function () {
            self.tab.send([94, self.tab.F], null);
          }
        }, this.o.ad), " "]);
      }
      if (this.o.Va) {
        $(marker, [$("button", {
          className: "minw",
          onclick: function () {
            if (0 <= self.b - 1) {
              close(self.tab, self.b - 1);
            }
          }
        }, "<"), " ", $("button", {
          className: "minw",
          onclick: function () {
            if (self.b + 1 < self.g) {
              close(self.tab, self.b + 1);
            }
          }
        }, ">")]);
      }
    }
    marker = this.o.Ta || this.o.Rc;
    /** @type {string} */
    this.B = "10%";
    /** @type {string} */
    this.G = this.C = (marker ? Math.round(90 / this.o.fc) : Math.round(100 / this.o.fc)) + "%";
    if (marker && 2 == this.o.fc) {
      /** @type {string} */
      this.C = "40%";
    }
  }
  /**
   * @param {!Object} where
   * @return {?}
   */
  function serialize(where) {
    return 11 < where.length ? where.substring(0, 9) + ".." : where;
  }
  /**
   * @param {!Object} o
   * @param {!Object} v
   * @return {undefined}
   */
  function add(o, v) {
    var g = o.g;
    var containerTR;
    var line;
    /** @type {number} */
    var p = 0;
    /** @type {number} */
    var i = 0;
    for (; i < v.length; i++) {
      if ("=" == v[i] || "_" == v[i]) {
        /** @type {number} */
        i = i + (o.A - 1);
      } else {
        if (0 == o.g % o.A && (containerTR = o.a.insertRow(-1), line = i + o.A < v.length ? v[i + 2] : "", p = "=" == line ? 2 : "_" == line ? 1 : 0, o.o.Ta && containerTR && (resolve(line = containerTR.insertCell(-1), [1 + Math.floor(o.g / o.A) + ""]), 1 == o.a.rows.length && (line.width = o.B))), o.g++, containerTR =
          o.a.rows[o.a.rows.length - 1]) {
          resolve(line = containerTR.insertCell(-1), [serialize(v[i])]);
          if (o.o.Va) {
            extend(line, {
              cursor: "pointer"
            });
          }
          if (1 == o.a.rows.length) {
            line.width = o.o.Rc && 0 == (o.g - 1) % o.A ? o.B : 2 == containerTR.cells.length ? o.C : o.G;
          }
          if (p) {
            /** @type {string} */
            line.style.borderBottom = 1 == p ? "dashed 1px #000" : "solid 2px #000";
          }
        }
      }
    }
    if (o.o.Va) {
      o.Aa(o.g - g);
    } else {
      /** @type {number} */
      o.sa.scrollTop = o.sa.scrollHeight - o.sa.clientHeight;
    }
  }
  /**
   * @param {!Object} options
   * @param {number} d
   * @param {boolean} e
   * @return {undefined}
   */
  function multiply(options, d, e) {
    var attributesRow = options.a.rows[Math.floor(d / options.A)];
    if (attributesRow) {
      var valueProgess;
      if (valueProgess = attributesRow.cells[d % 2 + options.o.Ta]) {
        callback(valueProgess, "fb", e);
      }
    }
  }
  /**
   * @param {!Object} a
   * @param {number} b
   * @return {undefined}
   */
  function testCircleCircle(a, b) {
    if (a.o.Va) {
      if (-1 != a.b) {
        multiply(a, a.b, false);
      }
      /** @type {number} */
      a.b = b;
      if (-1 != b) {
        multiply(a, a.b, true);
      }
    }
  }
  /**
   * @param {!Object} value
   * @return {?}
   */
  function setItem(value) {
    return value.b == value.g - 1;
  }
  /**
   * @param {!Object} data
   * @param {?} i
   * @return {undefined}
   */
  function test(data, i) {
    var fr = this;
    /** @type {!Object} */
    this.app = data;
    this.Ja = this.app.j("gname");
    this.N = i;
    this.a = {};
    /** @type {!Array} */
    this.b = [];
    /** @type {number} */
    this.B = 1;
    /** @type {boolean} */
    this.L = false;
    var self = this;
    this.f = document.getElementById("pretl") || $(this.app.B, "div", {
      style: {
        display: "none"
      }
    });
    extend(this.f, {
      display: "none"
    });
    callback(this.f, "tblobby");
    callback(this.f, "usno");
    this.S = $(this.f, "div", {
      className: "tbvusers"
    });
    this.A = document.getElementById("pretabs") || $(this.f, "div");
    callback(this.A, "tbvtabs");
    var inputElements;
    var buttonIndex;
    var elem;
    var node;
    html(this.f, data = $("div", {
      className: "newtab2 dcpd"
    }), this.f.firstChild);
    $(data, [inputElements = $("button", {
      className: "minwd",
      onclick: function () {
        self.app.send([payload], null);
      }
    }, this.app.j("bl_newtab")), " ", $("div", {
      className: "selcwr mro"
    }, [buttonIndex = $("button", {
      className: "selcbt min85"
    }, $("div", "-")), elem = $("select", {
      className: "selcsl",
      onchange: function (event) {
        if (event = event.target.options[event.target.selectedIndex]) {
          self.app.X("/join " + event.text.split(" ")[0]);
        }
      }
    })]), $("div", {
      className: "ib mro"
    }, [$("button", {
      className: "ubut",
      onclick: function () {
        debug(self.app, self.app.N + (self.L ? "" : "/p"));
        return false;
      }
    }, $("div", {
      className: "uicon"
    }))]), node = $("span", {
      className: "tuinfo",
      style: {
        color: "gray"
      }
    }, "-")]);
    expect(this.app, "rooms", function (self) {
      /** @type {number} */
      elem.options.length = 0;
      $(elem, self.list.map(function (body, state) {
        return $("option", state == self.rb ? {
          selected: true
        } : {}, body);
      }));
      resolve(buttonIndex, [self.list[0 <= self.rb ? self.rb : 0].split(" ")[0] || "-"]);
    });
    expect(this.app, "tumode", function (value) {
      extend(inputElements, {
        display: value ? "none" : "inline-block"
      });
    });
    expect(this.app, "tuinfo", function (ease) {
      return resolve(node, [ease || "-"]);
    });
    $(this.A, "div", {
      className: "tldeco"
    });
    this.M = $(this.A, "div");
    this.G = $(this.A, "div", {
      className: "chpan dcpd"
    });
    $(this.G, "div", {
      className: "chtop"
    }, [this.I = $("select", {
      className: "chgrlist minwd",
      onchange: function (event) {
        var NEGATIVE = event.target.selectedIndex;
        self.app.X("/g " + (0 == NEGATIVE ? "-" : 1 == NEGATIVE ? "+" : event.target.value.split(" ")[0]));
      }
    }), $("input", {
      type: "checkbox",
      onchange: function () {
        var c = this.checked;
        callback(self.G, "chopen", c);
        if (c) {
          self.g.Aa();
        }
      }
    }), this.app.j("bl_whisper"), " ", this.ba = $("span")]);
    data = $(this.G, "div", {
      className: "chsub"
    });
    this.g = new Controller(this.app, data, {
      Zd: function (left) {
        self.app.X(left);
      },
      He: {
        style: {
          display: "block",
          position: "relative",
          width: "100%",
          maxWidth: "640px",
          height: "140px"
        }
      }
    });
    extend(this.g.a, {
      border: "solid 1px rgba(0,0,0,0.3)"
    });
    expect(this.app, "chmode", function (a) {
      return callback(fr.f, "chmode", !!a);
    });
    this.C = $(this.A, "div", {
      className: "tlst"
    });
    expect(this.app, "tumode", function (server) {
      /** @type {number} */
      server = server ? 0 : 1;
      if (server != self.B) {
        /** @type {number} */
        self.B = server;
      }
    });
    var s = $(this.S, "div", {
      className: "ulpan",
      style: {
        display: "none"
      }
    }, [$("button", {
      onclick: function () {
        var sandbox = self.H;
        ok(sandbox, sandbox.Ra ? 0 : 1);
        return false;
      }
    }, this.app.j("tu_bcan"))]);
    expect(this.app, "urole", function (i) {
      extend(s, {
        display: i == cursel ? "block" : "none"
      });
    });
    var fProduction = this.app.o.ob;
    this.H = new Color(this.app, this.S, {
      oe: true,
      cols: fProduction ? [file, -1] : [F, file],
      ud: false,
      Pc: function (data, sandbox) {
        var name = data.name;
        var expectedChecksum = self.app.gc;
        if (1 == sandbox.Ra && expectedChecksum == cursel || ed) {
          ok(sandbox, 0);
          config(self.app, name);
        } else {
          if (expectedChecksum == readChecksum || ignoreChecksum) {
            compile(self.app, name, data.R);
          } else {
            self.app.send([action], [name]);
          }
        }
      },
      Ue: true
    });
    logger(this.app, this.H);
    expect(this.app, "tumode", function (result) {
      var opts = self.H;
      /** @type {!Array} */
      var cols = 0 != result ? [value, F] : fProduction ? [file, -1] : [F, file];
      /** @type {number} */
      result = 0 != result ? value : 0;
      /** @type {number} */
      result = void 0 === result ? 0 : result;
      /** @type {!Array} */
      opts.cols = cols;
      if (opts.B) {
        opts.B = overlay(opts, opts.B);
      }
      forward(opts, result);
    });
  }
  /**
   * @param {!Object} data
   * @param {number} x
   * @return {undefined}
   */
  function refresh(data, x) {
    var key = x.split(" ");
    if ("/c" == key[0]) {
      /** @type {string} */
      data.g.a.innerHTML = "";
    } else {
      if ("/g" == key[0]) {
        callback(data.G, "chgrp");
        /** @type {number} */
        x = parseInt(key[1], 10);
        key = key.slice(2);
        resolve(data.ba, [0 < key.length ? "(" + key.length + ")" : ""]);
        key.unshift("-", "START");
        /** @type {number} */
        data.I.options.length = 0;
        $(data.I, key.map(function (body) {
          return $("option", {}, body);
        }));
        if (0 <= x) {
          /** @type {number} */
          data.I.selectedIndex = x + 2;
        }
      } else {
        data.g.append(x);
      }
    }
  }
  /**
   * @param {!Object} options
   * @param {!Object} data
   * @param {!Object} template
   * @return {?}
   */
  function parse(options, data, template) {
    var i;
    /** @type {number} */
    var to = 0;
    if (1 == options.N) {
      var el = $("div", {
        className: "tplone tavail"
      }, [$("div", {
        className: "tpllist"
      }, data.O[position + 0] || "-"), $("div", {
        className: "tpar0"
      }, data.O[0])]);
    } else {
      el = $("div", {
        className: "tplbl"
      });
      /** @type {number} */
      i = data.nd - 1;
      for (; 0 <= i; i--) {
        var number = 0 == data.J[base + i] && pow(data) && data.app.Fa >= translate(data);
        var value = data.O[position + i];
        var r = process(options.app, value);
        html(el, $("div", {
          className: number ? "tplemp" : "" != value ? "tplnorm" : "tplunav"
        }, [$("div", {
          className: r ? "r" + merge(options.app, r.g) : "rnone"
        }), value || "\u2013", $("span", {
          className: "tplrn snum"
        }, r ? 0 != (r.a & arg) ? "X" : exec(options.app, r.g) : "-")]), number || "" == value ? null : el.firstChild);
        if (number) {
          to++;
        }
      }
      $(el, "div", {
        className: "tpar0"
      }, [$("div", {
        className: "rnone"
      }), data.O[0]]);
    }
    return fn(options.C, $("a", {
      className: "awrap dcpd" + (0 < to ? " tavail" : ""),
      onclick: function () {
        server(options.app, data.F);
        return false;
      }
    }, $("div", {
      className: "tmaxw"
    }, [$("div", {
      className: "tnum"
    }, 2 == data.J[2] ? "-" : "#" + data.F), $("div", {
      className: "tpar1"
    }, data.O[0]), el, 1 < options.N ? $("div", {
      className: "tjoin"
    }, $("button", {
      className: "butbl"
    }, ">>")) : null])), template);
  }
  /**
   * @param {!Object} req
   * @param {!Object} data
   * @param {!Object} options
   * @return {undefined}
   */
  function jsonp(req, data, options) {
    if (!req.a.hasOwnProperty(data[1])) {
      data = new Node(req.app, data, options);
      req.a[data.F] = {
        Ua: data
      };
      data.D = parse(req, data);
      error(req, data);
    }
  }
  /**
   * @param {!Object} item
   * @param {?} i
   * @return {undefined}
   */
  function inverse(item, i) {
    if (item.a.hasOwnProperty(i)) {
      var val = item.a[i].Ua;
      if (val.D) {
        item.C.removeChild(val.D);
        /** @type {null} */
        val.D = null;
      }
      val = item.b.indexOf(val);
      if (-1 != val) {
        item.b.splice(val, 1);
      }
      delete item.a[i];
    }
  }
  /**
   * @param {!Object} context
   * @param {string} data
   * @param {!Array} key
   * @return {undefined}
   */
  function message(context, data, key) {
    if (context.a.hasOwnProperty(data[1])) {
      var item = context.a[data[1]].Ua;
      getType(item, data, key);
      data = context.b.indexOf(item);
      if (-1 != data) {
        item.D = parse(context, item, item.D);
        context.b.splice(data, 1);
        error(context, item);
      }
    }
  }
  /**
   * @param {!Object} b
   * @param {!Array} data
   * @param {!Array} items
   * @return {undefined}
   */
  function insert(b, data, items) {
    plugin(b);
    b.a = {};
    /** @type {number} */
    b.b.length = 0;
    var tx = data[1];
    /** @type {number} */
    var i = 3;
    /** @type {number} */
    var index = 0;
    var offset = tx;
    var count = data[2];
    for (; !(i >= data.length);) {
      if (-1 == tx) {
        offset = data[i];
        i++;
        /** @type {number} */
        count = position + (offset - (base - 1));
      }
      if (i + offset > data.length || index + count > items.length) {
        break;
      }
      var a = new Node(b.app, data.slice(i - 1, i - 1 + offset + 1), items.slice(index, index + count));
      b.a[a.F] = {
        Ua: a
      };
      a.D = parse(b, a);
      error(b, a);
      index = index + count;
      i = i + offset;
    }
  }
  /**
   * @param {!Object} s
   * @param {!Object} key
   * @return {undefined}
   */
  function error(s, key) {
    var index = s.b.length;
    var sibling = s.app.Fa;
    for (; 0 < index;) {
      var data = s.b[index - 1];
      var x = data.J[1];
      var len = key.J[1];
      /** @type {number} */
      x = x < len ? -1 : x > len ? 1 : 0;
      if (0 != s.B) {
        len = Number(data);
        var size = Number(key);
        var ch = td(data);
        var owner = td(key);
        /** @type {number} */
        x = -x;
        /** @type {number} */
        data = 0 < size && 0 == owner ? 0 < len && 0 == ch ? x : 1 : 0 < len && 0 == ch ? -1 : report(key, sibling) && 0 < size ? report(data, sibling) && 0 < len ? x : -1 : report(data, sibling) && 0 < len ? 1 : 2 == key.J[2] ? 2 == data.J[2] ? x : 1 : 2 == data.J[2] ? -1 : x;
      } else {
        /** @type {number} */
        data = x;
      }
      if (0 == s.B) {
        /** @type {number} */
        data = -data;
      }
      if (0 <= data) {
        break;
      }
      index--;
    }
    sibling = s.b[index];
    s.b.splice(index, 0, key);
    if (key.D) {
      s.C.insertBefore(key.D, sibling ? sibling.D : null);
    }
  }
  /**
   * @param {!Object} f
   * @return {undefined}
   */
  function plugin(f) {
    Object.keys(f.a).forEach(function (index) {
      var c = this.a[index].Ua;
      if (c.D) {
        this.C.removeChild(c.D);
        /** @type {null} */
        c.D = null;
      }
      delete this.a[index];
    }, f);
  }
  /**
   * @return {undefined}
   */
  function kernel() {
  }
  /**
   * @param {!Object} options
   * @param {number} key
   * @param {number} target
   * @param {number} type
   * @return {undefined}
   */
  function equal(options, key, target, type) {
    if (options.time) {
      /** @type {string} */
      options.time[key] = 0 > target ? "(?)" : "" + Math.floor(target / 60) + ":" + Math.floor(target % 60 / 10) + target % 60 % 10 + ("undefined" != typeof type && 0 < type ? "/" + type : "");
    }
  }
  /**
   * @return {undefined}
   */
  function EventListener() {
  }
  /**
   * @param {!Object} a
   * @return {?}
   */
  function inject(a) {
    return 0 != (a.U & 4) && (0 == (a.U & 2) || 0 != (a.U & 2) && !isObject(a)) && (pow(a.ma) && a.app.Fa >= translate(a.ma) || 0 != (a.la & 2)) && !isString(a);
  }
  /**
   * @param {!Object} o
   * @return {?}
   */
  function isObject(o) {
    return 0 != (o.U & 1);
  }
  /**
   * @param {!Object} obj
   * @return {?}
   */
  function isString(obj) {
    return 0 != (obj.la & 4);
  }
  /**
   * @param {!Object} t
   * @param {string} id
   * @param {number} c
   * @return {undefined}
   */
  function setTimeout(t, id, c) {
    t.send([82, t.F, c], [id]);
  }
  /**
   * @return {undefined}
   */
  function Promise() {
  }
  /**
   * @param {!Object} self
   * @return {undefined}
   */
  function setupEvents(self) {
    if (self.app.P && self.a.je) {
      /**
       * @return {?}
       */
      window.onkeypress = function () {
        return true;
      };
    }
  }
  /**
   * @param {!Object} obj
   * @return {undefined}
   */
  function fail(obj) {
    $(obj.f, "div", {
      className: "thnavcont usno tama"
    }, [$("button", {
      className: "xbut hdhei",
      style: {
        zIndex: cluezIndex + 1,
        right: 0,
        width: "44px"
      },
      onclick: function () {
        copy(obj);
      }
    }, "X"), $("button", {
      className: "cmenubut hdhei",
      style: {
        zIndex: cluezIndex + 1,
        right: "38px",
        width: "44px"
      },
      onclick: function () {
        this.blur();
        _(obj, !transform(obj.f, "sbdropvis"));
      }
    }, $("div", {
      className: "cmenu"
    }))]);
    var w = $(obj.f, "div", {
      className: "thead bsbb usno hdhei",
      style: {
        zIndex: cluezIndex
      }
    });
    w = $(w, "div", {
      style: {
        display: "table",
        width: "100%",
        height: "100%"
      }
    });
    w = $(w, "div", {
      style: {
        display: "table-cell",
        verticalAlign: "middle",
        textAlign: "center"
      }
    });
    open(obj, w);
  }
  /**
   * @param {!Object} item
   * @param {!Object} hash
   * @return {undefined}
   */
  function open(item, hash) {
    expect(item.app, "tabplayers", function () {
      if (0 >= item.F) {
        resolve(hash, []);
      } else {
        var args = item.T;
        /** @type {number} */
        var cell_amount = 0;
        var i;
        var ix;
        /** @type {number} */
        i = 0;
        for (; i < args.ea; i++) {
          if (null != args.name[i]) {
            cell_amount++;
          }
        }
        /** @type {null} */
        var message = null;
        /** @type {null} */
        var masterPageName = null;
        /** @type {boolean} */
        var nameMatch = 2 < cell_amount;
        if (item.a.Jf) {
          /** @type {boolean} */
          nameMatch = false;
        }
        var labelMatch = 3 == cell_amount && item.a.If;
        if (!labelMatch) {
          /** @type {number} */
          var gridX = 2 < cell_amount ? 2 : 1;
          message = $("div", {
            className: nameMatch ? "tabmaxmed" : "tabany",
            style: 1 < gridX ? {
              margin: "0 auto",
              fontSize: "13px",
              lineHeight: "1.2",
              paddingRight: "8px"
            } : {
                margin: "0 auto",
                paddingRight: "8px"
              }
          });
          /** @type {number} */
          i = 0;
          for (; 2 > i; i++) {
            /** @type {number} */
            var size = i + (1 == gridX && null == args.name[i] ? 1 : 0);
            var x = $(message, "div", {
              style: {
                display: "table-cell",
                textAlign: 1 > i ? "right" : "left",
                width: "50%",
                padding: "0 .4em",
                whiteSpace: "nowrap"
              }
            });
            var node = $(x, "div", {
              className: "ib",
              style: {
                textAlign: "left",
                marginRight: ".4em"
              }
            });
            /** @type {number} */
            var index = size;
            /** @type {number} */
            ix = 0;
            for (; ix < gridX; ix++, index = index + 2) {
              if (0 < ix) {
                $(node, "br");
              }
              if (args.Sa && args.R) {
                $(node, "div", {
                  className: "ib",
                  style: {
                    width: "3px",
                    height: ".8em",
                    background: args.Sa[index],
                    marginRight: "3px"
                  }
                });
              }
              $(node, "b", {}, [args.R && args.Eb ? args.R[index] : "#" + (index + 1)]);
            }
            node = $(x, "div", {
              className: "ib",
              style: {
                textAlign: "right"
              }
            });
            /** @type {number} */
            index = size;
            /** @type {number} */
            ix = 0;
            for (; ix < gridX; ix++, index = index + 2) {
              if (0 < ix) {
                $(node, "br");
              }
              if (args.time) {
                $(node, [args.time[index] + ""]);
              }
            }
          }
        }
        if (nameMatch || labelMatch) {
          masterPageName = $("div", {
            className: labelMatch ? "fs15xs" : "minmed",
            style: {
              textAlign: "center",
              paddingRight: "21px"
            }
          });
          /** @type {number} */
          i = 0;
          for (; i < cell_amount; i++) {
            if (null != args.name[i]) {
              node = $(masterPageName, "div", {
                className: "ib",
                style: {
                  margin: "0 .4em"
                }
              });
              if (args.Sa) {
                $(node, "div", {
                  className: "ib",
                  style: {
                    width: "4px",
                    height: ".8em",
                    background: args.Sa[i],
                    marginRight: "4px",
                    verticalAlign: "middle"
                  }
                });
              }
              $(node, "b", {
                style: {
                  marginRight: ".4em"
                }
              }, [args.R && args.Eb ? args.R[i] : "#" + (i + 1)]);
              if (args.time) {
                $(node, [args.time[i] + ""]);
              }
            }
          }
        }
        resolve(hash, [message, masterPageName]);
      }
    });
  }
  /**
   * @param {!Object} obj
   * @return {undefined}
   */
  function done(obj) {
    var b;
    obj.xa = $(obj.f, "div", {
      className: "bsbb tsb sbclrd"
    });
    var one = $(obj.xa, "div", {
      className: "tsbinner bsbb"
    });
    $(one, "div", {
      className: "ttlcont"
    }, [$("div", {
      className: "ttlnav"
    }, [b = $("button", {
      className: "butsys butlh",
      onclick: function () {
        then(obj.app);
      }
    }, "\u2013"), $("button", {
      className: "butsys butlh",
      onclick: function () {
        copy(obj);
      }
    }, "X")]), obj.ba = $("div", ["-"])]);
    expect(obj.app, "chatalert", function (a) {
      return callback(b, "alert", a);
    });
    expect(obj.app, "tabopen", function () {
      return resolve(obj.ba, [upload(obj)]);
    });
    expect(obj.app, "tabstatus", function () {
      return resolve(obj.ba, [upload(obj)]);
    });
    createOption(obj, one);
    var t = $(one, "div", {
      className: "tsinsb lh1s",
      style: {
        textAlign: "center",
        background: "#f21",
        color: "#fff",
        paddingLeft: ".5em",
        paddingRight: ".5em"
      }
    });
    obj.va = $(t, "div", {
      className: "tstatlabl nowrel"
    }, "-");
    $(t, "div", {
      className: "tstatstrl"
    }, [$("button", {
      className: "butwb",
      style: {
        minWidth: "8em"
      },
      onclick: function () {
        obj.send([85, obj.F], null);
        return false;
      }
    }, obj.j("bl_start"))]);
    listen(obj, one);
    obj.b.qf();
  }
  /**
   * @param {!Object} value
   * @param {!Object} label
   * @return {undefined}
   */
  function createOption(value, label) {
    var input_container = $(label, "div");
    anonymous(value, input_container);
    expect(value.app, "tabplayers", function () {
      return anonymous(value, input_container);
    });
  }
  /**
   * @param {!Object} self
   * @param {!Array} options
   * @return {undefined}
   */
  function listen(self, options) {
    apply(self, options);
    self.b = factory(self, options);
    options = self.b;
    setup(self, options.pb);
    options.add(self.j("sw_chat"), self.ta.f);
    options = self.b;
    if (!self.a.ke) {
      self.history = new draw(self, options.pb, {
        fc: self.a.pe,
        Va: !self.a.me,
        ad: self.a.Wb,
        Rc: self.a.Ff ? true : false,
        Ta: self.a.le ? 0 : 1
      });
      self.oa = options.add(self.app.P ? self.j("sw_history") : null, self.history.f);
    }
    activate(self, self.b);
    check(self, self.b);
  }
  /**
   * @param {!Object} module
   * @param {?} url
   * @return {?}
   */
  function factory(module, url) {
    module = {
      Vc: {},
      Na: {},
      pf: module.app.G,
      ha: function (callback) {
        this.f = $(callback, "div", {
          className: "tcrdcont"
        }, [$("div", {
          className: "tcrdtabcont"
        }, this.sa = $("div", {
          className: "tcrdtab"
        })), this.pb = $("div", {
          className: "tcrdpan"
        })]);
      },
      add: function (a, b) {
        var c = this;
        /** @type {number} */
        var p = Object.keys(this.Na).length;
        /** @type {!Object} */
        this.Na[p] = b;
        extend(b, {
          visibility: p ? "hidden" : "inherit"
        });
        if (a) {
          $(this.sa, "div", {
            className: "tcrdcell"
          }, this.Vc[p] = $("button", {
            className: p ? "" : "active",
            onclick: function () {
              return c.show(p);
            }
          }, [a]));
        }
        return p;
      },
      show: function (i) {
        Object.keys(this.Vc).forEach(function (item) {
          callback(this.Vc[item], "active", i == item);
        }, this);
        Object.keys(this.Na).forEach(function (sub_id) {
          extend(this.Na[sub_id], {
            visibility: i == sub_id ? "inherit" : "hidden"
          });
        }, this);
        if (this.pf) {
          var contentScrollTop = this.Na[i].firstChild.scrollTop;
          this.pb.insertBefore(this.Na[i], null);
          this.Na[i].firstChild.scrollTop = contentScrollTop;
        }
      },
      qf: function () {
        this.pb.insertBefore(this.Na[0], null);
      }
    };
    module.ha(url);
    return module;
  }
  /**
   * @param {!Object} options
   * @return {undefined}
   */
  function createServer(options) {
    if (!options.H) {
      options.H = create(options.app, options.j("bl_invite"), {
        width: "81%",
        minWidth: "280px",
        maxWidth: "320px"
      }, {
        nopad: true
      });
      options.pa = new Color(options.app, options.H.Ma, {
        oe: true,
        cols: [F, file],
        vf: file,
        ud: true,
        Pc: function (catalogs) {
          success(options.app);
          options.send([95, options.F, 0], [catalogs.name]);
        },
        Ie: {
          className: "ovysct",
          style: {
            width: "100%",
            height: "300px",
            borderTop: "solid 1px #ddd"
          }
        }
      });
      logger(options.app, options.pa);
    }
    filter(options.app, options.H);
  }
  /**
   * @param {!Object} params
   * @return {undefined}
   */
  function page(params) {
    params.La = $(params.u.f, "div", {
      className: "tsinbo bsbb",
      style: {
        position: "absolute",
        width: "100%",
        height: "100%",
        textAlign: "center",
        zIndex: 70
      }
    });
    var element = $(params.La, "div", {
      style: {
        display: "table-cell",
        verticalAlign: "middle"
      }
    });
    element = $(element, "div", {
      className: "bsbb bs ib",
      style: {
        textAlign: "center",
        maxWidth: "80%",
        minWidth: "35%",
        padding: "0.75em 2em",
        border: "solid 3px #fff",
        background: "#e21",
        color: "#fff"
      }
    });
    params.qa = $(element, "div", {
      className: "tstatlabl fb"
    }, "");
    element = $(element, "div", {
      className: "tstatstrl"
    });
    $(element, "button", {
      className: "butwb",
      style: {
        marginTop: ".25em",
        minWidth: "8em"
      },
      onclick: function () {
        params.send([85, params.F], null);
        return false;
      }
    }, params.j("bl_start"));
  }
  /**
   * @param {!Object} options
   * @param {string} value
   * @param {boolean} once
   * @return {undefined}
   */
  function subscribe(options, value, once) {
    if (options.ya != value || options.Ka != once) {
      var ya = options.ya;
      var prefix = options.vc;
      /** @type {string} */
      options.ya = value;
      /** @type {boolean} */
      options.Ka = !!once;
      /** @type {string} */
      options.qa.innerHTML = options.ya ? (prefix ? prefix + "<br />" : "") + options.ya : "-";
      options.va.innerHTML = options.ya || "-";
      callback(options.f, "tstatstart", null != value && !!once);
      callback(options.f, "tstatact", null != options.ya);
      if (!(!!ya == (null != options.ya))) {
        _(options, null != options.ya);
      }
    }
  }
  /**
   * @param {!Object} node
   * @param {number} key
   * @param {number} a
   * @return {undefined}
   */
  function loop(node, key, a) {
    if (node.ka) {
      var tag = node.Ya;
      /** @type {number} */
      node.Ya = key;
      resolve(node.ua, [(a ? g(node, key, a[a.length - 1]) : "") || "-"]);
      node.ka.show(node.Ya ? 1 : 0);
      if (tag != node.Ya) {
        _(node, 0 != node.Ya);
      }
    }
  }
  /**
   * @param {!Object} d
   * @param {!Object} r
   * @param {number} b
   * @return {undefined}
   */
  function randomColor(d, r, b) {
    /** @type {!Array} */
    r = [93, d.F, r];
    if ("undefined" != typeof b) {
      r.push(b);
    }
    r.push(Math.floor((Date.now() - d.bb.g) / 100));
    d.send(r, null);
  }
  /**
   * @param {!Object} data
   * @return {?}
   */
  function bind(data) {
    /** @type {!Array} */
    var watchable = [];
    if (data.a.cf) {
      watchable.push(data.Lc = $("button", {
        className: "minw",
        onclick: function () {
          randomColor(data, 1, 1);
        }
      }, data.j("bl_draw")), " ");
    }
    if (data.a.ne) {
      watchable.push(data.Nc = $("button", {
        className: "minw",
        onclick: function () {
          data.ka.show(2);
        }
      }, data.j("bl_resign")), " ");
    }
    if (data.a.$c) {
      watchable.push(data.Oc = $("button", {
        className: "minw",
        onclick: function () {
          randomColor(data, 1, 2);
        }
      }, data.j("bl_undo")), " ");
    }
    if (data.a.Hf) {
      watchable.push(data.I = $("button", {
        onclick: function () {
          randomColor(data, 1, 10);
        }
      }, data.j("bl_resign") + ": 1"), " ", data.L = $("button", {
        onclick: function () {
          randomColor(data, 1, 11);
        }
      }, "2"), " ", data.M = $("button", {
        onclick: function () {
          randomColor(data, 1, 12);
        }
      }, "3"));
    }
    return watchable;
  }
  /**
   * @param {!Object} data
   * @param {!Object} options
   * @return {undefined}
   */
  function apply(data, options) {
    if (data.a.$c || data.a.cf || data.a.ne) {
      data.ka = new Layer(options, {
        className: "trqcont lh1s",
        style: {
          position: "relative"
        }
      });
      data.ka.add(0, {
        className: "nowrel"
      }, bind(data));
      data.ka.add(1, {
        className: "trqans dsp1 nowrel",
        style: {
          display: "none"
        }
      }, [$("button", {
        className: "minw",
        onclick: function () {
          randomColor(data, 2, data.Ya);
          loop(data, 0);
        }
      }, data.j("bl_yes")), " ", $("button", {
        className: "minw",
        onclick: function () {
          randomColor(data, 3, data.Ya);
          loop(data, 0);
        }
      }, data.j("bl_no")), " ", data.ua = $("span", ["..."])]);
      data.ka.add(2, {
        className: "nowrel",
        style: {
          display: "none"
        }
      }, [$("button", {
        className: "minw",
        onclick: function () {
          randomColor(data, 4);
          data.ka.show(0);
        }
      }, data.j("bl_yes")), " ", $("button", {
        className: "minw",
        onclick: function () {
          data.ka.show(0);
        }
      }, data.j("bl_no")), " ", $("span", {
        className: "ttup"
      }, data.j("bl_resign"))]);
    }
  }
  /**
   * @param {!Object} data
   * @param {string} content
   * @return {undefined}
   */
  function setup(data, content) {
    /**
     * @return {undefined}
     */
    function render() {
      /**
       * @return {undefined}
       */
      function callback() {
        /** @type {number} */
        var selected = 0;
        hooksWithName.forEach(function (radioItem) {
          if (radioItem.checked) {
            selected++;
          }
        });
        /** @type {boolean} */
        button.disabled = !selected;
      }
      /**
       * @return {undefined}
       */
      function update() {
        var text = hooksWithName.filter(function (radioItem) {
          return radioItem.checked;
        }).map(function (select_ele) {
          return select_ele.value;
        });
        if (0 < text.length) {
          this.send([96, this.F], text);
        }
      }
      var _this = this;
      if (!this.rd) {
        this.rd = create(this.app, this.j("t_rpin") || "-", {
          width: "80%",
          minHeight: "0",
          minWidth: "280px",
          maxWidth: "600px"
        }, {
          nopad: true
        });
      }
      var button;
      var hooksWithName = k.map(function (command_module_id, premadeCommentListId) {
        return $("input", {
          type: "checkbox",
          value: command_module_id,
          id: "_chrep" + premadeCommentListId,
          onchange: function () {
            return callback();
          }
        });
      });
      resolve(this.rd.Ma, [$("div", {
        className: "bsbb",
        style: {
          width: "100%",
          height: "220px",
          padding: "0 15px",
          borderTop: "solid 1px #ddd",
          borderBottom: "solid 1px #ddd",
          overflowY: "scroll"
        }
      }, 0 == hooksWithName.length ? $("p", "-") : hooksWithName.map(function (props) {
        return $("p", {}, [props, " ", $("label", {
          htmlFor: props.id
        }, props.value)]);
      })), $("p", {
        className: "bsbb",
        style: {
          padding: "0 15px"
        }
      }, button = $("button", {
        disabled: true,
        className: "minw",
        onclick: function () {
          success(_this.app);
          update.call(_this);
        }
      }, this.j("bl_ok")))]);
      filter(this.app, this.rd);
    }
    var target;
    var f;
    var scope = (window.k2prechat || "").split(" ");
    content = $(content, "div", {
      style: {
        position: "absolute",
        width: "100%",
        height: "100%"
      }
    }, [target = $("div", {
      className: "bsbb mb1s",
      style: {
        position: "absolute",
        top: 0,
        bottom: 0,
        left: 0,
        right: 0,
        background: "#fff",
        padding: "2px 4px 3px 8px"
      }
    }), $("div", {
      className: "h1s",
      style: {
        position: "absolute",
        bottom: 0,
        left: 0,
        right: 0,
        paddingTop: "4px"
      }
    }, [$("form", {
      onsubmit: function () {
        data.X(f.value.trim());
        /** @type {string} */
        f.value = "";
        return false;
      }
    }, f = $("input", {
      className: "bsbb",
      name: "somename",
      type: "text",
      autocomplete: "off",
      autocapitalize: "off",
      style: {
        width: "100%",
        border: "none"
      }
    }))]), $("div", {
      className: "bsbb",
      style: {
        position: "absolute",
        top: "100%",
        width: "100%",
        margin: "-2px 0 4px"
      }
    }, [$("button", {
      className: "ddbut butlh",
      onclick: function (event) {
        if (!data.app.G) {
          event.target.focus();
        }
      },
      style: {
        position: "absolute",
        top: 0,
        background: "#f8f8f8",
        border: "none",
        borderRadius: 0
      }
    }, "..."), $("div", {
      className: "ddcont bsbb bs dsp1",
      onmousedown: function (event) {
        if (data.app.lf && event.target.onclick) {
          event.target.onclick();
        }
      },
      ontouchend: function (event) {
        if (data.app.G) {
          if (event.target.onclick) {
            event.target.onclick();
          }
          return false;
        }
      },
      style: {
        position: "absolute",
        bottom: 0,
        width: "100%",
        background: "#f8f8f8",
        paddingTop: "1em"
      }
    }, [$("p", {}, [$("button", {
      onclick: function () {
        return render.call(data);
      },
      style: {
        color: "red",
        borderColor: "red"
      }
    }, data.j("t_rpin"))]), $("p", {}, [$("input", {
      type: "checkbox",
      checked: true,
      ontouchend: function (e) {
        if (data.app.G) {
          /** @type {boolean} */
          e.target.checked = !e.target.checked;
          e.target.onchange(e);
        }
      },
      onchange: function (b) {
        b = b.target.checked;
        data.X(b ? "/chat1" : "/chat0");
        if (!b) {
          data.ta.reset();
        }
      }
    }), " ", data.j("sw_chat")]), $("p", {}, "\ud83d\ude00 \ud83d\ude02 \u2639\ufe0f \ud83e\udd14 \ud83d\ude2d \ud83d\ude34 \ud83e\udd10 \ud83d\udc4d \ud83d\udc4e".split(" ").map(function (version) {
      return $("span", {
        className: "emo",
        style: {
          cursor: "pointer",
          marginRight: "1px"
        },
        onclick: function () {
          return data.X(version);
        }
      }, version);
    })), scope ? $("p", {
      style: {}
    }, scope.map(function (version) {
      return $("span", {
        style: {
          cursor: "pointer",
          marginRight: ".5em",
          padding: ".4em 0"
        },
        onclick: function () {
          return data.X(version);
        }
      }, version);
    })) : null])])]);
    extend(target, {
      wordWrap: "break-word",
      overflowY: "scroll",
      WebkitOverflowScrolling: "touch"
    });
    /**
     * @param {!Event} event
     * @return {?}
     */
    target.onselectstart = function (event) {
      event.stopPropagation();
      return true;
    };
    /** @type {!Array} */
    var k = [];
    expect(data.app, "tabchat", function (a) {
      if (0 != a.indexOf(element)) {
        k.push(a);
      }
    });
    expect(data.app, "tabopen", function (a) {
      if (0 == a) {
        /** @type {!Array} */
        k = [];
      }
    });
    var iConfig = data.a.Cf || 0;
    data.ta = {
      f: content,
      append: function (a) {
        /** @type {boolean} */
        var c = target.scrollTop + 2 >= target.scrollHeight - target.clientHeight;
        for (; 0 < iConfig && target.children.length > iConfig && target.firstChild;) {
          target.removeChild(target.firstChild);
        }
        var b = 0 != a.indexOf(element) ? a.indexOf(":") : 0;
        a = $(target, "div", {
          className: "tind"
        }, 0 < b ? [$("b", [a.substring(0, b)]), a.substring(b)] : a);
        write(a);
        if (c) {
          /** @type {number} */
          target.scrollTop = target.scrollHeight - target.clientHeight + 1;
        }
      },
      reset: function () {
        /** @type {string} */
        target.innerHTML = "";
        /** @type {string} */
        f.value = "";
      },
      Aa: function () {
        /** @type {number} */
        target.scrollTop = target.scrollHeight - target.clientHeight + 1;
      }
    };
  }
  /**
   * @param {!Object} self
   * @param {?} state
   * @return {undefined}
   */
  function activate(self, state) {
    var type = $(state.pb, "div", {
      style: {
        position: "absolute",
        width: "100%",
        height: "100%"
      }
    });
    self.cb = new Color(self.app, type, {
      cols: [F, -1],
      ud: true,
      Pc: function (b, act) {
        if (1 != act.Ra) {
          self.app.send([action], [b.name]);
        } else {
          ok(act, 0);
          self.X("/boot " + b.name);
        }
      },
      Ie: {
        className: "mb1s ulwp ovysct"
      }
    });
    state.add(self.j("sw_users"), type);
    $(type, "div", {
      className: "lh1s nowrel",
      style: {
        position: "absolute",
        left: 0,
        right: 0,
        bottom: 0
      }
    }, [self.Mc = $("button", {
      className: "minw",
      onclick: function () {
        createServer(self);
      }
    }, self.j("bl_invite")), " ", self.Kc = $("button", {
      className: "minw",
      onclick: function () {
        var cb = self.cb;
        var el = cb.Ra;
        if (!(0 != el && 1 != el)) {
          ok(cb, 1 - cb.Ra);
        }
      }
    }, self.j("bl_boot")), null, " ", self.W = $("span", {
      className: "mlh"
    }, ["-"])]);
  }
  /**
   * @param {!Object} t
   * @param {!Object} s
   * @param {!Object} r
   * @return {undefined}
   */
  function reset(t, s, r) {
    s = t.u.Dd(s, r);
    if (t.history && s) {
      if (Array.isArray(s)) {
        add(t.history, s);
      } else {
        if (r = t.history, 0 != r.a.rows.length) {
          var undelete_backup = r.a.rows[r.a.rows.length - 1];
          if (0 != undelete_backup.cells.length) {
            resolve(undelete_backup.cells[undelete_backup.cells.length - 1], [serialize(s)]);
            if (r.o.Va) {
              r.Aa(0);
            }
          }
        }
      }
      t.u.qb(t.history.b);
    }
  }
  /**
   * @param {!Object} obj
   * @param {!Object} args
   * @param {!Object} cb
   * @return {undefined}
   */
  function save(obj, args, cb) {
    cb = obj.Ee(args, cb);
    obj.u.history(args, cb);
    if (obj.history && cb) {
      args = obj.history;
      /** @type {number} */
      args.g = 0;
      /** @type {number} */
      args.b = -1;
      /** @type {number} */
      var i = args.a.rows.length - 1;
      for (; 0 <= i; i--) {
        args.a.deleteRow(i);
      }
      add(args, cb);
      obj.u.qb(obj.history.b);
    }
  }
  /**
   * @param {!Object} a
   * @param {?} b
   * @return {undefined}
   */
  function close(a, b) {
    if (!a.u.wb) {
      testCircleCircle(a.history, b);
      a.u.qb(a.history.b);
    }
  }
  /**
   * @param {!Object} self
   * @param {?} state
   * @return {undefined}
   */
  function check(self, state) {
    var c = $(state.pb, "div", {
      className: "bsbb dsp1",
      style: {
        background: "#fff",
        position: "absolute",
        top: "0",
        width: "100%",
        height: "100%"
      }
    });
    var node = $(c, "div", {
      style: {
        width: "50%",
        cssFloat: "left",
        marginTop: ".75em"
      }
    });
    var wrapper = $(node, "div", {
      className: "mbsp"
    });
    /** @type {!Array} */
    var navLinksArr = [];
    navLinksArr.push(self.j("tb_ttpub"));
    if (self.app.W) {
      /** @type {number} */
      var k = 0;
      for (; 7 > k; k++) {
        navLinksArr.push(exec(self.app, 0 == k % 2 ? join(self.app, 1 + (k >> 1)) : Math.floor((join(self.app, 1 + (k >> 1)) + join(self.app, (k >> 1) + 2)) / 2)) + "+");
      }
    }
    navLinksArr.push(self.j("tb_ttprv"));
    self.S = $(wrapper, "select", {
      onchange: function () {
        var lastSelected = this.selectedIndex;
        setTimeout(self, "ttype", self.app.W && 0 < lastSelected ? 8 <= lastSelected ? 2 : lastSelected + 2 : 2 * lastSelected);
      }
    }, navLinksArr.map(function (mei) {
      return $("option", mei);
    }));
    wrapper = $(node, "div", {
      className: "mbsp"
    });
    $(wrapper, "div", {}, self.j("tb_tr_game"));
    self.Xa.push(self.G = $(wrapper, "select", {
      onchange: function () {
        try {
          setTimeout(self, "tg", self.a.Ef ? -(this.selectedIndex + 1) : this.options[this.selectedIndex].text);
        } catch (m) {
        }
      }
    }, self.app.Hd.map(function (mei) {
      return $("option", mei);
    })));
    if (self.a.bf) {
      self.Xa.push(self.A = $(wrapper, "select", {
        onchange: function () {
          try {
            setTimeout(self, "tm", this.options[this.selectedIndex].text);
          } catch (m) {
          }
        }
      }, self.app.Md.map(function (mei) {
        return $("option", mei);
      })));
    }
    if (self.a.$c && "undefined" == typeof self.a.Gf) {
      $(node, "div", {
        className: "mbsp"
      }, [self.C = $("input", {
        type: "checkbox",
        onchange: function () {
          setTimeout(self, "ud", this.checked ? 1 : 0);
        }
      }), self.j("tb_noundo")]);
      self.Xa.push(self.C);
    }
    $(node, "div", {}, [$("input", {
      type: "checkbox",
      checked: self.app.qa,
      onchange: function () {
        var original = this.checked;
        var input = self.app;
        input.qa = original;
        clear(input);
        if (original) {
          play(self.app.qa);
        }
      }
    }), self.j("p_bp")]);
    node = $(c, "div", {
      style: {
        width: "50%",
        cssFloat: "right",
        marginTop: ".75em"
      }
    });
    $(node, "div", {
      className: "mbsp nowrel"
    }, [self.g = $("input", {
      type: "checkbox",
      onchange: function () {
        setTimeout(self, "gtype", this.checked ? 0 : 1);
      }
    }), self.j("tb_gtnrt") + " (x)"]);
    self.Ge(node);
    if (!self.app.P && self.history) {
      $(node, "div", $("button", {
        onclick: function () {
          self.b.show(self.oa);
        }
      }, [self.j("sw_history")]));
    }
    state.add(self.j("sw_setup"), c, true);
  }
  /**
   * @param {!Object} e
   * @param {number} t
   * @param {!Object} color
   * @return {undefined}
   */
  function validate(e, t, color) {
    /** @type {number} */
    var i = 0;
    for (; i < color.length - 2; i++) {
      /** @type {number} */
      var value = parseInt(color[i + 2], 10);
      if ("ttype" == t[i]) {
        /** @type {number} */
        e.S.selectedIndex = e.app.W && 1 < value ? 2 == value ? 8 : value - 2 : value >> 1;
      } else {
        if ("gtype" == t[i]) {
          if (!e.Tb) {
            /** @type {boolean} */
            e.g.checked = 0 == value;
          }
        } else {
          if ("tm" == t[i]) {
            if (e.A) {
              map(e.A, value.toString());
            }
          } else {
            if ("tg" == t[i]) {
              if (e.G) {
                if (0 > value) {
                  /** @type {number} */
                  e.G.selectedIndex = -value - 1;
                } else {
                  map(e.G, value.toString());
                }
              }
            } else {
              if ("ud" == t[i]) {
                if (e.C) {
                  /** @type {boolean} */
                  e.C.checked = 0 != value;
                }
              } else {
                if ("tch" != t[i]) {
                  if ("op:" == t[i].substring(0, 3)) {
                    if (e.W) {
                      resolve(e.W, ["op: " + t[i].substring(3)]);
                    }
                  } else {
                    e.Fe(t[i], value);
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  /**
   * @param {!Object} params
   * @return {?}
   */
  function upload(params) {
    return 0 >= params.F ? "-" : toString(params.j("l_tab"), params.F.toString()) + " \u00a0 " + params.ma.O[0];
  }
  /**
   * @param {!Object} args
   * @return {undefined}
   */
  function copy(args) {
    if (!(isString(args) && 0 != (args.la & 64) && (isObject(args) && 0 != (args.U & 2) || 0 != (args.la & 16) && 0 == (args.U & 16)))) {
      args.app.send([73, args.F], null);
      cb(args.app, "tabopen", 0);
      then(args.app);
      /** @type {number} */
      args.F = -1;
      setTimeout(function () {
        if (-1 == args.F) {
          args.reset();
        }
      }, 50);
    }
  }
  /**
   * @param {!Object} data
   * @param {!Array} options
   * @param {!Object} parent
   * @return {undefined}
   */
  function run(data, options, parent) {
    if (0 < data.F) {
      data.reset();
    }
    if (2 < options.length) {
      /** @type {number} */
      data.Yc = Math.floor(options[2] / 16) % 3;
      /** @type {boolean} */
      data.Tb = 0 != data.Yc;
      /** @type {boolean} */
      data.g.checked = 1 == data.Yc;
      /** @type {boolean} */
      data.g.disabled = data.Tb;
    }
    data.ma = new Node(data.app, options, parent);
    data.F = data.ma.J[1];
    cb(data.app, "tabopen", data.F);
    /** @type {boolean} */
    data.kc = true;
    if (window.k2spectm) {
      data.ta.append(element + window.k2spectm);
    }
    if (data.app.P && data.wa) {
      data.ta.append(element + "(info) " + data.wa);
    }
  }
  /**
   * @param {!Object} obj
   * @param {boolean} name
   * @return {undefined}
   */
  function _(obj, name) {
    callback(obj.f, "sbdropvis", name);
  }
  /**
   * @param {!Object} opts
   * @param {number} d
   * @param {number} i
   * @param {number} t
   * @return {undefined}
   */
  function runTest(opts, d, i, t) {
    if (1 == i) {
      opts = opts.T;
      if (opts.R) {
        opts.R[d] = t.toString();
      }
    } else {
      if (2 == i) {
        equal(opts.T, d, t);
      }
    }
  }
  /**
   * @param {!Object} data
   * @param {number} i
   * @return {undefined}
   */
  function moveTo(data, i) {
    if (i != data.Me) {
      /** @type {number} */
      data.Me = i;
      if (8 != i) {
        if (isObject(data)) {
          i = 9 == i ? data.j("win_draw") : 0 > i ? toString(data.j("los_pln"), (-i).toString()) : toString(data.j("win_pln"), (i + 1).toString());
          /** @type {number} */
          data.vc = i;
          data.ta.append(element + i);
        }
      } else {
        /** @type {null} */
        data.vc = null;
      }
    }
  }
  /**
   * @param {!Object} e
   * @param {!Object} container
   * @return {undefined}
   */
  function anonymous(e, container) {
    var options = e.T;
    var A = $("div", {
      className: "tplcont",
      style: {
        overflowY: "auto"
      }
    });
    /** @type {number} */
    var i = 0;
    for (; i < options.ea; i++) {
      var node = $(A, "div", {
        style: {
          cssFloat: 0 == i % 2 ? "left" : "right",
          width: "49.5%",
          overflowX: "hidden",
          marginTop: 2 <= i ? e.app.P ? "8px" : "4px" : 0
        }
      });
      var show = $(node, "div", {
        className: "f12",
        style: {
          verticalAlign: "middle",
          lineHeight: "12px",
          background: "rgba(0,0,0,0.8)",
          color: "rgba(255,255,255,0.95)",
          fontWeight: "bold",
          padding: "0 5px"
        }
      });
      if (options.Sa) {
        $(show, "div", {
          style: {
            display: "inline-block",
            width: "7px",
            height: "7px",
            background: options.Sa[(i + (options.bd ? 1 : 0)) % options.ea],
            marginRight: "5px"
          }
        });
      }
      $(show, ["#" + (i + 1)]);
      if (!e.a.od && options.R) {
        $(show, "div", {
          style: {
            display: "inline-block",
            marginLeft: "4px",
            width: 0,
            height: 0,
            borderLeft: "solid 4px transparent",
            borderRight: "solid 4px transparent",
            borderTop: "solid 7px rgba(255,255,255,0.9)",
            visibility: e.ga == i ? "inherit" : "hidden"
          }
        });
      }
      var overlay = $(node, "div", {
        style: {
          position: "relative"
        }
      });
      var ANCHOR = options.a[i];
      /** @type {boolean} */
      show = 0 == ANCHOR;
      $(overlay, "button", {
        className: "butsys butsit",
        disabled: 1 != ANCHOR,
        style: {
          display: show ? "none" : "block",
          position: "absolute",
          width: "100%",
          height: "100%"
        },
        onclick: function (b) {
          return function () {
            e.send([83, e.F, b], null);
          };
        }(i)
      }, ["#" + (i + 1)]);
      overlay = $(overlay, "div", {
        style: {
          background: "#fff",
          padding: "6px 6px 6px 0",
          visibility: show ? "inherit" : "hidden"
        }
      });
      $(overlay, "button", {
        style: {
          display: 0 != (options.pc & 1 << i) && null != options.name[i] ? "block" : "none",
          cssFloat: "right",
          border: 0,
          padding: "2px 6px",
          fontWeight: "bold",
          margin: "1px 0",
          background: "#bbb",
          color: "#fff"
        },
        onclick: function (b) {
          return function () {
            e.send([84, e.F, b], null);
          };
        }(i)
      }, ["X"]);
      $(overlay, "div", {
        className: "nowrel",
        style: {
          fontSize: "115%",
          color: options.focus[i] ? "inherit" : "#aaa",
          padding: "2px 6px"
        }
      }, [show ? options.name[i] || "--" : "-"]);
      node = $(node, "div", {
        className: "tplext",
        style: {
          marginTop: "6px",
          width: "100%"
        }
      });
      if (options.R && options.Eb) {
        $(node, "div", {
          style: {
            display: "table-cell",
            width: "52%",
            textAlign: "center"
          }
        }, $("div", {
          style: {
            background: "#fff",
            opacity: ".7",
            fontWeight: "bold",
            padding: "1px 0 2px"
          }
        }, [options.R[i]]));
      }
      if (options.time) {
        $(node, "div", {
          style: {
            display: "table-cell",
            verticalAlign: "middle",
            textAlign: options.Eb ? "center" : "left"
          }
        }, [options.time[i], e.a.od || options.R ? null : $("div", {
          style: {
            marginLeft: "8px",
            display: "inline-block",
            width: 0,
            height: 0,
            borderLeft: "4px solid transparent",
            borderRight: "4px solid transparent",
            borderBottom: "solid 8px rgba(0,0,0,0.8)",
            visibility: i == e.ga ? "inherit" : "hidden"
          }
        })]);
      }
      if (!options.time) {
        $(node, "div", {
          style: {
            display: "table-cell"
          }
        });
      }
    }
    resolve(container, A);
  }
  /**
   * @param {!Object} m
   * @return {undefined}
   */
  function slice(m) {
    cb(m.app, "tabplayers");
  }
  /**
   * @param {!Object} data
   * @return {undefined}
   */
  function clone(data) {
    /** @type {number} */
    var offset = 0;
    for (; offset < data.T.ea; offset++) {
      if (offset >= data.ma.nd) {
        var T = data.T;
        /** @type {number} */
        var i = offset;
        /** @type {null} */
        T.name[i] = null;
        /** @type {number} */
        T.a[i] = 0;
      } else {
        if (T = data.ma.J[base + offset], 1 == T) {
          T = data.T;
          /** @type {number} */
          i = offset;
          T.name[i] = data.ma.O[position + offset];
          /** @type {number} */
          T.a[i] = 0;
        } else {
          if (3 == T) {
            T = data.T;
            /** @type {number} */
            i = offset;
            /** @type {null} */
            T.name[i] = null;
            /** @type {number} */
            T.a[i] = 0;
          } else {
            i = data.T;
            /** @type {number} */
            var k = offset;
            /** @type {number} */
            i.a[k] = 0 == T && inject(data) ? 1 : -1;
            /** @type {string} */
            i.name[k] = "";
          }
        }
      }
    }
    constructor(data);
    slice(data);
  }
  /**
   * @param {!Object} data
   * @param {boolean} name
   * @param {!Array} options
   * @return {undefined}
   */
  function on(data, name, options) {
    getType(data.ma, name, options);
    cb(data.app, "tabstatus");
    clone(data);
    if (data.Qa) {
      isArray(data.u);
    }
  }
  /**
   * @param {!Object} t
   * @param {string} x
   * @param {number} s
   * @return {undefined}
   */
  function end(t, x, s) {
    t.u.setActive(x, true);
    if (s) {
      play(t.app.qa);
    }
    x = x || t.Hc;
    cb(t.app, "tabalert", x);
    t = t.app;
    /** @type {(null|string)} */
    x = x ? "\u25bc" : null;
    /** @type {boolean} */
    s = "hasFocus" in document ? document.hasFocus() : true;
    if (!(null != x && (null == x || s && t.tab && t.tab.f == t.C) || t.Lb == x)) {
      /** @type {string} */
      t.Lb = x;
      emit(t);
    }
  }
  /**
   * @param {!Object} obj
   * @return {undefined}
   */
  function stringify(obj) {
    var i = obj.ga;
    var c = obj.Za;
    /** @type {number} */
    var brick = 0;
    if (!(isObject(obj) && 0 != (obj.U & 2) || 0 != (obj.la & 16) && 0 == (obj.U & 16))) {
      /** @type {number} */
      brick = 0 == (obj.la & 1) || isObject(obj) && 0 == (obj.U & 8) ? isString(obj) ? 1 << c : 0 : 15;
    }
    /** @type {number} */
    obj.T.pc = brick;
    brick = isString(obj) && isObject(obj) && 0 == (obj.U & 8);
    if (obj.Nc) {
      /** @type {boolean} */
      obj.Nc.disabled = !brick;
    }
    if (obj.I) {
      /** @type {boolean} */
      obj.I.disabled = !brick;
    }
    if (obj.L) {
      /** @type {boolean} */
      obj.L.disabled = !brick;
    }
    if (obj.M) {
      /** @type {boolean} */
      obj.M.disabled = !brick;
    }
    brick = isString(obj) && isObject(obj) && 0 == (obj.U & 4) && 0 == (obj.U & 8);
    if (obj.Lc) {
      /** @type {boolean} */
      obj.Lc.disabled = !(brick && i != c);
    }
    if (obj.Oc) {
      /** @type {boolean} */
      obj.Oc.disabled = !(brick && !obj.Je);
    }
    /** @type {boolean} */
    var canBuild = 0 != (obj.la & 1);
    /** @type {boolean} */
    obj.S.disabled = !canBuild;
    if (obj.Mc) {
      /** @type {boolean} */
      obj.Mc.disabled = !canBuild;
    }
    if (obj.Kc) {
      /** @type {boolean} */
      obj.Kc.disabled = !canBuild;
    }
    /** @type {boolean} */
    canBuild = canBuild && !isObject(obj);
    if (!obj.Tb) {
      /** @type {boolean} */
      obj.g.disabled = !canBuild;
    }
    /** @type {number} */
    var j = 0;
    for (; j < obj.Xa.length; j++) {
      /** @type {boolean} */
      obj.Xa[j].disabled = !canBuild;
    }
    clone(obj);
    i = brick && i == c && -1 != i;
    end(obj, i, i && !obj.Vd);
    obj.Vd = i;
    isArray(obj.u);
  }
  /**
   * @param {!Object} self
   * @param {number} data
   * @param {number} s
   * @param {number} i
   * @return {?}
   */
  function finish(self, data, s, i) {
    if (1 == data[s] || 2 == data[s]) {
      /** @type {boolean} */
      var v = 2 == data[s];
      if (s + (3 + (v ? 1 : 0)) > i) {
        return s;
      }
      i = data[s + 1];
      /** @type {number} */
      var _channel = data[s + 2] % 2;
      var action = self.ga;
      if (0 > i) {
        i = i + 65536;
      }
      /** @type {(null|string)} */
      self.sb = -1 > data[s + 2] ? "0:00" : null;
      if (-1 > data[s + 2]) {
        /** @type {number} */
        action = -data[s + 2] >> 2;
      }
      self = self.bb;
      data = v ? data[s + 3] : 0;
      /** @type {number} */
      self.g = Date.now();
      self.a = action;
      /** @type {number} */
      self.I = i;
      self.b = "undefined" != typeof data ? data : 0;
      /** @type {number} */
      self.C = _channel;
      if (!self.H) {
        self.start();
      }
      if (0 < self.b) {
        /** @type {number} */
        self.G = -1;
        verify(self);
      }
      s = s + (3 + (v ? 1 : 0));
    } else {
      if (3 == data[s]) {
        if (s + 2 > i) {
          return s;
        }
        /** @type {boolean} */
        self.Je = 0 != (data[s + 1] & 1);
        /** @type {boolean} */
        v = 0 != (data[s + 1] & 2);
        i = self.T;
        if (i.bd != v) {
          /** @type {boolean} */
          i.bd = v;
        }
        i = self.u;
        if (i.Qc != v) {
          /** @type {boolean} */
          i.Qc = v;
        }
        /** @type {number} */
        v = data[s + 1] >> 3;
        /** @type {number} */
        data = 0;
        for (; data < self.T.ea; data++) {
          /** @type {boolean} */
          self.T.focus[data] = 0 == (v & 1 << data);
        }
        s = s + 2;
      }
    }
    return s;
  }
  /**
   * @param {!Object} data
   * @param {!Object} result
   * @param {number} value
   * @return {undefined}
   */
  function render(data, result, value) {
    if (!(2 > result.length)) {
      switch (result[0]) {
        case 92: // lk16:92 recv move i:[92,table ID,move] s:absent move=(white?64:0)+(8*row)+column
          reset(data, result, value);
          break;
        case 91: // lk16:91 recv move history i:[91,table ID, many moves]
          /** @type {boolean} */
          data.kc = false;
          save(data, result.slice(2), value);
          break;
        case 88: // lk16:88 recv ??? [88,table ID,?, ?, ?]
          if (5 > result.length) {
            break;
          }
          var index = data.Za;
          data.la = result[2];
          data.U = result[3];
          data.Za = result[4];
          data.u.qe = data.Za;
          data.Hc = isString(data) && 0 != (data.U & 8) && 0 == (data.la & 8) && 0 == (data.U & 16);
          stringify(data);
          if (data.Za != index && data.history && !setItem(data.history)) {
            add(data.history, []);
            data.u.qb(data.history.b);
          }
          break;
        case 90: // lk16:90 recv ??? i:[90,table ID, 15x ? ] s:absent
          if (3 > result.length) {
            break;
          }
          if (result.length < result[2] + 4) {
            break;
          }
          if (data.bb) {
            /** @type {number} */
            data.bb.a = -1;
          }
          value = result[3 + result[2]];
          if (0 == value) {
            break;
          }
          var p = result[2] + 4;
          var name = p + value;
          /** @type {number} */
          var expected = 0;
          if (result.length < p + value) {
            break;
          }
          for (; result.length >= name + value;) {
            /** @type {number} */
            index = 0;
            for (; index < value; index++) {
              runTest(data, expected, result[p + index], result[name + index]);
            }
            name = name + value;
            expected++;
          }
          if (0 < result[2]) {
            data.ga = result[3];
          }
          if (1 < result[2]) {
            moveTo(data, result[4]);
          }
          /** @type {null} */
          data.sb = null;
          index = 3 + result[2];
          /** @type {number} */
          value = 5;
          for (; value < index;) {
            p = 5 > result[value] ? finish(data, result, value, index) : data.ie(result, value, index);
            if (value == p) {
              break;
            }
            value = p;
          }
          stringify(data);
          break;
        case 93:
          if (3 > result.length) {
            break;
          }
          if (0 < value.length) {
            loop(data, result[2], value);
          } else {
            loop(data, 0);
          }
          break;
        case 81: // lk16:81 recv incoming chatmessage i:[81,table ID] s:[chat message]
          if (1 > value.length) {
            break;
          }
          cb(data.app, "tabchat", value[0]);
          data.ta.append(value[0]);
          if (data.b && 2 < result.length) {
            data.b.show(0);
          }
          break;
        case 89: // lk16:89 recv table settings i:[89, table ID, many setting values] s:[many setting names]
          if (2 > result.length || value.length < result.length - 2) {
            break;
          }
          validate(data, value, result);
          break;
        case 87:
          if (3 > result.length) {
            break;
          }
          /** @type {boolean} */
          data.T.focus[result[2] >> 1] = 0 != (result[2] & 1);
          slice(data);
          break;
        case 94:
          if (1 > value.length) {
            break;
          }
          if (!data.B) {
            data.B = create(data.app, data.a.Wb || "-", {
              width: "80%",
              minHeight: "0",
              minWidth: "280px",
              maxWidth: "600px"
            }, {
              nopad: true
            });
            /**
             * @param {!Event} event
             * @return {?}
             */
            data.B.Ma.onselectstart = function (event) {
              event.stopPropagation();
              return true;
            };
          }
          resolve(data.B.Ma, $("textarea", {
            value: value[0],
            className: "bsbb taplain",
            style: {
              width: "100%",
              height: "280px",
              borderTop: "solid 1px #ddd",
              padding: "4px 15px",
              fontFamily: "monospace"
            },
            spellcheck: false,
            readOnly: true
          }));
          filter(data.app, data.B, null, {
            okselect: true
          });
          break;
        case 84:
          if (1 > value.length) {
            break;
          }
          result = process(data.app, value[0]);
          if (null != result) {
            matches(data.cb, result);
          }
          if (!(0 == (data.la & 1) && !isString(data) || isObject(data))) {
            play(data.app.qa);
          }
          break;
        case 85:
          if (1 > value.length) {
            break;
          }
          result = process(data.app, value[0]);
          if (null != result) {
            complete(data.cb, result);
          }
          break;
        case 86: // lk16:86 recv ??? s:[86,table ID] i:[player, player]
          p = {};
          /** @type {number} */
          index = 0;
          for (; index < value.length; index++) {
            result = process(data.app, value[index]);
            if (!(null == result || p.hasOwnProperty(result.name))) {
              /** @type {!Object} */
              p[result.name] = result;
            }
          }
          func(data.cb, p);
      }
    }
  }
  /**
   * @param {!Object} params
   * @return {undefined}
   */
  function constructor(params) {
    /** @type {null} */
    var value = null;
    /** @type {boolean} */
    var onFulfillment = false;
    /** @type {boolean} */
    var d = false;
    if (0 != (params.U & 16)) {
      /** @type {string} */
      value = "aw_rnd";
    } else {
      if (0 != (params.U & 4)) {
        if (inject(params)) {
          /** @type {string} */
          value = "pr_sel";
        } else {
          if (isObject(params)) {
            /** @type {string} */
            value = "aw_pls";
          } else {
            /** @type {string} */
            value = "aw_opp";
            /** @type {boolean} */
            d = true;
          }
        }
      } else {
        if (0 != (params.U & 8)) {
          if (isString(params) && 0 == (params.la & 8)) {
            /** @type {string} */
            value = "";
            /** @type {boolean} */
            onFulfillment = true;
          } else {
            /** @type {string} */
            value = "aw_go";
          }
        }
      }
    }
    if (null != value) {
      value = "" != value ? params.j(value) : "-";
      var val;
      if (d) {
        if (pow(params.ma)) {
          if (0 < (val = translate(params.ma))) {
            /** @type {string} */
            value = value + (" (" + exec(params.app, val) + "+)");
          }
        } else {
          /** @type {string} */
          value = value + (" (" + params.j("bl_invite") + ")");
        }
      }
      if (null != params.sb) {
        /** @type {string} */
        value = "(" + params.sb + ")" + ("" != value && "-" != value ? " " + value : "");
      }
      subscribe(params, value, onFulfillment);
    } else {
      subscribe(params, null);
    }
  }
  /**
   * @param {!Object} a
   * @param {number} c
   * @param {string} b
   * @return {?}
   */
  function g(a, c, b) {
    switch (c) {
      case 1:
        return a.j("bl_draw");
      case 2:
        return a.j("bl_undo");
      case 4:
        return -1 != b.indexOf("/") ? b : a.j("bl_tram");
      case 10:
        return a.j("bl_resign") + " 1";
      case 11:
        return a.j("bl_resign") + " 2";
    }
    return "(?)";
  }
  /**
   * @param {number} i
   * @param {string} n
   * @param {?} fs
   * @return {undefined}
   */
  function Node(i, n, fs) {
    /** @type {number} */
    this.app = i;
    n[2] %= 16;
    /** @type {string} */
    this.J = n;
    this.O = fs;
    /** @type {number} */
    i = this.O.length - 1;
    for (; 1 <= i; i--) {
      n = this.O[i].indexOf("/");
      if (0 <= n) {
        this.O[i] = this.O[i].substring(0, n);
      }
    }
    this.F = this.J[1].toString();
    /** @type {number} */
    this.nd = this.J.length - base;
  }
  /**
   * @param {!Object} args
   * @return {?}
   */
  function translate(args) {
    /** @type {number} */
    var from = 1 + Math.floor((args.J[2] - 3) / 2);
    return 2 < args.J[2] ? 0 == (args.J[2] - 3) % 2 ? join(args.app, from) : join(args.app, from) + join(args.app, from + 1) >> 1 : 0;
  }
  /**
   * @param {!Object} data
   * @return {?}
   */
  function Number(data) {
    /** @type {number} */
    var buckets = 0;
    /** @type {number} */
    var counter = base;
    for (; counter < data.J.length; counter++) {
      if (0 == data.J[counter]) {
        buckets++;
      }
    }
    return buckets;
  }
  /**
   * @param {!Object} node
   * @return {?}
   */
  function td(node) {
    /** @type {number} */
    var self = 0;
    /** @type {number} */
    var a = base;
    for (; a < node.J.length; a++) {
      if (!(1 != node.J[a] && 2 != node.J[a])) {
        self++;
      }
    }
    return self;
  }
  /**
   * @param {!Object} key
   * @param {?} data
   * @return {?}
   */
  function report(key, data) {
    return 0 == key.J[2] || 3 <= key.J[2] && data >= translate(key);
  }
  /**
   * @param {!Object} args
   * @return {?}
   */
  function pow(args) {
    return 0 == args.J[2] || 3 <= args.J[2];
  }
  /**
   * @param {?} args
   * @param {string} e
   * @param {number} key
   * @return {undefined}
   */
  function getType(args, e, key) {
    if (!(3 > e.length)) {
      e[2] %= 16;
      /** @type {string} */
      args.J = e;
      /** @type {number} */
      args.O = key;
      /** @type {number} */
      key = args.O.length - 1;
      for (; 1 <= key; key--) {
        var d = args.O[key].indexOf("/");
        if (0 <= d) {
          args.O[key] = args.O[key].substring(0, d);
        }
      }
      /** @type {number} */
      args.nd = e.length - base;
    }
  }
  /**
   * @param {?} msg
   * @param {string} id
   * @param {number} d
   * @param {!Object} data
   * @return {undefined}
   */
  function Connection(msg, id, d, data) {
    /** @type {number} */
    var index = 0;
    /** @type {string} */
    this.name = id;
    this.a = data[index++];
    this.A = msg.Cc && msg.Cc[Math.floor(this.a / DyMilli)] || "";
    this.b = 0 != (d & h) && index < data.length ? data[index++] : 0;
    this.R = 0 != (d & value) && index < data.length ? data[index++] : 0;
    this.g = 0 != (d & F) && index < data.length ? data[index++] : 0;
    /** @type {!Array} */
    this.Pa = [];
  }
  /**
   * @param {!Object} context
   * @param {string} num
   * @return {undefined}
   */
  function forEach(context, num) {
    num = context.Pa.indexOf(num);
    if (-1 != num) {
      context.Pa.splice(num, 1);
    }
  }
  /**
   * @param {!Object} type
   * @return {undefined}
   */
  function attr(type) {
    for (; 0 < type.Pa.length;) {
      complete(type.Pa[0], type);
    }
  }
  /**
   * @param {!Object} o
   * @param {?} type
   * @return {?}
   */
  function length(o, type) {
    switch (type) {
      case h:
        return o.b;
      case F:
        return o.g;
      case value:
        return o.R;
      case file:
        return 0 == o.b ? 1 : 0;
    }
    return 0;
  }
  /**
   * @param {!Object} e
   * @param {number} d
   * @param {!Object} data
   * @return {undefined}
   */
  function tick(e, d, data) {
    /** @type {number} */
    var i = 0;
    e.a = data[i++];
    e.b = 0 != (d & h) && i < data.length ? data[i++] : 0;
    e.R = 0 != (d & value) && i < data.length ? data[i++] : 0;
    e.g = 0 != (d & F) && i < data.length ? data[i++] : 0;
    /** @type {number} */
    d = e.Pa.length - 1;
    for (; 0 <= d; d--) {
      if (data = e.Pa[d], i = e, data.b.hasOwnProperty(i.name)) {
        var item = data.b[i.name];
        item.D = load(data, i, item.D);
        item = data.g.indexOf(i);
        if (-1 != item) {
          data.g.splice(item, 1);
          push(data, i, data.g.length);
        }
      }
    }
  }
  /**
   * @param {!Object} a
   * @param {!Object} b
   * @param {!Object} options
   * @return {undefined}
   */
  function Color(a, b, options) {
    /** @type {!Object} */
    this.app = a;
    /** @type {!Object} */
    this.o = options;
    /** @type {number} */
    this.Ra = 0;
    this.cols = options.cols;
    this.A = options.vf || 0;
    /** @type {boolean} */
    this.C = 0 == this.A;
    this.b = {};
    /** @type {!Array} */
    this.g = [];
    this.f = $(b, "div", options.Ie || {});
    this.a = $(this.f, "table", {
      className: "ul " + (options.Ue ? "uls2" : "uls1")
    });
    this.B = options.oe ? overlay(this, null) : null;
  }
  /**
   * @param {!Object} d
   * @param {number} i
   * @return {undefined}
   */
  function forward(d, i) {
    if (d.A != i) {
      /** @type {number} */
      d.A = i;
      /** @type {boolean} */
      d.C = 0 == i;
      /** @type {number} */
      i = 1;
      var patchLen = d.g.length;
      for (; i < patchLen; i++) {
        var descriptor = d.g[i];
        d.g.splice(i, 1);
        push(d, descriptor, i);
      }
    }
  }
  /**
   * @param {string} a
   * @param {number} b
   * @return {undefined}
   */
  function ok(a, b) {
    if (a.Ra != b) {
      /** @type {number} */
      a.Ra = b;
      callback(a.a, "ulm1", 1 == b);
    }
  }
  /**
   * @param {!Object} obj
   * @param {!Object} a
   * @param {number} i
   * @return {undefined}
   */
  function push(obj, a, i) {
    var options = obj.A;
    for (; 0 < i;) {
      var b = obj.g[i - 1];
      if (0 != options) {
        var l = length(b, options);
        var d = length(a, options);
        /** @type {number} */
        l = l < d ? -1 : l > d ? 1 : 0;
      } else {
        /** @type {number} */
        l = b.name < a.name ? -1 : b.name > a.name ? 1 : 0;
      }
      if (obj.C) {
        /** @type {number} */
        l = -l;
      }
      if (0 > l || 0 == l && 0 < (b.name < a.name ? -1 : b.name > a.name ? 1 : 0)) {
        i--;
      } else {
        break;
      }
    }
    options = (options = obj.g[i]) ? obj.b[options.name] : null;
    obj.g.splice(i, 0, a);
    obj = obj.b[a.name];
    if (obj.D) {
      obj = obj.D;
      obj.parentNode.insertBefore(obj, options ? options.D : null);
    }
  }
  /**
   * @param {!Object} result
   * @param {!Object} value
   * @return {undefined}
   */
  function matches(result, value) {
    if (!result.b.hasOwnProperty(value.name)) {
      result.b[value.name] = {
        Ua: value
      };
      result.b[value.name].D = load(result, value);
      push(result, value, result.g.length);
      value.Pa.push(result);
    }
  }
  /**
   * @param {string} obj
   * @param {!Object} source
   * @return {undefined}
   */
  function complete(obj, source) {
    if (obj.b.hasOwnProperty(source.name)) {
      forEach(source, obj);
      var data = obj.b[source.name];
      if (data.D && data.D.parentNode) {
        data = data.D;
        data.parentNode.removeChild(data);
      }
      data = obj.g.indexOf(source);
      if (-1 != data) {
        obj.g.splice(data, 1);
      }
      delete obj.b[source.name];
    }
  }
  /**
   * @param {!Object} a
   * @param {!Object} item
   * @return {undefined}
   */
  function func(a, item) {
    a.reset();
    Object.keys(item || {}).forEach(function (value) {
      value = item[value];
      this.b[value.name] = {
        Ua: value
      };
      this.b[value.name].D = load(this, value);
      push(this, value, this.g.length);
      value.Pa.push(this);
    }, a);
  }
  /**
   * @param {!Object} x
   * @param {!Object} b
   * @return {?}
   */
  function overlay(x, b) {
    var el = $("tr", {
      className: "ulhead"
    });
    $(el, "td", {
      onclick: function () {
        forward(x, 0);
        return false;
      }
    }, $("div", {
      className: "darr"
    }));
    /** @type {number} */
    var i = 0;
    for (; i < x.cols.length; i++) {
      $(el, "td", {
        onclick: function (i) {
          return function () {
            forward(x, x.cols[i]);
          };
        }(i)
      }, $("div", {
        className: "darr"
      }));
    }
    fn(x.a, el, b);
    return el;
  }
  /**
   * @param {!Object} self
   * @param {!Object} options
   * @param {!Object} template
   * @return {?}
   */
  function load(self, options, template) {
    var all = self.o.ud && options.A != "(" + self.app.lang + ")";
    all = $("tr", {
      onclick: function () {
        self.o.Pc(options, self);
        return false;
      }
    }, [$("td", {}, [0 != (options.a & DIRECTION_HORIZONTAL) || 0 != (options.a & arg) ? $("div", {
      className: "ulsym",
      style: {
        cssFloat: "right"
      }
    }, 0 != (options.a & arg) ? "X" : 0 != (options.a & DIRECTION_HORIZONTAL) ? "\u2605" : "") : null, $("div", {
      className: "ulnm"
    }, [$("div", {
      className: "r" + merge(self.app, options.g)
    }), options.name, $("span", {
      className: "ulla" + (all ? "" : " ulla0")
    }, options.A)])]), $("td", {
      className: "m1ac"
    }, $("button", {
      className: "ulbx"
    }, "X"))]);
    /** @type {number} */
    var i = 0;
    for (; 2 > i; i++) {
      var key = self.cols[i];
      var span = {
        className: "m0ac ulnu"
      };
      if (key == h) {
        $(all, "td", span, val(options.b));
      } else {
        if (key == value) {
          $(all, "td", span, fetch(self.app, options.R));
        } else {
          if (key == F) {
            $(all, "td", span, exec(self.app, options.g));
          } else {
            if (key == file) {
              $(all, "td", {
                className: "m0ac ulnu",
                title: val(options.b)
              }, 0 != options.b ? "#" : "");
            }
          }
        }
      }
    }
    fn(self.a, all, template);
    return all;
  }
  /**
   * @param {string} t
   * @return {undefined}
   */
  function onload(t) {
    /**
     * @return {undefined}
     */
    function init() {
      target = context.createBuffer(1, 2048, context.sampleRate);
      var positions = target.getChannelData(0);
      /** @type {number} */
      var i = 0;
      var l = positions.length;
      for (; i < l; i++) {
        /** @type {number} */
        positions[i] = .6 * Math.sin(i * Math.PI * 2 / 52) * (l - i) / l;
      }
    }
    if (t) {
      scale(t);
    }
    /** @type {string} */
    t = window.navigator.userAgent || "";
    var Buffer = window.AudioContext || window.webkitAudioContext;
    if (Buffer) {
      try {
        context = new Buffer;
      } catch (d) {
        return;
      }
      if (0 < t.indexOf("Windows NT 5.1") && 0 < t.indexOf("Firefox/")) {
        init();
      } else {
        t = window.k2snd["a.mp3"];
        context.decodeAudioData(function (string) {
          /** @type {number} */
          var size = string.length / 4 * 3;
          string = window.atob(string);
          /** @type {!ArrayBuffer} */
          var array = new ArrayBuffer(size);
          /** @type {!Uint8Array} */
          var u8arr = new Uint8Array(array);
          /** @type {number} */
          var n = 0;
          for (; n < size; n++) {
            u8arr[n] = string.charCodeAt(n);
          }
          return array;
        }(t.substring(t.indexOf(",") + 1)), function (userElem) {
          /** @type {string} */
          target = userElem;
        }, function () {
          init();
        });
      }
    }
  }
  /**
   * @param {!Object} source
   * @return {undefined}
   */
  function play(source) {
    if (context) {
      if (target && source) {
        source = context.createBufferSource();
        source.buffer = target;
        source.connect(context.destination);
        if (source.noteOn) {
          source.noteOn(0);
        } else {
          source.start(0);
        }
      }
    } else {
      if (corsAvailable || loadAudio(), audioElement && source) {
        try {
          audioElement.play();
        } catch (b) {
        }
      }
    }
  }
  /**
   * @return {undefined}
   */
  function loadAudio() {
    if (audioElement = document.createElement("audio")) {
      if (0 < (window.navigator.userAgent || "").indexOf(" Firefox/4") || !audioElement.canPlayType || !audioElement.canPlayType("audio/mpeg")) {
        /** @type {null} */
        audioElement = null;
      } else {
        audioElement.src = window.k2snd["a.mp3"];
      }
    }
    /** @type {boolean} */
    corsAvailable = true;
  }
  /**
   * @param {string} type
   * @return {undefined}
   */
  function scale(type) {
    if (!we) {
      /** @type {boolean} */
      we = true;
      /**
       * @return {undefined}
       */
      var start = function () {
        document.removeEventListener(type, start, true);
        if (context) {
          var source = context.createBufferSource();
          try {
            source.buffer = context.createBuffer(1, 1, context.sampleRate);
          } catch (d) {
          }
          source.connect(context.destination);
          if (source.noteOn) {
            source.noteOn(0);
          } else {
            source.start(0);
          }
        } else {
          if (corsAvailable || loadAudio(), audioElement) {
            try {
              audioElement.play();
              audioElement.pause();
            } catch (d) {
            }
          }
        }
      };
      document.addEventListener(type, start, true);
    }
  }
  /**
   * @param {(Object|string)} o
   * @return {undefined}
   */
  function ready(o) {
    var h = this;
    /** @type {(Object|string)} */
    this.o = o;
    /** @type {number} */
    this.g = this.b = 0;
    /** @type {number} */
    this.L = -1;
    /** @type {string} */
    this.I = "";
    /** @type {number} */
    this.a = 0;
    /** @type {null} */
    this.M = this.G = null;
    /** @type {string} */
    this.B = "";
    /** @type {number} */
    this.H = 0;
    setTimeout(function () {
      if (0 < h.b || 0 == h.b && 1 == h.L) {
        h.o.Xc(runlist);
      }
    }, 500);
    /** @type {number} */
    this.N = Date.now();
    connect(this);
  }
  /**
   * @param {!Object} a
   * @return {undefined}
   */
  function connect(a) {
    a.b += 1;
    if (a.b > a.o.ports.length) {
      /** @type {number} */
      a.b = -1;
      a.o.Xc(encoding);
    } else {
      var y = a.b;
      /** @type {number} */
      a.g = setTimeout(function () {
        /** @type {number} */
        a.g = 0;
        strictEqual(a, y);
      }, 1E3 * (2 == y ? 2 : 6));
      var center = a.o.ports[y - 1].split(":");
      if (!(2 > center.length)) {
        if ("wss" == center[0] || "ws" == center[0]) {
          _connect(a, y, center[0], parseInt(center[1], 10));
        } else {
          if ("https" == center[0] || "http" == center[0]) {
            distance(a, y, center[0], parseInt(center[1], 10));
          }
        }
      }
    }
  }
  /**
   * @param {!Object} a
   * @param {string} b
   * @return {?}
   */
  function assert(a, b) {
    if (b != a.b) {
      return false;
    }
    if (a.g) {
      clearTimeout(a.g);
      /** @type {number} */
      a.g = 0;
    }
    /** @type {number} */
    a.b = 0;
    return true;
  }
  /**
   * @param {!Object} a
   * @param {string} b
   * @return {undefined}
   */
  function strictEqual(a, b) {
    if (b == a.b) {
      if (a.g) {
        clearTimeout(a.g);
        /** @type {number} */
        a.g = 0;
      }
      connect(a);
    }
  }
  /**
   * @param {!Object} value
   * @return {undefined}
   */
  function disconnect(value) {
    if (value.C) {
      clearTimeout(value.C);
      /** @type {null} */
      value.C = null;
    }
    value.o.Xc(or);
  }
  /**
   * @param {!Object} self
   * @param {string} p
   * @return {undefined}
   */
  function drop(self, p) {
    p = p.split("\n");
    /** @type {number} */
    var j = 0;
    var pl = p.length;
    for (; j < pl; j++) {
      try {
        /** @type {*} */
        var sourceCell = JSON.parse(p[j]);
        var selectedMediaItems = sourceCell.i || [];
        var QueryLanguageComponent = sourceCell.s || [];
      } catch (onwaiting) {
        split("PARSE/ds.l=" + p.length + " >" + p[j] + "< " + onwaiting);
        return;
      }
      if (0 != selectedMediaItems.length) {
        if (-1 == self.L) {
          /** @type {number} */
          self.L = 1 == selectedMediaItems[0] ? 1 : 0;
        }
        if (1 == selectedMediaItems[0]) {
          if (self.o.ee()) {
            self.send([2], null);
          }
        } else {
          self.o.Ye(selectedMediaItems, QueryLanguageComponent);
        }
      }
    }
    if (!self.C) {
      /** @type {number} */
      self.C = setTimeout(function () {
        return getFile(self);
      }, 3E4);
    }
  }
  /**
   * @param {!Object} self
   * @return {undefined}
   */
  function getFile(self) {
    if (self.o.ee()) {
      self.send([], null);
    }
    /** @type {number} */
    self.C = setTimeout(function () {
      return getFile(self);
    }, 3E4);
  }
  /**
   * @param {!Array} target
   * @param {!Array} a
   * @return {?}
   */
  function encode(target, a) {
    a = a ? a : [];
    /** @type {number} */
    var j = 0;
    var startLen = a.length;
    for (; j < startLen; j++) {
      /** @type {string} */
      a[j] = '"' + a[j].replace(/\\/g, "\\\\").replace(/"/g, '\\"') + '"';
    }
    return '{"i":[' + target.join() + "]" + (0 < a.length ? ',"s":[' + a.join() + "]" : "") + "}";
  }
  /**
   * @param {!Object} context
   * @param {string} a
   * @param {string} url
   * @param {number} host
   * @return {undefined}
   */
  function _connect(context, a, url, host) {
    try {
      /** @type {!WebSocket} */
      var socket = new WebSocket(url + "://" + context.o.host + ":" + host + "/ws/");
    } catch (g) {
      strictEqual(context, a);
      return;
    }
    /**
     * @return {undefined}
     */
    socket.onclose = function () {
      strictEqual(context, a);
    };
    /**
     * @return {undefined}
     */
    socket.onopen = function () {
      /**
       * @param {!Object} b
       * @return {undefined}
       */
      socket.onmessage = function (b) {
        if (assert(context, a)) {
          /** @type {!WebSocket} */
          context.A = socket;
          /**
           * @param {!Object} b
           * @return {undefined}
           */
          context.A.onmessage = function (b) {
            drop(context, b.data);
          };
          /**
           * @return {undefined}
           */
          context.A.onclose = function () {
            /** @type {null} */
            context.A = null;
            disconnect(context);
          };
          drop(context, b.data);
        } else {
          socket.close();
        }
      };
      var args = context.o.de(a, context.N);
      socket.send(encode(args.J, args.O));
    };
  }
  /**
   * @param {!Object} r
   * @param {string} val
   * @param {string} s
   * @param {number} dx
   * @return {undefined}
   */
  function distance(r, val, s, dx) {
    /** @type {string} */
    r.I = s + "://" + r.o.host + ":" + dx;
    new xhr({
      url: r.I + "/r/0",
      data: "1",
      Wa: function () {
        strictEqual(r, val);
      },
      onload: function (c) {
        if (assert(r, val)) {
          r.a = c || "X" + Math.random();
          /** @type {number} */
          r.H = 0;
          c = r.o.de(val, r.N);
          r.send(c.J, c.O);
        }
      }
    });
  }
  /**
   * @param {!Object} item
   * @return {undefined}
   */
  function registerAsLoaded(item) {
    if (item.a && !item.G) {
      item.G = new xhr({
        url: item.I + "/r/" + item.a,
        data: null,
        Wa: function () {
          /** @type {null} */
          item.G = null;
          if (0 < item.H) {
            if (item.a) {
              /** @type {null} */
              item.a = null;
              disconnect(item);
            }
          } else {
            item.H++;
            setTimeout(function () {
              registerAsLoaded(item);
            }, 25);
          }
        },
        onload: function (d) {
          /** @type {null} */
          item.G = null;
          /** @type {number} */
          item.H = 0;
          if (d && 0 < d.length) {
            drop(item, d);
          }
          setTimeout(function () {
            registerAsLoaded(item);
          }, 25);
        }
      });
    }
  }
  /**
   * @param {!Object} node
   * @param {string} name
   * @return {undefined}
   */
  function addOverlay(node, name) {
    if (node.a) {
      if (name && 0 < name.length) {
        node.B += (0 < node.B.length ? "\n" : "") + name;
      }
      if (!node.M) {
        name = node.B;
        /** @type {string} */
        node.B = "";
        node.M = new xhr({
          url: node.I + "/w/" + node.a,
          data: name,
          Wa: function () {
            if (node.a) {
              /** @type {null} */
              node.a = null;
              disconnect(node);
            }
          },
          onload: function () {
            if (!node.G) {
              registerAsLoaded(node);
            }
            /** @type {null} */
            node.M = null;
            setTimeout(function () {
              if (0 < node.B.length) {
                addOverlay(node);
              }
            }, 25);
          }
        });
      }
    }
  }
  /**
   * @param {!Object} options
   * @return {undefined}
   */
  function xhr(options) {
    setTimeout(function () {
      /** @type {!XMLHttpRequest} */
      var request = new XMLHttpRequest;
      if (options.Wc || "withCredentials" in request) {
        /**
         * @return {undefined}
         */
        request.onreadystatechange = function () {
          if (4 == request.readyState) {
            if (200 == request.status || 204 == request.status) {
              if (options.onload) {
                options.onload(request.responseText);
              }
            } else {
              if (options.Wa) {
                options.Wa(request.status);
              }
            }
          }
        };
      } else {
        if (window.XDomainRequest) {
          /** @type {!XDomainRequest} */
          request = new XDomainRequest;
          /**
           * @return {undefined}
           */
          request.onprogress = function () {
          };
          /**
           * @return {undefined}
           */
          request.onload = function () {
            if (options.onload) {
              options.onload(request.responseText);
            }
          };
          /** @type {function(): undefined} */
          request.onerror = request.ontimeout = function () {
            if (options.Wa) {
              options.Wa();
            }
          };
        }
      }
      request.open("POST", options.url, true);
      if (options.Wc && !options.Ze && request.setRequestHeader) {
        request.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      }
      request.send(options.data);
    }, 0);
  }
  /**
   * @param {string} options
   * @param {!Object} data
   * @return {undefined}
   */
  function View(options, data) {
    var $scope = this;
    new xhr({
      url: options,
      data: data,
      Wc: true,
      Ze: "string" !== typeof data,
      onload: function (text) {
        if ($scope.handle) {
          $scope.handle(true, JSON.parse(text));
        }
      },
      Wa: function () {
        if ($scope.handle) {
          $scope.handle(false, {});
        }
      }
    });
  }
  /**
   * @param {!Object} name
   * @return {?}
   */
  function $(name) {
    /** @type {!Arguments} */
    var a = arguments;
    /** @type {null} */
    var ret = null;
    /** @type {null} */
    var out = null;
    /** @type {number} */
    var i = 0;
    if (0 == a.length) {
      return ret;
    }
    if (a[i].nodeName) {
      out = a[i++];
    }
    if (Array.isArray(a[i])) {
      return ret = [], a[i].forEach(function (value) {
        if (null != value) {
          value = "string" == typeof value || "number" == typeof value ? document.createTextNode(value) : value;
          if (out) {
            out.appendChild(value);
          }
          ret.push(value);
        }
      }), ret;
    }
    if ("string" !== typeof a[i]) {
      return ret;
    }
    /** @type {!Element} */
    ret = document.createElement(a[i++]);
    if (i < a.length && "object" == typeof a[i] && !Array.isArray(a[i]) && !a[i].nodeName) {
      var attrs = a[i++];
      Object.keys(attrs).forEach(function (attr) {
        if ("style" == attr) {
          Object.keys(attrs[attr]).forEach(function (i) {
            ret.style[i] = attrs[attr][i];
          });
        } else {
          ret[attr] = attrs[attr];
        }
      });
    }
    if (out) {
      out.appendChild(ret);
    }
    if (i >= a.length) {
      return ret;
    }
    a = a[i];
    if ("string" == typeof a) {
      ret.appendChild(document.createTextNode(a));
    } else {
      if (Array.isArray(a)) {
        a.forEach(function (f) {
          if ("string" == typeof f) {
            ret.appendChild(document.createTextNode(f));
          } else {
            if (f) {
              ret.appendChild(f);
            }
          }
        });
      } else {
        ret.appendChild(a);
      }
    }
    return ret;
  }
  /**
   * @param {!Object} node
   * @param {!Array} value
   * @return {undefined}
   */
  function resolve(node, value) {
    /** @type {string} */
    node.innerHTML = "";
    if (Array.isArray(value)) {
      value.forEach(function (msg) {
        if ("string" == typeof msg || "number" == typeof msg) {
          node.appendChild(document.createTextNode(msg));
        } else {
          if (msg) {
            node.appendChild(msg);
          }
        }
      });
    } else {
      node.appendChild(value);
    }
  }
  /**
   * @param {!Object} target
   * @param {?} n
   * @param {?} s
   * @return {?}
   */
  function html(target, n, s) {
    return s ? target.insertBefore(n, s) : target.appendChild(n);
  }
  /**
   * @param {!Object} container
   * @param {?} p
   * @param {!Object} obj
   * @return {?}
   */
  function fn(container, p, obj) {
    if (obj) {
      container.replaceChild(p, obj);
    } else {
      container.appendChild(p);
    }
    return p;
  }
  /**
   * @param {!Object} a
   * @param {string} val
   * @return {undefined}
   */
  function map(a, val) {
    /** @type {number} */
    var i = 0;
    for (; i < a.options.length; i++) {
      if (a.options[i].text == val) {
        /** @type {number} */
        a.selectedIndex = i;
        break;
      }
    }
  }
  /**
   * @param {!Object} pattern
   * @param {!Object} data
   * @return {undefined}
   */
  function extend(pattern, data) {
    if ("undefined" !== typeof data) {
      Object.keys(data).forEach(function (k) {
        pattern.style[k] = data[k];
      });
    }
  }
  /**
   * @param {!Object} a
   * @param {string} v
   * @param {boolean} c
   * @return {undefined}
   */
  function callback(a, v, c) {
    if (2 == arguments.length || c) {
      a.classList.add(v);
    } else {
      a.classList.remove(v);
    }
  }
  /**
   * @param {!Object} t
   * @param {string} val
   * @return {?}
   */
  function transform(t, val) {
    return t.classList.contains(val);
  }
  /**
   * @param {!Object} arg
   * @return {undefined}
   */
  function acceptTheCall(arg) {
    callback(arg, "nav0open", !transform(arg, "nav0open"));
  }
  /**
   * @return {?}
   */
  function getUnderlineBackgroundPositionY() {
    return -1 != (window.navigator.userAgent || "").indexOf("IEMobile/") ? 2 : Math.max(window.devicePixelRatio || 1, 1);
  }
  /**
   * @param {!Object} container
   * @param {number} x
   * @param {!Date} y
   * @return {?}
   */
  function getMousePosition(container, x, y) {
    container = container.getBoundingClientRect();
    return {
      x: x - container.left,
      y: y - container.top
    };
  }
  /**
   * @param {!Object} label
   * @param {?} version
   * @return {undefined}
   */
  function Layer(label, version) {
    this.f = $(label, "div", version);
    /** @type {number} */
    this.a = 0;
  }
  /**
   * @param {string} buff
   * @param {string} name
   * @return {?}
   */
  function toString(buff, name) {
    var i = buff.indexOf("%s");
    return -1 != i ? buff.substring(0, i) + name + buff.substring(i + 2) : buff;
  }
  /**
   * @param {!Object} o
   * @return {undefined}
   */
  function write(o) {
    o.innerHTML = o.innerHTML.replace(/\uD83C\uDFF4(?:\uDB40\uDC67\uDB40\uDC62(?:\uDB40\uDC65\uDB40\uDC6E\uDB40\uDC67|\uDB40\uDC77\uDB40\uDC6C\uDB40\uDC73|\uDB40\uDC73\uDB40\uDC63\uDB40\uDC74)\uDB40\uDC7F|\u200D\u2620\uFE0F)|\uD83D\uDC69\u200D\uD83D\uDC69\u200D(?:\uD83D\uDC66\u200D\uD83D\uDC66|\uD83D\uDC67\u200D(?:\uD83D[\uDC66\uDC67]))|\uD83D\uDC68(?:\u200D(?:\u2764\uFE0F\u200D(?:\uD83D\uDC8B\u200D)?\uD83D\uDC68|(?:\uD83D[\uDC68\uDC69])\u200D(?:\uD83D\uDC66\u200D\uD83D\uDC66|\uD83D\uDC67\u200D(?:\uD83D[\uDC66\uDC67]))|\uD83D\uDC66\u200D\uD83D\uDC66|\uD83D\uDC67\u200D(?:\uD83D[\uDC66\uDC67])|\uD83C[\uDF3E\uDF73\uDF93\uDFA4\uDFA8\uDFEB\uDFED]|\uD83D[\uDCBB\uDCBC\uDD27\uDD2C\uDE80\uDE92]|\uD83E[\uDDB0-\uDDB3])|(?:\uD83C[\uDFFB-\uDFFF])\u200D(?:\uD83C[\uDF3E\uDF73\uDF93\uDFA4\uDFA8\uDFEB\uDFED]|\uD83D[\uDCBB\uDCBC\uDD27\uDD2C\uDE80\uDE92]|\uD83E[\uDDB0-\uDDB3]))|\uD83D\uDC69\u200D(?:\u2764\uFE0F\u200D(?:\uD83D\uDC8B\u200D(?:\uD83D[\uDC68\uDC69])|\uD83D[\uDC68\uDC69])|\uD83C[\uDF3E\uDF73\uDF93\uDFA4\uDFA8\uDFEB\uDFED]|\uD83D[\uDCBB\uDCBC\uDD27\uDD2C\uDE80\uDE92]|\uD83E[\uDDB0-\uDDB3])|\uD83D\uDC69\u200D\uD83D\uDC66\u200D\uD83D\uDC66|(?:\uD83D\uDC41\uFE0F\u200D\uD83D\uDDE8|\uD83D\uDC69(?:\uD83C[\uDFFB-\uDFFF])\u200D[\u2695\u2696\u2708]|\uD83D\uDC68(?:(?:\uD83C[\uDFFB-\uDFFF])\u200D[\u2695\u2696\u2708]|\u200D[\u2695\u2696\u2708])|(?:(?:\u26F9|\uD83C[\uDFCB\uDFCC]|\uD83D\uDD75)\uFE0F|\uD83D\uDC6F|\uD83E[\uDD3C\uDDDE\uDDDF])\u200D[\u2640\u2642]|(?:\u26F9|\uD83C[\uDFCB\uDFCC]|\uD83D\uDD75)(?:\uD83C[\uDFFB-\uDFFF])\u200D[\u2640\u2642]|(?:\uD83C[\uDFC3\uDFC4\uDFCA]|\uD83D[\uDC6E\uDC71\uDC73\uDC77\uDC81\uDC82\uDC86\uDC87\uDE45-\uDE47\uDE4B\uDE4D\uDE4E\uDEA3\uDEB4-\uDEB6]|\uD83E[\uDD26\uDD37-\uDD39\uDD3D\uDD3E\uDDB8\uDDB9\uDDD6-\uDDDD])(?:(?:\uD83C[\uDFFB-\uDFFF])\u200D[\u2640\u2642]|\u200D[\u2640\u2642])|\uD83D\uDC69\u200D[\u2695\u2696\u2708])\uFE0F|\uD83D\uDC69\u200D\uD83D\uDC67\u200D(?:\uD83D[\uDC66\uDC67])|\uD83D\uDC69\u200D\uD83D\uDC69\u200D(?:\uD83D[\uDC66\uDC67])|\uD83D\uDC68(?:\u200D(?:(?:\uD83D[\uDC68\uDC69])\u200D(?:\uD83D[\uDC66\uDC67])|\uD83D[\uDC66\uDC67])|\uD83C[\uDFFB-\uDFFF])|\uD83C\uDFF3\uFE0F\u200D\uD83C\uDF08|\uD83D\uDC69\u200D\uD83D\uDC67|\uD83D\uDC69(?:\uD83C[\uDFFB-\uDFFF])\u200D(?:\uD83C[\uDF3E\uDF73\uDF93\uDFA4\uDFA8\uDFEB\uDFED]|\uD83D[\uDCBB\uDCBC\uDD27\uDD2C\uDE80\uDE92]|\uD83E[\uDDB0-\uDDB3])|\uD83D\uDC69\u200D\uD83D\uDC66|\uD83C\uDDF6\uD83C\uDDE6|\uD83C\uDDFD\uD83C\uDDF0|\uD83C\uDDF4\uD83C\uDDF2|\uD83D\uDC69(?:\uD83C[\uDFFB-\uDFFF])|\uD83C\uDDED(?:\uD83C[\uDDF0\uDDF2\uDDF3\uDDF7\uDDF9\uDDFA])|\uD83C\uDDEC(?:\uD83C[\uDDE6\uDDE7\uDDE9-\uDDEE\uDDF1-\uDDF3\uDDF5-\uDDFA\uDDFC\uDDFE])|\uD83C\uDDEA(?:\uD83C[\uDDE6\uDDE8\uDDEA\uDDEC\uDDED\uDDF7-\uDDFA])|\uD83C\uDDE8(?:\uD83C[\uDDE6\uDDE8\uDDE9\uDDEB-\uDDEE\uDDF0-\uDDF5\uDDF7\uDDFA-\uDDFF])|\uD83C\uDDF2(?:\uD83C[\uDDE6\uDDE8-\uDDED\uDDF0-\uDDFF])|\uD83C\uDDF3(?:\uD83C[\uDDE6\uDDE8\uDDEA-\uDDEC\uDDEE\uDDF1\uDDF4\uDDF5\uDDF7\uDDFA\uDDFF])|\uD83C\uDDFC(?:\uD83C[\uDDEB\uDDF8])|\uD83C\uDDFA(?:\uD83C[\uDDE6\uDDEC\uDDF2\uDDF3\uDDF8\uDDFE\uDDFF])|\uD83C\uDDF0(?:\uD83C[\uDDEA\uDDEC-\uDDEE\uDDF2\uDDF3\uDDF5\uDDF7\uDDFC\uDDFE\uDDFF])|\uD83C\uDDEF(?:\uD83C[\uDDEA\uDDF2\uDDF4\uDDF5])|\uD83C\uDDF8(?:\uD83C[\uDDE6-\uDDEA\uDDEC-\uDDF4\uDDF7-\uDDF9\uDDFB\uDDFD-\uDDFF])|\uD83C\uDDEE(?:\uD83C[\uDDE8-\uDDEA\uDDF1-\uDDF4\uDDF6-\uDDF9])|\uD83C\uDDFF(?:\uD83C[\uDDE6\uDDF2\uDDFC])|\uD83C\uDDEB(?:\uD83C[\uDDEE-\uDDF0\uDDF2\uDDF4\uDDF7])|\uD83C\uDDF5(?:\uD83C[\uDDE6\uDDEA-\uDDED\uDDF0-\uDDF3\uDDF7-\uDDF9\uDDFC\uDDFE])|\uD83C\uDDE9(?:\uD83C[\uDDEA\uDDEC\uDDEF\uDDF0\uDDF2\uDDF4\uDDFF])|\uD83C\uDDF9(?:\uD83C[\uDDE6\uDDE8\uDDE9\uDDEB-\uDDED\uDDEF-\uDDF4\uDDF7\uDDF9\uDDFB\uDDFC\uDDFF])|\uD83C\uDDE7(?:\uD83C[\uDDE6\uDDE7\uDDE9-\uDDEF\uDDF1-\uDDF4\uDDF6-\uDDF9\uDDFB\uDDFC\uDDFE\uDDFF])|[#\*0-9]\uFE0F\u20E3|\uD83C\uDDF1(?:\uD83C[\uDDE6-\uDDE8\uDDEE\uDDF0\uDDF7-\uDDFB\uDDFE])|\uD83C\uDDE6(?:\uD83C[\uDDE8-\uDDEC\uDDEE\uDDF1\uDDF2\uDDF4\uDDF6-\uDDFA\uDDFC\uDDFD\uDDFF])|\uD83C\uDDF7(?:\uD83C[\uDDEA\uDDF4\uDDF8\uDDFA\uDDFC])|\uD83C\uDDFB(?:\uD83C[\uDDE6\uDDE8\uDDEA\uDDEC\uDDEE\uDDF3\uDDFA])|\uD83C\uDDFE(?:\uD83C[\uDDEA\uDDF9])|(?:\uD83C[\uDFC3\uDFC4\uDFCA]|\uD83D[\uDC6E\uDC71\uDC73\uDC77\uDC81\uDC82\uDC86\uDC87\uDE45-\uDE47\uDE4B\uDE4D\uDE4E\uDEA3\uDEB4-\uDEB6]|\uD83E[\uDD26\uDD37-\uDD39\uDD3D\uDD3E\uDDB8\uDDB9\uDDD6-\uDDDD])(?:\uD83C[\uDFFB-\uDFFF])|(?:\u26F9|\uD83C[\uDFCB\uDFCC]|\uD83D\uDD75)(?:\uD83C[\uDFFB-\uDFFF])|(?:[\u261D\u270A-\u270D]|\uD83C[\uDF85\uDFC2\uDFC7]|\uD83D[\uDC42\uDC43\uDC46-\uDC50\uDC66\uDC67\uDC70\uDC72\uDC74-\uDC76\uDC78\uDC7C\uDC83\uDC85\uDCAA\uDD74\uDD7A\uDD90\uDD95\uDD96\uDE4C\uDE4F\uDEC0\uDECC]|\uD83E[\uDD18-\uDD1C\uDD1E\uDD1F\uDD30-\uDD36\uDDB5\uDDB6\uDDD1-\uDDD5])(?:\uD83C[\uDFFB-\uDFFF])|(?:[\u231A\u231B\u23E9-\u23EC\u23F0\u23F3\u25FD\u25FE\u2614\u2615\u2648-\u2653\u267F\u2693\u26A1\u26AA\u26AB\u26BD\u26BE\u26C4\u26C5\u26CE\u26D4\u26EA\u26F2\u26F3\u26F5\u26FA\u26FD\u2705\u270A\u270B\u2728\u274C\u274E\u2753-\u2755\u2757\u2795-\u2797\u27B0\u27BF\u2B1B\u2B1C\u2B50\u2B55]|\uD83C[\uDC04\uDCCF\uDD8E\uDD91-\uDD9A\uDDE6-\uDDFF\uDE01\uDE1A\uDE2F\uDE32-\uDE36\uDE38-\uDE3A\uDE50\uDE51\uDF00-\uDF20\uDF2D-\uDF35\uDF37-\uDF7C\uDF7E-\uDF93\uDFA0-\uDFCA\uDFCF-\uDFD3\uDFE0-\uDFF0\uDFF4\uDFF8-\uDFFF]|\uD83D[\uDC00-\uDC3E\uDC40\uDC42-\uDCFC\uDCFF-\uDD3D\uDD4B-\uDD4E\uDD50-\uDD67\uDD7A\uDD95\uDD96\uDDA4\uDDFB-\uDE4F\uDE80-\uDEC5\uDECC\uDED0-\uDED2\uDEEB\uDEEC\uDEF4-\uDEF9]|\uD83E[\uDD10-\uDD3A\uDD3C-\uDD3E\uDD40-\uDD45\uDD47-\uDD70\uDD73-\uDD76\uDD7A\uDD7C-\uDDA2\uDDB0-\uDDB9\uDDC0-\uDDC2\uDDD0-\uDDFF])|(?:[#\*0-9\xA9\xAE\u203C\u2049\u2122\u2139\u2194-\u2199\u21A9\u21AA\u231A\u231B\u2328\u23CF\u23E9-\u23F3\u23F8-\u23FA\u24C2\u25AA\u25AB\u25B6\u25C0\u25FB-\u25FE\u2600-\u2604\u260E\u2611\u2614\u2615\u2618\u261D\u2620\u2622\u2623\u2626\u262A\u262E\u262F\u2638-\u263A\u2640\u2642\u2648-\u2653\u265F\u2660\u2663\u2665\u2666\u2668\u267B\u267E\u267F\u2692-\u2697\u2699\u269B\u269C\u26A0\u26A1\u26AA\u26AB\u26B0\u26B1\u26BD\u26BE\u26C4\u26C5\u26C8\u26CE\u26CF\u26D1\u26D3\u26D4\u26E9\u26EA\u26F0-\u26F5\u26F7-\u26FA\u26FD\u2702\u2705\u2708-\u270D\u270F\u2712\u2714\u2716\u271D\u2721\u2728\u2733\u2734\u2744\u2747\u274C\u274E\u2753-\u2755\u2757\u2763\u2764\u2795-\u2797\u27A1\u27B0\u27BF\u2934\u2935\u2B05-\u2B07\u2B1B\u2B1C\u2B50\u2B55\u3030\u303D\u3297\u3299]|\uD83C[\uDC04\uDCCF\uDD70\uDD71\uDD7E\uDD7F\uDD8E\uDD91-\uDD9A\uDDE6-\uDDFF\uDE01\uDE02\uDE1A\uDE2F\uDE32-\uDE3A\uDE50\uDE51\uDF00-\uDF21\uDF24-\uDF93\uDF96\uDF97\uDF99-\uDF9B\uDF9E-\uDFF0\uDFF3-\uDFF5\uDFF7-\uDFFF]|\uD83D[\uDC00-\uDCFD\uDCFF-\uDD3D\uDD49-\uDD4E\uDD50-\uDD67\uDD6F\uDD70\uDD73-\uDD7A\uDD87\uDD8A-\uDD8D\uDD90\uDD95\uDD96\uDDA4\uDDA5\uDDA8\uDDB1\uDDB2\uDDBC\uDDC2-\uDDC4\uDDD1-\uDDD3\uDDDC-\uDDDE\uDDE1\uDDE3\uDDE8\uDDEF\uDDF3\uDDFA-\uDE4F\uDE80-\uDEC5\uDECB-\uDED2\uDEE0-\uDEE5\uDEE9\uDEEB\uDEEC\uDEF0\uDEF3-\uDEF9]|\uD83E[\uDD10-\uDD3A\uDD3C-\uDD3E\uDD40-\uDD45\uDD47-\uDD70\uDD73-\uDD76\uDD7A\uDD7C-\uDDA2\uDDB0-\uDDB9\uDDC0-\uDDC2\uDDD0-\uDDFF])\uFE0F|(?:[\u261D\u26F9\u270A-\u270D]|\uD83C[\uDF85\uDFC2-\uDFC4\uDFC7\uDFCA-\uDFCC]|\uD83D[\uDC42\uDC43\uDC46-\uDC50\uDC66-\uDC69\uDC6E\uDC70-\uDC78\uDC7C\uDC81-\uDC83\uDC85-\uDC87\uDCAA\uDD74\uDD75\uDD7A\uDD90\uDD95\uDD96\uDE45-\uDE47\uDE4B-\uDE4F\uDEA3\uDEB4-\uDEB6\uDEC0\uDECC]|\uD83E[\uDD18-\uDD1C\uDD1E\uDD1F\uDD26\uDD30-\uDD39\uDD3D\uDD3E\uDDB5\uDDB6\uDDB8\uDDB9\uDDD1-\uDDDD])/g,
      w);
  }
  /**
   * @param {string} newWayId
   * @return {undefined}
   */
  function split(newWayId) {
    if (document.domain) {
      new View("/misc/k2rep.php", "fb=" + newWayId);
    }
  }
  /**
   * @param {!Object} data
   * @return {?}
   */
  function redraw(data) {
    /** @type {string} */
    var result = "";
    Object.keys(data || {}).forEach(function (key) {
      result = result + ((result ? "&" : "") + encodeURIComponent(key) + "=" + encodeURIComponent(data[key]));
    });
    return result;
  }
  /**
   * @param {?} callback
   * @param {number} triggerNow
   * @return {undefined}
   */
  function SvgShape(callback, triggerNow) {
    this.N = callback;
    this.S = "undefined" == typeof triggerNow ? 4 : triggerNow;
    /** @type {boolean} */
    this.H = false;
    /** @type {number} */
    this.g = this.A = this.B = 0;
    /** @type {number} */
    this.a = -1;
    /** @type {number} */
    this.C = this.b = this.I = 0;
    /** @type {number} */
    this.M = -1;
    /** @type {number} */
    this.L = this.G = 0;
  }
  /**
   * @param {!Object} c
   * @return {undefined}
   */
  function verify(c) {
    /** @type {number} */
    var y = Date.now();
    if (0 < c.A && y > c.A) {
      /** @type {number} */
      c.A = 0;
    }
    if (-1 != c.a) {
      /** @type {number} */
      var i = y - c.g;
      /** @type {number} */
      y = 0;
      if (0 < c.b) {
        /** @type {number} */
        var j = c.b * Math.floor(200);
        if (i < j) {
          /** @type {number} */
          y = Math.ceil((j - i) / 1E3);
          /** @type {number} */
          i = 0;
        } else {
          /** @type {number} */
          i = i - j;
        }
      }
      /** @type {number} */
      i = c.I * Math.floor(200) + i * (0 < c.C ? 1 : -1);
      /** @type {number} */
      i = 0 < c.C ? Math.floor(i / 1E3) : Math.ceil(i / 1E3);
      if (c.a != c.M || i != c.G || y != c.L) {
        j = c.N;
        var h = c.a;
        /** @type {number} */
        var exceptions = y;
        if (null != j.sb) {
          /** @type {string} */
          j.sb = 0 > i ? "(?)" : Math.floor(i / 60) + ":" + Math.floor(i % 60 / 10) % 10 + i % 60 % 10;
          constructor(j);
        } else {
          equal(j.T, h, i, exceptions);
          slice(j);
        }
        /** @type {number} */
        c.G = i;
        /** @type {number} */
        c.L = y;
        c.M = c.a;
      }
    }
  }
  /**
   * @param {!Object} value
   * @return {undefined}
   */
  function form(value) {
    var query = this;
    /** @type {!Object} */
    this.app = value;
    this.Ja = this.app.j("bl_buds");
    this.f = $(this.app.B, "div", {
      className: "stview usno",
      style: {
        display: "none"
      }
    });
    /** @type {!Array} */
    this.a = [];
    this.b = $(this.f, "div", {
      className: "clst btifp"
    });
    var c;
    $(this.f, "div", {
      className: "caddbox dcpd"
    }, [$("form", {
      onsubmit: function () {
        var a = c.value.trim();
        /** @type {string} */
        c.value = "";
        if (a) {
          query.app.X("/buddy " + a);
        }
        c.blur();
        return false;
      }
    }, [$("p", {}, c = $("input", {
      className: "aid",
      name: "x",
      autocomplete: "off"
    })), $("p", {}, $("input", {
      type: "submit",
      className: "minw",
      value: this.app.j("bl_ad")
    }))])]);
    m(this, this.app.b);
  }
  /**
   * @param {!Object} event
   * @param {string} text
   * @return {undefined}
   */
  function downandpress(event, text) {
    /** @type {number} */
    event.tab = 0;
    /** @type {string} */
    event.Ha = "";
    if (event.status) {
      /** @type {string} */
      event.Ha = text;
      text = 0 < event.$a.length ? event.$a[0] : 0;
      if ("#" == text) {
        event.tab = event.$a.substring(1);
      } else {
        if (":" == text) {
          event.Ha += " (" + event.$a.substring(1) + ")";
        }
      }
    } else {
      event.Ha = event.$a.substring(0, 10);
    }
  }
  /**
   * @param {!Object} options
   * @param {?} name
   * @return {undefined}
   */
  function handler(options, name) {
    if (!options.g) {
      options.g = create(options.app, "-");
    }
    var c;
    resolve(options.g.Ma, [options.app.S ? null : $("div", [$("p", $("a", {
      className: "lbut minwd",
      target: "_blank",
      href: toString(get(options.app, "stat"), encodeURIComponent(name)),
      onclick: function () {
        success(options.app);
      }
    }, options.app.j("ui_stats") + " >")), $("div", {
      className: "dtline"
    })]), $("p", [c = $("input", {
      type: "checkbox",
      checked: true
    }), options.app.j("bl_buds")]), $("p", $("button", {
      className: "minw",
      onclick: function () {
        success(options.app);
        if (!c.checked) {
          options.app.X("/unbuddy " + name);
        }
      }
    }, options.app.j("bl_ok")))]);
    filter(options.app, options.g, name);
  }
  /**
   * @param {!Object} data
   * @param {!Object} config
   * @param {!Object} value
   * @return {?}
   */
  function wrap(data, config, value) {
    return fn(data.b, $("a", {
      className: "awrap dcpd" + (config.status ? " st1" : ""),
      onclick: function () {
        handler(data, config.name);
        return false;
      }
    }, $("div", {
      className: "maxw"
    }, [$("div", {
      className: "chtbl",
      onclick: function (event) {
        each(data.app, config.name);
        event.stopPropagation();
        return false;
      }
    }, $("div", {
      className: "spbb"
    })), $("div", {
      className: "uname"
    }, [$("div", {
      className: "sta"
    }), config.name]), $("div", {
      className: "infbl",
      style: {
        verticalAlign: "top"
      }
    }, [$("span", {
      className: "inftx"
    }, config.Ha), config.tab ? $("button", {
      className: "butbl butlh",
      onclick: function (event) {
        if (config.tab) {
          server(data.app, config.tab);
        }
        event.stopPropagation();
        return false;
      }
    }, "#" + config.tab) : null])])), value);
  }
  /**
   * @param {string} t
   * @param {!Object} b
   * @return {undefined}
   */
  function append(t, b) {
    var i = t.a.length;
    /** @type {number} */
    var animation = b.status ? 1 : 0;
    for (; 0 < i;) {
      var a = t.a[i - 1];
      /** @type {number} */
      var object = a.status ? 1 : 0;
      /** @type {number} */
      var value = b.name < a.name ? -1 : b.name > a.name ? 1 : 0;
      if (0 <= (1 == animation ? 0 == object ? -1 : value : 1 == object ? 1 : b.Ha < a.Ha ? 1 : b.Ha > a.Ha ? -1 : value)) {
        break;
      }
      i--;
    }
    animation = t.a[i];
    t.a.splice(i, 0, b);
    if (b.D) {
      t.b.insertBefore(b.D, animation ? animation.D : null);
    }
  }
  /**
   * @param {undefined} result
   * @param {!Object} options
   * @return {undefined}
   */
  function getDebugTransaction(result, options) {
    options.D = wrap(result, options);
    append(result, options);
  }
  /**
   * @param {!Object} result
   * @param {?} item
   * @return {undefined}
   */
  function calc(result, item) {
    var disabledItemIndex = result.a.indexOf(item);
    if (-1 != disabledItemIndex) {
      if (item.D) {
        result.b.removeChild(item.D);
        /** @type {null} */
        item.D = null;
      }
      result.a.splice(disabledItemIndex, 1);
    }
  }
  /**
   * @param {string} data
   * @param {!Object} options
   * @return {undefined}
   */
  function initiate(data, options) {
    var c = data.a.indexOf(options);
    if (-1 != c) {
      options.D = wrap(data, options, options.D);
      data.a.splice(c, 1);
      append(data, options);
    }
  }
  /**
   * @param {string} a
   * @param {!Object} x
   * @return {undefined}
   */
  function m(a, x) {
    /** @type {number} */
    a.a.length = 0;
    Object.keys(x).forEach(function (data) {
      data = x[data];
      data.D = wrap(a, data);
      append(a, data);
    }, a);
  }
  /**
   * @param {!Object} value
   * @return {undefined}
   */
  function handle(value) {
    /**
     * @param {!Object} params
     * @return {?}
     */
    function init(params) {
      return [$("b", params.sd + (0 == params.Ga ? "" : " (" + params.j("t_gsmp") + ")")), params.o.ob ? null : $("div", {
        className: "mlo r" + merge(params, params.Fa)
      }), params.o.ob ? null : $("span", {
        className: "snum"
      }, exec(params, params.Fa))];
    }
    /** @type {!Object} */
    this.app = value;
    this.Ja = this.app.j("bl_mr");
    this.f = $(value.B, "div", {
      className: "stvxpad vnarrow",
      style: {
        display: "none"
      }
    });
    var $scope = this;
    $(this.f, "div", {
      className: "btifp"
    });
    if (0 == this.app.Ga) {
      $(this.f, [$("p", [$("button", {
        className: "minwd",
        onclick: function () {
          debug($scope.app, $scope.app.Ne);
        }
      }, this.app.j("bl_prefs")), " ", $("button", {
        className: "minwd",
        onclick: function () {
          debug($scope.app, $scope.app.Re);
        }
      }, this.app.j("t_prof"))]), this.app.P ? null : $("p", {
        style: {
          marginTop: "-.5em"
        }
      }, [$("button", {
        className: "minwd",
        onclick: function () {
          logout($scope.app);
        }
      }, this.app.j("t_lout"))]), $("hr")]);
    }
    var n = $(this.f, "p", init(this.app));
    expect(this.app, "urank", function () {
      resolve(n, init($scope.app));
    });
    if (this.app.S) {
      $(this.f, [$("hr"), $("p", $("button", {
        className: "minwd",
        onclick: function () {
          debug($scope.app, $scope.app.Pe);
        }
      }, this.app.j("t_fb")))]);
    }
  }
  /**
   * @param {!Object} value
   * @return {undefined}
   */
  function go(value) {
    /** @type {!Object} */
    this.app = value;
    this.Ja = this.app.j("t_fb");
    this.f = $(value.B, "div", {
      className: "stvxpad vnarrow",
      style: {
        display: "none"
      }
    });
    var root = this;
    var component;
    $(this.f, [$("p", {
      className: "fb"
    }, this.app.j("t_sf")), component = $("textarea", {
      className: "bsbb",
      rows: 6,
      style: {
        width: "100%"
      }
    }), $("p", $("button", {
      className: "minw",
      onclick: function () {
        var app = root.app;
        debug(app, app.N);
        app = component.value.trim();
        /** @type {string} */
        component.value = "";
        new xhr({
          url: get(root.app, "feedback"),
          data: "fb=" + app + "\n\n" + window.location.href + "\n" + screen.width + "x" + screen.height + "px",
          Wc: true
        });
        return false;
      }
    }, this.app.j("bl_ok")))]);
  }
  /**
   * @param {!Object} data
   * @return {undefined}
   */
  function start(data) {
    var menu = this;
    /** @type {!Object} */
    this.app = data;
    this.Ja = this.app.j("bl_cs");
    /** @type {null} */
    this.b = this.A = null;
    this.a = {};
    /** @type {!Array} */
    this.g = [];
    this.f = $(this.app.B, "div", {
      className: "stview",
      style: {
        display: "none"
      }
    });
    this.H = $(this.f, "div", {
      className: "imvfrm"
    });
    this.G = $(this.f, "div", {
      className: "imvlst"
    });
    this.B = $(this.app.H, "div", {
      className: "fb fl",
      style: {
        display: "none",
        color: "#fff"
      }
    }, [$("div", {
      className: "sta"
    }), this.I = $("div", {
      className: "ib"
    })]);
    this.C = $(this.G, "div", {
      className: "iml btifp"
    });
    $(this.G, "p", {
      className: "dcpd"
    }, this.A = $("select", {
      style: {
        margin: "1em 0 .25em",
        minWidth: "10em"
      },
      onchange: function (e) {
        e = e.target;
        if (0 < e.selectedIndex) {
          each(menu.app, e.options[e.selectedIndex].text);
        }
        /** @type {number} */
        menu.selectedIndex = 0;
      }
    }));
    attach(this);
    data = this.app.I;
    /** @type {number} */
    var c = 90;
    var subunit;
    for (subunit in data) {
      if (data.hasOwnProperty(subunit) && data[subunit].Oa && 0 < c) {
        var url = next(this, data[subunit].Ia, true);
        send(this, url);
        c--;
      }
    }
    if (0 < this.g.length) {
      this.app.send([reg, 1], this.g);
      /** @type {!Array} */
      this.g = [];
    }
  }
  /**
   * @param {!Object} e
   * @param {boolean} id
   * @return {undefined}
   */
  function normalize(e, id) {
    var key = e.app;
    if (key.C && key.C == e.f && (key = e.b)) {
      var els = e.app.I["_" + key.Ia];
      if (els && els.Oa) {
        e.app.send([25, els.Oa], [key.Ia]);
        /** @type {number} */
        els.Oa = 0;
        /** @type {number} */
        els.tc = 0;
        if (!(id && key == id)) {
          send(e, key);
        }
      }
    }
  }
  /**
   * @param {!Object} o
   * @return {undefined}
   */
  function _get(o) {
    if (o.app.oa && !o.b && o.C.firstChild) {
      var value = o.C.firstChild;
      var i;
      for (i in o.a) {
        if (o.a.hasOwnProperty(i) && o.a[i].fa == value) {
          request(o, o.a[i].Ia);
          break;
        }
      }
    }
  }
  /**
   * @param {!Object} e
   * @param {?} input
   * @return {undefined}
   */
  function send(e, input) {
    if (input) {
      var series = e.app.I["_" + input.Ia];
      if (!series) {
        return;
      }
      if (series.Oa) {
        resolve(input.fd, [series.tc]);
      }
      extend(input.fd, {
        visibility: series.Oa ? "inherit" : "hidden"
      });
    }
    rebuildModelFromFields(e.app);
  }
  /**
   * @param {!Object} e
   * @param {string} s
   * @param {!Object} c
   * @return {undefined}
   */
  function print(e, s, c) {
    var dataProcessor = e.app.pa["_" + s];
    if (dataProcessor && dataProcessor.Xd) {
      c.u.append("[" + e.app.j("chr") + "]", true);
    }
    if (e = e.app.I["_" + s]) {
      c.u.append(e.rc);
    }
  }
  /**
   * @param {!Object} options
   * @param {string} name
   * @param {boolean} cl
   * @return {?}
   */
  function next(options, name, cl) {
    /**
     * @return {undefined}
     */
    function callback() {
      var lastSelected = this.selectedIndex;
      if (lastSelected) {
        if (1 == lastSelected) {
          options.app.send([24], [name]);
        } else {
          if (2 == lastSelected) {
            options.app.X("/ignore " + name);
            remove(options, name);
          }
        }
        /** @type {number} */
        this.selectedIndex = 0;
      }
    }
    /** @type {string} */
    var obj = "_" + name;
    if (options.a.hasOwnProperty(obj)) {
      return options.a[obj];
    }
    options.a[obj] = obj = {
      f: null,
      u: null,
      Ia: name,
      Ke: false,
      fa: null,
      fd: null
    };
    obj.f = $(options.H, "div", {
      style: {
        display: "none"
      }
    });
    /** @type {!Array} */
    var navLinksArr = ["...", options.app.j("chd"), options.app.j("ui_block") + " (" + name + ")"];
    obj.u = new Controller(options.app, obj.f, {
      Zd: function (a) {
        if (a.length > n) {
          a = a.substring(0, n) + "...";
        }
        options.app.send([21], [name, a]);
      },
      De: function () {
        options.app.send([23], [name]);
      },
      ld: true,
      He: {
        className: "imfr"
      },
      Wd: {
        className: "imtx",
        style: {
          minHeight: "5em"
        }
      },
      Xe: {
        className: "imin"
      }
    });
    var title = $(obj.u.ab.parentNode.parentNode, "div", {
      className: "imo1"
    });
    $(title, "select", {
      className: "drops",
      onchange: callback
    }, navLinksArr.map(function (mei) {
      return $("option", mei);
    }));
    title = $(obj.f, "p", {
      className: "imo2"
    });
    $(title, "select", {
      className: "drops",
      onchange: callback
    }, navLinksArr.map(function (mei) {
      return $("option", mei);
    }));
    obj.fa = $(options.C, "a", {
      className: "awrap dcpd",
      onclick: function () {
        each(options.app, name);
        return false;
      }
    }, [$("div", {
      className: "clbt",
      onclick: function (event) {
        remove(options, name);
        event.stopPropagation();
        return false;
      }
    }, "X"), $("div", {
      className: "sta"
    }), $("div", {
      className: "ib"
    }, name), obj.fd = $("div", {
      className: "unrd",
      style: {
        visibility: "hidden"
      }
    }, "0")]);
    print(options, name, obj);
    if (cl) {
      options.g.push(name);
    } else {
      options.app.send([reg, 1], [name]);
    }
    return obj;
  }
  /**
   * @param {!Object} options
   * @param {string} name
   * @return {undefined}
   */
  function remove(options, name) {
    /** @type {string} */
    var index = "_" + name;
    if (options.a.hasOwnProperty(index)) {
      var p = options.a[index];
      if (p) {
        p.fa.parentNode.removeChild(p.fa);
        p.f.parentNode.removeChild(p.f);
        delete options.a[index];
        options.app.send([reg, 0], [name]);
        if (options.b == p) {
          request(options, null);
        }
      }
    }
  }
  /**
   * @param {!Object} obj
   * @param {string} name
   * @return {undefined}
   */
  function request(obj, name) {
    /** @type {null} */
    var p = null;
    /** @type {null} */
    var key = null;
    if (name) {
      p = obj.a.hasOwnProperty(key = "_" + name) ? obj.a[key] : next(obj, name);
    }
    if (p != obj.b) {
      if (obj.b) {
        extend(obj.b.f, {
          display: "none"
        });
        callback(obj.b.fa, "slctd", false);
      }
      if (obj.b = p) {
        callback(obj.f, "imact", true);
        resolve(obj.I, [p.Ia]);
        callback(obj.B, "st1", p.Ke);
        extend(p.f, {
          display: "block"
        });
        callback(p.fa, "slctd");
        p.u.Aa();
        if (obj.app.P) {
          p.u.ab.focus();
        }
        normalize(obj);
      } else {
        callback(obj.f, "imact", false);
      }
    }
  }
  /**
   * @param {!Object} options
   * @return {undefined}
   */
  function attach(options) {
    var keys = options.app.eb.slice(0);
    keys.unshift("-- " + options.app.j("bl_cs") + " --");
    /** @type {number} */
    options.A.options.length = 0;
    $(options.A, keys.map(function (mei) {
      return $("option", mei);
    }));
    keys = options.app.Od;
    var end;
    /** @type {number} */
    var i = 0;
    for (; i < keys.length && 90 > i; i++) {
      if (end = options.a["_" + keys[i]]) {
        print(options, keys[i], end);
      } else {
        next(options, keys[i], true);
      }
    }
    if (0 < options.g.length) {
      options.app.send([reg, 1], options.g);
      /** @type {!Array} */
      options.g = [];
    }
    _get(options);
  }
  /**
   * @param {!Object} data
   * @param {!Object} c
   * @param {boolean} a
   * @return {undefined}
   */
  function fmt(data, c, a) {
    if (c = data.a["_" + c]) {
      /** @type {boolean} */
      c.Ke = a;
      callback(c.fa, "st1", a);
      if (c == data.b) {
        callback(data.B, "st1", a);
      }
    }
  }
  /**
   * @param {!Object} options
   * @param {string} i
   * @param {number} type
   * @param {undefined} a
   * @param {?} f
   * @param {string} x0
   * @return {undefined}
   */
  function set(options, i, type, a, f, x0) {
    var data = options.a["_" + i];
    if (type == Z || type == No) {
      if (data) {
        data.u.append(a);
      } else {
        if (f) {
          data = next(options, i);
        }
      }
      if (f) {
        if (!normalize(options, data)) {
          send(options, data);
        }
      }
      if (x0) {
        options = options.A;
        html(options, $("option", i), 1 < options.length ? options.options[1] : null);
      }
    } else {
      if (-1 == type) {
        if (data) {
          data.u.append(a, true);
        }
      } else {
        if (type == SS) {
          if (data) {
            data.u.reset();
          }
          options = options.A.options;
          /** @type {number} */
          type = options.length - 1;
          for (; 0 <= type; type--) {
            if (options[type].text == i) {
              options.remove(type);
              break;
            }
          }
        }
      }
    }
  }
  /**
   * @param {!Object} value
   * @return {undefined}
   */
  function b(value) {
    /** @type {!Object} */
    this.app = value;
    this.Ja = this.app.j("bl_prefs");
    this.f = $(value.B, "div", {
      className: "stvxpad vnarrow",
      style: {
        display: "none"
      }
    });
    var result = this;
    $(this.f, [$("p", [this.b = $("input", {
      type: "checkbox"
    }), this.app.j("p_ignprv")]), $("p", [this.a = $("input", {
      type: "checkbox"
    }), this.app.j("p_prvbud")]), $("p", [this.g = $("input", {
      type: "checkbox"
    }), this.app.j("p_igninv")]), $("p", $("button", {
      className: "minw",
      onclick: function () {
        var data = result.app;
        var border = result.a.checked;
        data.cc = result.b.checked;
        data.bc = border;
        data = result.app;
        data.ec = result.g.checked;
        clear(data);
        debug(result.app, result.app.N);
        return false;
      }
    }, this.app.j("bl_ok")))]);
  }
  /**
   * @param {!Object} target
   * @return {undefined}
   */
  function Request(target) {
    /** @type {!Object} */
    this.app = target;
    /** @type {boolean} */
    this.yd = true;
    /** @type {boolean} */
    this.xd = false;
    this.f = $(target.B, "div", {
      className: "astat bsbb",
      style: {
        display: "none"
      },
      ontouchmove: function () {
        return false;
      }
    });
  }
  /**
   * @param {!Object} value
   * @param {!Object} q
   * @return {undefined}
   */
  function main(value, q) {
    var s = this;
    /** @type {!Object} */
    this.app = value;
    /** @type {!Object} */
    this.o = q;
    this.Ja = q.ub || "-";
    /** @type {number} */
    this.b = 0;
    /** @type {number} */
    this.B = 1;
    /** @type {null} */
    this.A = null;
    /** @type {number} */
    this.g = 0;
    this.f = $(value.B, "div", {
      className: "stview",
      style: {
        display: "none"
      }
    });
    this.C = $(this.f, "p", {
      className: "tac"
    }, $("div", {
      className: "loader"
    }));
    this.fa = $(this.f, "div", {
      className: "turs",
      style: {
        display: "none"
      }
    });
    /**
     * @param {string} e
     * @param {?} undefined
     * @return {?}
     */
    window[this.o.ib] = function (e, undefined) {
      return s.lb(e, undefined);
    };
  }
  /**
   * @param {!Object} data
   * @param {number} visible
   * @return {undefined}
   */
  function reposition(data, visible) {
    extend(data.fa, {
      display: visible ? "block" : "none"
    });
    extend(data.C, {
      display: "none"
    });
    /** @type {number} */
    data.g = 0;
    if (!visible) {
      /** @type {number} */
      var max = data.g = setTimeout(function () {
        if (data.g == max) {
          /** @type {number} */
          data.g = 0;
          extend(data.C, {
            display: "block"
          });
        }
      }, 500);
    }
  }
  /**
   * @param {!Object} self
   * @return {undefined}
   */
  function destroy(self) {
    var response = self.A = new View(self.o.url + (2 == self.b ? "&sk=2&page=" + self.B : ""), redraw(self.o.dd ? {
      jsget: 1,
      ksession: wrapped(self.app)
    } : {
        jsget: 1
      }));
    /**
     * @param {boolean} type
     * @param {?} data
     * @return {?}
     */
    response.handle = function (type, data) {
      return self.handle(response, type, data);
    };
  }
  /**
   * @param {!Object} state
   * @param {!Object} obj
   * @return {undefined}
   */
  function build(state, obj) {
    /**
     * @param {string} method
     * @return {?}
     */
    function maxpage(method) {
      return obj.tx && obj.tx[method] || method;
    }
    var result = obj.tl;
    $(state.fa, "div", {
      className: "bwrap dcpd"
    }, [$("a", {
      href: "",
      onclick: function () {
        debug(state.app, state.app.wd);
        return false;
      }
    }, maxpage("t_hupt")), " | ", $("a", {
      href: "",
      onclick: function () {
        debug(state.app, state.app.wd + "/f");
        return false;
      }
    }, maxpage("t_hfit"))]);
    state.a = $(state.fa, "table", {
      className: "tulst",
      style: {
        width: "100%",
        borderCollapse: "collapse"
      }
    });
    var j;
    for (j in result) {
      var res = result[j];
      var container = state.a.insertRow(-1);
      if (1 == state.b) {
        var element = $(container.insertCell(-1), "a", {
          className: "awrap dcpd hv"
        });
        container = $(element, "div", {
          className: "maxw"
        });
        var div = $(container, "div", {
          className: "bl1"
        });
        /** @type {string} */
        $(div, "div", {
          className: "tid"
        }).innerHTML = '<b class="lc">' + res.id + "</b> (" + res.nop + ")";
        div = $(div, "div", {
          className: "torg"
        });
        /** @type {string} */
        div.innerHTML = "" + res.onm + "";
        div = $(container, "div", {
          className: "bl2"
        });
        container = $(div, "div", {
          className: "tpar"
        });
        container.innerHTML = res.par;
        var overlay = $(div, "div", {
          className: "tdtup"
        });
        overlay.innerHTML = res.dt;
        (function (b) {
          /**
           * @return {?}
           */
          element.onclick = function () {
            state.app.X("/join " + b);
            return false;
          };
        })(res.id);
      } else {
        element = $(container.insertCell(-1), "div", {
          className: "awrap dcpd"
        });
        container = $(element, "div", {
          className: "maxw"
        });
        div = $(container, "div", {
          className: "bl1"
        });
        overlay = $(div, "div", {
          className: "tdtfin"
        });
        overlay.innerHTML = res.dt;
        div = $(div, "div", {
          className: "torg"
        });
        /** @type {string} */
        div.innerHTML = "" + res.onm + "";
        div = $(container, "div", {
          className: "bl2"
        });
        container = $(div, "div", {
          className: "tpar"
        });
        container.innerHTML = res.par;
        $(div, "a", {
          className: "tdt",
          href: res.resurl,
          target: "_blank"
        }, [maxpage("t_trsl"), " (" + res.nop + ")"]);
      }
    }
    result = $(state.fa, "div", {
      className: "dcpd"
    });
    if (2 == state.b && obj.page) {
      $(result, "p", {}, $("a", {
        className: "fb",
        href: "",
        onclick: function () {
          debug(state.app, state.app.wd + "/f/" + (obj.page + 1));
          return false;
        }
      }, maxpage("t_next") + " >"));
    }
    $(result, "p", {
      style: {
        marginTop: "2em"
      }
    }, $("button", {
      className: "minw",
      onclick: function () {
        debug(state.app, state.app.Qe);
      }
    }, maxpage("t_tune")));
    if ((j = obj.toft) && 0 < j.length) {
      res = $(result, "div", {
        style: {
          marginTop: "1.5em"
        }
      });
      $(res, "p", maxpage("t_oftt"));
      res = $(res, "p", {
        className: ""
      });
      var p;
      for (p in j) {
        res.innerHTML += j[p] + "<br />";
      }
    }
    if (obj.tcur) {
      $(result, "p", {
        marginTop: "1.5em"
      }, obj.tcur);
    }
  }
  /**
   * @param {!Object} value
   * @param {!Object} q
   * @return {undefined}
   */
  function r(value, q) {
    var s = this;
    /** @type {!Object} */
    this.app = value;
    /** @type {!Object} */
    this.o = q;
    this.Ja = q.ub || "-";
    /** @type {null} */
    this.b = null;
    /** @type {number} */
    this.a = 0;
    this.f = $(value.B, "div", {
      className: "stvxpad vnarrow",
      style: {
        display: "none"
      }
    });
    this.A = $(this.f, "p", {
      className: "tac"
    }, $("div", {
      className: "loader"
    }));
    this.g = $(this.f, "div", {
      style: {
        display: "none"
      }
    });
    /**
     * @param {string} e
     * @param {?} undefined
     * @return {?}
     */
    window[this.o.ib] = function (e, undefined) {
      return s.lb(e, undefined);
    };
  }
  /**
   * @param {!Object} node
   * @param {number} visible
   * @return {undefined}
   */
  function toggleVisibility(node, visible) {
    extend(node.g, {
      display: visible ? "block" : "none"
    });
    extend(node.A, {
      display: "none"
    });
    /** @type {number} */
    node.a = 0;
    if (!visible) {
      /** @type {number} */
      var el = node.a = setTimeout(function () {
        if (node.a == el) {
          /** @type {number} */
          node.a = 0;
          extend(node.A, {
            display: "block"
          });
        }
      }, 500);
    }
  }
  /**
   * @return {undefined}
   */
  function a() {
  }
  /**
   * @param {!Object} node
   * @param {number} value
   * @return {?}
   */
  function _build(node, value) {
    /** @type {number} */
    var instance = 0;
    /** @type {number} */
    var i = 0;
    for (; i < node.u.length; i++) {
      /** @type {number} */
      var column = 0;
      for (; column < node.u[i].length; column++) {
        if (node.u[i][column] == value) {
          instance++;
        }
      }
    }
    return instance;
  }
  /**
   * @param {?} self
   * @param {!Object} name
   * @param {number} e
   * @param {number} d
   * @return {undefined}
   */
  function format(self, name, e, d) {
    var x = self.ja;
    /** @type {number} */
    var y = self.ra * x;
    self.Fb.drawImage(self.Ac[name], Math.round(self.nb * x + e * y), Math.round(self.tb * x + d * y), y, y);
  }
  /**
   * @param {?} element
   * @param {!Object} width
   * @param {number} height
   * @return {undefined}
   */
  function select(element, width, height) {
    if (element.jc != width || element.qd != height) {
      if (element.jc = width, element.qd = height, element.Ca[1]) {
        if (-1 != element.jc) {
          var $scope = element.Ca[1];
          /** @type {boolean} */
          $scope.Gb = true;
          /** @type {!Object} */
          $scope.Ae = width;
          /** @type {number} */
          $scope.Be = height;
          if (!$scope.kd) {
            element.oc(1, $scope.Ob, width, height);
          }
        } else {
          if (element.Ca[1].Gb) {
            /** @type {boolean} */
            element.Ca[1].Gb = false;
            extend(element.Ca[1].Ob, {
              display: "none"
            });
          }
        }
      }
    }
  }
  /**
   * @param {?} options
   * @param {number} i
   * @param {number} x
   * @return {?}
   */
  function put(options, i, x) {
    i = 8 * x + i;
    /** @type {number} */
    x = 0;
    for (; x < options.uc.length; x++) {
      if (options.uc[x] == i) {
        return true;
      }
    }
    return false;
  }
  /**
   * @param {number} val
   * @return {?}
   */
  function recurse(val) {
    return 0 > val ? "--" : String.fromCharCode(val % 8 + 97, Math.floor(val / 8) % 8 + 49);
  }
  /**
   * @return {undefined}
   */
  function me() {
  }
  /**
   * @return {undefined}
   */
  function e() {
    show.call(this, {
      Te: a,
      tf: true
    });
  }
  /**
   * @param {number} a
   * @param {number} b
   * @return {?}
   */
  function String(a, b) {
    /** @type {!Array} */
    var h = [0, 1, 1, 1, 0, -1, -1, -1];
    /** @type {!Array} */
    var ev = [-1, -1, 0, 1, 1, 1, 0, -1];
    /** @type {!Array} */
    var str = [];
    var table = a.u.u;
    a = a.u.Ib;
    /** @type {number} */
    var c = 0;
    for (; c < a; c++) {
      /** @type {number} */
      var b = 0;
      for (; b < a; b++) {
        if (-1 == table[c][b]) {
          /** @type {number} */
          var j = 0;
          for (; 8 > j; j++) {
            /** @type {number} */
            var i = b;
            /** @type {number} */
            var t = c;
            /** @type {number} */
            var cols = 0;
            for (; ;) {
              i = i + h[j];
              t = t + ev[j];
              cols++;
              if (0 > i || i >= a || 0 > t || t >= a) {
                break;
              }
              if (table[t][i] == b) {
                if (1 < cols) {
                  str.push(c * a + b);
                  /** @type {number} */
                  j = 8;
                }
                break;
              } else {
                if (-1 == table[t][i]) {
                  break;
                }
              }
            }
          }
        }
      }
    }
    return str;
  }
  var self;
  /** @type {!Function} */
  var defineProperty = "function" == typeof Object.defineProperties ? Object.defineProperty : function (object, name, descriptor) {
    if (object != Array.prototype && object != Object.prototype) {
      object[name] = descriptor.value;
    }
  };
  var _key = "undefined" != typeof window && window === this ? this : "undefined" != typeof global && null != global ? global : this;
  analyzeAll("Array.prototype.fill", function (position) {
    return position ? position : function (posts, n, i) {
      var len = this.length || 0;
      if (0 > n) {
        /** @type {number} */
        n = Math.max(0, len + n);
      }
      if (null == i || i > len) {
        i = len;
      }
      /** @type {number} */
      i = Number(i);
      if (0 > i) {
        /** @type {number} */
        i = Math.max(0, len + i);
      }
      /** @type {number} */
      n = Number(n || 0);
      for (; n < i; n++) {
        this[n] = posts;
      }
      return this;
    };
  });
  analyzeAll("Object.assign", function (position) {
    return position ? position : function (win, canCreateDiscussions) {
      /** @type {number} */
      var i = 1;
      for (; i < arguments.length; i++) {
        var context = arguments[i];
        if (context) {
          var key;
          for (key in context) {
            if (Object.prototype.hasOwnProperty.call(context, key)) {
              win[key] = context[key];
            }
          }
        }
      }
      return win;
    };
  });
  self = init.prototype;
  /**
   * @param {number} i
   * @param {string} m
   * @return {undefined}
   */
  self.xc = function (i, m) {  // lk16:handle websocket message
    switch (i[0]) {
      case last: // lk16:31 recv get client info
        /** @type {number} */
        i = 0;
        for (; i + 1 < m.length; i = i + 2) {
          this.td(m[i], m[i + 1]);
        }
        /** @type {boolean} */
        this.ua = true;
        _init(this);
        break;
      case xa: // lk16:23 recv send some unknown options
        if (1 > m.length) {
          break;
        }
        m = m[0];
        /** @type {string} */
        this.nc = "";
        m = m.split("&");
        /** @type {number} */
        i = 0;
        for (; i < m.length; i++) {
          var str = m[i];
          var MimeTypePos = str.indexOf("=");
          if (-1 != MimeTypePos) {
            var keys1 = str.substr(0, MimeTypePos);
            str = str.substr(MimeTypePos + 1).replace(/%3F/g, "&");
            /** @type {boolean} */
            MimeTypePos = "1" == str || "true" == str;
            if ("noi" == keys1) {
              /** @type {boolean} */
              this.ec = MimeTypePos;
            } else {
              if ("nop" == keys1) {
                /** @type {boolean} */
                this.cc = MimeTypePos;
              } else {
                if ("prb" == keys1) {
                  /** @type {boolean} */
                  this.bc = MimeTypePos;
                } else {
                  if ("snd" == keys1) {
                    /** @type {boolean} */
                    this.qa = MimeTypePos;
                  } else {
                    if (null != this.Mb && this.Mb == keys1) {
                      /** @type {boolean} */
                      this.Ec = MimeTypePos;
                    } else {
                      if (0 < keys1.length && "_" == keys1.charAt(0)) {
                        this.nc += "&" + keys1 + "=" + str;
                      }
                    }
                  }
                }
              }
            }
          }
        }
        break;
      case ya: // lk16:51 recv key value pairs
        this.text = {};
        /** @type {number} */
        i = 0;
        for (; i < m.length; i = i + 2) {
          this.text[m[i]] = m[i + 1];
        }
        /** @type {string} */
        this.text.t_sf = "pl" != this.lang ? "Send feedback (English)" : "Prze\u015blij uwagi:";
        /** @type {string} */
        this.text.t_fb = "pl" != this.lang ? "feedback" : "uwagi";
        if ((m = this.text.gname) && 0 < m.length) {
          this.text.gname = m.toUpperCase();
        }
        if (!this.Ka) {
          this.Ic();
          /** @type {boolean} */
          this.Ka = true;
        }
        break;
      case za: // lk16:17 recv unknown
        if (1 > m.length) {
          break;
        }
        m = {
          gd: m[0],
          link: 3 <= m.length ? {
            href: m[1],
            target: "_blank"
          } : null,
          mf: 3 <= m.length ? m[2] : null
        };
        /** @type {boolean} */
        this.ua = false;
        /** @type {boolean} */
        this.Ad = true;
        generate(this, m);
        break;
      case Aa: // lk16:19 recv unknown
        if (2 > i.length || !this.S) {
          break;
        }
        try {
          window.localStorage.setItem("k2ver", i[1]);
          window.location.reload();
        } catch (g) {
          /** @type {boolean} */
          this.ua = false;
          /** @type {boolean} */
          this.Ad = true;
          generate(this, {
            gd: "VERSION"
          });
        }
        break;
      case Ba: // lk16:18 recv username,language,?,game
        if (!(2 > i.length || 2 > m.length)) {
          this.Ga = i[1];
          if (2 < i.length) {
            /** @type {boolean} */
            this.wc = 0 != i[2];
          }
          this.sd = m[0];
          this.lang = m[1];
          /** @type {string} */
          this.zd = "pl" == this.lang ? "KURNIK" : "PlayOK";
          if (2 < m.length && 0 < m[2].length) {
            try {
              window.localStorage.setItem("autoid", m[2]);
            } catch (g) {
            }
          }
          if (3 < m.length) {
            this.zb = m[3];
          }
          /** @type {number} */
          i = 4;
          for (; i < m.length; i++) {
            keys1 = m[i].split(":", 2);
            if (1 < keys1.length && !this.Ld.hasOwnProperty(keys1[0])) {
              this.Ld[keys1[0]] = keys1[1];
            }
          }
        }
    }
  };
  /**
   * @param {string} key
   * @param {string} e
   * @return {?}
   */
  self.j = function (key, e) {
    return this.text[key] || window.k2text && window.k2text[key] || e || key;
  };
  /**
   * @param {!Array} result
   * @param {?} data
   * @return {undefined}
   */
  self.send = function (result, data) {
    if (this.Bd) {
      this.Bd.send(result, data);
    }
  };
  /**
   * @param {string} a
   * @return {undefined}
   */
  self.X = function (a) {
    if (a.length > n) {
      a = a.substring(0, n) + "...";
    }
    this.send([workerId], [a]);
  };
  /**
   * @return {?}
   */
  self.af = function () {
    return this.Ec;
  };
  /**
   * @param {string} element
   * @param {string} style
   * @return {undefined}
   */
  self.td = function (element, style) {
    if ("set_langsymbols" == element) {
      this.Cc = style.split(" ").map(function (a) {
        return "(" + a + ")";
      });
    }
  };
  /**
   * @return {undefined}
   */
  self.se = function () {
  };
  /**
   * @return {undefined}
   */
  self.Ic = function () {
  };
  /**
   * @return {?}
   */
  self.re = function () {
    return "h";
  };
  /**
   * @param {string} done
   * @return {?}
   */
  self.cd = function (done) {
    return "status" == done ? this.sc = new Request(this) : null;
  };
  /**
   * @return {undefined}
   */
  self.yc = function () {
  };
  /**
   * @return {undefined}
   */
  self.na = function () {
    if (this.fb) {
      clearTimeout(this.fb);
      /** @type {number} */
      this.fb = 0;
    }
    /** @type {number} */
    var b = window.innerWidth;
    /** @type {number} */
    var i = window.innerHeight;
    /** @type {number} */
    b = this.Pd ? 500 >= b ? 0 : 860 > b ? 1 : 2 : 2;
    /** @type {boolean} */
    var totalContributionAmountSatoshis = false;
    if (this.xa != b) {
      callback(this.a.parentNode, "vm" + this.xa, false);
      /** @type {number} */
      this.xa = b;
      callback(this.a.parentNode, "vm" + this.xa);
      /** @type {boolean} */
      this.oa = 2 <= b;
      /** @type {boolean} */
      totalContributionAmountSatoshis = true;
    }
    if (this.Pb) {
      totalContributionAmountSatoshis = this.wa;
      /** @type {number} */
      var w = Math.round(totalContributionAmountSatoshis * this.hf);
      /** @type {number} */
      var j = Math.round(.95 * w);
      /** @type {number} */
      var d = i;
      if (w <= i - 48) {
        /** @type {number} */
        d = w;
      } else {
        if (j <= i - 48) {
          /** @type {number} */
          d = i - 48;
        }
      }
      if (2 > b) {
        /** @type {number} */
        d = i;
      }
      /** @type {number} */
      b = Math.min(Math.floor((i - d) / 2), 36);
      this.Yd = {
        Af: totalContributionAmountSatoshis,
        gf: Math.floor(d),
        ff: i - b,
        top: b
      };
      if (this.Rd && 0 < b) {
        extend(this.Rd, {
          top: -(b >> 1) + "px"
        });
      }
      /** @type {boolean} */
      totalContributionAmountSatoshis = true;
    }
    if (totalContributionAmountSatoshis) {
      _init(this);
    } else {
      this.yc(false);
    }
  };
  /** @type {!Object} */
  move.prototype = Object.create(init.prototype);
  /** @type {!Object} */
  self = move.prototype;
  /** @type {function(?): undefined} */
  self.constructor = move;
  /**
   * @return {undefined}
   */
  self.Ic = function () {
    /**
     * @return {?}
     */
    function handler() {
      return [$("b", params.sd), params.o.ob ? null : $("div", {
        className: "mlh r" + merge(params, params.Fa)
      }), params.o.ob ? null : $("span", {
        className: "snum"
      }, exec(params, params.Fa))];
    }
    /**
     * @param {number} relative
     * @return {?}
     */
    function parse(relative) {
      return [$("div", {
        className: "spbb"
      }), " (" + relative + ")"];
    }
    /**
     * @param {number} layout
     * @return {?}
     */
    function showLayoutPreview(layout) {
      return [$("div", {
        className: "spbb"
      }), " (" + layout + ")"];
    }
    var params = this;
    var data = this;
    this.$b = $(this.H, "div", {
      className: "dclpd bsbb navcont",
      style: {
        display: "none"
      }
    });
    var checkUncheckAllButton;
    var inputElements;
    var elem;
    var node;
    $(this.$b, "div", {
      className: "newtab1"
    }, [checkUncheckAllButton = $("button", {
      className: "butsys minwd",
      onclick: function () {
        data.send([payload], null);
      }
    }, this.j("bl_newtab")), " ", $("div", {
      className: "selcwr mro"
    }, [inputElements = $("button", {
      className: "selcbt butsys vsel"
    }, ["-"]), elem = $("select", {
      className: "selcsl",
      onchange: function (event) {
        if (event = event.target.options[event.target.selectedIndex]) {
          data.X("/join " + event.text.split(" ")[0]);
        }
      }
    })]), node = $("span", {
      className: "tuinfo fb"
    }, "-")]);
    expect(this, "rooms", function (self) {
      /** @type {number} */
      elem.options.length = 0;
      $(elem, self.list.map(function (body, state) {
        return $("option", state == self.rb ? {
          selected: true
        } : {}, body);
      }));
      resolve(inputElements, [self.list[0 <= self.rb ? self.rb : 0].split("(")[0] || "-"]);
    });
    expect(this, "tumode", function (a) {
      /** @type {boolean} */
      checkUncheckAllButton.disabled = !!a;
    });
    expect(this, "tuinfo", function (ease) {
      return resolve(node, [ease || "-"]);
    });
    var start = $(this.$b, "div", {
      className: "nav ib"
    });
    var self = {};
    self.h = $(start, "button", {
      className: "bmain",
      onclick: function () {
        debug(data, data.N);
      }
    }, this.j("bl_tbs"));
    if (0 == this.Ga) {
      var currentFlexChild = self.c = $(start, "button", {
        onclick: function () {
          debug(data, "c");
        }
      }, [data.j("bl_buds") + " (0)"]);
      var value = self.m = $(start, "button", {
        onclick: function () {
          debug(data, "m");
        }
      }, parse(0));
      if (this.wc) {
        self.t = $(start, "button", {
          onclick: function () {
            debug(data, "t");
          }
        }, this.j("t_turs"));
      }
    }
    self.e = $(start, "button", {
      onclick: function () {
        debug(data, "e");
      }
    }, this.j("bl_mr"));
    var y = self.x = $(start, "button", {
      className: "btab",
      onclick: function () {
        if (0 < data.tab.F) {
          debug(data, data.tab.F.toString());
        }
      }
    }, "#000");
    expect(this, "tabalert", function (a) {
      return callback(y, "alert", a);
    });
    expect(this, "tabopen", function (goodDots) {
      return resolve(y, [0 < goodDots ? "#" + goodDots : "#000"]);
    });
    if (currentFlexChild) {
      expect(this, "ncontacts", function (a) {
        return resolve(currentFlexChild, [data.j("bl_buds") + " (" + a + ")"]);
      });
    }
    if (value) {
      expect(this, "nmessages", function (dep) {
        resolve(value, parse(dep));
        callback(value, "alert", 0 < dep);
      });
    }
    expect(this, "nav", function (item) {
      Object.keys(self).forEach(function (name) {
        return callback(self[name], "active", name == item);
      });
    });
    this.Yb = $(this.H, "div", {
      className: "navttl fb fl",
      style: {
        display: "none"
      }
    });
    var a = $(this.a, "div", {
      className: "nav0 usno tama",
      style: {
        zIndex: cluezIndex + 1
      }
    });
    var download = $(a, "button", {
      className: "mbut hdhei hdbwd",
      onclick: function () {
        acceptTheCall(a);
        return false;
      }
    }, [$("div", {
      className: "micon"
    })]);
    document.addEventListener("click", function (e) {
      if (!(!transform(a, "nav0open") || a.contains(e.target) && !t.contains(e.target))) {
        acceptTheCall(a);
        if (!t.contains(e.target)) {
          e.stopPropagation();
        }
      }
    }, true);
    expect(this, "chatalert", function (a) {
      return callback(download, "alert", a);
    });
    start = $(a, "div", {
      className: "mcont"
    });
    var t = $(start, "div", {
      className: "mlst"
    });
    var d = {};
    var k = d.x = $(t, "button", {
      className: "btab",
      onclick: function () {
        if (0 < data.tab.F) {
          debug(data, data.tab.F.toString());
        }
      }
    }, "#000");
    d.h = $(t, "button", {
      onclick: function () {
        debug(data, data.N);
      }
    }, this.j("bl_tbs"));
    if (0 == this.Ga) {
      var defaultSounds = d.c = $(t, "button", {
        onclick: function () {
          debug(data, "c");
        }
      }, [data.j("bl_buds") + " (0)"]);
      var valueProgess = d.m = $(t, "button", {
        onclick: function () {
          debug(data, "m");
        }
      }, showLayoutPreview(0));
      if (this.wc) {
        d.t = $(t, "button", {
          onclick: function () {
            debug(data, "t");
          }
        }, this.j("t_turs"));
      }
    } else {
      if (!this.P) {
        d.l = $(t, "button", {
          onclick: function () {
            if (data.Ab) {
              debug(data, "l");
            } else {
              logout(data);
            }
          }
        }, this.j("t_lgin"));
      }
    }
    d.e = $(t, "button", {
      onclick: function () {
        debug(data, "e");
      }
    }, this.j("bl_mr"));
    expect(this, "tabalert", function (a) {
      return callback(k, "alert", a);
    });
    expect(this, "tabopen", function (goodDots) {
      return resolve(k, [0 < goodDots ? "#" + goodDots : "#000"]);
    });
    if (defaultSounds) {
      expect(this, "ncontacts", function (a) {
        return resolve(defaultSounds, [data.j("bl_buds") + " (" + a + ")"]);
      });
    }
    if (valueProgess) {
      expect(this, "nmessages", function (layout) {
        resolve(valueProgess, showLayoutPreview(layout));
        callback(valueProgess, "alert", 0 < layout);
      });
    }
    expect(this, "nav", function (item) {
      Object.keys(d).forEach(function (i) {
        return callback(d[i], "active", i == item);
      });
    });
    start = $(start, "div", {
      className: "msub"
    });
    var found = $(start, "p", handler());
    expect(this, "urank", function () {
      return resolve(found, handler());
    });
    if (!(this.S || this.P)) {
      $(start, "p", [$("button", {
        onclick: function () {
          /** @type {string} */
          document.body.innerHTML = "";
          window.location = window.k2url.home;
        }
      }, this.zd)]);
    }
    expect(this, "tabopen", function (a) {
      if (0 < a) {
        setTimeout(function () {
          return callback(data.a, "navtabopen", 0 < a);
        }, 0);
      } else {
        callback(data.a, "navtabopen", 0 < a);
      }
    });
    expect(this, "tumode", function (a) {
      return callback(data.a, "tumode", !!a);
    });
    text(this, "h");
  };
  /**
   * @return {undefined}
   */
  self.se = function () {
    insert(this.g, [0, 0, 0], []);
    if (this.L) {
      clearInterval(this.L);
      /** @type {number} */
      this.L = 0;
    }
    if (this.tab) {
      this.tab.reset();
    }
  };
  /**
   * @param {string} i
   * @param {string} element
   * @return {undefined}
   */
  self.td = function (i, element) {
    init.prototype.td.call(this, i, element);
    if ("set_tab_mt" == i) {
      this.Md = element.split(" ");
    } else {
      if ("set_tab_at" == i) {
        this.Ed = element.split(" ");
      } else {
        if ("set_tab_gt" == i) {
          this.Hd = element.split(" ");
        } else {
          if ("set_rank" == i) {
            this.W = element.split(" ").filter(function (a, b) {
              return 0 == b % 2;
            }).map(function (id_local) {
              return parseInt(id_local, 10);
            });
          }
        }
      }
    }
  };
  /**
   * @param {?} match
   * @param {!Array} data
   * @return {undefined}
   */
  self.xc = function (match, data) { // lk16:handle websocket message
    var t = this;
    var i;
    if (match[0] >= lb) {
      if (1 < match.length && this.tab && this.tab.F == match[1]) {
        render(this.tab, match, data);
      }
    } else {
      switch (match[0]) {
        case CSI:
          if (1 > data.length) {
            break;
          }
          /** @type {!Array} */
          this.eb = [];
          this.pa = {};
          match = data[0].split(" ");
          /** @type {number} */
          i = 0;
          for (; i < match.length; i++) {
            var event = match[i];
            /** @type {string} */
            var index = "_" + event;
            if (!this.pa.hasOwnProperty(index)) {
              this.pa[index] = {
                Ia: event,
                Xd: 1
              };
              this.eb.push(event);
            }
          }
          this.Od = 1 < data.length && 0 < data[1].length ? data[1].split(" ") : [];
          if (this.M) {
            attach(this.M);
          }
          break;
        case wink:
          if (2 > data.length) {
            break;
          }
          event = 1 < match.length ? match[1] : 0;
          match = 2 < match.length ? match[2] : No;
          /** @type {boolean} */
          index = false;
          var k = data[1];
          /** @type {string} */
          var key = "_" + k;
          var item = this.I.hasOwnProperty(key) ? this.I[key] : null;
          if (match == Z || match == No) {
            if (!item) {
              this.I[key] = item = {
                Ia: k,
                rc: "",
                Oa: 0,
                tc: 0
              };
            }
            item.rc = "" != item.rc ? item.rc + "\n" + data[0] : data[0];
            if (0 != event) {
              item.Oa = event;
              /** @type {number} */
              var shiftTc = 0;
              var queryStringKeysArray = data[0].split("\n");
              /** @type {number} */
              i = 0;
              for (; i < queryStringKeysArray.length; i++) {
                if (0 != queryStringKeysArray[i].indexOf(element)) {
                  shiftTc++;
                }
              }
              item.tc += shiftTc;
            }
            if (!this.pa.hasOwnProperty(key)) {
              this.pa[key] = {
                Ia: k,
                Xd: 0
              };
              this.eb.unshift(k);
              /** @type {boolean} */
              index = true;
            }
          } else {
            if (match == SS) {
              if (item) {
                delete this.I[key];
              }
              if (this.pa.hasOwnProperty(key)) {
                delete this.pa[key];
                if (-1 != (i = this.eb.indexOf(k))) {
                  this.eb.splice(i, 1);
                }
              }
            }
          }
          if (this.M) {
            set(this.M, k, match, data[0], event, index);
          }
          if (!this.M) {
            rebuildModelFromFields(this);
          }
          break;
        case ease_in:
          if (1 > match.length || !this.M) {
            break;
          }
          /** @type {number} */
          i = 1;
          for (; i < match.length && !(i - 1 >= data.length); i++) {
            fmt(this.M, data[i - 1], match[i]);
          }
          break;
        case ease_out: // lk16:20 recv "+ <game name>"
          if (1 > data.length || !this.g) {
            break;
          }
          if (this.g) {
            refresh(this.g, data[0]);
          }
          break;
        case read_write: // lk16:73 recv table settings on join? i:[73,room ID, ?, ?, ?] s:[table settings, player, player]
          if (2 > match.length) {
            break;
          }
          if (2 < match.length) {
            if (!this.tab) {
              text(this, "x");
            }
            run(this.tab, match, data);
          }
          debug(this, match[1].toString(), true);
          break;
        case eslint:
          if (2 > match.length) {
            break;
          }
          if (proceed() == match[1]) {
            then(this);
          }
          if (this.tab && this.tab.F == match[1]) {
            this.tab.reset();
          }
          break;
        case visited: // lk16:71 recv initial table settings i:[71, ?, ?] + many [room id, ?, ?, ?] s: many [table settings, player, player]
          if (3 > match.length) {
            break;
          }
          insert(this.g, match, data);
          if (!this.tab) {
            setTimeout(function () {
              if (!t.tab) {
                text(t, "x");
              }
            }, 100);
          }
          break;
        case valid: // lk16:70 recv table settings update i:[70, room ID, ?, ?, ?] s:[table settings, player, player]
          if (1 > data.length || 3 > match.length) {
            break;
          }
          if (this.g.a.hasOwnProperty(match[1])) {
            message(this.g, match, data);
          } else {
            jsonp(this.g, match, data);
          }
          if (this.tab && this.tab.F == match[1]) {
            on(this.tab, match, data);
          }
          break;
        case active: // lk16:72 recv i:[72,room ID] s absent
          if (2 > match.length) {
            break;
          }
          inverse(this.g, match[1]);
          break;
        case focus:
          if (1 > data.length) {
            break;
          }
          doSearch(this, data[0]);
          break;
        case u:
          if (2 > data.length || 2 > match.length) {
            break;
          }
          doSearch(this, data[0], [overwrite, match[1]], null, [Kb, match[1]], [data[1]]);
          break;
        case x: // lk16:24 recv logout? s:[username]
          if (1 > data.length) {
            break;
          }
          if (this.A.hasOwnProperty(data[0])) {
            attr(this.A[data[0]]);
            delete this.A[data[0]];
          }
          break;
        case X: // lk16:25 recv i:[25, ?, table id?, rating] s:[username]
          if (1 > data.length || 2 > match.length) {
            break;
          }
          if (this.A.hasOwnProperty(data[0])) {
            tick(this.A[data[0]], this.mc, match.slice(1));
          } else {
            match = new Connection(this, data[0], this.mc, match.slice(1));
            this.A[data[0]] = match;
            /** @type {number} */
            i = this.vb.length - 1;
            for (; 0 <= i; i--) {
              matches(this.vb[i], match);
            }
          }
          break;
        case v: // lk16:27 recv active users, i:[27, ?, ?] + for each user [?, ?, rating] s: many [username]
          if (3 > match.length) {
            break;
          }
          drawTextCss(this);
          i = match[1];
          event = match[2];
          /** @type {number} */
          index = 3;
          /** @type {number} */
          k = 0;
          for (; index + i <= match.length && k + 1 + event <= data.length;) {
            this.A[data[k]] = new Connection(this, data[k], this.mc, match.slice(index, index + i));
            index = index + i;
            k = k + (1 + event);
          }
          /** @type {number} */
          i = this.vb.length - 1;
          for (; 0 <= i; i--) {
            func(this.vb[i], this.A);
          }
          break;
        case Rb:
          /** @type {number} */
          i = 0;
          for (; i < match.length - 1 && 2 * i < data.length; i++) {
            if (this.b.hasOwnProperty(data[2 * i]) && (event = this.b[data[2 * i]])) {
              event.status = match[1 + i];
              if (2 * i + 1 < data.length) {
                event.$a = data[2 * i + 1];
              }
              downandpress(event, this.j("st_st1"));
              if (this.ba) {
                initiate(this.ba, event);
              }
            }
          }
          /** @type {number} */
          var url_str_dir = 0;
          Object.keys(this.b).forEach(function (n) {
            if (t.b[n].status) {
              url_str_dir++;
            }
          }, this);
          cb(this, "ncontacts", url_str_dir);
          break;
        case Ub: // lk16:28 recv friends
          if (2 > match.length) {
            break;
          }
          if (0 == match[1]) {
            Object.keys(this.b).forEach(function (n) {
              delete t.b[n];
            }, this);
            /** @type {number} */
            i = 0;
            for (; i < data.length - 1; i = i + 2) {
              if (!this.b.hasOwnProperty(data[i])) {
                event = {
                  name: data[i],
                  status: 0,
                  $a: data[i + 1]
                };
                downandpress(event, this.j("st_st1"));
                this.b[data[i]] = event;
              }
            }
            if (this.ba) {
              m(this.ba, this.b);
            }
          } else {
            if (1 == match[1]) {
              if (3 > match.length || 2 > data.length) {
                break;
              }
              if (this.b.hasOwnProperty(data[0])) {
                break;
              }
              event = {
                name: data[0],
                status: match[2],
                $a: data[1]
              };
              downandpress(event, this.j("st_st1"));
              this.b[data[0]] = event;
              if (this.ba) {
                getDebugTransaction(this.ba, event); // lk16:todo check
              }
            } else {
              if (-1 == match[1]) {
                if (1 > data.length) {
                  break;
                }
                if (!this.b.hasOwnProperty(data[0])) {
                  break;
                }
                event = this.b[data[0]];
                if (!event) {
                  break;
                }
                delete this.b[data[0]];
                if (this.ba) {
                  calc(this.ba, event);
                }
              }
            }
          }
          /** @type {number} */
          var jobsFinishPass = 0;
          Object.keys(this.b).forEach(function (n) {
            if (t.b[n].status) {
              jobsFinishPass++;
            }
          }, this);
          cb(this, "ncontacts", jobsFinishPass);
          break;
        case Yb:
          break;
        case Zb:
          if (1 > data.length || !this.g) {
            break;
          }
          /** @type {number} */
          i = 0;
          if (1 < match.length) {
            i = match[1];
          }
          if (this.L) {
            clearInterval(this.L);
            /** @type {number} */
            this.L = 0;
          }
          if (0 > i) {
            /**
             * @param {number} k
             * @return {undefined}
             */
            var func = function (k) {
              cb(t, "tuinfo", 0 > k ? "(?)" : Math.floor(k / 60) + ":" + Math.floor(k % 60 / 10) % 10 + k % 60 % 10);
            };
            func(-i);
            /** @type {number} */
            this.Sd = Date.now() + 1E3 * -i + 100;
            /** @type {number} */
            this.L = setInterval(function () {
              /** @type {number} */
              var ii = Math.floor((t.Sd - Date.now()) / 1E3);
              func(ii);
              if (0 > ii && t.L) {
                clearInterval(t.L);
                /** @type {number} */
                t.L = 0;
              }
            }, 1E3);
          }
          if (0 <= i) {
            cb(this, "tuinfo", data[0]);
          }
          break;
        case $b: // lk16:30 recv set_tournaments, set_cols, set_chat
          plugin(this.g);
          drawTextCss(this);
          /** @type {number} */
          i = 0;
          match = data.length;
          for (; i + 1 < match; i = i + 2) {
            unwrap(this, data[i], data[i + 1]);
          }
          break;
        case bc: // lk16:32 recv i[32, ?] s:[room info separated by \n in one string]
          if (2 > match.length || 1 > data.length) {
            break;
          }
          data = data[0].split("\n");
          if (this.lc) {
            data.unshift("-");
          }
          cb(this, "rooms", {
            list: data,
            rb: match[1]
          });
          if (this.df) {
            do_export(this);
            debug(this, this.N);
          } else {
            /** @type {boolean} */
            this.df = true;
          }
          break;
        case cc:
          if (2 > match.length || 1 > data.length) {
            break;
          }
          if (1 == data.length) {
            doSearch(this, data[0], [prepend, 1], null, [prepend, 0], null);
          } else {
            doSearch(this, data[0], [prepend, 2], [data[1]], [prepend, 0], [data[1]]);
          }
          break;
        case gc: // lk16:22 recv role?
          if (2 > match.length) {
            break;
          }
          if (this.gc == match[1]) {
            break;
          }
          this.gc = match[1];
          cb(this, "urole", this.gc);
          break;
        case hc: // lk16:33 recv user rating
          if (2 > match.length) {
            break;
          }
          this.Fa = match[1];
          cb(this, "urank", this.Fa);
          break;
        case ic:
          if (1 > data.length) {
            break;
          }
          initialize(this, data[0], match, data);
          break;
        default:
          init.prototype.xc.call(this, match, data);
      }
    }
  };
  /**
   * @return {undefined}
   */
  self.yc = function () {
    if (this.tab && this.tab.f == this.C) {
      var item = this.tab;
      /** @type {boolean} */
      item.dc = !item.app.P;
      if (item.dc) {
        if (transform(item.f, "sbfixed")) {
          callback(item.f, "sbfixed", false);
        }
        callback(item.f, "gvhead", true);
        callback(item.f, "sbdrop");
      } else {
        if (transform(item.f, "sbdrop")) {
          callback(item.f, "sbdrop", false);
        }
        callback(item.f, "gvhead", false);
        callback(item.f, "sbfixed");
      }
      callback(item.app.a, "doddmenu", item.dc);
      item.u.na();
    }
  };
  /**
   * @param {string} undefined
   * @return {?}
   */
  self.re = function (undefined) {
    /** @type {number} */
    var b = parseInt(undefined, 10);
    return !isNaN(b) && 0 < b ? "x" : "f" == undefined || "p" == undefined || "c" == undefined || "m" == undefined || "e" == undefined || 0 == this.Ga && "i" == undefined || 0 == this.Ga && ("t" == undefined || "n" == undefined) || this.Ab && ("l" == undefined || "s" == undefined) ? undefined : "h";
  };
  /**
   * @param {string} undefined
   * @return {?}
   */
  self.cd = function (undefined) {
    var message = this;
    return "h" == undefined ? this.g = new test(this, this.hc) : "x" == undefined ? (this.tab = new (this.o.wf || Object.constructor)(this), this.tab.ha(this), this.tab) : "c" == undefined ? this.ba = new form(this) : "m" == undefined ? this.M = new start(this) : "e" == undefined ? new handle(this) : "f" == undefined ? new go(this) : "p" == undefined ? new b(this) : "l" == undefined && this.Ab ? new r(this, {
      url: get(this, "login"),
      ib: "k2qlog",
      ub: this.j("t_lgin") + " (" + this.zd + ")",
      lb: function (a22) {
        if ("v_signup" == a22) {
          debug(message, "s");
        }
      }
    }) : "s" == undefined && this.Ab ? new r(this, {
      url: get(this, "register"),
      ib: "k2qreg",
      ub: this.j("t_hdrg")
    }) : "i" == undefined ? new r(this, {
      url: get(this, "profile"),
      ib: "k2qprof",
      dd: true,
      ub: this.j("t_prof")
    }) : "t" == undefined ? new main(this, {
      url: get(this, "tourns"),
      ib: "k2qtl",
      ub: this.j("t_turs")
    }) : "n" == undefined ? new r(this, {
      url: get(this, "newtourn"),
      ib: "k2qnt",
      dd: true,
      ub: this.j("t_turs")
    }) : init.prototype.cd.call(this, undefined);
  };
  self = show.prototype;
  /**
   * @return {undefined}
   */
  self.xe = function () {
  };
  /**
   * @return {?}
   */
  self.j = function () {
    return "";
  };
  /**
   * @return {undefined}
   */
  self.na = function () {
    this.u.na();
  };
  /**
   * @param {number} fallback
   * @return {?}
   */
  self.mb = function (fallback) {
    return fallback;
  };
  /**
   * @return {undefined}
   */
  self.ce = function () {
  };
  /**
   * @return {?}
   */
  self.he = function () {
    return "";
  };
  /**
   * @return {?}
   */
  self.fe = function () {
    return "";
  };
  /**
   * @return {undefined}
   */
  self.start = function () {
    var local = this;
    this.na();
    /** @type {boolean} */
    this.I = window.k2pback.m2 ? true : false;
    resolve(this.M, this.I ? [this.B[1] = $("input", {
      type: "radio",
      name: "pbpt",
      onchange: function () {
        update(local, 1, 0);
      }
    }), "1", this.B[2] = $("input", {
      type: "radio",
      name: "pbpt",
      onchange: function () {
        update(local, 2, 0);
      }
    }), "2"] : []);
    var selector = search();
    update(this, selector.Xb, selector.Cb);
    createElement(this, false);
  };
  /**
   * @param {number} x
   * @return {undefined}
   */
  self.ye = function (x) {
    var valueProgess = this;
    var width = this.a.length;
    this.Y = x >= width ? width - 1 : x;
    /** @type {number} */
    this.ga = (this.Y + 1) % 2;
    if (width = this.he(this.Y)) {
      resolve(this.L, $("b", [width + (this.b ? "\u00a0 " + this.b : "")]));
    }
    resolve(this.N, this.a.map(function (targetFormIdAndName, deps) {
      return $("span", {
        style: this.Y == deps ? {
          color: "#c22"
        } : {},
        onclick: function () {
          require(valueProgess, deps);
        }
      }, [("undefined" === typeof this.Sc.Mf && 0 == deps % 2 ? 1 + (deps >> 1) + ". " : "") + targetFormIdAndName, this.Sc.Of ? $("br") : " "]);
    }, this));
    if ((x || window.location.hash) && window.history.replaceState) {
      window.history.replaceState({}, "", "#" + x + (1 < this.A ? "/" + this.A : ""));
    }
  };
  self = Tool.prototype;
  /** @type {number} */
  self.ja = 1;
  /** @type {null} */
  self.$ = null;
  /** @type {null} */
  self.f = null;
  /** @type {null} */
  self.Jb = null;
  /** @type {null} */
  self.Fb = null;
  /** @type {number} */
  self.hb = 0;
  /** @type {number} */
  self.gb = 0;
  /** @type {number} */
  self.nb = 0;
  /** @type {number} */
  self.tb = 0;
  /** @type {number} */
  self.Hb = -1;
  /** @type {boolean} */
  self.wb = false;
  /** @type {boolean} */
  self.ia = false;
  /** @type {boolean} */
  self.Rb = false;
  /** @type {boolean} */
  self.Qc = false;
  /** @type {!Array} */
  self.Ac = [];
  /** @type {number} */
  self.Zb = 0;
  /** @type {boolean} */
  self.we = false;
  /** @type {!Array} */
  self.u = [];
  /** @type {!Array} */
  self.Ea = [];
  /** @type {!Array} */
  self.Da = [];
  /** @type {null} */
  self.K = null;
  /** @type {boolean} */
  self.ca = false;
  /** @type {boolean} */
  self.jb = false;
  /** @type {boolean} */
  self.za = false;
  /** @type {boolean} */
  self.ze = false;
  /** @type {number} */
  self.$d = 0;
  /** @type {number} */
  self.ae = 0;
  /** @type {number} */
  self.Tc = 0;
  /** @type {number} */
  self.Uc = 0;
  /** @type {number} */
  self.Ba = 0;
  /** @type {boolean} */
  self.be = false;
  /** @type {boolean} */
  self.V = false;
  /** @type {boolean} */
  self.kb = false;
  /** @type {!Array} */
  self.Ca = [];
  /** @type {number} */
  self.vd = 0;
  /**
   * @param {!Object} callback
   * @param {string} hash
   * @return {undefined}
   */
  self.ha = function (callback, hash) {
    /** @type {!Object} */
    this.$ = callback;
    this.Ba = hash || 0;
    this.ja = getUnderlineBackgroundPositionY();
    this.b = this.$.app.P;
    if (this.Rb && !this.b) {
      /** @type {boolean} */
      this.Rb = false;
    }
    this.f = $(this.$.f, "div", {
      className: "bcont noth usno"
    });
    if (this.$.app.Nb) {
      /** @type {boolean} */
      this.be = true;
    } else {
      if (window.PointerEvent) {
        /** @type {boolean} */
        this.V = true;
      } else {
        /** @type {boolean} */
        this.kb = true;
      }
    }
    link(this);
    this.Jb = $(this.f, "canvas", {
      className: "noth",
      style: {
        position: "absolute",
        left: 0,
        top: 0,
        width: "100%",
        height: "100%",
        zIndex: 1
      }
    });
    this.Fb = this.Jb.getContext("2d");
  };
  /**
   * @return {undefined}
   */
  self.qb = function () {
  };
  /**
   * @return {undefined}
   */
  self.oc = function () {
  };
  /**
   * @return {undefined}
   */
  self.Dd = function () {
  };
  /**
   * @return {undefined}
   */
  self.history = function () {
  };
  /**
   * @return {undefined}
   */
  self.jd = function () {
    if (0 != this.Da.length) {
      /** @type {number} */
      var a = this.Da.length - 1;
      for (; 0 <= a; a--) {
      }
      this.Ca.forEach(function (current, cX1) {
        if (current.Gb && !current.kd) {
          this.oc(cX1, current.Ob, current.Ae, current.Be);
        }
      }, this);
    }
  };
  /**
   * @return {undefined}
   */
  self.na = function () {
    var img_width = this.f.offsetWidth;
    var randomBasedOnCookie = this.f.offsetHeight;
    var c = getUnderlineBackgroundPositionY();
    if (img_width != this.hb || randomBasedOnCookie != this.gb || this.ja != c) {
      this.ja = c;
      this.hb = img_width;
      this.gb = randomBasedOnCookie;
      /** @type {number} */
      this.Jb.width = Math.floor(img_width * this.ja);
      /** @type {number} */
      this.Jb.height = Math.floor(randomBasedOnCookie * this.ja);
      this.Fb.fillRect(0, 0, 1 * this.ja, 1 * this.ja);
      this.jd();
      this.ic(true);
    }
  };
  /**
   * @return {undefined}
   */
  self.Vb = function () {
  };
  /**
   * @return {undefined}
   */
  self.Zc = function () {
  };
  /**
   * @param {string} val
   * @return {undefined}
   */
  self.rf = function (val) {
    var c = this;
    /** @type {!Arguments} */
    var names = arguments;
    var cover;
    /** @type {number} */
    var i = 0;
    /** @type {number} */
    var length = names.length;
    for (; i < length; i++) {
      this.Ac.push(cover = $("img", {
        src: window.k2img[names[i]]
      }));
      if (!cover.complete) {
        this.Zb++;
        /**
         * @return {undefined}
         */
        cover.onload = function () {
          c.Zb--;
          if (0 == c.Zb && c.we) {
            c.ic();
          }
        };
      }
    }
  };
  /**
   * @return {undefined}
   */
  self.ic = function () {
    if (loadPlugin(this) && 0 != this.Ea.length) {
      var i;
      /** @type {boolean} */
      var b = !!this.K && !!this.K.Z;
      /** @type {number} */
      i = this.Ea.length - 1;
      for (; 0 <= i; i--) {
        /** @type {null} */
        this.Ea[i].Z = null;
      }
      /** @type {number} */
      i = this.Da.length - 1;
      for (; 0 <= i; i--) {
        var item = this.Da[i];
        if (!(0 > item.y || item.y >= this.u.length || 0 > item.x || item.x >= this.u[item.y].length)) {
          var result = this.u[item.y][item.x];
          if (-1 != item.F) {
            /** @type {number} */
            item.x = item.y = -1;
          } else {
            result.Z = item;
          }
        }
      }
      /** @type {number} */
      i = this.Ea.length - 1;
      for (; 0 <= i; i--) {
        result = this.Ea[i];
        if (!(0 > result.x || 0 > result.y || null != result.Z)) {
          result.Z = f(this, result.x, result.y);
        }
      }
      if (b && !this.K.Z) {
        split("ERR " + window.k2url.game + " PI/NULL/1 s.x,y=" + this.K.x + "," + this.K.y + ", act:" + this.wb);
      }
      result = this.K && this.ca ? this.K.Z : null;
      /** @type {number} */
      i = this.Da.length - 1;
      for (; 0 <= i; i--) {
        item = this.Da[i];
        if (0 <= item.x && 0 <= item.y) {
          if (item != result) {
            extend(item.a, {
              display: "block"
            });
          }
        } else {
          extend(item.a, {
            display: "none"
          });
        }
      }
      this.Ca.forEach(function (current, cX1) {
        if (current.kd && current.Gb) {
          this.oc(cX1, current.Ob, current.Ae, current.Be);
        }
      }, this);
    }
  };
  /**
   * @param {string} b
   * @param {string} t
   * @return {undefined}
   */
  self.setActive = function (b, t) {
    if (this.wb != b) {
      this.ia = this.wb = b;
      if (!b && this.K) {
        animate(this, null);
        if ("undefined" != typeof t && t) {
          isArray(this);
        }
      }
      if (!b && this.za) {
        replace(this);
      }
      if (!b && this.ze) {
        /** @type {boolean} */
        this.ze = false;
      }
    }
  };
  /** @type {string} */
  var element = "+ ";
  /**
   * @return {undefined}
   */
  Controller.prototype.reset = function () {
    /** @type {string} */
    this.a.innerHTML = "";
    /** @type {string} */
    this.ab.value = "";
  };
  /**
   * @return {undefined}
   */
  Controller.prototype.Aa = function () {
    if (this.o.ld && !this.app.oa) {
      if (0 < this.f.clientHeight) {
        /** @type {number} */
        (document.scrollingElement || document.documentElement).scrollTop = document.documentElement.scrollHeight - document.documentElement.clientHeight;
      }
    } else {
      var t = this.a;
      /** @type {number} */
      var center = t.scrollHeight - t.clientHeight;
      if (0 < center) {
        /** @type {number} */
        t.scrollTop = center;
      }
    }
  };
  /**
   * @param {!Object} props
   * @param {string} root
   * @return {undefined}
   */
  Controller.prototype.append = function (props, root) {
    var bng1 = this;
    var commaPos;
    props = props.split("\n");
    var data = this.a;
    var a = this.o.ld && !this.app.oa ? document.body : this.a;
    /** @type {boolean} */
    var k = a.scrollTop + 2 >= a.scrollHeight - a.clientHeight;
    if (root) {
      props.reverse();
    }
    /** @type {number} */
    var i = 0;
    for (; i < props.length; i++) {
      var value = props[i];
      var v = root ? data.firstChild : null;
      if (0 == value.indexOf(element)) {
        var f = value.substring(element.length);
        if (2 < f.length && "[" == f[0] && "]" == f[f.length - 1]) {
          /** @type {number} */
          value = parseInt(f.substring(1, f.length - 1), 10);
          if (!isNaN(value)) {
            /**
             * @param {number} i
             * @return {?}
             */
            f = function (i) {
              /** @type {number} */
              i = parseInt(i, 10);
              return 10 > i ? "0" + i : i;
            };
            /** @type {!Date} */
            value = new Date(1E3 * value);
            /** @type {string} */
            f = value.toLocaleDateString() == (new Date).toLocaleDateString() ? f(value.getHours()) + ":" + f(value.getMinutes()) : value.getFullYear() + "-" + f(value.getMonth() + 1) + "-" + f(value.getDate());
          }
        }
        write(html(data, $("div", {
          className: "tind"
        }, ["+ " + f]), v));
      } else {
        if (1 < value.length && "[" == value[0] && "]" == value[value.length - 1]) {
          html(data, $("div", {
            className: "mtbq"
          }, $("button", {
            onclick: function (event) {
              this.parentNode.parentNode.removeChild(this.parentNode);
              if (bng1.o.De) {
                bng1.o.De();
              }
              event.stopPropagation();
              return false;
            }
          }, [value.substring(1, value.length - 1)])), v);
        } else {
          if (-1 != (commaPos = value.indexOf(":"))) {
            write(html(data, $("div", {
              className: "tind"
            }, [$("b", [value.substring(0, commaPos)]), value.substring(commaPos)]), v));
          } else {
            $(data, "div", [value]);
          }
        }
      }
    }
    if (this.o.ld && !this.app.oa) {
      if (!root && 0 < this.f.clientHeight && document.documentElement.scrollHeight > document.documentElement.clientHeight) {
        /** @type {number} */
        (document.scrollingElement || document.documentElement).scrollTop = document.documentElement.scrollHeight;
      }
    } else {
      if (k) {
        /** @type {number} */
        a.scrollTop = a.scrollHeight - a.clientHeight + 1;
      }
    }
  };
  /** @type {string} */
  var w = '<span class="emo">$&</span>';
  /** @type {number} */ // lk16 constants
  var cluezIndex = 99;
  /** @type {number} */
  var default_zIndex = 101;
  /** @type {number} */
  var n = 512;
  /** @type {number} */
  var za = 17;
  /** @type {number} */
  var Ba = 18;
  /** @type {number} */
  var Aa = 19;
  /** @type {number} */
  var ease_out = 20;
  /** @type {number} */
  var wink = 21;
  /** @type {number} */
  var gc = 22;
  /** @type {number} */
  var xa = 23;
  /** @type {number} */
  var x = 24;
  /** @type {number} */
  var X = 25;
  /** @type {number} */
  var v = 27;
  /** @type {number} */
  var Ub = 28;
  /** @type {number} */
  var Rb = 29;
  /** @type {number} */
  var $b = 30;
  /** @type {number} */
  var last = 31;
  /** @type {number} */
  var bc = 32;
  /** @type {number} */
  var hc = 33;
  /** @type {number} */
  var ease_in = 34;
  /** @type {number} */
  var CSI = 35;
  /** @type {number} */
  var workerId = 20;
  /** @type {number} */
  var reg = 22;
  /** @type {number} */
  var ya = 51;
  /** @type {number} */
  var focus = 52;
  /** @type {number} */
  var Yb = 53;
  /** @type {number} */
  var id = 52;
  /** @type {number} */
  var ic = 61;
  /** @type {number} */
  var action = 61;
  /** @type {number} */
  var ia = 63;
  /** @type {number} */
  var below_centered = 64;
  /** @type {number} */
  var valid = 70;
  /** @type {number} */
  var visited = 71;
  /** @type {number} */
  var active = 72;
  /** @type {number} */
  var read_write = 73;
  /** @type {number} */
  var eslint = 74;
  /** @type {number} */
  var u = 75;
  /** @type {number} */
  var Zb = 76;
  /** @type {number} */
  var cc = 77;
  /** @type {number} */
  var payload = 71;
  /** @type {number} */
  var overwrite = 72;
  /** @type {number} */
  var Kb = 74;
  /** @type {number} */
  var prepend = 76;
  /** @type {number} */
  var event = 77;
  /** @type {number} */
  var lb = 80;
  /** @type {number} */
  var SS = -2;
  /** @type {number} */
  var Z = 0;
  /** @type {number} */
  var No = 1;
  /** @type {number} */
  var cursel = 1;
  /** @type {number} */
  var readChecksum = 2;
  var ctx;
  /**
   * @return {undefined}
   */
  draw.prototype.reset = function () {
    /** @type {number} */
    this.g = 0;
    /** @type {number} */
    this.b = -1;
    /** @type {number} */
    var index = this.a.rows.length - 1;
    for (; 0 <= index; index--) {
      this.a.deleteRow(index);
    }
  };
  /**
   * @param {number} addedRenderer
   * @return {undefined}
   */
  draw.prototype.Aa = function (addedRenderer) {
    if (!(-1 == this.tab.Za && this.b < this.g - 1 - addedRenderer)) {
      testCircleCircle(this, this.g - 1);
      /** @type {number} */
      this.sa.scrollTop = this.sa.scrollHeight - this.sa.clientHeight;
    }
  };
  /** @type {number} */
  var ignoreChecksum = 0;
  /** @type {number} */
  var ed = 0;
  /**
   * @return {undefined}
   */
  test.prototype.onshow = function () {
    if (this.g) {
      this.g.Aa();
    }
    if (!this.W) {
      /** @type {boolean} */
      this.W = true;
      if ("function" === typeof window.k2adload) {
        setTimeout(function () {
          window.k2adload();
        }, 0);
      }
    }
  };
  /**
   * @return {undefined}
   */
  test.prototype.Db = function () {
    /** @type {boolean} */
    var l = "p" == y();
    if (this.L != l) {
      /** @type {boolean} */
      this.L = l;
      callback(this.f, "tbact", l);
    }
  };
  self = kernel.prototype;
  /** @type {number} */
  self.ea = 0;
  /** @type {number} */
  self.pc = 0;
  /** @type {boolean} */
  self.Eb = false;
  /** @type {boolean} */
  self.bd = false;
  /** @type {null} */
  self.name = null;
  /** @type {null} */
  self.Jc = null;
  /** @type {null} */
  self.Sa = null;
  /**
   * @param {number} index
   * @param {number} ch
   * @param {boolean} hash
   * @param {string} callback
   * @return {undefined}
   */
  self.ha = function (index, ch, hash, callback) {
    /** @type {boolean} */
    this.Eb = ch = !!ch;
    this.ea = index + index % 2;
    /** @type {number} */
    this.pc = 0;
    this.Sa = callback || null;
    /** @type {!Array} */
    this.name = Array(this.ea);
    /** @type {!Array} */
    this.Jc = Array(this.ea);
    /** @type {!Array} */
    this.a = Array(this.ea);
    /** @type {!Array} */
    this.focus = Array(this.ea);
    if (ch) {
      /** @type {!Array} */
      this.R = Array(this.ea);
    }
    if (hash) {
      /** @type {!Array} */
      this.time = Array(this.ea);
    }
    /** @type {number} */
    index = 0;
    for (; index < this.ea; index++) {
      /** @type {string} */
      this.name[index] = "";
      /** @type {number} */
      this.a[index] = 0;
      /** @type {string} */
      this.Jc[index] = "";
      if (this.time) {
        /** @type {string} */
        this.time[index] = "0:00";
      }
      if (this.R) {
        /** @type {string} */
        this.R[index] = "-";
      }
      /** @type {boolean} */
      this.focus[index] = true;
    }
  };
  /**
   * @return {undefined}
   */
  self.reset = function () {
    /** @type {number} */
    var index = 0;
    for (; index < this.ea; index++) {
      /** @type {number} */
      var i = index;
      /** @type {number} */
      this.a[i] = 1;
      /** @type {string} */
      this.name[i] = "";
      if (this.R) {
        /** @type {string} */
        this.R[index] = "-";
      }
      equal(this, index, 0);
      /** @type {string} */
      this.Jc[index] = "";
      /** @type {boolean} */
      this.focus[index] = true;
    }
    /** @type {number} */
    this.pc = 0;
  };
  self = EventListener.prototype;
  /** @type {number} */
  self.F = 0;
  /** @type {number} */
  self.U = 0;
  /** @type {number} */
  self.la = 0;
  /** @type {number} */
  self.Za = -1;
  /** @type {number} */
  self.ga = -1;
  /** @type {number} */
  self.Me = 0;
  /** @type {null} */
  self.vc = null;
  /** @type {boolean} */
  self.Je = false;
  /** @type {boolean} */
  self.Tb = false;
  /** @type {number} */
  self.Yc = 0;
  /** @type {null} */
  self.app = null;
  /**
   * @param {string} s
   * @return {?}
   */
  self.j = function (s) {
    return this.app.j(s);
  };
  /**
   * @param {!Array} item
   * @param {?} type
   * @return {undefined}
   */
  self.send = function (item, type) {
    this.app.send(item, type);
  };
  /**
   * @param {string} a
   * @return {undefined}
   */
  self.X = function (a) {
    if (a.length) {
      if (a.length > n) {
        a = a.substring(0, n) + "...";
      }
      this.send([81, this.F], [a]);
    }
  };
  Promise.prototype = new EventListener;
  self = Promise.prototype;
  /** @type {function(): undefined} */
  self.constructor = Promise;
  /** @type {boolean} */
  self.dc = false;
  /** @type {null} */
  self.f = null;
  /** @type {null} */
  self.u = null;
  /** @type {null} */
  self.ma = null;
  /** @type {null} */
  self.T = null;
  /** @type {null} */
  self.ta = null;
  /** @type {null} */
  self.history = null;
  /** @type {null} */
  self.cb = null;
  /** @type {null} */
  self.bb = null;
  /** @type {!Array} */
  self.Xa = [];
  /** @type {null} */
  self.Nc = null;
  /** @type {null} */
  self.Lc = null;
  /** @type {null} */
  self.Oc = null;
  /** @type {null} */
  self.Mc = null;
  /** @type {null} */
  self.Kc = null;
  /** @type {null} */
  self.ka = null;
  /** @type {number} */
  self.Ya = 0;
  /** @type {null} */
  self.sb = null;
  /** @type {boolean} */
  self.kc = true;
  /** @type {boolean} */
  self.Vd = false;
  /** @type {boolean} */
  self.Hc = false;
  /** @type {boolean} */
  self.ue = false;
  /**
   * @param {!Object} callback
   * @param {?} hash
   * @param {string} a
   * @return {undefined}
   */
  self.ha = function (callback, hash, a) {
    if (!this.f) {
      /** @type {!Object} */
      this.app = callback;
      /** @type {boolean} */
      this.xf = this.xd = this.yf = this.yd = true;
      /** @type {string} */
      this.a = a;
      this.a.pd = this.a.pd || false;
      this.a.md = this.a.md || null;
      this.a.Wb = this.a.Wb || null;
      this.a.od = this.a.od || false;
      this.a.ke = this.a.ke || false;
      this.a.me = this.a.me || false;
      this.a.le = this.a.le || false;
      this.a.pe = this.a.pe || 2;
      this.a.je = this.a.je || false;
      this.T = new kernel;
      this.T.ha(this.app.hc, this.a.pd, true, this.a.md);
      this.f = $(this.app.B, "div", {
        className: "gview",
        style: {
          display: "none"
        }
      });
      if (this.app.P) {
        extend(this.f, {
          minHeight: "460px"
        });
      }
      fail(this);
      done(this);
      this.u = new hash;
      this.u.ha(this);
      extend(this.f, {
        background: this.u.Oe
      });
      page(this);
      new Layer(this.u.f, {
        className: "ctcont cttac ctst0",
        style: {
          zIndex: 70
        }
      });
      this.bb = new SvgShape(this);
      setupEvents(this);
      callback = this.app.qa;
      onload(callback && this.app.kf ? "click" : callback && this.app.G ? "touchstart" : null);
    }
  };
  /**
   * @return {undefined}
   */
  self.ye = function () {
    callback(this.f, "tstatinbohide", !setItem(this.history));
  };
  /** @type {null} */
  self.ya = null;
  /**
   * @param {!Object} x
   * @param {number} c
   * @param {?} name
   * @param {?} b
   * @return {undefined}
   */
  self.Ud = function (x, c, name, b) {
    this.u.setActive(false, void 0, "_move");
    /** @type {!Array} */
    x = [92, this.F, x];
    if ("undefined" != typeof c) {
      x.push(c);
    }
    if ("undefined" != typeof name) {
      x.push(name);
    }
    x.push(Math.floor((Date.now() - this.bb.g) / 100));
    this.send(x, "undefined" != typeof b ? [b] : null);
  };
  /**
   * @param {!Object} scope
   * @return {undefined}
   */
  self.Ge = function (scope) {
    if (!this.a.bf) {
      var prev = this.app.Ed;
      if (prev && !(1 >= prev.length)) {
        var c = this;
        $(scope, "div", {
          className: "mbsp"
        }, [$("div", {
          className: "nowrel"
        }, this.j("tb_tr_add")), this.A = $("select", {
          onchange: function () {
            try {
              setTimeout(c, "tm", this.options[this.selectedIndex].text);
            } catch (d) {
            }
          }
        }, prev.map(function (mei) {
          return $("option", mei);
        }))]);
        this.Xa.push(this.A);
      }
    }
  };
  /**
   * @return {undefined}
   */
  self.Fe = function () {
  };
  /**
   * @return {undefined}
   */
  self.reset = function () {
    if (0 != this.F) {
      /** @type {number} */
      this.F = 0;
      cb(this.app, "tabopen", 0);
      /** @type {boolean} */
      this.kc = true;
      /** @type {boolean} */
      this.Hc = false;
      /** @type {null} */
      this.vc = null;
      var o = this.bb;
      if (o.B) {
        clearInterval(o.B);
        /** @type {number} */
        o.B = 0;
        /** @type {boolean} */
        o.H = false;
      }
      /** @type {number} */
      o.a = -1;
      /** @type {number} */
      o.A = 0;
      _(this, false);
      this.T.reset();
      slice(this);
      this.ta.reset();
      if (this.history) {
        this.history.reset();
      }
      this.cb.reset();
      if (this.ka) {
        this.ka.show(0);
      }
      if (this.b) {
        this.b.show(0);
      }
      subscribe(this, null);
      end(this, false, false);
      this.u.reset(true);
      isArray(this.u);
    }
  };
  /**
   * @param {number} fallback
   * @return {?}
   */
  self.mb = function (fallback) {
    return fallback;
  };
  /**
   * @param {!Object} el
   * @param {number} e
   * @return {?}
   */
  self.ie = function (el, e) {
    return e;
  };
  /**
   * @param {!Object} auth
   * @param {!Object} connection
   * @return {?}
   */
  self.Ee = function (auth, connection) {
    return connection;
  };
  /**
   * @param {string} config
   * @return {undefined}
   */
  self.Db = function (config) {
    if (config) {
      config = this.app;
      /** @type {boolean} */
      this.ue = !!config.C && (!config.sc || config.C != config.sc.f);
    }
  };
  /**
   * @return {undefined}
   */
  self.onshow = function () {
    /** @type {number} */
    var port = parseInt(proceed(), 10);
    if (port != this.F) {
      server(this.app, port, true);
    }
    if (this.ta) {
      this.ta.Aa();
    }
  };
  /** @type {number} */
  var base = 3;
  /** @type {number} */
  var position = 1;
  /** @type {number} */
  var h = 1;
  /** @type {number} */
  var value = 2;
  /** @type {number} */
  var F = 4;
  /** @type {number} */
  var file = 16;
  /** @type {number} */
  var DIRECTION_HORIZONTAL = 1;
  /** @type {number} */
  var arg = 2;
  /** @type {number} */
  var DyMilli = 16;
  /**
   * @return {undefined}
   */
  Color.prototype.reset = function () {
    /** @type {number} */
    this.g.length = 0;
    Object.keys(this.b).forEach(function (pos) {
      var data = this.b[pos];
      forEach(data.Ua, this);
      if (data.D) {
        data = data.D;
        data.parentNode.removeChild(data);
      }
      delete this.b[pos];
    }, this);
    ok(this, 0);
  };
  /** @type {null} */
  var audioElement = null;
  /** @type {null} */
  var context = null;
  /** @type {null} */
  var target = null;
  /** @type {boolean} */
  var corsAvailable = false;
  /** @type {boolean} */
  var we = false;
  /** @type {number} */
  var runlist = -1;
  /** @type {number} */
  var encoding = 1;
  /** @type {number} */
  var or = 2;
  /**
   * @param {!Array} result
   * @param {?} data
   * @return {undefined}
   */
  ready.prototype.send = function (result, data) {
    if (this.A || this.a) {
      if (result = encode(result, data), this.A) {
        try {
          this.A.send(result);
        } catch (c) {
        }
      } else {
        if (this.a) {
          addOverlay(this, result);
        }
      }
    }
  };
  if (!("classList" in document.createElement("_"))) {
    /**
     * @param {!Object} a
     * @param {string} name
     * @param {boolean} content
     * @return {undefined}
     */
    callback = function (a, name, content) {
      if (2 == arguments.length || content) {
        if (!transform(a, name)) {
          a.className += (0 < a.className.length ? " " : "") + name;
        }
      } else {
        if (transform(a, name)) {
          a.className = a.className.replace(new RegExp("(\\s|^)" + name + "(\\s|$)"), " ");
        }
      }
    };
    /**
     * @param {!Object} t
     * @param {string} val
     * @return {?}
     */
    transform = function (t, val) {
      return !!t.className.match(new RegExp("(\\s|^)" + val + "(\\s|$)"));
    };
  }
  /**
   * @param {string} i
   * @return {undefined}
   */
  Layer.prototype.show = function (i) {
    if (this.a != i) {
      var children = this.f.children;
      /** @type {number} */
      var j = 0;
      var il = children.length;
      for (; j < il; j++) {
        if ("undefined" !== typeof children[j]["data-tab"]) {
          extend(children[j], {
            display: children[j]["data-tab"] == i ? "block" : "none"
          });
        }
      }
      /** @type {string} */
      this.a = i;
    }
  };
  /**
   * @param {?} data
   * @param {?} result
   * @param {?} element
   * @return {?}
   */
  Layer.prototype.add = function (data, result, element) {
    result = element ? $(this.f, "div", result, element) : $(this.f, "div", result);
    result["data-tab"] = data;
    return result;
  };
  /**
   * @return {undefined}
   */
  SvgShape.prototype.start = function () {
    var ast = this;
    /** @type {boolean} */
    this.H = true;
    /** @type {number} */
    this.B = setInterval(function () {
      return verify(ast);
    }, Math.floor(1E3 / this.S));
  };
  /** @type {number} */
  window.k2ver = 233;
  /**
   * @return {?}
   */
  start.prototype.Le = function () {
    return this.b ? this.B : null;
  };
  /**
   * @return {undefined}
   */
  start.prototype.onshow = function () {
    if (this.b) {
      this.b.u.Aa();
    }
    normalize(this);
    var data = this.b;
    _get(this);
    if (data && this.app.P) {
      data.u.ab.focus();
    }
  };
  /**
   * @return {undefined}
   */
  start.prototype.Db = function () {
    var x = y();
    if (x && !/^[a-z0-9~]+$/i.test(x)) {
      /** @type {null} */
      x = null;
    }
    if (!this.app.oa) {
      request(this, x);
    }
  };
  /**
   * @return {undefined}
   */
  b.prototype.onshow = function () {
    this.b.checked = this.app.cc;
    this.a.checked = this.app.bc;
    this.g.checked = this.app.ec;
  };
  /**
   * @param {string} options
   * @return {undefined}
   */
  main.prototype.Db = function (options) {
    var w = y();
    /** @type {number} */
    var b = 1;
    /** @type {number} */
    var value = 1;
    var f;
    if (w && (f = w.split("/")) && "f" == f[0]) {
      /** @type {number} */
      b = 2;
      /** @type {number} */
      w = parseInt(f[1], 10);
      if (!isNaN(w)) {
        /** @type {number} */
        value = w;
      }
    }
    if (options || b != this.b || value != this.B) {
      if (0 != this.b) {
        reposition(this, 0);
      }
      /** @type {number} */
      this.b = b;
      /** @type {number} */
      this.B = value;
      destroy(this);
    }
  };
  /**
   * @return {undefined}
   */
  main.prototype.lb = function () {
  };
  /**
   * @param {string} res
   * @param {boolean} err
   * @param {?} data
   * @return {undefined}
   */
  main.prototype.handle = function (res, err, data) {
    if (res == this.A) {
      /** @type {null} */
      this.A = null;
      if (err) {
        this.fa.innerHTML = data.html || "";
        build(this, data);
        reposition(this, 1);
      }
    }
  };
  /**
   * @return {undefined}
   */
  r.prototype.onshow = function () {
    var a = this;
    toggleVisibility(this, 0);
    var b = this.b = new View(this.o.url, redraw(this.o.dd ? {
      jsget: 1,
      ksession: wrapped(this.app)
    } : {
        jsget: 1
      }));
    /**
     * @param {boolean} error
     * @param {?} options
     * @return {?}
     */
    b.handle = function (error, options) {
      return a.handle(b, error, options);
    };
  };
  /**
   * @param {string} e
   * @param {?} data
   * @return {undefined}
   */
  r.prototype.lb = function (e, data) {
    var eventManager = this;
    if ("fsub" == e) {
      toggleVisibility(this, 0);
      var b = this.b = new View(this.o.url, new FormData(data));
      /**
       * @param {boolean} data
       * @param {?} type
       * @return {?}
       */
      b.handle = function (data, type) {
        return eventManager.handle(b, data, type);
      };
    }
    if (this.o.lb) {
      this.o.lb(e, data);
    }
  };
  /**
   * @param {string} value
   * @param {boolean} type
   * @param {!Object} data
   * @return {undefined}
   */
  r.prototype.handle = function (value, type, data) {
    if (value == this.b && (this.b = null, type)) {
      if (data.ksession) {
        value = this.app;
        try {
          window.localStorage.setItem("ksession", data.ksession);
        } catch (d) {
        }
        if (value.Ub) {
          /** @type {string} */
          document.cookie = "kguest=1;path=/";
        }
        find(value, true);
      } else {
        if (data.done) {
          data = this.app;
          if (data.C && data.C == this.f) {
            data = this.app;
            debug(data, data.N);
          }
        } else {
          this.g.innerHTML = data.html || "";
          if (data.script) {
            $(this.g, "script", {
              type: "text/javascript",
              text: data.script
            });
          }
          toggleVisibility(this, 1);
        }
      }
    }
  };
  window.k2snd = {};
  /** @type {string} */
  window.k2snd["a.mp3"] = "data:audio/mpeg;base64,SUQzAwAAAAADWVRBTEIAAAABAAAAVENPTgAAAAEAAABUSVQyAAAAAQAAAFRQRTEAAAABAAAAVFJDSwAAAAEAAABUWUVSAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//uQxAAAAAAAAAAAAAAAAAAAAAAAWGluZwAAAA8AAAADAAAGHQCioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqLd3d3d3d3d3d3d3d3d3d3d3d3d3d3d3d3d3d3d3d3d3d3///////////////////////////////////////////8AAABQTEFNRTMuOTlyBLkAAAAALjMAADUgJAMlQQAB4AAABh3H83fyAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//vAxAAADKhVSnWGACMLKCr/P0DRAACd3SHRXRXQDoB0U2nswLkGlJtic6nep1ibTgoam7XwOCQJAkCQJBgYLF69evMz9evfpRYsWLAQBMHwfB8PqBAEDlYPg/lHZcHwfygY1g+H8oGNYPn8oGPz/KcoH+BOg/wJ5/oAAAAEAJwJ4Np00doeAoHATSRSLpgYC5XiIQEYwB0AqMBUAXTEJAMUwexRsMC9AiTDgQz8wGkASAIAIYPCECGASAOpg9gCGUAO4ciGWQMiwAYX4ZFHKGaAxgcAgSBgAmLoQsQ0c0ZUMVB+wqQhUl7pCzhKQfEOaLlIaSwh4QByerGVBvE5ATIvEWMS6YmLJFIMN8lyZMi8TRiXS7+WT3zpkXiaMS6XTI2Lxr+Yf+tFFSSS0UVJfW3//MgUkFRSQVF7//1hAAAAB43IIIAWipavJAMy4KAOmCUBiYfQOphMALmGIMaZA0pJlnERA4V8wPANzAsAnDgJU5gSAABQR4mTurQrdRkHYTtnXZN61N/UoaRaf//1k0SCP//8xHV///8pf//82/8r/66YAAAAQDiKAAAMABHADkOZCABAYgAOIACzCcAhMDIFowqQgzDMYdNCYD4w9AXFVhoA0MALkzjiXqL//oCGDMmv//skw0h5Q//+qYBus/4Jf+JP/R/6VYWQCyC3mkmFdd8oAAAAAAAAAxi6OpdMsCSIgFDAhMPxUNTplMKAzMUzIMWS0MaTeMWBtFAKMlwAIgAMyxCMWirt6OGTUyQQmIC/+i3gkHZ4y2Iu7f//fyB5RLObx1/2oj9HSS2GqhIV4cKTqkxBTUUzLjk5LjWqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqr/+2DEzwAMuSEr/eaAIVYTpX680ASqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqr/+zDE6QANbG8v+d0AAAAAP8OAAASqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpUQUcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/w==";
  /** @type {string} */
  window.k2style = '.h1s{height:44px;box-sizing:border-box}.aleg .h1s{height:40px;box-sizing:border-box}.lh1s{line-height:44px}.aleg .lh1s{line-height:40px}.mb1s{margin-bottom:44px}.aleg .mb1s{margin-bottom:40px}.hdhei{height:46px}.hdbwd{width:54px}.k2base{background:#fff;color:#000;position:relative}.k2base{height:auto;min-height:100%}.k2base.asizing{margin:0 auto}.k2base{-webkit-tap-highlight-color:rgba(0,0,0,0)}.k2base.devios{cursor:pointer}.k2base{font:16px/1.4 arial,sans-serif}.k2base.aleg{font:14px/1.4 arial,sans-serif}.win .k2base.aleg{font:13px/1.5 verdana,sans-serif}.k2base{-webkit-text-size-adjust:100%}.k2base{-ms-touch-action:pan-y}.anav{display:none;position:fixed;top:0;left:0;right:0;height:46px;line-height:46px;text-align:center;background:#2258bb;overflow:hidden}.asizing .anav{position:absolute}.donav .anav{display:block}.aleg .anav{position:absolute;background:#9ce}.acon{position:absolute;top:0;bottom:0;left:0;right:0}.donav:not(.aleg) .acon{position:static;padding-top:46px}.aleg.donav .acon{top:46px;overflow-y:auto}.vm2 .aleg.donav.k2base{border-style:solid;border-color:#9ce;border-width:0 8px 0 8px}.vm2 .aleg.donav .acon{border-bottom:solid 44px #9ce}.stview{padding:0}.stvxpad{padding-left:40px;padding-right:40px}.vm0 .stvxpad,.vm1 .stvxpad{padding-left:15px;padding-right:15px}.vm2 .stview,.vm2 .stvxpad{padding:37px 40px}.stview,.stvxpad{border-bottom:solid 1px transparent}.vm2 .btifp{border-top:solid 1px #e4e4e4}.dclpd,.dcpd{padding-left:40px}.dcrpd,.dcpd{padding-right:40px}.vm0 .dclpd,.vm0 .dcpd{padding-left:15px}.vm0 .dcrpd,.vm0 .dcpd{padding-right:15px}.vnarrow{max-width:520px;margin:0 auto}.vm0 .nifvm0{display:none}.ib{display:inline-block}.fb{font-weight:bold}.drops{width:3em}.nowrel{white-space:nowrap;overflow:hidden;text-overflow:ellipsis}.wsnw{white-space:nowrap}hr{border:0;height:1px;background-color:#e4e4e4}.snum{font-size:14px}.aleg .snum{font-size:13px}.win .snum{font-size:12px}.f12{font-size:12px}.win .f12{font-size:11px}.fl{font-size:18px}.emo{font-size:125%;line-height:.9;vertical-align:-0.22em}.win .emo{vertical-align:-0.1em}a:link,a:visited,.lc{color:#25b;text-decoration:none}input:not([type]),input[type=text],input[type=password],textarea{font:inherit;line-height:1.4;color:inherit;margin:0;border:solid 1px rgba(0,0,0,0.3);border-radius:0;padding:6px 2px;outline:none;-webkit-appearance:none}.aleg input:not([type]),.aleg input[type=text],.aleg input[type=password],.aleg textarea{padding:4px 2px;line-height:1.4}button,input[type="submit"],a.lbut{font:inherit;line-height:1.4;background:none;color:inherit;margin:0;border:solid 1px rgba(0,0,0,0.4);border-radius:4px;padding:6px 12px;outline:none;display:inline-block}.devtch button:active{opacity:.5}button:not(.butsys)[disabled],input[type="submit"][disabled]{opacity:.6}.aleg button,.aleg input[type="submit"],.aleg a.lbut{padding:4px 10px}.aleg button:not([disabled]){cursor:pointer}.butsys{background:#f8f8f8;border-color:rgba(0,0,0,0.3)}.butsys[disabled]{color:#888}.butwb{background:#fff;color:#222;border-color:transparent}.butlh,.aleg .butlh{padding:1px 10px}.butbl{background:rgba(32,96,180,0.2);color:rgba(32,96,180,0.8);border-color:transparent}.ddcont{display:none}.ddbut:focus+.ddcont,.devios .ddbut:hover+.ddcont{display:block}.ddcont.ddopen{display:block}.ddcont:active{display:block}a.lbut{color:inherit;text-align:center}select{font:inherit;margin:0;height:1.9em;outline:none}input[type=file]{font:inherit;box-sizing:border-box;max-width:100%}input[type=checkbox],input[type=radio]{margin-left:0;vertical-align:middle}button::-moz-focus-inner{padding:0;border:none}.selcwr{position:relative;display:inline-block;line-height:normal}.selcbt,.aleg .selcbt{text-align:left;overflow:hidden;text-overflow:ellipsis;padding-right:20px;white-space:nowrap}.selcbt::before{content:"";width:0;height:0;display:inline-block;position:absolute;right:8px;top:44%;border-top:solid 5px #555;border-left:solid 4px transparent;border-right:solid 4px transparent}.selcsl{opacity:0;position:absolute;box-sizing:border-box;left:0;top:0;width:100%;height:100%}.noth{-webkit-tap-highlight-color:rgba(0,0,0,0)}.tama{touch-action:manipulation}.bs{box-shadow:0 0 1px 1px rgba(0,0,0,0.15);-webkit-box-shadow:0 0 1px 1px rgba(0,0,0,0.15)}.usno{-webkit-user-select:none;-moz-user-select:none;-ms-user-select:none}.bsbb,.k2base{box-sizing:border-box;-webkit-box-sizing:border-box}.mlh{margin-left:.5em}.mlo{margin-left:1em}.mro{margin-right:1em}.mbsp{margin-bottom:.5em}.mbh{margin-bottom:.5em}.mtbq{margin:.25em 0}.mtoh{margin-top:1.5em}.minm{min-width:50px}.minw{min-width:75px}.minwd{min-width:140px;box-sizing:border-box}.tind{text-indent:-0.85em;margin-left:.85em}.win .aleg .tind{text-indent:-1.1em;margin-left:1.1em}.tac{text-align:center}.ttup{text-transform:uppercase}.dsp1,.aleg .dsp1{padding-left:10px;padding-right:10px}.fw{max-width:520px}.fw label{display:block;margin:4px 0}.fw input:not([type]),.fw input[type=text],.fw input[type=password]{box-sizing:border-box;width:100%;max-width:360px}.la{display:inline-block;-webkit-touch-callout:none;text-align:center}.spbb{position:relative;display:inline-block;width:20px;height:11px;border-radius:2px}.spbb:after{content:"";position:absolute;bottom:-5px;left:6px;border-width:5px 5px 0 0;border-style:solid}.aleg .spbb{position:relative;display:inline-block;width:18px;height:10px;border-radius:2px}.aleg .spbb:after{content:"";position:absolute;bottom:-4px;left:5px;border-width:4px 4px 0 0;border-style:solid}.spbb{background:rgba(32,96,180,0.4)}.spbb:after{border-color:rgba(32,96,180,0.4) transparent}.uicon{display:inline-block;position:relative;width:12px;height:3px;border-radius:2px 2px 0 0}.uicon:after{content:"";position:absolute;left:3px;width:6px;height:6px;margin-top:-7px;border-radius:6px}.uicon,.uicon:after{background:#aaa}button .cmenu{display:inline-block;width:5px;height:4px;border-top:4px solid #fff;border-bottom:12px double #fff;margin-top:4px}button.active .cmenu{border-color:#888}.sta{display:inline-block;width:11px;height:11px;border-radius:5px;background:#ddd;margin-right:8px}.aleg .sta{width:10px;height:10px}.st1 .sta{background:#2c2}.darr{display:inline-block;width:0;height:0;border-left:solid 5px transparent;border-right:solid 5px transparent;border-top:solid 6px #e8e8e8}.loader{display:inline-block;width:12px;height:6px;background-image:url(data:image/gif;base64,R0lGODlhDAACAIABAAAAAP///yH/C05FVFNDQVBFMi4wAwEAAAAh+QQBMgABACwAAAAADAACAAACBIyPqVcAIfkEATIAAQAsAAAAAAwAAgAAAgYEgqmmHQUAIfkEATIAAQAsAAAAAAwAAgAAAgaMDXB7yQUAIfkEATIAAQAsAAAAAAwAAgAAAgaMjwDIoQUAOw==)}#pread{margin:0;padding:10px 0 6px 40px;border-bottom:solid 1px #e4e4e4}.vm0 #pread,.vm1 #pread{padding-left:20px}.szpan{display:none;position:absolute;top:0;left:50%;margin-top:-12px;line-height:24px;width:160px;margin-left:-80px;text-align:center}.dosize .szpan{display:block}.ifdosize{display:none}.dosize .ifdosize{display:block}.szpan .bmax{height:13px;line-height:13px;width:32px;border:none;border-radius:0;padding:0 0 1px 0;background:#e4e4e4;color:#fff;margin:0 1px;cursor:pointer}table.br{-webkit-tap-highlight-color:transparent}table.br td{padding:4px 8px;border-bottom:solid 1px #eee;overflow:hidden}.astat{height:100%;text-align:center}.aleg .astat{background:#def}.astat>table{width:100%;height:100%}.aleg .astat>table{height:auto}.astat>table td{vertical-align:middle;padding:10px 0}.dtline{border-top:solid 1px rgba(0,0,0,0.2);margin:1em 0}.taplain{border:none;margin:0;outline:0}.ovysct{overflow-y:scroll;-webkit-overflow-scrolling:touch}.scln{margin:.25em 0}.scp1{min-width:6em}.bdtab{table-layout:fixed;border-collapse:collapse;background:#fff;text-align:center;width:280px;margin:0 auto;border-radius:3px;box-shadow:0 0 1px 1px rgba(0,0,0,0.1)}.bdtab{line-height:42px}.aleg .bdtab{line-height:40px}.bdmpart .bdp2{display:none}.bdmp2on .bdp1{display:none}.bdmp2on .bdp2{display:table-row}.bdtab td{border:solid 1px #e8e8e8;padding:0}.bdtab tr:first-child td,.bdtab tr.trlfst td{border-top:none}.bdtab tr:last-child td{border-bottom:none}.bdtab td:first-child{border-left:none}.bdtab td:last-child{border-right:none}.bdtab td button{width:100%;line-height:inherit;padding:0;border:none;border-radius:0;background:none;color:#222;outline:none}.bdtab td button:active{background:rgba(0,0,0,0.1)}.bdtab td button[disabled]{opacity:.5}.bdtab td button[disabled]:active{background:none}.bdtab tr.invm0{display:none}.vm0 .bdtab tr.invm0{display:table-row}.bdtab tr.invm0{display:none}.vm0 .bdtab tr.invm0{display:table-row}.vm1 .bdtab tr.trlfstnotvm0 td,.vm2 .bdtab tr.trlfstnotvm0 td{border-top:none}.clst .awrap{display:block;border-bottom:solid 1px #e4e4e4;line-height:44px;white-space:nowrap}.devtch .clst .awrap:active{background:rgba(0,0,0,0.15)}.aleg .clst .awrap{line-height:34px}.clst .awrap button{margin-left:.5em}.clst .maxw{max-width:420px}.clst .uname{display:inline-block;width:47%;font-weight:bold}.clst .infbl{display:inline-block}.clst .chtbl{display:inline-block;float:right;padding:0 16px;margin-right:-15px}.aleg .clst .uname{font-weight:normal}.aleg .clst .st1 .uname{font-weight:bold}.caddbox{margin:2.5em 0}.caddbox .aid{min-width:200px}.vm0 .clst .uname{width:auto}.vm0 .clst .inftx{display:none}.hvok .clst .awrap:hover{background:#eaecef;cursor:pointer}.ctcont{position:absolute;margin:0 auto;left:0;right:0}.ctcont button{padding:4px 10px}.cttal{text-align:left}.cttac{text-align:center}.ctcont{line-height:40px}.aleg .ctcont{line-height:36px}.ctst0{background:#fff;color:#000;width:300px}.ctst1{background:#333;color:#eee;width:300px}.ctst1 button{border-color:rgba(255,255,255,0.8);border-radius:2px;margin-right:.2em}.ctst0 .ctmsg+div,.ctst1 .ctmsg+div{margin-top:-6px;margin-bottom:6px}.ctcontcard{position:absolute;margin:0 auto;left:0;right:0;text-align:center}.ctcontcard button{border-color:#fff;color:#fff;border-radius:2px;padding:3px 10px}.ctcontlim4{white-space:nowrap}.ctcontlim4 button{min-width:45px;max-width:95px;white-space:nowrap;overflow:hidden}.ctnomsg .ctmsg{display:none}.ctstgs{background:#fff0b0}.ctstgs .ctpan{margin:0 1em}.ctstcol,.aleg .ctstcol{min-width:1.7em;margin-right:3px;padding:0;border-color:transparent}.vm2 .aleg .tblobby{height:100%}.tbvusers{display:none}.tbact .tbvtabs{display:none}.tbact .tbvusers{display:block}.vm2 .aleg .tbvtabs{display:block;float:left;width:71%}.vm2 .aleg .tbvusers{display:block;float:right;width:29%}.vm2 .aleg .tbvtabs{height:100%;overflow-y:scroll}.vm2 .aleg .tbvusers{height:100%;overflow-y:scroll}.newtab2{display:block;border-bottom:solid 1px #e4e4e4;padding-top:12px;padding-bottom:12px;white-space:nowrap}.tumode .newtab1 .vsel{width:5em}.vm2 .aleg .newtab2{display:none}.min85{min-width:85px}.tldeco{display:none}.vm2 .aleg .tldeco{display:block;height:16px;border-bottom:solid 1px #e4e4e4}.vm2 .aleg #pread~.tldeco{display:none}.alrt{padding-top:10px;padding-bottom:10px;border-bottom:solid 1px #e4e4e4;background:#fff0b0;color:rgba(0,0,0,0.9)}.tuinfo{display:none}.tumode .tuinfo{display:inline-block}.chpan{display:none;padding-top:10px;padding-bottom:10px;border-bottom:solid 1px #e4e4e4}.chmode .chpan{display:block}.chtop{color:#808080}.chsub{display:none;margin-top:10px}.chopen .chsub{display:block}.chgrlist{display:none;margin-right:15px}.chgrp.chopen .chgrlist{display:inline-block}.tabany{display:table}.tabmaxmed{display:table}.minmed{display:none}@media (min-width:501px){.tabmaxmed{display:none}.minmed{display:block}}@media (max-width:340px){.fs15xs{font-size:15px}}.imvfrm{display:none}.imact .imvlst{display:none}.imact .imvfrm{display:block}.vm2 .imvfrm{display:block;float:left;width:60%}.vm2 .imvlst{display:block;float:right;width:35%;margin-left:5%}.iml .awrap{display:block;line-height:44px;border-bottom:solid 1px #e4e4e4;font-weight:bold;cursor:pointer}.devtch .iml .awrap:active{background:rgba(0,0,0,0.15)}.aleg .iml .awrap{line-height:34px;font-weight:normal}.iml .slctd{background:#e8e8e8}.aleg .iml .st1{font-weight:bold}.iml .clbt{float:right;font-weight:normal;color:#ccc;padding:0 15px;margin-right:-15px}.iml .unrd{display:inline-block;margin-left:8px;background:#e22;color:#fff;line-height:1.2;font-weight:bold;border-radius:14px;padding:0 12px 1px 12px}.imtx{padding:4px 15px}.vm2 .imtx{border:solid 1px rgba(0,0,0,0.3);height:150px;overflow-y:scroll;-webkit-overflow-scrolling:touch}.imin{background:#f8f8f8;border-top:solid 1px #ddd;padding:8px 15px}.vm2 .imin{background:transparent;border-top:none;padding:8px 0}.imo1{display:table-cell;width:1%;padding-left:1em}.imo2{display:none}.vm2 .imo1{display:none}.vm2 .imo2{display:block;margin-top:.5em}.navttl{color:#fff;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;margin:0 80px}.navcont{width:100%;text-align:left;white-space:nowrap}.newtab1{display:inline-block;min-width:340px}.newtab1 .vsel{width:160px}.nav button{background:rgba(0,20,60,0.2);color:#fff;border-color:transparent;border-radius:0;margin-right:1px;min-width:5em}.nav button.bmain{min-width:6em}.nav button.btab{min-width:6em;margin-left:1.5em;background:#fc4;color:#fff;display:none}.nav button.alert{background:#e21;color:#fff}.nav button.active{background:#fff;color:#4c7a9c}.nav button .spbb{background:rgba(255,255,255,0.9)}.nav button .spbb:after{border-color:rgba(255,255,255,0.9) transparent}.nav button.alert .spbb{background:#fff}.nav button.alert .spbb:after{border-color:#fff transparent}.nav button.active .spbb{background:#4c7a9c}.nav button.active .spbb:after{border-color:#4c7a9c transparent}.navtabopen .nav button.btab{display:inline-block}.nav0{position:fixed;left:0;top:0;display:none}.asizing .nav0{position:absolute}.doddmenu .nav0{display:block}.nav0>.mbut{border:none;border-radius:0;padding:0;margin:0;color:#fff;border-bottom:solid 1px transparent;position:relative}.nav0>.mbut .micon{display:inline-block;width:20px;height:3px;border-top:3px solid #fff;border-bottom:9px double #fff;margin-top:3px}.thcol1 .mbut .micon{border-color:#999}.thcol2 .mbut .micon{border-color:#378a4a}.nav0>.mbut.alert::after{content:"";width:12px;height:12px;background:#f22;position:absolute;left:60%;top:33%;border-radius:6px}.nav0open{background:#fff;bottom:0;border-right:solid 1px #ddd;overflow-y:auto}.nav0open>.mbut .micon{border-color:#777}.nav0 .mcont{display:none}.nav0open .mcont{display:block;min-width:250px;margin-top:-1px}.nav0 .mlst{border-top:solid 1px #eee}.nav0 .mlst button{font:inherit;display:block;width:100%;text-align:left;color:#000;border:none;border-radius:0;border-bottom:solid 1px #eee;padding:0 17px;line-height:44px}.nav0 .mlst button:active{background:rgba(0,0,0,0.15)}.nav0 .mlst button:xlast-child{border-bottom:solid 1px #ddd}.nav0 .mlst button.btab{display:none}.nav0 .mlst button.alert{color:#e21;font-weight:bold}.nav0 .mlst button.active{font-weight:bold}.nav0 .mlst button .spbb{background:#aaa}.nav0 .mlst button .spbb:after{border-color:#aaa transparent}.nav0 .mlst button.alert .spbb{background:#e21}.nav0 .mlst button.alert .spbb:after{border-color:#e21 transparent}.navtabopen .nav0 .mlst button.btab{display:block}.nav0 .msub{padding:.5em 15px}button.ubut{border-color:transparent;background:#f8f8f8}.tbact button.ubut{background:#e4e4e4}.r0,.r1,.r2,.r3,.r4,.r5,.rnone{display:inline-block;vertical-align:baseline;border-radius:2px;margin-right:6px}.r0,.r1,.r2,.r3,.r4,.r5,.rnone{width:10px;height:8px;border:solid 1px rgba(0,0,0,0.01)}.aleg .r0,.aleg .r1,.aleg .r2,.aleg .r3,.aleg .r4,.aleg .r5,.aleg .rnone{width:8px;height:7px}.rnone,.aleg .rnone{border-color:transparent}.r0{background:#d0d0d0}.r1{background:#00d0f8}.r2{background:#00bf00}.r3{background:#eaea10}.r4{background:#ff9910}.r5{background:#ee4242}.bcont{position:absolute;left:0;top:0;right:0;bottom:0}.bcont{-ms-touch-action:none}.bhead{display:none;height:46px}.gvhead .bhead{display:table}.gvhead .bcont{top:46px}.gvnohead{display:none}.gview:not(.gvhead) .gvnohead{display:block}.wcol,.wcolhd{color:#fff}.wcol button{border-color:#fff}.wcolhd button{padding:3px 12px;border-color:#fff;margin-right:4px}.wcol button.alert,.wcolhd button.alert{background:#e11}.gview{position:relative;width:100%;height:100%;overflow:hidden}.iosfix .gview{position:fixed}@media screen and (orientation:landscape) and (min-width:598px){.devmob.h100vh .gview{height:100vh;min-height:314px}}@media screen and (orientation:portrait) and (device-height:480px){.gview{min-height:416px}}.thead{display:none;position:absolute;top:0;left:0;right:0}.thead{color:#fff;border-bottom:solid 1px #fff}.thcol1 .thead{color:#999;border-bottom-color:#999}.thcol2 .thead{color:#378a4a;border-bottom-color:#378a4a}.gvhead .thead{display:block}.thnavcont{display:none;color:#fff}.thcol1 .thnavcont{color:#999}.thcol2 .thnavcont{color:#378a4a}.sbdrop .thnavcont{display:block}.xbut{background:transparent;border:none;color:inherit;position:absolute;top:0}.cmenubut{background:transparent;border:none;color:inherit;position:absolute;top:0}.thcol1 .cmenu{border-color:#999}.thcol2 .cmenu{border-color:#378a4a}.sbdropvis .cmenubut{background:#fff}.sbdropvis .cmenu{border-color:#888}.tsb{position:absolute;top:0;right:0;bottom:0;overflow-y:auto;display:none}.sbfixed .tsb{width:310px}.sbdrop .tsb{width:320px;top:46px}.sbdrop .tsb{display:block;visibility:hidden;z-index:98}.sbdropvis .tsb{visibility:inherit}.sbfixed .tsb{display:block}.sbfixed .bcont{margin-right:310px}.tsbinner{padding:5px 10px 10px}.sbfixed .tsbinner{min-height:100%}.sbclrd{background:#9ce;color:#123}.tsinbo{display:none}.sbfixed .tsinbo{display:table;visibility:hidden}.tstatact .tsinbo{visibility:inherit}.tstatinbohide .tsinbo{visibility:hidden}.tsinsb{display:none}.sbdrop .tsinsb{display:block;margin:4px 0 3px}.tstatstart .tsinsb .tstatlabl{display:none}.tstatstrl{display:none}.tstatstart .tstatstrl{display:block}.ttlcont{white-space:nowrap}.sbdrop .ttlcont{font-weight:bold;text-align:center;line-height:44px}.sbfixed .ttlcont{font-weight:bold;border-bottom:solid 1px rgba(0,0,0,0.25);line-height:46px}.ttlnav{display:none}.ttlnav button{min-width:42px;margin-left:2px}.sbfixed .ttlnav{display:block;float:right;font-weight:normal;text-align:right}.ttlnav button.alert{background:#e21;color:#fff}.sbfixed .tplcont{margin-top:15px}.sbfixed .trqcont{margin-top:20px;border-top:solid 1px rgba(0,0,0,0.25)}.trqcont .trqans{background:#234;color:#fff}.tctres{display:none}.sbfixed .tctres{display:block;margin-top:15px}.trk{display:none}.sbfixed .trk{display:block;margin-top:15px;border-top:solid 1px rgba(0,0,0,0.25)}.sbfixed .tcrdcont{margin-top:15px}.sbfixed .tcrdspecgs{margin-top:4px}.sbfixed .tcrdpan{height:160px}.sbdrop .tcrdpan{height:160px}.tsbchat{display:none}.sbfixed .tsbchat{display:block}.tinbcht{display:none}.sbdrop .tinbcht{display:block}.tinbopt{display:none}.sbfixed .tinbopt{display:block}.vm2 .ctgstools{float:right;width:48%}.vm2 .ctgsword{float:right;width:50%;max-width:300px}.vm0 .ctgstools,.vm1 .ctgstools{border-bottom:solid 1px rgba(0,0,0,0.1)}.tplext{display:none}.sbfixed .tplext{display:table}.tpltab{xborder-top:solid 1px rgba(0,0,0,0.4);overflow-y:auto}.tpltcl,.tpltcr{box-sizing:border-box;width:50%;position:relative;margin-bottom:2px;xborder-top:solid 1px rgba(0,0,0,0.4);background:#eee;padding-left:10px;padding-right:6px}.tpltcl{float:left;border-right:solid 1px #fff;a:transparent}.tpltcr{float:right;border-left:solid 1px #fff;a:transparent}button.wh100{width:100%;height:100%}.tcrdtabcont{margin:0 -3px}.tcrdtab{width:100%;display:table;border-collapse:separate;border-spacing:3px;table-layout:fixed}.tcrdcell{display:table-cell}.tcrdtab button{width:100%;border-radius:0;padding-left:4px;padding-right:4px;overflow-x:hidden;text-overflow:ellipsis;white-space:nowrap}.tcrdtab button.active{border-color:#000}.tcrdpan{position:relative}.sbclrd .tcrdtabcont{margin:0 -1px}.sbclrd .tcrdtab{border-spacing:1px}.sbclrd .tcrdtab button{background:rgba(0,30,90,0.25);color:#fff;border-color:transparent;padding:6px 4px}.tcrdspecgs .tcrdtab button{background:#f8f8f8;color:#333;padding:5px 4px}.sbclrd .tcrdtab button.active{background:#fff;color:inherit}.butsit,.aleg .butsit{background:#eee;border-bottom-color:rgba(0,0,0,0.5)}.butsit[disabled],.aleg butsit[disabled]{opacity:.8}.tlst .awrap{display:block;padding-top:8px;padding-bottom:8px;white-space:nowrap;border-bottom:solid 1px #e4e4e4}.devtch .tlst .awrap:active{background:rgba(0,0,0,0.15)}.aleg .tlst .awrap{padding-top:8px;padding-bottom:8px}.tlst .tmaxw{display:table;width:100%;max-width:600px;table-layout:fixed}.tlst .tnum{display:table-cell;width:15%}.tlst .tplbl{display:table-cell;width:40%}.tlst .tplone{display:table-cell;width:60%}.tlst .tpllist{font-weight:bold;overflow-x:hidden;text-overflow:ellipsis}.tlst .tplnorm{font-weight:bold;overflow:hidden;text-overflow:ellipsis}.tlst .tplemp{color:#aaa}.vm0 .tlst .tplemp{display:none}.tlst .tplunav{display:none}.tlst .tpar1{display:table-cell;width:25%}.vm0 .tlst .tpar1{display:none}.tlst .tpar0{display:none}.vm0 .tlst .tavail .tpar0{display:block}.tlst .tplrn{display:none}.tlst .tavail .tplnorm .tplrn{display:inline;font-weight:normal;margin-left:.4em}.tlst .tjoin{display:table-cell;width:20%;visibility:hidden}.vm0 .tlst .tjoin,.vm1 .tlst .tjoin{width:15%}.tlst .tjoin button{padding-top:3px;padding-bottom:3px;line-height:.9}.tlst .tavail .tjoin{visibility:inherit}.hvok .tlst .awrap:hover{background:#eaecef;cursor:pointer}.turs .bwrap{border-bottom:solid 1px #e4e4e4;padding-top:1em;padding-bottom:1em}.tulst td{padding:0}.tulst .awrap{display:block;border-bottom:solid 1px #e4e4e4;white-space:nowrap;padding-top:12px;padding-bottom:12px}.vm0 .tulst .awrap{padding-top:10px;padding-bottom:10px}.aleg .tulst .awrap{padding-top:8px;padding-bottom:8px}.tulst .maxw{max-width:600px}.tulst .bl1,.tulst .bl2{display:inline-block;width:50%}.vm0 .tulst .b1{width:50%}.vm0 .tulst .b2{width:50%}.tulst .tid,.tulst .torg,.tulst .tpar,.tulst .tdtup,.tulst .tdtfin{display:inline-block}.tulst .tid{width:50%;max-width:130px}.tulst .tpar{min-width:50%}.vm0 .tulst .tid,.vm0 .tulst .tpar{display:block;width:auto}.tulst .tdtfin{width:50%}.vm0 .tulst .tdtfin,.vm0 .tulst .tpar{display:block;width:auto}.hvok .tulst tr .hv:hover{background:#eaecef;cursor:pointer}.ul{max-width:400px;width:100%;border-collapse:collapse}.ulnm{overflow:hidden;text-overflow:ellipsis}.ul td:first-child{width:70%;white-space:nowrap;padding-left:15px}.uls1{margin:2px 0}.uls1 td{padding-top:4px;padding-bottom:4px}.uls2 td{padding-top:1px;padding-bottom:1px}.uls2 .ulnm{font-weight:bold}.vm2 .aleg .uls2 .ulnm{font-weight:normal}.ulsym{margin:0 5px;font-size:12px}.ulla{color:#aaa;margin-left:4px}.ulla0{display:none}.hvok .ul tr:hover .ulla0{display:inline;color:#fff}button.ulbx{border:0;padding:0 1em;font-weight:bold;background:#e22;color:#fff}.ulnu{padding-right:10px}.ul tr.ulhead{color:#eee}.ul tr.ulhead td{padding-top:0;padding-bottom:0}.ulpan{padding:10px 15px;border-bottom:solid 1px #e4e4e4;margin-bottom:1px}td.m1ac{display:none}.ulm1 td.m1ac{display:table-cell}.ulm1 td.m0ac{display:none}.ulwp{background:#fff;position:absolute;top:0;left:0;right:0;bottom:0}.ulost{overflow-y:scroll;-webkit-overflow-scrolling:touch}.hvok .ul tr:hover{background:#eaecef;cursor:pointer}';
  if ("undefined" === typeof window.k2img) {
    window.k2img = {};
  }
  /** @type {string} */
  window.k2img["0.png"] = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAMAAABHPGVmAAAABGdBTUEAALGPC/xhBQAAAYBQTFRFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFSMVAAAAChIKAAAAAAAAAAAAAAAAFygXAAAAAAAAAAAAAAAAAAAAAAAAGCcYFycXAAAAAAAAGiwaAAAAAAAAAAAAAAAAAAAAFiYWEBoQAAAAAAAAFyYXFycXFSMVAAAAAAAAFycXFSMVFiYWFycXFycXEyATEyETAAAAAAAADxkPFicWFiYWERwRBgsGAQIBDRYNEyATFyYXFycXFiYWAAAAAAAAAAAAAAAAAAAAAAAAAAAAFycXAgUCCQ8JFSUVFiQWFyYXFygXAAAAExwTICAgAAAAAAAAAAAAFicWFyYXGCcYAAAAAAAAAAAAAAAADRQNAAAAAAAAHTAdHTQdGCoYHzQfGywbHjIeGSwZGzEbGy8bIDYgGSgZHjUeGy4bHTEdGi0aHjceHzcfHzUfGisaHDEcHTMdHjQeGSkZHDAcGSsZGCkYGiwaHTIdIDcgGSoZHzYfIDggGCgYEEyKiAAAAF90Uk5TAE0BAghMSwkDSkIGkEU/GgUEHDg8Ig4HGBEj5z8nDDZIOi4w4YEsDPPwsyUqgcDL/N6od0MeVOvQm2JPa6Lb98VHEiZJNSBB8VRrndbuVxUSBxMLCurl9R9EODIoFj4ADaGbAAAHNUlEQVRo3u1aZ1viWBQ2oQmIoSkLYgEVC4Ogow+WtTfEgo4FZ3TcdWfnmcGyoQUCITd/fRMUvTe0FObbnI8J4c055z3lnpOurt/yW1rJmeMkdLgRnh04Ph6YDW8chk4cZx0G6AlnIqkUR3E1obhI5Fs45DjTdAbC0ROlKxQPIZYUlXoc6HGoRxgKrVEVimsqmQy9FhrC1UB4Dte59kIlezwqtEimOCmSoqZCQ4ogbCcDnGRJpaInCijgma5w8mTaI9c1JzFKJgZHJ+Upc3bIKZLDU8kQ3UMbyjC4VPhUqjKeMU6xjN1qJXnGEW0SDs+FXJplGYZl0+Ufz415cR9zmHEpGI2Co/TEApGw2VIjmMleQzuT4Z7L+ucqVQQmV3gsVSiapiqZx0KO4S/9fKrXJzPVe6Fp54866hbLJADpQrHueiENAFmuu165uu1viXIWrvNElgHMv6XGbiqVGUA+iV+ruNh33grl8JvogWfeUOUWsV8s80Z7Fr/YnL25xfAT0UtROQD+LrYmbSYPQE6szEeiKZM9U6LXfABMoX1s/ADgQaRsrJdoxmRRoJcYwBalRGCFBYzIa4uBT7bGSVFkBxL8Q0vMjXnAfEcvjQcbUuwUrR9FEpSlZ+AyIBGl6UvvgqGBwUL3IhPk5CStPGARv9Cro1+tYgjNaRL5zQPI05w8lIcUUl0OBusNFkJeJAv+o2XWLBY8IRdWh+0iHmuGYvAPHgGTkZvliwxAKJY8CGx1o4p8hGOd+gke5deSR0AiQbn05Rr1Pb4G334CeSUVKw+ySAH77EdVcVCI4vKN9fIciTwXd17DXsF74JtlGRHCoQ8iqkzPB9yQKt0DFOJBShlIBZAwR4+NXwgzZC2UvgoVEVSBEyod3x+ceLdXKALdI0FRKUgJsHB4rVpGE+95cgauU+BBeUuUBnAoLBqdC663qpusNFNZphQQU6+vmN7t5YB1ZJVbq2ovhMSW0U81fsFltwIY5RgczQCYX+O6YaKWi0NIcsipAOGjHk5Ic/q7Pm2tSaFgq2bVgBQQj87o52tO0WwgVH9WA4Ia4giz+N2vVeWIgklYUgOSAWk4R+o2b2qRMoaQq6IGpILQawAbefP8MXSdAbQaEAohZwwz+uyv6Qs+rQOgBoOjkecjmN6556oHIUFKHQisSUSn99Y4HO2kuUjEXLo3kNlf53jdftDwUhcXuV9E4VlMZ3oF6dqALJRT0qjAwQin4SPdmyZdh0he+EsNSBZOK5EZyCch6hclyCXsPUOeoPRQQa8UiaZ6Pk6uX+PEgxYtFZ4vIX7n4lDEDyW5DpXfLNJ1J1ewkd23qjV937SCyhPUDGEM2xxN1OrvR64z9hK94BIG1ZMuRwZRWTG/RKaO87E4eF5rV3D4vEgxSjNLBjBwg3uM8eR6q/FdXT0VRJWsUkWQw9YcBtUscTNcJJmiMo+gJogLLoGPKIsUklqUdKqpB9Qjs5jgEqjjRvlFseCHkm4I5T7PLaPvvRcW4vGSRnIpkH3WypAo9acOeGu9R0ldFyn4Pq36iM1zy7t9jhxNT6foTg4LuMmDKrfM6LALVeVe5dhD4K+OPzOKhhKns8ibVEiQl3x0pPLiqdUVr8iOb6FutiYaRRVZWaMoUb4b5xURn+NfcjElnpWx3yXx6qfYVtwGJihyXT8kxD0xUXuTlzweFBl2cpmnlgk5xb8POiN18QUe2uT9b2kACmKzCsba3CVcjeaQmh5Oyci2rofimYUZ77abzKBPmwyfv8sZPnNHPLP0n/0Ja5PVyeka13CMzmYzoj48VcqyDcfo3OyBYKxhu6HZaNh2e0U3WQiQucJjpsLfpasLAbLJQoC7XBaY5Qy2GNjbemMNDZOVvNqI63iHeD9s2VrsNsyOqyZRXV3S8AqQLZY03JWgh97k/9pyV4P33x4pb4nWqhjzo0SbrRN+fruoFEPgFU+sG8LQbufUPXGtcAXIlxDeH/M39ov2SzrNBfHnpXyI4/EdAWN/lLiQsgjUaIle2bvG8LKej48dr58wSFs24to/An/KWjeOjRsFd1icH75qpe6AcbM7eDe3LhVife5AUMM4v9vnNkvfM+O2frt/cykqBSK6JJAK02/e+e0TNlmLbI0hEdwdWZptGxovELod024wYZC7lcet/UTANx+fPm7BqOk49gLx2Rcg+q0KvpbQuLauAz7TyPLSzGQ9QGxmaRnTVQ01YvIFFrZcyr4vwbu1bntg2GsxrsTH52aOxqL8wS8ZHTuaWR2PH7wg6IwW73DA7tZ2K/7oQ4Ahtked85tGffU/a7JSBdCPWPZ9o4OEGogqjMY8kVjY9u86TZbNEaNer9NhOp2e//9Ni8m5699eSEyYO/AhDm5znbuJvUH/zbDP6d03zZv2vU7f8I1/cI9wn7tseFdnBNdYtRPuBGHf6wsGt4PBvj07kXBPaK2aTiG8AdmsZpfWIIjWZbbaJAL8D5piaAopA1R1AAAAAElFTkSuQmCC";
  /** @type {string} */
  window.k2img["1.png"] = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAMAAABHPGVmAAAABGdBTUEAALGPC/xhBQAAAYBQTFRFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA7+/v/////Pz8x8fH/f395ubm////////AAAA5eXlLi4u////////eHh4/f39rq6u/v7+9/f3hYWF////RUVF29vb/f39lZWV////VFRU6+vrAAAA+vr6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA////AAAAAAAAAAAAAAAAAAAA29vb////9PT0AAAABwcHAAAAAAAAAAAAAAAAAAAAAAAAAAAAEhISAAAAAAAAAAAAICAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgICAAAAAAAAAAAA9/f3hoaG////////////+vr6fn5+7+/vAAAAAAAAAAAAAAAAAAAAAAAA9vb2NTU1AAAA////////3t7eAAAAAAAABwcHBAQEBgYGBgYGBgYG+fn59PT0+/v79vb28vLy/f398PDw////ZOZMlwAAAHh0Uk5TAE0BAghMSwkGQkoDBAUHETpIGxkMCufBydt0kDIl5XEhgJPNu6jkLUkhzfyo7oDsJ/VDISZHSUESHhEVRh01K90G8k9UCkUTDyNEC2E9H0BpNh87MhwNOC5TNzAt7qLydVrxne4sLz8WPjH0Vyn4CuAaF0BKRz4i1ZOOMwAABi9JREFUaN7tWelX2lgU54WEHcKiILLJomgFsYBFQQHF2o64Hq0FtS6jra12mHraORMgmfzr84Ib2SAJ4VvvBz94ePm9313fvVej+S2/pZes+i8vJj1RrzuZdHujnsmLS/+qugDT4WiD5kkj6ptWCcjs97lpEWnQbp9/cIh82Ev3EW84PxDEhK9BS5Bz34RiiL/DkiA6MOFRZda+dNPSpT01PqNAU5O0LJk9nxxFZWJcntMyhWqmxrWyrDFJK5GGHJXlPbQyaXniUslMeGmlQgTjZkRShCdFlN4imiRJ/vcf/NMkWpSIM8cNElCEMVpt+HW2kO2W0C9fxxf7qUw7IYBBPSA02wRFQQIU5NRuPuBQAoRn3vZBmXALQDCfI3ifowgGWgDm00yhJ8pqlB/LDIeWiDcxfNq8f5f/rfZC8fE+QwpelkWT5F0hVhDXGHIpQKNJ9Yt0Pplk5B4V87EJTi6hIA1CQnBAMpybTH3HRTx5lRPoLf5hETLwMhyVjZXeWYSTIg+jKTk58lAizvdCZsmzvZeSjgEForBIN4I/l/QCCgtzVdCUk7WaXNUmal+tvFDPs2utTAwGhX3g9fo8X2Fhju+ScjMwyfHkRPED14/z5xyjU7JLI8f4n9Y/L5rZIBecWxEKigmH/dXHHbbtV71sZTWVlKwmS2Gzwc0am4qfQ5xSAsI9F7LvdFtF62MTaSurvpyDifRnvIvKqlsFIryT19hHp+El/fpVIcI7GsrNjyDCQUIqJcJQYTnY1fJt4SVPelgxQtKKhWTFypjpYMn2bJJzFmVCOQhbX3d144u+/Cppi6evveWTd2aBSkINoi2aZvvXHDi+twrYvaUs2l+ivtsoZ7pvmaeGYlIdB+Yfj4H0s1E8rKsQg4AQ3YqgXoGNU/yxqnjVsjs3AMqVhcOnSFERhO03QeDYdz1aPslyD1o9kDVgsq88pq+GeiDs87NAt52xDRsE0908+fDQ1HUNsN0/0KEbHssF9A9tYnRoLgww4yPI0IKRHqs/M1E3rRCstPJik4thJcgr8JwhzdPsx4CaqV63vWNTu2hxavdeV8QPrfxWgKP4lLtYllfzIQHAwm3hsf5qL8R/OIi2IqCrnmj8Kjy3+U9ueq+OGeeryHCfqRUmCaPCowjVHtyVrpo1rNZhD4DlL90tSnTAjlGgaywDmB67X9xDaecADMWXtzAzQE0O2pi2OI3pdR1q66RgVrPFprgt9hkAut1sldWa5pPqDgvW1gFwHHd1WgIDtUHHHlQCmj1dwjlDidHgACi80UqwAsAb+xJvtjZNK0Xhj6KIOUhk85Q7koBp0jPLRSElpcoWyZu+jWEMkR3+kBAdTSodD3IHlWtb0LWMJdzMn3Oi42v8QWc/Mp15K1d9jLIWik6b0BzSOq7OyBZ6FjB9yx4Jzm2RmSilxvAZepZu80vBKjK1jU+JjNH50/+HHYEAzXKdUdb+il50/TDjFtLJ00LgYaHRcyFAp6DRwZuDwJH4KL0a/1NQMZJXG3d7jEF2S4uWHruNajwlYuPOkqbz+R5LmuDWL1h0jbWvPXc1lnczQeVPovIWY/T0rQvtvQ9Cqt9fNRRivKp3HOvQqe+/c0Jjimr87FkFYpjShx/e9l/SaVF8bko+RipUZ2yeO3G+lbII1Jot38fkYoxt/YLx8eam5tJLWzYitveBuR9yIIKhdQDdavngr68oInGhiRje3/9M3EmFuEvUGRqmdDGDG6RiQBTL0X3tn6uUJGNcQWNAGgvfvnwYschaZGv1hcDhr0i5b2hEGHtDaxiLgYJe1iKbIWM9cpYONkKJa3GE60QIcmAgNu0l55EV0cgWrW1xp2TfNG1FYmsChohFmFwIRecw2j8vLdq0GiWCmFF8pXScc4B6aO4sVi6noC/cpX6UY2dzoUr9gYRpefe4tIKjZkSjUBgYZ/b2wLihAxh4Egx0vg8wnWMjZz/JOgeB6MBoDSOFpWytuG3cWHCYdDoMAxim05kcCxvG7WItu1QYMWg1AwtisVVxV2b+9HDfvn2TM6aNuZtt+/7h6XzGhVdtFkSjjiBaKzqCF1wrmUwgkA0EMpkVVwEfQa1atRCegSxWgw3VM4LaDFaLRID/AdOAjlobcgtJAAAAAElFTkSuQmCC";
  a.prototype = new Tool;
  /** @type {function(): undefined} */
  a.prototype.constructor = a;
  /** @type {!Array} */
  var data = ["#555", "#fff"];
  self = a.prototype;
  /** @type {string} */
  self.Oe = "#45925d";
  /** @type {string} */
  self.xb = "--";
  /** @type {number} */
  self.Ib = 8;
  /** @type {number} */
  self.jc = -1;
  /** @type {number} */
  self.qd = -1;
  /** @type {number} */
  self.ed = -1;
  /** @type {number} */
  self.te = -1;
  /** @type {number} */
  self.ra = 0;
  /** @type {!Array} */
  self.da = [];
  /** @type {!Array} */
  self.uc = [];
  /**
   * @param {number} callback
   * @return {undefined}
   */
  self.ha = function (callback) {
    Tool.prototype.ha.call(this, callback);
    /** @type {!Array} */
    this.u = [];
    /** @type {number} */
    callback = 0;
    for (; 8 > callback; callback++) {
      this.u.push(Array(8));
    }
    this.rf("0.png", "1.png");
    this.Ca[1] = {
      Gb: false,
      Ob: $(this.f, "div", {
        className: "bsbb",
        style: {
          pointerEvents: "none",
          position: "absolute",
          display: "none",
          zIndex: 5,
          background: "#fff",
          opacity: "0.5"
        }
      })
    };
    /** @type {boolean} */
    this.Ca[1].kd = false;
    this.reset();
  };
  /**
   * @return {undefined}
   */
  self.jd = function () {
    var flatness_db = this.hb;
    var widnowHeight = this.gb;
    /** @type {number} */
    this.ra = Math.round(Math.min(flatness_db / 9, widnowHeight / 9));
    /** @type {number} */
    this.nb = Math.round((flatness_db - 8 * this.ra) / 2);
    /** @type {number} */
    this.tb = Math.round((widnowHeight - 8 * this.ra) / 2);
    Tool.prototype.jd.call(this);
  };
  /**
   * @return {undefined}
   */
  self.ic = function () {
    if (loadPlugin(this)) {
      ctx = this.Fb;
      var size = this.ja;
      /** @type {number} */
      var row = this.hb * size;
      /** @type {number} */
      var start = this.gb * size;
      /** @type {number} */
      var scale = this.ra * size;
      /** @type {number} */
      var width = this.nb * size;
      /** @type {number} */
      var height = this.tb * size;
      var s = this.Fb;
      var i = this.ja;
      /** @type {number} */
      var r = this.hb * i;
      /** @type {number} */
      i = i * this.gb;
      /** @type {string} */
      s.fillStyle = "#45925d";
      s.fillRect(0, 0, r, i);
      /** @type {number} */
      s = Math.max(1 * size, Math.min(row, start) / 400);
      /** @type {string} */
      ctx.fillStyle = "#346f47";
      /** @type {number} */
      size = 0;
      for (; 8 >= size; size++) {
        ctx.fillRect(Math.round(width + size * scale - s / 2), height, s, 8 * scale);
        ctx.fillRect(width, Math.round(height + size * scale - s / 2), 8 * scale, s);
      }
      var n;
      s = this.ed;
      r = this.te;
      /** @type {number} */
      size = 0;
      for (; 8 > size; size++) {
        /** @type {number} */
        i = 0;
        for (; 8 > i; i++) {
          if (-1 != (n = this.u[size][i])) {
            format(this, n, i, size);
          }
        }
      }
      /** @type {number} */
      size = scale / 7;
      if (-1 != s && -1 != r && -1 != (n = this.u[r][s])) {
        /** @type {string} */
        ctx.fillStyle = n != this.$.mb(this.$.Za) ? "#f00" : n ? "#222" : "#ddd";
        ctx.fillRect(Math.round(width + (s + .5) * scale - size / 2), Math.round(height + (r + .5) * scale - size / 2), size, size);
      }
      size = this.$.ga;
      if (this.$.dc && -1 != size) {
        /** @type {number} */
        n = (size + (1 == this.qe ? 1 : 0)) % 2;
        size = data[this.$.mb(size)];
        /** @type {number} */
        s = height / 2;
        /** @type {number} */
        scale = height + 2 * scale;
        if (height > width) {
          paint(this, row / 2, 0 == n ? start - s : s, 1 == n, size);
        } else {
          paint(this, row - width / 2, 0 == n ? start - scale : scale, 1 == n, size);
        }
      }
    }
  };
  /**
   * @param {number} value
   * @param {number} i
   * @return {undefined}
   */
  self.Vb = function (value, i) {
    /** @type {number} */
    value = Math.floor((value - this.nb) / this.ra);
    /** @type {number} */
    i = Math.floor((i - this.tb) / this.ra);
    if (0 <= value && 8 > value && 0 <= i && 8 > i && put(this, value, i)) {
      this.$.Ud(2, value + 8 * i);
      if (-1 == this.u[i][value]) {
        format(this, this.$.mb(this.$.ga), value, i);
      }
    }
  };
  /**
   * @param {number} r
   * @param {number} val
   * @return {undefined}
   */
  self.Zc = function (r, val) {
    /** @type {number} */
    r = Math.floor((r - this.nb) / this.ra);
    /** @type {number} */
    val = Math.floor((val - this.tb) / this.ra);
    if (this.jc != r || this.qd != val) {
      if (0 <= r && 8 > r && 0 <= val && 8 > val && put(this, r, val)) {
        select(this, r, val);
      } else {
        select(this, -1, -1);
      }
    }
  };
  /**
   * @param {number} a
   * @param {!Object} doc
   * @param {number} c
   * @param {number} x0
   * @return {undefined}
   */
  self.oc = function (a, doc, c, x0) {
    a = this.ra;
    extend(doc, {
      display: "block",
      left: Math.floor(this.nb + c * a) + "px",
      top: Math.floor(this.tb + x0 * a) + "px",
      width: Math.round(a) + "px",
      height: Math.round(a) + "px"
    });
  };
  /**
   * @param {string} flag
   * @return {undefined}
   */
  self.setActive = function (flag) {
    Tool.prototype.setActive.call(this, flag);
    if (!flag) {
      select(this, -1, -1);
    }
  };
  /**
   * @return {undefined}
   */
  self.reset = function () {
    /** @type {number} */
    var i = 0;
    for (; 8 > i; i++) {
      /** @type {number} */
      var textMethod = 0;
      for (; 8 > textMethod; textMethod++) {
        /** @type {number} */
        this.u[i][textMethod] = -1;
      }
    }
    /** @type {number} */
    this.ed = -1;
    select(this, -1, -1);
    if (!this.$.kc) {
      /** @type {number} */
      this.u[3][3] = 1;
      /** @type {number} */
      this.u[3][4] = 0;
      /** @type {number} */
      this.u[4][3] = 0;
      /** @type {number} */
      this.u[4][4] = 1;
    }
  };
  /**
   * @param {number} offset
   * @return {undefined}
   */
  self.qb = function (offset) {
    this.reset();
    /** @type {number} */
    var endOfCentralDirOffset = offset;
    /** @type {number} */
    var display_name = -1;
    /** @type {number} */
    var layer_i = 0;
    for (; layer_i < this.da.length; layer_i++) {
      if (-1 == this.da[layer_i]) {
        if (endOfCentralDirOffset--, 0 > endOfCentralDirOffset) {
          break;
        }
      } else {
        /** @type {number} */
        var value = this.da[layer_i] % 8;
        /** @type {number} */
        var i = Math.floor(this.da[layer_i] / 8) % 8;
        /** @type {number} */
        var key = Math.floor(this.da[layer_i] / 64) % 2;
        if (!(0 > value || 8 <= value || 0 > i || 8 <= i || 0 > key || 1 < key)) {
          if (key == display_name) {
            endOfCentralDirOffset--;
          }
          if (0 > endOfCentralDirOffset) {
            break;
          }
          if (-1 != this.da[layer_i]) {
            /** @type {number} */
            this.ed = value;
            /** @type {number} */
            this.te = i;
            /** @type {number} */
            this.u[i][value] = key;
            /** @type {number} */
            display_name = value;
            /** @type {number} */
            value = key;
            /** @type {!Array} */
            var possible = [0, 1, 1, 1, 0, -1, -1, -1];
            /** @type {!Array} */
            var error = [-1, -1, 0, 1, 1, 1, 0, -1];
            /** @type {number} */
            var index = 0;
            for (; 8 > index; index++) {
              /** @type {number} */
              var key = display_name;
              /** @type {number} */
              var val = i;
              /** @type {number} */
              var x = 0;
              for (; ;) {
                key = key + possible[index];
                val = val + error[index];
                x++;
                if (0 > key || 8 <= key || 0 > val || 8 <= val) {
                  break;
                }
                if (this.u[val][key] == value) {
                  if (1 < x) {
                    for (; ;) {
                      /** @type {number} */
                      key = key - possible[index];
                      /** @type {number} */
                      val = val - error[index];
                      if (key == display_name && val == i) {
                        break;
                      }
                      /** @type {number} */
                      this.u[val][key] = value;
                    }
                  }
                  break;
                } else {
                  if (-1 == this.u[val][key]) {
                    break;
                  }
                }
              }
            }
            endOfCentralDirOffset--;
            if (0 > endOfCentralDirOffset) {
              break;
            }
            /** @type {number} */
            display_name = key;
          }
        }
      }
    }
    this.$.ye(offset);
    isArray(this);
  };
  /**
   * @param {!Object} val
   * @return {undefined}
   */
  self.history = function (val) {
    /** @type {!Object} */
    this.da = val;
  };
  /**
   * @param {!Object} c
   * @return {?}
   */
  self.Dd = function (c) {
    if (!(3 > c.length)) {
      c = c[2];
      /** @type {number} */
      var headerEndIndex = 0;
      /** @type {number} */
      var item = -2;
      if (0 < this.da.length && -1 == this.da[0]) {
        headerEndIndex++;
      }
      this.da.push(c);
      if (this.da.length > headerEndIndex + 1) {
        /** @type {number} */
        item = Math.floor(this.da[this.da.length - 2] / 64) % 2;
      }
      /** @type {!Array} */
      var html = [];
      if (this.da.length > headerEndIndex + 1 && item == Math.floor(c / 64) % 2) {
        html.push("--");
      }
      html.push(recurse(c));
      return html;
    }
  };
  /**
   * @return {undefined}
   */
  window.k2start = function () {
    new move({
      wf: me,
      zf: 1713,
      nf: 2
    });
  };
  /**
   * @return {undefined}
   */
  window.k2play = function () {
    new addRoomButton;
  };
  me.prototype = new Promise;
  self = me.prototype;
  /** @type {function(): undefined} */
  self.constructor = me;
  /**
   * @param {?} callback
   * @return {undefined}
   */
  self.ha = function (callback) {
    Promise.prototype.ha.call(this, callback, a, {
      $c: true,
      ne: true,
      Wb: "PGN",
      pd: true,
      md: data
    });
  };
  /**
   * @param {number} bytes
   * @return {?}
   */
  self.mb = function (bytes) {
    return this.u.Qc ? 1 - bytes : bytes;
  };
  /**
   * @param {!Object} a
   * @param {number} e
   * @return {?}
   */
  self.ie = function (a, e) {
    if (5 == a[e]) {
      var i = a[e + 1];
      this.u.uc = a.slice(e + 2, e + 2 + i);
      return e + 2 + i;
    }
    return e;
  };
  /**
   * @param {!Object} scope
   * @return {undefined}
   */
  self.Ge = function (scope) {
    var b = this;
    $(scope, "div", {
      className: "mbsp"
    }, [this.N = $("input", {
      type: "checkbox",
      onchange: function () {
        setTimeout(b, "xot", this.checked ? 1 : 0);
      }
    }), "random 8 (xot)"]);
    this.Xa.push(this.N);
  };
  /**
   * @param {string} f
   * @param {number} args
   * @return {undefined}
   */
  self.Fe = function (f, args) {
    if ("xot" == f) {
      /** @type {boolean} */
      this.N.checked = 0 != args;
    }
  };
  /**
   * @param {!Object} a
   * @return {?}
   */
  self.Ee = function (a) {
    /** @type {!Array} */
    var children = [];
    /** @type {number} */
    var curr_val = -2;
    /** @type {number} */
    var j = 0;
    var startLen = a.length;
    for (; j < startLen; j++) {
      var i = a[j];
      /** @type {number} */
      var new_curr_val = -1 != i ? Math.floor(i / 64) % 2 : -1;
      if (new_curr_val == curr_val) {
        children.push(this.u.xb);
      }
      children.push(recurse(i));
      /** @type {number} */
      curr_val = new_curr_val;
    }
    return children;
  };
  /** @type {!Object} */
  e.prototype = Object.create(show.prototype);
  /** @type {!Object} */
  self = e.prototype;
  /** @type {function(): undefined} */
  self.constructor = e;
  /**
   * @return {?}
   */
  self.he = function () {
    return _build(this.u, 0) + ":" + _build(this.u, 1);
  };
  /**
   * @return {?}
   */
  self.fe = function () {
    return this.W ? "xot" : "";
  };
  /**
   * @param {string} s
   * @return {?}
   */
  self.xe = function (s) {
    /** @type {!Array} */
    var next = [];
    /** @type {!Array} */
    var jisps = [];
    /** @type {number} */
    var value = 1;
    var dy = this.u.Ib;
    this.W = window.k2pback.v;
    /** @type {!Array} */
    var line = [];
    if (0 <= s.indexOf(" ")) {
      line = s.split(" ");
    } else {
      var l = s.length;
      /** @type {number} */
      var i = 0;
      for (; i < l && !(i + 1 >= l);) {
        if ("--" == s.substring(i, i + 2)) {
          line.push(this.u.xb);
        } else {
          var e = s.charAt(i).toLowerCase();
          var side = s.charAt(i + 1);
          if (!("a" <= e && "h" >= e && "1" <= side && "8" >= side)) {
            break;
          }
          line.push(e + side);
        }
        /** @type {number} */
        i = i + 2;
      }
    }
    for (; line.length;) {
      s = line.shift();
      if (!line.length && 0 < s.indexOf("-")) {
        break;
      }
      if (0 != s.length && "." != s.charAt(s.length - 1)) {
        jisps.push(s);
        /** @type {number} */
        value = 1 - value;
        if (!("-" == s.charAt(0) || 2 > s.length)) {
          next.push(s.charCodeAt(0) - 97 + (s.charCodeAt(1) - 49) * dy + value * dy * dy);
        }
      }
    }
    return {
      J: next,
      O: jisps
    };
  };
  /**
   * @param {number} f
   * @param {number} a
   * @return {undefined}
   */
  self.Ud = function (f, a) {
    var valueProgess = this;
    var h = this.ga;
    var currentIndex = this.Y;
    /** @type {number} */
    f = this.u.Ib * this.u.Ib;
    if (!this.C) {
      createElement(this, true);
    }
    this.a = this.a.slice(0, currentIndex + 1);
    /** @type {number} */
    var index = 0;
    /** @type {number} */
    var i = this.a.length - 1;
    for (; 0 <= i; i--) {
      if (this.a[i] == this.u.xb) {
        index++;
      }
    }
    this.g = this.g.slice(0, currentIndex + 1 - index);
    /** @type {number} */
    var l = 0 < this.g.length ? Math.floor(this.g[this.g.length - 1] / f) % 2 : -1;
    if (h == l && 0 < this.a.length && this.a[this.a.length - 1] != this.u.xb) {
      this.a.push(this.u.xb);
    }
    this.g.push(a + h * f);
    this.a.push(recurse(a));
    this.u.history(this.g, this.a);
    if (window.location && window.history.replaceState && window.k2pback.x) {
      try {
        window.history.replaceState({}, "", location.protocol + "//" + location.hostname + window.k2pback.x + "+" + this.a.join("") + "#");
      } catch (l) {
      }
    }
    setTimeout(function () {
      return require(valueProgess, currentIndex + 1 + (h == l ? 1 : 0));
    }, 0);
  };
  /**
   * @return {undefined}
   */
  self.ce = function () {
    var newValue = this.ga;
    var text = String(this, newValue);
    if (0 == text.length) {
      text = String(this, 1 - newValue);
      if (0 < text.length) {
        /** @type {number} */
        this.ga = 1 - newValue;
      }
    }
    this.u.uc = text;
  };
}).call(this);

// lk16:0 send keep-alive every 30 sec i:[] s:absent
// lk16:2 send keep-alive every 30 sec i:[2] s:absent
// lk16:73 send leave room i:[73,room ID] s:absent
// lk16:83 send sit down