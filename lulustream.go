// Package lulustream provides a wrapper for lulustream.com
package lulustream

import (
	"fmt"
	"html"
	"net/url"
	"os"
	"os/exec"
	"regexp"

	"github.com/dop251/goja"
	"github.com/kissanjamgit/ext"
	"resty.dev/v3"
)

// map[
// Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
// Accept-Encoding:gzip, deflate Accept-Language:en-US,en;q=0.5
// Connection:keep-alive
//
//	DNT:1 Priority:u=0, i
//	Sec-Fetch-Dest:document
//	Sec-Fetch-Mode:navigate
//	Sec-Fetch-Site:cross-site
//	Sec-GPC:1
//	TE:trailers
//	Upgrade-Insecure-Requests:1
//	User-Agent :Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:135.0) Gecko/20100101 Firefox/135.0]
var header = map[string]string{
	`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:152.0) Gecko/20100101 Firefox/152.0`,
	`Accept`:          `*/*`,
	`Accept-Language`: `en-US,en;q=0.9`,
	`Accept-Encoding`: `gzip, deflate, br, zstd`,
	`Origin`:          `https://luluvdo.com`,
	`Sec-GPC`:         `1`,
	`Connection`:      `keep-alive`,
	`Referer`:         `https://luluvdo.com/`,
	`Sec-Fetch-Dest`:  `empty`,
	`Sec-Fetch-Mode`:  `cors`,
	`Sec-Fetch-Site`:  `cross-site`,
	`Priority`:        `u=0`,
	`Pragma`:          `no-cache`,
	`Cache-Control`:   `no-cache`,
}

type Lulustream struct {
	Source string
}

func New(source string) ext.Site {
	return &Lulustream{Source: source}
}

func (l *Lulustream) Resource(client *resty.Client) (cr ext.ContentResource, err error) {
	js := `(function(p,a,c,k,e,d){while(c--)if(k[c])p=p.replace(new RegExp('\\b'+c.toString(a)+'\\b','g'),k[c]);return p}('l("36").82({80:[{28:"z://7z.7y.1w/7x/7w/7v/7u/7t.7s?t=7r-7q-7p&s=7o&e=7n&f=7m&i=0.3&7l=0"}],7k:"z://7j.7i.1h/7h.7g",1z:"y%",1y:"y%",7f:{1x:"/b/1i/1i-3s.20",13:"1i-3s"},7e:"7d",7c:"7b.81",7a:\'79\',77:"1b",g:[{28:"/3r/3i.3r",76:"75 1c",74:"1c"}],1c:{73:1,71:\'#70\',6z:\'#6y\',6x:"6w",6v:30,6u:\'y\',},\'6t\':{"6s":"6r"},6q:"6p",6o:"z://27.1h",3q:{28:"z://27.1h/b/1i/b-3q.n",6n:"z://27.1h",1f:"6m-6l",3b:"5"}});m 25,26;m 6k=0,6j=0;m b=l();m 3p=0,6i=0,6h=0,6g=0;$.6f({6e:{\'6d-6b\':\'39-6a\'}});b.o(\'2h\',c(x){h(5>0&&x.1f>=5&&26!=1){26=1;$(\'q.69\').68(\'67\')}});b.o(\'1u\',c(x){3p=x.1f});b.o(\'38\',c(x){3o(x)});b.o(\'66\',c(){$(\'q.3n\').65()});b.o(\'3h\',c(x){});c 3o(x){$(\'q.3n\').1s();$(\'#64\').1s();h(25)1k;25=1;1g=0;h(63.62===61){1g=1}$.5z(\'/5y?3g=5x&3f=3e&5w=&5v=&5u=&1g=\'+1g,c(3m){$(\'#5t\').5s(3m)})}c 5r(){m g=b.1l(3l);3k.3j(g);h(g.v>1){2b(i=0;i<g.v;i++){h(g[i].13==3l){3k.3j(\'!!=\'+i);b.29(i)}}}}b.o(\'5q\',c(){b.o(\'5p\',c(21){h(5o(\'3i\').5n(21.g[21.5m].5l)){l().3h(1b);l().5k(0);3d(\'/?3g=5j&3f=3e\')}});c 3d(1x){m $1e=$("<q />").20({1f:"5i",1z:"y%",1y:"y%",5h:0,3a:0,3c:5g,5f:"5e(10%, 10%, 10%, 0.4)","2g-5d":"5c"});$("<5b />").20({1z:"60%",1y:"60%",3c:5a,"3b-3a":"59"}).58({\'57\':1x,\'56\':\'0\',\'55\':\'39\'}).37($1e);$1e.54(c(){$(53).52();l().38()});$1e.37($(\'#36\'))}l().1r(\'<n 33="31://2z.2y.1w/2x/n" 2w="a-n-w" 2v="0 0 2u 2t"><2s d="51.5 2r.6c-9.5 7.9-22.8 9.7-34.1 4.50 4z.4 0 4y 4x.6 7.2 72.3 18.4 4w.5-3.6 34.1 4.4v 4u 2q-17.7 14.3-32 32-2p 14.3 32 4t 17.7-14.3 32-32 4s-32-14.3-32-4r-11.5 9.6-1v 2o"/></n>\',"4q 10 2n",c(){l().1u(l().2m()+10)},"35");$("q[2l=35]").2j().2i(\'.a-w-1t\');l().1r(\'<n 33="31://2z.2y.1w/2x/n" 2w="a-n-w" 2v="0 0 2u 2t"><2s d="4p.5 2r.4o.5 7.9 22.8 9.7 34.1 4.4n.4-16.6 18.4-4m-12.4-7.2-23.7-18.4-4k-24.5-3.6-34.1 4.4l-1v 4j 2q-17.7-14.3-32-32-4i 78.3 0 4h 17.7 14.3 32 32 2p-14.3 32-4g.5 9.6 1v 2o"/></n>\',"4f 10 2n",c(){m 1d=l().2m()-10;h(1d<0)1d=0;l().1u(1d)},"2k");$("q[2l=2k]").2j().2i(\'.a-w-1t\');$("q.a-w-1t").1s();$(\'.a-4e-2h\').4d($(\'.a-2g-4c\'))});b.o("p",c(15){m g=b.1l();h(g.v<2)1k;$(\'.a-k-4b-4a\').49(c(){$(\'#a-k-j-p\').1o(\'a-k-j-19\');$(\'.a-j-p\').u(\'r-1p\',\'1a\')});b.1r("/48/47.n","46 45",c(){$(\'.a-2f\').44(\'a-k-2e\');$(\'.a-k-1c, .a-k-43\').u(\'r-1q\',\'1a\');h($(\'.a-2f\').42(\'a-k-2e\')){$(\'.a-j-p\').u(\'r-1q\',\'1b\');$(\'.a-j-p\').u(\'r-1p\',\'1b\');$(\'.a-k-j-41\').1o(\'a-k-j-19\');$(\'.a-k-j-p\').40(\'a-k-j-19\')}3z{$(\'.a-j-p\').u(\'r-1q\',\'1a\');$(\'.a-j-p\').u(\'r-1p\',\'1a\');$(\'.a-k-j-p\').1o(\'a-k-j-19\')}},"3y");b.o("3x",c(15){1n.3w(\'1m\',15.g[15.3v].13)});h(1n.2d(\'1m\')){3u("2c(1n.2d(\'1m\'));",3t)}});m 1j;c 2c(2a){m g=b.1l();h(g.v>1){2b(i=0;i<g.v;i++){h(g[i].13==2a){h(i==1j){1k}1j=i;b.29(i)}}}}',36,291,'||||||||||jw|player|function||||tracks|if||submenu|settings|jwplayer|var|svg|on|audioTracks|div|aria|||attr|length|icon||100|https||||name||event||||active|false|true|captions|tt|dd|position|adb|com|jw8|current_audio|return|getAudioTracks|default_audio|localStorage|removeClass|expanded|checked|addButton|hide|rewind|seek|192|org|url|height|width|css|tr||||vvplay|vvad|lulustream|file|setCurrentAudioTrack|audio_name|for|audio_set|getItem|open|controls|text|time|insertAfter|detach|ff00|button|getPosition|sec|160z|32s32|241V96c0|440|path|512|320|viewBox|class|2000|w3|www||http||xmlns||ff11|vplayer|appendTo|play|no|top|margin|zIndex|openIframeOverlay|lrp9b43lb4fb|file_code|op|pause|empty|log|console|track_name|data|video_ad|doPlay|prevt|logo|srt|theme|300|setTimeout|currentTrack|setItem|audioTrackChanged|dualSound|else|addClass|quality|hasClass|playbackRates|toggleClass|Track|Audio|dualy|images|mousedown|buttons|topbar|countdown|append|slider|Rewind|32V271l11|96V416c0|32S0|160L64|29s||29V96c0|4s18|6c9|M267|Forward|32V271l|32s|32V416c0|160L256|4l192|67s24|83|416V96C0|428|4S0|M52|remove|this|click|scrolling|frameborder|src|prop|50px|1000001|iframe|center|align|rgba|background|1000000|left|absolute|upload_srt|setCurrentCaptions|id|track|test|RegExp|captionsChanged|ready|set_audio_track|html|fviews|referer|embed|hash|view|dl|get||undefined|cRAds|window|over_player_msg|show|complete|slow|fadeIn|video_ad_fadein|cache|Cache||Content|headers|ajaxSetup|lastt|v2done|tott|vastdone2|vastdone1|bar|control|link|aboutlink|LuluStream|abouttext|HD|4340|qualityLabels|fontOpacity|backgroundOpacity|Tahoma|fontFamily|303030|backgroundColor|FFFFFF|color||userFontScale|kind|Upload|label|androidhls||none|preload|1146|duration|uniform|stretching|skin|jpg|lrp9b43lb4fb_xt|lulucdn|img|image|sp|18761758|28800|1779366890|A|OUiz6R5JcqiQ_I7nrdYe0|TYo5euZUUZ9_uUmvHO4|m3u8|master|lrp9b43lb4fb_h|03752|03|hls2|tnmr|wkw3dwshigvf|sources||setup'.split('|')))`
	ja := goja.New()

	value, err := ja.RunString(js)
	if err != nil {
		return
	}

	u, err := url.Parse(l.Source)
	if err != nil {
		return
	}

	// Create a local copy of headers to avoid modifying the global map
	reqHeaders := make(map[string]string)
	for k, v := range header {
		reqHeaders[k] = v
	}
	reqHeaders[`Referer`] = `https://` + u.Host + `/`
	reqHeaders[`Origin`] = `https://` + u.Host

	res, err := client.R().SetHeaders(reqHeaders).Get(l.Source)
	if err != nil {
		return
	}

	scrabbledFunSubmatch := regexp.MustCompile(`<script type='text/javascript'>eval([^<]*)`).FindStringSubmatch(res.String())
	if len(scrabbledFunSubmatch) < 2 {
		err = fmt.Errorf(`len(scrabbledFunSubmatch) < 2 `)
		return
	}

	var URL string
	if urlSubmatch := regexp.MustCompile(`file:"([^"]+\.m3u8[^"]*)"`).FindStringSubmatch(value.String()); len(urlSubmatch) < 2 {
		err = fmt.Errorf("len(urlSubmatch) <=2")
		return
	} else {
		URL = urlSubmatch[1]
	}

	var Name string

	nameRe := regexp.MustCompile(`<h1\s+class="h5">([^<]+)<`)
	if nameSubmatch := nameRe.FindStringSubmatch(res.String()); len(nameSubmatch) < 2 {
		err = fmt.Errorf("len(nameSubmatch) <=2")
		return
	} else {
		Name = nameSubmatch[1]
	}
	Name = html.UnescapeString(Name)

	cr = ext.ContentResource{URL: URL, Name: Name}
	return
}

func (l *Lulustream) Download(cr ext.ContentResource) (err error) {
	// Use the page URL as the Referer, as many sites validate this
	referer := l.Source

	// Construct command with --referer flag instead of just a header
	cmd := exec.Command("yt-dlp.exe", cr.URL, "-o", cr.Name, "--referer", referer)

	var arg []string
	for k, v := range header {
		// Skip Referer here as we use the dedicated --referer flag
		if k == "Referer" {
			continue
		}
		arg = append(arg, "--add-header", fmt.Sprintf("%s: %s", k, v))
	}
	cmd.Args = append(cmd.Args, arg...)

	fmt.Println(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	return
}
