package main

var apps = map[string]interface{}{
	"apps": map[string]interface{}{
		"1C-Bitrix": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"Set-Cookie":    "BITRIX_",
				"X-Powered-CMS": "Bitrix Site Manager",
			},
			"html":    "(?:<link[^>]+components/bitrix|(?:src|href)=\"/bitrix/(?:js|templates))",
			"icon":    "1C-Bitrix.png",
			"implies": "PHP",
			"script":  "1c-bitrix",
			"website": "http://www.1c-bitrix.ru",
		},
		"91App": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "91app.png",
			"script":  "https\\:\\/\\/track\\.91app\\.io\\/track\\.js\\?",
			"website": "https://www.91app.com/",
		},
		"3dCart": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"cookies": map[string]interface{}{
				"3dvisit": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "3DCART",
			},
			"icon":    "3dCart.png",
			"script":  "(?:twlh(?:track)?\\.asp|3d_upsell\\.js)",
			"website": "http://www.3dcart.com",
		},
		"A-Frame": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"html":    "<a-scene[^<>]*>",
			"icon":    "A-Frame.svg",
			"implies": "three.js",
			"js": map[string]interface{}{
				"AFRAME.version": "^(.+)$\\;version:\\1",
			},
			"script":  "/?([\\d.]+)?/aframe(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://aframe.io",
		},
		"AD EBiS": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": []interface{}{
				"<!-- EBiS contents tag",
				"<!--EBiS tag",
				"<!-- Tag EBiS",
				"<!-- EBiS common tag",
			},
			"icon":    "ebis.png",
			"website": "http://www.ebis.ne.jp",
		},
		"AOLserver": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "AOLserver/?([\\d.]+)?\\;version:\\1",
			},
			"icon":    "AOLserver.png",
			"website": "http://aolserver.com",
		},
		"AT Internet Analyzer": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "AT Internet.png",
			"js": map[string]interface{}{
				"ATInternet": "",
				"xtsite":     "",
			},
			"website": "http://atinternet.com/en",
		},
		"AT Internet XiTi": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "AT Internet.png",
			"js": map[string]interface{}{
				"xt_click": "",
			},
			"script":  "xiti\\.com/hit\\.xiti",
			"website": "http://atinternet.com/en",
		},
		"AWStats": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "AWStats.png",
			"implies": "Perl",
			"meta": map[string]interface{}{
				"generator": "AWStats ([\\d.]+(?: \\(build [\\d.]+\\))?)\\;version:\\1",
			},
			"website": "http://awstats.sourceforge.net",
		},
		"AMP": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html":    "<html[^>]* (?:amp|⚡)[^-]",
			"icon":    "Accelerated-Mobile-Pages.svg",
			"website": "https://www.amp.dev",
		},
		"AMP Plugin": map[string]interface{}{
			"cats": []interface{}{
				1,
				5,
			},
			"icon": "Accelerated-Mobile-Pages.svg",
			"implies": []interface{}{
				"WordPress",
			},
			"meta": map[string]interface{}{
				"generator": "^AMP Plugin v(\\d+\\.\\d+.*)$\\;version:\\1",
			},
			"website": "https://amp-wp.org",
		},
		"Azure": map[string]interface{}{
			"cats": []interface{}{
				62,
			},
			"headers": map[string]interface{}{
				"azure-regionname": "",
				"azure-sitename":   "",
				"azure-slotname":   "",
				"azure-version":    "",
			},
			"cookies": map[string]interface{}{
				"ARRAffinity": "",
				"TiPMix":      "",
			},
			"icon":    "azure.svg",
			"website": "https://azure.microsoft.com",
		},
		"Azure CDN": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"server":     "^(?:ECAcc|ECS|ECD)",
				"X-EC-Debug": "",
			},
			"icon":    "azure.svg",
			"website": "https://azure.microsoft.com/en-us/services/cdn/",
		},
		"Acquia Cloud": map[string]interface{}{
			"cats": []interface{}{
				62,
			},
			"headers": map[string]interface{}{
				"X-AH-Environment": "^\\w+$",
			},
			"icon": "acquia-cloud.png",
			"implies": []interface{}{
				"Drupal\\;confidence:95",
				"Apache",
				"Percona",
				"Amazon EC2",
			},
			"website": "https://www.acquia.com/",
		},
		"Act-On": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"icon": "ActOn.png",
			"js": map[string]interface{}{
				"ActOn": "",
			},
			"website": "http://act-on.com",
		},
		"AdInfinity": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon":    "AdInfinity.png",
			"script":  "adinfinity\\.com\\.au",
			"website": "http://adinfinity.com.au",
		},
		"AdRiver": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": "(?:<embed[^>]+(?:src=\"https?://mh\\d?\\.adriver\\.ru/|flashvars=\"[^\"]*(?:http:%3A//(?:ad|mh\\d?)\\.adriver\\.ru/|adriver_banner))|<(?:(?:iframe|img)[^>]+src|a[^>]+href)=\"https?://ad\\.adriver\\.ru/)",
			"icon": "AdRiver.png",
			"js": map[string]interface{}{
				"adriver": "",
			},
			"script":  "(?:adriver\\.core\\.\\d\\.js|https?://(?:content|ad|masterh\\d)\\.adriver\\.ru/)",
			"website": "http://adriver.ru",
		},
		"AdRoll": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "AdRoll.svg",
			"js": map[string]interface{}{
				"adroll_adv_id": "",
				"adroll_pix_id": "",
			},
			"script":  "(?:a|s)\\.adroll\\.com",
			"website": "http://adroll.com",
		},
		"Adcash": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Adcash.svg",
			"js": map[string]interface{}{
				"SuLoaded":       "",
				"SuUrl":          "",
				"ac_bgclick_URL": "",
				"ct_nOpp":        "",
				"ct_nSuUrl":      "",
				"ct_siteunder":   "",
				"ct_tag":         "",
			},
			"script":  "^[^\\/]*//(?:[^\\/]+\\.)?adcash\\.com/(?:script|ad)/",
			"url":     "^https?://(?:[^\\/]+\\.)?adcash\\.com/script/pop_",
			"website": "http://adcash.com",
		},
		"AddShoppers": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "AddShoppers.png",
			"script":  "cdn\\.shop\\.pe/widget/",
			"website": "http://www.addshoppers.com",
		},
		"AddThis": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon": "AddThis.svg",
			"js": map[string]interface{}{
				"addthis": "",
			},
			"script":  "addthis\\.com/js/",
			"website": "http://www.addthis.com",
		},
		"AddToAny": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon": "AddToAny.png",
			"js": map[string]interface{}{
				"a2apage_init": "",
			},
			"script":  "addtoany\\.com/menu/page\\.js",
			"website": "http://www.addtoany.com",
		},
		"Adminer": map[string]interface{}{
			"cats": []interface{}{
				3,
			},
			"html": []interface{}{
				"Adminer</a> <span class=\"version\">([\\d.]+)</span>\\;version:\\1",
				"onclick=\"bodyClick\\(event\\);\" onload=\"verifyVersion\\('([\\d.]+)'\\);\">\\;version:\\1",
			},
			"icon":    "adminer.png",
			"implies": "PHP",
			"website": "http://www.adminer.org",
		},
		"Adnegah": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"headers": map[string]interface{}{
				"X-Advertising-By": "adnegah\\.net",
			},
			"html":    "<iframe [^>]*src=\"[^\"]+adnegah\\.net",
			"icon":    "adnegah.png",
			"script":  "[^a-z]adnegah.*\\.js$",
			"website": "https://Adnegah.net",
		},
		"Adobe ColdFusion": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"Cookie": "CFTOKEN=",
			},
			"html":    "<!-- START headerTags\\.cfm",
			"icon":    "Adobe ColdFusion.svg",
			"implies": "CFML",
			"js": map[string]interface{}{
				"_cfEmails": "",
			},
			"script":  "/cfajax/",
			"url":     "\\.cfm(?:$|\\?)",
			"website": "http://adobe.com/products/coldfusion-family.html",
		},
		"Adobe DTM": map[string]interface{}{
			"cats": []interface{}{
				42,
			},
			"script":  "//assets.adobedtm.com/",
			"icon":    "adobedmt.png",
			"website": "https://marketing.adobe.com/resources/help/en_US/dtm/c_overview.html",
		},
		"Adobe Experience Manager": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": []interface{}{
				"<div class=\"[^\"]*parbase",
				"<div[^>]+data-component-path=\"[^\"+]jcr:",
				"<div class=\"[^\"]*aem-Grid",
			},
			"icon":    "Adobe Experience Manager.svg",
			"implies": "Java",
			"script": []interface{}{
				"/etc/designs/",
				"/etc/clientlibs/",
				"/etc.clientlibs/",
			},
			"website": "https://www.adobe.com/marketing/experience-manager.html",
		},
		"Adobe GoLive": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"icon": "Adobe GoLive.png",
			"meta": map[string]interface{}{
				"generator": "Adobe GoLive(?:\\s([\\d.]+))?\\;version:\\1",
			},
			"website": "http://www.adobe.com/products/golive",
		},
		"Adobe Muse": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"icon": "Adobe Muse.svg",
			"meta": map[string]interface{}{
				"generator": "^Muse(?:$| ?/?(\\d[\\d.]+))\\;version:\\1",
			},
			"website": "http://muse.adobe.com",
		},
		"Adobe RoboHelp": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"icon": "Adobe RoboHelp.svg",
			"js": map[string]interface{}{
				"gbWhLang":  "",
				"gbWhMsg":   "",
				"gbWhProxy": "",
				"gbWhUtil":  "",
				"gbWhVer":   "",
			},
			"meta": map[string]interface{}{
				"generator": "^Adobe RoboHelp(?: ([\\d]+))?\\;version:\\1",
			},
			"script":  "(?:wh(?:utils|ver|proxy|lang|topic|msg)|ehlpdhtm)\\.js",
			"website": "http://adobe.com/products/robohelp.html",
		},
		"ADPLAN": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "ADPLAN.png",
			"script": []interface{}{
				"^https?://[^.]+\\.adplan7\\.com/\\;version:7",
				"^https?://(?!o\\.)\\w+\\.advg\\.jp/",
			},
			"website": "https://www.adplan7.com/",
		},
		"Advanced Web Stats": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html":    "aws\\.src = [^<]+caphyon-analytics",
			"icon":    "Advanced Web Stats.png",
			"implies": "Java",
			"website": "http://www.advancedwebstats.com",
		},
		"Advert Stream": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Advert Stream.png",
			"js": map[string]interface{}{
				"advst_is_above_the_fold": "",
			},
			"script":  "(?:ad\\.advertstream\\.com|adxcore\\.com)",
			"website": "http://www.advertstream.com",
		},
		"Adyen": map[string]interface{}{
			"cats": []interface{}{
				41,
			},
			"icon": "Adyen.svg",
			"js": map[string]interface{}{
				"adyen.encrypt.version": "^(.+)$\\;version:\\1",
			},
			"website": "https://www.adyen.com",
		},
		"Adzerk": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": "<iframe [^>]*src=\"[^\"]+adzerk\\.net",
			"icon": "Adzerk.png",
			"js": map[string]interface{}{
				"ados":        "",
				"adosResults": "",
			},
			"script":  "adzerk\\.net/ados\\.js",
			"website": "http://adzerk.com",
		},
		"Aegea": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^E2 Aegea v(\\d+)$\\;version:\\1",
			},
			"icon": "Aegea.png",
			"implies": []interface{}{
				"PHP",
				"jQuery",
			},
			"website": "http://blogengine.ru",
		},
		"Afosto": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Afosto SaaS BV",
			},
			"icon":    "Afosto.svg",
			"website": "http://afosto.com",
		},
		"AfterBuy": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": []interface{}{
				"<dd>This OnlineStore is brought to you by ViA-Online GmbH Afterbuy\\. Information and contribution at https://www\\.afterbuy\\.de</dd>",
			},
			"icon":    "after-buy.png",
			"script":  "shop-static\\.afterbuy\\.de",
			"website": "http://www.afterbuy.de",
		},
		"Ahoy": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"js": map[string]interface{}{
				"ahoy": "",
			},
			"cookies": map[string]interface{}{
				"ahoy_track":   "",
				"ahoy_visit":   "",
				"ahoy_visitor": "",
			},
			"website": "https://github.com/ankane/ahoy",
		},
		"Aircall": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon":    "aircall.png",
			"script":  "^https?://cdn\\.aircall\\.io/",
			"website": "http://aircall.io",
		},
		"Airee": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server": "^Airee",
			},
			"icon":    "Airee.png",
			"website": "http://xn--80aqc2a.xn--p1ai",
		},
		"Akamai": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"X-Akamai-Transformed": "",
			},
			"icon":    "akamai.svg",
			"website": "http://akamai.com",
		},
		"Akaunting": map[string]interface{}{
			"cats": []interface{}{
				55,
			},
			"headers": map[string]interface{}{
				"X-Akaunting": "^Free Accounting Software$",
			},
			"html": []interface{}{
				"<link[^>]+akaunting-green\\.css",
				"Powered By Akaunting: <a [^>]*href=\"https?://(?:www\\.)?akaunting\\.com[^>]+>",
			},
			"icon":    "akaunting.svg",
			"implies": "Laravel",
			"website": "https://akaunting.com",
		},
		"Akka HTTP": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"Server": "akka-http(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "akka-http.png",
			"website": "http://akka.io",
		},
		"Algolia Realtime Search": map[string]interface{}{
			"cats": []interface{}{
				29,
			},
			"icon": "Algolia Realtime Search.svg",
			"js": map[string]interface{}{
				"AlgoliaSearch":         "",
				"algoliasearch.version": "^(.+)$\\;version:\\1",
			},
			"website": "http://www.algolia.com",
		},
		"All in One SEO Pack": map[string]interface{}{
			"cats": []interface{}{
				54,
			},
			"html":    "<!-- All in One SEO Pack ([\\d.]+) \\;version:\\1",
			"icon":    "all-in-One-SEO-Pack.png",
			"implies": "WordPress",
			"website": "https://wordpress.org/plugins/all-in-one-seo-pack/",
		},
		"Allegro RomPager": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Allegro-Software-RomPager(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Allegro RomPager.png",
			"website": "http://allegrosoft.com/embedded-web-server-s2",
		},
		"AlloyUI": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "AlloyUI.png",
			"implies": []interface{}{
				"Bootstrap",
				"YUI",
			},
			"js": map[string]interface{}{
				"AUI": "",
			},
			"script":  "^https?://cdn\\.alloyui\\.com/",
			"website": "http://www.alloyui.com",
		},
		"Amaya": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"icon": "Amaya.png",
			"meta": map[string]interface{}{
				"generator": "Amaya(?: V?([\\d.]+[a-z]))?\\;version:\\1",
			},
			"website": "http://www.w3.org/Amaya",
		},
		"Amazon Cloudfront": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Via":         "\\(CloudFront\\)$",
				"X-Amz-Cf-Id": "",
			},
			"icon":    "Amazon-Cloudfront.svg",
			"implies": "Amazon Web Services",
			"website": "http://aws.amazon.com/cloudfront/",
		},
		"Amazon EC2": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "\\(Amazon\\)",
			},
			"icon":    "aws-ec2.svg",
			"implies": "Amazon Web Services",
			"website": "http://aws.amazon.com/ec2/",
		},
		"Amazon Web Services": map[string]interface{}{
			"cats": []interface{}{
				62,
			},
			"icon":    "aws.svg",
			"website": "https://aws.amazon.com/",
		},
		"Amazon ECS": map[string]interface{}{
			"cats": []interface{}{
				63,
			},
			"headers": map[string]interface{}{
				"Server": "^ECS",
			},
			"icon": "aws.svg",
			"implies": []interface{}{
				"Amazon Web Services",
				"Docker",
			},
			"website": "https://aws.amazon.com/elasticloadbalancing/",
		},
		"Amazon ELB": map[string]interface{}{
			"cats": []interface{}{
				65,
			},
			"cookies": map[string]interface{}{
				"AWSELB": "",
			},
			"icon":    "aws-elb.png",
			"implies": "Amazon Web Services",
			"website": "https://aws.amazon.com/elasticloadbalancing/",
		},
		"Amazon S3": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"headers": map[string]interface{}{
				"Server": "^AmazonS3$",
			},
			"icon":    "aws-s3.svg",
			"implies": "Amazon Web Services",
			"website": "http://aws.amazon.com/s3/",
		},
		"Amber": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Amber$",
			},
			"icon":    "amber.png",
			"website": "https://amberframework.org",
		},
		"Ametys": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Ametys.png",
			"implies": "Java",
			"meta": map[string]interface{}{
				"generator": "(?:Ametys|Anyware Technologies)",
			},
			"script":  "ametys\\.js",
			"website": "http://ametys.org",
		},
		"Amiro.CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Amiro.CMS.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Amiro",
			},
			"website": "http://amirocms.com",
		},
		"Amplitude": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "amplitude.png",
			"script": []interface{}{
				"cdn\\.amplitude\\.com",
			},
			"website": "https://amplitude.com/",
		},
		"Analysys Ark": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"js": map[string]interface{}{
				"AnalysysAgent": "",
			},
			"cookies": map[string]interface{}{
				"ARK_ID": "",
			},
			"icon":    "Analysys Ark.svg",
			"script":  "AnalysysFangzhou_JS_SDK\\.min\\.js\\?v=([\\d.]+)\\;version:\\1",
			"website": "https://ark.analysys.cn",
		},
		"Anetwork": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon":    "Anetwork.png",
			"script":  "static-cdn\\.anetwork\\.ir/",
			"website": "https://www.anetwork.ir",
		},
		"Angular": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"excludes": []interface{}{
				"AngularDart",
				"AngularJS",
			},
			"html": "<[^>]+ ng-version=\"([\\d.]+)\"\\;version:\\1",
			"icon": "Angular.svg",
			"js": map[string]interface{}{
				"ng.coreTokens": "",
				"ng.probe":      "",
			},
			"website": "https://angular.io",
		},
		"Angular Material": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"icon":    "AngularJS.svg",
			"implies": "AngularJS",
			"js": map[string]interface{}{
				"ngMaterial": "",
			},
			"script":  "/([\\d.rc-]+)?/angular-material(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://material.angularjs.org",
		},
		"AngularDart": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"excludes": []interface{}{
				"Angular",
				"AngularJS",
			},
			"icon":    "AngularDart.svg",
			"implies": "Dart",
			"js": map[string]interface{}{
				"ngTestabilityRegistries": "",
			},
			"website": "https://webdev.dartlang.org/angular/",
		},
		"AngularJS": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"excludes": []interface{}{
				"Angular",
				"AngularDart",
			},
			"icon": "AngularJS.svg",
			"html": []interface{}{
				"<(?:div|html)[^>]+ng-app=",
				"<ng-app",
			},
			"js": map[string]interface{}{
				"angular":              "",
				"angular.version.full": "^(.+)$\\;version:\\1",
			},
			"script": []interface{}{
				"angular[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"/([\\d.]+(?:-?rc[.\\d]*)*)/angular(?:\\.min)?\\.js\\;version:\\1",
				"angular.*\\.js",
			},
			"website": "https://angularjs.org",
		},
		"Ant Design": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html": []interface{}{
				"<[^>]*class=\"ant-(?:btn|col|row|layout|breadcrumb|menu|pagination|steps|select|cascader|checkbox|calendar|form|input-number|input|mention|rate|radio|slider|switch|tree-select|time-picker|transfer|upload|avatar|badge|card|carousel|collapse|list|popover|tooltip|table|tabs|tag|timeline|tree|alert|modal|message|notification|progress|popconfirm|spin|anchor|back-top|divider|drawer)",
				"<i class=\"anticon anticon-",
			},
			"icon": "Ant Design.svg",
			"js": map[string]interface{}{
				"antd": "",
			},
			"website": "https://ant.design",
		},
		"Apache": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "(?:Apache(?:$|/([\\d.]+)|[^/-])|(?:^|\\b)HTTPD)\\;version:\\1",
			},
			"icon":    "Apache.svg",
			"website": "http://apache.org",
		},
		"Apache HBase": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"html":    "<style[^>]+static/hbase",
			"icon":    "Apache HBase.png",
			"implies": "Java",
			"website": "http://hbase.apache.org",
		},
		"Apache Hadoop": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"html":    "<style[^>]+static/hadoop",
			"icon":    "Apache Hadoop.svg",
			"website": "http://hadoop.apache.org",
		},
		"Apache JSPWiki": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"html":    "<html[^>]* xmlns:jspwiki=",
			"icon":    "Apache JSPWiki.png",
			"implies": "Apache Tomcat",
			"script":  "jspwiki",
			"url":     "wiki\\.jsp",
			"website": "http://jspwiki.org",
		},
		"Apache Tomcat": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server":       "^Apache-Coyote(?:/([\\d.]+))?\\;version:\\1",
				"X-Powered-By": "\\bTomcat\\b(?:-([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Apache Tomcat.svg",
			"implies": "Java",
			"website": "http://tomcat.apache.org",
		},
		"Apache Traffic Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "ATS/?([\\d.]+)?\\;version:\\1",
			},
			"icon":    "Apache Traffic Server.png",
			"website": "http://trafficserver.apache.org/",
		},
		"Apache Wicket": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"icon":    "Apache Wicket.svg",
			"implies": "Java",
			"js": map[string]interface{}{
				"Wicket": "",
			},
			"website": "http://wicket.apache.org",
		},
		"ApexPages": map[string]interface{}{
			"cats": []interface{}{
				51,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Salesforce\\.com ApexPages",
			},
			"icon":    "ApexPages.png",
			"implies": "Salesforce",
			"website": "https://developer.salesforce.com/docs/atlas.en-us.apexcode.meta/apexcode/apex_intro.htm",
		},
		"Apigee": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"script":  "/profiles/apigee",
			"html":    "<script>[^>]{0,50}script src=[^>]/profiles/apigee",
			"icon":    "apigee.svg",
			"website": "https://cloud.google.com/apigee/",
		},
		"Apostrophe CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<[^>]+data-apos-refreshable[^>]",
			"icon":    "apostrophecms.svg",
			"implies": "Node.js",
			"website": "http://apostrophecms.org",
		},
		"AppNexus": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html":    "<(?:iframe|img)[^>]+adnxs\\.(?:net|com)",
			"icon":    "AppNexus.svg",
			"script":  "adnxs\\.(?:net|com)",
			"website": "http://appnexus.com",
		},
		"Appcues": map[string]interface{}{
			"cats": []interface{}{
				58,
			},
			"icon":    "Appcues.svg",
			"script":  "fast\\.appcues.com*\\.js",
			"website": "https://appcues.com",
		},
		"Arastta": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"excludes": "OpenCart",
			"headers": map[string]interface{}{
				"Arastta":   "^(.+)$\\;version:\\1",
				"X-Arastta": "",
			},
			"html":    "Powered by <a [^>]*href=\"https?://(?:www\\.)?arastta\\.org[^>]+>Arastta",
			"icon":    "Arastta.svg",
			"implies": "PHP",
			"script":  "arastta\\.js",
			"website": "http://arastta.org",
		},
		"ArcGIS API for JavaScript": map[string]interface{}{
			"cats": []interface{}{
				35,
			},
			"icon": "arcgis_icon.png",
			"script": []interface{}{
				"js\\.arcgis\\.com",
				"basemaps\\.arcgis\\.com",
			},
			"website": "https://developers.arcgis.com/javascript/",
		},
		"Artifactory": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"html": []interface{}{
				"<span class=\"version\">Artifactory(?: Pro)?(?: Power Pack)?(?: ([\\d.]+))?\\;version:\\1",
			},
			"icon": "Artifactory.svg",
			"js": map[string]interface{}{
				"ArtifactoryUpdates": "",
			},
			"script": []interface{}{
				"wicket/resource/org\\.artifactory\\.",
			},
			"website": "http://jfrog.com/open-source/#os-arti",
		},
		"Artifactory Web Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Artifactory(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon": "Artifactory.svg",
			"implies": []interface{}{
				"Artifactory",
			},
			"website": "http://jfrog.com/open-source/#os-arti",
		},
		"ArvanCloud": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"AR-PoweredBy": "Arvan Cloud \\(arvancloud\\.com\\)",
			},
			"icon": "ArvanCloud.png",
			"js": map[string]interface{}{
				"ArvanCloud": "",
			},
			"website": "http://www.ArvanCloud.com",
		},
		"AsciiDoc": map[string]interface{}{
			"cats": []interface{}{
				1,
				20,
				27,
			},
			"icon": "AsciiDoc.png",
			"js": map[string]interface{}{
				"asciidoc": "",
			},
			"meta": map[string]interface{}{
				"generator": "^AsciiDoc ([\\d.]+)\\;version:\\1",
			},
			"website": "http://www.methods.co.nz/asciidoc",
		},
		"Asciinema": map[string]interface{}{
			"cats": []interface{}{
				14,
			},
			"html": "<asciinema-player",
			"icon": "asciinema.png",
			"js": map[string]interface{}{
				"asciinema": "",
			},
			"script":  "asciinema\\.org/",
			"website": "https://asciinema.org/",
		},
		"Atlassian Bitbucket": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"html":    "<li>Atlassian Bitbucket <span title=\"[a-z0-9]+\" id=\"product-version\" data-commitid=\"[a-z0-9]+\" data-system-build-number=\"[a-z0-9]+\"> v([\\d.]+)<\\;version:\\1",
			"icon":    "Atlassian Bitbucket.svg",
			"implies": "Python",
			"js": map[string]interface{}{
				"bitbucket": "",
			},
			"meta": map[string]interface{}{
				"application-name": "Bitbucket",
			},
			"website": "http://www.atlassian.com/software/bitbucket/overview/",
		},
		"Atlassian Confluence": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"headers": map[string]interface{}{
				"X-Confluence-Request-Time": "",
			},
			"html":    "Powered by <a href=[^>]+atlassian\\.com/software/confluence(?:[^>]+>Atlassian Confluence</a> ([\\d.]+))?\\;version:\\1",
			"icon":    "Atlassian Confluence.svg",
			"implies": "Java",
			"meta": map[string]interface{}{
				"confluence-request-time": "",
			},
			"website": "http://www.atlassian.com/software/confluence/overview/team-collaboration-software",
		},
		"Atlassian FishEye": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"cookies": map[string]interface{}{
				"FESESSIONID": "",
			},
			"html":    "<title>(?:Log in to )?FishEye (?:and Crucible )?([\\d.]+)?</title>\\;version:\\1",
			"icon":    "Atlassian FishEye.svg",
			"website": "http://www.atlassian.com/software/fisheye/overview/",
		},
		"Atlassian Jira": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"html":    "Powered by\\s+<a href=[^>]+atlassian\\.com/(?:software/jira|jira-bug-tracking/)[^>]+>Atlassian\\s+JIRA(?:[^v]*v(?:ersion: )?(\\d+\\.\\d+(?:\\.\\d+)?))?\\;version:\\1",
			"icon":    "Atlassian Jira.svg",
			"implies": "Java",
			"js": map[string]interface{}{
				"jira": "",
			},
			"meta": map[string]interface{}{
				"ajs-version-number": "^(.+)$\\;version:\\1",
				"application-name":   "JIRA",
			},
			"website": "http://www.atlassian.com/software/jira/overview/",
		},
		"Atlassian Jira Issue Collector": map[string]interface{}{
			"cats": []interface{}{
				13,
				47,
			},
			"icon": "Atlassian Jira.svg",
			"script": []interface{}{
				"jira-issue-collector-plugin",
				"atlassian\\.jira\\.collector\\.plugin",
			},
			"website": "http://www.atlassian.com/software/jira/overview/",
		},
		"Aurelia": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html": []interface{}{
				"<[^>]+aurelia-app=[^>]",
				"<[^>]+data-main=[^>]aurelia-bootstrapper",
				"<[^>]+au-target-id=[^>]\\d",
			},
			"icon": "Aurelia.svg",
			"script": []interface{}{
				"aurelia(?:\\.min)?\\.js",
			},
			"website": "http://aurelia.io",
		},
		"Avangate": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<link[^>]* href=\"^https?://edge\\.avangate\\.net/",
			"icon": "Avangate.svg",
			"js": map[string]interface{}{
				"__avng8_": "",
				"avng8_":   "",
			},
			"script":  "^https?://edge\\.avangate\\.net/",
			"website": "http://avangate.com",
		},
		"Awesomplete": map[string]interface{}{
			"cats": []interface{}{
				29,
			},
			"html": "<link[^>]+href=\"[^>]*awesomplete(?:\\.min)?\\.css",
			"js": map[string]interface{}{
				"awesomplete": "",
			},
			"script":  "/awesomplete\\.js(?:$|\\?)",
			"website": "https://leaverou.github.io/awesomplete/",
		},
		"BEM": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html":    "<[^>]+data-bem",
			"icon":    "BEM.png",
			"website": "http://en.bem.info",
		},
		"BIGACE": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "(?:Powered by <a href=\"[^>]+BIGACE|<!--\\s+Site is running BIGACE)",
			"icon":    "BIGACE.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "BIGACE ([\\d.]+)\\;version:\\1",
			},
			"website": "http://bigace.de",
		},
		"Bablic": map[string]interface{}{
			"cats": []interface{}{
				3,
				9,
			},
			"icon": "bablic.png",
			"js": map[string]interface{}{
				"bablic": "",
			},
			"website": "https://www.bablic.com/",
		},
		"Backbone.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon":    "Backbone.js.png",
			"implies": "Underscore.js",
			"js": map[string]interface{}{
				"Backbone":         "",
				"Backbone.VERSION": "^(.+)$\\;version:\\1",
			},
			"script":  "backbone.*\\.js",
			"website": "http://backbonejs.org",
		},
		"Backdrop": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"excludes": "Drupal",
			"icon":     "Backdrop.png",
			"implies":  "PHP",
			"js": map[string]interface{}{
				"Backdrop": "",
			},
			"meta": map[string]interface{}{
				"generator": "Backdrop CMS(?: (\\d))?\\;version:\\1",
			},
			"website": "http://backdropcms.org",
		},
		"Backpack": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"cookies": map[string]interface{}{
				"backpack_session=": "",
			},
			"icon":    "Backpack.png",
			"implies": "Laravel",
			"website": "https://backpackforlaravel.com",
		},
		"Backtory": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Backtory",
			},
			"icon":    "Backtory.svg",
			"website": "https://backtory.com",
		},
		"Banshee": map[string]interface{}{
			"cats": []interface{}{
				1,
				18,
			},
			"html":    "Built upon the <a href=\"[^>]+banshee-php\\.org/\">[a-z]+</a>(?:v([\\d.]+))?\\;version:\\1",
			"icon":    "Banshee.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Banshee PHP",
			},
			"website": "http://www.banshee-php.org",
		},
		"BaseHTTP": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "BaseHTTP\\/?([\\d\\.]+)?\\;version:\\1",
			},
			"icon":    "BaseHTTP.png",
			"implies": "Python",
			"website": "http://docs.python.org/2/library/basehttpserver.html",
		},
		"BigBangShop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-SERVER": "BIGBANGSHOP",
			},
			"icon":    "bigbangshop.svg",
			"website": "https://www.bigbangshop.com.br",
		},
		"BigDump": map[string]interface{}{
			"cats": []interface{}{
				3,
			},
			"html": "<!-- <h1>BigDump: Staggered MySQL Dump Importer ver\\. ([\\d.b]+)\\;version:\\1",
			"implies": []interface{}{
				"MySQL",
				"PHP",
			},
			"website": "http://www.ozerov.de/bigdump.php",
		},
		"Bigcommerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<link href=[^>]+cdn\\d+\\.bigcommerce\\.com/",
			"icon":    "Bigcommerce.png",
			"script":  "cdn\\d+\\.bigcommerce\\.com/",
			"url":     "mybigcommerce\\.com",
			"website": "http://www.bigcommerce.com",
		},
		"Bigware": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"bigWAdminID": "",
				"bigwareCsid": "",
			},
			"html":    "(?:Diese <a href=[^>]+bigware\\.de|<a href=[^>]+/main_bigware_\\d+\\.php)",
			"icon":    "Bigware.png",
			"implies": "PHP",
			"url":     "(?:\\?|&)bigWAdminID=",
			"website": "http://bigware.de",
		},
		"BittAds": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "BittAds.png",
			"js": map[string]interface{}{
				"bitt": "",
			},
			"script":  "bittads\\.com/js/bitt\\.js$",
			"website": "http://bittads.com",
		},
		"Bizweb": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "bizweb.png",
			"js": map[string]interface{}{
				"Bizweb": "",
			},
			"website": "https://www.bizweb.vn",
		},
		"Blade": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "blade-([\\w.]+)?\\;version:\\1",
			},
			"icon":    "Blade.png",
			"implies": "Java",
			"website": "https://lets-blade.com",
		},
		"Blesta": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"blesta_sid": "",
			},
			"icon":    "Blesta.png",
			"website": "http://www.blesta.com",
		},
		"Blip.tv": map[string]interface{}{
			"cats": []interface{}{
				14,
			},
			"html":    "<(?:param|embed|iframe)[^>]+blip\\.tv/play",
			"icon":    "Blip.tv.png",
			"website": "http://blip.tv",
		},
		"Blogger": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"icon":    "Blogger.png",
			"implies": "Python",
			"meta": map[string]interface{}{
				"generator": "^Blogger$",
			},
			"url":     "^https?://[^/]+\\.blogspot\\.com",
			"website": "http://www.blogger.com",
		},
		"Bluefish": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"icon": "Bluefish.png",
			"meta": map[string]interface{}{
				"generator": "Bluefish(?:\\s([\\d.]+))?\\;version:\\1",
			},
			"website": "http://sourceforge.net/projects/bluefish",
		},
		"Boa": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Boa\\/?([\\d\\.a-z]+)?\\;version:\\1",
			},
			"website": "http://www.boa.org",
		},
		"Boba.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"implies": "Google Analytics",
			"script":  "boba(?:\\.min)?\\.js",
			"website": "http://boba.space150.com",
		},
		"Bold Chat": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon":    "BoldChat.png",
			"script":  "^https?://vmss\\.boldchat\\.com/aid/\\d{18}/bc\\.vms4/vms\\.js",
			"website": "https://www.boldchat.com/",
		},
		"BoldGrid": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"html": []interface{}{
				"<link rel=[\"']stylesheet[\"'] [^>]+boldgrid",
				"<link rel=[\"']stylesheet[\"'] [^>]+post-and-page-builder",
				"<link[^>]+s\\d+\\.boldgrid\\.com",
			},
			"script":  "/wp-content/plugins/post-and-page-builder",
			"implies": "WordPress",
			"website": "https://boldgrid.com",
		},
		"Bolt": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Bolt.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Bolt",
			},
			"website": "http://bolt.cm",
		},
		"Bonfire": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"bf_session": "",
			},
			"html":    "Powered by <a[^>]+href=\"https?://(?:www\\.)?cibonfire\\.com[^>]*>Bonfire v([^<]+)\\;version:\\1",
			"icon":    "Bonfire.png",
			"implies": "CodeIgniter",
			"website": "http://cibonfire.com",
		},
		"Bootstrap": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": []interface{}{
				"<style>/\\*!\\* Bootstrap v(\\d\\.\\d\\.\\d)\\;version:\\1",
				"<link[^>]+?href=[^\"]/css/([\\d.]+)/bootstrap\\.(?:min\\.)?css\\;version:\\1",
				"<link[^>]+?href=\"[^\"]*bootstrap(?:\\.min)?\\.css",
				"<div[^>]+class=\"[^\"]*glyphicon glyphicon-",
			},
			"icon": "Bootstrap.png",
			"js": map[string]interface{}{
				"bootstrap.Alert.VERSION":               "^(.+)$\\;version:\\1",
				"jQuery.fn.tooltip.Constructor.VERSION": "^(.+)$\\;version:\\1",
			},
			"script": []interface{}{
				"twitter\\.github\\.com/bootstrap",
				"bootstrap[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"(?:/([\\d.]+))?(?:/js)?/bootstrap(?:\\.min)?\\.js\\;version:\\1",
			},
			"website": "https://getbootstrap.com",
		},
		"Bootstrap Table": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"html": "<link[^>]+href=\"[^>]*bootstrap-table(?:\\.min)?\\.css",
			"icon": "Bootstrap Table.svg",
			"implies": []interface{}{
				"Bootstrap",
				"jQuery",
			},
			"script":  "bootstrap-table(?:\\.min)?\\.js",
			"website": "http://bootstrap-table.wenzhixin.net.cn/",
		},
		"Bounce Exchange": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"script": "^https?://tag\\.bounceexchange\\.com/",
			"icon":   "Bounce Exchange.svg",
			"js": map[string]interface{}{
				"bouncex": "",
			},
			"website": "http://www.bounceexchange.com",
		},
		"Braintree": map[string]interface{}{
			"cats": []interface{}{
				41,
			},
			"icon": "Braintree.svg",
			"js": map[string]interface{}{
				"Braintree":         "",
				"Braintree.version": "^(.+)$\\;version:\\1",
			},
			"website": "https://www.braintreepayments.com",
		},
		"Brightspot": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Brightspot$",
			},
			"icon":    "Brightspot.svg",
			"implies": "Java",
			"website": "https://www.brightspot.com",
		},
		"BrowserCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "BrowserCMS.png",
			"implies": "Ruby",
			"meta": map[string]interface{}{
				"generator": "BrowserCMS ([\\d.]+)\\;version:\\1",
			},
			"website": "http://browsercms.org",
		},
		"Bubble": map[string]interface{}{
			"cats": []interface{}{
				1,
				3,
				18,
				22,
			},
			"icon":    "bubble.png",
			"implies": "Node.js",
			"js": map[string]interface{}{
				"appquery": "",
			},
			"website": "http://bubble.is",
		},
		"BugSnag": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "BugSnag.png",
			"js": map[string]interface{}{
				"Bugsnag":       "",
				"bugsnag":       "",
				"bugsnagClient": "",
			},
			"script":  "/bugsnag.*\\.js",
			"website": "http://bugsnag.com",
		},
		"Bugzilla": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"html": []interface{}{
				"href=\"enter_bug\\.cgi\">",
				"<main id=\"bugzilla-body\"",
				"<a href=\"https?://www\\.bugzilla\\.org/docs/([0-9.]+)/[^>]+>Help<\\;version:\\1",
				"<span id=\"information\" class=\"header_addl_info\">version ([\\d.]+)<\\;version:\\1",
			},
			"cookies": map[string]interface{}{
				"Bugzilla_login_request_cookie": "",
			},
			"icon":    "Bugzilla.png",
			"implies": "Perl",
			"js": map[string]interface{}{
				"BUGZILLA": "",
			},
			"meta": map[string]interface{}{
				"generator": "Bugzilla ?([\\d.]+)?\\;version:\\1",
			},
			"website": "http://www.bugzilla.org",
		},
		"Bulma": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "<link[^>]+?href=\"[^\"]+bulma(?:\\.min)?\\.css",
			"icon":    "Bulma.png",
			"website": "http://bulma.io",
		},
		"Burning Board": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html": "<a href=\"[^>]+woltlab\\.com[^<]+<strong>Burning Board",
			"icon": "Burning Board.png",
			"implies": []interface{}{
				"PHP",
				"Woltlab Community Framework",
			},
			"website": "http://www.woltlab.com",
		},
		"Business Catalyst": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<!-- BC_OBNW -->",
			"icon":    "Business Catalyst.png",
			"script":  "CatalystScripts",
			"website": "http://businesscatalyst.com",
		},
		"BuySellAds": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"script": "^https?://s\\d\\.buysellads\\.com/",
			"icon":   "BuySellAds.png",
			"js": map[string]interface{}{
				"_bsa":                   "",
				"_bsaPRO":                "",
				"_bsap":                  "",
				"_bsap_serving_callback": "",
			},
			"website": "http://buysellads.com",
		},
		"CDN77": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server": "^CDN77-Turbo$",
			},
			"icon":    "CDN77.png",
			"website": "https://www.cdn77.com",
		},
		"CFML": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"icon":    "CFML.png",
			"website": "http://adobe.com/products/coldfusion-family.html",
		},
		"CKEditor": map[string]interface{}{
			"cats": []interface{}{
				24,
			},
			"icon": "CKEditor.png",
			"js": map[string]interface{}{
				"CKEDITOR":          "",
				"CKEDITOR.version":  "^(.+)$\\;version:\\1",
				"CKEDITOR_BASEPATH": "",
			},
			"website": "http://ckeditor.com",
		},
		"CMS Made Simple": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"CMSSESSID": "",
			},
			"icon":    "CMS Made Simple.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "CMS Made Simple",
			},
			"website": "http://cmsmadesimple.org",
		},
		"CMSimple": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "CMSimple( [\\d.]+)?\\;version:\\1",
			},
			"website": "http://www.cmsimple.org/en",
		},
		"CPG Dragonfly": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Dragonfly CMS",
			},
			"icon":    "CPG Dragonfly.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "CPG Dragonfly",
			},
			"website": "http://dragonflycms.org",
		},
		"CS Cart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": []interface{}{
				"&nbsp;Powered by (?:<a href=[^>]+cs-cart\\.com|CS-Cart)",
				"\\.cm-noscript[^>]+</style>",
			},
			"icon":    "CS Cart.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"fn_compare_strings": "",
			},
			"website": "http://www.cs-cart.com",
		},
		"CacheFly": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server": "^CFS ",
				"X-CF1":  "",
				"X-CF2":  "",
			},
			"icon":    "CacheFly.png",
			"website": "http://www.cachefly.com",
		},
		"Caddy": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^Caddy$",
			},
			"icon":    "caddy.svg",
			"implies": "Go",
			"website": "http://caddyserver.com",
		},
		"CakePHP": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"cakephp": "",
			},
			"icon":    "CakePHP.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"application-name": "CakePHP",
			},
			"website": "http://cakephp.org",
		},
		"Captch Me": map[string]interface{}{
			"cats": []interface{}{
				16,
				36,
			},
			"icon": "Captch Me.svg",
			"js": map[string]interface{}{
				"Captchme": "",
			},
			"script":  "^https?://api\\.captchme\\.net/",
			"website": "http://captchme.com",
		},
		"Carbon Ads": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": "<[a-z]+ [^>]*id=\"carbonads-container\"",
			"icon": "Carbon Ads.png",
			"js": map[string]interface{}{
				"_carbonads": "",
			},
			"script":  "//(?:engine|srv)\\.carbonads\\.com\\/",
			"website": "http://carbonads.net",
		},
		"Cargo": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<link [^>]+Cargo feed",
			"icon":    "Cargo.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"cargo_title": "",
			},
			"script":  "/cargo\\.",
			"website": "http://cargocollective.com",
		},
		"Catberry.js": map[string]interface{}{
			"cats": []interface{}{
				12,
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Catberry",
			},
			"icon":    "Catberry.js.png",
			"implies": "Node.js",
			"js": map[string]interface{}{
				"catberry":         "",
				"catberry.version": "^(.+)$\\;version:\\1",
			},
			"website": "http://catberry.org",
		},
		"CentOS": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":       "CentOS",
				"X-Powered-By": "CentOS",
			},
			"icon":    "CentOS.png",
			"website": "http://centos.org",
		},
		"Chameleon": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Chameleon.png",
			"implies": []interface{}{
				"Apache",
				"PHP",
			},
			"meta": map[string]interface{}{
				"generator": "chameleon-cms",
			},
			"website": "http://chameleon-system.de",
		},
		"Chamilo": map[string]interface{}{
			"cats": []interface{}{
				21,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Chamilo ([\\d.]+)\\;version:\\1",
			},
			"html":    "\">Chamilo ([\\d.]+)</a>\\;version:\\1",
			"icon":    "Chamilo.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Chamilo ([\\d.]+)\\;version:\\1",
			},
			"website": "http://www.chamilo.org",
		},
		"Chart.js": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "Chart.js.svg",
			"js": map[string]interface{}{
				"Chart":                   "\\;confidence:50",
				"Chart.defaults.doughnut": "",
				"chart.ctx.bezierCurveTo": "",
			},
			"script": []interface{}{
				"/Chart(?:\\.bundle)?(?:\\.min)?\\.js\\;confidence:75",
				"chartjs\\.org/dist/([\\d.]+(?:-[^/]+)?|master|latest)/Chart.*\\.js\\;version:\\1",
				"cdnjs\\.cloudflare\\.com/ajax/libs/Chart\\.js/([\\d.]+(?:-[^/]+)?)/Chart.*\\.js\\;version:\\1",
				"cdn\\.jsdelivr\\.net/(?:npm|gh/chartjs)/chart\\.js@([\\d.]+(?:-[^/]+)?|latest)/dist/Chart.*\\.js\\;version:\\1",
			},
			"website": "https://www.chartjs.org",
		},
		"Chartbeat": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Chartbeat.png",
			"js": map[string]interface{}{
				"_sf_async_config": "",
				"_sf_endpt":        "",
			},
			"script":  "chartbeat\\.js",
			"website": "http://chartbeat.com",
		},
		"Cherokee": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^Cherokee(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Cherokee.png",
			"website": "http://www.cherokee-project.com",
		},
		"CherryPy": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"Server": "CherryPy\\/?([\\d\\.]+)?\\;version:\\1",
			},
			"icon":    "CherryPy.png",
			"implies": "Python",
			"website": "http://www.cherrypy.org",
		},
		"Chevereto": map[string]interface{}{
			"cats": []interface{}{
				7,
			},
			"meta": map[string]interface{}{
				"generator": "^Chevereto ?([0-9.]+)?$\\;version:\\1",
			},
			"script":  "/chevereto\\.js",
			"html":    "Powered by <a href=\"https?://chevereto\\.com\">",
			"icon":    "chevereto.png",
			"implies": "PHP",
			"website": "https://chevereto.com/",
		},
		"Chitika": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Chitika.png",
			"js": map[string]interface{}{
				"ch_client":          "",
				"ch_color_site_link": "",
			},
			"script":  "scripts\\.chitika\\.net/",
			"website": "http://chitika.com",
		},
		"Ckan": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"Access-Control-Allow-Headers": "X-CKAN-API-KEY",
				"Link":                         "<http://ckan\\.org/>; rel=shortlink",
			},
			"icon": "Ckan.png",
			"implies": []interface{}{
				"Python",
				"Solr",
				"Java",
				"PostgreSQL",
			},
			"meta": map[string]interface{}{
				"generator": "^ckan ?([0-9.]+)$\\;version:\\1",
			},
			"website": "http://ckan.org/",
		},
		"Clarity": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": []interface{}{
				"<clr-main-container",
				"<link [^>]*href=\"[^\"]*clr-ui(?:\\.min)?\\.css",
			},
			"js": map[string]interface{}{
				"ClarityIcons": "",
			},
			"script": "clr-angular(?:\\.umd)?(?:\\.min)?\\.js",
			"icon":   "clarity.svg",
			"implies": []interface{}{
				"Angular",
			},
			"website": "https://clarity.design/",
		},
		"ClickHeat": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "ClickHeat.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"clickHeatServer": "",
			},
			"script":  "clickheat.*\\.js",
			"website": "http://www.labsmedia.com/clickheat/index.html",
		},
		"ClickTale": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "ClickTale.png",
			"js": map[string]interface{}{
				"clickTaleStartEventSignal": "",
			},
			"website": "http://www.clicktale.com",
		},
		"Clicky": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Clicky.png",
			"js": map[string]interface{}{
				"clicky": "",
			},
			"script":  "static\\.getclicky\\.com",
			"website": "http://getclicky.com",
		},
		"Clientexec": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "clientexec\\.[^>]*\\s?=\\s?[^>]*;",
			"icon":    "Clientexec.png",
			"website": "http://www.clientexec.com",
		},
		"Clipboard.js": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon":    "Clipboard.js.svg",
			"script":  "clipboard(?:-([\\d.]+))?(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://clipboardjs.com/",
		},
		"CloudCart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "cloudcart.svg",
			"meta": map[string]interface{}{
				"author": "^CloudCart LLC$",
			},
			"script":  "/cloudcart-(?:assets|storage)/",
			"website": "http://cloudcart.com",
		},
		"CloudFlare": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server":          "^cloudflare$",
				"cf-cache-status": "",
				"cf-ray":          "",
			},
			"icon": "CloudFlare.svg",
			"cookies": map[string]interface{}{
				"__cfduid": "",
			},
			"js": map[string]interface{}{
				"CloudFlare": "",
			},
			"website": "http://www.cloudflare.com",
		},
		"Cloudcoins": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"js": map[string]interface{}{
				"CLOUDCOINS": "",
			},
			"script":  "https?://cdn\\.cloudcoins\\.co/javascript/cloudcoins\\.min\\.js",
			"website": "https://cloudcoins.co",
		},
		"Cloudera": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"headers": map[string]interface{}{
				"Server": "cloudera",
			},
			"icon":    "Cloudera.png",
			"website": "http://www.cloudera.com",
		},
		"Coaster CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "coaster-cms.png",
			"implies": "Laravel",
			"meta": map[string]interface{}{
				"generator": "^Coaster CMS v([\\d.]+)$\\;version:\\1",
			},
			"website": "https://www.coastercms.org",
		},
		"CodeIgniter": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"ci_csrf_token":     "^(.+)$\\;version:\\1?2+:",
				"ci_session":        "",
				"exp_last_activity": "",
				"exp_tracker":       "",
			},
			"html":    "<input[^>]+name=\"ci_csrf_token\"\\;version:2+",
			"icon":    "CodeIgniter.png",
			"implies": "PHP",
			"website": "http://codeigniter.com",
		},
		"CodeMirror": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon": "CodeMirror.png",
			"js": map[string]interface{}{
				"CodeMirror":         "",
				"CodeMirror.version": "^(.+)$\\;version:\\1",
			},
			"website": "http://codemirror.net",
		},
		"CoinHive": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon": "CoinHive.svg",
			"js": map[string]interface{}{
				"CoinHive": "",
			},
			"script": []interface{}{
				"\\/(?:coinhive|(authedmine))(?:\\.min)?\\.js\\;version:\\1?opt-in:",
				"coinhive\\.com/lib",
			},
			"url":     "https?://cnhv\\.co/",
			"website": "https://coinhive.com",
		},
		"CoinHive Captcha": map[string]interface{}{
			"cats": []interface{}{
				16,
				56,
			},
			"html":    "(?:<div[^>]+class=\"coinhive-captcha[^>]+data-key|<div[^>]+data-key[^>]+class=\"coinhive-captcha)",
			"icon":    "CoinHive.svg",
			"script":  "https?://authedmine\\.com/(?:lib/captcha|captcha)",
			"website": "https://coinhive.com",
		},
		"Coinhave": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon":    "coinhave.png",
			"script":  "https?://coin-have\\.com/c/[0-9a-zA-Z]{4}\\.js",
			"website": "https://coin-have.com/",
		},
		"Coinimp": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon": "coinimp.png",
			"js": map[string]interface{}{
				"Client.Anonymous": "\\;confidence:50",
			},
			"script":  "https?://www\\.hashing\\.win/scripts/min\\.js",
			"website": "https://www.coinimp.com",
		},
		"Coinlab": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon": "coinlab.png",
			"js": map[string]interface{}{
				"Coinlab": "",
			},
			"script":  "https?://coinlab\\.biz/lib/coinlab\\.js\\?id=",
			"website": "https://coinlab.biz/en",
		},
		"ColorMeShop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "colormeshop.png",
			"js": map[string]interface{}{
				"Colorme": "",
			},
			"website": "https://shop-pro.jp",
		},
		"Comandia": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<link[^>]+=['\"]//cdn\\.mycomandia\\.com",
			"icon": "Comandia.svg",
			"js": map[string]interface{}{
				"Comandia": "",
			},
			"website": "http://comandia.com",
		},
		"Combeenation": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<iframe[^>]+src=\"[^>]+portal\\.combeenation\\.com",
			"icon":    "Combeenation.png",
			"website": "https://www.combeenation.com",
		},
		"Commerce Server": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"COMMERCE-SERVER-SOFTWARE": "",
			},
			"icon":    "Commerce Server.png",
			"implies": "Microsoft ASP.NET",
			"website": "http://commerceserver.net",
		},
		"CompaqHTTPServer": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "CompaqHTTPServer\\/?([\\d\\.]+)?\\;version:\\1",
			},
			"icon":    "HP.svg",
			"website": "http://www.hp.com",
		},
		"Concrete5": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Concrete5.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"CCM_IMAGE_PATH": "",
			},
			"cookies": map[string]interface{}{
				"CONCRETE5": "",
			},
			"meta": map[string]interface{}{
				"generator": "^concrete5 - ([\\d.]+)$\\;version:\\1",
			},
			"script":  "/concrete/js/",
			"website": "https://concrete5.org",
		},
		"Connect": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Connect$",
			},
			"icon":    "Connect.png",
			"implies": "Node.js",
			"website": "http://www.senchalabs.org/connect",
		},
		"Contao": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": []interface{}{
				"<!--[^>]+powered by (?:TYPOlight|Contao)[^>]*-->",
				"<link[^>]+(?:typolight|contao)\\.css",
			},
			"icon":    "Contao.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "^Contao Open Source CMS$",
			},
			"website": "http://contao.org",
		},
		"Contenido": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Contenido.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Contenido ([\\d.]+)\\;version:\\1",
			},
			"website": "http://contenido.org/en",
		},
		"Contensis": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Contensis.png",
			"implies": []interface{}{
				"Java",
				"CFML",
			},
			"meta": map[string]interface{}{
				"generator": "Contensis CMS Version ([\\d.]+)\\;version:\\1",
			},
			"website": "https://zengenti.com/en-gb/products/contensis",
		},
		"ContentBox": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"icon":    "ContentBox.png",
			"implies": "Adobe ColdFusion",
			"meta": map[string]interface{}{
				"generator": "ContentBox powered by ColdBox",
			},
			"website": "http://www.gocontentbox.org",
		},
		"Contentful": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<[^>]+(?:https?:)?//(?:assets|downloads|images|videos)\\.(?:ct?fassets\\.net|contentful\\.com)",
			"icon":    "Contentful.svg",
			"website": "http://www.contentful.com",
		},
		"ConversionLab": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "ConversionLab.png",
			"script":  "conversionlab\\.trackset\\.com/track/tsend\\.js",
			"website": "http://www.trackset.it/conversionlab",
		},
		"Coppermine": map[string]interface{}{
			"cats": []interface{}{
				7,
			},
			"html":    "<!--Coppermine Photo Gallery ([\\d.]+)\\;version:\\1",
			"icon":    "Coppermine.png",
			"implies": "PHP",
			"website": "http://coppermine-gallery.net",
		},
		"Cosmoshop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "Cosmoshop.png",
			"script":  "cosmoshop_functions\\.js",
			"website": "http://cosmoshop.de",
		},
		"Cotonti": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Cotonti.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Cotonti",
			},
			"website": "http://www.cotonti.com",
		},
		"CouchDB": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "CouchDB/([\\d.]+)\\;version:\\1",
			},
			"icon":    "CouchDB.png",
			"website": "http://couchdb.apache.org",
		},
		"Countly": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Countly.png",
			"js": map[string]interface{}{
				"Countly": "",
			},
			"website": "https://count.ly",
		},
		"Cowboy": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^Cowboy$",
			},
			"icon":    "Cowboy.png",
			"implies": "Erlang",
			"website": "http://ninenines.eu",
		},
		"CppCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^CppCMS/([\\d.]+)$\\;version:\\1",
			},
			"icon":    "CppCMS.png",
			"implies": "C\\+\\+",
			"website": "http://cppcms.com",
		},
		"Craft CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"CraftSessionId": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "\\bCraft CMS\\b",
			},
			"icon":    "Craft CMS.svg",
			"implies": "Yii",
			"website": "https://craftcms.com",
		},
		"Craft Commerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "\\bCraft Commerce\\b",
			},
			"icon":    "Craft CMS.svg",
			"implies": "Craft CMS",
			"website": "https://craftcommerce.com",
		},
		"Crazy Egg": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Crazy Egg.png",
			"js": map[string]interface{}{
				"CE2": "",
			},
			"script":  "script\\.crazyegg\\.com/pages/scripts/\\d+/\\d+\\.js",
			"website": "http://crazyegg.com",
		},
		"Criteo": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Criteo.svg",
			"js": map[string]interface{}{
				"Criteo":        "",
				"criteo_pubtag": "",
				"criteo_q":      "",
			},
			"script": []interface{}{
				"//(?:cas\\.criteo\\.com|(?:[^/]\\.)?criteo\\.net)/",
				"//static.criteo.net/js/ld/ld.js",
			},
			"website": "http://criteo.com",
		},
		"Cross Pixel": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Cross Pixel.png",
			"js": map[string]interface{}{
				"cp_C4w1ldN2d9PmVrkN": "",
			},
			"script":  "tag\\.crsspxl\\.com/s1\\.js",
			"website": "http://datadesk.crsspxl.com",
		},
		"CrossBox": map[string]interface{}{
			"cats": []interface{}{
				30,
			},
			"icon":    "CrossBox.png",
			"html":    "<span class=\"product-name-loading\">CrossBox Premium Webmail",
			"website": "https://crossbox.io",
		},
		"Crypto-Loot": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon": "Crypto-Loot.png",
			"js": map[string]interface{}{
				"CRLT.CONFIG.ASMJS_NAME": "",
				"CryptoLoot":             "",
			},
			"script": []interface{}{
				"^/crypto-loot\\.com/lib/",
				"^/webmine\\.pro/",
				"^/cryptoloot\\.pro/",
				"/crlt\\.js\\;confidence:75",
			},
			"website": "https://crypto-loot.com/",
		},
		"CubeCart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "(?:Powered by <a href=[^>]+cubecart\\.com|<p[^>]+>Powered by CubeCart)",
			"icon":    "CubeCart.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "cubecart",
			},
			"website": "http://www.cubecart.com",
		},
		"Cufon": map[string]interface{}{
			"cats": []interface{}{
				17,
			},
			"icon": "Cufon.png",
			"js": map[string]interface{}{
				"Cufon": "",
			},
			"script":  "cufon-yui\\.js",
			"website": "http://cufon.shoqolate.com",
		},
		"D3": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "D3.png",
			"js": map[string]interface{}{
				"d3.version": "^(.+)$\\;version:\\1",
			},
			"script":  "/d3(?:\\. v\\d+)?(?:\\.min)?\\.js",
			"website": "http://d3js.org",
		},
		"DHTMLX": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "DHTMLX.png",
			"script":  "dhtmlxcommon\\.js",
			"website": "http://dhtmlx.com",
		},
		"DERAK.CLOUD": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server":        "^DERAK.CLOUD$",
				"Derak-Umbrage": "",
			},
			"icon": "DerakCloud.png",
			"cookies": map[string]interface{}{
				"__derak_auth": "",
				"__derak_user": "",
			},
			"js": map[string]interface{}{
				"derakCloud.init": "",
			},
			"website": "https://derak.cloud/",
		},
		"DM Polopoly": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<(?:link [^>]*href|img [^>]*src)=\"/polopoly_fs/",
			"icon":    "DM Polopoly.png",
			"implies": "Java",
			"website": "http://www.atex.com/products/dm-polopoly",
		},
		"DNN": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"DotNetNukeAnonymous": "",
			},
			"headers": map[string]interface{}{
				"Cookie":          "dnn_IsMobile=",
				"DNNOutputCache":  "",
				"X-Compressed-By": "DotNetNuke",
			},
			"html": []interface{}{
				"<!-- by DotNetNuke Corporation",
				"<!-- DNN Platform",
			},
			"icon":    "DNN.png",
			"implies": "Microsoft ASP.NET",
			"js": map[string]interface{}{
				"DotNetNuke":     "",
				"dnn.apiversion": "^(.+)$\\;version:\\1",
			},
			"meta": map[string]interface{}{
				"generator": "DotNetNuke",
			},
			"script": []interface{}{
				"/js/dnncore\\.js",
				"/js/dnn\\.js",
			},
			"website": "http://dnnsoftware.com",
		},
		"DTG": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": []interface{}{
				"<a[^>]+Site Powered by DTG",
			},
			"icon":    "DTG.png",
			"implies": "Mono.net",
			"website": "https://www.dtg.nl",
		},
		"Dancer": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"Server":       "Perl Dancer ([\\d.]+)\\;version:\\1",
				"X-Powered-By": "Perl Dancer ([\\d.]+)\\;version:\\1",
			},
			"icon":    "Dancer.png",
			"implies": "Perl",
			"website": "http://perldancer.org",
		},
		"Danneo CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "CMS Danneo ([\\d.]+)\\;version:\\1",
			},
			"icon": "Danneo CMS.png",
			"implies": []interface{}{
				"Apache",
				"PHP",
			},
			"meta": map[string]interface{}{
				"generator": "Danneo CMS ([\\d.]+)\\;version:\\1",
			},
			"website": "http://danneo.com",
		},
		"Dart": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"excludes": []interface{}{
				"Angular",
				"AngularJS",
			},
			"html":    "/(?:<script)[^>]+(?:type=\"application/dart\")/",
			"icon":    "Dart.svg",
			"implies": "AngularDart",
			"js": map[string]interface{}{
				"___dart__$dart_dartObject_ZxYxX_0_": "",
				"___dart_dispatch_record_ZxYxX_0_":   "",
			},
			"script": []interface{}{
				"/(?:\\.)?(?:dart)(?:\\.js)?/",
				"packages/browser/dart\\.js",
			},
			"website": "https://www.dartlang.org",
		},
		"Darwin": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":       "Darwin",
				"X-Powered-By": "Darwin",
			},
			"icon":    "Apple.svg",
			"website": "https://opensource.apple.com",
		},
		"Datadome": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"cookies": map[string]interface{}{
				"datadome": "",
			},
			"script": "^https://ct\\.datadome\\.co/[a-z]\\.js$",
			"headers": map[string]interface{}{
				"X-DataDome":     "",
				"Server":         "^DataDome$",
				"X-DataDome-CID": "",
			},
			"icon":    "datadome.png",
			"website": "https://datadome.co/",
		},
		"DataLife Engine": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "DataLife Engine.png",
			"implies": []interface{}{
				"PHP",
				"Apache",
			},
			"js": map[string]interface{}{
				"dle_root": "",
			},
			"meta": map[string]interface{}{
				"generator": "DataLife Engine",
			},
			"website": "https://dle-news.ru",
		},
		"DataTables": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "DataTables.png",
			"implies": "jQuery",
			"script":  "dataTables.*\\.js",
			"website": "http://datatables.net",
		},
		"Day.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "Day.js.svg",
			"js": map[string]interface{}{
				"dayjs": "",
			},
			"website": "https://github.com/iamkun/dayjs",
		},
		"Debian": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":       "Debian",
				"X-Powered-By": "(?:Debian|dotdeb|(potato|woody|sarge|etch|lenny|squeeze|wheezy|jessie|stretch|buster|sid))\\;version:\\1",
			},
			"icon":    "Debian.png",
			"website": "https://debian.org",
		},
		"DedeCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "DedeCMS.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"DedeContainer": "",
			},
			"script":  "dedeajax",
			"website": "http://dedecms.com",
		},
		"DirectAdmin": map[string]interface{}{
			"cats": []interface{}{
				9,
			},
			"headers": map[string]interface{}{
				"Server": "DirectAdmin Daemon v([\\d.]+)\\;version:\\1",
			},
			"html": "<a[^>]+>DirectAdmin</a> Web Control Panel",
			"icon": "DirectAdmin.png",
			"implies": []interface{}{
				"PHP",
				"Apache",
			},
			"website": "https://www.directadmin.com",
		},
		"Discourse": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"icon":    "Discourse.png",
			"implies": "Ruby on Rails",
			"js": map[string]interface{}{
				"Discourse": "",
			},
			"meta": map[string]interface{}{
				"generator": "Discourse(?: ?/?([\\d.]+\\d))?\\;version:\\1",
			},
			"website": "https://discourse.org",
		},
		"Discuz! X": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"icon":    "Discuz X.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"DISCUZCODE":    "",
				"discuzVersion": "^(.+)$\\;version:\\1",
				"discuz_uid":    "",
			},
			"meta": map[string]interface{}{
				"generator": "Discuz! X([\\d\\.]+)?\\;version:\\1",
			},
			"website": "http://www.discuz.net",
		},
		"Disqus": map[string]interface{}{
			"cats": []interface{}{
				15,
			},
			"html": "<div[^>]+id=\"disqus_thread\"",
			"icon": "Disqus.svg",
			"js": map[string]interface{}{
				"DISQUS":           "",
				"disqus_shortname": "",
				"disqus_url":       "",
			},
			"script":  "disqus_url",
			"website": "https://disqus.com",
		},
		"Django": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "(?:powered by <a[^>]+>Django ?([\\d.]+)?<\\/a>|<input[^>]*name=[\"']csrfmiddlewaretoken[\"'][^>]*>)\\;version:\\1",
			"icon":    "Django.png",
			"implies": "Python",
			"js": map[string]interface{}{
				"__admin_media_prefix__": "",
				"django":                 "",
			},
			"website": "https://djangoproject.com",
		},
		"Django CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Django CMS.png",
			"implies": "Django",
			"website": "https://django-cms.org",
		},
		"Docusaurus": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"icon": "docusaurus.svg",
			"implies": []interface{}{
				"React",
				"webpack",
			},
			"js": map[string]interface{}{
				"search.indexName": "",
			},
			"meta": map[string]interface{}{
				"generator": "^Docusaurus$",
			},
			"website": "https://docusaurus.io/",
		},
		"Docker": map[string]interface{}{
			"cats": []interface{}{
				60,
			},
			"icon":    "Docker.svg",
			"implies": "Linux",
			"html":    "<!-- This comment is expected by the docker HEALTHCHECK  -->",
			"website": "https://www.docker.com/",
		},
		"Dojo": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "Dojo.png",
			"js": map[string]interface{}{
				"dojo":               "",
				"dojo.version.major": "^(.+)$\\;version:\\1",
			},
			"script":  "([\\d.]+)/dojo/dojo(?:\\.xd)?\\.js\\;version:\\1",
			"website": "https://dojotoolkit.org",
		},
		"Dokeos": map[string]interface{}{
			"cats": []interface{}{
				21,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Dokeos",
			},
			"html": "(?:Portal <a[^>]+>Dokeos|@import \"[^\"]+dokeos_blue)",
			"icon": "Dokeos.png",
			"implies": []interface{}{
				"PHP",
				"Xajax",
				"jQuery",
				"CKEditor",
			},
			"meta": map[string]interface{}{
				"generator": "Dokeos",
			},
			"website": "https://dokeos.com",
		},
		"DokuWiki": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"cookies": map[string]interface{}{
				"DokuWiki": "",
			},
			"html": []interface{}{
				"<div[^>]+id=\"dokuwiki__>",
				"<a[^>]+href=\"#dokuwiki__",
			},
			"icon":    "DokuWiki.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "^DokuWiki( Release [\\d-]+)?\\;version:\\1",
			},
			"website": "https://www.dokuwiki.org",
		},
		"Dotclear": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Dotclear-Static-Cache": "",
			},
			"icon":    "Dotclear.png",
			"implies": "PHP",
			"website": "http://dotclear.org",
		},
		"DoubleClick Ad Exchange (AdX)": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "DoubleClick.svg",
			"script": []interface{}{
				"googlesyndication\\.com/pagead/show_ads\\.js",
				"tpc\\.googlesyndication\\.com/safeframe",
				"googlesyndication\\.com.*abg\\.js",
			},
			"website": "http://www.doubleclickbygoogle.com/solutions/digital-marketing/ad-exchange/",
		},
		"DoubleClick Campaign Manager (DCM)": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon":    "DoubleClick.svg",
			"script":  "2mdn\\.net",
			"website": "http://www.doubleclickbygoogle.com/solutions/digital-marketing/campaign-manager/",
		},
		"DoubleClick Floodlight": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon":    "DoubleClick.svg",
			"script":  "https?://fls\\.doubleclick\\.net",
			"website": "http://support.google.com/ds/answer/6029713?hl=en",
		},
		"DoubleClick for Publishers (DFP)": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon":    "DoubleClick.svg",
			"script":  "googletagservices\\.com/tag/js/gpt(?:_mobile)?\\.js",
			"website": "http://www.google.com/dfp",
		},
		"DovetailWRP": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<link[^>]* href=\"\\/DovetailWRP\\/",
			"icon":    "DovetailWRP.png",
			"implies": "Microsoft ASP.NET",
			"script":  "\\/DovetailWRP\\/",
			"website": "http://www.dovetailinternet.com",
		},
		"Doxygen": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"html": "(?:<!-- Generated by Doxygen ([\\d.]+)|<link[^>]+doxygen\\.css)\\;version:\\1",
			"icon": "Doxygen.png",
			"meta": map[string]interface{}{
				"generator": "Doxygen ([\\d.]+)\\;version:\\1",
			},
			"website": "http://www.doxygen.nl/",
		},
		"DreamWeaver": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"html": "<!--[^>]*(?:InstanceBeginEditable|Dreamweaver([^>]+)target|DWLayoutDefaultTable)\\;version:\\1",
			"js": map[string]interface{}{
				"MM_showMenu":       "",
				"MM_preloadImages":  "",
				"MM_showHideLayers": "",
			},
			"icon":    "DreamWeaver.png",
			"website": "https://www.adobe.com/products/dreamweaver.html",
		},
		"Drupal": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"Expires":        "19 Nov 1978",
				"X-Drupal-Cache": "",
				"X-Generator":    "^Drupal(?:\\s([\\d.]+))?\\;version:\\1",
			},
			"html":    "<(?:link|style)[^>]+\"/sites/(?:default|all)/(?:themes|modules)/",
			"icon":    "Drupal.svg",
			"implies": "PHP",
			"js": map[string]interface{}{
				"Drupal": "",
			},
			"meta": map[string]interface{}{
				"generator": "^Drupal(?:\\s([\\d.]+))?\\;version:\\1",
			},
			"script":  "drupal\\.js",
			"website": "https://drupal.org",
		},
		"Drupal Commerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<[^>]+(?:id=\"block[_-]commerce[_-]cart[_-]cart|class=\"commerce[_-]product[_-]field)",
			"icon":    "Drupal Commerce.png",
			"implies": "Drupal",
			"website": "http://drupalcommerce.org",
		},
		"Dynamicweb": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
				10,
			},
			"cookies": map[string]interface{}{
				"Dynamicweb": "",
			},
			"icon":    "Dynamicweb.png",
			"implies": "Microsoft ASP.NET",
			"meta": map[string]interface{}{
				"generator": "Dynamicweb ([\\d.]+)\\;version:\\1",
			},
			"website": "http://www.dynamicweb.dk",
		},
		"Dynatrace": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"headers": map[string]interface{}{
				"X-dynaTrace-JS-Agent": "",
			},
			"icon":    "Dynatrace.png",
			"script":  "dtagent.*\\.js",
			"website": "http://dynatrace.com",
		},
		"EasyEngine": map[string]interface{}{
			"cats": []interface{}{
				47,
				9,
			},
			"icon": "EasyEngine.png",
			"implies": []interface{}{
				"Docker",
			},
			"headers": map[string]interface{}{
				"x-powered-by": "^EasyEngine (.*)$\\;version:\\1",
			},
			"website": "https://easyengine.io",
		},
		"EC-CUBE": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "ec-cube.png",
			"implies": "PHP",
			"script": []interface{}{
				"eccube\\.js",
				"win_op\\.js",
			},
			"website": "http://www.ec-cube.net",
		},
		"Elementor": map[string]interface{}{
			"cats": []interface{}{
				51,
			},
			"html": []interface{}{
				"<div class=(?:\"|')[^\"']*elementor",
				"<section class=(?:\"|')[^\"']*elementor",
				"<link [^>]*href=(?:\"|')[^\"']*elementor/assets",
				"<link [^>]*href=(?:\"|')[^\"']*uploads/elementor/css",
			},
			"js": map[string]interface{}{
				"elementorFrontend.getElements": "",
			},
			"script":  "elementor/assets/js/[^/]+\\.js\\?ver=([\\d.]+)$\\;version:\\1",
			"icon":    "Elementor.png",
			"implies": "WordPress",
			"website": "https://elementor.com",
		},
		"ELOG": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html":    "<title>ELOG Logbook Selection</title>",
			"icon":    "ELOG.png",
			"website": "http://midas.psi.ch/elog",
		},
		"ELOG HTTP": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "ELOG HTTP ?([\\d.-]+)?\\;version:\\1",
			},
			"icon":    "ELOG.png",
			"implies": "ELOG",
			"website": "http://midas.psi.ch/elog",
		},
		"EPages": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "epages 6",
			},
			"html":    "<div class=\"BoxContainer\">",
			"icon":    "epages.png",
			"website": "http://www.epages.com/",
		},
		"EPiServer": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"EPiServer": "",
				"EPiTrace":  "",
			},
			"icon":    "EPiServer.png",
			"implies": "Microsoft ASP.NET",
			"meta": map[string]interface{}{
				"generator": "EPiServer",
			},
			"website": "http://episerver.com",
		},
		"EPrints": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon":    "EPrints.png",
			"implies": "Perl",
			"js": map[string]interface{}{
				"EPJS_menu_template": "",
				"EPrints":            "",
			},
			"meta": map[string]interface{}{
				"generator": "EPrints ([\\d.]+)\\;version:\\1",
			},
			"website": "http://www.eprints.org",
		},
		"EdgeCast": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server": "^ECD\\s\\(\\S+\\)",
			},
			"icon":    "EdgeCast.png",
			"url":     "https?://(?:[^/]+\\.)?edgecastcdn\\.net/",
			"website": "http://www.edgecast.com",
		},
		"Elcodi": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Elcodi": "",
			},
			"icon": "Elcodi.png",
			"implies": []interface{}{
				"PHP",
				"Symfony",
			},
			"website": "http://elcodi.io",
		},
		"Eleanor CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Eleanor CMS.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Eleanor",
			},
			"website": "http://eleanor-cms.ru",
		},
		"Element UI": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "ElementUI.svg",
			"implies": []interface{}{
				"Vue",
			},
			"html": []interface{}{
				"<(?:div|button) class=\"el-(?:table-column|table-filter|popper|pagination|pager|select-group|form|form-item|color-predefine|color-hue-slider|color-svpanel|color-alpha-slider|color-dropdown|color-picker|badge|tree|tree-node|select|message|dialog|checkbox|checkbox-button|checkbox-group|container|steps|carousel|menu|menu-item|submenu|menu-item-group|button|button-group|card|table|select-dropdown|row|tabs|notification|radio|progress|progress-bar|tag|popover|tooltip|cascader|cascader-menus|cascader-menu|time-spinner|spinner|spinner-inner|transfer|transfer-panel|rate|slider|dropdown|dropdown-menu|textarea|input|input-group|popup-parent|radio-group|main|breadcrumb|time-range-picker|date-range-picker|year-table|date-editor|range-editor|time-spinner|date-picker|time-panel|date-table|month-table|picker-panel|collapse|collapse-item|alert|select-dropdown|select-dropdown__empty|select-dropdown__wrap|select-dropdown__list|scrollbar|switch|carousel|upload|upload-dragger|upload-list|upload-cover|aside|input-number|header|message-box|footer|radio-button|step|autocomplete|autocomplete-suggestion|loading-parent|loading-mask|loading-spinner|)",
			},
			"website": "https://element.eleme.io/",
		},
		"Eloqua": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"icon": "Oracle.png",
			"js": map[string]interface{}{
				"elqCurESite": "",
				"elqLoad":     "",
				"elqSiteID":   "",
				"elq_global":  "",
			},
			"script":  "elqCfg\\.js",
			"website": "http://eloqua.com",
		},
		"EmbedThis Appweb": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Mbedthis-Appweb(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Embedthis.png",
			"website": "http://embedthis.com/appweb",
		},
		"Ember.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon":    "Ember.js.png",
			"implies": "Handlebars",
			"js": map[string]interface{}{
				"Ember":         "",
				"Ember.VERSION": "^(.+)$\\;version:\\1",
			},
			"website": "http://emberjs.com",
		},
		"Ensighten": map[string]interface{}{
			"cats": []interface{}{
				42,
			},
			"script":  "//nexus\\.ensighten\\.com/",
			"icon":    "ensighten.png",
			"website": "https://success.ensighten.com/hc/en-us",
		},
		"Envoy": map[string]interface{}{
			"cats": []interface{}{
				64,
			},
			"icon": "Envoy.png",
			"headers": map[string]interface{}{
				"Server":                        "^envoy$",
				"x-envoy-upstream-service-time": "",
			},
			"website": "https://www.envoyproxy.io/",
		},
		"Enyo": map[string]interface{}{
			"cats": []interface{}{
				12,
				26,
			},
			"icon": "Enyo.png",
			"js": map[string]interface{}{
				"enyo": "",
			},
			"script":  "enyo\\.js",
			"website": "http://enyojs.com",
		},
		"Epoch": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"html":    "<link[^>]+?href=\"[^\"]+epoch(?:\\.min)?\\.css",
			"implies": "D3",
			"script":  "epoch(?:\\.min)?\\.js",
			"website": "https://fastly.github.io/epoch",
		},
		"Epom": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Epom.png",
			"js": map[string]interface{}{
				"epomCustomParams": "",
			},
			"url":     "^https?://(?:[^/]+\\.)?ad(?:op)?shost1\\.com/",
			"website": "http://epom.com",
		},
		"Erlang": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"headers": map[string]interface{}{
				"Server": "Erlang( OTP/(?:[\\d.ABR-]+))?\\;version:\\1",
			},
			"icon":    "Erlang.png",
			"website": "http://www.erlang.org",
		},
		"Essential JS 2": map[string]interface{}{
			"cats": []interface{}{
				12,
				18,
				59,
			},
			"icon":    "syncfusion.svg",
			"website": "https://www.syncfusion.com/javascript-ui-controls",
			"js": map[string]interface{}{
				"ej.base": "",
			},
			"html": "<[^<]* class=\"[^<]*\be-control\b[^-][^<]*\be-lib\b[^>]*\"[^>]*>",
		},
		"Etherpad": map[string]interface{}{
			"cats": []interface{}{
				24,
			},
			"headers": map[string]interface{}{
				"Server": "^Etherpad",
			},
			"icon":    "etherpad.png",
			"implies": "Node.js",
			"js": map[string]interface{}{
				"padeditbar": "",
				"padimpexp":  "",
			},
			"script": []interface{}{
				"/ep_etherpad-lite/",
			},
			"website": "https://etherpad.org",
		},
		"Exhibit": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "Exhibit.png",
			"js": map[string]interface{}{
				"Exhibit":         "",
				"Exhibit.version": "^(.+)$\\;version:\\1",
			},
			"script":  "exhibit.*\\.js",
			"website": "http://simile-widgets.org/exhibit/",
		},
		"Express": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Express$",
			},
			"icon":    "Express.png",
			"implies": "Node.js",
			"website": "http://expressjs.com",
		},
		"ExpressionEngine": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"exp_csrf_token":    "",
				"exp_last_activity": "",
				"exp_tracker":       "",
			},
			"icon":    "ExpressionEngine.png",
			"implies": "PHP",
			"website": "http://expressionengine.com",
		},
		"ExtJS": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "ExtJS.png",
			"js": map[string]interface{}{
				"Ext":                        "",
				"Ext.version":                "^(.+)$\\;version:\\1",
				"Ext.versions.extjs.version": "^(.+)$\\;version:\\1",
			},
			"script":  "ext-base\\.js",
			"website": "https://www.sencha.com",
		},
		"F5 BigIP": map[string]interface{}{
			"cats": []interface{}{
				64,
			},
			"headers": map[string]interface{}{
				"server": "^big-?ip$",
			},
			"cookies": map[string]interface{}{
				"F5_ST":           "",
				"MRHSHint":        "",
				"F5_HT_shrinked":  "",
				"F5_fullWT":       "",
				"MRHSequence":     "",
				"MRHSession":      "",
				"LastMRH_Session": "",
				"TIN":             "",
			},
			"icon":    "F5.png",
			"website": "https://www.f5.com/products/big-ip-services",
		},
		"FAST ESP": map[string]interface{}{
			"cats": []interface{}{
				29,
			},
			"html":    "<form[^>]+id=\"fastsearch\"",
			"icon":    "FAST ESP.png",
			"website": "http://microsoft.com/enterprisesearch",
		},
		"FAST Search for SharePoint": map[string]interface{}{
			"cats": []interface{}{
				29,
			},
			"html": "<input[^>]+ name=\"ParametricSearch",
			"icon": "FAST Search for SharePoint.png",
			"implies": []interface{}{
				"Microsoft SharePoint",
				"Microsoft ASP.NET",
			},
			"url":     "Pages/SearchResults\\.aspx\\?k=",
			"website": "http://sharepoint.microsoft.com/en-us/product/capabilities/search/Pages/Fast-Search.aspx",
		},
		"FWP": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<!--\\s+FwP Systems",
			"icon": "FWP.png",
			"meta": map[string]interface{}{
				"generator": "FWP Shop",
			},
			"website": "http://fwpshop.org",
		},
		"Facebook": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "Facebook.svg",
			"script":  "//connect\\.facebook\\.net/[^/]*/[a-z]*\\.js",
			"website": "http://facebook.com",
		},
		"Fact Finder": map[string]interface{}{
			"cats": []interface{}{
				29,
			},
			"html":    "<!-- Factfinder",
			"icon":    "Fact Finder.png",
			"script":  "Suggest\\.ff",
			"url":     "(?:/ViewParametricSearch|ffsuggest\\.[a-z]htm)",
			"website": "http://fact-finder.com",
		},
		"FancyBox": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "FancyBox.png",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"$.fancybox.version": "^(.+)$\\;version:\\1",
			},
			"script":  "jquery\\.fancybox(?:\\.pack|\\.min)?\\.js(?:\\?v=([\\d.]+))?$\\;version:\\1",
			"website": "http://fancyapps.com/fancybox",
		},
		"Fastcommerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Fastcommerce.png",
			"meta": map[string]interface{}{
				"generator": "^Fastcommerce",
			},
			"website": "https://www.fastcommerce.com.br",
		},
		"Fastly": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Fastly-Debug-Digest": "",
				"Vary":                "Fastly-SSL",
				"x-via-fastly:":       "",
				"X-Fastly-Request-ID": "",
			},
			"icon":    "Fastly.svg",
			"website": "https://www.fastly.com",
		},
		"Fat-Free Framework": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Fat-Free Framework$",
			},
			"icon":    "Fat-Free Framework.png",
			"implies": "PHP",
			"website": "http://fatfreeframework.com",
		},
		"Fbits": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Fbits.png",
			"js": map[string]interface{}{
				"fbits": "",
			},
			"website": "https://www.traycorp.com.br",
		},
		"Fedora": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server": "Fedora",
			},
			"icon":    "Fedora.png",
			"website": "http://fedoraproject.org",
		},
		"Fingerprintjs": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"js": map[string]interface{}{
				"Fingerprint":          "(\\d)?$\\;version:\\1",
				"Fingerprint2":         "",
				"Fingerprint2.VERSION": "^(.+)$\\;version:\\1",
			},
			"script":  "fingerprint(\\d)?(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://valve.github.io/fingerprintjs2/",
		},
		"Firebase": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon": "Firebase.png",
			"js": map[string]interface{}{
				"firebase.SDK_VERSION": "([\\d.]+)$\\;version:\\1",
			},
			"script":  "/(?:([\\d.]+)/)?firebase(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://firebase.com",
		},
		"Fireblade": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server": "fbs",
			},
			"icon":    "Fireblade.png",
			"website": "http://fireblade.com",
		},
		"Flarum": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html": "<div id=\"flarum-loading\"",
			"icon": "flarum.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"js": map[string]interface{}{
				"app.cache.discussionList": "",
				"app.forum.freshness":      "",
			},
			"website": "http://flarum.org/",
		},
		"Flask": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Werkzeug/?([\\d\\.]+)?\\;version:\\1",
			},
			"icon":    "Flask.png",
			"implies": "Python",
			"website": "http://flask.pocoo.org",
		},
		"Flat UI": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "<link[^>]* href=[^>]+flat-ui(?:\\.min)?\\.css",
			"icon":    "Flat UI.png",
			"implies": "Bootstrap",
			"website": "https://designmodo.github.io/Flat-UI/",
		},
		"FlexCMP": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Flex-Lang":  "",
				"X-Powered-By": "FlexCMP.+\\[v\\. ([\\d.]+)\\;version:\\1",
			},
			"html": "<!--[^>]+FlexCMP[^>v]+v\\. ([\\d.]+)\\;version:\\1",
			"icon": "FlexCMP.png",
			"meta": map[string]interface{}{
				"generator": "^FlexCMP",
			},
			"website": "http://www.flexcmp.com/cms/home",
		},
		"FlexSlider": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "FlexSlider.png",
			"implies": "jQuery",
			"script": []interface{}{
				"jquery\\.flexslider(?:\\.min)?\\.js$",
			},
			"website": "https://woocommerce.com/flexslider/",
		},
		"Flickity": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"js": map[string]interface{}{
				"Flickity": "",
			},
			"script":  "/flickity(?:\\.pkgd)?(?:\\.min)?\\.js",
			"website": "https://flickity.metafizzy.co/",
		},
		"FluxBB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html":    "<p id=\"poweredby\">[^<]+<a href=\"https?://fluxbb\\.org/\">",
			"icon":    "FluxBB.png",
			"implies": "PHP",
			"website": "https://fluxbb.org",
		},
		"Flyspray": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"cookies": map[string]interface{}{
				"flyspray_project": "",
			},
			"html":    "(?:<a[^>]+>Powered by Flyspray|<map id=\"projectsearchform)",
			"icon":    "Flyspray.png",
			"implies": "PHP",
			"website": "http://flyspray.org",
		},
		"Font Awesome": map[string]interface{}{
			"cats": []interface{}{
				17,
			},
			"html": []interface{}{
				"<link[^>]* href=[^>]+(?:([\\d.]+)/)?(?:css/)?font-awesome(?:\\.min)?\\.css\\;version:\\1",
				"<script[^>]* src=[^>]+fontawesome(?:\\.js)?",
			},
			"icon":    "Font Awesome.png",
			"website": "http://fontawesome.io",
		},
		"Fork CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "ForkCMS.png",
			"implies": "Symfony",
			"meta": map[string]interface{}{
				"generator": "^Fork CMS$",
			},
			"website": "http://www.fork-cms.com/",
		},
		"Fortune3": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "(?:<link [^>]*href=\"[^\\/]*\\/\\/www\\.fortune3\\.com\\/[^\"]*siterate\\/rate\\.css|Powered by <a [^>]*href=\"[^\"]+fortune3\\.com)",
			"icon":    "Fortune3.png",
			"script":  "cartjs\\.php\\?(?:.*&)?s=[^&]*myfortune3cart\\.com",
			"website": "http://fortune3.com",
		},
		"Foswiki": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"cookies": map[string]interface{}{
				"FOSWIKISTRIKEONE": "",
				"SFOSWIKISID":      "",
			},
			"headers": map[string]interface{}{
				"X-Foswikiaction": "",
				"X-Foswikiuri":    "",
			},
			"html": []interface{}{
				"<div class=\"foswiki(?:Copyright|Page|Main)\">",
			},
			"icon":    "foswiki.png",
			"implies": "Perl",
			"js": map[string]interface{}{
				"foswiki": "",
			},
			"meta": map[string]interface{}{
				"foswiki.SERVERTIME": "",
				"foswiki.WIKINAME":   "",
			},
			"website": "http://foswiki.org",
		},
		"FreeBSD": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server": "FreeBSD(?: ([\\d.]+))?\\;version:\\1",
			},
			"icon":    "FreeBSD.png",
			"website": "http://freebsd.org",
		},
		"FreeTextBox": map[string]interface{}{
			"cats": []interface{}{
				24,
			},
			"html":    "/<!--\\s*\\*\\s*FreeTextBox v\\d+ \\(([.\\d]+)(?:(?:.|\\n)+?<!--\\s*\\*\\s*License Type: (Distribution|Professional)License)?/i\\;version:\\1 \\2",
			"icon":    "FreeTextBox.png",
			"implies": "Microsoft ASP.NET",
			"js": map[string]interface{}{
				"FTB_API":      "",
				"FTB_AddEvent": "",
			},
			"website": "http://freetextbox.com",
		},
		"Freespee": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "Freespee.svg",
			"script":  "analytics\\.freespee\\.com/js/external/fs\\.(?:min\\.)?js",
			"website": "https://www.freespee.com",
		},
		"Freshchat": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon":    "freshchat.png",
			"script":  "wchat\\.freshchat\\.com/js/widget\\.js",
			"website": "https://www.freshworks.com/live-chat-software/",
		},
		"Freshmarketer": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "freshmarketer.png",
			"script":  "cdn\\.freshmarketer\\.com",
			"website": "https://www.freshworks.com/marketing-automation/conversion-rate-optimization/",
		},
		"Froala Editor": map[string]interface{}{
			"cats": []interface{}{
				24,
			},
			"html": "<[^>]+class=\"[^\"]*(?:fr-view|fr-box)",
			"icon": "Froala.svg",
			"implies": []interface{}{
				"jQuery",
				"Font Awesome",
			},
			"website": "http://froala.com/wysiwyg-editor",
		},
		"FrontPage": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"icon": "FrontPage.png",
			"meta": map[string]interface{}{
				"ProgId":    "^FrontPage\\.",
				"generator": "Microsoft FrontPage(?:\\s((?:Express )?[\\d.]+))?\\;version:\\1",
			},
			"website": "http://office.microsoft.com/frontpage",
		},
		"Fusion Ads": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Fusion Ads.png",
			"js": map[string]interface{}{
				"_fusion": "",
			},
			"script":  "^[^\\/]*//[ac]dn\\.fusionads\\.net/(?:api/([\\d.]+)/)?\\;version:\\1",
			"website": "http://fusionads.net",
		},
		"Future Shop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "futureshop.png",
			"script":  "future-shop.*\\.js",
			"website": "https://www.future-shop.jp",
		},
		"G-WAN": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "G-WAN",
			},
			"icon":    "G-WAN.png",
			"website": "http://gwan.com",
		},
		"GX WebManager": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<!--\\s+Powered by GX",
			"icon": "GX WebManager.png",
			"meta": map[string]interface{}{
				"generator": "GX WebManager(?: ([\\d.]+))?\\;version:\\1",
			},
			"website": "http://www.gxsoftware.com/en/products/web-content-management.htm",
		},
		"Gallery": map[string]interface{}{
			"cats": []interface{}{
				7,
			},
			"html": []interface{}{
				"<div id=\"gsNavBar\" class=\"gcBorder1\">",
				"<a href=\"http://gallery\\.sourceforge\\.net\"><img[^>]+Powered by Gallery\\s*(?:(?:v|Version)\\s*([0-9.]+))?\\;version:\\1",
			},
			"icon": "Gallery.png",
			"js": map[string]interface{}{
				"$.fn.gallery_valign": "",
				"galleryAuthToken":    "",
			},
			"website": "http://galleryproject.org/",
		},
		"Gambio": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "(?:<link[^>]* href=\"templates/gambio/|<a[^>]content\\.php\\?coID=\\d|<!-- gambio eof -->|<!--[\\s=]+Shopsoftware by Gambio GmbH \\(c\\))",
			"icon":    "Gambio.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"gambio": "",
			},
			"script":  "gm_javascript\\.js\\.php",
			"website": "http://gambio.de",
		},
		"Gatsby": map[string]interface{}{
			"cats": []interface{}{
				57,
				12,
			},
			"html": []interface{}{
				"<div id=\"___gatsby\">",
				"<style id=\"gatsby-inlined-css\">",
			},
			"meta": map[string]interface{}{
				"generator": "^Gatsby(?: ([0-9.]+))?$\\;version:\\1",
			},
			"icon": "Gatsby.svg",
			"implies": []interface{}{
				"React",
				"webpack",
			},
			"website": "https://www.gatsbyjs.org/",
		},
		"Gauges": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"cookies": map[string]interface{}{
				"_gauges_": "",
			},
			"icon": "Gauges.png",
			"js": map[string]interface{}{
				"_gauges": "",
			},
			"website": "https://get.gaug.es",
		},
		"Gentoo": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "gentoo",
			},
			"icon":    "Gentoo.png",
			"website": "http://www.gentoo.org",
		},
		"Gerrit": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"html": []interface{}{
				">Gerrit Code Review</a>\\s*\"\\s*\\(([0-9.]+)\\)\\;version:\\1",
				"<(?:div|style) id=\"gerrit_",
			},
			"icon": "gerrit.svg",
			"implies": []interface{}{
				"Java",
				"git",
			},
			"js": map[string]interface{}{
				"Gerrit":    "",
				"gerrit_ui": "",
			},
			"meta": map[string]interface{}{
				"title": "^Gerrit Code Review$",
			},
			"script":  "^gerrit_ui/gerrit_ui",
			"website": "http://www.gerritcodereview.com",
		},
		"Get Satisfaction": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"icon": "Get Satisfaction.png",
			"js": map[string]interface{}{
				"GSFN": "",
			},
			"website": "https://getsatisfaction.com/corp/",
		},
		"GetSimple CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "GetSimple CMS.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "GetSimple",
			},
			"website": "http://get-simple.info",
		},
		"Ghost": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"headers": map[string]interface{}{
				"X-Ghost-Cache-Status": "",
			},
			"icon":    "Ghost.png",
			"implies": "Node.js",
			"meta": map[string]interface{}{
				"generator": "Ghost(?:\\s([\\d.]+))?\\;version:\\1",
			},
			"website": "http://ghost.org",
		},
		"GitBook": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"icon": "GitBook.png",
			"meta": map[string]interface{}{
				"generator": "GitBook(?:.([\\d.]+))?\\;version:\\1",
			},
			"url":     "^https?://[^/]+\\.gitbook\\.com/",
			"website": "https://www.gitbook.com",
		},
		"GitHub Pages": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server":              "^GitHub\\.com$",
				"X-GitHub-Request-Id": "",
			},
			"icon":    "GitHub.svg",
			"implies": "Ruby on Rails",
			"url":     "^https?://[^/]+\\.github\\.io/",
			"website": "https://pages.github.com/",
		},
		"GitLab": map[string]interface{}{
			"cats": []interface{}{
				13,
				47,
			},
			"cookies": map[string]interface{}{
				"_gitlab_session": "",
			},
			"html": []interface{}{
				"<meta content=\"https?://[^/]+/assets/gitlab_logo-",
				"<header class=\"navbar navbar-fixed-top navbar-gitlab with-horizontal-nav\">",
			},
			"icon":    "GitLab.svg",
			"implies": "Ruby on Rails",
			"js": map[string]interface{}{
				"GitLab":              "",
				"gl.dashboardOptions": "",
			},
			"meta": map[string]interface{}{
				"og:site_name": "^GitLab$",
			},
			"website": "https://about.gitlab.com",
		},
		"GitLab CI": map[string]interface{}{
			"cats": []interface{}{
				44,
				47,
			},
			"icon":    "GitLab CI.png",
			"implies": "Ruby on Rails",
			"meta": map[string]interface{}{
				"description": "GitLab Continuous Integration",
			},
			"website": "http://about.gitlab.com/gitlab-ci",
		},
		"Gitea": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"cookies": map[string]interface{}{
				"i_like_gitea": "",
			},
			"html": []interface{}{
				"<div class=\"ui left\">\\n\\s+© Gitea Version: ([\\d.]+)\\;version:\\1",
			},
			"icon": "gitea.svg",
			"meta": map[string]interface{}{
				"keywords": "^go,git,self-hosted,gitea$",
			},
			"website": "https://gitea.io",
		},
		"Gitiles": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"html": "Powered by <a href=\"https://gerrit\\.googlesource\\.com/gitiles/\">Gitiles<",
			"implies": []interface{}{
				"Java",
				"git",
			},
			"website": "http://gerrit.googlesource.com/gitiles/",
		},
		"GlassFish": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "GlassFish(?: Server)?(?: Open Source Edition)?(?: ?/?([\\d.]+))?\\;version:\\1",
			},
			"icon": "GlassFish.png",
			"implies": []interface{}{
				"Java",
			},
			"website": "http://glassfish.java.net",
		},
		"Glyphicons": map[string]interface{}{
			"cats": []interface{}{
				17,
			},
			"html":    "(?:<link[^>]* href=[^>]+glyphicons(?:\\.min)?\\.css|<img[^>]* src=[^>]+glyphicons)",
			"icon":    "Glyphicons.png",
			"website": "http://glyphicons.com",
		},
		"Go": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"icon":    "Go.svg",
			"website": "https://golang.org",
		},
		"GoAhead": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "GoAhead",
			},
			"icon":    "GoAhead.png",
			"website": "http://embedthis.com/products/goahead/index.html",
		},
		"GoJS": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon":    "GoJS.png",
			"website": "https://gojs.net/",
			"js": map[string]interface{}{
				"go.version":     "(.*)\\;version:\\1",
				"go.GraphObject": "",
			},
		},
		"GoSquared": map[string]interface{}{
			"cats": []interface{}{
				10,
				52,
				53,
			},
			"icon": "gosquared.png",
			"js": map[string]interface{}{
				"_gs": "\\;confidence:30",
			},
			"website": "https://www.gosquared.com/",
		},
		"GoStats": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "GoStats.png",
			"js": map[string]interface{}{
				"_goStatsRun":   "",
				"_go_track_src": "",
				"go_msie":       "",
			},
			"website": "http://gostats.com/",
		},
		"Gogs": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"cookies": map[string]interface{}{
				"i_like_gogits": "",
			},
			"html": []interface{}{
				"<div class=\"ui left\">\\n\\s+© \\d{4} Gogs Version: ([\\d.]+) Page:\\;version:\\1",
				"<button class=\"ui basic clone button\" id=\"repo-clone-ssh\" data-link=\"gogs@",
			},
			"icon": "gogs.png",
			"meta": map[string]interface{}{
				"keywords": "go, git, self-hosted, gogs",
			},
			"script":  "js/gogs\\.js",
			"website": "http://gogs.io",
		},
		"Google AdSense": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Google AdSense.svg",
			"js": map[string]interface{}{
				"Goog_AdSense_":    "",
				"__google_ad_urls": "",
				"google_ad_":       "",
			},
			"script": []interface{}{
				"googlesyndication\\.com/",
				"ad\\.ca\\.doubleclick\\.net",
				"2mdn\\.net",
				"ad\\.ca\\.doubleclick\\.net",
			},
			"website": "https://www.google.fr/adsense/start/",
		},
		"Google Analytics": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"cookies": map[string]interface{}{
				"__utma": "",
				"_ga":    "",
				"_gat":   "",
			},
			"icon": "Google Analytics.svg",
			"html": "<amp-analytics [^>]*type=[\"']googleanalytics[\"']",
			"js": map[string]interface{}{
				"GoogleAnalyticsObject": "",
				"gaGlobal":              "",
			},
			"script":  "google-analytics\\.com\\/(?:ga|urchin|analytics)\\.js",
			"website": "http://google.com/analytics",
		},
		"Google Analytics Enhanced eCommerce": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "Google Analytics.svg",
			"implies": "Google Analytics",
			"js": map[string]interface{}{
				"gaplugins.EC": "",
			},
			"script":  "google-analytics\\.com\\/plugins\\/ua\\/(?:ec|ecommerce)\\.js",
			"website": "https://developers.google.com/analytics/devguides/collection/analyticsjs/enhanced-ecommerce",
		},
		"Google App Engine": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Google Frontend",
			},
			"icon":    "Google App Engine.png",
			"website": "http://code.google.com/appengine",
		},
		"Google Charts": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "Google Charts.png",
			"js": map[string]interface{}{
				"__googleVisualizationAbstractRendererElementsCount__": "",
				"__gvizguard__": "",
			},
			"website": "http://developers.google.com/chart/",
		},
		"Google Cloud": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Via": "^1\\.1 google$",
			},
			"icon":    "google_cloud.svg",
			"website": "https://cloud.google.com",
		},
		"Google Code Prettify": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon": "Google.svg",
			"js": map[string]interface{}{
				"prettyPrint": "",
			},
			"website": "http://code.google.com/p/google-code-prettify",
		},
		"Google Font API": map[string]interface{}{
			"cats": []interface{}{
				17,
			},
			"html": "<link[^>]* href=[^>]+fonts\\.(?:googleapis|google)\\.com",
			"icon": "Google Font API.png",
			"js": map[string]interface{}{
				"WebFonts": "",
			},
			"script":  "googleapis\\.com/.+webfont",
			"website": "http://google.com/fonts",
		},
		"Google Maps": map[string]interface{}{
			"cats": []interface{}{
				35,
			},
			"icon": "Google Maps.png",
			"script": []interface{}{
				"(?:maps\\.google\\.com/maps\\?file=api(?:&v=([\\d.]+))?|maps\\.google\\.com/maps/api/staticmap)\\;version:API v\\1",
				"//maps\\.googleapis\\.com/maps/api/js",
			},
			"website": "http://maps.google.com",
		},
		"Google PageSpeed": map[string]interface{}{
			"cats": []interface{}{
				23,
				33,
			},
			"headers": map[string]interface{}{
				"X-Mod-Pagespeed": "([\\d.]+)\\;version:\\1",
				"X-Page-Speed":    "(.+)\\;version:\\1",
			},
			"icon":    "Google PageSpeed.png",
			"website": "http://developers.google.com/speed/pagespeed/mod",
		},
		"Google Plus": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "Google Plus.svg",
			"script":  "apis\\.google\\.com/js/[a-z]*\\.js",
			"website": "http://plus.google.com",
		},
		"Google Sites": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Google Sites.png",
			"url":     "^https?://sites\\.google\\.com",
			"website": "http://sites.google.com",
		},
		"Google Tag Manager": map[string]interface{}{
			"cats": []interface{}{
				42,
			},
			"html": []interface{}{
				"googletagmanager\\.com/ns\\.html[^>]+></iframe>",
				"<!-- (?:End )?Google Tag Manager -->",
			},
			"icon": "Google Tag Manager.png",
			"js": map[string]interface{}{
				"google_tag_manager": "",
				"googletag":          "",
			},
			"website": "http://www.google.com/tagmanager",
		},
		"Google Wallet": map[string]interface{}{
			"cats": []interface{}{
				41,
			},
			"icon": "Google Wallet.png",
			"script": []interface{}{
				"checkout\\.google\\.com",
				"wallet\\.google\\.com",
			},
			"website": "http://wallet.google.com",
		},
		"Google Web Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "gws",
			},
			"icon":    "Google.svg",
			"website": "http://en.wikipedia.org/wiki/Google_Web_Server",
		},
		"Google Web Toolkit": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"icon":    "Google Web Toolkit.png",
			"implies": "Java",
			"js": map[string]interface{}{
				"__gwt_": "",
			},
			"meta": map[string]interface{}{
				"gwt:property": "",
			},
			"website": "http://developers.google.com/web-toolkit",
		},
		"Graffiti CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"graffitibot": "",
			},
			"icon":    "Graffiti CMS.png",
			"implies": "Microsoft ASP.NET",
			"meta": map[string]interface{}{
				"generator": "Graffiti CMS ([^\"]+)\\;version:\\1",
			},
			"script":  "/graffiti\\.js",
			"website": "http://graffiticms.codeplex.com",
		},
		"Grav": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Grav.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "GravCMS(?:\\s([\\d.]+))?\\;version:\\1",
			},
			"website": "http://getgrav.org",
		},
		"Gravatar": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html": "<[^>]+gravatar\\.com/avatar/",
			"icon": "Gravatar.png",
			"js": map[string]interface{}{
				"Gravatar": "",
			},
			"website": "http://gravatar.com",
		},
		"Gravity Forms": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html": []interface{}{
				"<div class=(?:\"|')[^>]*gform_wrapper",
				"<div class=(?:\"|')[^>]*gform_body",
				"<ul [^>]*class=(?:\"|')[^>]*gform_fields",
				"<link [^>]*href=(?:\"|')[^>]*wp-content/plugins/gravityforms/css/",
			},
			"icon":    "gravityforms.svg",
			"implies": "WordPress",
			"script":  "/wp-content/plugins/gravityforms/js/[^/]+\\.js\\?ver=([\\d.]+)$\\;version:\\1",
			"website": "http://gravityforms.com",
		},
		"Green Valley CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<img[^>]+/dsresource\\?objectid=",
			"icon":    "Green Valley CMS.png",
			"implies": "Apache Tomcat",
			"meta": map[string]interface{}{
				"DC.identifier": "/content\\.jsp\\?objectid=",
			},
			"website": "http://www.greenvalley.nl/Public/Producten/Content_Management/CMS",
		},
		"Gridsome": map[string]interface{}{
			"cats": []interface{}{
				57,
			},
			"icon":    "Gridsome.svg",
			"implies": "Vue.js",
			"meta": map[string]interface{}{
				"generator": "^Gridsome v([\\d.]+)$\\;version:\\1",
			},
			"website": "https://gridsome.org",
		},
		"GrowingIO": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"js": map[string]interface{}{
				"gio": "",
			},
			"cookies": map[string]interface{}{
				"grwng_uid":  "",
				"gr_user_id": "",
			},
			"icon":    "GrowingIO.png",
			"script":  "assets\\.growingio\\.com/([\\d.]+)/gio.js\\;version:\\1",
			"website": "https://www.growingio.com/",
		},
		"HERE": map[string]interface{}{
			"cats": []interface{}{
				35,
			},
			"icon":    "HERE.png",
			"script":  "https?://js\\.cit\\.api\\.here\\.com/se/([\\d.]+)\\/\\;version:\\1",
			"website": "http://developer.here.com",
		},
		"HHVM": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "HHVM/?([\\d.]+)?\\;version:\\1",
			},
			"icon":    "HHVM.png",
			"implies": "PHP\\;confidence:75",
			"website": "http://hhvm.com",
		},
		"HP ChaiServer": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "HP-Chai(?:Server|SOE)(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "HP.svg",
			"website": "http://hp.com",
		},
		"HP Compact Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "HP_Compact_Server(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "HP.svg",
			"website": "http://hp.com",
		},
		"HP ProCurve": map[string]interface{}{
			"cats": []interface{}{
				37,
			},
			"icon":    "HP.svg",
			"website": "http://hp.com/networking",
		},
		"HP System Management": map[string]interface{}{
			"cats": []interface{}{
				46,
			},
			"headers": map[string]interface{}{
				"Server": "HP System Management",
			},
			"icon":    "HP.svg",
			"website": "http://hp.com",
		},
		"HP iLO": map[string]interface{}{
			"cats": []interface{}{
				22,
				46,
			},
			"headers": map[string]interface{}{
				"Server": "HP-iLO-Server(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "HP.svg",
			"website": "http://hp.com",
		},
		"HTTP/2": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"excludes": "SPDY",
			"headers": map[string]interface{}{
				"X-Firefox-Spdy": "h2",
			},
			"icon":    "http2.png",
			"website": "https://http2.github.io",
		},
		"Haddock": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"html":    "<p>Produced by <a href=\"http://www\\.haskell\\.org/haddock/\">Haddock</a> version ([0-9.]+)</p>\\;version:\\1",
			"script":  "haddock-util\\.js",
			"website": "http://www.haskell.org/haddock/",
		},
		"Hammer.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "Hammer.js.png",
			"js": map[string]interface{}{
				"Ha.VERSION":     "^(.+)$\\;version:\\1",
				"Hammer":         "",
				"Hammer.VERSION": "^(.+)$\\;version:\\1",
			},
			"script":  "hammer(?:\\.min)?\\.js",
			"website": "https://hammerjs.github.io",
		},
		"Handlebars": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html": "<[^>]*type=[^>]text\\/x-handlebars-template",
			"icon": "Handlebars.png",
			"js": map[string]interface{}{
				"Handlebars":         "",
				"Handlebars.VERSION": "^(.+)$\\;version:\\1",
			},
			"script":  "handlebars(?:\\.runtime)?(?:-v([\\d.]+?))?(?:\\.min)?\\.js\\;version:\\1",
			"website": "http://handlebarsjs.com",
		},
		"Haravan": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Haravan.png",
			"js": map[string]interface{}{
				"Haravan": "",
			},
			"script":  "haravan.*\\.js",
			"website": "https://www.haravan.com",
		},
		"Haskell": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"icon":    "Haskell.png",
			"website": "http://wiki.haskell.org/Haskell",
		},
		"HeadJS": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"html": "<[^>]*data-headjs-load",
			"icon": "HeadJS.png",
			"js": map[string]interface{}{
				"head.browser.name": "",
			},
			"script":  "head\\.(?:core|load)(?:\\.min)?\\.js",
			"website": "http://headjs.com",
		},
		"Heap": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Heap.png",
			"js": map[string]interface{}{
				"heap": "",
			},
			"script":  "heap-\\d+\\.js",
			"website": "http://heapanalytics.com",
		},
		"Hello Bar": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon": "Hello Bar.png",
			"js": map[string]interface{}{
				"HelloBar": "",
			},
			"script":  "hellobar\\.js",
			"website": "http://hellobar.com",
		},
		"Hexo": map[string]interface{}{
			"cats": []interface{}{
				57,
			},
			"html": []interface{}{
				"Powered by <a href=\"https?://hexo\\.io/?\"[^>]*>Hexo</",
			},
			"icon": "Hexo.png",
			"meta": map[string]interface{}{
				"generator": "Hexo(?: v?([\\d.]+))?\\;version:\\1",
			},
			"website": "https://hexo.io",
		},
		"Hiawatha": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Hiawatha v([\\d.]+)\\;version:\\1",
			},
			"icon":    "Hiawatha.png",
			"website": "http://hiawatha-webserver.org",
		},
		"Highcharts": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"html": "<svg[^>]*><desc>Created with Highcharts ([\\d.]*)\\;version:\\1",
			"icon": "Highcharts.png",
			"js": map[string]interface{}{
				"Highcharts":         "",
				"Highcharts.version": "^(.+)$\\;version:\\1",
			},
			"script":  "highcharts.*\\.js",
			"website": "https://www.highcharts.com",
		},
		"Highlight.js": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon": "Highlight.js.png",
			"js": map[string]interface{}{
				"hljs.highlightBlock": "",
				"hljs.listLanguages":  "",
			},
			"script":  "/(?:([\\d.])+/)?highlight(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://highlightjs.org/",
		},
		"Highstock": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"html":    "<svg[^>]*><desc>Created with Highstock ([\\d.]*)\\;version:\\1",
			"icon":    "Highcharts.png",
			"script":  "highstock[.-]?([\\d\\.]*\\d).*\\.js\\;version:\\1",
			"website": "http://highcharts.com/products/highstock",
		},
		"Hinza Advanced CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"icon":    "hinza_advanced_cms.svg",
			"implies": "Laravel",
			"meta": map[string]interface{}{
				"generator": "hinzacms",
			},
			"website": "http://hinzaco.com",
		},
		"Bloomreach": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<[^>]+/binaries/(?:[^/]+/)*content/gallery/",
			"icon":    "Bloomreach.png",
			"website": "https://developers.bloomreach.com",
		},
		"Hogan.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "Hogan.js.png",
			"js": map[string]interface{}{
				"Hogan": "",
			},
			"script": []interface{}{
				"hogan-[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"([\\d.]+)/hogan(?:\\.min)?\\.js\\;version:\\1",
			},
			"website": "https://twitter.github.io/hogan.js/",
		},
		"Homeland": map[string]interface{}{
			"cats": []interface{}{
				1,
				2,
			},
			"cookies": map[string]interface{}{
				"_homeland_": "",
			},
			"icon":    "Homeland.png",
			"implies": "Ruby on Rails",
			"website": "https://gethomeland.com",
		},
		"Hotaru CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"hotaru_mobile": "",
			},
			"icon":    "Hotaru CMS.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Hotaru CMS",
			},
			"website": "http://hotarucms.org",
		},
		"Hotjar": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Hotjar.png",
			"js": map[string]interface{}{
				"HotLeadfactory":    "",
				"HotleadController": "",
				"hj.apiUrlBase":     "",
			},
			"script":  "^//static\\.hotjar\\.com/c/hotjar-",
			"website": "https://www.hotjar.com",
		},
		"HubSpot": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"html": "<!-- Start of Async HubSpot",
			"icon": "HubSpot.png",
			"js": map[string]interface{}{
				"_hsq":    "",
				"hubspot": "",
			},
			"website": "https://www.hubspot.com",
		},
		"Hugo": map[string]interface{}{
			"cats": []interface{}{
				57,
			},
			"html": "powered by <a [^>]*href=\"http://hugo.spf13.com",
			"icon": "Hugo.png",
			"meta": map[string]interface{}{
				"generator": "Hugo ([\\d.]+)?\\;version:\\1",
			},
			"website": "http://gohugo.io",
		},
		"Hybris": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"_hybris": "",
			},
			"html":    "<[^>]+/(?:sys_master|hybr|_ui/(?:responsive/)?(?:desktop|common(?:/images|/img)?))/",
			"icon":    "Hybris.png",
			"implies": "Java",
			"website": "https://hybris.com",
		},
		"IBM Coremetrics": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "IBM.svg",
			"script":  "cmdatatagutils\\.js",
			"website": "http://ibm.com/software/marketing-solutions/coremetrics",
		},
		"IBM DataPower": map[string]interface{}{
			"cats": []interface{}{
				64,
			},
			"headers": map[string]interface{}{
				"X-Backside-Transport":    "",
				"X-Global-Transaction-ID": "",
			},
			"icon":    "DataPower.png",
			"website": "https://www.ibm.com/products/datapower-gateway",
		},
		"IBM HTTP Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "IBM_HTTP_Server(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "IBM.svg",
			"website": "http://ibm.com/software/webservers/httpservers",
		},
		"IBM Tivoli Storage Manager": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "TSM_HTTP(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "IBM.svg",
			"website": "http://ibm.com",
		},
		"IBM WebSphere Commerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "href=\"(?:\\/|[^>]+)webapp\\/wcs\\/",
			"icon":    "IBM.svg",
			"implies": "Java",
			"url":     "/wcs/",
			"website": "http://ibm.com/software/genservers/commerceproductline",
		},
		"IBM WebSphere Portal": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"IBM-Web2-Location":       "",
				"Itx-Generated-Timestamp": "",
			},
			"icon":    "IBM.svg",
			"implies": "Java",
			"url":     "/wps/",
			"website": "http://ibm.com/software/websphere/portal",
		},
		"IIS": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^(?:Microsoft-)?IIS(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "IIS.png",
			"implies": "Windows Server",
			"website": "http://www.iis.net",
		},
		"INFOnline": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "INFOnline.png",
			"js": map[string]interface{}{
				"iam_data": "",
				"szmvars":  "",
			},
			"script":  "^https?://(?:[^/]+\\.)?i(?:oam|v)wbox\\.de/",
			"website": "https://www.infonline.de",
		},
		"INTI": map[string]interface{}{
			"cats": []interface{}{
				6,
				53,
			},
			"icon":    "byINTI.svg",
			"url":     "\\.byinti\\.com",
			"website": "https://byinti.com",
		},
		"IPB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"cookies": map[string]interface{}{
				"ipbWWLmodpids":    "",
				"ipbWWLsession_id": "",
			},
			"html": "<link[^>]+ipb_[^>]+\\.css",
			"icon": "IPB.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"js": map[string]interface{}{
				"IPBoard":     "",
				"ipb_var":     "",
				"ipsSettings": "",
			},
			"script":  "jscripts/ips_",
			"website": "https://invisioncommunity.com/",
		},
		"Ideasoft": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Ideasoft.png",
			"script": []interface{}{
				"\\.myideasoft\\.com/",
			},
			"website": "https://www.ideasoft.com",
		},
		"IdoSell Shop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "idosellshop.png",
			"js": map[string]interface{}{
				"IAI_Ajax": "",
			},
			"website": "https://www.idosell.com",
		},
		"Immutable.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "Immutable.js.png",
			"js": map[string]interface{}{
				"Immutable":         "",
				"Immutable.version": "^(.+)$\\;version:\\1",
			},
			"script":  "^immutable\\.(?:min\\.)?js$",
			"website": "https://facebook.github.io/immutable-js/",
		},
		"ImpressCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"ICMSSession": "",
				"ImpressCMS":  "",
			},
			"icon":    "ImpressCMS.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "ImpressCMS",
			},
			"script":  "include/linkexternal\\.js",
			"website": "http://www.impresscms.org",
		},
		"ImpressPages": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "ImpressPages.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "ImpressPages(?: CMS)?( [\\d.]*)?\\;version:\\1",
			},
			"website": "http://impresspages.org",
		},
		"InProces": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<!-- CSS InProces Portaal default -->",
			"icon":    "InProces.png",
			"script":  "brein/inproces/website/websitefuncties\\.js",
			"website": "http://www.brein.nl/oplossing/product/website",
		},
		"Incapsula": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"X-CDN": "Incapsula",
			},
			"icon":    "Incapsula.png",
			"website": "http://www.incapsula.com",
		},
		"Includable": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Includable-Version": "",
			},
			"icon":    "Includable.svg",
			"website": "http://includable.com",
		},
		"Indexhibit": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<(?:link|a href) [^>]+ndxz-studio",
			"implies": []interface{}{
				"PHP",
				"Apache",
				"Exhibit",
			},
			"meta": map[string]interface{}{
				"generator": "Indexhibit",
			},
			"website": "http://www.indexhibit.org",
		},
		"Indico": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"MAKACSESSION": "",
			},
			"html":    "Powered by\\s+(?:CERN )?<a href=\"http://(?:cdsware\\.cern\\.ch/indico/|indico-software\\.org|cern\\.ch/indico)\">(?:CDS )?Indico( [\\d\\.]+)?\\;version:\\1",
			"icon":    "Indico.png",
			"website": "http://indico-software.org",
		},
		"Indy": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Indy(?:/([\\d.]+))?\\;version:\\1",
			},
			"website": "http://indyproject.org",
		},
		"InfernoJS": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "InfernoJS.png",
			"js": map[string]interface{}{
				"Inferno":         "",
				"Inferno.version": "^(.+)$\\;version:\\1",
			},
			"website": "https://infernojs.org",
		},
		"Infusionsoft": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"html": []interface{}{
				"<input [^>]*name=\"infusionsoft_version\" [^>]*value=\"([^>]*)\" [^>]*\\/>\\;version:\\1",
				"<input [^>]*value=\"([^>]*)\" [^>]*name=\"infusionsoft_version\" [^>]*\\/>\\;version:\\1",
			},
			"icon":    "infusionsoft.svg",
			"website": "http://infusionsoft.com",
		},
		"Inspectlet": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": []interface{}{
				"<!-- (?:Begin|End) Inspectlet Embed Code -->",
			},
			"icon": "inspectlet.png",
			"js": map[string]interface{}{
				"__insp":   "",
				"__inspld": "",
			},
			"script": []interface{}{
				"cdn\\.inspectlet\\.com",
			},
			"website": "https://www.inspectlet.com/",
		},
		"Instabot": map[string]interface{}{
			"cats": []interface{}{
				5,
				10,
				32,
				52,
				58,
			},
			"icon": "Instabot.png",
			"js": map[string]interface{}{
				"Instabot": "",
			},
			"script":  "/rokoInstabot\\.js",
			"website": "https://instabot.io/",
		},
		"InstantCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"InstantCMS[logdate]": "",
			},
			"icon":    "InstantCMS.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "InstantCMS",
			},
			"website": "http://www.instantcms.ru",
		},
		"Intel Active Management Technology": map[string]interface{}{
			"cats": []interface{}{
				22,
				46,
			},
			"headers": map[string]interface{}{
				"Server": "Intel\\(R\\) Active Management Technology(?: ([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Intel Active Management Technology.png",
			"website": "http://intel.com",
		},
		"IntenseDebate": map[string]interface{}{
			"cats": []interface{}{
				15,
			},
			"icon":    "IntenseDebate.png",
			"script":  "intensedebate\\.com",
			"website": "http://intensedebate.com",
		},
		"Intercom": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Intercom.png",
			"js": map[string]interface{}{
				"Intercom": "",
			},
			"script":  "(?:api\\.intercom\\.io/api|static\\.intercomcdn\\.com/intercom\\.v1)",
			"website": "https://www.intercom.com",
		},
		"Intershop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "Intershop.png",
			"script":  "(?:is-bin|INTERSHOP)",
			"website": "http://intershop.com",
		},
		"Invenio": map[string]interface{}{
			"cats": []interface{}{
				50,
			},
			"cookies": map[string]interface{}{
				"INVENIOSESSION": "",
			},
			"html":    "(?:Powered by|System)\\s+(?:CERN )?<a (?:class=\"footer\" )?href=\"http://(?:cdsware\\.cern\\.ch(?:/invenio)?|invenio-software\\.org|cern\\.ch/invenio)(?:/)?\">(?:CDS )?Invenio</a>\\s*v?([\\d\\.]+)?\\;version:\\1",
			"icon":    "Invenio.png",
			"website": "http://invenio-software.org",
		},
		"Inwemo": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon": "inwemo.png",
			"js": map[string]interface{}{
				"Inwemo": "",
			},
			"script":  "https?://cdn\\.inwemo\\.com/inwemo\\.min\\.js",
			"website": "https://inwemo.com/",
		},
		"Ionic": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"icon":    "ionic.png",
			"implies": "Angular",
			"js": map[string]interface{}{
				"Ionic.config":  "",
				"Ionic.version": "^(.+)$\\;version:\\1",
			},
			"website": "https://ionicframework.com",
		},
		"Ionicons": map[string]interface{}{
			"cats": []interface{}{
				17,
			},
			"html":    "<link[^>]* href=[^>]+ionicons(?:\\.min)?\\.css",
			"icon":    "Ionicons.png",
			"website": "http://ionicons.com",
		},
		"JAlbum": map[string]interface{}{
			"cats": []interface{}{
				7,
			},
			"icon":    "JAlbum.png",
			"implies": "Java",
			"meta": map[string]interface{}{
				"generator": "JAlbum( [\\d.]+)?\\;version:\\1",
			},
			"website": "http://jalbum.net/en",
		},
		"JBoss Application Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "JBoss(?:-([\\d.]+))?\\;version:\\1",
			},
			"icon":    "JBoss Application Server.png",
			"website": "http://jboss.org/jbossas.html",
		},
		"JBoss Web": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"excludes": "Apache Tomcat",
			"headers": map[string]interface{}{
				"X-Powered-By": "JBossWeb(?:-([\\d.]+))?\\;version:\\1",
			},
			"icon":    "JBoss Web.png",
			"implies": "JBoss Application Server",
			"website": "http://jboss.org/jbossweb",
		},
		"JET Enterprise": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"powered": "jet-enterprise",
			},
			"icon":    "JET Enterprise.svg",
			"website": "http://www.jetecommerce.com.br/",
		},
		"JS Charts": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "JS Charts.png",
			"js": map[string]interface{}{
				"JSChart": "",
			},
			"script":  "jscharts.*\\.js",
			"website": "http://www.jscharts.com",
		},
		"JSEcoin": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon": "JSEcoin.png",
			"js": map[string]interface{}{
				"jseMine": "",
			},
			"script":  "^(?:https):?//load\\.jsecoin\\.com/load/",
			"website": "https://jsecoin.com/",
		},
		"JTL Shop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"JTLSHOP": "",
			},
			"html":    "(?:<input[^>]+name=\"JTLSHOP|<a href=\"jtl\\.php)",
			"icon":    "JTL Shop.png",
			"website": "http://www.jtl-software.de/produkte/jtl-shop3",
		},
		"Jahia DX": map[string]interface{}{
			"cats": []interface{}{
				1,
				47,
			},
			"html":    "<script id=\"staticAssetAggregatedJavascrip",
			"icon":    "JahiaDX.svg",
			"website": "http://www.jahia.com/dx",
		},
		"Jalios": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Jalios.png",
			"meta": map[string]interface{}{
				"generator": "Jalios",
			},
			"website": "http://www.jalios.com",
		},
		"Java": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"cookies": map[string]interface{}{
				"JSESSIONID": "",
			},
			"icon":    "Java.png",
			"website": "http://java.com",
		},
		"Java Servlet": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Servlet(?:.([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Java.png",
			"implies": "Java",
			"website": "http://www.oracle.com/technetwork/java/index-jsp-135475.html",
		},
		"JavaScript Infovis Toolkit": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "JavaScript Infovis Toolkit.png",
			"js": map[string]interface{}{
				"$jit":         "",
				"$jit.version": "^(.+)$\\;version:\\1",
			},
			"script":  "jit(?:-yc)?\\.js",
			"website": "https://philogb.github.io/jit/",
		},
		"JavaServer Faces": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "JSF(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "JavaServer Faces.png",
			"implies": "Java",
			"website": "http://javaserverfaces.java.net",
		},
		"JavaServer Pages": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "JSP(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Java.png",
			"implies": "Java",
			"website": "http://www.oracle.com/technetwork/java/javaee/jsp/index.html",
		},
		"Jekyll": map[string]interface{}{
			"cats": []interface{}{
				57,
			},
			"html": []interface{}{
				"Powered by <a href=\"https?://jekyllrb\\.com\"[^>]*>Jekyll</",
				"<!-- Created with Jekyll Now -",
				"<!-- Begin Jekyll SEO tag",
			},
			"icon": "Jekyll.png",
			"meta": map[string]interface{}{
				"generator": "Jekyll (v[\\d.]+)?\\;version:\\1",
			},
			"website": "http://jekyllrb.com",
		},
		"Jenkins": map[string]interface{}{
			"cats": []interface{}{
				44,
			},
			"headers": map[string]interface{}{
				"X-Jenkins": "([\\d.]+)\\;version:\\1",
			},
			"html":    "<span class=\"jenkins_ver\"><a href=\"https://jenkins\\.io/\">Jenkins ver\\. ([\\d.]+)\\;version:\\1",
			"icon":    "Jenkins.png",
			"implies": "Java",
			"js": map[string]interface{}{
				"jenkinsCIGlobal": "",
				"jenkinsRules":    "",
			},
			"website": "https://jenkins.io/",
		},
		"Jetshop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<(?:div|aside) id=\"jetshop-branding\">",
			"icon": "Jetshop.png",
			"js": map[string]interface{}{
				"JetshopData": "",
			},
			"website": "http://jetshop.se",
		},
		"Jetty": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Jetty(?:\\(([\\d\\.]*\\d+))?\\;version:\\1",
			},
			"icon":    "Jetty.png",
			"implies": "Java",
			"website": "http://www.eclipse.org/jetty",
		},
		"Jimdo": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Jimdo-Instance": "",
				"X-Jimdo-Wid":      "",
			},
			"icon": "jimdo.png",
			"url":  "\\.jimdo\\.com/",
			"js": map[string]interface{}{
				"jimdoData":  "",
				"jimdo_Data": "",
			},
			"website": "https://www.jimdo.com",
		},
		"Jirafe": map[string]interface{}{
			"cats": []interface{}{
				10,
				32,
			},
			"icon": "Jirafe.png",
			"js": map[string]interface{}{
				"jirafe": "",
			},
			"script":  "/jirafe\\.js",
			"website": "https://docs.jirafe.com",
		},
		"Jive": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"headers": map[string]interface{}{
				"X-JIVE-USER-ID":        "",
				"X-JSL":                 "",
				"X-Jive-Flow-Id":        "",
				"X-Jive-Request-Id":     "",
				"x-jive-chrome-wrapped": "",
			},
			"icon":    "Jive.png",
			"website": "http://www.jivesoftware.com",
		},
		"JobberBase": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon":    "JobberBase.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"Jobber": "",
			},
			"meta": map[string]interface{}{
				"generator": "Jobberbase",
			},
			"website": "http://www.jobberbase.com",
		},
		"Joomla": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Content-Encoded-By": "Joomla! ([\\d.]+)\\;version:\\1",
			},
			"html":    "(?:<div[^>]+id=\"wrapper_r\"|<(?:link|script)[^>]+(?:feed|components)/com_|<table[^>]+class=\"pill)\\;confidence:50",
			"icon":    "Joomla.svg",
			"implies": "PHP",
			"js": map[string]interface{}{
				"Joomla":    "",
				"jcomments": "",
			},
			"meta": map[string]interface{}{
				"generator": "Joomla!(?: ([\\d.]+))?\\;version:\\1",
			},
			"url":     "option=com_",
			"website": "https://www.joomla.org",
		},
		"K2": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html":    "<!--(?: JoomlaWorks \"K2\"| Start K2)",
			"icon":    "K2.png",
			"implies": "Joomla",
			"js": map[string]interface{}{
				"K2RatingURL": "",
			},
			"website": "https://getk2.org",
		},
		"KISSmetrics": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "KISSmetrics.png",
			"js": map[string]interface{}{
				"KM_COOKIE_DOMAIN": "",
			},
			"website": "https://www.kissmetrics.com",
		},
		"Kajabi": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"_kjb_session": "",
			},
			"icon": "Kajabi.svg",
			"js": map[string]interface{}{
				"Kajabi": "",
			},
			"website": "https://newkajabi.com",
		},
		"Kampyle": map[string]interface{}{
			"cats": []interface{}{
				10,
				13,
			},
			"cookies": map[string]interface{}{
				"k_visit": "",
			},
			"icon": "Kampyle.png",
			"js": map[string]interface{}{
				"KAMPYLE_COMMON": "",
				"k_track":        "",
				"kampyle":        "",
			},
			"script":  "cf\\.kampyle\\.com/k_button\\.js",
			"website": "http://www.kampyle.com",
		},
		"Kamva": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Kamva.svg",
			"js": map[string]interface{}{
				"Kamva": "",
			},
			"meta": map[string]interface{}{
				"generator": "[CK]amva",
			},
			"script":  "cdn\\.mykamva\\.ir",
			"website": "https://kamva.ir",
		},
		"Kemal": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Kemal",
			},
			"icon":    "kemalcr.png",
			"website": "http://kemalcr.com",
		},
		"Kendo UI": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "<link[^>]*\\s+href=[^>]*styles/kendo\\.common(?:\\.min)?\\.css[^>]*/>",
			"icon":    "Kendo UI.png",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"kendo":         "",
				"kendo.version": "^(.+)$\\;version:\\1",
			},
			"website": "https://www.telerik.com/kendo-ui",
		},
		"Kentico CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"CMSPreferredCulture": "",
			},
			"icon": "Kentico CMS.png",
			"meta": map[string]interface{}{
				"generator": "Kentico CMS ([\\d.R]+ \\(build [\\d.]+\\))\\;version:\\1",
			},
			"website": "http://www.kentico.com",
		},
		"Kestrel": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^Kestrel",
			},
			"implies": []interface{}{
				"Microsoft ASP.NET",
			},
			"icon":    "kestrel.svg",
			"website": "https://docs.microsoft.com/en-us/aspnet/core/fundamentals/servers/kestrel",
		},
		"KeyCDN": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server": "^keycdn-engine$",
			},
			"icon":    "KeyCDN.png",
			"website": "http://www.keycdn.com",
		},
		"Kibana": map[string]interface{}{
			"cats": []interface{}{
				29,
				25,
			},
			"headers": map[string]interface{}{
				"kbn-name":    "kibana",
				"kbn-version": "^([\\d.]+)$\\;version:\\1",
			},
			"html":    "<title>Kibana</title>",
			"icon":    "kibana.svg",
			"implies": "Node.js",
			"url":     "kibana#/dashboard/",
			"website": "http://www.elastic.co/products/kibana",
		},
		"KineticJS": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "KineticJS.png",
			"js": map[string]interface{}{
				"Kinetic":         "",
				"Kinetic.version": "^(.+)$\\;version:\\1",
			},
			"script":  "kinetic(?:-v?([\\d.]+))?(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://github.com/ericdrowell/KineticJS/",
		},
		"Klarna Checkout": map[string]interface{}{
			"cats": []interface{}{
				41,
				6,
				5,
			},
			"icon": "Klarna.svg",
			"js": map[string]interface{}{
				"_klarnaCheckout": "",
			},
			"website": "https://www.klarna.com/international/",
		},
		"Knockout.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "Knockout.js.png",
			"js": map[string]interface{}{
				"ko.version": "^(.+)$\\;version:\\1",
			},
			"website": "http://knockoutjs.com",
		},
		"Koa": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^koa$",
			},
			"icon":    "Koa.png",
			"implies": "Node.js",
			"website": "http://koajs.com",
		},
		"Koala Framework": map[string]interface{}{
			"cats": []interface{}{
				1,
				18,
			},
			"html":    "<!--[^>]+This website is powered by Koala Web Framework CMS",
			"icon":    "Koala Framework.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "^Koala Web Framework CMS",
			},
			"website": "http://koala-framework.org",
		},
		"KobiMaster": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "Kobimaster.png",
			"implies": "Microsoft ASP.NET",
			"js": map[string]interface{}{
				"kmGetSession": "",
				"kmPageInfo":   "",
			},
			"website": "https://www.kobimaster.com.tr",
		},
		"Koha": map[string]interface{}{
			"cats": []interface{}{
				21,
			},
			"html": []interface{}{
				"<input name=\"koha_login_context\" value=\"intranet\" type=\"hidden\">",
				"<a href=\"/cgi-bin/koha/",
			},
			"icon":    "koha.png",
			"implies": "Perl",
			"js": map[string]interface{}{
				"KOHA": "",
			},
			"meta": map[string]interface{}{
				"generator": "^Koha ([\\d.]+)$\\;version:\\1",
			},
			"website": "https://koha-community.org/",
		},
		"Kohana": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"kohanasession": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Kohana Framework ([\\d.]+)\\;version:\\1",
			},
			"icon":    "Kohana.png",
			"implies": "PHP",
			"website": "http://kohanaframework.org",
		},
		"Koken": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"koken_referrer": "",
			},
			"html": []interface{}{
				"<html lang=\"en\" class=\"k-source-essays k-lens-essays\">",
				"<!--\\s+KOKEN DEBUGGING",
			},
			"icon": "Koken.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"meta": map[string]interface{}{
				"generator": "Koken ([\\d.]+)\\;version:\\1",
			},
			"script":  "koken(?:\\.js\\?([\\d.]+)|/storage)\\;version:\\1",
			"website": "http://koken.me",
		},
		"Kolibri CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Kolibri",
			},
			"meta": map[string]interface{}{
				"generator": "Kolibri",
			},
			"website": "http://alias.io",
		},
		"Komodo CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Komodo CMS.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "^Komodo CMS",
			},
			"website": "http://www.komodocms.com",
		},
		"Kontaktify": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "Kontaktify.png",
			"script":  "//(?:www\\.)?kontaktify\\.com/embed\\.js",
			"website": "https://www.kontaktify.com",
		},
		"Koobi": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<!--[^K>-]+Koobi ([a-z\\d.]+)\\;version:\\1",
			"icon": "Koobi.png",
			"meta": map[string]interface{}{
				"generator": "Koobi",
			},
			"website": "http://dream4.de/cms",
		},
		"Kooboo CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-KoobooCMS-Version": "^(.+)$\\;version:\\1",
			},
			"icon":    "Kooboo CMS.png",
			"implies": "Microsoft ASP.NET",
			"script":  "/Kooboo",
			"website": "http://kooboo.com",
		},
		"Kotisivukone": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Kotisivukone.png",
			"script":  "kotisivukone(?:\\.min)?\\.js",
			"website": "http://www.kotisivukone.fi",
		},
		"Kubernetes Dashboard": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"html":    "<html ng-app=\"kubernetesDashboard\">",
			"icon":    "Kubernetes.svg",
			"website": "https://kubernetes.io/",
		},
		"LEPTON": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "LEPTON.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "LEPTON",
			},
			"website": "http://www.lepton-cms.org",
		},
		"LabVIEW": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "LabVIEW(?:/([\\d\\.]+))?\\;version:\\1",
			},
			"icon":    "LabVIEW.png",
			"website": "http://ni.com/labview",
		},
		"Laravel": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"laravel_session": "",
			},
			"icon":    "Laravel.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"Laravel": "",
			},
			"website": "https://laravel.com",
		},
		"Laterpay": map[string]interface{}{
			"cats": []interface{}{
				41,
			},
			"icon": "laterpay.png",
			"meta": map[string]interface{}{
				"laterpay:connector:callbacks:on_user_has_access": "deobfuscateText",
			},
			"script":  "https?://connectormwi\\.laterpay\\.net/([0-9.]+)[a-zA-z-]*/live/[\\w-]+\\.js\\;version:\\1",
			"website": "https://www.laterpay.net/",
		},
		"Lazy.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"script":  "lazy(?:\\.browser)?(?:\\.min)?\\.js",
			"website": "http://danieltao.com/lazy.js",
		},
		"Leaflet": map[string]interface{}{
			"cats": []interface{}{
				35,
			},
			"icon": "Leaflet.png",
			"js": map[string]interface{}{
				"L.DistanceGrid": "",
				"L.PosAnimation": "",
				"L.version":      "^(.+)$\\;version:\\1\\;confidence:0",
			},
			"script":  "leaflet.*\\.js",
			"website": "http://leafletjs.com",
		},
		"Less": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html":    "<link[^>]+ rel=\"stylesheet/less\"",
			"icon":    "Less.png",
			"website": "http://lesscss.org",
		},
		"Liferay": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"Liferay-Portal": "[a-z\\s]+([\\d.]+)\\;version:\\1",
			},
			"icon": "Liferay.png",
			"js": map[string]interface{}{
				"Liferay": "",
			},
			"website": "https://www.liferay.com",
		},
		"Lift": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Lift-Version": "(.+)\\;version:\\1",
			},
			"icon":    "Lift.png",
			"implies": "Scala",
			"website": "http://liftweb.net",
		},
		"LightMon Engine": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"lm_online": "",
			},
			"html":    "<!-- Lightmon Engine Copyright Lightmon",
			"icon":    "LightMon Engine.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "LightMon Engine",
			},
			"website": "http://lightmon.ru",
		},
		"Lightbox": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"html":    "<link [^>]*href=\"[^\"]+lightbox(?:\\.min)?\\.css",
			"icon":    "Lightbox.png",
			"script":  "lightbox.*\\.js",
			"website": "http://lokeshdhakar.com/projects/lightbox2/",
		},
		"Lightspeed eCom": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<!-- \\[START\\] 'blocks/head\\.rain' -->",
			"icon":    "Lightspeed.svg",
			"script":  "http://assets\\.webshopapp\\.com",
			"url":     "seoshop.webshopapp.com",
			"website": "http://www.lightspeedhq.com/products/ecommerce/",
		},
		"Lighty": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"lighty_version": "",
			},
			"icon":    "Lighty.png",
			"implies": "PHP",
			"website": "https://gitlab.com/lighty/framework",
		},
		"LimeSurvey": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"headers": map[string]interface{}{
				"generator": "LimeSurvey",
			},
			"icon":    "LimeSurvey.png",
			"website": "http://limesurvey.org/",
		},
		"LinkSmart": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "LinkSmart.png",
			"js": map[string]interface{}{
				"LS_JSON":       "",
				"LinkSmart":     "",
				"_mb_site_guid": "",
			},
			"script":  "^https?://cdn\\.linksmart\\.com/linksmart_([\\d.]+?)(?:\\.min)?\\.js\\;version:\\1",
			"website": "http://linksmart.com",
		},
		"Linkedin": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "Linkedin.svg",
			"script":  "//platform\\.linkedin\\.com/in\\.js",
			"website": "http://linkedin.com",
		},
		"List.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "List.js.png",
			"js": map[string]interface{}{
				"List": "",
			},
			"script":  "^list\\.(?:min\\.)?js$",
			"website": "http://listjs.com",
		},
		"LiteSpeed": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^LiteSpeed$",
			},
			"icon":    "LiteSpeed.svg",
			"website": "http://litespeedtech.com",
		},
		"Lithium": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"LithiumVisitor": "",
			},
			"html":    " <a [^>]+Powered by Lithium",
			"icon":    "Lithium.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"LITHIUM": "",
			},
			"website": "https://www.lithium.com",
		},
		"LiveAgent": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon": "LiveAgent.png",
			"js": map[string]interface{}{
				"LiveAgent": "",
			},
			"website": "https://www.ladesk.com",
		},
		"LiveChat": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon":    "LiveChat.png",
			"script":  "cdn\\.livechatinc\\.com/.*tracking\\.js",
			"website": "http://livechatinc.com",
		},
		"LiveHelp": map[string]interface{}{
			"cats": []interface{}{
				52,
				53,
			},
			"icon":    "LiveHelp.png",
			"script":  "^https?://server\\.livehelp\\.it/widgetjs/[0-9]{5}/[0-9]{1,3}\\.js",
			"website": "http://www.livehelp.it",
		},
		"LiveJournal": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"icon":    "LiveJournal.png",
			"url":     "\\.livejournal\\.com",
			"website": "http://www.livejournal.com",
		},
		"LivePerson": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon":    "LivePerson.png",
			"script":  "^https?://lptag\\.liveperson\\.net/tag/tag\\.js",
			"website": "https://www.liveperson.com/",
		},
		"LiveStreet CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "LiveStreet CMS",
			},
			"icon":    "LiveStreet CMS.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"LIVESTREET_SECURITY_KEY": "",
			},
			"website": "http://livestreetcms.com",
		},
		"Livefyre": map[string]interface{}{
			"cats": []interface{}{
				15,
			},
			"html": "<[^>]+(?:id|class)=\"livefyre",
			"icon": "Livefyre.png",
			"js": map[string]interface{}{
				"FyreLoader":      "",
				"L.version":       "^(.+)$\\;confidence:0\\;version:\\1",
				"LF.CommentCount": "",
				"fyre":            "",
			},
			"script":  "livefyre_init\\.js",
			"website": "http://livefyre.com",
		},
		"Liveinternet": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": []interface{}{
				"<script[^<>]*>[^]{0,128}?src\\s*=\\s*['\"]//counter\\.yadro\\.ru/hit(?:;\\S+)?\\?(?:t\\d+\\.\\d+;)?r",
				"<!--LiveInternet counter-->",
				"<!--/LiveInternet-->",
				"<a href=\"http://www\\.liveinternet\\.ru/click\"",
			},
			"icon":    "Liveinternet.png",
			"script":  "/js/al/common\\.js\\?[0-9_]+",
			"website": "http://liveinternet.ru/rating/",
		},
		"LocalFocus": map[string]interface{}{
			"cats": []interface{}{
				61,
			},
			"html": "<iframe[^>]+localfocus",
			"icon": "LocalFocus.png",
			"implies": []interface{}{
				"Angular",
				"D3",
			},
			"website": "https://www.localfocus.nl/en/",
		},
		"Locomotive": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<link[^>]*/sites/[a-z\\d]{24}/theme/stylesheets",
			"icon": "Locomotive.png",
			"implies": []interface{}{
				"Ruby on Rails",
				"MongoDB",
			},
			"website": "http://www.locomotivecms.com",
		},
		"Lodash": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"excludes": "Underscore.js",
			"icon":     "Lo-dash.png",
			"js": map[string]interface{}{
				"_.VERSION":      "^(.+)$\\;confidence:0\\;version:\\1",
				"_.differenceBy": "",
			},
			"script":  "lodash.*\\.js",
			"website": "http://www.lodash.com",
		},
		"Logitech Media Server": map[string]interface{}{
			"cats": []interface{}{
				22,
				38,
			},
			"headers": map[string]interface{}{
				"Server": "Logitech Media Server(?: \\(([\\d\\.]+))?\\;version:\\1",
			},
			"icon":    "Logitech Media Server.png",
			"website": "http://www.mysqueezebox.com",
		},
		"Lotus Domino": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Lotus-Domino",
			},
			"icon":    "Lotus Domino.png",
			"implies": "Java",
			"website": "http://www-01.ibm.com/software/lotus/products/domino",
		},
		"LOU": map[string]interface{}{
			"cats": []interface{}{
				58,
			},
			"icon":    "LOU.png",
			"script":  "cdn\\.louassist\\.com*",
			"website": "https://www.louassist.com",
		},
		"Lua": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "\\bLua(?: ([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Lua.png",
			"website": "http://www.lua.org",
		},
		"Lucene": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon":    "Lucene.png",
			"implies": "Java",
			"website": "http://lucene.apache.org/core/",
		},
		"Luigi’s Box": map[string]interface{}{
			"cats": []interface{}{
				10,
				29,
			},
			"icon": "Luigisbox.svg",
			"js": map[string]interface{}{
				"Luigis": "",
			},
			"website": "https://www.luigisbox.com",
		},
		"M.R. Inc BoxyOS": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"icon":    "M.R. Inc.png",
			"website": "http://mrincworld.com",
		},
		"M.R. Inc SiteFrame": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"Powered-By": "M\\.R\\. Inc SiteFrame",
			},
			"icon":    "M.R. Inc.png",
			"website": "http://mrincworld.com",
		},
		"M.R. Inc Webserver": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "M\\.R\\. Inc Webserver",
			},
			"icon": "M.R. Inc.png",
			"implies": []interface{}{
				"M.R. Inc BoxyOS",
			},
			"website": "http://mrincworld.com",
		},
		"MHonArc": map[string]interface{}{
			"cats": []interface{}{
				50,
			},
			"html":    "<!-- MHonArc v([0-9.]+) -->\\;version:\\1",
			"icon":    "mhonarc.png",
			"website": "http://www.mhonarc.at",
		},
		"MODX": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^MODX",
			},
			"html": []interface{}{
				"<a[^>]+>Powered by MODX</a>",
				"<(?:link|script)[^>]+assets/snippets/\\;confidence:20",
				"<form[^>]+id=\"ajaxSearch_form\\;confidence:20",
				"<input[^>]+id=\"ajaxSearch_input\\;confidence:20",
			},
			"icon":    "MODX.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"MODX":            "",
				"MODX_MEDIA_PATH": "",
			},
			"meta": map[string]interface{}{
				"generator": "MODX[^\\d.]*([\\d.]+)?\\;version:\\1",
			},
			"website": "http://modx.com",
		},
		"MYPAGE Platform": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"cookies": map[string]interface{}{
				"mypage_session": "",
			},
			"headers": map[string]interface{}{
				"CMS-Version": "^(.+)$\\;version:\\1\\;confidence:0",
			},
			"icon":    "mypage-platform.png",
			"implies": "Laravel",
			"website": "https://www.mypage.vn",
		},
		"Botble CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"cookies": map[string]interface{}{
				"botble_session": "",
			},
			"headers": map[string]interface{}{
				"CMS-Version": "^(.+)$\\;version:\\1\\;confidence:0",
			},
			"icon":    "Botble-CMS.png",
			"implies": "Laravel",
			"website": "https://botble.com",
		},
		"MadAdsMedia": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "MadAdsMedia.png",
			"js": map[string]interface{}{
				"setMIframe": "",
				"setMRefURL": "",
			},
			"script":  "^https?://(?:ads-by|pixel)\\.madadsmedia\\.com/",
			"website": "http://madadsmedia.com",
		},
		"Magento": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"frontend": "\\;confidence:50",
			},
			"html": []interface{}{
				"<script [^>]+data-requiremodule=\"mage/\\;version:2",
				"<script [^>]+data-requiremodule=\"Magento_\\;version:2",
				"<script type=\"text/x-magento-init\">",
			},
			"icon": "Magento.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"js": map[string]interface{}{
				"Mage":       "",
				"VarienForm": "",
			},
			"script": []interface{}{
				"js/mage",
				"skin/frontend/(?:default|(enterprise))\\;version:\\1?Enterprise:Community",
				"static/_requirejs\\;confidence:50\\;version:2",
			},
			"website": "https://magento.com",
		},
		"MailChimp": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"html": []interface{}{
				"<form [^>]*data-mailchimp-url",
				"<form [^>]*id=\"mc-embedded-subscribe-form\"",
				"<form [^>]*name=\"mc-embedded-subscribe-form\"",
				"<input [^>]*id=\"mc-email\"\\;confidence:20",
				"<!-- Begin MailChimp Signup Form -->",
			},
			"icon": "mailchimp.svg",
			"script": []interface{}{
				"s3\\.amazonaws\\.com/downloads\\.mailchimp\\.com/js/mc-validate\\.js",
				"cdn-images\\.mailchimp\\.com/[^>]*\\.css",
			},
			"website": "http://mailchimp.com",
		},
		"MakeShopKorea": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "MakeShopKorea.png",
			"js": map[string]interface{}{
				"Makeshop":            "",
				"MakeshopLogUniqueId": "",
			},
			"website": "https://www.makeshop.co.kr",
		},
		"Mambo": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"excludes": "Joomla",
			"icon":     "Mambo.png",
			"meta": map[string]interface{}{
				"generator": "Mambo",
			},
			"website": "http://mambo-foundation.org",
		},
		"MantisBT": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"html":    "<img[^>]+ alt=\"Powered by Mantis Bugtracker",
			"icon":    "MantisBT.png",
			"implies": "PHP",
			"website": "http://www.mantisbt.org",
		},
		"ManyContacts": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "ManyContacts.png",
			"script":  "\\/assets\\/js\\/manycontacts\\.min\\.js",
			"website": "http://www.manycontacts.com",
		},
		"MariaDB": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon":    "mariadb.svg",
			"website": "https://mariadb.org",
		},
		"Marionette.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "Marionette.js.svg",
			"implies": []interface{}{
				"Underscore.js",
				"Backbone.js",
			},
			"js": map[string]interface{}{
				"Marionette":         "",
				"Marionette.VERSION": "^(.+)$\\;version:\\1",
			},
			"script":  "backbone\\.marionette.*\\.js",
			"website": "https://marionettejs.com",
		},
		"Marked": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "marked.svg",
			"js": map[string]interface{}{
				"marked": "",
			},
			"script":  "/marked(?:\\.min)?\\.js",
			"website": "https://marked.js.org",
		},
		"Marketo": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"icon": "Marketo.png",
			"js": map[string]interface{}{
				"Munchkin": "",
			},
			"script":  "munchkin\\.marketo\\.net/munchkin\\.js",
			"website": "https://www.marketo.com",
		},
		"Material Design Lite": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": "<link[^>]* href=\"[^\"]*material(?:\\.[\\w]+-[\\w]+)?(?:\\.min)?\\.css",
			"icon": "Material Design Lite.png",
			"js": map[string]interface{}{
				"MaterialIconToggle": "",
			},
			"script":  "(?:/([\\d.]+))?/material(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://getmdl.io",
		},
		"Materialize CSS": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "<link[^>]* href=\"[^\"]*materialize(?:\\.min)?\\.css",
			"icon":    "Materialize CSS.png",
			"implies": "jQuery",
			"script":  "materialize(?:\\.min)?\\.js",
			"website": "http://materializecss.com",
		},
		"MathJax": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "MathJax.png",
			"js": map[string]interface{}{
				"MathJax":         "",
				"MathJax.version": "^(.+)$\\;version:\\1",
			},
			"script":  "([\\d.]+)?/mathjax\\.js\\;version:\\1",
			"website": "https://www.mathjax.org",
		},
		"Matomo": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"cookies": map[string]interface{}{
				"PIWIK_SESSID": "",
			},
			"icon": "Piwik.png",
			"js": map[string]interface{}{
				"Matomo": "",
				"Piwik":  "",
				"_paq":   "",
			},
			"meta": map[string]interface{}{
				"apple-itunes-app": "app-id=737216887",
				"generator":        "(?:Matomo|Piwik) - Open Source Web Analytics",
				"google-play-app":  "app-id=org\\.piwik\\.mobile2",
			},
			"script":  "piwik\\.js|piwik\\.php",
			"website": "http://piwik.org",
		},
		"Mattermost": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html": "<noscript> To use Mattermost, please enable JavaScript\\. </noscript>",
			"icon": "mattermost.png",
			"implies": []interface{}{
				"Go",
				"React",
			},
			"js": map[string]interface{}{
				"mm_config":          "",
				"mm_current_user_id": "",
				"mm_license":         "",
				"mm_user":            "",
			},
			"website": "https://about.mattermost.com",
		},
		"Mautic": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"icon": "mautic.svg",
			"js": map[string]interface{}{
				"MauticTrackingObject": "",
			},
			"script":  "[^a-z]mtc.*\\.js",
			"website": "https://www.mautic.org/",
		},
		"MaxCDN": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"Server":        "^NetDNA",
				"X-CDN-Forward": "^maxcdn$",
			},
			"icon":    "MaxCDN.png",
			"website": "http://www.maxcdn.com",
		},
		"MaxSite CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "MaxSite CMS.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "MaxSite CMS",
			},
			"website": "http://max-3000.com",
		},
		"Mean.io": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"headers": map[string]interface{}{
				"X-Powered-CMS": "Mean\\.io",
			},
			"icon": "Mean.io.png",
			"implies": []interface{}{
				"MongoDB",
				"Express",
				"Angular",
			},
			"website": "http://mean.io",
		},
		"MediaElement.js": map[string]interface{}{
			"cats": []interface{}{
				14,
			},
			"icon": "MediaElement.js.png",
			"js": map[string]interface{}{
				"mejs":         "",
				"mejs.version": "^(.+)$\\;version:\\1",
			},
			"website": "http://www.mediaelementjs.com",
		},
		"MediaTomb": map[string]interface{}{
			"cats": []interface{}{
				38,
			},
			"headers": map[string]interface{}{
				"Server": "MediaTomb(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "MediaTomb.png",
			"website": "http://mediatomb.cc",
		},
		"MediaWiki": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"html": []interface{}{
				"<body[^>]+class=\"mediawiki\"",
				"<(?:a|img)[^>]+>Powered by MediaWiki</a>",
				"<a[^>]+/Special:WhatLinksHere/",
			},
			"icon":    "MediaWiki.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"mw.util.toggleToc": "",
			},
			"meta": map[string]interface{}{
				"generator": "^MediaWiki ?(.+)$\\;version:\\1",
			},
			"website": "https://www.mediawiki.org",
		},
		"Medium": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Medium$",
			},
			"icon":    "Medium.svg",
			"implies": "Node.js",
			"script":  "medium\\.com",
			"url":     "^https?://(?:www\\.)?medium\\.com",
			"website": "https://medium.com",
		},
		"Meebo": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"html":    "(?:<iframe id=\"meebo-iframe\"|Meebo\\('domReady'\\))",
			"icon":    "Meebo.png",
			"website": "http://www.meebo.com",
		},
		"Melis CMS V2": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"html": "<!-- Rendered with Melis CMS V2",
			"icon": "meliscmsv2.png",
			"meta": map[string]interface{}{
				"powered-by": "^Melis CMS",
			},
			"website": "http://www.melistechnology.com/",
		},
		"Mermaid": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"html": "<div [^>]*class=[\"']mermaid[\"']>\\;confidence:90",
			"js": map[string]interface{}{
				"mermaid": "",
			},
			"script":  "/mermaid(?:\\.min)?\\.js",
			"website": "https://mermaidjs.github.io/",
		},
		"Meteor": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html": "<link[^>]+__meteor-css__",
			"icon": "Meteor.png",
			"implies": []interface{}{
				"MongoDB",
				"Node.js",
			},
			"js": map[string]interface{}{
				"Meteor":         "",
				"Meteor.release": "^METEOR@([\\d.]+)\\;version:\\1",
			},
			"website": "https://www.meteor.com",
		},
		"Methode": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<!-- Methode uuid: \"[a-f\\d]+\" ?-->",
			"icon": "Methode.png",
			"meta": map[string]interface{}{
				"eomportal-id":         "\\d+",
				"eomportal-instanceid": "\\d+",
				"eomportal-lastUpdate": "",
				"eomportal-loid":       "[\\d.]+",
				"eomportal-uuid":       "[a-f\\d]+",
			},
			"website": "https://www.eidosmedia.com/",
		},
		"Microsoft ASP.NET": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"ASP.NET_SessionId": "",
				"ASPSESSION":        "",
			},
			"headers": map[string]interface{}{
				"X-AspNet-Version": "(.+)\\;version:\\1",
				"X-Powered-By":     "^ASP\\.NET",
			},
			"html":    "<input[^>]+name=\"__VIEWSTATE",
			"icon":    "Microsoft ASP.NET.png",
			"implies": "IIS\\;confidence:50",
			"url":     "\\.aspx?(?:$|\\?)",
			"website": "https://www.asp.net",
		},
		"Microsoft Excel": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"html": "(?:<html [^>]*xmlns:w=\"urn:schemas-microsoft-com:office:excel\"|<!--\\s*(?:START|END) OF OUTPUT FROM EXCEL PUBLISH AS WEB PAGE WIZARD\\s*-->|<div [^>]*x:publishsource=\"?Excel\"?)",
			"icon": "Microsoft Excel.svg",
			"meta": map[string]interface{}{
				"ProgId":    "^Excel\\.",
				"generator": "Microsoft Excel( [\\d.]+)?\\;version:\\1",
			},
			"website": "https://office.microsoft.com/excel",
		},
		"Microsoft HTTPAPI": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Microsoft-HTTPAPI(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Microsoft.png",
			"website": "http://microsoft.com",
		},
		"Microsoft PowerPoint": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"html": "(?:<html [^>]*xmlns:w=\"urn:schemas-microsoft-com:office:powerpoint\"|<link rel=\"?Presentation-XML\"? href=\"?[^\"]+\\.xml\"?>|<o:PresentationFormat>[^<]+</o:PresentationFormat>[^!]+<o:Slides>\\d+</o:Slides>(?:[^!]+<o:Version>([\\d.]+)</o:Version>)?)\\;version:\\1",
			"icon": "Microsoft PowerPoint.svg",
			"meta": map[string]interface{}{
				"ProgId":    "^PowerPoint\\.",
				"generator": "Microsoft PowerPoint ( [\\d.]+)?\\;version:\\1",
			},
			"website": "https://office.microsoft.com/powerpoint",
		},
		"Microsoft Publisher": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"html": "(?:<html [^>]*xmlns:w=\"urn:schemas-microsoft-com:office:publisher\"|<!--[if pub]><xml>)",
			"icon": "Microsoft Publisher.svg",
			"meta": map[string]interface{}{
				"ProgId":    "^Publisher\\.",
				"generator": "Microsoft Publisher( [\\d.]+)?\\;version:\\1",
			},
			"website": "https://office.microsoft.com/publisher",
		},
		"Microsoft SharePoint": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"MicrosoftSharePointTeamServices": "^(.+)$\\;version:\\1",
				"SPRequestGuid":                   "",
				"SharePointHealthScore":           "",
				"X-SharePointHealthScore":         "",
			},
			"icon": "Microsoft SharePoint.png",
			"js": map[string]interface{}{
				"SPDesignerProgID":    "",
				"_spBodyOnLoadCalled": "",
			},
			"meta": map[string]interface{}{
				"generator": "Microsoft SharePoint",
			},
			"website": "https://sharepoint.microsoft.com",
		},
		"Microsoft Word": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"html": "(?:<html [^>]*xmlns:w=\"urn:schemas-microsoft-com:office:word\"|<w:WordDocument>|<div [^>]*class=\"?WordSection1[\" >]|<style[^>]*>[^>]*@page WordSection1)",
			"icon": "Microsoft Word.svg",
			"meta": map[string]interface{}{
				"ProgId":    "^Word\\.",
				"generator": "Microsoft Word( [\\d.]+)?\\;version:\\1",
			},
			"website": "https://office.microsoft.com/word",
		},
		"Mietshop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<a href=\"https://ssl\\.mietshop\\.d",
			"icon": "mietshop.png",
			"meta": map[string]interface{}{
				"generator": "Mietshop",
			},
			"website": "http://www.mietshop.de/",
		},
		"Milligram": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": []interface{}{
				"<link[^>]+?href=\"[^\"]+milligram(?:\\.min)?\\.css",
			},
			"icon":    "Milligram.png",
			"website": "https://milligram.github.io",
		},
		"Minero.cc": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"script": []interface{}{
				"//minero\\.cc/lib/minero(?:-miner|-hidden)?\\.min\\.js",
			},
			"website": "http://minero.cc/",
		},
		"MiniBB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html":    "<a href=\"[^\"]+minibb[^<]+</a>[^<]+\\n<!--End of copyright link",
			"icon":    "MiniBB.png",
			"website": "http://www.minibb.com",
		},
		"MiniServ": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "MiniServ\\/?([\\d\\.]+)?\\;version:\\1",
			},
			"website": "http://sourceforge.net/projects/miniserv",
		},
		"Mint": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Mint.png",
			"js": map[string]interface{}{
				"Mint": "",
			},
			"script":  "mint/\\?js",
			"website": "https://haveamint.com",
		},
		"Mithril": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon":    "Mithril.svg",
			"script":  "mithril/\\?js",
			"website": "https://mithril.js.org",
		},
		"Mixpanel": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Mixpanel.png",
			"js": map[string]interface{}{
				"mixpanel": "",
			},
			"script":  "api\\.mixpanel\\.com/track",
			"website": "https://mixpanel.com",
		},
		"MkDocs": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"icon": "mkdocs.png",
			"meta": map[string]interface{}{
				"generator": "^mkdocs-([\\d.]+)\\;version:\\1",
			},
			"website": "http://www.mkdocs.org/",
		},
		"Mobify": map[string]interface{}{
			"cats": []interface{}{
				26,
			},
			"icon": "Mobify.png",
			"js": map[string]interface{}{
				"Mobify": "",
			},
			"script":  "//cdn\\.mobify\\.com/",
			"website": "https://www.mobify.com",
		},
		"Mobirise": map[string]interface{}{
			"cats": []interface{}{
				51,
			},
			"html": []interface{}{
				"<!-- Site made with Mobirise Website Builder v([\\d.]+)\\;version:\\1",
			},
			"icon": "mobirise.png",
			"meta": map[string]interface{}{
				"generator": "^Mobirise v([\\d.]+)\\;version:\\1",
			},
			"website": "https://mobirise.com",
		},
		"MochiKit": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "MochiKit.png",
			"js": map[string]interface{}{
				"MochiKit":                  "",
				"MochiKit.MochiKit.VERSION": "^(.+)$\\;version:\\1",
			},
			"script":  "MochiKit(?:\\.min)?\\.js",
			"website": "https://mochi.github.io/mochikit/",
		},
		"MochiWeb": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "MochiWeb(?:/([\\d.]+))?\\;version:\\1",
			},
			"website": "https://github.com/mochi/mochiweb",
		},
		"Modernizr": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "Modernizr.svg",
			"js": map[string]interface{}{
				"Modernizr._version": "^(.+)$\\;version:\\1",
			},
			"script": []interface{}{
				"([\\d.]+)?/modernizr(?:.([\\d.]+))?.*\\.js\\;version:\\1?\\1:\\2",
			},
			"website": "https://modernizr.com",
		},
		"Modified": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "modified.png",
			"meta": map[string]interface{}{
				"generator": "\\(c\\) by modified eCommerce Shopsoftware ------ http://www\\.modified-shop\\.org",
			},
			"website": "http://www.modified-shop.org/",
		},
		"Moguta.CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"html":    "<link[^>]+href=[\"'][^\"]+mg-(?:core|plugins|templates)/",
			"script":  "mg-(?:core|plugins|templates)/",
			"icon":    "Moguta.CMS.png",
			"implies": "PHP",
			"website": "https://moguta.ru",
		},
		"MoinMoin": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"cookies": map[string]interface{}{
				"MOIN_SESSION": "",
			},
			"icon":    "MoinMoin.png",
			"implies": "Python",
			"js": map[string]interface{}{
				"show_switch2gui": "",
			},
			"script":  "moin(?:_static(\\d)(\\d)(\\d)|.+)/common/js/common\\.js\\;version:\\1.\\2.\\3",
			"website": "https://moinmo.in",
		},
		"Mojolicious": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"server":       "^mojolicious",
				"x-powered-by": "mojolicious",
			},
			"icon":    "Mojolicious.png",
			"implies": "Perl",
			"website": "http://mojolicio.us",
		},
		"Mollom": map[string]interface{}{
			"cats": []interface{}{
				16,
			},
			"html":    "<img[^>]+\\.mollom\\.com",
			"icon":    "Mollom.png",
			"script":  "mollom(?:\\.min)?\\.js",
			"website": "http://mollom.com",
		},
		"Moment Timezone": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "Moment.js.svg",
			"implies": "Moment.js",
			"script":  "moment-timezone(?:-data)?(?:\\.min)?\\.js",
			"website": "http://momentjs.com/timezone/",
		},
		"Moment.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "Moment.js.svg",
			"js": map[string]interface{}{
				"moment":         "",
				"moment.version": "^(.+)$\\;version:\\1",
			},
			"script":  "moment(?:\\.min)?\\.js",
			"website": "https://momentjs.com",
		},
		"Mondo Media": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Mondo Media.png",
			"meta": map[string]interface{}{
				"generator": "Mondo Shop",
			},
			"website": "http://mondo-media.de",
		},
		"Monerominer": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"html":    "<iframe[^>]+src=[\\'\"]https?://monerominer\\.rocks/miner\\.php\\?siteid=",
			"icon":    "monerominer.png",
			"website": "https://monerominer.rocks/",
		},
		"MongoDB": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon":    "MongoDB.png",
			"website": "http://www.mongodb.org",
		},
		"Mongrel": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Mongrel",
			},
			"icon":    "Mongrel.png",
			"implies": "Ruby",
			"website": "http://mongrel2.org",
		},
		"Monkey HTTP Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Monkey/?([\\d.]+)?\\;version:\\1",
			},
			"icon":    "Monkey HTTP Server.png",
			"website": "http://monkey-project.com",
		},
		"Mono": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Mono",
			},
			"icon":    "Mono.png",
			"website": "http://mono-project.com",
		},
		"Mono.net": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Mono.net.png",
			"implies": "Matomo",
			"js": map[string]interface{}{
				"_monoTracker": "",
			},
			"script":  "monotracker(?:\\.min)?\\.js",
			"website": "https://www.mono.net/en",
		},
		"MooTools": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "MooTools.png",
			"js": map[string]interface{}{
				"MooTools":         "",
				"MooTools.version": "^(.+)$\\;version:\\1",
			},
			"script":  "mootools.*\\.js",
			"website": "https://mootools.net",
		},
		"Moodle": map[string]interface{}{
			"cats": []interface{}{
				21,
			},
			"cookies": map[string]interface{}{
				"MOODLEID_":     "",
				"MoodleSession": "",
			},
			"html":    "<img[^>]+moodlelogo",
			"icon":    "Moodle.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"M.core":   "",
				"Y.Moodle": "",
			},
			"meta": map[string]interface{}{
				"keywords": "^moodle",
			},
			"website": "http://moodle.org",
		},
		"Moon": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon":    "moon.svg",
			"script":  "/moon(?:\\.min)?\\.js$",
			"website": "https://kbrsh.github.io/moon/",
		},
		"MotoCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<link [^>]*href=\"[^>]*\\/mt-content\\/[^>]*\\.css",
			"icon": "MotoCMS.svg",
			"implies": []interface{}{
				"PHP",
				"AngularJS",
				"jQuery",
			},
			"script":  "/mt-includes/js/website(?:assets)?\\.(?:min)?\\.js",
			"website": "http://motocms.com",
		},
		"Mouse Flow": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "mouseflow.png",
			"js": map[string]interface{}{
				"_mfq": "",
			},
			"script": []interface{}{
				"cdn\\.mouseflow\\.com",
			},
			"website": "https://mouseflow.com/",
		},
		"Movable Type": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Movable Type.png",
			"meta": map[string]interface{}{
				"generator": "Movable Type",
			},
			"website": "http://movabletype.org",
		},
		"Mozard Suite": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Mozard Suite.png",
			"meta": map[string]interface{}{
				"author": "Mozard",
			},
			"url":     "/mozard/!suite",
			"website": "http://mozard.nl",
		},
		"Mura CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"icon":    "Mura CMS.png",
			"implies": "Adobe ColdFusion",
			"meta": map[string]interface{}{
				"generator": "Mura CMS ([\\d]+)\\;version:\\1",
			},
			"website": "http://www.getmura.com",
		},
		"Mustache": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "Mustache.png",
			"js": map[string]interface{}{
				"Mustache.version": "^(.+)$\\;version:\\1",
			},
			"script":  "mustache(?:\\.min)?\\.js",
			"website": "https://mustache.github.io",
		},
		"MyBB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html": "(?:<script [^>]+\\s+<!--\\s+lang\\.no_new_posts|<a[^>]* title=\"Powered By MyBB)",
			"icon": "MyBB.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"js": map[string]interface{}{
				"MyBB": "",
			},
			"website": "https://mybb.com",
		},
		"MyBlogLog": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "MyBlogLog.png",
			"script":  "pub\\.mybloglog\\.com",
			"website": "http://www.mybloglog.com",
		},
		"MySQL": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon":    "MySQL.svg",
			"website": "http://mysql.com",
		},
		"Mynetcap": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Mynetcap.png",
			"meta": map[string]interface{}{
				"generator": "Mynetcap",
			},
			"website": "http://www.netcap-creation.fr",
		},
		"NEO - Omnichannel Commerce Platform": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"powered": "jet-neo",
			},
			"icon":    "Plataforma NEO.svg",
			"website": "http://www.jetecommerce.com.br/",
		},
		"NVD3": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"html":    "<link[^>]* href=[^>]+nv\\.d3(?:\\.min)?\\.css",
			"icon":    "NVD3.png",
			"implies": "D3",
			"js": map[string]interface{}{
				"nv.addGraph": "",
				"nv.version":  "^(.+)$\\;confidence:0\\;version:\\1",
			},
			"script":  "nv\\.d3(?:\\.min)?\\.js",
			"website": "http://nvd3.org",
		},
		"Navegg": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "Navegg.png",
			"script":  "tag\\.navdmp\\.com",
			"website": "https://www.navegg.com/",
		},
		"Neos CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"excludes": "TYPO3 CMS",
			"headers": map[string]interface{}{
				"X-Flow-Powered": "Neos/?(.+)?$\\;version:\\1",
			},
			"icon":    "Neos.svg",
			"implies": "Neos Flow",
			"url":     "/neos/",
			"website": "https://neos.io",
		},
		"Neos Flow": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"excludes": "TYPO3 CMS",
			"headers": map[string]interface{}{
				"X-Flow-Powered": "Flow/?(.+)?$\\;version:\\1",
			},
			"icon":    "Neos.svg",
			"implies": "PHP",
			"website": "https://flow.neos.io",
		},
		"Nepso": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-CMS": "Nepso",
			},
			"icon":    "Nepso.png",
			"website": "http://nepso.com",
		},
		"Netlify": map[string]interface{}{
			"cats": []interface{}{
				22,
				31,
			},
			"headers": map[string]interface{}{
				"X-NF-Request-ID": "",
				"Server":          "^Netlify",
			},
			"icon":    "Netlify.svg",
			"website": "https://www.netlify.com/",
		},
		"Neto": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Neto.svg",
			"js": map[string]interface{}{
				"NETO": "",
			},
			"script":  "jquery\\.neto.*\\.js",
			"website": "https://www.neto.com.au",
		},
		"Netsuite": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"NS_VER": "",
			},
			"icon":    "Netsuite.png",
			"website": "http://netsuite.com",
		},
		"Nette Framework": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"nette-browser": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Nette Framework",
			},
			"html": []interface{}{
				"<input[^>]+data-nette-rules",
				"<div[^>]+id=\"snippet-",
				"<input[^>]+id=\"frm-",
			},
			"icon":    "Nette Framework.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"Nette":         "",
				"Nette.version": "^(.+)$\\;version:\\1",
			},
			"website": "https://nette.org",
		},
		"New Relic": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "New Relic.png",
			"js": map[string]interface{}{
				"NREUM":    "",
				"newrelic": "",
			},
			"website": "https://newrelic.com",
		},
		"Next.js": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"x-powered-by": "^Next\\.js ?([0-9.]+)?\\;version:\\1",
			},
			"icon": "zeit.svg",
			"implies": []interface{}{
				"React",
				"webpack",
				"Node.js",
			},
			"js": map[string]interface{}{
				"__NEXT_DATA__": "",
			},
			"website": "https://github.com/zeit/next.js",
		},
		"NextGEN Gallery": map[string]interface{}{
			"cats": []interface{}{
				7,
			},
			"html": []interface{}{
				"<!-- <meta name=\"NextGEN\" version=\"([\\d.]+)\" /> -->\\;version:\\1",
			},
			"icon": "NextGEN Gallery.png",
			"implies": []interface{}{
				"WordPress",
			},
			"script":  "/nextgen-gallery/js/",
			"website": "https://www.imagely.com/wordpress-gallery-plugin",
		},
		"Nginx": map[string]interface{}{
			"cats": []interface{}{
				22,
				64,
			},
			"headers": map[string]interface{}{
				"Server":          "nginx(?:/([\\d.]+))?\\;version:\\1",
				"X-Fastcgi-Cache": "",
			},
			"icon":    "Nginx.svg",
			"website": "http://nginx.org/en",
		},
		"Node.js": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"icon":    "node.js.png",
			"website": "http://nodejs.org",
		},
		"NodeBB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^NodeBB$",
			},
			"icon":    "NodeBB.png",
			"implies": "Node.js",
			"script":  "^/nodebb\\.min\\.js\\?",
			"website": "https://nodebb.org",
		},
		"Now": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"server":      "^now$",
				"x-now-trace": "",
				"x-now-id":    "",
				"x-now-cache": "",
			},
			"icon":    "zeit.svg",
			"website": "https://zeit.co/now",
		},
		"OWL Carousel": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"html":    "<link [^>]*href=\"[^\"]+owl\\.carousel(?:\\.min)?\\.css",
			"icon":    "OWL Carousel.png",
			"implies": "jQuery",
			"script":  "owl\\.carousel.*\\.js",
			"website": "https://owlcarousel2.github.io/OwlCarousel2/",
		},
		"OXID eShop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<!--[^-]*OXID eShop",
			"icon": "OXID eShop.png",
			"js": map[string]interface{}{
				"oxCookieNote":     "",
				"oxInputValidator": "",
				"oxLoginBox":       "",
				"oxModalPopup":     "",
				"oxTopMenu":        "",
			},
			"website": "https://en.oxid-esales.com/en/home.html",
		},
		"October CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"october_session=": "",
			},
			"icon":    "October CMS.png",
			"implies": "Laravel",
			"website": "http://octobercms.com",
		},
		"Octopress": map[string]interface{}{
			"cats": []interface{}{
				57,
			},
			"html":    "Powered by <a href=\"http://octopress\\.org\">",
			"icon":    "octopress.png",
			"implies": "Jekyll",
			"meta": map[string]interface{}{
				"generator": "Octopress",
			},
			"script":  "/octopress\\.js",
			"website": "http://octopress.org",
		},
		"Odoo": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"html": "<link[^>]* href=[^>]+/web/css/(?:web\\.assets_common/|website\\.assets_frontend/)\\;confidence:25",
			"icon": "Odoo.png",
			"implies": []interface{}{
				"Python",
				"PostgreSQL",
				"Node.js",
				"Less",
			},
			"meta": map[string]interface{}{
				"generator": "Odoo",
			},
			"script":  "/web/js/(?:web\\.assets_common/|website\\.assets_frontend/)\\;confidence:25",
			"website": "http://odoo.com",
		},
		"Olark": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon":    "Olark.png",
			"script":  "^https?:\\/\\/static\\.olark\\.com\\/jsclient\\/loader1\\.js",
			"website": "https://www.olark.com/",
		},
		"OneAPM": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "OneAPM.png",
			"js": map[string]interface{}{
				"BWEUM": "",
			},
			"website": "http://www.oneapm.com",
		},
		"OneStat": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "OneStat.png",
			"js": map[string]interface{}{
				"OneStat_Pageview": "",
			},
			"website": "http://www.onestat.com",
		},
		"Open AdStream": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Open AdStream.png",
			"js": map[string]interface{}{
				"OAS_AD": "",
			},
			"website": "https://www.xaxis.com",
		},
		"Open Classifieds": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Open Classifieds.png",
			"meta": map[string]interface{}{
				"author":    "open-classifieds\\.com",
				"copyright": "Open Classifieds ?([0-9.]+)?\\;version:\\1",
			},
			"website": "http://open-classifieds.com",
		},
		"Open Journal Systems": map[string]interface{}{
			"cats": []interface{}{
				50,
			},
			"cookies": map[string]interface{}{
				"OJSSID": "",
			},
			"icon":    "Open Journal Systems.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Open Journal Systems(?: ([\\d.]+))?\\;version:\\1",
			},
			"website": "http://pkp.sfu.ca/ojs",
		},
		"Open Web Analytics": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": "<!-- (?:Start|End) Open Web Analytics Tracker -->",
			"icon": "Open Web Analytics.png",
			"js": map[string]interface{}{
				"OWA.config.baseUrl": "",
				"owa_baseUrl":        "",
				"owa_cmds":           "",
			},
			"website": "http://www.openwebanalytics.com",
		},
		"Open eShop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "Open eShop.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"author":    "open-eshop\\.com",
				"copyright": "Open eShop ?([0-9.]+)?\\;version:\\1",
			},
			"website": "http://open-eshop.com/",
		},
		"OpenCart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"OCSESSID": "",
			},
			"icon":    "OpenCart.png",
			"implies": "PHP",
			"website": "http://www.opencart.com",
		},
		"OpenCms": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"Server": "OpenCms",
			},
			"html":    "<link href=\"/opencms/",
			"icon":    "OpenCms.png",
			"implies": "Java",
			"script":  "opencms",
			"website": "http://www.opencms.org",
		},
		"OpenGSE": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "GSE",
			},
			"icon":    "Google.svg",
			"implies": "Java",
			"website": "http://code.google.com/p/opengse",
		},
		"OpenGrok": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"cookies": map[string]interface{}{
				"OpenGrok": "",
			},
			"icon":    "OpenGrok.png",
			"implies": "Java",
			"meta": map[string]interface{}{
				"generator": "OpenGrok(?: v?([\\d.]+))?\\;version:\\1",
			},
			"website": "http://hub.opensolaris.org/bin/view/Project+opengrok/WebHome",
		},
		"OpenLayers": map[string]interface{}{
			"cats": []interface{}{
				35,
			},
			"icon": "OpenLayers.png",
			"js": map[string]interface{}{
				"OpenLayers.VERSION_NUMBER": "([\\d.]+)\\;version:\\1",
				"ol.CanvasMap":              "",
			},
			"script":  "openlayers",
			"website": "https://openlayers.org",
		},
		"OpenNemas": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "OpenNemas",
			},
			"icon": "OpenNemas.png",
			"meta": map[string]interface{}{
				"generator": "OpenNemas",
			},
			"website": "http://www.opennemas.com",
		},
		"OpenResty": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "openresty(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon": "OpenResty.png",
			"implies": []interface{}{
				"Nginx",
				"Lua",
			},
			"website": "http://openresty.org",
		},
		"OpenSSL": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server": "OpenSSL(?:/([\\d.]+[a-z]?))?\\;version:\\1",
			},
			"icon":    "OpenSSL.png",
			"website": "http://openssl.org",
		},
		"OpenText Web Solutions": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<!--[^>]+published by Open Text Web Solutions",
			"icon":    "OpenText Web Solutions.png",
			"implies": "Microsoft ASP.NET",
			"website": "http://websolutions.opentext.com",
		},
		"OpenUI5": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "OpenUI5.png",
			"js": map[string]interface{}{
				"sap.ui.version": "^(.+)$\\;version:\\1",
			},
			"script":  "sap-ui-core\\.js",
			"website": "http://openui5.org/",
		},
		"OpenX": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "OpenX.png",
			"script": []interface{}{
				"https?://[^/]*\\.openx\\.net",
				"https?://[^/]*\\.servedbyopenx\\.com",
			},
			"website": "http://openx.com",
		},
		"Ophal": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Ophal(?: (.+))? \\(ophal\\.org\\)\\;version:\\1",
			},
			"icon":    "Ophal.png",
			"implies": "Lua",
			"meta": map[string]interface{}{
				"generator": "Ophal(?: (.+))? \\(ophal\\.org\\)\\;version:\\1",
			},
			"script":  "ophal\\.js",
			"website": "http://ophal.org",
		},
		"Optimizely": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Optimizely.png",
			"js": map[string]interface{}{
				"optimizely": "",
			},
			"script":  "optimizely\\.com.*\\.js",
			"website": "https://www.optimizely.com",
		},
		"Oracle Application Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Oracle[- ]Application[- ]Server(?: Containers for J2EE)?(?:[- ](\\d[\\da-z./]+))?\\;version:\\1",
			},
			"icon":    "Oracle.png",
			"website": "http://www.oracle.com/technetwork/middleware/ias/overview/index.html",
		},
		"Oracle Commerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-ATG-Version": "(?:ATGPlatform/([\\d.]+))?\\;version:\\1",
			},
			"html":    "<[^>]+_dyncharset",
			"icon":    "Oracle.png",
			"website": "http://www.oracle.com/applications/customer-experience/commerce/products/commerce-platform/index.html",
		},
		"Oracle Commerce Cloud": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"OracleCommerceCloud-Version": "^(.+)$\\;version:\\1",
			},
			"html":    "<[^>]+id=\"oracle-cc\"",
			"icon":    "Oracle.png",
			"website": "http://cloud.oracle.com/commerce-cloud",
		},
		"Oracle Dynamic Monitoring Service": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"headers": map[string]interface{}{
				"x-oracle-dms-ecid": "",
			},
			"icon":    "Oracle.png",
			"website": "http://oracle.com",
		},
		"Oracle HTTP Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Oracle-HTTP-Server(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Oracle.png",
			"website": "http://oracle.com",
		},
		"Oracle Recommendations On Demand": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "Oracle.png",
			"script":  "atgsvcs.+atgsvcs\\.js",
			"website": "http://www.oracle.com/us/products/applications/commerce/recommendations-on-demand/index.html",
		},
		"Oracle Web Cache": map[string]interface{}{
			"cats": []interface{}{
				23,
			},
			"headers": map[string]interface{}{
				"Server": "Oracle(?:AS)?[- ]Web[- ]Cache(?:[- /]([\\da-z./]+))?\\;version:\\1",
			},
			"icon":    "Oracle.png",
			"website": "http://oracle.com",
		},
		"Orchard CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Orchard CMS.png",
			"implies": "Microsoft ASP.NET",
			"meta": map[string]interface{}{
				"generator": "Orchard",
			},
			"website": "http://orchardproject.net",
		},
		"Outbrain": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon": "Outbrain.png",
			"js": map[string]interface{}{
				"OB_releaseVer":     "^(.+)$\\;version:\\1",
				"OutbrainPermaLink": "",
			},
			"script":  "widgets\\.outbrain\\.com/outbrain\\.js",
			"website": "https://www.outbrain.com",
		},
		"Outlook Web App": map[string]interface{}{
			"cats": []interface{}{
				30,
			},
			"html":    "<link\\s[^>]*href=\"[^\"]*?([\\d.]+)/themes/resources/owafont\\.css\\;version:\\1",
			"icon":    "Outlook.svg",
			"implies": "Microsoft ASP.NET",
			"js": map[string]interface{}{
				"IsOwaPremiumBrowser": "",
			},
			"url":     "/owa/auth/log(?:on|off)\\.aspx",
			"website": "http://help.outlook.com",
		},
		"PANSITE": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "PANSITE.png",
			"meta": map[string]interface{}{
				"generator": "PANSITE",
			},
			"website": "http://panvision.de/Produkte/Content_Management/index.asp",
		},
		"PDF.js": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html": "<\\/div>\\s*<!-- outerContainer -->\\s*<div\\s*id=\"printContainer\"><\\/div>",
			"icon": "PDF.js.svg",
			"js": map[string]interface{}{
				"PDFJS":         "",
				"PDFJS.version": "^(.+)$\\;version:\\1",
			},
			"url":     "/web/viewer\\.html?file=[^&]\\.pdf",
			"website": "https://mozilla.github.io/pdf.js/",
		},
		"PHP": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"cookies": map[string]interface{}{
				"PHPSESSID": "",
			},
			"headers": map[string]interface{}{
				"Server":       "php/?([\\d.]+)?\\;version:\\1",
				"X-Powered-By": "^php/?([\\d.]+)?\\;version:\\1",
			},
			"icon":    "PHP.svg",
			"url":     "\\.php(?:$|\\?)",
			"website": "http://php.net",
		},
		"PHP-Fusion": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "Powered by <a href=\"[^>]+php-fusion",
			"headers": map[string]interface{}{
				"X-Powered-By": "PHP-Fusion (.+)$\\;version:\\1",
			},
			"icon": "PHP-Fusion.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"website": "https://www.php-fusion.co.uk",
		},
		"PHP-Nuke": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<[^>]+Powered by PHP-Nuke",
			"icon":    "PHP-Nuke.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "PHP-Nuke",
			},
			"website": "http://phpnuke.org",
		},
		"PHPDebugBar": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"icon": "phpdebugbar.png",
			"js": map[string]interface{}{
				"PhpDebugBar": "",
				"phpdebugbar": "",
			},
			"script": []interface{}{
				"debugbar.*\\.js",
			},
			"website": "http://phpdebugbar.com/",
		},
		"Cecil": map[string]interface{}{
			"cats": []interface{}{
				57,
			},
			"icon": "Cecil.png",
			"meta": map[string]interface{}{
				"generator": "^Cecil|PHPoole$",
			},
			"website": "https://cecil.app",
		},
		"Pagekit": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Pagekit.png",
			"meta": map[string]interface{}{
				"generator": "Pagekit",
			},
			"website": "http://pagekit.com",
		},
		"Pagevamp": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-ServedBy": "pagevamp",
			},
			"icon": "Pagevamp.png",
			"js": map[string]interface{}{
				"Pagevamp": "",
			},
			"website": "https://www.pagevamp.com",
		},
		"Pantheon": map[string]interface{}{
			"cats": []interface{}{
				62,
			},
			"headers": map[string]interface{}{
				"x-pantheon-styx-hostname": "",
				"Server":                   "^Pantheon",
			},
			"implies": []interface{}{
				"PHP",
				"Nginx",
				"MariaDB",
			},
			"icon":    "pantheon.svg",
			"website": "https://pantheon.io/",
		},
		"Paper.js": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "paperjs.png",
			"js": map[string]interface{}{
				"paper.version": "^(.+)$\\;version:\\1",
			},
			"website": "http://paperjs.org/",
		},
		"Pardot": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"headers": map[string]interface{}{
				"X-Pardot-LB":    "",
				"X-Pardot-Route": "",
				"X-Pardot-Rsp":   "",
			},
			"icon": "Pardot.png",
			"js": map[string]interface{}{
				"piAId":      "",
				"piCId":      "",
				"piHostname": "",
				"piProtocol": "",
				"piTracker":  "",
			},
			"website": "https://www.pardot.com",
		},
		"Pars Elecom Portal": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Pars Elecom Portal",
			},
			"icon": "parselecom.png",
			"implies": []interface{}{
				"Microsoft ASP.NET",
				"IIS",
				"Windows Server",
			},
			"meta": map[string]interface{}{
				"copyright": "Pars Elecom Portal",
			},
			"website": "http://parselecom.net",
		},
		"Parse.ly": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Parse.ly.png",
			"js": map[string]interface{}{
				"PARSELY": "",
			},
			"website": "https://www.parse.ly",
		},
		"Paths.js": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"script":  "paths(?:\\.min)?\\.js",
			"website": "https://github.com/andreaferretti/paths-js",
		},
		"PayPal": map[string]interface{}{
			"cats": []interface{}{
				41,
			},
			"html": "<input[^>]+_s-xclick",
			"icon": "PayPal.svg",
			"js": map[string]interface{}{
				"PAYPAL": "",
			},
			"script":  "paypalobjects\\.com/js",
			"website": "https://paypal.com",
		},
		"Pelican": map[string]interface{}{
			"cats": []interface{}{
				57,
			},
			"html": []interface{}{
				"powered by <a href=\"[^>]+getpelican\\.com",
				"powered by <a href=\"https?://pelican\\.notmyidea\\.org",
			},
			"icon":    "pelican.png",
			"implies": "Python",
			"website": "https://blog.getpelican.com/",
		},
		"PencilBlue": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "PencilBlue",
			},
			"icon":    "PencilBlue.png",
			"implies": "Node.js",
			"website": "http://pencilblue.org",
		},
		"Pendo": map[string]interface{}{
			"cats": []interface{}{
				58,
			},
			"icon":    "Pendo.svg",
			"script":  "cdn\\.pendo\\.io*\\.js",
			"website": "https://pendo.io",
		},
		"Percona": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon":    "percona.svg",
			"website": "https://www.percona.com",
		},
		"Percussion": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<[^>]+class=\"perc-region\"",
			"icon": "Percussion.png",
			"meta": map[string]interface{}{
				"generator": "(?:Percussion|Rhythmyx)",
			},
			"website": "http://percussion.com",
		},
		"Perl": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"headers": map[string]interface{}{
				"Server": "\\bPerl\\b(?: ?/?v?([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Perl.png",
			"website": "http://perl.org",
		},
		"Phabricator": map[string]interface{}{
			"cats": []interface{}{
				13,
				47,
			},
			"cookies": map[string]interface{}{
				"phsid": "",
			},
			"html": "<[^>]+(?:class|id)=\"phabricator-",
			"icon": "Phabricator.png",
			"implies": []interface{}{
				"PHP",
			},
			"script":  "/phabricator/[a-f0-9]{8}/rsrc/js/phui/[a-z-]+\\.js$",
			"website": "http://phacility.com",
		},
		"Phaser": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "Phaser.png",
			"js": map[string]interface{}{
				"Phaser":         "",
				"Phaser.VERSION": "^(.+)$\\;version:\\1",
			},
			"website": "https://phaser.io",
		},
		"Phenomic": map[string]interface{}{
			"cats": []interface{}{
				57,
			},
			"html": []interface{}{
				"<[^>]+id=\"phenomic(?:root)?\"",
			},
			"icon": "Phenomic.svg",
			"implies": []interface{}{
				"React",
			},
			"script":  "/phenomic\\.browser\\.[a-f0-9]+\\.js",
			"website": "https://phenomic.io/",
		},
		"Phusion Passenger": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server":       "Phusion Passenger ([\\d.]+)\\;version:\\1",
				"X-Powered-By": "Phusion Passenger ?([\\d.]+)?\\;version:\\1",
			},
			"icon":    "Phusion Passenger.png",
			"website": "https://phusionpassenger.com",
		},
		"Pimcore": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^pimcore$",
			},
			"icon":    "pimcore.svg",
			"implies": "PHP",
			"website": "http://pimcore.org",
		},
		"Pingoteam": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Pingoteam.svg",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"designer": "Pingoteam",
			},
			"website": "https://www.pingoteam.ir/",
		},
		"Pinterest": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "Pinterest.svg",
			"script":  "//assets\\.pinterest\\.com/js/pinit\\.js",
			"website": "http://pinterest.com",
		},
		"Planet": map[string]interface{}{
			"cats": []interface{}{
				49,
			},
			"icon": "Planet.png",
			"meta": map[string]interface{}{
				"generator": "^Planet(?:/([\\d.]+))?\\;version:\\1",
			},
			"website": "http://planetplanet.org",
		},
		"PlatformOS": map[string]interface{}{
			"cats": []interface{}{
				1,
				62,
			},
			"icon": "PlatformOS.svg",
			"headers": map[string]interface{}{
				"x-powered-by": "^platformOS$",
			},
			"website": "https://www.platform-os.com",
		},
		"Platform.sh": map[string]interface{}{
			"cats": []interface{}{
				62,
			},
			"icon": "platformsh.svg",
			"headers": map[string]interface{}{
				"x-platform-cluster":   "",
				"x-platform-processor": "",
				"x-platform-router":    "",
				"x-platform-server":    "",
			},
			"website": "https://platform.sh",
		},
		"Play": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"PLAY_SESSION": "",
			},
			"icon":    "Play.svg",
			"implies": "Scala",
			"website": "https://www.playframework.com",
		},
		"Plentymarkets": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Plentymarkets.png",
			"meta": map[string]interface{}{
				"generator": "plentymarkets",
			},
			"website": "http://plentymarkets.eu",
		},
		"Plesk": map[string]interface{}{
			"cats": []interface{}{
				9,
			},
			"headers": map[string]interface{}{
				"X-Powered-By":       "^Plesk(?:L|W)in",
				"X-Powered-By-Plesk": "^Plesk",
			},
			"icon":    "Plesk.png",
			"script":  "common\\.js\\?plesk",
			"website": "https://www.plesk.com/",
		},
		"Pligg": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<span[^>]+id=\"xvotes-0",
			"icon": "Pligg.png",
			"js": map[string]interface{}{
				"pligg_": "",
			},
			"meta": map[string]interface{}{
				"generator": "Pligg",
			},
			"website": "http://pligg.com",
		},
		"Plone": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Plone.png",
			"implies": "Python",
			"meta": map[string]interface{}{
				"generator": "Plone",
			},
			"website": "http://plone.org",
		},
		"Plotly": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon":    "Plotly.png",
			"implies": "D3",
			"js": map[string]interface{}{
				"Plotly.version": "([\\d.])\\;version:\\1",
			},
			"script":  "https?://cdn\\.plot\\.ly/plotly",
			"website": "https://plot.ly/javascript/",
		},
		"Po.st": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon": "Po.st.png",
			"js": map[string]interface{}{
				"pwidget_config": "",
			},
			"website": "http://www.po.st/",
		},
		"Polyfill": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "polyfill.svg",
			"script": []interface{}{
				"^https?://cdn\\.polyfill\\.io/",
				"/polyfill\\.min\\.js",
			},
			"website": "https://polyfill.io",
		},
		"Polymer": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html": "(?:<polymer-[^>]+|<link[^>]+rel=\"import\"[^>]+/polymer\\.html\")",
			"icon": "Polymer.png",
			"js": map[string]interface{}{
				"Polymer.version": "^(.+)$\\;version:\\1",
			},
			"script":  "polymer\\.js",
			"website": "http://polymer-project.org",
		},
		"Posterous": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"html": "<div class=\"posterous",
			"icon": "Posterous.png",
			"js": map[string]interface{}{
				"Posterous": "",
			},
			"website": "http://posterous.com",
		},
		"PostgreSQL": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon":    "PostgreSQL.png",
			"website": "http://www.postgresql.org/",
		},
		"Powergap": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": []interface{}{
				"<a[^>]+title=\"POWERGAP",
				"<input type=\"hidden\" name=\"shopid\"",
			},
			"icon":    "Powergap.png",
			"website": "http://powergap.de",
		},
		"Prebid": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Prebid.png",
			"js": map[string]interface{}{
				"PREBID_TIMEOUT": "",
				"pbjs":           "",
			},
			"script": []interface{}{
				"/prebid\\.js",
				"adnxs\\.com/[^\"]*(?:prebid|/pb\\.js)",
			},
			"website": "http://prebid.org",
		},
		"Prefix-Free": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon": "Prefix-Free.png",
			"js": map[string]interface{}{
				"PrefixFree": "",
			},
			"script":  "prefixfree\\.js",
			"website": "https://leaverou.github.io/prefixfree/",
		},
		"PrestaShop": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"PrestaShop": "",
			},
			"headers": map[string]interface{}{
				"Powered-By": "^Prestashop$",
			},
			"html": []interface{}{
				"Powered by <a\\s+[^>]+>PrestaShop",
				"<!-- /Block [a-z ]+ module (?:HEADER|TOP)?\\s?-->",
				"<!-- /Module Block [a-z ]+ -->",
			},
			"icon": "PrestaShop.svg",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"js": map[string]interface{}{
				"freeProductTranslation": "\\;confidence:25",
				"priceDisplayMethod":     "\\;confidence:25",
				"priceDisplayPrecision":  "\\;confidence:25",
			},
			"meta": map[string]interface{}{
				"generator": "PrestaShop",
			},
			"website": "http://www.prestashop.com",
		},
		"Prism": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon": "Prism.svg",
			"js": map[string]interface{}{
				"Prism": "",
			},
			"script":  "prism\\.js",
			"website": "http://prismjs.com",
		},
		"Project Wonderful": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": "<div[^>]+id=\"pw_adbox_",
			"icon": "Project Wonderful.png",
			"js": map[string]interface{}{
				"pw_adloader": "",
			},
			"script":  "^https?://(?:www\\.)?projectwonderful\\.com/(?:pwa\\.js|gen\\.php)",
			"website": "http://projectwonderful.com",
		},
		"ProjectPoi": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon": "ProjectPoi.png",
			"js": map[string]interface{}{
				"ProjectPoi": "",
			},
			"script":  "^(?:https):?//ppoi\\.org/lib/",
			"website": "https://ppoi.org/",
		},
		"Projesoft": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "projesoft.png",
			"script": []interface{}{
				"projesoft\\.js",
			},
			"website": "https://www.projesoft.com.tr",
		},
		"Prototype": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "Prototype.png",
			"js": map[string]interface{}{
				"Prototype.Version": "^(.+)$\\;version:\\1",
			},
			"script":  "(?:prototype|protoaculous)(?:-([\\d.]*[\\d]))?.*\\.js\\;version:\\1",
			"website": "http://www.prototypejs.org",
		},
		"Protovis": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"js": map[string]interface{}{
				"protovis": "",
			},
			"script":  "protovis.*\\.js",
			"website": "http://mbostock.github.io/protovis",
		},
		"Proximis Omnichannel": map[string]interface{}{
			"cats": []interface{}{
				6,
				1,
			},
			"html": "<html[^>]+data-ng-app=\"RbsChangeApp\"",
			"icon": "Proximis Omnichannel.png",
			"implies": []interface{}{
				"PHP",
				"AngularJS",
			},
			"js": map[string]interface{}{
				"__change": "",
			},
			"meta": map[string]interface{}{
				"generator": "Proximis Omnichannel",
			},
			"website": "https://www.proximis.com",
		},
		"Proximis Web to Store": map[string]interface{}{
			"cats": []interface{}{
				5,
				6,
			},
			"icon":    "Proximis Omnichannel.png",
			"script":  "widget-commerce(?:\\.min)?\\.js",
			"website": "https://www.proximis.com",
		},
		"PubMatic": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon":    "PubMatic.png",
			"script":  "https?://[^/]*\\.pubmatic\\.com",
			"website": "http://www.pubmatic.com/",
		},
		"Public CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"PUBLICCMS_USER": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-PublicCMS": "^(.+)$\\;version:\\1",
			},
			"icon":    "Public CMS.png",
			"implies": "Java",
			"website": "http://www.publiccms.com",
		},
		"Pure CSS": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": []interface{}{
				"<link[^>]+(?:([\\d.])+/)?pure(?:-min)?\\.css\\;version:\\1",
				"<div[^>]+class=\"[^\"]*pure-u-(?:sm-|md-|lg-|xl-)?\\d-\\d",
			},
			"icon":    "Pure CSS.png",
			"website": "http://purecss.io",
		},
		"Pygments": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html":    "<link[^>]+pygments\\.css[\"']",
			"icon":    "pygments.png",
			"website": "http://pygments.org",
		},
		"PyroCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"pyrocms": "",
			},
			"headers": map[string]interface{}{
				"X-Streams-Distribution": "PyroCMS",
			},
			"icon":    "PyroCMS.png",
			"implies": "Laravel",
			"website": "http://pyrocms.com",
		},
		"Python": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"headers": map[string]interface{}{
				"Server": "(?:^|\\s)Python(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Python.png",
			"website": "http://python.org",
		},
		"Quantcast": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Quantcast.png",
			"js": map[string]interface{}{
				"quantserve": "",
			},
			"script":  "\\.quantserve\\.com/quant\\.js",
			"website": "http://www.quantcast.com",
		},
		"Question2Answer": map[string]interface{}{
			"cats": []interface{}{
				15,
			},
			"html":    "<!-- Powered by Question2Answer",
			"icon":    "question2answer.png",
			"implies": "PHP",
			"script":  "\\./qa-content/qa-page\\.js\\?([0-9.]+)\\;version:\\1",
			"website": "http://www.question2answer.org",
		},
		"Quick.CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<a href=\"[^>]+opensolution\\.org/\">CMS by",
			"icon": "Quick.CMS.png",
			"meta": map[string]interface{}{
				"generator": "Quick\\.CMS(?: v([\\d.]+))?\\;version:\\1",
			},
			"website": "http://opensolution.org",
		},
		"Quick.Cart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<a href=\"[^>]+opensolution\\.org/\">(?:Shopping cart by|Sklep internetowy)",
			"icon": "Quick.Cart.png",
			"meta": map[string]interface{}{
				"generator": "Quick\\.Cart(?: v([\\d.]+))?\\;version:\\1",
			},
			"website": "http://opensolution.org",
		},
		"Quill": map[string]interface{}{
			"cats": []interface{}{
				24,
			},
			"icon": "Quill.png",
			"js": map[string]interface{}{
				"Quill": "",
			},
			"website": "http://quilljs.com",
		},
		"RBS Change": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"html":    "<html[^>]+xmlns:change=",
			"icon":    "RBS Change.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "RBS Change",
			},
			"website": "http://www.rbschange.fr",
		},
		"RCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "RCMS.png",
			"meta": map[string]interface{}{
				"generator": "^(?:RCMS|ReallyCMS)",
			},
			"website": "http://www.rcms.fi",
		},
		"RD Station": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"icon": "RD Station.png",
			"js": map[string]interface{}{
				"RDStation": "",
			},
			"script":  "d335luupugsy2\\.cloudfront\\.net/js/loader-scripts/.*-loader\\.js",
			"website": "http://rdstation.com.br",
		},
		"RDoc": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"html": []interface{}{
				"<link[^>]+href=\"[^\"]*rdoc-style\\.css",
				"Generated by <a[^>]+href=\"https?://rdoc\\.rubyforge\\.org[^>]+>RDoc</a> ([\\d.]*\\d)\\;version:\\1",
			},
			"icon":    "RDoc.png",
			"implies": "Ruby",
			"website": "https://github.com/RDoc/RDoc",
		},
		"RackCache": map[string]interface{}{
			"cats": []interface{}{
				23,
			},
			"headers": map[string]interface{}{
				"X-Rack-Cache": "",
			},
			"icon":    "RackCache.png",
			"implies": "Ruby",
			"website": "https://github.com/rtomayko/rack-cache",
		},
		"RainLoop": map[string]interface{}{
			"cats": []interface{}{
				30,
			},
			"headers": map[string]interface{}{
				"Server": "^RainLoop",
			},
			"html": []interface{}{
				"<link[^>]href=\"rainloop/v/([0-9.]+)/static/apple-touch-icon\\.png/>\\;version:\\1",
			},
			"meta": map[string]interface{}{
				"rlAppVersion": "^([0-9.]+)$\\;version:\\1",
			},
			"icon":    "RainLoop.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"rainloopI18N": "",
				"rainloop":     "",
			},
			"script":  "^rainloop/v/([0-9.]+)/\\;version:\\1",
			"website": "https://www.rainloop.net/",
		},
		"Rakuten DBCore": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Rakuten DBCore.png",
			"meta": map[string]interface{}{
				"generator":      "Rakuten DBCore",
				"generator:site": "http://ecservice\\.rakuten\\.com\\.br",
			},
			"website": "http://ecservice.rakuten.com.br",
		},
		"Rakuten Digital Commerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "RakutenDigitalCommerce.png",
			"js": map[string]interface{}{
				"RakutenApplication": "",
			},
			"website": "https://digitalcommerce.rakuten.com.br",
		},
		"Ramda": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "Ramda.png",
			"script":  "ramda.*\\.js",
			"website": "http://ramdajs.com",
		},
		"Raphael": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "Raphael.png",
			"js": map[string]interface{}{
				"Raphael.version": "^(.+)$\\;version:\\1",
			},
			"script":  "raphael(?:-([\\d.]+))?(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://dmitrybaranovskiy.github.io/raphael/",
		},
		"Raspbian": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":       "Raspbian",
				"X-Powered-By": "Raspbian",
			},
			"icon":    "Raspbian.svg",
			"website": "https://www.raspbian.org/",
		},
		"Raychat": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon": "raychat.png",
			"js": map[string]interface{}{
				"Raychat": "",
			},
			"script":  "app\\.raychat\\.io/scripts/js",
			"website": "https://raychat.io",
		},
		"Rayo": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Rayo.png",
			"implies": []interface{}{
				"AngularJS",
				"Microsoft ASP.NET",
			},
			"js": map[string]interface{}{
				"Rayo": "",
			},
			"meta": map[string]interface{}{
				"generator": "^Rayo",
			},
			"website": "http://www.rayo.ir",
		},
		"Rdf": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"website": "https://www.w3.org/RDF/",
		},
		"ReDoc": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"html":    "<redoc ",
			"icon":    "redoc.png",
			"implies": "React",
			"js": map[string]interface{}{
				"Redoc.version": "^(.+)$\\;version:\\1",
			},
			"script":  "/redoc\\.(?:min\\.)?js",
			"website": "https://github.com/Rebilly/ReDoc",
		},
		"React": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html": "<[^>]+data-react",
			"icon": "React.png",
			"js": map[string]interface{}{
				"React.version": "^(.+)$\\;version:\\1",
				"react.version": "^(.+)$\\;version:\\1",
			},
			"script": []interface{}{
				"react(?:-with-addons)?[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"/([\\d.]+)/react(?:\\.min)?\\.js\\;version:\\1",
				"react.*\\.js",
			},
			"website": "https://reactjs.org",
		},
		"Red Hat": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":       "Red Hat",
				"X-Powered-By": "Red Hat",
			},
			"icon":    "Red Hat.png",
			"website": "http://redhat.com",
		},
		"Reddit": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html":    "(?:<a[^>]+Powered by Reddit|powered by <a[^>]+>reddit<)",
			"icon":    "Reddit.png",
			"implies": "Python",
			"js": map[string]interface{}{
				"reddit": "",
			},
			"url":     "^https?://(?:www\\.)?reddit\\.com",
			"website": "http://code.reddit.com",
		},
		"Redmine": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"cookies": map[string]interface{}{
				"_redmine_session": "",
			},
			"html":    "Powered by <a href=\"[^>]+Redmine",
			"icon":    "Redmine.png",
			"implies": "Ruby on Rails",
			"meta": map[string]interface{}{
				"description": "Redmine",
			},
			"website": "http://www.redmine.org",
		},
		"Reinvigorate": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Reinvigorate.png",
			"js": map[string]interface{}{
				"reinvigorate": "",
			},
			"website": "http://www.reinvigorate.net",
		},
		"RequireJS": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "RequireJS.png",
			"js": map[string]interface{}{
				"requirejs.version": "^(.+)$\\;version:\\1",
			},
			"script":  "require.*\\.js",
			"website": "http://requirejs.org",
		},
		"Resin": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^Resin(?:/(\\S*))?\\;version:\\1",
			},
			"icon":    "Resin.png",
			"implies": "Java",
			"website": "http://caucho.com",
		},
		"Reveal.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon":    "Reveal.js.png",
			"implies": "Highlight.js",
			"js": map[string]interface{}{
				"Reveal.VERSION": "^(.+)$\\;version:\\1",
			},
			"script":  "(?:^|/)reveal(?:\\.min)?\\.js",
			"website": "http://lab.hakim.se/reveal-js",
		},
		"Revel": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"REVEL_FLASH":   "",
				"REVEL_SESSION": "",
			},
			"icon":    "Revel.png",
			"implies": "Go",
			"website": "https://revel.github.io",
		},
		"Revslider": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html": []interface{}{
				"<link[^>]* href=[\\'\"][^']+revslider[/\\w-]+\\.css\\?ver=([0-9.]+)[\\'\"]\\;version:\\1",
			},
			"icon":    "revslider.png",
			"implies": "WordPress",
			"script":  "/revslider/[/\\w-]+/js",
			"website": "https://revolution.themepunch.com/",
		},
		"Rickshaw": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"implies": "D3",
			"js": map[string]interface{}{
				"Rickshaw": "",
			},
			"script":  "rickshaw(?:\\.min)?\\.js",
			"website": "http://code.shutterstock.com/rickshaw/",
		},
		"RightJS": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "RightJS.png",
			"js": map[string]interface{}{
				"RightJS": "",
			},
			"script":  "right\\.js",
			"website": "http://rightjs.org",
		},
		"Riot": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "Riot.png",
			"js": map[string]interface{}{
				"riot": "",
			},
			"script":  "riot(?:\\+compiler)?(?:\\.min)?\\.js",
			"website": "https://riot.js.org/",
		},
		"RiteCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "RiteCMS.png",
			"implies": []interface{}{
				"PHP",
				"SQLite\\;confidence:80",
			},
			"meta": map[string]interface{}{
				"generator": "^RiteCMS(?: (.+))?\\;version:\\1",
			},
			"website": "http://ritecms.com",
		},
		"Roadiz CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Roadiz CMS",
			},
			"icon": "Roadiz CMS.png",
			"implies": []interface{}{
				"PHP",
				"Symfony",
			},
			"meta": map[string]interface{}{
				"generator": "^Roadiz ([a-z0-9\\s\\.]+) - \\;version:\\1",
			},
			"website": "http://www.roadiz.io",
		},
		"Robin": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Robin.png",
			"js": map[string]interface{}{
				"_robin_getRobinJs":      "",
				"robin_settings":         "",
				"robin_storage_settings": "",
			},
			"website": "http://www.robinhq.com",
		},
		"RockRMS": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
				32,
			},
			"icon": "RockRMS.svg",
			"implies": []interface{}{
				"Windows Server",
				"IIS",
				"Microsoft ASP.NET",
			},
			"meta": map[string]interface{}{
				"generator": "^Rock v([0-9.]+)\\;version:\\1",
			},
			"website": "http://www.rockrms.com",
		},
		"RoundCube": map[string]interface{}{
			"cats": []interface{}{
				30,
			},
			"html":    "<title>RoundCube",
			"icon":    "RoundCube.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"rcmail":    "",
				"roundcube": "",
			},
			"website": "http://roundcube.net",
		},
		"Rubicon Project": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon":    "Rubicon Project.png",
			"script":  "https?://[^/]*\\.rubiconproject\\.com",
			"website": "http://rubiconproject.com/",
		},
		"Ruby": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"headers": map[string]interface{}{
				"Server": "(?:Mongrel|WEBrick|Ruby)",
			},
			"icon":    "Ruby.png",
			"website": "http://ruby-lang.org",
		},
		"Ruby on Rails": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"Server":       "mod_(?:rails|rack)",
				"X-Powered-By": "mod_(?:rails|rack)",
			},
			"icon":    "Ruby on Rails.png",
			"implies": "Ruby",
			"meta": map[string]interface{}{
				"csrf-param": "^authenticity_token$\\;confidence:50",
			},
			"cookies": map[string]interface{}{
				"_session_id": "\\;confidence:75",
			},
			"script":  "/assets/application-[a-z\\d]{32}/\\.js\\;confidence:50",
			"website": "https://rubyonrails.org",
		},
		"Ruxit": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "Ruxit.png",
			"script":  "ruxitagentjs",
			"website": "http://ruxit.com",
		},
		"RxJS": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "RxJS.png",
			"js": map[string]interface{}{
				"Rx.CompositeDisposable": "",
				"Rx.Symbol":              "",
			},
			"script":  "rx(?:\\.\\w+)?(?:\\.compat|\\.global)?(?:\\.min)?\\.js",
			"website": "http://reactivex.io",
		},
		"S.Builder": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "S.Builder.png",
			"meta": map[string]interface{}{
				"generator": "S\\.Builder",
			},
			"website": "http://www.sbuilder.ru",
		},
		"SAP": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "SAP NetWeaver Application Server",
			},
			"icon":    "SAP.png",
			"website": "http://sap.com",
		},
		"SDL Tridion": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<img[^>]+_tcm\\d{2,3}-\\d{6}\\.",
			"icon":    "SDL Tridion.png",
			"website": "http://www.sdl.com/products/tridion",
		},
		"Sensors Data": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"js": map[string]interface{}{
				"sa.lib_version":                    "([\\d.]+)\\;version:\\1",
				"sensorsdata_app_js_bridge_call_js": "",
			},
			"cookies": map[string]interface{}{
				"sensorsdata2015session":    "",
				"sensorsdata2015jssdkcross": "",
			},
			"icon":    "Sensors Data.svg",
			"script":  "sensorsdata",
			"website": "https://www.sensorsdata.cn",
		},
		"Sentry": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"html": "<script[^>]*>\\s*Raven\\.config\\('[^']*', {\\s+release: '([0-9\\.]+)'\\;version:\\1",
			"js": map[string]interface{}{
				"Sentry":                     "",
				"Sentry.SDK_VERSION":         "(.+)\\;version:\\1",
				"Raven.config":               "",
				"ravenOptions.whitelistUrls": "",
			},
			"icon":    "Sentry.svg",
			"website": "https://sentry.io/",
		},
		"SIMsite": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "SIMsite.png",
			"meta": map[string]interface{}{
				"SIM.medium": "",
			},
			"script":  "/sim(?:site|core)/js",
			"website": "http://simgroep.nl/internet/portfolio-contentbeheer_41623/",
		},
		"SMF": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html":    "credits/?\" title=\"Simple Machines Forum\" target=\"_blank\" class=\"new_win\">SMF ([0-9.]+)</a>\\;version:\\1",
			"icon":    "SMF.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"smf_": "",
			},
			"website": "http://www.simplemachines.org",
		},
		"SOBI 2": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html":    "(?:<!-- start of Sigsiu Online Business Index|<div[^>]* class=\"sobi2)",
			"icon":    "SOBI 2.png",
			"implies": "Joomla",
			"website": "http://www.sigsiu.net/sobi2.html",
		},
		"SPDY": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"excludes": "HTTP/2",
			"headers": map[string]interface{}{
				"X-Firefox-Spdy": "\\d\\.\\d",
			},
			"icon":    "SPDY.png",
			"website": "http://chromium.org/spdy",
		},
		"SPIP": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"Composed-By":  "SPIP ([\\d.]+) @\\;version:\\1",
				"X-Spip-Cache": "",
			},
			"icon":    "spip.svg",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "(?:^|\\s)SPIP(?:\\s([\\d.]+(?:\\s\\[\\d+\\])?))?\\;version:\\1",
			},
			"website": "http://www.spip.net",
		},
		"SQL Buddy": map[string]interface{}{
			"cats": []interface{}{
				3,
			},
			"html":    "(?:<title>SQL Buddy</title>|<[^>]+onclick=\"sideMainClick\\(\"home\\.php)",
			"icon":    "SQL Buddy.png",
			"implies": "PHP",
			"website": "http://www.sqlbuddy.com",
		},
		"SQLite": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon":    "SQLite.png",
			"website": "http://www.sqlite.org",
		},
		"SUSE": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":       "SUSE(?:/?\\s?-?([\\d.]+))?\\;version:\\1",
				"X-Powered-By": "SUSE(?:/?\\s?-?([\\d.]+))?\\;version:\\1",
			},
			"icon":    "SUSE.png",
			"website": "http://suse.com",
		},
		"SWFObject": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon": "SWFObject.png",
			"js": map[string]interface{}{
				"SWFObject": "",
			},
			"script":  "swfobject.*\\.js",
			"website": "https://github.com/swfobject/swfobject",
		},
		"Saia PCD": map[string]interface{}{
			"cats": []interface{}{
				45,
			},
			"headers": map[string]interface{}{
				"Server": "Saia PCD(?:([/a-z\\d.]+))?\\;version:\\1",
			},
			"icon":    "Saia PCD.png",
			"website": "http://saia-pcd.com",
		},
		"Sails.js": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"sails.sid": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Sails(?:$|[^a-z0-9])",
			},
			"icon":    "Sails.js.svg",
			"implies": "Express",
			"website": "http://sailsjs.org",
		},
		"Salesforce": map[string]interface{}{
			"cats": []interface{}{
				53,
			},
			"cookies": map[string]interface{}{
				"com.salesforce": "",
			},
			"html": "<[^>]+=\"brandQuaternaryFgrs\"",
			"icon": "Salesforce.svg",
			"js": map[string]interface{}{
				"SFDCApp":         "",
				"SFDCCmp":         "",
				"SFDCPage":        "",
				"SFDCSessionVars": "",
			},
			"website": "https://www.salesforce.com",
		},
		"Salesforce Commerce Cloud": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"Server": "Demandware eCommerce Server",
			},
			"html": "<[^>]+demandware\\.edgesuite",
			"icon": "Salesforce.svg",
			"js": map[string]interface{}{
				"dwAnalytics": "",
			},
			"script":  "/demandware\\.static/",
			"website": "http://demandware.com",
		},
		"Sarka-SPIP": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Sarka-SPIP.png",
			"implies": "SPIP",
			"meta": map[string]interface{}{
				"generator": "Sarka-SPIP(?:\\s([\\d.]+))?\\;version:\\1",
			},
			"website": "http://sarka-spip.net",
		},
		"Sazito": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Sazito.svg",
			"js": map[string]interface{}{
				"Sazito": "",
			},
			"meta": map[string]interface{}{
				"generator": "^Sazito",
			},
			"website": "http://sazito.com",
		},
		"Scala": map[string]interface{}{
			"cats": []interface{}{
				27,
			},
			"icon":    "Scala.png",
			"website": "http://www.scala-lang.org",
		},
		"Scholica": map[string]interface{}{
			"cats": []interface{}{
				21,
			},
			"headers": map[string]interface{}{
				"X-Scholica-Version": "",
			},
			"icon":    "Scholica.svg",
			"website": "http://scholica.com",
		},
		"Scientific Linux": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":       "Scientific Linux",
				"X-Powered-By": "Scientific Linux",
			},
			"icon":    "Scientific Linux.png",
			"website": "http://scientificlinux.org",
		},
		"SeamlessCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "SeamlessCMS.png",
			"meta": map[string]interface{}{
				"generator": "^Seamless\\.?CMS",
			},
			"website": "http://www.seamlesscms.com",
		},
		"Segment": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Segment.png",
			"js": map[string]interface{}{
				"analytics": "",
			},
			"script":  "cdn\\.segment\\.com/analytics\\.js",
			"website": "https://segment.com",
		},
		"Select2": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "Select2.png",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"jQuery.fn.select2": "",
			},
			"script":  "select2(?:\\.min|\\.full)?\\.js",
			"website": "https://select2.org/",
		},
		"Semantic-ui": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": []interface{}{
				"<link[^>]+semantic(?:\\.min)\\.css\"",
			},
			"icon":    "Semantic-ui.png",
			"script":  "/semantic(?:-([\\d.]+))?(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://semantic-ui.com",
		},
		"Sencha Touch": map[string]interface{}{
			"cats": []interface{}{
				12,
				26,
			},
			"icon":    "Sencha Touch.png",
			"script":  "sencha-touch.*\\.js",
			"website": "http://sencha.com/products/touch",
		},
		"Serendipity": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"icon":    "Serendipity.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"Powered-By": "Serendipity v\\.([\\d.]+)\\;version:\\1",
				"generator":  "Serendipity(?: v\\.([\\d.]+))?\\;version:\\1",
			},
			"website": "http://s9y.org",
		},
		"Shadow": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "ShadowFramework",
			},
			"icon":    "Shadow.png",
			"implies": "PHP",
			"website": "http://shadow-technologies.co.uk",
		},
		"Shapecss": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": "<link[^>]* href=\"[^\"]*shapecss(?:\\.min)?\\.css",
			"icon": "Shapecss.svg",
			"js": map[string]interface{}{
				"Shapecss": "",
			},
			"script": []interface{}{
				"shapecss[-.]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"/([\\d.]+)/shapecss(?:\\.min)?\\.js\\;version:\\1",
				"shapecss.*\\.js",
			},
			"website": "https://shapecss.com",
		},
		"ShareThis": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon": "ShareThis.png",
			"js": map[string]interface{}{
				"SHARETHIS": "",
			},
			"script":  "w\\.sharethis\\.com/",
			"website": "http://sharethis.com",
		},
		"ShellInABox": map[string]interface{}{
			"cats": []interface{}{
				46,
			},
			"html": []interface{}{
				"<title>Shell In A Box</title>",
				"must be enabled for ShellInABox</noscript>",
			},
			"icon": "ShellInABox.png",
			"js": map[string]interface{}{
				"ShellInABox": "",
			},
			"website": "http://shellinabox.com",
		},
		"Shiny": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"icon": "Shiny.png",
			"js": map[string]interface{}{
				"Shiny.addCustomMessageHandler": "",
			},
			"website": "https://shiny.rstudio.com",
		},
		"ShinyStat": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": "<img[^>]*\\s+src=['\"]?https?://www\\.shinystat\\.com/cgi-bin/shinystat\\.cgi\\?[^'\"\\s>]*['\"\\s/>]",
			"icon": "ShinyStat.png",
			"js": map[string]interface{}{
				"SSsdk": "",
			},
			"script":  "^https?://codice(?:business|ssl|pro|isp)?\\.shinystat\\.com/cgi-bin/getcod\\.cgi",
			"website": "http://shinystat.com",
		},
		"Shopatron": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": []interface{}{
				"<body class=\"shopatron",
				"<img[^>]+mediacdn\\.shopatron\\.com\\;confidence:50",
			},
			"icon": "Shopatron.png",
			"js": map[string]interface{}{
				"shptUrl": "",
			},
			"meta": map[string]interface{}{
				"keywords": "Shopatron",
			},
			"script":  "mediacdn\\.shopatron\\.com",
			"website": "http://ecommerce.shopatron.com",
		},
		"Shopcada": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Shopcada.png",
			"js": map[string]interface{}{
				"Shopcada": "",
			},
			"website": "http://shopcada.com",
		},
		"Shoper": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Shoper.svg",
			"js": map[string]interface{}{
				"shoper": "",
			},
			"website": "https://www.shoper.pl",
		},
		"Shopery": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Shopery": "",
			},
			"icon": "Shopery.svg",
			"implies": []interface{}{
				"PHP",
				"Symfony",
				"Elcodi",
			},
			"website": "http://shopery.com",
		},
		"Shopfa": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"js": map[string]interface{}{
				"shopfa": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^ShopFA ([\\d.]+)$\\;version:\\1",
			},
			"meta": map[string]interface{}{
				"generator": "^ShopFA ([\\d.]+)$\\;version:\\1",
			},
			"icon":    "Shopfa.svg",
			"website": "https://shopfa.com",
		},
		"Shopify": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<link[^>]+=['\"]//cdn\\.shopify\\.com\\;confidence:25",
			"icon": "Shopify.svg",
			"js": map[string]interface{}{
				"Shopify": "\\;confidence:25",
			},
			"headers": map[string]interface{}{
				"x-shopid":        "\\;confidence:50",
				"x-shopify-stage": "\\;confidence:50",
			},
			"url":     "^https?//.+\\.myshopify\\.com",
			"website": "http://shopify.com",
		},
		"Shopline": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "shopline.png",
			"meta": map[string]interface{}{
				"og:image": "https\\:\\/\\/img\\.shoplineapp\\.com",
			},
			"website": "https://shoplineapp.com/",
		},
		"Shoptet": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<link [^>]*href=\"https?://cdn\\.myshoptet\\.com/",
			"icon":    "Shoptet.svg",
			"implies": "PHP",
			"js": map[string]interface{}{
				"shoptet": "",
			},
			"meta": map[string]interface{}{
				"web_author": "^Shoptet",
			},
			"script": []interface{}{
				"^https?://cdn\\.myshoptet\\.com/",
			},
			"website": "http://www.shoptet.cz",
		},
		"Shopware": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<title>Shopware ([\\d\\.]+) [^<]+\\;version:\\1",
			"icon": "Shopware.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
				"jQuery",
			},
			"meta": map[string]interface{}{
				"application-name": "Shopware",
			},
			"script": []interface{}{
				"(?:(shopware)|/web/cache/[0-9]{10}_.+)\\.js\\;version:\\1?4:5",
				"/jquery\\.shopware\\.min\\.js",
				"/engine/Shopware/",
			},
			"website": "http://shopware.com",
		},
		"Signal": map[string]interface{}{
			"cats": []interface{}{
				32,
			},
			"script": []interface{}{
				"//s\\.btstatic\\.com/tag\\.js",
				"//s\\.thebrighttag\\.com/iframe\\?",
			},
			"js": map[string]interface{}{
				"signalData": "",
			},
			"icon":    "signal.png",
			"website": "https://www.signal.co/",
		},
		"Silva": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "SilvaCMS",
			},
			"icon":    "Silva.png",
			"website": "http://silvacms.org",
		},
		"SilverStripe": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "Powered by <a href=\"[^>]+SilverStripe",
			"icon": "SilverStripe.svg",
			"meta": map[string]interface{}{
				"generator": "^SilverStripe",
			},
			"implies": "PHP",
			"website": "https://www.silverstripe.org",
		},
		"SimpleHTTP": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "SimpleHTTP(?:/([\\d.]+))?\\;version:\\1",
			},
			"website": "http://example.com",
		},
		"Simplébo": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-ServedBy": "simplebo",
			},
			"icon":    "Simplebo.png",
			"website": "https://www.simplebo.fr",
		},
		"Site Meter": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "Site Meter.png",
			"script":  "sitemeter\\.com/js/counter\\.js\\?site=",
			"website": "http://www.sitemeter.com",
		},
		"SiteCatalyst": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "SiteCatalyst.png",
			"js": map[string]interface{}{
				"s_INST":     "",
				"s_account":  "",
				"s_code":     "",
				"s_objectID": "",
			},
			"script":  "/s[_-]code.*\\.js",
			"website": "http://www.adobe.com/solutions/digital-marketing.html",
		},
		"SiteEdit": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "SiteEdit.png",
			"meta": map[string]interface{}{
				"generator": "SiteEdit",
			},
			"website": "http://www.siteedit.ru",
		},
		"Sitecore": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"SC_ANALYTICS_GLOBAL_COOKIE": "",
			},
			"html":    "<img[^>]+src=\"[^>]*/~/media/[^>]+\\.ashx",
			"icon":    "Sitecore.png",
			"website": "http://sitecore.net",
		},
		"Sitefinity": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "Sitefinity.svg",
			"implies": "Microsoft ASP.NET",
			"meta": map[string]interface{}{
				"generator": "^Sitefinity (.+)$\\;version:\\1",
			},
			"website": "http://www.sitefinity.com",
		},
		"Sivuviidakko": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Sivuviidakko.png",
			"meta": map[string]interface{}{
				"generator": "Sivuviidakko",
			},
			"website": "http://sivuviidakko.fi",
		},
		"Sizmek": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html":    "(?:<a [^>]*href=\"[^/]*//[^/]*serving-sys\\.com/|<img [^>]*src=\"[^/]*//[^/]*serving-sys\\.com/)",
			"icon":    "Sizmek.png",
			"script":  "serving-sys\\.com/",
			"website": "http://sizmek.com",
		},
		"Slick": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"html":    "<link [^>]+(?:/([\\d.]+)/)?slick-theme\\.css\\;version:\\1",
			"implies": "jQuery",
			"script":  "(?:/([\\d.]+))?/slick(?:\\.min)?\\.js\\;version:\\1",
			"website": "https://kenwheeler.github.io/slick",
		},
		"Slimbox": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"html":    "<link [^>]*href=\"[^/]*slimbox(?:-rtl)?\\.css",
			"icon":    "Slimbox.png",
			"implies": "MooTools",
			"script":  "slimbox\\.js",
			"website": "http://www.digitalia.be/software/slimbox",
		},
		"Slimbox 2": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"html":    "<link [^>]*href=\"[^/]*slimbox2(?:-rtl)?\\.css",
			"icon":    "Slimbox 2.png",
			"implies": "jQuery",
			"script":  "slimbox2\\.js",
			"website": "http://www.digitalia.be/software/slimbox2",
		},
		"Smart Ad Server": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": "<img[^>]+smartadserver\\.com\\/call",
			"icon": "Smart Ad Server.png",
			"js": map[string]interface{}{
				"SmartAdServer": "",
			},
			"website": "http://smartadserver.com",
		},
		"SmartSite": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": "<[^>]+/smartsite\\.(?:dws|shtml)\\?id=",
			"icon": "SmartSite.png",
			"meta": map[string]interface{}{
				"author": "Redacteur SmartInstant",
			},
			"website": "http://www.seneca.nl/pub/Smartsite/Smartsite-Smartsite-iXperion",
		},
		"Smartstore": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "Smartstore.png",
			"script":  "smjslib\\.js",
			"website": "http://smartstore.com",
		},
		"Snap": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Snap/([.\\d]+)\\;version:\\1",
			},
			"icon":    "Snap.png",
			"implies": "Haskell",
			"website": "http://snapframework.com",
		},
		"Snap.svg": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "Snap.svg.png",
			"js": map[string]interface{}{
				"Snap.version": "^(.+)$\\;version:\\1",
			},
			"script":  "snap\\.svg(?:-min)?\\.js",
			"website": "http://snapsvg.io",
		},
		"Snoobi": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Snoobi.png",
			"js": map[string]interface{}{
				"snoobi": "",
			},
			"script":  "snoobi\\.com/snoop\\.php",
			"website": "http://www.snoobi.com",
		},
		"SobiPro": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon":    "SobiPro.png",
			"implies": "Joomla",
			"js": map[string]interface{}{
				"SobiProUrl": "",
			},
			"website": "http://sigsiu.net/sobipro.html",
		},
		"Socket.io": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon":    "Socket.io.png",
			"implies": "Node.js",
			"js": map[string]interface{}{
				"io.Socket":  "",
				"io.version": "^(.+)$\\;version:\\1",
			},
			"script":  "socket\\.io.*\\.js",
			"website": "https://socket.io",
		},
		"SoftTr": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "softtr.png",
			"meta": map[string]interface{}{
				"author": "SoftTr E-Ticaret Sitesi Yazılımı",
			},
			"website": "http://www.softtr.com",
		},
		"Solodev": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"solodev_session": "",
			},
			"html":    "<div class=[\"']dynamicDiv[\"'] id=[\"']dd\\.\\d\\.\\d(?:\\.\\d)?[\"']>",
			"icon":    "Solodev.png",
			"implies": "PHP",
			"website": "http://www.solodev.com",
		},
		"Solr": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"icon":    "Solr.png",
			"implies": "Lucene",
			"website": "http://lucene.apache.org/solr/",
		},
		"Solusquare OmniCommerce Cloud": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"_solusquare": "",
			},
			"icon":    "Solusquare.png",
			"implies": "Adobe ColdFusion",
			"meta": map[string]interface{}{
				"generator": "^Solusquare$",
			},
			"website": "https://www.solusquare.com",
		},
		"Solve Media": map[string]interface{}{
			"cats": []interface{}{
				16,
				36,
			},
			"icon": "Solve Media.png",
			"js": map[string]interface{}{
				"ACPuzzle":                   "",
				"_ACPuzzle":                  "",
				"_adcopy-puzzle-image-image": "",
				"adcopy-puzzle-image-image":  "",
			},
			"script":  "^https?://api\\.solvemedia\\.com/",
			"website": "http://solvemedia.com",
		},
		"SonarQubes": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"html": []interface{}{
				"<link href=\"/css/sonar\\.css\\?v=([\\d.]+)\\;version:\\1",
				"<title>SonarQube</title>",
			},
			"icon":    "sonar.png",
			"implies": "Java",
			"js": map[string]interface{}{
				"SonarMeasures": "",
				"SonarRequest":  "",
			},
			"meta": map[string]interface{}{
				"application-name": "^SonarQubes$",
			},
			"script":  "^/js/bundles/sonar\\.js?v=([\\d.]+)$\\;version:\\1",
			"website": "https://www.sonarqube.org/",
		},
		"SoundManager": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "SoundManager.png",
			"js": map[string]interface{}{
				"BaconPlayer":          "",
				"SoundManager":         "",
				"soundManager.version": "V(.+) \\;version:\\1",
			},
			"website": "http://www.schillmania.com/projects/soundmanager2",
		},
		"Sphinx": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"html": "Created using <a href=\"https?://sphinx-doc\\.org/\">Sphinx</a> ([0-9.]+)\\.\\;version:\\1",
			"icon": "Sphinx.png",
			"js": map[string]interface{}{
				"DOCUMENTATION_OPTIONS": "",
			},
			"website": "http://sphinx.pocoo.org",
		},
		"SpiderControl iniNet": map[string]interface{}{
			"cats": []interface{}{
				45,
			},
			"icon": "SpiderControl iniNet.png",
			"meta": map[string]interface{}{
				"generator": "iniNet SpiderControl",
			},
			"website": "http://spidercontrol.net/ininet",
		},
		"SpinCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"spincms_session": "",
			},
			"icon":    "SpinCMS.png",
			"implies": "PHP",
			"website": "http://www.spin.cw",
		},
		"Splunk": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html": "<p class=\"footer\">&copy; [-\\d]+ Splunk Inc\\.(?: Splunk ([\\d\\.]+(?: build [\\d\\.]*\\d)?))?[^<]*</p>\\;version:\\1",
			"icon": "Splunk.png",
			"meta": map[string]interface{}{
				"author": "Splunk Inc\\;confidence:50",
			},
			"website": "http://splunk.com",
		},
		"Splunkd": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Splunkd",
			},
			"icon":    "Splunk.png",
			"website": "http://splunk.com",
		},
		"Spree": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "(?:<link[^>]*/assets/store/all-[a-z\\d]{32}\\.css[^>]+>|<script>\\s*Spree\\.(?:routes|translations|api_key))",
			"icon":    "Spree.png",
			"implies": "Ruby on Rails",
			"website": "http://spreecommerce.com",
		},
		"Sqreen": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"headers": map[string]interface{}{
				"X-Protected-By": "^Sqreen$",
			},
			"icon":    "Sqreen.png",
			"website": "https://sqreen.io",
		},
		"Squarespace": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-ServedBy": "squarespace",
			},
			"html": "<!-- This is Squarespace\\. -->",
			"icon": "Squarespace.png",
			"js": map[string]interface{}{
				"Squarespace": "",
			},
			"website": "http://www.squarespace.com",
		},
		"SquirrelMail": map[string]interface{}{
			"cats": []interface{}{
				30,
			},
			"html":    "<small>SquirrelMail version ([.\\d]+)[^<]*<br \\;version:\\1",
			"icon":    "SquirrelMail.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"squirrelmail_loginpage_onload": "",
			},
			"url":     "/src/webmail\\.php(?:$|\\?)",
			"website": "http://squirrelmail.org",
		},
		"Squiz Matrix": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Squiz Matrix",
			},
			"html":    "<!--\\s+Running (?:MySource|Squiz) Matrix",
			"icon":    "Squiz Matrix.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Squiz Matrix",
			},
			"website": "http://squiz.net",
		},
		"Stackla": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon": "Stackla.png",
			"js": map[string]interface{}{
				"Stackla": "",
			},
			"script":  "assetscdn\\.stackla\\.com\\/media\\/js\\/widget\\/(?:[a-zA-Z0-9.]+)?\\.js",
			"website": "http://stackla.com/",
		},
		"Starlet": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^Plack::Handler::Starlet",
			},
			"icon":    "Starlet.png",
			"implies": "Perl",
			"website": "http://metacpan.org/pod/Starlet",
		},
		"Statcounter": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Statcounter.svg",
			"js": map[string]interface{}{
				"_statcounter": "\\;confidence:100",
				"sc_project":   "\\;confidence:50",
				"sc_security":  "\\;confidence:50",
			},
			"script":  "statcounter\\.com/counter/counter",
			"website": "http://www.statcounter.com",
		},
		"Store Systems": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "Shopsystem von <a href=[^>]+store-systems\\.de\"|\\.mws_boxTop",
			"icon":    "Store Systems.png",
			"website": "http://store-systems.de",
		},
		"Storeden": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Storeden",
			},
			"icon":    "storeden.svg",
			"website": "https://www.storeden.com",
		},
		"Storyblok": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "storyblok.png",
			"meta": map[string]interface{}{
				"generator": "storyblok",
			},
			"website": "https://www.storyblok.com",
		},
		"Strapdown.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "strapdown.js.png",
			"implies": []interface{}{
				"Bootstrap",
				"Google Code Prettify",
			},
			"script":  "strapdown\\.js",
			"website": "http://strapdownjs.com",
		},
		"Strapi": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Strapi",
			},
			"icon":    "Strapi.png",
			"website": "https://strapi.io",
		},
		"Strato": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<a href=\"http://www\\.strato\\.de/\" target=\"_blank\">",
			"icon":    "strato.png",
			"website": "http://shop.strato.com",
		},
		"Stripe": map[string]interface{}{
			"cats": []interface{}{
				41,
			},
			"html": "<input[^>]+data-stripe",
			"icon": "Stripe.png",
			"js": map[string]interface{}{
				"Stripe.version": "^(.+)$\\;version:\\1",
			},
			"script":  "js\\.stripe\\.com",
			"website": "http://stripe.com",
		},
		"SublimeVideo": map[string]interface{}{
			"cats": []interface{}{
				14,
			},
			"icon": "SublimeVideo.png",
			"js": map[string]interface{}{
				"sublimevideo": "",
			},
			"script":  "cdn\\.sublimevideo\\.net/js/[a-z\\d]+\\.js",
			"website": "http://sublimevideo.net",
		},
		"Subrion": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-CMS": "Subrion CMS",
			},
			"icon":    "Subrion.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "^Subrion ",
			},
			"website": "http://subrion.com",
		},
		"Sucuri": map[string]interface{}{
			"cats": []interface{}{
				31,
			},
			"headers": map[string]interface{}{
				"x-sucuri-cache:": "",
				"x-sucuri-id":     "",
			},
			"icon":    "sucuri.png",
			"website": "https://sucuri.net/",
		},
		"Sulu": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Generator": "Sulu/?(.+)?$\\;version:\\1",
			},
			"icon":    "Sulu.svg",
			"implies": "Symfony",
			"website": "http://sulu.io",
		},
		"SumoMe": map[string]interface{}{
			"cats": []interface{}{
				5,
				32,
			},
			"icon":    "SumoMe.png",
			"script":  "load\\.sumome\\.com",
			"website": "http://sumome.com",
		},
		"SunOS": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":         "SunOS( [\\d\\.]+)?\\;version:\\1",
				"Servlet-engine": "SunOS( [\\d\\.]+)?\\;version:\\1",
			},
			"icon":    "Oracle.png",
			"website": "http://oracle.com/solaris",
		},
		"Supersized": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon":    "Supersized.png",
			"script":  "supersized(?:\\.([\\d.]*[\\d]))?.*\\.js\\;version:\\1",
			"website": "http://buildinternet.com/project/supersized",
		},
		"Svbtle": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"icon": "svbtle.png",
			"meta": map[string]interface{}{
				"generator": "^Svbtle\\.com$",
			},
			"url":     "^https?://[^/]+\\.svbtle\\.com",
			"website": "https://www.svbtle.com",
		},
		"Svelte": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html":    "<[^>]+class=\"[^\"]*svelte-",
			"icon":    "Svelte.svg",
			"website": "https://svelte.dev",
		},
		"SweetAlert": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"html": "<link[^>]+?href=\"[^\"]+sweet-alert(?:\\.min)?\\.css",
			"icon": "SweetAlert.png",
			"js": map[string]interface{}{
				"swal": "",
			},
			"script":  "sweet-alert(?:\\.min)?\\.js",
			"website": "https://t4t5.github.io/sweetalert/",
		},
		"SweetAlert2": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"excludes": "SweetAlert",
			"html":     "<link[^>]+?href=\"[^\"]+sweetalert2(?:\\.min)?\\.css",
			"icon":     "SweetAlert2.png",
			"js": map[string]interface{}{
				"Sweetalert2": "",
			},
			"script":  "sweetalert2(?:\\.all)?(?:\\.min)?\\.js",
			"website": "https://sweetalert2.github.io/",
		},
		"Swiftlet": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Generator":      "Swiftlet",
				"X-Powered-By":     "Swiftlet",
				"X-Swiftlet-Cache": "",
			},
			"html":    "Powered by <a href=\"[^>]+Swiftlet",
			"icon":    "Swiftlet.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "Swiftlet",
			},
			"website": "http://swiftlet.org",
		},
		"Swiftype": map[string]interface{}{
			"cats": []interface{}{
				29,
			},
			"icon": "swiftype.png",
			"js": map[string]interface{}{
				"Swiftype": "",
			},
			"script":  "swiftype\\.com/embed\\.js$",
			"website": "http://swiftype.com",
		},
		"Symfony": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"icon":    "Symfony.png",
			"implies": "PHP",
			"website": "http://symfony.com",
		},
		"Sympa": map[string]interface{}{
			"cats": []interface{}{
				30,
			},
			"meta": map[string]interface{}{
				"generator": "^Sympa$",
			},
			"html":    "<a href=\"https?://www\\.sympa\\.org\">\\s*Powered by Sympa\\s*</a>",
			"icon":    "sympa.png",
			"implies": "Perl",
			"website": "https://www.sympa.org/",
		},
		"Synology DiskStation": map[string]interface{}{
			"cats": []interface{}{
				48,
			},
			"html": "<noscript><div class='syno-no-script'",
			"icon": "Synology DiskStation.png",
			"meta": map[string]interface{}{
				"application-name": "Synology DiskStation",
				"description":      "^DiskStation provides a full-featured network attached storage",
			},
			"script":  "webapi/entry\\.cgi\\?api=SYNO\\.(?:Core|Filestation)\\.Desktop\\.",
			"website": "http://synology.com",
		},
		"SyntaxHighlighter": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html": "<(?:script|link)[^>]*sh(?:Core|Brush|ThemeDefault)",
			"icon": "SyntaxHighlighter.png",
			"js": map[string]interface{}{
				"SyntaxHighlighter": "",
			},
			"website": "https://github.com/syntaxhighlighter",
		},
		"TWiki": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"cookies": map[string]interface{}{
				"TWIKISID": "",
			},
			"html":    "<img [^>]*(?:title|alt)=\"This site is powered by the TWiki collaboration platform",
			"icon":    "TWiki.png",
			"implies": "Perl",
			"script":  "(?:TWikiJavascripts|twikilib(?:\\.min)?\\.js)",
			"website": "http://twiki.org",
		},
		"tailwindcss": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "<link[^>]+?href=\"[^\"]+tailwindcss(?:\\.min)?\\.css",
			"icon":    "tailwindcss.svg",
			"website": "https://tailwindcss.com/",
		},
		"TYPO3 CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": []interface{}{
				"<link[^>]+ href=\"typo3(?:conf|temp)/",
				"<img[^>]+ src=\"typo3(?:conf|temp)/",
			},
			"icon":    "TYPO3.svg",
			"implies": "PHP",
			"script":  "^typo3(?:conf|temp)/",
			"meta": map[string]interface{}{
				"generator": "TYPO3\\s+(?:CMS\\s+)?([\\d.]+)?(?:\\s+CMS)?\\;version:\\1",
			},
			"url":     "/typo3/",
			"website": "https://typo3.org/",
		},
		"Taiga": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"icon": "Taiga.png",
			"implies": []interface{}{
				"Django",
				"AngularJS",
			},
			"js": map[string]interface{}{
				"taigaConfig": "",
			},
			"website": "http://taiga.io",
		},
		"Tawk.to": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon":    "TawkTo.png",
			"script":  "//embed\\.tawk\\.to",
			"website": "http://tawk.to",
		},
		"Tealeaf": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Tealeaf.png",
			"js": map[string]interface{}{
				"TeaLeaf": "",
			},
			"website": "http://www.tealeaf.com",
		},
		"Tealium": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Tealium.png",
			"js": map[string]interface{}{
				"TEALIUMENABLED": "",
			},
			"script": []interface{}{
				"^(?:https?:)?//tags\\.tiqcdn\\.com/",
				"/tealium/utag\\.js$",
			},
			"website": "http://tealium.com",
		},
		"TeamCity": map[string]interface{}{
			"cats": []interface{}{
				44,
			},
			"html": "<span class=\"versionTag\"><span class=\"vWord\">Version</span> ([\\d\\.]+)\\;version:\\1",
			"icon": "TeamCity.svg",
			"implies": []interface{}{
				"Apache Tomcat",
				"Java",
				"jQuery",
				"Moment.js",
				"Prototype",
				"React",
				"Underscore.js",
			},
			"meta": map[string]interface{}{
				"application-name": "TeamCity",
			},
			"website": "https://www.jetbrains.com/teamcity/",
		},
		"Telescope": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Telescope.png",
			"implies": []interface{}{
				"Meteor",
				"React",
			},
			"js": map[string]interface{}{
				"Telescope": "",
			},
			"website": "http://telescopeapp.org",
		},
		"TN Express Web": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"TNEW": "",
			},
			"implies": []interface{}{
				"Tessitura",
			},
			"icon":    "tessitura.svg",
			"website": "https://www.tessituranetwork.com",
		},
		"Tessitura": map[string]interface{}{
			"cats": []interface{}{
				53,
			},
			"html": "<!--[^>]+Tessitura Version: (\\d*\\.\\d*\\.\\d*)?\\;version:\\1",
			"implies": []interface{}{
				"Microsoft ASP.NET",
				"IIS",
				"Windows Server",
			},
			"icon":    "tessitura.svg",
			"website": "https://www.tessituranetwork.com",
		},
		"Tengine": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Tengine",
			},
			"icon":    "Tengine.png",
			"website": "http://tengine.taobao.org",
		},
		"Textalk": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "textalk.png",
			"meta": map[string]interface{}{
				"generator": "Textalk Webshop",
			},
			"website": "https://www.textalk.se",
		},
		"Textpattern CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Textpattern CMS.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"meta": map[string]interface{}{
				"generator": "Textpattern",
			},
			"website": "http://textpattern.com",
		},
		"Thelia": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"html": "<(?:link|style|script)[^>]+/assets/frontOffice/",
			"icon": "Thelia.png",
			"implies": []interface{}{
				"PHP",
				"Symfony",
			},
			"website": "http://thelia.net",
		},
		"ThinkPHP": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "ThinkPHP",
			},
			"icon":    "ThinkPHP.png",
			"implies": "PHP",
			"website": "http://www.thinkphp.cn",
		},
		"Ticimax": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Ticimax.png",
			"script": []interface{}{
				"cdn\\.ticimax\\.com/",
			},
			"website": "https://www.ticimax.com",
		},
		"Tictail": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<link[^>]*tictail\\.com",
			"icon":    "tictail.png",
			"website": "https://tictail.com",
		},
		"TiddlyWiki": map[string]interface{}{
			"cats": []interface{}{
				1,
				2,
				4,
				8,
			},
			"html": "<[^>]*type=[^>]text\\/vnd\\.tiddlywiki",
			"icon": "TiddlyWiki.png",
			"js": map[string]interface{}{
				"tiddler": "",
			},
			"meta": map[string]interface{}{
				"application-name":   "^TiddlyWiki$",
				"copyright":          "^TiddlyWiki created by Jeremy Ruston",
				"generator":          "^TiddlyWiki$",
				"tiddlywiki-version": "^(.+)$\\;version:\\1",
			},
			"website": "http://tiddlywiki.com",
		},
		"Tiki Wiki CMS Groupware": map[string]interface{}{
			"cats": []interface{}{
				1,
				2,
				8,
				11,
				13,
			},
			"icon": "Tiki Wiki CMS Groupware.png",
			"meta": map[string]interface{}{
				"generator": "^Tiki",
			},
			"script":  "(?:/|_)tiki",
			"website": "http://tiki.org",
		},
		"Tilda": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<link[^>]* href=[^>]+tilda(?:cdn|\\.ws|-blocks)",
			"icon":    "Tilda.svg",
			"script":  "tilda(?:cdn|\\.ws|-blocks)",
			"website": "https://tilda.cc",
		},
		"Timeplot": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "Timeplot.png",
			"js": map[string]interface{}{
				"Timeplot": "",
			},
			"script":  "timeplot.*\\.js",
			"website": "http://www.simile-widgets.org/timeplot/",
		},
		"TinyMCE": map[string]interface{}{
			"cats": []interface{}{
				24,
			},
			"icon": "TinyMCE.png",
			"js": map[string]interface{}{
				"tinyMCE.majorVersion": "([\\d.]+)\\;version:\\1",
			},
			"script":  "/tiny_?mce(?:\\.min)?\\.js",
			"website": "http://tinymce.com",
		},
		"Titan": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "Titan.png",
			"js": map[string]interface{}{
				"titan":        "",
				"titanEnabled": "",
			},
			"website": "http://titan360.com",
		},
		"TomatoCart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "TomatoCart.png",
			"js": map[string]interface{}{
				"AjaxShoppingCart": "",
			},
			"meta": map[string]interface{}{
				"generator": "TomatoCart",
			},
			"website": "http://tomatocart.com",
		},
		"TornadoServer": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "TornadoServer(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "TornadoServer.png",
			"website": "http://tornadoweb.org",
		},
		"TotalCode": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^TotalCode$",
			},
			"icon":    "TotalCode.png",
			"website": "http://www.totalcode.com",
		},
		"Trac": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"html": []interface{}{
				"<a id=\"tracpowered",
				"Powered by <a href=\"[^\"]*\"><strong>Trac(?:[ /]([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Trac.png",
			"implies": "Python",
			"website": "http://trac.edgewall.org",
		},
		"TrackJs": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "TrackJs.png",
			"js": map[string]interface{}{
				"TrackJs": "",
			},
			"script":  "tracker\\.js",
			"website": "http://trackjs.com",
		},
		"Transifex": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "transifex.png",
			"js": map[string]interface{}{
				"Transifex.live.lib_version": "(.+)\\;version:\\1",
			},
			"website": "https://www.transifex.com",
		},
		"Translucide": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "translucide.svg",
			"implies": []interface{}{
				"PHP",
				"jQuery",
			},
			"script":  "lucide\\.init(?:\\.min)?\\.js",
			"website": "http://www.translucide.net",
		},
		"Tray": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "tray.png",
			"script":  "tcdn\\.com\\.br",
			"website": "https://www.tray.com.br",
		},
		"Tumblr": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"headers": map[string]interface{}{
				"X-Tumblr-User": "",
			},
			"html":    "<iframe src=\"[^>]+tumblr\\.com",
			"icon":    "Tumblr.png",
			"url":     "^https?://(?:www\\.)?[^/]+\\.tumblr\\.com/",
			"website": "http://www.tumblr.com",
		},
		"TweenMax": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "TweenMax.png",
			"js": map[string]interface{}{
				"TweenMax.version": "^(.+)$\\;version:\\1",
			},
			"script":  "TweenMax(?:\\.min)?\\.js",
			"website": "http://greensock.com/tweenmax",
		},
		"Twilight CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-CMS": "Twilight CMS",
			},
			"icon":    "Twilight CMS.png",
			"website": "http://www.twilightcms.com",
		},
		"TwistPHP": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "TwistPHP",
			},
			"icon":    "TwistPHP.png",
			"implies": "PHP",
			"website": "http://twistphp.com",
		},
		"TwistedWeb": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "TwistedWeb(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "TwistedWeb.png",
			"website": "http://twistedmatrix.com/trac/wiki/TwistedWeb",
		},
		"Twitter": map[string]interface{}{
			"cats": []interface{}{
				5,
			},
			"icon":    "Twitter.svg",
			"script":  "//platform\\.twitter\\.com/widgets\\.js",
			"website": "http://twitter.com",
		},
		"Twitter Emoji (Twemoji)": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"js": map[string]interface{}{
				"twemoji": "",
			},
			"script":  "twemoji(?:\\.min)?\\.js",
			"website": "https://twitter.github.io/twemoji/",
		},
		"Twitter Flight": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon":    "Twitter Flight.png",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"flight": "",
			},
			"website": "https://flightjs.github.io/",
		},
		"Twitter typeahead.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "Twitter typeahead.js.png",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"typeahead": "",
			},
			"script":  "(?:typeahead|bloodhound)\\.(?:jquery|bundle)?(?:\\.min)?\\.js",
			"website": "https://twitter.github.io/typeahead.js",
		},
		"TypePad": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"icon": "TypePad.png",
			"meta": map[string]interface{}{
				"generator": "typepad",
			},
			"url":     "typepad\\.com",
			"website": "http://www.typepad.com",
		},
		"Typecho": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"icon":    "typecho.svg",
			"implies": "PHP",
			"js": map[string]interface{}{
				"TypechoComment": "",
			},
			"meta": map[string]interface{}{
				"generator": "Typecho( [\\d.]+)?\\;version:\\1",
			},
			"url":     "/admin/login\\.php?referer=http%3A%2F%2F",
			"website": "http://typecho.org/",
		},
		"Typekit": map[string]interface{}{
			"cats": []interface{}{
				17,
			},
			"icon": "Typekit.png",
			"js": map[string]interface{}{
				"Typekit.config.js": "^(.+)$\\;version:\\1",
			},
			"script":  "use\\.typekit\\.com",
			"website": "http://typekit.com",
		},
		"UIKit": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "<[^>]+class=\"[^\"]*(?:uk-container|uk-section)",
			"icon":    "UIKit.png",
			"script":  "uikit.*\\.js",
			"website": "https://getuikit.com",
		},
		"UMI.CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Generated-By": "UMI\\.CMS",
			},
			"icon":    "UMI.CMS.png",
			"implies": "PHP",
			"website": "https://www.umi-cms.ru",
		},
		"UNIX": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server": "Unix",
			},
			"icon":    "UNIX.png",
			"website": "http://unix.org",
		},
		"Ubercart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "Ubercart.png",
			"implies": "Drupal",
			"script":  "uc_cart/uc_cart_block\\.js",
			"website": "http://www.ubercart.org",
		},
		"Ubuntu": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server":       "Ubuntu",
				"X-Powered-By": "Ubuntu",
			},
			"icon":    "Ubuntu.png",
			"website": "http://www.ubuntu.com/server",
		},
		"UltraCart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<form [^>]*action=\"[^\"]*\\/cgi-bin\\/UCEditor\\?(?:[^\"]*&)?merchantId=[^\"]",
			"icon": "UltraCart.png",
			"js": map[string]interface{}{
				"ucCatalog": "",
			},
			"script":  "cgi-bin\\/UCJavaScript\\?",
			"url":     "/cgi-bin/UCEditor\\?",
			"website": "http://ultracart.com",
		},
		"Umbraco": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Umbraco-Version": "^(.+)$\\;version:\\1",
			},
			"html":    "powered by <a href=[^>]+umbraco",
			"icon":    "Umbraco.png",
			"implies": "Microsoft ASP.NET",
			"js": map[string]interface{}{
				"UC_IMAGE_SERVICE|ITEM_INFO_SERVICE": "",
				"UC_ITEM_INFO_SERVICE":               "",
				"UC_SETTINGS":                        "",
				"Umbraco":                            "",
			},
			"meta": map[string]interface{}{
				"generator": "umbraco",
			},
			"url":     "/umbraco/login\\.aspx(?:$|\\?)",
			"website": "http://umbraco.com",
		},
		"Unbounce": map[string]interface{}{
			"cats": []interface{}{
				20,
				51,
			},
			"headers": map[string]interface{}{
				"X-Unbounce-PageId": "",
			},
			"icon":    "Unbounce.png",
			"script":  "ubembed\\.com",
			"website": "http://unbounce.com",
		},
		"Underscore.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"excludes": "Lodash",
			"icon":     "Underscore.js.png",
			"js": map[string]interface{}{
				"_.VERSION":       "^(.+)$\\;confidence:0\\;version:\\1",
				"_.restArguments": "",
			},
			"script":  "underscore.*\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1",
			"website": "http://underscorejs.org",
		},
		"Usabilla": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"icon": "Usabilla.svg",
			"js": map[string]interface{}{
				"usabilla_live": "",
			},
			"website": "http://usabilla.com",
		},
		"user.com": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": "<div[^>]+/id=\"ue_widget\"",
			"icon": "user.com.svg",
			"js": map[string]interface{}{
				"UserEngage": "",
			},
			"website": "https://user.com",
		},
		"UserGuiding": map[string]interface{}{
			"cats": []interface{}{
				58,
				10,
			},
			"icon": "UserGuiding.svg",
			"implies": []interface{}{
				"React",
				"webpack",
				"Node.js",
			},
			"script":  "static\\.userguiding*\\.js",
			"website": "https://userguiding.com",
		},
		"UserLike": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon": "UserLike.svg",
			"script": []interface{}{
				"userlike\\.min\\.js",
				"userlikelib\\.min\\.js",
			},
			"website": "http://userlike.com",
		},
		"UserRules": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"icon": "UserRules.png",
			"js": map[string]interface{}{
				"_usrp": "",
			},
			"website": "http://www.userrules.com",
		},
		"UserVoice": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"icon": "UserVoice.png",
			"js": map[string]interface{}{
				"UserVoice": "",
			},
			"website": "http://uservoice.com",
		},
		"Ushahidi": map[string]interface{}{
			"cats": []interface{}{
				1,
				35,
			},
			"cookies": map[string]interface{}{
				"ushahidi": "",
			},
			"icon": "Ushahidi.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
				"OpenLayers",
			},
			"js": map[string]interface{}{
				"Ushahidi": "",
			},
			"script":  "/js/ushahidi\\.js$",
			"website": "http://www.ushahidi.com",
		},
		"VIVVO": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"VivvoSessionId": "",
			},
			"icon": "VIVVO.png",
			"js": map[string]interface{}{
				"vivvo": "",
			},
			"website": "http://vivvo.net",
		},
		"VP-ASP": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<a[^>]+>Powered By VP-ASP Shopping Cart</a>",
			"icon":    "VP-ASP.png",
			"implies": "Microsoft ASP.NET",
			"script":  "vs350\\.js",
			"website": "http://www.vpasp.com",
		},
		"VTEX": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"VtexWorkspace":   "",
				"VtexFingerPrint": "",
				"vtex_session":    "",
			},
			"headers": map[string]interface{}{
				"Server":  "^VTEX IO$",
				"powered": "vtex",
			},
			"icon":    "VTEX.svg",
			"website": "https://vtex.com/",
		},
		"VTEX Integrated Store": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "vtex-integrated-store",
			},
			"icon":    "VTEX Integrated Store.png",
			"website": "http://lojaintegrada.com.br",
		},
		"Vaadin": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"icon":    "Vaadin.svg",
			"implies": "Java",
			"js": map[string]interface{}{
				"vaadin": "",
			},
			"script":  "vaadinBootstrap\\.js(?:\\?v=([\\d.]+))?\\;version:\\1",
			"website": "https://vaadin.com",
		},
		"Vanilla": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Vanilla",
			},
			"html":    "<body id=\"(?:DiscussionsPage|vanilla)",
			"icon":    "Vanilla.png",
			"implies": "PHP",
			"website": "http://vanillaforums.org",
		},
		"Varnish": map[string]interface{}{
			"cats": []interface{}{
				23,
			},
			"headers": map[string]interface{}{
				"Via":                "varnish(?: \\(Varnish/([\\d.]+)\\))?\\;version:\\1",
				"X-Varnish":          "",
				"X-Varnish-Action":   "",
				"X-Varnish-Age":      "",
				"X-Varnish-Cache":    "",
				"X-Varnish-Hostname": "",
			},
			"icon":    "Varnish.svg",
			"website": "http://www.varnish-cache.org",
		},
		"Venda": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-venda-hitid": "",
			},
			"icon":    "Venda.png",
			"website": "http://venda.com",
		},
		"Veoxa": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": "<img [^>]*src=\"[^\"]+tracking\\.veoxa\\.com",
			"icon": "Veoxa.png",
			"js": map[string]interface{}{
				"VuVeoxaContent": "",
			},
			"script":  "tracking\\.veoxa\\.com",
			"website": "http://veoxa.com",
		},
		"VideoJS": map[string]interface{}{
			"cats": []interface{}{
				14,
			},
			"html": "<div[^>]+class=\"video-js+\">",
			"icon": "VideoJS.png",
			"js": map[string]interface{}{
				"VideoJS": "",
			},
			"script":  "zencdn\\.net/c/video\\.js",
			"website": "http://videojs.com",
		},
		"VigLink": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon": "VigLink.png",
			"js": map[string]interface{}{
				"vglnk":      "",
				"vl_cB":      "",
				"vl_disable": "",
			},
			"script":  "(?:^[^/]*//[^/]*viglink\\.com/api/|vglnk\\.js)",
			"website": "http://viglink.com",
		},
		"Vigbo": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"_gphw_mode": "",
			},
			"html":    "<link[^>]* href=[^>]+(?:\\.vigbo\\.com|\\.gophotoweb\\.com)",
			"icon":    "vigbo.png",
			"script":  "(?:\\.vigbo\\.com|\\.gophotoweb\\.com)",
			"website": "https://vigbo.com",
		},
		"Vignette": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<[^>]+=\"vgn-?ext",
			"icon":    "Vignette.png",
			"website": "http://www.vignette.com",
		},
		"Vimeo": map[string]interface{}{
			"cats": []interface{}{
				14,
			},
			"html":    "(?:<(?:param|embed)[^>]+vimeo\\.com/moogaloop|<iframe[^>]player\\.vimeo\\.com)",
			"icon":    "Vimeo.png",
			"website": "http://vimeo.com",
		},
		"VirtueMart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<div id=\"vmMainPage",
			"icon":    "VirtueMart.png",
			"implies": "Joomla",
			"website": "http://virtuemart.net",
		},
		"Virtuoso": map[string]interface{}{
			"cats": []interface{}{
				34,
			},
			"headers": map[string]interface{}{
				"Server": "Virtuoso/?([0-9.]+)?\\;version:\\1",
			},
			"meta": map[string]interface{}{
				"Copyright": "^Copyright &copy; \\d{4} OpenLink Software",
				"Keywords":  "^OpenLink Virtuoso Sparql",
			},
			"url":     "/sparql",
			"website": "https://virtuoso.openlinksw.com/",
		},
		"Visual WebGUI": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"icon":    "Visual WebGUI.png",
			"implies": "Microsoft ASP.NET",
			"js": map[string]interface{}{
				"VWGEventArgs": "",
			},
			"meta": map[string]interface{}{
				"generator": "^Visual WebGUI",
			},
			"script":  "\\.js\\.wgx$",
			"url":     "\\.wgx$",
			"website": "http://www.gizmox.com/products/visual-web-gui/",
		},
		"Visual Website Optimizer": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": []interface{}{
				"<!-- (?:Start|End) Visual Website Optimizer A?Synchronous Code -->",
			},
			"icon": "vwo.svg",
			"js": map[string]interface{}{
				"VWO":   "",
				"__vwo": "",
			},
			"script": []interface{}{
				"dev\\.visualwebsiteoptimizer\\.com",
			},
			"website": "https://vwo.com/",
		},
		"VisualPath": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "VisualPath.png",
			"script":  "visualpath[^/]*\\.trackset\\.it/[^/]+/track/include\\.js",
			"website": "http://www.trackset.com/web-analytics-software/visualpath",
		},
		"Volusion (V1)": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<link [^>]*href=\"[^\"]*/vspfiles/",
			"icon":    "Volusion.svg",
			"implies": "Microsoft ASP.NET",
			"js": map[string]interface{}{
				"volusion": "",
			},
			"script":  "/volusion\\.js(?:\\?([\\d.]*))?\\;version:\\1",
			"website": "https://www.volusion.com",
		},
		"Volusion (V2)": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<body [^>]*data-vn-page-name",
			"icon":    "Volusion.svg",
			"implies": "AngularJS",
			"website": "https://www.volusion.com",
		},
		"Vue.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html": "<[^>]+data-v(?:ue)-",
			"icon": "Vue.js.png",
			"js": map[string]interface{}{
				"Vue.version": "^(.+)$\\;version:\\1",
			},
			"script": []interface{}{
				"vue[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"(?:/([\\d.]+))?/vue(?:\\.min)?\\.js\\;version:\\1",
			},
			"website": "http://vuejs.org",
		},
		"Nuxt.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"html": []interface{}{
				"<div [^>]*id=\"__nuxt\"",
				"<script [^>]*>window\\.__NUXT__",
			},
			"icon": "Nuxt.js.svg",
			"js": map[string]interface{}{
				"$nuxt": "",
			},
			"script": []interface{}{
				"/_nuxt/",
			},
			"implies": "Vue.js",
			"website": "https://nuxtjs.org",
		},
		"W3 Total Cache": map[string]interface{}{
			"cats": []interface{}{
				23,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "W3 Total Cache(?:/([\\d.]+))?\\;version:\\1",
			},
			"html":    "<!--[^>]+W3 Total Cache",
			"icon":    "W3 Total Cache.png",
			"implies": "WordPress",
			"website": "http://www.w3-edge.com/wordpress-plugins/w3-total-cache",
		},
		"W3Counter": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "W3Counter.png",
			"script":  "w3counter\\.com/tracker\\.js",
			"website": "http://www.w3counter.com",
		},
		"WEBXPAY": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "Powered by <a href=\"https://www\\.webxpay\\.com\">WEBXPAY<",
			"icon": "WEBXPAY.png",
			"js": map[string]interface{}{
				"WEBXPAY": "",
			},
			"website": "https://webxpay.com",
		},
		"WHMCS": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"WHMCS": "",
			},
			"icon":    "WHMCS.png",
			"website": "http://www.whmcs.com",
		},
		"WP Rocket": map[string]interface{}{
			"cats": []interface{}{
				23,
			},
			"headers": map[string]interface{}{
				"X-Powered-By":          "WP Rocket(?:/([\\d.]+))?\\;version:\\1",
				"X-Rocket-Nginx-Bypass": "",
			},
			"html":    "<!--[^>]+WP Rocket",
			"icon":    "WP Rocket.png",
			"implies": "WordPress",
			"website": "http://wp-rocket.me",
		},
		"WP Engine": map[string]interface{}{
			"cats": []interface{}{
				62,
			},
			"headers": map[string]interface{}{
				"wpe-backend": "",
			},
			"icon":    "wpengine.svg",
			"implies": "WordPress",
			"website": "https://wpengine.com",
		},
		"Warp": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^Warp/(\\d+(?:\\.\\d+)+)?$\\;version:\\1",
			},
			"icon":    "Warp.png",
			"implies": "Haskell",
			"website": "http://www.stackage.org/package/warp",
		},
		"Web2py": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "web2py",
			},
			"icon": "Web2py.png",
			"implies": []interface{}{
				"Python",
				"jQuery",
			},
			"meta": map[string]interface{}{
				"generator": "^Web2py",
			},
			"script":  "web2py\\.js",
			"website": "http://web2py.com",
		},
		"WebGUI": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"wgSession": "",
			},
			"icon":    "WebGUI.png",
			"implies": "Perl",
			"meta": map[string]interface{}{
				"generator": "^WebGUI ([\\d.]+)\\;version:\\1",
			},
			"website": "http://www.webgui.org",
		},
		"WebPublisher": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "WebPublisher.png",
			"meta": map[string]interface{}{
				"generator": "WEB\\|Publisher",
			},
			"website": "http://scannet.dk",
		},
		"WebSite X5": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"icon": "WebSite X5.png",
			"meta": map[string]interface{}{
				"generator": "Incomedia WebSite X5 (\\w+ [\\d.]+)\\;version:\\1",
			},
			"website": "http://websitex5.com",
		},
		"Webdev": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"headers": map[string]interface{}{
				"WebDevSrc": "",
			},
			"html": "<!-- [a-zA-Z0-9_]+ [\\d/]+ [\\d:]+ WebDev \\d\\d ([\\d.]+) -->\\;version:\\1",
			"icon": "webdev.png",
			"meta": map[string]interface{}{
				"generator": "^WEBDEV$",
			},
			"website": "https://www.windev.com/webdev/index.html",
		},
		"Webix": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "Webix.png",
			"js": map[string]interface{}{
				"webix": "",
			},
			"script":  "\\bwebix\\.js",
			"website": "http://webix.com",
		},
		"Webmine": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"html":    "<iframe[^>]+src=[\\'\"]https://webmine\\.cz/miner\\?key=",
			"icon":    "webmine.png",
			"website": "https://webmine.cz/",
		},
		"Webs": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"Server": "Webs\\.com/?([\\d\\.]+)?\\;version:\\1",
			},
			"icon":    "Webs.png",
			"website": "http://webs.com",
		},
		"Websocket": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html": []interface{}{
				"<link[^>]+rel=[\"']web-socket[\"']",
				"<(?:link|a)[^>]+href=[\"']wss?://",
			},
			"icon":    "websocket.png",
			"website": "https://en.wikipedia.org/wiki/WebSocket",
		},
		"WebsPlanet": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "WebsPlanet.png",
			"meta": map[string]interface{}{
				"generator": "WebsPlanet",
			},
			"website": "http://websplanet.com",
		},
		"Websale": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "Websale.png",
			"url":     "/websale7/",
			"website": "http://websale.de",
		},
		"Website Creator": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "WebsiteCreator.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
				"Vue.js",
			},
			"meta": map[string]interface{}{
				"generator":      "Website Creator by hosttech",
				"wsc_rendermode": "",
			},
			"website": "https://www.hosttech.ch/websitecreator",
		},
		"WebsiteBaker": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "WebsiteBaker.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"meta": map[string]interface{}{
				"generator": "WebsiteBaker",
			},
			"website": "http://websitebaker2.org/en/home.php",
		},
		"Webtrekk": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Webtrekk.png",
			"js": map[string]interface{}{
				"webtrekk": "",
			},
			"website": "http://www.webtrekk.com",
		},
		"Webtrends": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": "<img[^>]+id=\"DCSIMG\"[^>]+webtrends",
			"icon": "Webtrends.png",
			"js": map[string]interface{}{
				"WTOptimize": "",
				"WebTrends":  "",
			},
			"website": "http://worldwide.webtrends.com",
		},
		"Weebly": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Weebly.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"js": map[string]interface{}{
				"_W.configDomain": "",
			},
			"script":  "cdn\\d+\\.editmysite\\.com",
			"website": "https://www.weebly.com",
		},
		"Weglot": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"headers": map[string]interface{}{
				"Weglot-Translated": "",
			},
			"icon": "Weglot.png",
			"script": []interface{}{
				"cdn\\.weglot\\.com",
				"wp-content/plugins/weglot",
			},
			"website": "https://www.weglot.com",
		},
		"Webzi": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "Webzi.svg",
			"js": map[string]interface{}{
				"Webzi": "",
			},
			"meta": map[string]interface{}{
				"generator": "^Webzi",
			},
			"script":  "cdn\\.6th\\.ir",
			"website": "https://webzi.ir",
		},
		"Wikinggruppen": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": []interface{}{
				"<!-- WIKINGGRUPPEN",
			},
			"icon":    "wikinggruppen.png",
			"website": "https://wikinggruppen.se/",
		},
		"WikkaWiki": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"html": "Powered by <a href=\"[^>]+WikkaWiki",
			"icon": "WikkaWiki.png",
			"meta": map[string]interface{}{
				"generator": "WikkaWiki",
			},
			"website": "http://wikkawiki.org",
		},
		"Windows CE": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server": "\\bWinCE\\b",
			},
			"icon":    "Microsoft.png",
			"website": "http://microsoft.com",
		},
		"Windows Server": map[string]interface{}{
			"cats": []interface{}{
				28,
			},
			"headers": map[string]interface{}{
				"Server": "Win32|Win64",
			},
			"icon":    "WindowsServer.png",
			"website": "http://microsoft.com/windowsserver",
		},
		"Wink": map[string]interface{}{
			"cats": []interface{}{
				26,
				12,
			},
			"icon": "Wink.png",
			"js": map[string]interface{}{
				"wink.version": "^(.+)$\\;version:\\1",
			},
			"script":  "(?:_base/js/base|wink).*\\.js",
			"website": "http://winktoolkit.org",
		},
		"Winstone Servlet Container": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server":       "Winstone Servlet (?:Container|Engine) v?([\\d.]+)?\\;version:\\1",
				"X-Powered-By": "Winstone(?:.([\\d.]+))?\\;version:\\1",
			},
			"website": "http://winstone.sourceforge.net",
		},
		"Wix": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
				11,
			},
			"cookies": map[string]interface{}{
				"Domain": "\\.wix\\.com",
			},
			"headers": map[string]interface{}{
				"X-Wix-Renderer-Server":    "",
				"X-Wix-Request-Id":         "",
				"X-Wix-Server-Artifact-Id": "",
			},
			"icon": "Wix.png",
			"implies": []interface{}{
				"React",
			},
			"js": map[string]interface{}{
				"wixBiSession": "",
			},
			"meta": map[string]interface{}{
				"generator": "Wix\\.com Website Builder",
			},
			"script":  "static\\.parastorage\\.com",
			"website": "https://www.wix.com",
		},
		"Wolf CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "(?:<a href=\"[^>]+wolfcms\\.org[^>]+>Wolf CMS(?:</a>)? inside|Thank you for using <a[^>]+>Wolf CMS)",
			"icon":    "Wolf CMS.png",
			"implies": "PHP",
			"website": "http://www.wolfcms.org",
		},
		"Woltlab Community Framework": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "var WCF_PATH[^>]+",
			"icon":    "Woltlab Community Framework.png",
			"implies": "PHP",
			"script":  "WCF\\..*\\.js",
			"website": "http://www.woltlab.com",
		},
		"WooCommerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": []interface{}{
				"<!-- WooCommerce",
				"<link rel='[^']+' id='woocommerce-(?:layout|smallscreen|general)-css'  href='https?://[^/]+/wp-content/plugins/woocommerce/assets/css/woocommerce(?:-layout|-smallscreen)?\\.css?ver=([\\d.]+)'\\;version:\\1",
			},
			"icon":    "WooCommerce.png",
			"implies": "WordPress",
			"js": map[string]interface{}{
				"woocommerce_params": "",
			},
			"meta": map[string]interface{}{
				"generator": "WooCommerce ([\\d.]+)\\;version:\\1",
			},
			"script":  "/woocommerce(?:\\.min)?\\.js(?:\\?ver=([0-9.]+))?\\;version:\\1",
			"website": "http://www.woothemes.com/woocommerce",
		},
		"Woopra": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon":    "Woopra.png",
			"script":  "static\\.woopra\\.com",
			"website": "http://www.woopra.com",
		},
		"Woosa": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"excludes": []interface{}{
				"WordPress",
				"WooCommerce",
			},
			"icon": "Woosa.png",
			"meta": map[string]interface{}{
				"generator": "^Woosa$",
			},
			"website": "https://woosa.nl",
		},
		"WordPress": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"html": []interface{}{
				"<link rel=[\"']stylesheet[\"'] [^>]+/wp-(?:content|includes)/",
				"<link[^>]+s\\d+\\.wp\\.com",
			},
			"icon": "WordPress.svg",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"headers": map[string]interface{}{
				"link": "rel=\"https://api\\.w\\.org/\"",
			},
			"js": map[string]interface{}{
				"wp_username": "",
			},
			"meta": map[string]interface{}{
				"generator": "^WordPress ?([\\d.]+)?\\;version:\\1",
			},
			"script":  "/wp-(?:content|includes)/",
			"website": "https://wordpress.org",
		},
		"WordPress Super Cache": map[string]interface{}{
			"cats": []interface{}{
				23,
			},
			"headers": map[string]interface{}{
				"WP-Super-Cache": "",
			},
			"html":    "<!--[^>]+WP-Super-Cache",
			"icon":    "wp_super_cache.png",
			"implies": "WordPress",
			"website": "http://z9.io/wp-super-cache/",
		},
		"Wowza Media Server": map[string]interface{}{
			"cats": []interface{}{
				38,
			},
			"html":    "<title>Wowza Media Server \\d+ ((?:\\w+ Edition )?\\d+\\.[\\d\\.]+(?: build\\d+)?)?\\;version:\\1",
			"icon":    "Wowza Media Server.png",
			"website": "http://www.wowza.com",
		},
		"X-Cart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"xid": "[a-z\\d]{32}(?:;|$)",
			},
			"html": []interface{}{
				"Powered by X-Cart(?: (\\d+))? <a[^>]+href=\"http://www\\.x-cart\\.com/\"[^>]*>\\;version:\\1",
				"<a[^>]+href=\"[^\"]*(?:\\?|&)xcart_form_id=[a-z\\d]{32}(?:&|$)",
			},
			"icon":    "X-Cart.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"xcart_web_dir": "",
				"xliteConfig":   "",
			},
			"meta": map[string]interface{}{
				"generator": "X-Cart(?: (\\d+))?\\;version:\\1",
			},
			"script":  "/skin/common_files/modules/Product_Options/func\\.js",
			"website": "http://x-cart.com",
		},
		"XAMPP": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"html": "<title>XAMPP(?: Version ([\\d\\.]+))?</title>\\;version:\\1",
			"icon": "XAMPP.png",
			"implies": []interface{}{
				"Apache",
				"MySQL",
				"PHP",
				"Perl",
			},
			"meta": map[string]interface{}{
				"author": "Kai Oswald Seidler\\;confidence:10",
			},
			"website": "http://www.apachefriends.org/en/xampp.html",
		},
		"XMB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html":    "<!-- Powered by XMB",
			"icon":    "XMB.png",
			"website": "http://www.xmbforum.com",
		},
		"XOOPS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "XOOPS.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"xoops": "",
			},
			"meta": map[string]interface{}{
				"generator": "XOOPS",
			},
			"website": "http://xoops.org",
		},
		"XRegExp": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "XRegExp.png",
			"js": map[string]interface{}{
				"XRegExp.version": "^(.+)$\\;version:\\1",
			},
			"script": []interface{}{
				"xregexp[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"/([\\d.]+)/xregexp(?:\\.min)?\\.js\\;version:\\1",
				"xregexp.*\\.js",
			},
			"website": "http://xregexp.com",
		},
		"XWiki": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"excludes": "MediaWiki",
			"html": []interface{}{
				"<html[^>]data-xwiki-[^>]>",
			},
			"icon":    "xwiki.png",
			"implies": "Java\\;confidence:99",
			"meta": map[string]interface{}{
				"wiki": "xwiki",
			},
			"website": "http://www.xwiki.org",
		},
		"Xajax": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "Xajax.png",
			"script":  "xajax_core.*\\.js",
			"website": "http://xajax-project.org",
		},
		"Xanario": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Xanario.png",
			"meta": map[string]interface{}{
				"generator": "xanario shopsoftware",
			},
			"website": "http://xanario.de",
		},
		"XenForo": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"cookies": map[string]interface{}{
				"xf_csrf":    "",
				"xf_session": "",
			},
			"html": []interface{}{
				"(?:jQuery\\.extend\\(true, XenForo|Forum software by XenForo™|<!--XF:branding|<html[^>]+id=\"XenForo\")",
				"<html id=\"XF\" ",
			},
			"icon": "XenForo.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"js": map[string]interface{}{
				"XF.GuestUsername": "",
			},
			"website": "http://xenforo.com",
		},
		"Xeora": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
				27,
			},
			"headers": map[string]interface{}{
				"Server":       "XeoraEngine",
				"X-Powered-By": "XeoraCube",
			},
			"html":    "<input type=\"hidden\" name=\"_sys_bind_\\d+\" id=\"_sys_bind_\\d+\" />",
			"icon":    "xeora.png",
			"script":  "/_bi_sps_v.+\\.js",
			"website": "http://www.xeora.org",
		},
		"Xitami": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Xitami(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Xitami.png",
			"website": "http://xitami.com",
		},
		"Xonic": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": []interface{}{
				"Powered by <a href=\"http://www\\.xonic-solutions\\.de/index\\.php\" target=\"_blank\">xonic-solutions Shopsoftware</a>",
			},
			"icon": "xonic.png",
			"meta": map[string]interface{}{
				"keywords": "xonic-solutions",
			},
			"script":  "core/jslib/jquery\\.xonic\\.js\\.php",
			"website": "http://www.xonic-solutions.de",
		},
		"XpressEngine": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "XpressEngine.png",
			"meta": map[string]interface{}{
				"generator": "XpressEngine",
			},
			"website": "http://www.xpressengine.com/",
		},
		"YUI": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "YUI.png",
			"js": map[string]interface{}{
				"YAHOO.VERSION": "^(.+)$\\;version:\\1",
				"YUI.version":   "^(.+)$\\;version:\\1",
			},
			"script":  "(?:/yui/|yui\\.yahooapis\\.com)",
			"website": "http://yuilibrary.com",
		},
		"YUI Doc": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"html":    "(?:<html[^>]* yuilibrary\\.com/rdf/[\\d.]+/yui\\.rdf|<body[^>]+class=\"yui3-skin-sam)",
			"icon":    "yahoo.png",
			"website": "http://developer.yahoo.com/yui/yuidoc",
		},
		"YaBB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html":    "Powered by <a href=\"[^>]+yabbforum",
			"icon":    "YaBB.png",
			"website": "http://www.yabbforum.com",
		},
		"Yahoo Advertising": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": []interface{}{
				"<iframe[^>]+adserver\\.yahoo\\.com",
				"<img[^>]+clicks\\.beap\\.bc\\.yahoo\\.com",
			},
			"icon": "yahoo.png",
			"js": map[string]interface{}{
				"adxinserthtml": "",
			},
			"script":  "adinterax\\.com",
			"website": "http://advertising.yahoo.com",
		},
		"Yahoo! Ecommerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-XRDS-Location": "/ystore/",
			},
			"html": "<link[^>]+store\\.yahoo\\.net",
			"icon": "yahoo.png",
			"js": map[string]interface{}{
				"YStore": "",
			},
			"website": "http://smallbusiness.yahoo.com/ecommerce",
		},
		"Yahoo! Tag Manager": map[string]interface{}{
			"cats": []interface{}{
				42,
			},
			"html":    "<!-- (?:End )?Yahoo! Tag Manager -->",
			"script":  "b\\.yjtag\\.jp/iframe",
			"icon":    "yahoo.png",
			"website": "https://tagmanager.yahoo.co.jp/",
		},
		"Yahoo! Web Analytics": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "yahoo.png",
			"js": map[string]interface{}{
				"YWA": "",
			},
			"script":  "d\\.yimg\\.com/mi/ywa\\.js",
			"website": "http://web.analytics.yahoo.com",
		},
		"Yandex.Direct": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": "<yatag class=\"ya-partner__ads\">",
			"icon": "Yandex.Direct.png",
			"js": map[string]interface{}{
				"yandex_ad_format":  "",
				"yandex_partner_id": "",
			},
			"script":  "https?://an\\.yandex\\.ru/",
			"website": "http://partner.yandex.com",
		},
		"Yandex.Metrika": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "Yandex.Metrika.png",
			"js": map[string]interface{}{
				"yandex_metrika": "",
			},
			"script": []interface{}{
				"mc\\.yandex\\.ru\\/metrika\\/watch\\.js",
				"cdn\\.jsdelivr\\.net\\/npm\\/yandex-metrica-watch\\/watch\\.js",
			},
			"website": "http://metrika.yandex.com",
		},
		"Yaws": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "Yaws(?: ([\\d.]+))?\\;version:\\1",
			},
			"icon": "Yaws.png",
			"implies": []interface{}{
				"Erlang",
			},
			"website": "http://yaws.hyber.org",
		},
		"Yieldlab": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"icon":    "Yieldlab.png",
			"script":  "^https?://(?:[^/]+\\.)?yieldlab\\.net/",
			"website": "http://yieldlab.de",
		},
		"Yii": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"cookies": map[string]interface{}{
				"YII_CSRF_TOKEN": "",
			},
			"html": []interface{}{
				"Powered by <a href=\"http://www\\.yiiframework\\.com/\" rel=\"external\">Yii Framework</a>",
				"<input type=\"hidden\" value=\"[a-zA-Z0-9]{40}\" name=\"YII_CSRF_TOKEN\" \\/>",
				"<!\\[CDATA\\[YII-BLOCK-(?:HEAD|BODY-BEGIN|BODY-END)\\]",
			},
			"icon": "Yii.png",
			"implies": []interface{}{
				"PHP",
			},
			"script": []interface{}{
				"/assets/[a-zA-Z0-9]{8}\\/yii\\.js$",
				"/yii\\.(?:validation|activeForm)\\.js",
			},
			"website": "https://www.yiiframework.com",
		},
		"Yoast SEO": map[string]interface{}{
			"cats": []interface{}{
				54,
			},
			"html": []interface{}{
				"<!-- This site is optimized with the Yoast (?:WordPress )?SEO plugin v([\\d.]+) -\\;version:\\1",
			},
			"icon":    "Yoast SEO.png",
			"implies": "WordPress",
			"website": "http://yoast.com",
		},
		"WP-Statistics": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": []interface{}{
				"<!-- Analytics by WP-Statistics v([\\d.]+) -\\;version:\\1",
			},
			"icon":    "WP-Statistics.png",
			"implies": "WordPress",
			"website": "https://wp-statistics.com",
		},
		"YouTrack": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"html": []interface{}{
				"no-title=\"YouTrack\">",
				"data-reactid=\"[^\"]+\">youTrack ([0-9.]+)<\\;version:\\1",
				"type=\"application/opensearchdescription\\+xml\" title=\"YouTrack\"/>",
			},
			"icon":    "YouTrack.png",
			"website": "http://www.jetbrains.com/youtrack/",
		},
		"YouTube": map[string]interface{}{
			"cats": []interface{}{
				14,
			},
			"html":    "<(?:param|embed|iframe)[^>]+youtube(?:-nocookie)?\\.com/(?:v|embed)",
			"icon":    "YouTube.png",
			"website": "http://www.youtube.com",
		},
		"iEXExchanger": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"implies": []interface{}{
				"PHP",
				"Apache",
				"Angular",
			},
			"cookies": map[string]interface{}{
				"iexexchanger_session": "",
			},
			"meta": map[string]interface{}{
				"generator": "iEXExchanger",
			},
			"icon":    "iEXExchanger.png",
			"website": "https://exchanger.iexbase.com",
		},
		"ZK": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html":    "<!-- ZK [.\\d\\s]+-->",
			"icon":    "ZK.png",
			"implies": "Java",
			"script":  "zkau/",
			"website": "http://zkoss.org",
		},
		"ZURB Foundation": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": []interface{}{
				"<link[^>]+foundation[^>\"]+css",
				"<div [^>]*class=\"[^\"]*(?:small|medium|large)-\\d{1,2} columns",
			},
			"icon": "ZURB Foundation.png",
			"js": map[string]interface{}{
				"Foundation.version": "([\\d.]+)\\;version:\\1",
			},
			"website": "http://foundation.zurb.com",
		},
		"Zabbix": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html":    "<body[^>]+zbxCallPostScripts",
			"icon":    "Zabbix.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"zbxCallPostScripts": "",
			},
			"meta": map[string]interface{}{
				"Author": "ZABBIX SIA\\;confidence:70",
			},
			"url":     "\\/zabbix\\/\\;confidence:30",
			"website": "http://zabbix.com",
		},
		"Zanox": map[string]interface{}{
			"cats": []interface{}{
				36,
			},
			"html": "<img [^>]*src=\"[^\"]+ad\\.zanox\\.com",
			"icon": "Zanox.png",
			"js": map[string]interface{}{
				"zanox": "",
			},
			"script":  "zanox\\.com/scripts/zanox\\.js$",
			"website": "http://zanox.com",
		},
		"Zen Cart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "Zen Cart.png",
			"meta": map[string]interface{}{
				"generator": "Zen Cart",
			},
			"website": "http://www.zen-cart.com",
		},
		"Zend": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"cookies": map[string]interface{}{
				"ZENDSERVERSESSID": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "Zend(?:Server)?(?:[\\s/]?([0-9.]+))?\\;version:\\1",
			},
			"icon":    "Zend.png",
			"website": "http://zend.com",
		},
		"Zendesk Chat": map[string]interface{}{
			"cats": []interface{}{
				52,
			},
			"icon":    "Zendesk Chat.png",
			"script":  "v2\\.zopim\\.com",
			"website": "http://zopim.com",
		},
		"Zenfolio": map[string]interface{}{
			"cats": []interface{}{
				7,
			},
			"icon": "Zenfolio.png",
			"js": map[string]interface{}{
				"Zenfolio": "",
			},
			"website": "https://zenfolio.com",
		},
		"Zepto": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "Zepto.png",
			"js": map[string]interface{}{
				"Zepto": "",
			},
			"script":  "zepto.*\\.js",
			"website": "http://zeptojs.com",
		},
		"Zeuscart": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<form name=\"product\" method=\"post\" action=\"[^\"]+\\?do=addtocart&prodid=\\d+\"(?!<\\/form>.)+<input type=\"hidden\" name=\"addtocart\" value=\"\\d+\">",
			"icon":    "Zeuscart.png",
			"implies": "PHP",
			"url":     "\\?do=prodetail&action=show&prodid=\\d+",
			"website": "http://zeuscart.com",
		},
		"Zinnia": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"icon":    "Zinnia.png",
			"implies": "Django",
			"meta": map[string]interface{}{
				"generator": "Zinnia",
			},
			"website": "http://django-blog-zinnia.com",
		},
		"Zone.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"js": map[string]interface{}{
				"Zone.root": "",
			},
			"website": "https://github.com/angular/zone.js/",
		},
		"Zope": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "^Zope/",
			},
			"icon":    "Zope.png",
			"website": "http://zope.org",
		},
		"a-blog cms": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "a-blog cms.svg",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "a-blog cms",
			},
			"website": "http://www.a-blogcms.jp",
		},
		"actionhero.js": map[string]interface{}{
			"cats": []interface{}{
				1,
				18,
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "actionhero API",
			},
			"icon":    "actionhero.js.png",
			"implies": "Node.js",
			"js": map[string]interface{}{
				"actionheroClient": "",
			},
			"script":  "actionheroClient\\.js",
			"website": "http://www.actionherojs.com",
		},
		"amCharts": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "amCharts.png",
			"js": map[string]interface{}{
				"AmCharts": "",
			},
			"script":  "amcharts.*\\.js",
			"website": "http://amcharts.com",
		},
		"animate.css": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"html": []interface{}{
				"<link [^>]+(?:/([\\d.]+)/)?animate\\.(?:min\\.)?css\\;version:\\1",
			},
			"website": "https://daneden.github.io/animate.css/",
		},
		"basket.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "basket.js.png",
			"js": map[string]interface{}{
				"basket.isValidItem": "",
			},
			"script":  "basket.*\\.js\\;confidence:10",
			"website": "https://addyosmani.github.io/basket.js/",
		},
		"cPanel": map[string]interface{}{
			"cats": []interface{}{
				9,
			},
			"headers": map[string]interface{}{
				"Server": "cpsrvd/([\\d.]+)\\;version:\\1",
			},
			"cookies": map[string]interface{}{
				"cpsession": "",
				"cprelogin": "",
			},
			"html":    "<!-- cPanel",
			"icon":    "cPanel.png",
			"website": "http://www.cpanel.net",
		},
		"cgit": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html": []interface{}{
				"<[^>]+id='cgit'",
				"generated by <a href='http://git\\.zx2c4\\.com/cgit/about/'>cgit v([\\d.a-z-]+)</a>\\;version:\\1",
			},
			"icon":    "cgit.png",
			"implies": "git",
			"meta": map[string]interface{}{
				"generator": "^cgit v([\\d.a-z-]+)$\\;version:\\1",
			},
			"website": "http://git.zx2c4.com/cgit",
		},
		"comScore": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"html": "<iframe[^>]* (?:id=\"comscore\"|scr=[^>]+comscore)|\\.scorecardresearch\\.com/beacon\\.js|COMSCORE\\.beacon",
			"icon": "comScore.png",
			"js": map[string]interface{}{
				"COMSCORE":  "",
				"_COMSCORE": "",
			},
			"script":  "\\.scorecardresearch\\.com/beacon\\.js|COMSCORE\\.beacon",
			"website": "http://comscore.com",
		},
		"debut": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "debut\\/?([\\d\\.]+)?\\;version:\\1",
			},
			"icon":    "debut.png",
			"implies": "Brother",
			"website": "http://www.brother.com",
		},
		"deepMiner": map[string]interface{}{
			"cats": []interface{}{
				56,
			},
			"icon": "deepminer.png",
			"js": map[string]interface{}{
				"deepMiner": "",
			},
			"script":  "deepMiner\\.js",
			"website": "https://github.com/deepwn/deepMiner",
		},
		"e107": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"e107_tz": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "e107",
			},
			"icon":    "e107.png",
			"implies": "PHP",
			"script":  "[^a-z\\d]e107\\.js",
			"website": "http://e107.org",
		},
		"eSyndiCat": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Drectory-Script": "^eSyndiCat",
			},
			"icon":    "eSyndiCat.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"esyndicat": "",
			},
			"meta": map[string]interface{}{
				"generator": "^eSyndiCat ",
			},
			"website": "http://esyndicat.com",
		},
		"eZ Publish": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"cookies": map[string]interface{}{
				"eZSESSID": "",
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^eZ Publish",
			},
			"icon":    "eZ Publish.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "eZ Publish",
			},
			"website": "http://ez.no",
		},
		"ef.js": map[string]interface{}{
			"cats": []interface{}{
				12,
			},
			"icon": "ef.js.svg",
			"js": map[string]interface{}{
				"ef.version": "^(.+)$\\;version:\\1",
				"efCore":     "",
			},
			"script":  "/ef(?:-core)?(?:\\.min|\\.dev)?\\.js",
			"website": "http://ef.js.org",
		},
		"enduro.js": map[string]interface{}{
			"cats": []interface{}{
				1,
				18,
				47,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^enduro\\.js$",
			},
			"icon":    "enduro.js.svg",
			"implies": "Node.js",
			"website": "http://endurojs.com",
		},
		"git": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"icon": "git.svg",
			"meta": map[string]interface{}{
				"generator": "\\bgit/([\\d.]+\\d)\\;version:\\1",
			},
			"website": "http://git-scm.com",
		},
		"gitlist": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"html": "<p>Powered by <a[^>]+>GitList ([\\d.]+)\\;version:\\1",
			"implies": []interface{}{
				"PHP",
				"git",
			},
			"website": "http://gitlist.org",
		},
		"gitweb": map[string]interface{}{
			"cats": []interface{}{
				47,
			},
			"html": "<!-- git web interface version ([\\d.]+)?\\;version:\\1",
			"icon": "git.svg",
			"implies": []interface{}{
				"Perl",
				"git",
			},
			"meta": map[string]interface{}{
				"generator": "gitweb(?:/([\\d.]+\\d))?\\;version:\\1",
			},
			"script":  "static/gitweb\\.js$",
			"website": "http://git-scm.com",
		},
		"govCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "govCMS.svg",
			"implies": []interface{}{
				"Drupal",
			},
			"meta": map[string]interface{}{
				"generator": "Drupal ([\\d]+) \\(http:\\/\\/drupal\\.org\\) \\+ govCMS\\;version:\\1",
			},
			"website": "https://www.govcms.gov.au",
		},
		"gunicorn": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "gunicorn(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "gunicorn.png",
			"implies": "Python",
			"website": "http://gunicorn.org",
		},
		"hapi.js": map[string]interface{}{
			"cats": []interface{}{
				18,
				22,
			},
			"cookies": map[string]interface{}{
				"Fe26.2**": "\\;confidence:50",
			},
			"icon":    "hapi.js.png",
			"implies": "Node.js",
			"website": "http://hapijs.com",
		},
		"iCongo": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon":    "Hybris.png",
			"implies": "Adobe ColdFusion",
			"meta": map[string]interface{}{
				"iCongo": "",
			},
			"website": "http://hybris.com/icongo",
		},
		"iPresta": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "iPresta.png",
			"implies": []interface{}{
				"PHP",
				"PrestaShop",
			},
			"meta": map[string]interface{}{
				"designer": "iPresta",
			},
			"website": "http://ipresta.ir",
		},
		"iWeb": map[string]interface{}{
			"cats": []interface{}{
				20,
			},
			"icon": "iWeb.png",
			"meta": map[string]interface{}{
				"generator": "^iWeb( [\\d.]+)?\\;version:\\1",
			},
			"website": "http://apple.com/ilife/iweb",
		},
		"ikiwiki": map[string]interface{}{
			"cats": []interface{}{
				8,
			},
			"html": []interface{}{
				"<link rel=\"alternate\" type=\"application/x-wiki\" title=\"Edit this page\" href=\"[^\"]*/ikiwiki\\.cgi",
				"<a href=\"/(?:cgi-bin/)?ikiwiki\\.cgi\\?do=",
			},
			"icon":    "ikiwiki.png",
			"website": "http://ikiwiki.info",
		},
		"imperia CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<imp:live-info sysid=\"[0-9a-f-]+\"(?: node_id=\"[0-9/]*\")? *\\/>",
			"icon":    "imperiaCMS.svg",
			"implies": "Perl",
			"meta": map[string]interface{}{
				"GENERATOR":           "^IMPERIA ([0-9.]{2,})+$\\;version:\\1",
				"X-Imperia-Live-Info": "",
			},
			"url":     "imperia/md/",
			"website": "https://www.pirobase-imperia.com/de/produkte/produktuebersicht/imperia-cms",
		},
		"io4 CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "io4 CMS.png",
			"meta": map[string]interface{}{
				"generator": "GO[ |]+CMS Enterprise",
			},
			"website": "http://notenbomer.nl/Producten/Content_management/io4_|_cms",
		},
		"ip-label": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"icon": "iplabel.svg",
			"js": map[string]interface{}{
				"clobs": "",
			},
			"script":  "clobs\\.js",
			"website": "http://www.ip-label.com",
		},
		"jQTouch": map[string]interface{}{
			"cats": []interface{}{
				26,
			},
			"icon": "jQTouch.png",
			"js": map[string]interface{}{
				"jQT": "",
			},
			"script":  "jqtouch.*\\.js",
			"website": "http://jqtouch.com",
		},
		"jQuery": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "jQuery.svg",
			"js": map[string]interface{}{
				"jQuery.fn.jquery": "([\\d.]+)\\;version:\\1",
			},
			"script": []interface{}{
				"jquery[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"/([\\d.]+)/jquery(?:\\.min)?\\.js\\;version:\\1",
				"jquery.*\\.js(?:\\?ver(?:sion)?=([\\d.]+))?\\;version:\\1",
			},
			"website": "https://jquery.com",
		},
		"jQuery Migrate": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "jQuery.svg",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"jQuery.migrateVersion":  "([\\d.]+)\\;version:\\1",
				"jQuery.migrateWarnings": "",
				"jqueryMigrate":          "",
			},
			"script":  "jquery[.-]migrate(?:-([\\d.]+))?(?:\\.min)?\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1?\\1:\\2",
			"website": "https://github.com/jquery/jquery-migrate",
		},
		"jQuery Mobile": map[string]interface{}{
			"cats": []interface{}{
				26,
			},
			"icon":    "jQuery Mobile.svg",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"jQuery.mobile.version": "^(.+)$\\;version:\\1",
			},
			"script":  "jquery[.-]mobile(?:-([\\d.]))?(?:\\.min)?\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1?\\1:\\2",
			"website": "https://jquerymobile.com",
		},
		"jQuery-pjax": map[string]interface{}{
			"cats": []interface{}{
				26,
			},
			"implies": "jQuery",
			"meta": map[string]interface{}{
				"pjax-timeout": "",
				"pjax-replace": "",
				"pjax-push":    "",
			},
			"js": map[string]interface{}{
				"jQuery.pjax": "",
			},
			"script":  "jquery[.-]pjax(?:-([\\d.]))?(?:\\.min)?\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1?\\1:\\2",
			"html":    "<div[^>]+data-pjax-container",
			"website": "https://github.com/defunkt/jquery-pjax",
		},
		"jQuery Sparklines": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"implies": "jQuery",
			"script":  "jquery\\.sparkline.*\\.js",
			"website": "http://omnipotent.net/jquery.sparkline/",
		},
		"jQuery UI": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon":    "jQuery UI.svg",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"jQuery.ui.version": "^(.+)$\\;version:\\1",
			},
			"script": []interface{}{
				"jquery-ui[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
				"([\\d.]+)/jquery-ui(?:\\.min)?\\.js\\;version:\\1",
				"jquery-ui.*\\.js",
			},
			"website": "http://jqueryui.com",
		},
		"jqPlot": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon":    "jqPlot.png",
			"implies": "jQuery",
			"script":  "jqplot.*\\.js",
			"website": "http://www.jqplot.com",
		},
		"libwww-perl-daemon": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "libwww-perl-daemon(?:/([\\d\\.]+))?\\;version:\\1",
			},
			"icon":    "libwww-perl-daemon.png",
			"implies": "Perl",
			"website": "http://metacpan.org/pod/HTTP::Daemon",
		},
		"lighttpd": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "lighttpd(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "lighttpd.png",
			"website": "http://www.lighttpd.net",
		},
		"math.js": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "math.js.png",
			"js": map[string]interface{}{
				"mathjs": "",
			},
			"script":  "math(?:\\.min)?\\.js",
			"website": "http://mathjs.org",
		},
		"mini_httpd": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "mini_httpd(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "mini_httpd.png",
			"website": "http://acme.com/software/mini_httpd",
		},
		"mod_auth_pam": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server": "mod_auth_pam(?:/([\\d\\.]+))?\\;version:\\1",
			},
			"icon":    "Apache.svg",
			"implies": "Apache",
			"website": "http://pam.sourceforge.net/mod_auth_pam",
		},
		"mod_dav": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server": "\\b(?:mod_)?DAV\\b(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Apache.svg",
			"implies": "Apache",
			"website": "http://webdav.org/mod_dav",
		},
		"mod_fastcgi": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server": "mod_fastcgi(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "Apache.svg",
			"implies": "Apache",
			"website": "http://www.fastcgi.com/mod_fastcgi/docs/mod_fastcgi.html",
		},
		"mod_jk": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server": "mod_jk(?:/([\\d\\.]+))?\\;version:\\1",
			},
			"icon": "Apache.svg",
			"implies": []interface{}{
				"Apache Tomcat",
				"Apache",
			},
			"website": "http://tomcat.apache.org/tomcat-3.3-doc/mod_jk-howto.html",
		},
		"mod_perl": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server": "mod_perl(?:/([\\d\\.]+))?\\;version:\\1",
			},
			"icon": "mod_perl.png",
			"implies": []interface{}{
				"Perl",
				"Apache",
			},
			"website": "http://perl.apache.org",
		},
		"mod_python": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server": "mod_python(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon": "mod_python.png",
			"implies": []interface{}{
				"Python",
				"Apache",
			},
			"website": "http://www.modpython.org",
		},
		"mod_rack": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server":       "mod_rack(?:/([\\d.]+))?\\;version:\\1",
				"X-Powered-By": "mod_rack(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon": "Phusion Passenger.png",
			"implies": []interface{}{
				"Ruby on Rails\\;confidence:50",
				"Apache",
			},
			"website": "http://phusionpassenger.com",
		},
		"mod_rails": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server":       "mod_rails(?:/([\\d.]+))?\\;version:\\1",
				"X-Powered-By": "mod_rails(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon": "Phusion Passenger.png",
			"implies": []interface{}{
				"Ruby on Rails\\;confidence:50",
				"Apache",
			},
			"website": "http://phusionpassenger.com",
		},
		"mod_ssl": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server": "mod_ssl(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "mod_ssl.png",
			"implies": "Apache",
			"website": "http://modssl.org",
		},
		"mod_wsgi": map[string]interface{}{
			"cats": []interface{}{
				33,
			},
			"headers": map[string]interface{}{
				"Server":       "mod_wsgi(?:/([\\d.]+))?\\;version:\\1",
				"X-Powered-By": "mod_wsgi(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon": "mod_wsgi.png",
			"implies": []interface{}{
				"Python\\;confidence:50",
				"Apache",
			},
			"website": "https://code.google.com/p/modwsgi",
		},
		"nopCommerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"Nop.customer": "",
			},
			"html":    "(?:<!--Powered by nopCommerce|Powered by: <a[^>]+nopcommerce)",
			"icon":    "nopCommerce.png",
			"implies": "Microsoft ASP.NET",
			"meta": map[string]interface{}{
				"generator": "^nopCommerce$",
			},
			"website": "http://www.nopcommerce.com",
		},
		"openEngine": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "openEngine.png",
			"meta": map[string]interface{}{
				"openEngine": "",
			},
			"website": "http://openengine.de/html/pages/de/",
		},
		"osCSS": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html":    "<body onload=\"window\\.defaultStatus='oscss templates';\"",
			"icon":    "osCSS.png",
			"website": "http://www.oscss.org",
		},
		"osCommerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"cookies": map[string]interface{}{
				"osCsid": "",
			},
			"html": []interface{}{
				"<br />Powered by <a href=\"https?://www\\.oscommerce\\.com",
				"<(?:input|a)[^>]+name=\"osCsid\"",
				"<(?:tr|td|table)class=\"[^\"]*infoBoxHeading",
			},
			"icon": "osCommerce.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"website": "https://www.oscommerce.com",
		},
		"osTicket": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"cookies": map[string]interface{}{
				"OSTSESSID": "",
			},
			"icon": "osTicket.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"website": "http://osticket.com",
		},
		"otrs": map[string]interface{}{
			"cats": []interface{}{
				13,
			},
			"html":    "<!--\\s+OTRS: Copyright \\d+-\\d+, OTRS AG",
			"icon":    "otrs.png",
			"implies": "Perl",
			"script":  "^/otrs-web/js/",
			"website": "https://www.otrs.com",
		},
		"ownCloud": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"html":    "<a href=\"https://owncloud\\.com\" target=\"_blank\">ownCloud Inc\\.</a><br/>Your Cloud, Your Data, Your Way!",
			"icon":    "ownCloud.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"apple-itunes-app": "app-id=543672169",
			},
			"website": "https://owncloud.org",
		},
		"papaya CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html":    "<link[^>]*/papaya-themes/",
			"icon":    "papaya CMS.png",
			"implies": "PHP",
			"website": "https://papaya-cms.com",
		},
		"particles.js": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"html": "<div id=\"particles-js\">",
			"js": map[string]interface{}{
				"particlesJS": "",
			},
			"script":  "/particles(?:\\.min)?\\.js",
			"website": "https://vincentgarreau.com/particles.js/",
		},
		"PhotoShelter": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": []interface{}{
				"<!--\\s+#+ Powered by the PhotoShelter Beam platform",
				"<link rel=[\"']dns-prefetch[\"'] [^>]+photoshelter.com",
			},
			"implies": []interface{}{
				"PHP",
				"MySQL",
				"jQuery",
			},
			"icon":    "PhotoShelter.png",
			"website": "https://www.photoshelter.com",
		},
		"phpAlbum": map[string]interface{}{
			"cats": []interface{}{
				7,
			},
			"html":    "<!--phpalbum ([.\\d\\s]+)-->\\;version:\\1",
			"icon":    "phpAlbum.png",
			"implies": "PHP",
			"website": "http://phpalbum.net",
		},
		"phpBB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"cookies": map[string]interface{}{
				"phpbb": "",
			},
			"html": []interface{}{
				"Powered by <a[^>]+phpBB",
				"<div class=phpbb_copyright>",
				"<[^>]+styles/(?:sub|pro)silver/theme",
				"<img[^>]+i_icon_mini",
				"<table class=\"[^\"]*forumline",
			},
			"icon":    "phpBB.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"phpbb":                 "",
				"style_cookie_settings": "",
			},
			"meta": map[string]interface{}{
				"copyright": "phpBB Group",
			},
			"website": "https://phpbb.com",
		},
		"phpCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon":    "phpCMS.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"phpcms": "",
			},
			"website": "http://phpcms.de",
		},
		"phpDocumentor": map[string]interface{}{
			"cats": []interface{}{
				4,
			},
			"html":    "<!-- Generated by phpDocumentor",
			"icon":    "phpDocumentor.png",
			"implies": "PHP",
			"website": "https://www.phpdoc.org",
		},
		"phpMyAdmin": map[string]interface{}{
			"cats": []interface{}{
				3,
			},
			"html": "(?: \\| phpMyAdmin ([\\d.]+)<\\/title>|PMA_sendHeaderLocation\\(|<link [^>]*href=\"[^\"]*phpmyadmin\\.css\\.php)\\;version:\\1",
			"icon": "phpMyAdmin.png",
			"implies": []interface{}{
				"PHP",
				"MySQL",
			},
			"js": map[string]interface{}{
				"pma_absolute_uri": "",
			},
			"website": "https://www.phpmyadmin.net",
		},
		"phpPgAdmin": map[string]interface{}{
			"cats": []interface{}{
				3,
			},
			"html":    "(?:<title>phpPgAdmin</title>|<span class=\"appname\">phpPgAdmin)",
			"icon":    "phpPgAdmin.png",
			"implies": "PHP",
			"website": "http://phppgadmin.sourceforge.net",
		},
		"phpSQLiteCMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "phpSQLiteCMS.png",
			"implies": []interface{}{
				"PHP",
				"SQLite",
			},
			"meta": map[string]interface{}{
				"generator": "^phpSQLiteCMS(?: (.+))?$\\;version:\\1",
			},
			"website": "http://phpsqlitecms.net",
		},
		"phpliteadmin": map[string]interface{}{
			"cats": []interface{}{
				3,
			},
			"html": []interface{}{
				"<span id='logo'>phpLiteAdmin</span> <span id='version'>v([0-9.]+)<\\;version:\\1",
				"<!-- Copyright [0-9]+ phpLiteAdmin (?:https?://www\\.phpliteadmin\\.org/) -->",
				"Powered by <a href='https?://www\\.phpliteadmin\\.org/'",
			},
			"icon": "phpliteadmin.png",
			"implies": []interface{}{
				"PHP",
				"SQLite",
			},
			"website": "https://www.phpliteadmin.org/",
		},
		"phpwind": map[string]interface{}{
			"cats": []interface{}{
				1,
				2,
			},
			"html":    "(?:Powered|Code) by <a href=\"[^\"]+phpwind\\.net",
			"icon":    "phpwind.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "^phpwind(?: v([0-9-]+))?\\;version:\\1",
			},
			"website": "https://www.phpwind.net",
		},
		"pirobase CMS": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"html": []interface{}{
				"<(?:script|link)[^>]/site/[a-z0-9/._-]+/resourceCached/[a-z0-9/._-]+",
				"<input[^>]+cbi:///cms/",
			},
			"icon":    "pirobaseCMS.svg",
			"implies": "Java",
			"website": "https://www.pirobase-imperia.com/de/produkte/produktuebersicht/pirobase-cms",
		},
		"prettyPhoto": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"html":    "(?:<link [^>]*href=\"[^\"]*prettyPhoto(?:\\.min)?\\.css|<a [^>]*rel=\"prettyPhoto)",
			"icon":    "prettyPhoto.png",
			"implies": "jQuery",
			"js": map[string]interface{}{
				"pp_alreadyInitialized": "",
				"pp_descriptions":       "",
				"pp_images":             "",
				"pp_titles":             "",
			},
			"script":  "jquery\\.prettyPhoto\\.js",
			"website": "http://no-margin-for-errors.com/projects/prettyphoto-jquery-lightbox-clone/",
		},
		"punBB": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"js": map[string]interface{}{
				"PUNBB": "",
			},
			"html":    "Powered by <a href=\"[^>]+punbb",
			"icon":    "punBB.png",
			"implies": "PHP",
			"website": "http://punbb.informer.com",
		},
		"reCAPTCHA": map[string]interface{}{
			"cats": []interface{}{
				16,
			},
			"html": []interface{}{
				"<div[^>]+id=\"recaptcha_image",
				"<link[^>]+recaptcha",
				"<div[^>]+class=\"g-recaptcha\"",
			},
			"icon": "reCAPTCHA.png",
			"js": map[string]interface{}{
				"Recaptcha": "",
				"recaptcha": "",
			},
			"script": []interface{}{
				"api-secure\\.recaptcha\\.net",
				"recaptcha_ajax\\.js",
				"/recaptcha/api\\.js",
			},
			"website": "https://www.google.com/recaptcha/",
		},
		"sIFR": map[string]interface{}{
			"cats": []interface{}{
				17,
			},
			"icon":    "sIFR.png",
			"script":  "sifr\\.js",
			"website": "https://www.mikeindustries.com/blog/sifr",
		},
		"sNews": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "sNews.png",
			"meta": map[string]interface{}{
				"generator": "sNews",
			},
			"website": "https://snewscms.com",
		},
		"script.aculo.us": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "script.aculo.us.png",
			"js": map[string]interface{}{
				"Scriptaculous.Version": "^(.+)$\\;version:\\1",
			},
			"script":  "/(?:scriptaculous|protoaculous)(?:\\.js|/)",
			"website": "https://script.aculo.us",
		},
		"scrollreveal": map[string]interface{}{
			"cats": []interface{}{
				59,
			},
			"icon": "scrollreveal.svg",
			"html": "<[^>]+data-sr(?:-id)",
			"js": map[string]interface{}{
				"ScrollReveal().version": "^(.+)$\\;version:\\1",
			},
			"script":  "scrollreveal(?:\\.min)(?:\\.js)",
			"website": "https://scrollrevealjs.org",
		},
		"shine.js": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"js": map[string]interface{}{
				"Shine": "",
			},
			"script":  "shine(?:\\.min)?\\.js",
			"website": "https://bigspaceship.github.io/shine.js/",
		},
		"styled-components": map[string]interface{}{
			"cats": []interface{}{
				12,
				47,
			},
			"html": []interface{}{
				"<style[^>]*data-styled(?:-components)?[\\s\"]",
				"<style[^>]+data-styled-version=\"([0-9]+)\"\\;version:\\1",
				"<[^>]+sc-component-id: sc-",
			},
			"icon": "styled-components.png",
			"implies": []interface{}{
				"React",
			},
			"js": map[string]interface{}{
				"styled": "",
			},
			"website": "https://styled-components.com",
		},
		"swift.engine": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "swift\\.engine",
			},
			"icon":    "swift.engine.png",
			"website": "http://mittec.ru/default",
		},
		"three.js": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"icon": "three.js.png",
			"js": map[string]interface{}{
				"THREE.REVISION": "^(.+)$\\;version:\\1",
			},
			"script":  "three(?:\\.min)?\\.js",
			"website": "https://threejs.org",
		},
		"thttpd": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"Server": "\\bthttpd(?:/([\\d.]+))?\\;version:\\1",
			},
			"icon":    "thttpd.png",
			"website": "https://acme.com/software/thttpd",
		},
		"total.js": map[string]interface{}{
			"cats": []interface{}{
				18,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^total\\.js",
			},
			"icon":    "total.js.png",
			"implies": "Node.js",
			"website": "https://totaljs.com",
		},
		"uCore": map[string]interface{}{
			"cats": []interface{}{
				1,
				18,
			},
			"cookies": map[string]interface{}{
				"ucore": "",
			},
			"icon":    "uCore.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "uCore PHP Framework",
			},
			"website": "http://ucore.io",
		},
		"uCoz": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"cookies": map[string]interface{}{
				"uCoz": "",
			},
			"icon":    "uCoz.svg",
			"website": "https://ucoz.ru",
		},
		"uKnowva": map[string]interface{}{
			"cats": []interface{}{
				1,
				2,
				18,
				50,
			},
			"headers": map[string]interface{}{
				"X-Content-Encoded-By": "uKnowva ([\\d.]+)\\;version:\\1",
			},
			"html":    "<a[^>]+>Powered by uKnowva</a>",
			"icon":    "uKnowva.png",
			"implies": "PHP",
			"meta": map[string]interface{}{
				"generator": "uKnowva (?: ([\\d.]+))?\\;version:\\1",
			},
			"script":  "/media/conv/js/jquery\\.js",
			"website": "https://uknowva.com",
		},
		"vBulletin": map[string]interface{}{
			"cats": []interface{}{
				2,
			},
			"html":    "<div id=\"copyright\">Powered by vBulletin",
			"icon":    "vBulletin.png",
			"implies": "PHP",
			"js": map[string]interface{}{
				"vBulletin": "",
			},
			"cookies": map[string]interface{}{
				"bbsessionhash":  "",
				"bblastactivity": "",
				"bblastvisit":    "",
			},
			"meta": map[string]interface{}{
				"generator": "vBulletin ?([\\d.]+)?\\;version:\\1",
			},
			"website": "https://www.vbulletin.com",
		},
		"vibecommerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"excludes": "PrestaShop",
			"icon":     "vibecommerce.png",
			"implies":  "PHP",
			"meta": map[string]interface{}{
				"designer":  "vibecommerce",
				"generator": "vibecommerce",
			},
			"website": "http://vibecommerce.com.br",
		},
		"Virgool": map[string]interface{}{
			"cats": []interface{}{
				11,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Virgool$",
			},
			"icon":    "Virgool.svg",
			"url":     "^https?://(?:www\\.)?virgool\\.io",
			"website": "https://virgool.io",
		},
		"shoperfa": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "^Shoperfa$",
			},
			"icon":    "Shoperfa.png",
			"url":     "^https?://(?:www\\.)?shoperfa\\.com",
			"website": "https://shoperfa.com",
		},
		"webEdition": map[string]interface{}{
			"cats": []interface{}{
				1,
			},
			"icon": "webEdition.png",
			"meta": map[string]interface{}{
				"DC.title":  "webEdition",
				"generator": "webEdition",
			},
			"website": "http://webedition.de/en",
		},
		"webpack": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon": "webpack.svg",
			"js": map[string]interface{}{
				"webpackJsonp": "",
			},
			"website": "https://webpack.js.org/",
		},
		"wisyCMS": map[string]interface{}{
			"cats": []interface{}{
				"1",
			},
			"icon": "wisyCMS.svg",
			"meta": map[string]interface{}{
				"generator": "^wisy CMS[ v]{0,3}([0-9.,]*)\\;version:\\1",
			},
			"website": "https://wisy.3we.de",
		},
		"parcel": map[string]interface{}{
			"cats": []interface{}{
				19,
			},
			"icon": "Parcel.png",
			"js": map[string]interface{}{
				"parcelRequire": "",
			},
			"website": "https://parceljs.org/",
		},
		"wpCache": map[string]interface{}{
			"cats": []interface{}{
				23,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "wpCache(?:/([\\d.]+))?\\;version:\\1",
			},
			"html": "<!--[^>]+wpCache",
			"icon": "wpCache.png",
			"implies": []interface{}{
				"WordPress",
				"PHP",
			},
			"meta": map[string]interface{}{
				"generator": "wpCache",
				"keywords":  "wpCache",
			},
			"url":     "^https?://[^/]+\\.wpcache\\.co",
			"website": "https://wpcache.co",
		},
		"xCharts": map[string]interface{}{
			"cats": []interface{}{
				25,
			},
			"html":    "<link[^>]* href=\"[^\"]*xcharts(?:\\.min)?\\.css",
			"implies": "D3",
			"js": map[string]interface{}{
				"xChart": "",
			},
			"script":  "xcharts\\.js",
			"website": "https://tenxer.github.io/xcharts/",
		},
		"xtCommerce": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"html": "<div class=\"copyright\">[^<]+<a[^>]+>xt:Commerce",
			"icon": "xtCommerce.png",
			"meta": map[string]interface{}{
				"generator": "xt:Commerce",
			},
			"website": "https://www.xt-commerce.com",
		},
		"Yepcomm": map[string]interface{}{
			"cats": []interface{}{
				6,
			},
			"icon": "yepcomm.png",
			"meta": map[string]interface{}{
				"copyright": "Yepcomm Tecnologia",
				"author":    "Yepcomm Tecnologia",
			},
			"website": "https://www.yepcomm.com.br",
		},
		"Halo": map[string]interface{}{
			"cats": []interface{}{
				1,
				11,
			},
			"icon": "Halo.svg",
			"meta": map[string]interface{}{
				"generator": "Halo ([\\d.]+)?\\;version:\\1",
			},
			"implies": "Java",
			"website": "https://halo.run",
		},
		"Rocket": map[string]interface{}{
			"cats": []interface{}{
				1,
				6,
			},
			"headers": map[string]interface{}{
				"x-powered-by": "^Rocket=https://rocketcms.io/",
			},
			"icon": "Rocket.svg",
			"implies": []interface{}{
				"webpack",
				"Node.js",
				"MySQL",
				"Less",
			},
			"website": "https://rocketcms.io",
		},
		"Zipkin": map[string]interface{}{
			"cats": []interface{}{
				10,
			},
			"headers": map[string]interface{}{
				"X-B3-TraceId":      "",
				"X-B3-SpanId":       "",
				"X-B3-ParentSpanId": "",
				"X-B3-Sampled":      "",
				"X-B3-Flags":        "",
			},
			"icon":    "Zipkin.png",
			"website": "https://zipkin.io/",
		},
		"RX Web Server": map[string]interface{}{
			"cats": []interface{}{
				22,
			},
			"headers": map[string]interface{}{
				"X-Powered-By": "RX-WEB",
			},
			"icon":    "RXWeb.svg",
			"website": "http://developers.rokitax.co.uk/projects/rxweb",
		},
	},
	"categories": map[string]interface{}{
		"1": map[string]interface{}{
			"name":     "CMS",
			"priority": 1,
		},
		"2": map[string]interface{}{
			"name":     "Message Boards",
			"priority": 1,
		},
		"3": map[string]interface{}{
			"name":     "Database Managers",
			"priority": 2,
		},
		"4": map[string]interface{}{
			"name":     "Documentation Tools",
			"priority": 2,
		},
		"5": map[string]interface{}{
			"name":     "Widgets",
			"priority": 9,
		},
		"6": map[string]interface{}{
			"name":     "Ecommerce",
			"priority": 1,
		},
		"7": map[string]interface{}{
			"name":     "Photo Galleries",
			"priority": 1,
		},
		"8": map[string]interface{}{
			"name":     "Wikis",
			"priority": 1,
		},
		"9": map[string]interface{}{
			"name":     "Hosting Panels",
			"priority": 1,
		},
		"10": map[string]interface{}{
			"name":     "Analytics",
			"priority": 9,
		},
		"11": map[string]interface{}{
			"name":     "Blogs",
			"priority": 1,
		},
		"12": map[string]interface{}{
			"name":     "JavaScript Frameworks",
			"priority": 8,
		},
		"13": map[string]interface{}{
			"name":     "Issue Trackers",
			"priority": 2,
		},
		"14": map[string]interface{}{
			"name":     "Video Players",
			"priority": 7,
		},
		"15": map[string]interface{}{
			"name":     "Comment Systems",
			"priority": 9,
		},
		"16": map[string]interface{}{
			"name":     "Captchas",
			"priority": 9,
		},
		"17": map[string]interface{}{
			"name":     "Font Scripts",
			"priority": 9,
		},
		"18": map[string]interface{}{
			"name":     "Web Frameworks",
			"priority": 7,
		},
		"19": map[string]interface{}{
			"name":     "Miscellaneous",
			"priority": 9,
		},
		"20": map[string]interface{}{
			"name":     "Editors",
			"priority": 4,
		},
		"21": map[string]interface{}{
			"name":     "LMS",
			"priority": 1,
		},
		"22": map[string]interface{}{
			"name":     "Web Servers",
			"priority": 8,
		},
		"23": map[string]interface{}{
			"name":     "Cache Tools",
			"priority": 7,
		},
		"24": map[string]interface{}{
			"name":     "Rich Text Editors",
			"priority": 5,
		},
		"25": map[string]interface{}{
			"name":     "JavaScript Graphics",
			"priority": 6,
		},
		"26": map[string]interface{}{
			"name":     "Mobile Frameworks",
			"priority": 8,
		},
		"27": map[string]interface{}{
			"name":     "Programming Languages",
			"priority": 5,
		},
		"28": map[string]interface{}{
			"name":     "Operating Systems",
			"priority": 6,
		},
		"29": map[string]interface{}{
			"name":     "Search Engines",
			"priority": 4,
		},
		"30": map[string]interface{}{
			"name":     "Web Mail",
			"priority": 2,
		},
		"31": map[string]interface{}{
			"name":     "CDN",
			"priority": 9,
		},
		"32": map[string]interface{}{
			"name":     "Marketing Automation",
			"priority": 9,
		},
		"33": map[string]interface{}{
			"name":     "Web Server Extensions",
			"priority": 7,
		},
		"34": map[string]interface{}{
			"name":     "Databases",
			"priority": 5,
		},
		"35": map[string]interface{}{
			"name":     "Maps",
			"priority": 6,
		},
		"36": map[string]interface{}{
			"name":     "Advertising Networks",
			"priority": 9,
		},
		"37": map[string]interface{}{
			"name":     "Network Devices",
			"priority": 2,
		},
		"38": map[string]interface{}{
			"name":     "Media Servers",
			"priority": 1,
		},
		"39": map[string]interface{}{
			"name":     "Webcams",
			"priority": 9,
		},
		"41": map[string]interface{}{
			"name":     "Payment Processors",
			"priority": 8,
		},
		"42": map[string]interface{}{
			"name":     "Tag Managers",
			"priority": 9,
		},
		"44": map[string]interface{}{
			"name":     "Build CI Systems",
			"priority": 3,
		},
		"45": map[string]interface{}{
			"name":     "Control Systems",
			"priority": 2,
		},
		"46": map[string]interface{}{
			"name":     "Remote Access",
			"priority": 1,
		},
		"47": map[string]interface{}{
			"name":     "Dev Tools",
			"priority": 2,
		},
		"48": map[string]interface{}{
			"name":     "Network Storage",
			"priority": 2,
		},
		"49": map[string]interface{}{
			"name":     "Feed Readers",
			"priority": 1,
		},
		"50": map[string]interface{}{
			"name":     "Document Management Systems",
			"priority": 1,
		},
		"51": map[string]interface{}{
			"name":     "Landing Page Builders",
			"priority": 2,
		},
		"52": map[string]interface{}{
			"name":     "Live Chat",
			"priority": 8,
		},
		"53": map[string]interface{}{
			"name":     "CRM",
			"priority": 5,
		},
		"54": map[string]interface{}{
			"name":     "SEO",
			"priority": 8,
		},
		"55": map[string]interface{}{
			"name":     "Accounting",
			"priority": 1,
		},
		"56": map[string]interface{}{
			"name":     "Cryptominer",
			"priority": 5,
		},
		"57": map[string]interface{}{
			"name":     "Static Site Generator",
			"priority": 1,
		},
		"58": map[string]interface{}{
			"name":     "User Onboarding",
			"priority": 8,
		},
		"59": map[string]interface{}{
			"name":     "JavaScript Libraries",
			"priority": 9,
		},
		"60": map[string]interface{}{
			"name":     "Containers",
			"priority": 8,
		},
		"61": map[string]interface{}{
			"name":     "SaaS",
			"priority": 8,
		},
		"62": map[string]interface{}{
			"name":     "PaaS",
			"priority": 8,
		},
		"63": map[string]interface{}{
			"name":     "IaaS",
			"priority": 8,
		},
		"64": map[string]interface{}{
			"name":     "Reverse Proxy",
			"priority": 7,
		},
		"65": map[string]interface{}{
			"name":     "Load Balancer",
			"priority": 7,
		},
	},
}
