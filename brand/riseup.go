// PhishDetect
// Copyright (C) 2018  Claudio Guarnieri
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package brand

// RiseUp brand properties.
func RiseUp() *Brand {
	name := "riseup"
	original := []string{"riseup"}
	whitelist := []string{"riseup.net"}
	suspicious := []string{
		"riseupa", "riseupb",
		"riseupc", "riseupd", "riseupe",
		"riseupf", "riseupg", "riseuph",
		"riseupi", "riseupj", "riseupk",
		"riseupl", "riseupm", "riseupn",
		"riseupo", "riseupp", "riseupq",
		"riseupr", "riseups", "riseupt",
		"riseupu", "riseupv", "riseupw",
		"riseupx", "riseupy", "riseupz",
		"siseup", "piseup", "viseup",
		"ziseup", "biseup", "2iseup",
		"rhseup", "rkseup", "rmseup",
		"raseup", "ryseup", "rireup",
		"riqeup", "riweup", "riceup",
		"ri3eup", "risdup", "risgup",
		"risaup", "rismup", "risuup",
		"risetp", "risewp", "riseqp",
		"riseep", "rise5p", "riseuq",
		"riseur", "riseut", "riseux",
		"riseu0", "xn--riep-shf471t", "xn--rseu-bnb924a",
		"xn--isup-cva3324b", "xn--rsep-shf3808v", "xn--riup-mva30t",
		"xn--rieu-m5a265b", "xn--rsup-hwa61r", "xn--rsep-vpa421e",
		"xn--rseu-5pa56j", "xn--seup-zya9704b", "xn--risp-d9d69o",
		"xn--seup-48b16y", "xn--risp-nva610c", "xn--iseu-esa7534b",
		"xn--r1sep-0fb", "xn--riep-n53a203i", "xn--ieup-zcc74z",
		"xn--rseu-bnb990a", "xn--isep-fdc9891b", "xn--rseu-esa067z",
		"xn--seup-48b3j", "xn--risp-xva078d", "xn--risp-xod3152b",
		"xn--isup-hpa32v", "xn--risp-nmd05y", "xn--risu-xva735b",
		"xn--isup-cva10t", "xn--iseu-fdc47t", "xn--reup-qdc534x",
		"xn--rsep-58b40s", "xn--rsep-g9t602f", "xn--seup-f9t932f",
		"xn--rsep-5pa211e", "xn--risp-iwa8214b", "xn--rsup-vpa7w",
		"xn--rsup-hpa41b", "xn--rlsep-x49a", "xn--lseup-1rc",
		"xn--isep-k4d1j", "xn--rsep-nnd945v", "xn--rseu-8vd6140w",
		"xn--risp-dva452b", "xn--seup-zya92r", "xn--rise-8vd74u",
		"xn--rlsep-zze", "xn--risu-dpa318b", "xn--rise-87a096a",
		"xn--rlsep-cce", "xn--isep-n53a8b", "xn--isup-k4d44f",
		"xn--riup-cva7578d", "xn--rsep-5pa211e", "xn--rsup-g9tz85g",
		"xn--riep-h9d3719a", "xn--risu-8vd46c", "xn--ieup-0huw60f",
		"xn--rsup-wva651b", "xn--isep-xod07d", "xn--rsup-wva0484w",
		"xn--risu-eod49f", "xn--isup-mva56s", "xn--rsup-hpa288z",
		"xn--riup-rdc65p", "xn--rsup-w4d2500w", "xn--risu-dpa384b",
		"xn--riup-c9d759u", "xn--rise-eod77x", "xn--riup-wva12b",
		"xn--risu-npa9i", "xn--isep-xnb24h", "xn--isep-fdc76q",
		"xn--rsup-lza339b", "xn--riup-1hu404g", "xn--rsup-bnb1255b",
		"xn--riup-w4d5493d", "xn--rise-ykb761a", "xn--rseu-58b34z",
		"xn--isup-6va74s", "xn--riup-mpa753c", "xn--risp-dva6714b",
		"xn--r1sup-kza", "xn--rsep-5mdu", "xn--isup-mpa71v",
		"xn--reup-5md05g", "xn--rsup-hwa6184w", "xn--riup-hwa79a",
		"xn--rsep-xnb50g", "xn--risp-xva400c", "xn--rieu-j6d7j",
		"xn--reup-58b660a", "xn--risp-iwa251b", "xn--isup-hpa92v",
		"xn--riup-c9df", "xn--risp-nva288d", "xn--rsup-lza5746b",
		"xn--isep-zcc58q", "xn--risp-77a175b", "xn--isup-cpa5844b",
		"xn--ieup-zcc587x", "xn--rsup-5pa5v", "xn--riep-m5a5473b",
		"xn--r1sep-bjg", "xn--iseu-fdc60y", "xn--seup-zya32r",
		"xn--rsup-cpa61u", "xn--risp-dva498d", "xn--seup-9mb15h",
		"xn--riup-mmd3174d", "xn--rsup-mmd295v", "xn--risp-npa431e",
		"xn--rsup-6va2j", "xn--rsup-hpa2a", "xn--rise-j6d7929a",
		"xn--riup-wva19s", "xn--rsup-58b160a", "xn--rsep-lza8993b",
		"xn--riup-hwa960c", "xn--risp-x4d28q", "xn--risp-xnb5945b",
		"xn--risp-iwa648d", "xn--isup-zcc4844b", "xn--rlsep-oxb",
		"xn--reup-5pa2710a", "xn--risu-esa2m", "xn--r1seu-zye",
		"xn--reup-58b6076d", "xn--risp-npa49e", "xn--risu-nmd94f",
		"xn--isup-mva98s", "xn--isep-ucc73q", "xn--isep-zcc13q",
		"xn--reup-5md0964d", "xn--rsup-6va8284w", "xn--seup-kza5504b",
		"xn--rieu-ykb3277d", "xn--iseu-ykb9913b", "xn--rsup-5pa923c",
		"xn--rieu-esa1888d", "xn--reup-lza839b", "xn--rise-ynb204a",
		"xn--risu-nva189b", "xn--seup-fdc29p", "xn--seup-9mb54h",
		"xn--risu-j6d4771b", "xn--rsep-vpa034b", "xn--rise-ynb270a",
		"xn--riup-6va90b", "xn--rsep-5pa274b", "xn--rsep-d7d127u",
		"xn--isup-k4do", "xn--riup-w4d341v", "xn--isup-mpa32v",
		"xn--rsup-vpa179c", "xn--lseup-j69a", "xn--isup-mpa3744b",
		"xn--risu-iwa315b", "xn--rseu-esa8894w", "xn--rieu-eod738v",
		"xn--iseu-zcc26q", "xn--iseu-ucc86q", "xn--rseu-vpa77j",
		"xn--riup-h9d83d", "xn--rsup-mva6l", "xn--riup-cpa7020a",
		"xn--rise-87a020b", "xn--risu-ipa932c", "xn--risp-d7d72e",
		"xn--rseu-8vd881v", "xn--reup-bnb228y", "xn--riep-xnb697y",
		"xn--rieu-rdc06t", "xn--rieu-eod9364d", "xn--riup-mva5478d",
		"xn--isep-nnd6410b", "xn--rise-esa523b", "xn--seup-upa13v",
		"xn--rlsup-kza", "xn--reup-5pa26d", "xn--risu-iwa381b",
		"xn--rsup-cpa44b", "xn--iseu-ucc89t", "xn--iseu-zcc29t",
		"xn--risp-nva881b", "xn--risu-xva702b", "xn--risu-7va759b",
		"xn--seup-4pa10v", "xn--rise-j6d9a", "xn--riup-6va1278d",
		"xn--rlseu-3tb", "xn--rieu-rdc03q", "xn--isup-6va56s",
		"xn--riup-m5a494b", "xn--rise-87a60c", "xn--rsep-0ya047d",
		"xn--reup-vpa47d", "xn--risu-npa78j", "xn--isup-c9d5329a",
		"xn--risu-esa050c", "xn--isup-ucc875a", "xn--rsup-0ya300b",
		"xn--seup-4md54e", "xn--r1eup-gsc", "xn--risu-eod80e",
		"xn--rsep-lza220b", "xn--rseu-eod815v", "xn--rsep-5md82y",
		"xn--rsep-0ya640b", "xn--rieu-ykb18i", "xn--rlsup-39d",
		"xn--isup-k4d4r", "xn--r1eup-i2e", "xn--iseu-k4d6f",
		"xn--rsup-mva0o", "xn--rieu-esa9700a", "xn--riep-77a47n",
		"xn--risp-7va858d", "xn--seup-k4d6600w", "xn--risu-dva399b",
		"xn--r1sep-x49a", "xn--isep-nnd42e", "xn--rsep-58b05r",
		"xn--rsup-mpak", "xn--risu-8vd49i", "xn--reup-g9d336u",
		"xn--reup-l5a33n", "xn--risp-dpa641e", "xn--isup-cva119b",
		"xn--rleup-i2e", "xn--rsup-hpa01u", "xn--seup-4pa921c",
		"xn--rleup-gsc", "xn--rseu-lza518b", "xn--reup-0ya008z",
		"xn--riep-d7d2g", "xn--rsup-mmd0280w", "xn--isup-cpa361c",
		"xn--1seup-j69a", "xn--riup-hwa709z", "xn--rseu-ykb2h",
		"xn--rsup-wva22i", "xn--riep-d7d2483d", "xn--rsup-cva4684w",
		"xn--isup-mva998b", "xn--rsep-lza670b", "xn--rsup-hpa0115w",
		"xn--seup-kza32r", "xn--risp-dpa60f", "xn--rsup-cva081b",
		"xn--rsup-0ya785c", "xn--rseu-0ya704b", "xn--rsup-hwa89h",
		"xn--rise-8vd18b", "xn--ieup-fdc767x", "xn--rsup-mva426z",
		"xn--rise-eod77x", "xn--riup-mu5ap81h", "xn--rseu-vpa322c",
		"xn--rlsup-jsa", "xn--rsup-vpa551c", "xn--reup-qdc3662w",
		"xn--reup-5pa4798d", "xn--risu-ykb807b", "xn--isep-77a63n",
		"xn--risp-shf3789a", "xn--rsup-hpa29k", "xn--reup-vpa4810a",
		"xn--rise-yod25e", "xn--isup-cpa53v", "xn--seup-48b3812b",
		"xn--riup-hpa3009d", "xn--riup-6va180c", "xn--reup-lza677z",
		"xn--rsup-5pa9x", "xn--risp-nmd6b", "xn--isep-ucc3102b",
		"xn--isep-fdc707b", "xn--isup-k4d6581b", "xn--ieup-fdc92z",
		"xn--rsep-77a3234w", "xn--rsep-5pa4934b", "xn--rsup-cva8m",
		"xn--riep-xod418v", "xn--risu-esa6o", "xn--rsep-bnb267c",
		"xn--isep-k4d79q", "xn--riep-m5a364d", "xn--riep-1hut30f",
		"xn--rise-ond4a", "xn--rsep-xnb774y", "xn--rseu-58b11v",
		"xn--rsep-5pa814b", "xn--rise-ykb168c", "xn--risu-dpa99j",
		"xn--rsup-vpa1z", "xn--riup-cva559z", "xn--risu-ykb870b",
		"xn--rsup-vpa9x", "xn--reup-bnb489a", "xn--seup-ucc0072w",
		"xn--rlsup-1we", "xn--rsup-cpa02b", "xn--rsep-lza49g",
		"xn--riup-hpa16v", "xn--rleup-vcb", "xn--riep-rdc396b",
		"xn--risp-7va81c", "xn--risp-npa6144b", "xn--rsup-mpa68k",
		"xn--rlsup-psa", "xn--isup-6va578b", "xn--rsup-bnb906b",
		"xn--rsup-6va441b", "xn--r1sep-cce", "xn--riep-m5a795b",
		"xn--riup-hpa19d", "xn--risu-j6d2j", "xn--iseu-8vd5879a",
		"xn--risu-ykb298a", "xn--risp-dpa8244b", "xn--risp-npa044b",
		"xn--riup-wva139z", "xn--risp-dpa605b", "xn--riup-m5a627a",
		"xn--isep-fdc31q", "xn--rsep-0ya001b", "xn--r1eup-vcb",
		"xn--risu-esa668c", "xn--isep-zcc95y", "xn--risp-dpa082c",
		"xn--risp-npa494b", "xn--rsep-58b6212b", "xn--rsup-mpa678z",
		"xn--rseu-esa07j", "xn--seup-48b1l", "xn--isup-6va14s",
		"xn--iseu-eod5110b", "xn--risp-xva03c", "xn--isep-shf2547a",
		"xn--seup-4pa1544b", "xn--rsup-cpe800u", "xn--rsup-lza747b",
		"xn--rsup-w4d428u", "xn--rsep-d7d9499v", "xn--isup-mpa151c",
		"xn--rseu-5pa112c", "xn--rsup-vpa5v", "xn--rsup-bnb397a",
		"xn--r1sep-vde", "xn--1seup-1rc", "xn--riup-cpa963c",
		"xn--riup-rdc2544b", "xn--rsup-cva636z", "xn--rsup-6va01i",
		"xn--risp-xva078d", "xn--reup-0ya269b", "xn--lseup-drc",
		"xn--rsup-mva25r", "xn--rieu-esa93u", "xn--rsup-5pa573b",
		"xn--rseu-vpa164b", "xn--rsup-cpa8a", "xn--rieu-ykb129y",
		"xn--seup-ucc274x", "xn--risp-xnb359a", "xn--seup-zya74r",
		"xn--riup-m5a2226b", "xn--riep-shf471t", "xn--isup-mmd9810b",
		"xn--risu-npa108b", "xn--riup-h9d0561b", "xn--rlseu-nme",
		"xn--risp-d9d8719a", "xn--ieup-k4d751v", "xn--rsup-wva04r",
		"xn--rsup-mva861b", "xn--risp-xnb385b", "xn--risp-dpa42l",
		"xn--risu-npa174b", "xn--rise-esa910e", "xn--riup-cpa9009d",
		"xn--risu-dpa1k", "xn--rlseu-7ce", "xn--rsep-xnb5073w",
		"xn--risp-npa21l", "xn--risu-esa632c", "xn--ieup-zcc7856d",
		"xn--riep-rdc5791b", "xn--reup-vv0c6755e", "xn--risu-xva32h",
		"xn--riup-mva501c", "xn--seup-zcc01q", "xn--isup-hwa92s",
		"xn--rlsep-0fb", "xn--rseu-eod6470w", "xn--risu-iwa99g",
		"xn--rsep-lza617d", "xn--rsep-nnd7770w", "xn--isup-cva38s",
		"xn--rseu-0ya938b", "xn--rseup-hoc", "xn--rseu-ykb206y",
		"xn--isup-cpe5068a", "xn--rsep-bnb4403b", "xn--r1sup-jsa",
		"xn--1seup-drc", "xn--r1sep-oxb", "xn--rise-j6d51q",
		"xn--rlsup-l0a", "xn--isup-ucc0944b", "xn--risu-esa8776b",
		"xn--r1sup-p51b", "xn--riup-c9d9573d", "xn--risp-d7d9371b",
		"xn--seup-zya757b", "xn--rsup-6va82r", "xn--iseu-ykb34i",
		"xn--iseu-ykb94i", "xn--riup-mva33b", "xn--rlsep-bjg",
		"xn--riep-h9d19o", "xn--riup-wva3378d", "xn--isup-wva35s",
		"xn--rseu-5md72f", "xn--ieup-qdc2302b", "xn--isup-wva9024b",
		"xn--rseu-ykb641a", "xn--reup-5md888v", "xn--rsep-0ya047d",
		"xn--risp-dva27i", "xn--riup-hpa1020a", "xn--rsep-lza058b",
		"xn--risp-nmd05y", "xn--riup-hwa9078d", "xn--riep-m5a967a",
		"xn--r1sep-zze", "xn--isup-fdc07p", "xn--risu-npa332c",
		"xn--rlseu-0va", "xn--rsup-cva46r", "xn--ieup-qdc01x",
		"xn--rlseu-zye", "xn--risp-dva820c", "xn--isup-mpa14v",
		"xn--rsep-bnb221a", "xn--seup-z63ay786g", "xn--rsep-5pa642c",
		"xn--riep-m5a328a", "xn--riup-mva349z", "xn--rise-ykb599a",
		"xn--rsep-0ya00b", "xn--seup-48b7i", "xn--ieup-l5a084b",
		"xn--isup-wva77s", "xn--rsep-shf3808v", "xn--rsup-5pa1786b",
		"xn--risp-nva24c", "xn--riep-xod6164d", "xn--rseu-5md59a",
		"xn--rsep-vpa421e", "xn--riup-cva51t", "xn--isup-ucc49p",
		"xn--rseu-esa88s", "xn--rise-8vd74u", "xn--rseu-lza96f",
		"xn--rseu-0ya770b", "xn--rleup-769c", "xn--isep-xnb42h",
		"xn--r1sup-dsa", "xn--seup-upa3644b", "xn--riep-77a414z",
		"xn--reup-l5a502z", "xn--seup-9mb96h", "xn--risu-ipa39j",
		"xn--risp-nmd48f", "xn--rsup-vpa3u", "xn--risp-7va299b",
		"xn--ieup-k4d9593d", "xn--isup-hwa5814b", "xn--rsup-mva2584w",
		"xn--isup-fdc6644b", "xn--risp-nva288d", "xn--seup-k4d838u",
		"xn--rsep-vpa48e", "xn--riep-shf6791d", "xn--1seup-mwe",
		"xn--rieu-esa142c", "xn--rseu-5md5c", "xn--riep-xod67f",
		"xn--seup-upa70v", "xn--risp-ipa644b", "xn--ieup-ucc1b",
		"xn--rieu-esa96c", "xn--rsep-bnb698a", "xn--risu-eod42m",
		"xn--isup-ucc26x", "xn--rlsup-d2e", "xn--risp-d9d69o",
		"xn--rsup-wva8m", "xn--rsup-vpa3886b", "xn--seup-4pa91v",
		"xn--risp-7va461b", "xn--riep-rdc396b", "xn--risp-npa431e",
		"xn--risu-ykb0265b", "xn--risp-dpe66i", "xn--risu-nva53h",
		"xn--isup-hpa751c", "xn--rseu-vpa197b", "xn--riep-xnb8957d",
		"xn--iseu-j6d3539a", "xn--risp-dva002b", "xn--riup-cva54b",
		"xn--riep-nnd03g", "xn--reup-58b409x", "xn--risp-nva242b",
		"xn--riup-mmd37g", "xn--risu-ykb421a", "xn--isup-hwa35s",
		"xn--riep-m5a14e", "xn--r1seu-3tb", "xn--r1sep-bjg",
		"xn--rieu-8vd9034d", "xn--rseu-5pa944b", "xn--risp-77a102c",
		"xn--rsup-5md57m", "xn--rsup-c9d836u", "xn--r1sup-8za",
		"xn--rsup-hpa83b", "xn--rsep-bnb267c", "xn--seup-kza99q",
		"xn--risp-77a716a", "xn--rseu-5pa977b", "xn--rise-ykb3413b",
		"xn--rsup-cpa234b", "xn--riup-rdc42x", "xn--rsep-5md8e",
		"xn--rseu-ykb0383w", "xn--risp-nnd7652b", "xn--ieup-l5a2083b",
		"xn--r1sup-x0a", "xn--rsup-mpa4015w", "xn--iseu-ykb76i",
		"xn--rlsup-x0a", "r1seup", "xn--seup-9mb1013b",
		"xn--seup-ucc61q", "xn--reup-5pa433c", "xn--rsep-5pa09k",
		"xn--reup-0ya06r", "xn--risp-n53a95i", "xn--riep-rdc72y",
		"xn--riup-cpa79d", "xn--rsup-cpaw", "xn--rieu-ykb380b",
		"xn--rsup-5pa341c", "xn--riup-hwa76s", "xn--isep-zcc7002b",
		"xn--reup-g9d1689v", "xn--rseup-k7a", "xn--isup-w4d1249a",
		"xn--rsup-cpa6115w", "xn--riup-cpa76v", "xn--rseu-ykb03h",
		"xn--risp-dva498d", "xn--rise-ond60f", "xn--rsup-hwa231b",
		"xn--ieup-zcc5a", "xn--risp-d7d7f", "xn--rsup-6va006z",
		"xn--isep-xnb257a", "xn--risu-8vd87a", "xn--riup-mpa55v",
		"xn--ieup-l5a64o", "xn--rlsup-dlf", "xn--rise-esa351c",
		"xn--rsup-58b57y", "xn--risu-esa4n", "xn--isup-ucc84z",
		"xn--risp-ipa2244b", "xn--risp-dva45c", "xn--riup-6va919z",
		"xn--r1eup-ljy", "xn--riup-rdc045a", "xn--rsep-xnb120a",
		"xn--r1eup-769c", "xn--reup-l5a947a", "xn--risp-7va0414b",
		"xn--seup-kza337b", "xn--risu-x4d1e", "xn--riep-xnb859a",
		"xn--rseu-lza383b", "xn--rise-esa1034b", "xn--riup-mpa58d",
		"xn--rise-esa97d", "xn--isep-xod2900b", "xn--rsup-bnb520a",
		"xn--iseu-eod39d", "xn--risu-j6d26e", "xn--isup-cpa92v",
		"xn--reup-0ya0z", "xn--isep-zcc527b", "xn--risu-dva132b",
		"xn--iseu-esa52u", "xn--risu-esa282b", "xn--iseu-8vd36a",
		"xn--r1seu-0va", "xn--rlsup-8za", "xn--risu-dva165b",
		"xn--iseu-esa530c", "xn--risp-xnb99z", "xn--risu-nva945b",
		"xn--riup-wva390c", "xn--rseu-lza350b", "xn--risp-dpa641e",
		"xn--reup-5pa23v", "xn--risp-77a3116b", "xn--risp-nmd0h",
		"xn--risu-esa0l", "xn--reup-58b4m", "xn--rsep-77a590z",
		"xn--isep-ucc56y", "xn--rieu-j6d7783d", "xn--rise-yod02a",
		"xn--rsup-5md54g", "xn--risp-iwa60c", "xn--rsep-bnb860a",
		"xn--isup-z63aw3i", "xn--rlsep-vde", "xn--riep-d7d040v",
		"xn--rsep-58b448b", "xn--rsep-58b87z", "xn--rsup-5pa7w",
		"xn--isep-xnb81h", "xn--risp-dpe66i", "xn--r1sup-dlf",
		"xn--rsup-hpaq", "xn--rsup-mpa024b", "xn--rieu-m5a61d",
		"xn--risp-x4d6h", "xn--isep-ucc19q", "xn--riup-6va97s",
		"xn--rsep-shf558s", "xn--rsup-5md95e", "xn--ieup-k4d9r",
		"xn--risp-npa862c", "xn--risp-iwa648d", "xn--riup-cva711c",
		"xn--isep-k4d9739a", "xn--rise-eodx", "xn--riep-shf6791d",
		"xn--rsep-vpa20l", "xn--iseu-zcc42y", "xn--iseu-ucc03y",
		"xn--rsup-hwa0i", "xn--riup-w4d5q", "xn--rsup-mpa40u",
		"xn--risu-ipa774b", "xn--rise-ykb168c", "xn--rsup-5pa1t",
		"xn--reup-lza6w", "xn--isup-fdc42z", "xn--isep-xnb4703b",
		"xn--rsep-58b448b", "xn--risp-77a583b", "xn--rsep-5md26f",
		"xn--isep-fdc707b", "xn--rsep-0ya2204b", "xn--rseu-j6d657u",
		"xn--risu-ipa5j", "xn--rleup-ljy", "xn--rseu-0ya39f",
		"xn--risu-ipa708b", "xn--ieup-ucc197x", "xn--riep-77a6128d",
		"xn--rseu-j6d4899v", "xn--rsup-vpa783b", "xn--risu-iwa549b",
		"xn--isup-mmd76e", "xn--r1sup-1we", "xn--rsep-vpa852c",
		"xn--seup-upa31v", "xn--risu-xva969b", "xn--isup-fdc83x",
		"xn--risu-7va591b", "xn--rise-eod11f", "xn--rsup-5pa959c",
		"xn--risp-7va858d", "xn--risp-nva06i", "xn--reup-0hu9706j",
		"xn--seup-4md7610b", "xn--rieu-8vd705v", "xn--reup-vpa6898d",
		"xn--risp-x4d28q", "xn--seup-zcc664x", "xn--rsup-cva64i",
		"xn--isup-zcc24z", "xn--risp-xva84i", "xn--risp-xva032b",
		"xn--rsup-mpaw", "xn--rsup-cpa888z", "xn--r1sup-psa",
		"xn--isup-mva17s", "xn--rsep-vpa484b", "xn--risu-7va525b",
		"xn--isup-fdc455a", "xn--rsup-hpa624b", "xn--rise-ykb9j",
		"xn--risp-dpe8458a", "xn--r1seu-nme", "xn--risu-esa8j",
		"xn--seup-kza50r", "xn--rsep-0ya81h", "xn--rsep-xod3270w",
		"xn--risu-nmd71b", "xn--rlsup-p51b", "xn--riep-nnd868v",
		"xn--rise-eod9500b", "xn--reup-l5a51e", "xn--risp-nnd55m",
		"xn--rsup-5md7852b", "xn--reup-lza8758d", "xn--rise-8vd9279a",
		"xn--rsup-wva4k", "xn--risu-7va11h", "xn--rseu-esa61a",
		"xn--riup-mmd119v", "xn--rlsup-wza", "xn--risp-ipa00f",
		"xn--rsup-0ya177b", "xn--rise-esa79j", "xn--reup-vpa44v",
		"xn--risp-xod10m", "xn--r1seu-7ce", "xn--iseu-esa70u",
		"xn--risu-nmd7e", "xn--rsup-bnb979a", "xn--ieup-l5a25o",
		"xn--rsep-xod594v", "xn--rsep-shf558s", "xn--iseu-fdc44q",
		"xn--risu-8vd6022b", "xn--rsep-5md82y", "xn--risp-xva671b",
		"xn--ieup-ucc3956d", "xn--riup-rdc01z", "xn--rsup-lza365c",
		"xn--risu-nva912b", "xn--rise-esa973b", "xn--ieup-ucc35z",
		"xn--risu-eod6352b", "xn--reup-vpa643c", "xn--risp-shf3789a",
		"xn--ieup-fdcn", "xn--rise-ynb438a", "xn--seup-9mb977a",
		"xn--riup-m5a013c", "xn--rsup-lza979a", "xn--risp-xva2514b",
		"xn--isep-77a06n", "xn--rsup-cva2p", "xn--rsep-0ya478b",
		"xn--isup-hpa9744b", "xn--rise-87a254b", "xn--risp-ipa041e",
		"xn--reup-0ya2068d", "xn--rsup-mu5ay775g", "xn--risp-xod17f",
		"xn--rise-ond47a", "xn--rsup-58b70r", "xn--riep-nnd0764d",
		"xn--rsup-5pa3u", "xn--reup-bnb4267d", "xn--isep-77a24n",
		"xn--risp-ipa041e", "xn--risp-nva4614b", "xn--r1sup-wza",
		"xn--isup-zcc275a", "xn--rsup-wva216z", "xn--isep-k4d79q",
		"xn--riep-m5a364d", "xn--risp-7va63i", "xn--riup-m5a086b",
		"xn--risp-nnd52g", "xn--isup-wva95s", "xn--rlsep-bjg",
		"xn--reup-bnb28h", "xn--risp-dpa254b", "xn--isup-zcc88p",
		"xn--rsup-c9d6689v", "xn--rsep-lza67a", "xn--seup-upa141c",
		"xn--risu-dpa542c", "xn--risp-x4d4639a", "xn--rise-j6d51q",
		"xn--rieu-j6d570v", "xn--ieup-z63ap82i", "xn--isup-wva788b",
		"xn--rsup-vpa143c", "xn--lseup-jrc", "xn--isep-zcc527b",
		"xn--isep-ucc137b", "xn--rsep-77a32m", "xn--ieup-fdc9656d",
		"xn--r1sup-l0a", "xn--rsep-77a50d", "xn--rseu-esa2z",
		"xn--isep-77a073b", "xn--riep-m5a3i", "xn--rseu-58b18r",
		"xn--rsup-58b3064b", "xn--risp-nnd93e", "xn--rieu-8vd96c",
		"xn--isep-d7d8139a", "xn--isep-fdc14y", "xn--rseu-esa403b",
		"xn--rieu-eod99f", "xn--seup-fdc6762w", "xn--riep-rdc99p",
		"xn--rsup-mpa23b", "xn--iseu-esa10u", "xn--isup-6va7914b",
		"xn--rsup-cpe6329v", "xn--riup-cpe723u", "xn--rsup-mva43i",
		"xn--risp-nmd2310b", "xn--rsep-vpa6044b", "xn--rise-esa910e",
		"xn--isep-ucc137b", "xn--riup-mpa5910a", "xn--iseu-ykb778a",
		"xn--riup-hpa363c", "xn--rise-ykb122a", "xn--seup-4pa59u",
		"xn--1seup-jrc", "xn--rsup-6va6l", "xn--seup-zcc4962w",
		"xn--rsep-5pa27e", "xn--risp-iwa089b", "xn--rseu-vpa9h",
		"xn--lseup-mwe", "xn--rieu-m5a008a", "xn--riep-77a675b",
		"xn--isup-cpa35v", "xn--ieup-g9d0329a", "xn--rieu-m5a031b",
		"xn--risp-7va812b", "xn--risp-iwa42i", "xn--rsep-bnb0a",
		"xn--reup-g9t4q", "xn--rsup-58b196a", "xn--reup-qdc97p",
		"xn--isep-shf2547a", "xn--rsup-0ya9946b", "xn--riup-cpe9213d",
		"xn--rseu-bnb168a", "xn--risp-iwa602b", "xn--rsep-5md0110b",
		"xn--risp-ipa005b", "xn--rsup-mmdw", "xn--r1sup-d2e",
		"xn--seup-fdc844x", "xn--riep-h9d19o", "xn--isup-hwa53s",
		"xn--risp-ipa81l", "xn--rsup-cpa89k", "xn--rseu-5pa7g",
		"xn--reup-l5a3344w", "xn--riep-rdc35q", "xn--ieup-l5a07o",
		"xn--isep-77a2963b", "xn--rlsup-dsa", "xn--isup-hpa74v",
		"xn--rsup-mpa80b", "xn--riep-xnb65h", "xn--risu-dva74h",
		"xn--risp-ipa472c", "xn--rsep-77a936a", "xn--rsup-hwa4k",
		"xn--riup-mpa7998d", "xn--isup-zcc65x", "xn--risp-xnb767a",
		"xn--isup-cva77s", "xn--rsep-n53ap096g", "xn--rieu-rdc29x",
		"xn--isup-mva1224b", "xn--risp-xod58d", "xn--r1sup-39d",
		"xn--reup-lza63r", "xn--rsup-hwa885z", "xn--isup-hwa368b",
		"xn--reup-g9ty68s", "xn--rsep-lza617d", "xn--rsup-0ya759b",
		"r-iseup", "ri-seup", "ris-eup",
		"rise-up", "riseu-p", "ri9seup",
		"rise7up", "riqseup", "ris3eup",
		"riseuhp", "risqeup", "rise4up",
		"riserup", "risyeup", "rijseup",
		"risehup", "risaeup", "risweup",
		"r9iseup", "risreup", "rixseup",
		"rizseup", "riseiup", "r8iseup",
		"riseyup", "riseeup", "rikseup",
		"riaseup", "rise3up", "riuseup",
		"risedup", "risdeup", "riyseup",
		"ridseup", "risewup", "riseuzp",
		"risseup", "riwseup", "riseujp",
		"risxeup", "rioseup", "risejup",
		"risesup", "riseuyp", "ri8seup",
		"roiseup", "risezup", "riszeup",
		"riseu8p", "rieseup", "ris4eup",
		"ruiseup", "rise8up", "riseuip",
		"rkiseup", "rjiseup", "riseu7p",
		"rieup", "risup", "iseup",
		"rseup", "risep", "riseu",
		"riseuup", "rriseup", "riiseup",
		"eiseup", "riaeup", "risrup",
		"rise8p", "riseyp", "rideup",
		"ris3up", "riswup", "rixeup",
		"rizeup", "ruseup", "r8seup",
		"risezp", "tiseup", "rise7p",
		"riszup", "ris4up", "rieeup",
		"fiseup", "rissup", "r9seup",
		"5iseup", "rjseup", "riseum",
		"riyeup", "riseip", "roseup",
		"riseuo", "riseul", "risejp",
		"diseup", "4iseup", "risehp",
		"irseup", "rsieup", "riesup",
		"risuep", "risepu", "risiup",
		"riseop", "riseap", "risoup",
		"reseup", "riseupcom",
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Whitelist:  whitelist,
		Suspicious: suspicious,
	}
}
