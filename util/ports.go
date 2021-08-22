package util

// cat ports.txt| grep -v '^#\|udp' > all-tcp-ports.txt
// cat all-tcp-ports.txt | awk '{print $1,$2}' | cut -d '/' -f1 > tcp-ports-filtered.txt
// cat all-tcp-ports.txt | awk '{print $2}' | cut -d '/' -f1 | head -n 1024 > 1024-ports.txt
// cat all-tcp-ports.txt | awk '{print $1}' | cut -d '/' -f1 | head -n 1024 > 1024-services.txt
// paste 1024-services.txt 1024-ports.txt > t
// cat t| tr '\t' ':' | uniq > 1024-ports-services.txt

var KnownPorts = map[string]string{
	"1":    "tcpmux",
	"2":    "compressnet",
	"3":    "compressnet",
	"4":    "unknown",
	"5":    "rje",
	"6":    "unknown",
	"7":    "echo",
	"8":    "unknown",
	"9":    "discard",
	"10":   "unknown",
	"11":   "systat",
	"12":   "unknown",
	"13":   "daytime",
	"14":   "unknown",
	"15":   "netstat",
	"16":   "unknown",
	"17":   "qotd",
	"18":   "msp",
	"19":   "chargen",
	"20":   "ftp-data",
	"21":   "ftp",
	"22":   "ssh",
	"23":   "telnet",
	"24":   "priv-mail",
	"25":   "smtp",
	"26":   "rsftp",
	"27":   "nsw-fe",
	"28":   "unknown",
	"29":   "msg-icp",
	"30":   "unknown",
	"31":   "msg-auth",
	"32":   "unknown",
	"33":   "dsp",
	"34":   "unknown",
	"35":   "priv-print",
	"37":   "time",
	"38":   "rap",
	"39":   "rlp",
	"40":   "unknown",
	"41":   "graphics",
	"42":   "nameserver",
	"43":   "whois",
	"44":   "mpm-flags",
	"45":   "mpm",
	"46":   "mpm-snd",
	"47":   "ni-ftp",
	"48":   "auditd",
	"49":   "tacacs",
	"50":   "re-mail-ck",
	"51":   "la-maint",
	"52":   "xns-time",
	"53":   "domain",
	"54":   "xns-ch",
	"55":   "isi-gl",
	"56":   "xns-auth",
	"57":   "priv-term",
	"58":   "xns-mail",
	"59":   "priv-file",
	"60":   "unknown",
	"61":   "ni-mail",
	"62":   "acas",
	"63":   "via-ftp",
	"64":   "covia",
	"65":   "tacacs-ds",
	"66":   "sqlnet",
	"67":   "dhcps",
	"68":   "dhcpc",
	"69":   "tftp",
	"70":   "gopher",
	"71":   "netrjs-1",
	"72":   "netrjs-2",
	"73":   "netrjs-3",
	"74":   "netrjs-4",
	"75":   "priv-dial",
	"76":   "deos",
	"77":   "priv-rje",
	"78":   "vettcp",
	"79":   "finger",
	"80":   "http",
	"81":   "hosts2-ns",
	"82":   "xfer",
	"83":   "mit-ml-dev",
	"84":   "ctf",
	"85":   "mit-ml-dev",
	"86":   "mfcobol",
	"87":   "priv-term-l",
	"88":   "kerberos-sec",
	"89":   "su-mit-tg",
	"90":   "dnsix",
	"91":   "mit-dov",
	"92":   "npp",
	"93":   "dcp",
	"94":   "objcall",
	"95":   "supdup",
	"96":   "dixie",
	"97":   "swift-rvf",
	"98":   "linuxconf",
	"99":   "metagram",
	"100":  "newacct",
	"101":  "hostname",
	"102":  "iso-tsap",
	"103":  "gppitnp",
	"104":  "acr-nema",
	"105":  "csnet-ns",
	"106":  "pop3pw",
	"107":  "rtelnet",
	"108":  "snagas",
	"109":  "pop2",
	"110":  "pop3",
	"111":  "rpcbind",
	"112":  "mcidas",
	"113":  "ident",
	"114":  "audionews",
	"115":  "sftp",
	"116":  "ansanotify",
	"117":  "uucp-path",
	"118":  "sqlserv",
	"119":  "nntp",
	"120":  "cfdptkt",
	"121":  "erpc",
	"122":  "smakynet",
	"123":  "ntp",
	"124":  "ansatrader",
	"125":  "locus-map",
	"126":  "unitary",
	"127":  "locus-con",
	"128":  "gss-xlicen",
	"129":  "pwdgen",
	"130":  "cisco-fna",
	"131":  "cisco-tna",
	"132":  "cisco-sys",
	"133":  "statsrv",
	"134":  "ingres-net",
	"135":  "msrpc",
	"136":  "profile",
	"137":  "netbios-ns",
	"138":  "netbios-dgm",
	"139":  "netbios-ssn",
	"140":  "emfis-data",
	"141":  "emfis-cntl",
	"142":  "bl-idm",
	"143":  "imap",
	"144":  "news",
	"145":  "uaac",
	"146":  "iso-tp0",
	"147":  "iso-ip",
	"148":  "cronus",
	"149":  "aed-512",
	"150":  "sql-net",
	"151":  "hems",
	"152":  "bftp",
	"153":  "sgmp",
	"154":  "netsc-prod",
	"155":  "netsc-dev",
	"156":  "sqlsrv",
	"157":  "knet-cmp",
	"158":  "pcmail-srv",
	"159":  "nss-routing",
	"160":  "sgmp-traps",
	"161":  "snmp",
	"162":  "snmptrap",
	"163":  "cmip-man",
	"164":  "cmip-agent",
	"165":  "xns-courier",
	"166":  "s-net",
	"167":  "namp",
	"168":  "rsvd",
	"169":  "send",
	"170":  "print-srv",
	"171":  "multiplex",
	"172":  "cl-1",
	"173":  "xyplex-mux",
	"174":  "mailq",
	"175":  "vmnet",
	"176":  "genrad-mux",
	"177":  "xdmcp",
	"178":  "nextstep",
	"179":  "bgp",
	"180":  "ris",
	"181":  "unify",
	"182":  "audit",
	"183":  "ocbinder",
	"184":  "ocserver",
	"185":  "remote-kis",
	"186":  "kis",
	"187":  "aci",
	"188":  "mumps",
	"189":  "qft",
	"190":  "gacp",
	"191":  "prospero",
	"192":  "osu-nms",
	"193":  "srmp",
	"194":  "irc",
	"195":  "dn6-nlm-aud",
	"196":  "dn6-smm-red",
	"197":  "dls",
	"198":  "dls-mon",
	"199":  "smux",
	"200":  "src",
	"201":  "at-rtmp",
	"202":  "at-nbp",
	"203":  "at-3",
	"204":  "at-echo",
	"205":  "at-5",
	"206":  "at-zis",
	"207":  "at-7",
	"208":  "at-8",
	"209":  "tam",
	"210":  "z39.50",
	"211":  "914c-g",
	"212":  "anet",
	"213":  "ipx",
	"214":  "vmpwscs",
	"215":  "softpc",
	"216":  "atls",
	"217":  "dbase",
	"218":  "mpp",
	"219":  "uarps",
	"220":  "imap3",
	"221":  "fln-spx",
	"222":  "rsh-spx",
	"223":  "cdc",
	"224":  "masqdialer",
	"225":  "unknown",
	"226":  "unknown",
	"228":  "unknown",
	"229":  "unknown",
	"230":  "unknown",
	"231":  "unknown",
	"233":  "unknown",
	"234":  "unknown",
	"235":  "unknown",
	"236":  "unknown",
	"237":  "unknown",
	"238":  "unknown",
	"242":  "direct",
	"243":  "sur-meas",
	"244":  "dayna",
	"245":  "link",
	"246":  "dsp3270",
	"247":  "subntbcst_tftp",
	"248":  "bhfhs",
	"249":  "unknown",
	"250":  "unknown",
	"251":  "unknown",
	"252":  "unknown",
	"253":  "unknown",
	"254":  "unknown",
	"255":  "unknown",
	"256":  "fw1-secureremote",
	"257":  "fw1-mc-fwmodule",
	"258":  "fw1-mc-gui",
	"259":  "esro-gen",
	"260":  "openport",
	"261":  "nsiiops",
	"262":  "arcisdms",
	"263":  "hdap",
	"264":  "bgmp",
	"265":  "maybe-fw1",
	"266":  "sst",
	"267":  "td-service",
	"268":  "td-replica",
	"269":  "manet",
	"270":  "gist",
	"271":  "pt-tls",
	"273":  "unknown",
	"276":  "unknown",
	"277":  "unknown",
	"280":  "http-mgmt",
	"281":  "personal-link",
	"282":  "cableport-ax",
	"283":  "rescap",
	"284":  "corerjd",
	"286":  "fxp",
	"287":  "k-block",
	"288":  "unknown",
	"289":  "unknown",
	"293":  "unknown",
	"294":  "unknown",
	"295":  "unknown",
	"300":  "unknown",
	"301":  "unknown",
	"303":  "unknown",
	"305":  "unknown",
	"306":  "unknown",
	"308":  "novastorbakcup",
	"309":  "entrusttime",
	"310":  "bhmds",
	"311":  "asip-webadmin",
	"312":  "vslmp",
	"313":  "magenta-logic",
	"314":  "opalis-robot",
	"315":  "dpsi",
	"316":  "decauth",
	"317":  "zannet",
	"318":  "pkix-timestamp",
	"319":  "ptp-event",
	"320":  "ptp-general",
	"321":  "pip",
	"322":  "rtsps",
	"323":  "rpki-rtr",
	"324":  "rpki-rtr-tls",
	"325":  "unknown",
	"326":  "unknown",
	"329":  "unknown",
	"333":  "texar",
	"334":  "unknown",
	"336":  "unknown",
	"337":  "unknown",
	"340":  "unknown",
	"343":  "unknown",
	"344":  "pdap",
	"345":  "pawserv",
	"346":  "zserv",
	"347":  "fatserv",
	"348":  "csi-sgwp",
	"349":  "mftp",
	"350":  "matip-type-a",
	"351":  "matip-type-b",
	"352":  "dtag-ste-sb",
	"353":  "ndsauth",
	"354":  "bh611",
	"355":  "datex-asn",
	"356":  "cloanto-net-1",
	"357":  "bhevent",
	"358":  "shrinkwrap",
	"359":  "tenebris_nts",
	"360":  "scoi2odialog",
	"361":  "semantix",
	"362":  "srssend",
	"363":  "rsvp_tunnel",
	"364":  "aurora-cmgr",
	"365":  "dtk",
	"366":  "odmr",
	"367":  "mortgageware",
	"368":  "qbikgdp",
	"369":  "rpc2portmap",
	"370":  "codaauth2",
	"371":  "clearcase",
	"372":  "ulistserv",
	"373":  "legent-1",
	"374":  "legent-2",
	"375":  "hassle",
	"376":  "nip",
	"377":  "tnETOS",
	"378":  "dsETOS",
	"379":  "is99c",
	"380":  "is99s",
	"381":  "hp-collector",
	"382":  "hp-managed-node",
	"383":  "hp-alarm-mgr",
	"384":  "arns",
	"385":  "ibm-app",
	"386":  "asa",
	"387":  "aurp",
	"388":  "unidata-ldm",
	"389":  "ldap",
	"390":  "uis",
	"391":  "synotics-relay",
	"392":  "synotics-broker",
	"393":  "dis",
	"394":  "embl-ndt",
	"395":  "netcp",
	"396":  "netware-ip",
	"397":  "mptn",
	"398":  "kryptolan",
	"399":  "iso-tsap-c2",
	"400":  "work-sol",
	"401":  "ups",
	"402":  "genie",
	"403":  "decap",
	"404":  "nced",
	"405":  "ncld",
	"406":  "imsp",
	"407":  "timbuktu",
	"408":  "prm-sm",
	"409":  "prm-nm",
	"410":  "decladebug",
	"411":  "rmt",
	"412":  "synoptics-trap",
	"413":  "smsp",
	"414":  "infoseek",
	"415":  "bnet",
	"416":  "silverplatter",
	"417":  "onmux",
	"418":  "hyper-g",
	"419":  "ariel1",
	"420":  "smpte",
	"421":  "ariel2",
	"422":  "ariel3",
	"423":  "opc-job-start",
	"424":  "opc-job-track",
	"425":  "icad-el",
	"426":  "smartsdp",
	"427":  "svrloc",
	"428":  "ocs_cmu",
	"429":  "ocs_amu",
	"430":  "utmpsd",
	"431":  "utmpcd",
	"432":  "iasd",
	"433":  "nnsp",
	"434":  "mobileip-agent",
	"435":  "mobilip-mn",
	"436":  "dna-cml",
	"437":  "comscm",
	"438":  "dsfgw",
	"439":  "dasp",
	"440":  "sgcp",
	"441":  "decvms-sysmgt",
	"442":  "cvc_hostd",
	"443":  "https",
	"444":  "snpp",
	"445":  "microsoft-ds",
	"446":  "ddm-rdb",
	"447":  "ddm-dfm",
	"448":  "ddm-ssl",
	"449":  "as-servermap",
	"450":  "tserver",
	"451":  "sfs-smp-net",
	"452":  "sfs-config",
	"453":  "creativeserver",
	"454":  "contentserver",
	"455":  "creativepartnr",
	"457":  "scohelp",
	"458":  "appleqtc",
	"459":  "ampr-rcmd",
	"460":  "skronk",
	"461":  "datasurfsrv",
	"462":  "datasurfsrvsec",
	"463":  "alpes",
	"464":  "kpasswd5",
	"465":  "smtps",
	"466":  "digital-vrc",
	"467":  "mylex-mapd",
	"468":  "photuris",
	"469":  "rcp",
	"470":  "scx-proxy",
	"471":  "mondex",
	"472":  "ljk-login",
	"473":  "hybrid-pop",
	"474":  "tn-tl-w1",
	"475":  "tcpnethaspsrv",
	"476":  "tn-tl-fd1",
	"477":  "ss7ns",
	"478":  "spsc",
	"479":  "iafserver",
	"480":  "loadsrv",
	"481":  "dvs",
	"482":  "bgs-nsi",
	"483":  "ulpnet",
	"484":  "integra-sme",
	"485":  "powerburst",
	"486":  "sstats",
	"487":  "saft",
	"488":  "gss-http",
	"489":  "nest-protocol",
	"490":  "micom-pfs",
	"491":  "go-login",
	"492":  "ticf-1",
	"493":  "ticf-2",
	"494":  "pov-ray",
	"495":  "intecourier",
	"496":  "pim-rp-disc",
	"497":  "retrospect",
	"498":  "siam",
	"499":  "iso-ill",
	"500":  "isakmp",
	"501":  "stmf",
	"502":  "mbap",
	"503":  "intrinsa",
	"504":  "citadel",
	"505":  "mailbox-lm",
	"506":  "ohimsrv",
	"507":  "crs",
	"508":  "xvttp",
	"509":  "snare",
	"510":  "fcp",
	"511":  "passgo",
	"512":  "exec",
	"513":  "login",
	"514":  "shell",
	"515":  "printer",
	"516":  "videotex",
	"517":  "talk",
	"518":  "ntalk",
	"519":  "utime",
	"520":  "efs",
	"521":  "ripng",
	"522":  "ulp",
	"523":  "ibm-db2",
	"524":  "ncp",
	"525":  "timed",
	"526":  "tempo",
	"527":  "stx",
	"528":  "custix",
	"529":  "irc",
	"530":  "courier",
	"531":  "conference",
	"532":  "netnews",
	"533":  "netwall",
	"534":  "mm-admin",
	"535":  "iiop",
	"536":  "opalis-rdv",
	"537":  "nmsp",
	"538":  "gdomap",
	"539":  "apertus-ldp",
	"540":  "uucp",
	"541":  "uucp-rlogin",
	"542":  "commerce",
	"543":  "klogin",
	"544":  "kshell",
	"545":  "ekshell",
	"546":  "dhcpv6-client",
	"547":  "dhcpv6-server",
	"548":  "afp",
	"549":  "idfp",
	"550":  "new-rwho",
	"551":  "cybercash",
	"552":  "deviceshare",
	"553":  "pirp",
	"554":  "rtsp",
	"555":  "dsf",
	"556":  "remotefs",
	"557":  "openvms-sysipc",
	"558":  "sdnskmp",
	"559":  "teedtap",
	"560":  "rmonitor",
	"561":  "monitor",
	"562":  "chshell",
	"563":  "snews",
	"564":  "9pfs",
	"565":  "whoami",
	"566":  "streettalk",
	"567":  "banyan-rpc",
	"568":  "ms-shuttle",
	"569":  "ms-rome",
	"570":  "meter",
	"571":  "umeter",
	"572":  "sonar",
	"573":  "banyan-vip",
	"574":  "ftp-agent",
	"575":  "vemmi",
	"576":  "ipcd",
	"577":  "vnas",
	"578":  "ipdd",
	"579":  "decbsrv",
	"580":  "sntp-heartbeat",
	"581":  "bdp",
	"582":  "scc-security",
	"583":  "philips-vc",
	"584":  "keyserver",
	"585":  "imap4-ssl",
	"586":  "password-chg",
	"587":  "submission",
	"588":  "cal",
	"589":  "eyelink",
	"590":  "tns-cml",
	"591":  "http-alt",
	"592":  "eudora-set",
	"593":  "http-rpc-epmap",
	"594":  "tpip",
	"595":  "cab-protocol",
	"596":  "smsd",
	"597":  "ptcnameservice",
	"598":  "sco-websrvrmg3",
	"599":  "acp",
	"600":  "ipcserver",
	"601":  "syslog-conn",
	"602":  "xmlrpc-beep",
	"603":  "mnotes",
	"604":  "tunnel",
	"605":  "soap-beep",
	"606":  "urm",
	"607":  "nqs",
	"608":  "sift-uft",
	"609":  "npmp-trap",
	"610":  "npmp-local",
	"611":  "npmp-gui",
	"612":  "hmmp-ind",
	"613":  "hmmp-op",
	"614":  "sshell",
	"615":  "sco-inetmgr",
	"616":  "sco-sysmgr",
	"617":  "sco-dtmgr",
	"618":  "dei-icda",
	"619":  "compaq-evm",
	"620":  "sco-websrvrmgr",
	"621":  "escp-ip",
	"622":  "collaborator",
	"623":  "oob-ws-http",
	"624":  "cryptoadmin",
	"625":  "apple-xsrvr-admin",
	"626":  "apple-imap-admin",
	"627":  "passgo-tivoli",
	"628":  "qmqp",
	"629":  "3com-amp3",
	"630":  "rda",
	"631":  "ipp",
	"632":  "bmpp",
	"633":  "servstat",
	"634":  "ginad",
	"635":  "rlzdbase",
	"636":  "ldapssl",
	"637":  "lanserver",
	"638":  "mcns-sec",
	"639":  "msdp",
	"640":  "entrust-sps",
	"641":  "repcmd",
	"642":  "esro-emsdp",
	"643":  "sanity",
	"644":  "dwr",
	"645":  "pssc",
	"646":  "ldp",
	"647":  "dhcp-failover",
	"648":  "rrp",
	"649":  "cadview-3d",
	"650":  "obex",
	"651":  "ieee-mms",
	"652":  "hello-port",
	"653":  "repscmd",
	"654":  "aodv",
	"655":  "tinc",
	"656":  "spmp",
	"657":  "rmc",
	"658":  "tenfold",
	"659":  "unknown",
	"660":  "mac-srvr-admin",
	"661":  "hap",
	"662":  "pftp",
	"663":  "purenoise",
	"664":  "secure-aux-bus",
	"665":  "sun-dr",
	"666":  "doom",
	"667":  "disclose",
	"668":  "mecomm",
	"669":  "meregister",
	"670":  "vacdsm-sws",
	"671":  "vacdsm-app",
	"672":  "vpps-qua",
	"673":  "cimplex",
	"674":  "acap",
	"675":  "dctp",
	"676":  "vpps-via",
	"677":  "vpp",
	"678":  "ggf-ncp",
	"679":  "mrm",
	"680":  "entrust-aaas",
	"681":  "entrust-aams",
	"682":  "xfr",
	"683":  "corba-iiop",
	"684":  "corba-iiop-ssl",
	"685":  "mdc-portmapper",
	"686":  "hcp-wismar",
	"687":  "asipregistry",
	"688":  "realm-rusd",
	"689":  "nmap",
	"690":  "vatp",
	"691":  "resvc",
	"692":  "hyperwave-isp",
	"693":  "connendp",
	"694":  "ha-cluster",
	"695":  "ieee-mms-ssl",
	"696":  "rushd",
	"697":  "uuidgen",
	"698":  "olsr",
	"699":  "accessnetwork",
	"700":  "epp",
	"701":  "lmp",
	"702":  "iris-beep",
	"703":  "unknown",
	"704":  "elcsd",
	"705":  "agentx",
	"706":  "silc",
	"707":  "borland-dsj",
	"708":  "unknown",
	"709":  "entrustmanager",
	"710":  "entrust-ash",
	"711":  "cisco-tdp",
	"712":  "tbrpf",
	"713":  "iris-xpc",
	"714":  "iris-xpcs",
	"715":  "iris-lwz",
	"716":  "pana",
	"717":  "unknown",
	"718":  "unknown",
	"719":  "unknown",
	"720":  "unknown",
	"721":  "unknown",
	"722":  "unknown",
	"723":  "omfs",
	"724":  "unknown",
	"725":  "unknown",
	"726":  "unknown",
	"727":  "unknown",
	"728":  "unknown",
	"729":  "netviewdm1",
	"730":  "netviewdm2",
	"731":  "netviewdm3",
	"732":  "unknown",
	"733":  "unknown",
	"734":  "unknown",
	"735":  "unknown",
	"736":  "unknown",
	"737":  "unknown",
	"738":  "unknown",
	"739":  "unknown",
	"740":  "netcp",
	"741":  "netgw",
	"742":  "netrcs",
	"743":  "unknown",
	"744":  "flexlm",
	"745":  "unknown",
	"746":  "unknown",
	"747":  "fujitsu-dev",
	"748":  "ris-cm",
	"749":  "kerberos-adm",
	"750":  "kerberos",
	"751":  "kerberos_master",
	"752":  "qrh",
	"753":  "rrh",
	"754":  "krb_prop",
	"755":  "unknown",
	"756":  "unknown",
	"757":  "unknown",
	"758":  "nlogin",
	"759":  "con",
	"760":  "krbupdate",
	"761":  "kpasswd",
	"762":  "quotad",
	"763":  "cycleserv",
	"764":  "omserv",
	"765":  "webster",
	"766":  "unknown",
	"767":  "phonebook",
	"768":  "unknown",
	"769":  "vid",
	"770":  "cadlock",
	"771":  "rtip",
	"772":  "cycleserv2",
	"773":  "submit",
	"774":  "rpasswd",
	"775":  "entomb",
	"776":  "wpages",
	"777":  "multiling-http",
	"778":  "unknown",
	"779":  "unknown",
	"780":  "wpgs",
	"781":  "hp-collector",
	"782":  "hp-managed-node",
	"783":  "spamassassin",
	"784":  "unknown",
	"785":  "unknown",
	"786":  "concert",
	"787":  "qsc",
	"788":  "unknown",
	"789":  "unknown",
	"790":  "unknown",
	"791":  "unknown",
	"792":  "unknown",
	"793":  "unknown",
	"794":  "unknown",
	"795":  "unknown",
	"796":  "unknown",
	"797":  "unknown",
	"798":  "unknown",
	"799":  "controlit",
	"800":  "mdbs_daemon",
	"801":  "device",
	"802":  "mbap-s",
	"803":  "unknown",
	"804":  "unknown",
	"805":  "unknown",
	"806":  "unknown",
	"807":  "unknown",
	"808":  "ccproxy-http",
	"809":  "unknown",
	"811":  "unknown",
	"812":  "unknown",
	"813":  "unknown",
	"814":  "unknown",
	"815":  "unknown",
	"816":  "unknown",
	"817":  "unknown",
	"818":  "unknown",
	"819":  "unknown",
	"820":  "unknown",
	"821":  "unknown",
	"822":  "unknown",
	"823":  "unknown",
	"824":  "unknown",
	"825":  "unknown",
	"826":  "unknown",
	"827":  "unknown",
	"828":  "itm-mcell-s",
	"829":  "pkix-3-ca-ra",
	"830":  "netconf-ssh",
	"831":  "netconf-beep",
	"832":  "netconfsoaphttp",
	"833":  "netconfsoapbeep",
	"834":  "unknown",
	"835":  "unknown",
	"836":  "unknown",
	"837":  "unknown",
	"838":  "unknown",
	"839":  "unknown",
	"840":  "unknown",
	"841":  "unknown",
	"842":  "unknown",
	"843":  "unknown",
	"844":  "unknown",
	"845":  "unknown",
	"846":  "unknown",
	"847":  "dhcp-failover2",
	"848":  "gdoi",
	"849":  "unknown",
	"850":  "unknown",
	"851":  "unknown",
	"852":  "unknown",
	"853":  "domain-s",
	"854":  "dlep",
	"855":  "unknown",
	"856":  "unknown",
	"857":  "unknown",
	"858":  "unknown",
	"859":  "unknown",
	"860":  "iscsi",
	"861":  "owamp-control",
	"862":  "twamp-control",
	"863":  "unknown",
	"864":  "unknown",
	"865":  "unknown",
	"866":  "unknown",
	"867":  "unknown",
	"868":  "unknown",
	"869":  "unknown",
	"870":  "unknown",
	"871":  "supfilesrv",
	"872":  "unknown",
	"873":  "rsync",
	"874":  "unknown",
	"875":  "unknown",
	"876":  "unknown",
	"877":  "unknown",
	"878":  "unknown",
	"879":  "unknown",
	"880":  "unknown",
	"881":  "unknown",
	"882":  "unknown",
	"883":  "unknown",
	"884":  "unknown",
	"885":  "unknown",
	"886":  "iclcnet-locate",
	"887":  "iclcnet_svinfo",
	"888":  "accessbuilder",
	"889":  "unknown",
	"890":  "unknown",
	"891":  "unknown",
	"892":  "unknown",
	"893":  "unknown",
	"894":  "unknown",
	"895":  "unknown",
	"896":  "unknown",
	"897":  "unknown",
	"898":  "sun-manageconsole",
	"899":  "unknown",
	"900":  "omginitialrefs",
	"901":  "samba-swat",
	"902":  "iss-realsecure",
	"903":  "iss-console-mgr",
	"904":  "unknown",
	"905":  "unknown",
	"906":  "unknown",
	"907":  "unknown",
	"908":  "unknown",
	"909":  "unknown",
	"910":  "kink",
	"911":  "xact-backup",
	"912":  "apex-mesh",
	"913":  "apex-edge",
	"914":  "unknown",
	"915":  "unknown",
	"916":  "unknown",
	"918":  "unknown",
	"919":  "unknown",
	"920":  "unknown",
	"921":  "unknown",
	"922":  "unknown",
	"923":  "unknown",
	"924":  "unknown",
	"925":  "unknown",
	"926":  "unknown",
	"927":  "unknown",
	"928":  "unknown",
	"929":  "unknown",
	"930":  "unknown",
	"931":  "unknown",
	"932":  "unknown",
	"933":  "unknown",
	"934":  "unknown",
	"935":  "unknown",
	"936":  "unknown",
	"937":  "unknown",
	"938":  "unknown",
	"939":  "unknown",
	"941":  "unknown",
	"942":  "unknown",
	"943":  "unknown",
	"944":  "unknown",
	"945":  "unknown",
	"946":  "unknown",
	"947":  "unknown",
	"948":  "unknown",
	"949":  "unknown",
	"950":  "oftep-rpc",
	"951":  "unknown",
	"952":  "unknown",
	"953":  "rndc",
	"954":  "unknown",
	"955":  "unknown",
	"956":  "unknown",
	"957":  "unknown",
	"958":  "unknown",
	"959":  "unknown",
	"960":  "unknown",
	"961":  "unknown",
	"962":  "unknown",
	"963":  "unknown",
	"964":  "unknown",
	"965":  "unknown",
	"966":  "unknown",
	"967":  "unknown",
	"968":  "unknown",
	"969":  "unknown",
	"970":  "unknown",
	"971":  "unknown",
	"972":  "unknown",
	"973":  "unknown",
	"974":  "unknown",
	"975":  "securenetpro-sensor",
	"976":  "unknown",
	"977":  "unknown",
	"978":  "unknown",
	"979":  "unknown",
	"980":  "unknown",
	"981":  "unknown",
	"982":  "unknown",
	"983":  "unknown",
	"984":  "unknown",
	"985":  "unknown",
	"986":  "unknown",
	"987":  "unknown",
	"988":  "unknown",
	"989":  "ftps-data",
	"990":  "ftps",
	"991":  "nas",
	"992":  "telnets",
	"993":  "imaps",
	"994":  "ircs",
	"995":  "pop3s",
	"996":  "xtreelic",
	"997":  "maitrd",
	"998":  "busboy",
	"999":  "garcon",
	"1000": "cadlock",
	"1001": "webpush",
	"1002": "windows-icfw",
	"1003": "unknown",
	"1004": "unknown",
	"1005": "unknown",
	"1006": "unknown",
	"1007": "unknown",
	"1008": "ufsd",
	"1009": "unknown",
	"1010": "surf",
	"1011": "unknown",
	"1012": "unknown",
	"1013": "unknown",
	"1014": "unknown",
	"1015": "unknown",
	"1016": "unknown",
	"1017": "unknown",
	"1018": "unknown",
	"1019": "unknown",
	"1020": "unknown",
	"1021": "exp1",
	"1022": "exp2",
	"1023": "netvenuechat",
	"1024": "kdm",
	"1025": "NFS-or-IIS",
	"1026": "LSA-or-nterm",
	"1027": "IIS",
	"1028": "unknown",
	"1029": "ms-lsa",
	"1030": "iad1",
	"1031": "iad2",
	"1032": "iad3",
	"1033": "netinfo",
	"1034": "zincite-a",
	"1035": "multidropper",
	"1036": "nsstp",
	"1037": "ams",
	"1038": "mtqp",
	"1039": "sbl",
	"1040": "netsaint",
	"1041": "danf-ak2",
	"1042": "afrog",
	"1043": "boinc",
	"1044": "dcutility",
	"1045": "fpitp",
	"1046": "wfremotertm",
	"1047": "neod1",
	"1048": "neod2",
	"1049": "td-postman",
	"1050": "java-or-OTGfileshare",
	"1051": "optima-vnet",
	"1052": "ddt",
}
