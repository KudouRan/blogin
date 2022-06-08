package blogin

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"

	"github.com/imroc/req/v3"
)

var Html1 = `<!doctype html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>扫码登录</title>
    <style>
        .qr {
            width: 30%;
            min-width: 300px;
        }

        .center {
            margin: 10px auto;
            text-align: center;
        }
		
		.msg span {
			margin: 0;
			font-size: 14px;
		}
    </style>
</head>
<body>
<script>
    // @formatter:off
    !function(r,b){"use strict";for(var x="length",m=[null,[[10,7,17,13],[1,1,1,1],[]],[[16,10,28,22],[1,1,1,1],[4,16]],[[26,15,22,18],[1,1,2,2],[4,20]],[[18,20,16,26],[2,1,4,2],[4,24]],[[24,26,22,18],[2,1,4,4],[4,28]],[[16,18,28,24],[4,2,4,4],[4,32]],[[18,20,26,18],[4,2,5,6],[4,20,36]],[[22,24,26,22],[4,2,6,6],[4,22,40]],[[22,30,24,20],[5,2,8,8],[4,24,44]],[[26,18,28,24],[5,4,8,8],[4,26,48]],[[30,20,24,28],[5,4,11,8],[4,28,52]],[[22,24,28,26],[8,4,11,10],[4,30,56]],[[22,26,22,24],[9,4,16,12],[4,32,60]],[[24,30,24,20],[9,4,16,16],[4,24,44,64]],[[24,22,24,30],[10,6,18,12],[4,24,46,68]],[[28,24,30,24],[10,6,16,17],[4,24,48,72]],[[28,28,28,28],[11,6,19,16],[4,28,52,76]],[[26,30,28,28],[13,6,21,18],[4,28,54,80]],[[26,28,26,26],[14,7,25,21],[4,28,56,84]],[[26,28,28,30],[16,8,25,20],[4,32,60,88]],[[26,28,30,28],[17,8,25,23],[4,26,48,70,92]],[[28,28,24,30],[17,9,34,23],[4,24,48,72,96]],[[28,30,30,30],[18,9,30,25],[4,28,52,76,100]],[[28,30,30,30],[20,10,32,27],[4,26,52,78,104]],[[28,26,30,30],[21,12,35,29],[4,30,56,82,108]],[[28,28,30,28],[23,12,37,34],[4,28,56,84,112]],[[28,30,30,30],[25,12,40,34],[4,32,60,88,116]],[[28,30,30,30],[26,13,42,35],[4,24,48,72,96,120]],[[28,30,30,30],[28,14,45,38],[4,28,52,76,100,124]],[[28,30,30,30],[29,15,48,40],[4,24,50,76,102,128]],[[28,30,30,30],[31,16,51,43],[4,28,54,80,106,132]],[[28,30,30,30],[33,17,54,45],[4,32,58,84,110,136]],[[28,30,30,30],[35,18,57,48],[4,28,56,84,112,140]],[[28,30,30,30],[37,19,60,51],[4,32,60,88,116,144]],[[28,30,30,30],[38,19,63,53],[4,28,52,76,100,124,148]],[[28,30,30,30],[40,20,66,56],[4,22,48,74,100,126,152]],[[28,30,30,30],[43,21,70,59],[4,26,52,78,104,130,156]],[[28,30,30,30],[45,22,74,62],[4,30,56,82,108,134,160]],[[28,30,30,30],[47,24,77,65],[4,24,52,80,108,136,164]],[[28,30,30,30],[49,25,81,68],[4,28,56,84,112,140,168]]],l=/^\d*$/,s=/^[A-Za-z0-9 $%*+\-./:]*$/,c=/^[A-Z0-9 $%*+\-./:]*$/,v=[],h=[-1],t=0,e=1;t<255;++t)v.push(e),h[e]=t,e=2*e^(128<=e?285:0);for(var p=[[]],n=0;n<30;++n){for(var a=p[n],o=[],u=0;u<=n;++u){var i=u<n?v[a[u]]:0,f=v[(n+(a[u-1]||0))%255];o.push(h[i^f])}p.push(o)}for(var d={},g=0;g<45;++g)d["0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:".charAt(g)]=g;function w(r){return 6<r}function F(r,t){var e=-8&function(r){var t=m[r],e=16*r*r+128*r+64;return w(r)&&(e-=36),t[2][x]&&(e-=25*t[2][x]*t[2][x]-10*t[2][x]-55),e}(r),n=m[r];return e-=8*n[0][t]*n[1][t]}function N(r,t){switch(t){case 1:return r<10?10:r<27?12:14;case 2:return r<10?9:r<27?11:13;case 4:return r<10?8:16;case 8:return r<10?8:r<27?10:12}}function A(r,t,e){var n=F(r,e)-4-N(r,t);switch(t){case 1:return 3*(n/10|0)+(n%10<4?0:n%10<7?1:2);case 2:return 2*(n/11|0)+(n%11<6?0:1);case 4:return n/8|0;case 8:return n/13|0}}function C(r,t){for(var e=r.slice(0),n=r[x],a=t[x],o=0;o<a;++o)e.push(0);for(var u=0;u<n;){var i=h[e[u++]];if(0<=i)for(var f=0;f<a;++f)e[u+f]^=v[(i+t[f])%255]}return e.slice(n)}function S(r,t,e,n){for(var a=r<<n,o=t-1;0<=o;--o)a>>n+o&1&&(a^=e<<o);return r<<n|a}function y(r,t,e){for(var n=L[e],a=r[x],o=0;o<a;++o)for(var u=0;u<a;++u)t[o][u]||(r[o][u]^=n(o,u));return r}function M(r,t,e,n){for(var a=r[x],o=21522^S(e<<3|n,5,1335,10),u=0;u<15;++u){var i=[a-1,a-2,a-3,a-4,a-5,a-6,a-7,a-8,7,5,4,3,2,1,0][u];r[[0,1,2,3,4,5,7,8,a-7,a-6,a-5,a-4,a-3,a-2,a-1][u]][8]=r[8][i]=o>>u&1}return r}function E(r){for(var t=function(r){for(var t=0,e=0;e<r[x];++e)5<=r[e]&&(t+=r[e]-5+3);for(var n=5;n<r[x];n+=2){var a=r[n];r[n-1]===a&&r[n-2]===3*a&&r[n-3]===a&&r[n-4]===a&&(r[n-5]>=4*a||r[n+1]>=4*a)&&(t+=40)}return t},e=r[x],n=0,a=0,o=0;o<e;++o){var u,i=r[o];u=[0];for(var f=0;f<e;){var l;for(l=0;f<e&&i[f];++l)++f;for(u.push(l),l=0;f<e&&!i[f];++l)++f;u.push(l)}n+=t(u),u=[0];for(var s=0;s<e;){var c;for(c=0;s<e&&r[s][o];++c)++s;for(u.push(c),c=0;s<e&&!r[s][o];++c)++s;u.push(c)}n+=t(u);var v=r[o+1]||[];a+=i[0];for(var h=1;h<e;++h){var p=i[h];a+=p,i[h-1]===p&&v[h]===p&&v[h-1]===p&&(n+=3)}}return n+=10*(Math.abs(a/e/e-.5)/.05|0)}function k(r,t,e,n,a){var o=m[t],u=function(r,t,e,n){function a(r,t){if(i<=t){for(o.push(u|r>>(t-=i));8<=t;)o.push(r>>(t-=8)&255);u=0,i=8}0<t&&(u|=(r&(1<<t)-1)<<(i-=t))}var o=[],u=0,i=8,f=e[x],l=N(r,t);switch(a(t,4),a(f,l),t){case 1:for(var s=2;s<f;s+=3)a(parseInt(e.substring(s-2,s+1),10),10);a(parseInt(e.substring(s-2),10),[0,4,7][f%3]);break;case 2:for(var c=1;c<f;c+=2)a(45*d[e.charAt(c-1)]+d[e.charAt(c)],11);f%2==1&&a(d[e.charAt(c-1)],6);break;case 4:for(var v=0;v<f;++v)a(e[v],8)}for(a(0,4),i<8&&o.push(u);o[x]+1<n;)o.push(236,17);return o[x]<n&&o.push(236),o}(t,e,r,F(t,n)>>3);u=function(r,t,e){for(var n=[],a=r[x]/t|0,o=0,u=t-r[x]%t,i=0;i<u;++i)n.push(o),o+=a;for(var f=u;f<t;++f)n.push(o),o+=1+a;n.push(o);for(var l=[],s=0;s<t;++s)l.push(C(r.slice(n[s],n[s+1]),e));for(var c=[],v=r[x]/t|0,h=0;h<v;++h)for(var p=0;p<t;++p)c.push(r[n[p]+h]);for(var d=u;d<t;++d)c.push(r[n[d+1]-1]);for(var g=0;g<e[x];++g)for(var m=0;m<t;++m)c.push(l[m][g]);return c}(u,o[1][n],p[o[0][n]]);var i=function(r){for(var t=m[r],e=function(r){return 4*r+17}(r),i=[],f=[],n=0;n<e;++n)i.push([]),f.push([]);function a(r,t,e,n,a){for(var o=0;o<e;++o)for(var u=0;u<n;++u)i[r+o][t+u]=a[o]>>u&1,f[r+o][t+u]=1}a(0,0,9,9,[127,65,93,93,93,65,383,0,64]),a(e-8,0,8,9,[256,127,65,93,93,93,65,127]),a(0,e-8,9,8,[254,130,186,186,186,130,254,0,0]);for(var o=9;o<e-8;++o)i[6][o]=i[o][6]=1&~o,f[6][o]=f[o][6]=1;for(var u=t[2],l=u[x],s=0;s<l;++s)for(var c=0===s?l-1:l,v=0===s||s===l-1?1:0;v<c;++v)a(u[s],u[v],5,5,[31,17,21,17,31]);if(w(r))for(var h=S(r,6,7973,12),p=0,d=0;d<6;++d)for(var g=0;g<3;++g)i[d][e-11+g]=i[e-11+g][d]=h>>p++&1,f[d][e-11+g]=f[e-11+g][d]=1;return{matrix:i,reserved:f}}(t),f=i.matrix,l=i.reserved;if(function(r,t,e){for(var n=r[x],a=0,o=-1,u=n-1;0<=u;u-=2){6===u&&--u;for(var i=o<0?n-1:0,f=0;f<n;++f){for(var l=u;u-2<l;--l)t[i][l]||(r[i][l]=e[a>>3]>>(7&~a)&1,++a);i+=o}o=-o}}(f,l,u),a<0){y(f,l,0),M(f,0,n,0);var s=0,c=E(f);for(y(f,l,0),a=1;a<8;++a){y(f,l,a),M(f,0,n,a);var v=E(f);v<c&&(c=v,s=a),y(f,l,a)}a=s}return y(f,l,a),M(f,0,n,a),f}var L=[function(r,t){return(r+t)%2==0},function(r){return r%2==0},function(r,t){return t%3==0},function(r,t){return(r+t)%3==0},function(r,t){return((r/2|0)+(t/3|0))%2==0},function(r,t){return r*t%2+r*t%3==0},function(r,t){return(r*t%2+r*t%3)%2==0},function(r,t){return((r+t)%2+r*t%3)%2==0}],R={generate:function(r,t){var e=t||{},n={numeric:1,alphanumeric:2,octet:4},a={L:1,M:0,Q:3,H:2},o=e.version||-1,u=a[(e.ecclevel||"L").toUpperCase()],i=e.mode?n[e.mode.toLowerCase()]:-1,f="mask"in e?e.mask:-1;if(i<0)i="string"==typeof r?r.match(l)?1:r.match(c)?2:4:4;else if(1!==i&&2!==i&&4!==i)throw"invalid or unsupported mode";if(null===(r=function(r,t){switch(r){case 1:return t.match(l)?t:null;case 2:return t.match(s)?t.toUpperCase():null;case 4:if("string"!=typeof t)return t;for(var e=[],n=0;n<t[x];++n){var a=t.charCodeAt(n);a<128?e.push(a):a<2048?e.push(192|a>>6,128|63&a):a<65536?e.push(224|a>>12,128|a>>6&63,128|63&a):e.push(240|a>>18,128|a>>12&63,128|a>>6&63,128|63&a)}return e}}(i,r)))throw"invalid data format";if(u<0||3<u)throw"invalid ECC level";if(o<0){for(o=1;o<=40&&!(r[x]<=A(o,i,u));++o);if(40<o)throw"too large data"}else if(o<1||40<o)throw"invalid version";if(-1!==f&&(f<0||8<f))throw"invalid mask";return k(r,o,i,u,f)},generateHTML:function(r,t){for(var e=t||{},n=e.fillcolor?e.fillcolor:"#FFFFFF",a=e.textcolor?e.textcolor:"#000000",o=R.generate(r,e),u=Math.max(e.modulesize||5,.5),i=Math.max(null!==e.margin?e.margin:4,0),f=b.createElement("div"),l=o[x],s=['<table border="0" cellspacing="0" cellpadding="0" style="border:'+u*i+"px solid "+n+";background:"+n+'">'],c=0;c<l;++c){s.push("<tr>");for(var v=0;v<l;++v)s.push('<td style="width:'+u+"px;height:"+u+"px"+(o[c][v]?";background:"+a:"")+'"></td>');s.push("</tr>")}f.className="qrcode";var h=b.createRange();h.selectNodeContents(f);var p=h.createContextualFragment(s.join("")+"</table>");return f.appendChild(p),f},generateSVG:function(r,t){var e=t||{},n=e.fillcolor?e.fillcolor:"#FFFFFF",a=e.textcolor?e.textcolor:"#000000",o=R.generate(r,e),u=o[x],i=Math.max(e.modulesize||5,.5),f=Math.max(e.margin?e.margin:4,0),l=i*(u+2*f),s=b.createElementNS("http://www.w3.org/2000/svg","svg");s.setAttributeNS(null,"viewBox","0 0 "+l+" "+l),s.setAttributeNS(null,"style","shape-rendering:crispEdges");var c="qrcode"+Date.now();s.setAttributeNS(null,"id",c);var v=b.createDocumentFragment(),h=b.createElementNS("http://www.w3.org/2000/svg","style");h.appendChild(b.createTextNode("#"+c+" .bg{fill:"+n+"}#"+c+" .fg{fill:"+a+"}")),v.appendChild(h);function p(r,t,e,n,a){var o=b.createElementNS("http://www.w3.org/2000/svg","rect")||"";return o.setAttributeNS(null,"class",r),o.setAttributeNS(null,"fill",t),o.setAttributeNS(null,"x",e),o.setAttributeNS(null,"y",n),o.setAttributeNS(null,"width",a),o.setAttributeNS(null,"height",a),o}v.appendChild(p("bg","none",0,0,l));for(var d=f*i,g=0;g<u;++g){for(var m=f*i,w=0;w<u;++w)o[g][w]&&v.appendChild(p("fg","none",m,d,i)),m+=i;d+=i}return s.appendChild(v),s},generatePNG:function(r,t){var e,n=t||{},a=n.fillcolor||"#FFFFFF",o=n.textcolor||"#000000",u=R.generate(r,n),i=Math.max(n.modulesize||5,.5),f=Math.max(null!==n.margin&&void 0!==n.margin?n.margin:4,0),l=u[x],s=i*(l+2*f),c=b.createElement("canvas");if(c.width=c.height=s,!(e=c.getContext("2d")))throw"canvas support is needed for PNG output";e.fillStyle=a,e.fillRect(0,0,s,s),e.fillStyle=o;for(var v=0;v<l;++v)for(var h=0;h<l;++h)u[v][h]&&e.fillRect(i*(f+h),i*(f+v),i,i);return c.toDataURL()}};r.QRCode=R}("undefined"!=typeof window?window:this,document);
    // @formatter:on
</script>
<script>
    // 包装一下 fetch
    function get(url, data) {
        return fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        }).then(function (response) {
            return response.json();
        });
    }
</script>

<h1 class="center">请使用客户端扫描后登录</h1>
<p>想要重新生成请刷新页面...</p>
<div class="qr center"></div>
<div class="msg"></div>
<script>
    var div = document.querySelector('.qr'),
        qr = QRCode.generateSVG("`
var Html2 = `", {
            ecclevel: "M",
            fillcolor: "#F2F2F2",
            textcolor: "#fd84b0",
            margin: 1,
            modulesize: 8
        });
    div.appendChild(qr);
    var msg = document.querySelector('.msg');
    var count = 0;
    var timer = setInterval(function () {
        get('/poll').then(function (data) {
            console.log(data);
            if (data.code === 0) {
                msg.innerHTML = "<div class=\"center\">"
                + "<h2>登录成功</h2>"
                + "<p>用户 mid：" + data.data.mid + "</p>"
                + "<p>用户 access_token：" + data.data.access_token + "</p>"
                + "<p>用户 refresh_token：" + data.data.refresh_token + "</p>"
                + "<p>过期时间 expires_in：" + data.data.expires_in + "</p>"
                + "</div>";
                div.remove();
                clearInterval(timer);
                return;
            }
            if (count > 6) {
                msg.innerHTML = '';
                count = 0;
            }
            if (data.code === 86038) {
                msg.innerHTML = '';
                clearInterval(timer);
            }
            var item = document.createElement('p')
            item.innerHTML = data.message;
            msg.appendChild(item);
            count++;
        });
    }, 3000);
</script>
</body>
</html>`

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
}

type AuthCodeData struct {
	Url      string `json:"url"`
	AuthCode string `json:"auth_code"`
}

type AuthCode struct {
	*BaseResponse
	Data AuthCodeData `json:"data"`
}

type QrLoginData struct {
	Mid          int    `json:"mid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type QrLogin struct {
	*BaseResponse
	Data QrLoginData `json:"data"`
}

func (l QrLogin) ToJson() []byte {
	b, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return b
}

func GetClient() *req.Client {
	return req.C().SetCommonHeaders(map[string]string{
		"user-agent":   "Mozilla/5.0 BiliDroid/6.74.0 (bbcallen@gmail.com) os/android model/MI 10 Pro mobi_app/android build/6740400 channel/xiaomi innerVer/xiaomi osVer/10 network/2",
		"content-type": "application/x-www-form-urlencoded; charset=utf-8",
	})
}

/**
 * app sign 签名
 * @param params *map[string]string
 */
func AppSign(params map[string]string) string {
	appkey := "4409e2ce8ffd12b8"
	appsec := "59b43e04ad6965f34319062b478f83dd"
	// 给 params 加入 appkey
	(params)["appkey"] = appkey
	// 对 params 进行排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 将 params 编码成 query string
	var query string
	for _, k := range keys {
		query += fmt.Sprintf("%s=%s&", k, url.QueryEscape((params)[k]))
	}
	// 对 query 进行 md5
	hash := md5.New()
	hash.Write([]byte(query[:len(query)-1] + appsec))
	(params)["sign"] = hex.EncodeToString(hash.Sum(nil))
	return query + "sign=" + (params)["sign"]
}
