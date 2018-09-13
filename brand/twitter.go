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

// Twitter brand properties.
func Twitter() *Brand {
	name := "twitter"
	original := []string{"twitter"}
	whitelist := []string{"twitter.com", "ads-twitter.com"}
	suspicious := []string{
		"twittera", "twitterb",
		"twitterc", "twitterd", "twittere",
		"twitterf", "twitterg", "twitterh",
		"twitteri", "twitterj", "twitterk",
		"twitterl", "twitterm", "twittern",
		"twittero", "twitterp", "twitterq",
		"twitterr", "twitters", "twittert",
		"twitteru", "twitterv", "twitterw",
		"twitterx", "twittery", "twitterz",
		"uwitter", "vwitter", "pwitter",
		"dwitter", "4witter", "tvitter",
		"tuitter", "tsitter", "tgitter",
		"t7itter", "twhtter", "twktter",
		"twmtter", "twatter", "twytter",
		"twiuter", "twivter", "twipter",
		"twidter", "twi4ter", "twituer",
		"twitver", "twitper", "twitder",
		"twit4er", "twittdr", "twittgr",
		"twittar", "twittmr", "twittur",
		"twittes", "twittep", "twittev",
		"twittez", "twitteb", "twitte2",
		"xn--witr-p6ac961c", "xn--titte-rwe74n", "xn--twite-6rc162a",
		"xn--wer-2ta2yca", "xn--twter-8ye25410a", "xn--wter-5pa264bca",
		"xn--twitt-mza14w", "xn--twitt-mza84w", "xn--twitt-3we9085b",
		"xn--vvitter-dqf", "xn--tittr-y0a292e", "xn--witr-ipa642cca",
		"xn--wtter-n4a706c", "xn--twttr-xza9n", "xn--wer-7zaba0383a",
		"xn--twter-dta091c", "xn--twir-s6aa272b", "xn--wittr-49d4h",
		"xn--twter-5nc624a", "xn--wer-zma737acaa", "xn--twlttr-mva",
		"xn--ttter-y3a546e", "xn--wie-fxb08tcaa", "xn--wir-7edca9358b",
		"xn--wir-lra661bcaa", "xn--twitr-9db372b", "xn--twir-s6da9358b",
		"xn--twitr-49d99g", "xn--twir-nva199ba", "xn--twitt-fsa36z",
		"xn--wittr-esa622c", "xn--twtte-0sa389c", "xn--twttr-0sa83a",
		"xn--twir-iwa991ba", "xn--twttr-qsa37b", "xn--wer-pdb21wcaa",
		"xn--tittr-2we81a", "xn--witr-nva87bc", "xn--twitr-9za097c",
		"xn--wtter-cta42f", "xn--twttr-ksa24470a", "xn--wittr-esa36f",
		"xn--wer-lyc79dca", "xn--wter-podb43320a", "xn--wittr-49d69g",
		"xn--tvitter-vvg", "xn--vvitter-5jg", "xn--witr-p6ac683c",
		"xn--wer-7edba586y", "xn--twttr-o4a274d", "xn--witer-6db162c",
		"xn--wittr-6db983c", "xn--twttr-y0a99070a", "xn--witr-7va45bc",
		"xn--twtte-rwe45510a", "xn--tittr-9za697e", "xn--twttr-ksa80c",
		"xn--wie-fxb05ncaa", "xn--wie-jxb54ncaa", "xn--twir-s6aa4863c",
		"xn--wer-6ub74ocaa", "xn--twitr-9ye1k", "xn--twitr-9db121d",
		"xn--tw1ttr-fhg", "xn--wer-jua704acaa", "xn--wier-r6a125bda",
		"xn--wter-bnb211aca", "xn--wtter-zsa380d", "xn--witte-lde5066b",
		"xn--tvvier-srfa", "xn--witr-iwa259bca", "xn--twitr-beb861c",
		"xn--twitr-y0a71c", "xn--witr-npa042cca", "xn--twttr-xza19u",
		"xn--twite-pde29e", "xn--titte-6rc862c", "xn--twitt-6rc271a",
		"xn--wtter-lde63320a", "xn--tier-r6da44m", "xn--twttr-dta680c",
		"xn--tittr-m0a099c", "xn--twter-9ye15410a", "xn--twttr-dta01a",
		"xn--twter-o4a727b", "xn--titer-9db562e", "xn--wittr-6ye3458b",
		"tw1tter", "xn--twer-6pa022ca", "xn--wer-rma747acaa",
		"xn--twer-cnb078aa", "xn--witr-p6dc1k", "xn--twttr-xwb3292c",
		"xn--twitr-esa76f", "xn--wtter-zsa102c", "xn--twitr-y0a61c",
		"xn--twttr-o4a025c", "xn--twlter-k1e", "xn--wittr-xza008c",
		"xn--vvitter-5qb", "xn--twltte-855b", "xn--twlter-rrf",
		"xn--wittr-lza239b", "xn--wier-r6aa552c", "xn--wlter-6yec",
		"xn--twttr-9za928b", "xn--twtte-i9x846f", "xn--twer-mza41aa",
		"xn--twier-9yea", "xn--tvviter-iqf", "xn--ier-7edba02r",
		"xn--twtte-0sa2961c", "xn--t1tter-wxf", "xn--witr-npa484bca",
		"xn--titte-irc692c", "xn--titte-orc982c", "xn--witr-x4d8ec",
		"xn--twitt-mza96w", "tvv1tter", "xn--twitr-beb4863c",
		"xn--wier-podc59f", "xn--wir-tzcca22h", "xn--twttr-elf98500a",
		"xn--twir-soda22h", "xn--twite-irc982a", "xn--twite-orc282a",
		"xn--vviter-orfd", "xn--twttr-esa99x", "xn--tittr-e2e39p",
		"xn--wir-5xc61eca", "xn--witr-xva022bca", "xn--tw1ttr-mye",
		"xn--wier-r6a125bca", "xn--tittr-qsa401f", "xn--tier-roda200a",
		"xn--wier-p6ac852c", "xn--witer-ldec", "xn--twitr-beb021d",
		"xn--twitt-b0a31w", "xn--twitt-b0a02w", "xn--titte-j4e5055b",
		"xn--w1ter-ldec", "xn--tw1ttr-mva", "xn--witr-ipa63ec",
		"xn--tvvittr-dhg", "xn--twite-9db11r", "xn--witr-nva232bca",
		"xn--tittr-m0a287e", "xn--twitr-9za44c", "xn--wier-p6ac673b",
		"xn--twttr-qsal", "xn--twie-ycc08qa", "xn--twie-3cc47qa",
		"xn--tw1ter-rrf", "xn--ttter-j4e074y", "xn--twer-6pa01ea",
		"xn--wir-1ra641bcaa", "xn--twttr-0sa312d", "xn--titte-r1f9814b",
		"xn--twter-dta981c", "xn--twlter-rkb", "xn--wittr-y0a31c",
		"xn--twir-s6da57f", "xn--ier-tzcba01i", "xn--twitr-qsa512c",
		"xn--wittr-qsa490d", "xn--twttr-z3a256b", "xn--witer-6db262c",
		"xn--twttr-0sa01a", "xn--wier-p6a077aa", "xn--wter-p6ab79560a",
		"xn--twitr-y0a267c", "xn--wer-6ub77ucaa", "xn--witer-6db973b",
		"xn--wir-kma657acaa", "xn--tittr-lza230d", "xn--twtte-dta69y",
		"xn--twtte-dta30z", "xn--itter-lde31i", "xn--twitr-ksa401d",
		"xn--wter-5pa80ec", "xn--twitt-o69ao6j", "xn--twitr-ksa222c",
		"xn--witer-lded", "xn--witer-8db652c", "xn--wtter-qbe17g",
		"xn--twtte-o4a6121c", "xn--twttr-o4a6304c", "xn--wir-kma11dca",
		"xn--twitt-rsa12z", "xn--twitt-rsa82z", "xn--twttr-ksa12n",
		"xn--wtter-wwb494b", "xn--twttr-m0a7434a", "xn--twir-s6aa021d",
		"xn--twttr-rbe0059b", "xn--twttr-qsa731c", "xn--itter-6ye74m",
		"xn--twir-nva18ba", "xn--tvvittr-gdh", "xn--twter-z3a39a",
		"xn--w1er-podca", "xn--twter-z3a657b", "xn--twitr-2we3g",
		"xn--twttr-xza5734a", "xn--ier-7zaba374c", "xn--witr-iwa24bc",
		"xn--tittr-e2e71l", "xn--twttr-y0a108b", "xn--twttr-m0a36u",
		"xn--twitr-9za809b", "xn--wer-pdb242acaa", "xn--tvvtter-9ya",
		"xn--twter-dta170d", "xn--wter-5md43fca", "xn--twter-rbe1f",
		"xn--titer-8db662e", "xn--wir-tra625acaa", "xn--twter-0sa402c",
		"xn--twitt-z0a7631c", "xn--twtte-xwb82j", "xn--twtte-xwb12j",
		"xn--tvviter-9jg", "xn--twitr-bze0k", "xn--twer-r6aa64p",
		"xn--tittr-lza843e", "xn--iter-p6ab474c", "xn--twttr-ksa6764a",
		"xn--wittr-9za409b", "xn--wlter-ldec", "xn--wir-7edca57f",
		"xn--twite-irc80u", "xn--twitr-ksa95f", "xn--witte-lde69e",
		"xn--tw1ttr-7of", "xn--witr-p6ac372b", "xn--twitr-xza129b",
		"xn--wer-jua2wca", "xn--twie-s6aa28q", "xn--tvvier-skba",
		"xn--t1tter-3eh", "xn--wtter-6db103b", "xn--tvvtter-web",
		"xn--twie-jdc65qa", "xn--twitr-esa211d", "xn--twter-5nc44v",
		"xn--twitr-49d00h", "xn--vvier-6dbda", "xn--twite-9db09q",
		"xn--wite-3cc72yca", "xn--twir-x4d1fa", "xn--twttr-e2e42310a",
		"xn--ttter-4nc146a", "xn--twitr-m0a577c", "xn--twitt-b0a14w",
		"xn--wtter-cta960d", "xn--twitr-xza95c", "xn--twlter-srf",
		"xn--tvvitter-q0e", "xn--wer-lyc7cca", "xn--twttr-5nc633a",
		"xn--twlttr-fva", "xn--tittr-j4e5238b", "xn--tvvtter-qza",
		"xn--titter-w1fa", "xn--twer-r6aa69560a", "xn--twitr-lza27c",
		"xn--twttr-o4a426b", "xn--wie-7edca9175b", "xn--tvvitte-lld",
		"xn--tvvitte-tld", "xn--tvvtter-4gd", "xn--twer-1ya848ba",
		"xn--wittr-ksa912c", "xn--witr-7va802bca", "xn--tvitter-sof",
		"xn--itter-lde500a", "xn--wter-lza660bca", "xn--wtter-4nc324a",
		"xn--twite-pde1066b", "xn--titer-ode200a", "xn--twir-npa784ba",
		"xn--witte-irc682a", "xn--witte-orc972a", "xn--twttr-9za1634a",
		"xn--twttr-qsa53470a", "xn--twite-ode2066b", "xn--wter-vpa02ec",
		"xn--tw1ttr-73a", "xn--twir-npa342ca", "xn--twie-jdc21ya",
		"xn--wter-p6ab67e", "xn--wtter-qbe9e", "xn--wir-kma683bcaa",
		"xn--wite-podc2066b", "xn--twttr-49dz", "xn--tw1tte-u6c",
		"xn--twttr-lza758b", "xn--wittr-lde16o", "xn--witer-6yed",
		"xn--twttr-0sa2y", "xn--twitt-lsa65z", "xn--twitr-49d8h",
		"xn--witer-6db083b", "xn--wir-cma693bcaa", "xn--tvviter-brb",
		"xn--twttr-dta991d", "xn--twite-rwe7h", "xn--twttr-i9x850h",
		"xn--twite-irc70u", "xn--twite-orc00u", "xn--wer-tzcba765z",
		"xn--tvvittr-3mf", "xn--twitt-59d69f", "xn--twitt-irc251b",
		"xn--twitt-orc541b", "xn--wtter-zsa83f", "xn--ttter-4nc759b",
		"xn--ttter-cta306e", "xn--twitt-yza55w", "xn--twttr-qsaz",
		"xn--tvitter-uvg", "xn--tvvittr-i8a", "xn--wier-podca",
		"xn--twier-9db773b", "xn--twtte-5nc0l", "xn--twlttr-m4a",
		"xn--twitt-rsa94z", "xn--twttr-z3a845c", "xn--tittr-esa623d",
		"xn--wltter-orf", "xn--twter-dta62f", "xn--tvvittr-ws4c",
		"xn--wer-zma29cca", "xn--witr-podc32h", "xn--tittr-49d41j",
		"xn--twter-dta270d", "xn--wir-gma193bcaa", "xn--twir-dpa994ba",
		"xn--twitt-fsa800d", "xn--witte-6db951c", "xn--twttr-lza1s",
		"xn--twttr-elf328x", "xn--twitt-n0a1831c", "xn--twitr-pde22h",
		"xn--iter-p6ab662e", "xn--wer-7zaba64p", "xn--twitt-lsa53z",
		"xn--twitt-lsa82z", "xn--tittr-ksa536e", "xn--titer-nde300a",
		"xn--ttter-zsa103d", "xn--twtte-z3a4421c", "xn--twttr-ksa29x",
		"xn--tw1tte-15c", "xn--tw1tte-85c", "xn--twitr-m0a398b",
		"xn--twttr-49d73420a", "xn--twir-ipa93ea", "xn--witr-p6dc0458b",
		"xn--twttr-rbe48f", "xn--wer-tzcba33320a", "xn--twttr-rbe63p",
		"xn--twitt-f2e2955b", "xn--ttter-n4a716e", "xn--tw1ter-srf",
		"xn--twer-68b25za", "xn--twttr-xza02k", "xn--twie-ycc63ya",
		"xn--twie-3cc03ya", "xn--twir-nmd3ga", "xn--twttr-qsa10c",
		"xn--tltter-i0g", "xn--ttter-wwb820d", "xn--twlttr-mye",
		"xn--twitr-xza308c", "xn--titer-9ye44m", "xn--witte-6rc852a",
		"xn--twttr-dta289c", "xn--w1er-p6aca", "xn--witr-podc85o",
		"xn--twlttr-04a", "xn--ttter-j4e63210a", "xn--itter-lde92v",
		"xn--ttter-qbe98i", "xn--wtter-y3a357b", "xn--twie-s6aa4680c",
		"xn--tw1er-9dba", "xn--tvvitte-imd", "xn--twitt-yza007c",
		"xn--twitr-ode60f", "xn--tvvittr-gog", "xn--twttr-lza9834a",
		"xn--twttr-2we14510a", "xn--tvvittr-wya", "xn--twttr-esa82n",
		"xn--wite-ycc77qca", "xn--wite-3cc17qca", "xn--twttr-z3a005d",
		"xn--twter-0sa04f", "xn--wter-58b49rca", "xn--tw1ter-k1e",
		"xn--tw1ter-rkb", "xn--wie-tzcca1066b", "xn--witer-ode29f",
		"xn--twitr-qsa790d", "xn--twer-r6aa892b", "xn--twir-7va769ba",
		"xn--tvvitter-p0e", "xn--twir-nmd95fa", "xn--twite-9ye0275b",
		"xn--twier-odea", "xn--wittr-lza96c", "xn--titer-nde72v",
		"xn--wie-ldd5eca", "xn--ttter-y3a358c", "xn--twttr-xwb704b",
		"xn--witr-nva889bca", "xn--twer-wpa22ea", "xn--twter-9db57e",
		"xn--tier-r6aa374c", "xn--vviter-h1ed", "xn--wittr-lde62h",
		"xn--twter-nde43320a", "xn--twitr-9ye0458b", "xn--wittr-qsa94f",
		"xn--titer-ode01i", "xn--twite-beb01r", "xn--tittr-2we01s",
		"xn--wite-ycc33yca", "xn--iter-p6ab097d", "xn--twier-ode69f",
		"xn--wittr-lde5249b", "xn--twite-9db38q", "xn--twitt-irc6171c",
		"xn--twitt-orc9071c", "xn--twite-bze9175b", "xn--tw1tte-855b",
		"xn--wittr-y0a857c", "xn--twtte-rwe887y", "xn--twer-r6da15410a",
		"xn--twitt-59d5076b", "xn--tittr-9za400d", "xn--tvvittr-f9a",
		"xn--wite-p6ac38q", "xn--wite-p6ac09q", "xn--twttr-0sa001c",
		"xn--ier-7zaba987d", "xn--twitt-z0a856c", "xn--wer-7zaba892b",
		"xn--iter-p6db54m", "xn--wite-p6ac11r", "xn--ier-7zaba562e",
		"xn--twttr-dta2y", "xn--twlttr-t3a", "xn--twitr-pde50f",
		"xn--twttr-dta8944c", "xn--twtte-z3a535c", "xn--twttr-y0a7j",
		"xn--twttr-e2e855y", "xn--witr-nmd65fca", "xn--twitt-rwen",
		"xn--twitt-6rc8861c", "xn--titer-8ye54m", "xn--witer-lde99f",
		"xn--twitr-y0a088b", "xn--twter-9db64p", "xn--twttr-9za3p",
		"xn--twite-beb98q", "xn--twite-beb28q", "xn--wittr-ksa65f",
		"xn--twltte-uof", "xn--twttr-xwb953c", "xn--twitt-mza3241c",
		"xn--twtte-irc92t", "xn--twtte-orc22t", "xn--twter-xwb416a",
		"xn--twitt-irc002a", "xn--witr-nmd0gc", "xn--wir-tzcca1249b",
		"xn--wite-p6dc0275b", "xn--twter-ode33320a", "xn--tier-r6da8s",
		"xn--titer-8ye9s", "xn--twttr-0sa2154c", "xn--tw1ttr-f5a",
		"xn--tvitter-tof", "xn--twitt-yza9041c", "xn--twer-6pa464ba",
		"xn--wer-7edba15410a", "xn--twtte-o4a705c", "xn--wittr-6db8863c",
		"xn--wtter-4nc14v", "xn--twttr-m0a29j", "xn--tw1ttr-tva",
		"xn--twir-dva39ba", "xn--twttr-xza14170a", "xn--wittr-m0a72c",
		"xn--wir-dra10bca", "xn--iter-podb72v", "xn--wter-podb865z",
		"xn--twttr-0sa42a", "xn--twttr-ksa441c", "xn--twir-7va112ba",
		"xn--twitt-n0a276c", "xn--twitt-rwe97g", "xn--wter-lza228bca",
		"xn--twitr-pde1249b", "xn--twitr-beb583c", "xn--tltter-3eh",
		"xn--wite-jdc35qca", "xn--wir-dra645acaa", "xn--twter-xwb794b",
		"xn--titer-9ye02r", "xn--wir-7zaca272b", "xn--wier-p6ac952c",
		"xn--tvvier-l1ea", "xn--twitr-lza639b", "xn--tvvittr-17a",
		"xn--wier-r6a567ada", "xn--t1tter-i0g", "xn--twer-roda33320a",
		"xn--titer-8db097d", "xn--twter-o4a46a", "xn--tittr-49d03w",
		"xn--witr-dpa694bca", "xn--wer-2ta724acaa", "xn--wir-1ra615acaa",
		"xn--twttr-lza50v", "xn--wittr-9za14c", "xn--wite-p6ac5680c",
		"xn--wier-podc69f", "xn--wtter-n4a527b", "xn--wier-r6a567aca",
		"xn--twttr-dta42a", "xn--wir-7edca0k", "xn--wir-4qa655acaa",
		"xn--witr-7va469bca", "xn--twite-6rc08t", "xn--witte-6db41r",
		"xn--twitt-irc40t", "xn--twitt-orc79s", "xn--titer-ode62v",
		"xn--twttr-5nc04u", "xn--twite-ode39e", "xn--twite-irc092a",
		"xn--twite-orc382a", "xn--tittr-r1f9007b", "xn--wir-7zaca861c",
		"xn--tittr-ksa913d", "xn--twttr-m0a9n", "xn--twltte-u6c",
		"xn--twtte-dta8761c", "xn--twter-o4a827b", "xn--wir-gma61dca",
		"xn--ttter-cta970f", "xn--twlter-skb", "xn--twir-s6aa861c",
		"xn--wtter-wwb216a", "xn--twir-7va75ba", "xn--wter-5md8dc",
		"xn--wittr-9za687c", "xn--wie-vxb03ncaa", "xn--wier-p6a625ba",
		"xn--wir-7zaca021d", "xn--twitr-esa66f", "xn--twttr-qsa9664a",
		"xn--tvvitte-xgg", "xn--twttr-xza348b", "xn--twttr-ksas",
		"xn--twitr-esa032c", "xn--twtter-4db", "xn--twler-9yea",
		"xn--witte-rwe3h", "xn--ttter-zsa390f", "xn--tittr-elf28d",
		"xn--itter-6ye1t", "xn--tittr-m0a603e", "xn--twir-npa33ea",
		"xn--twttr-9za5m", "xn--twer-roda765z", "xn--twter-0sa14f",
		"xn--iter-podb300a", "xn--twitr-y0a167c", "xn--wter-lza21ac",
		"xn--twer-wpa674ba", "xn--twitt-yza72w", "xn--twitt-yza43w",
		"xn--tittr-y0a867e", "xn--twitr-m0a03c", "xn--witr-npa03ec",
		"xn--witr-dva442bca", "xn--titer-8db474c", "xn--wir-dra671bcaa",
		"xn--twttr-dta439d", "xn--twlttr-fhg", "xn--twir-nva532ba",
		"xn--wier-p6aca", "xn--wittr-qsa212c", "xn--twitr-m0a498b",
		"xn--twter-ode765z", "xn--twie-o4d2ga", "xn--twtte-6rc3381a",
		"xn--twttr-esaz", "xn--wter-bnb868aca", "xn--wler-p6dca",
		"xn--twitr-xza408c", "xn--wtter-6db94p", "xn--twite-6rc262a",
		"xn--twttr-0sa849d", "xn--tw1ttr-fvf", "xn--twitt-6rc67s",
		"xn--twttr-y0a5m", "xn--wir-5xc6eca", "xn--twir-xva979ba",
		"xn--witte-6db39q", "xn--witte-6db68q", "xn--twtte-6rc10t",
		"xn--twitr-m0a13c", "xn--tw1tte-uof", "xn--twttr-m0a31170a",
		"xn--twter-8db67e", "xn--tvviter-hqf", "xn--witr-iwa691bca",
		"xn--twitt-lsa100d", "xn--witer-8db473b", "xn--twttr-0sa6z",
		"xn--tittr-j4e16d", "xn--twitr-pde75o", "xn--ttter-4nc334c",
		"xn--twtte-5nc323a", "xn--wie-7zaca551c", "xn--twtte-rbe0866b",
		"xn--twitr-ksa501d", "xn--twitt-fsa53z", "xn--twitt-fsa24z",
		"xn--twir-soda75o", "xn--twttr-esa151c", "xn--twter-5nc524a",
		"xn--wer-7zaba57e", "xn--twir-xva96ba", "xn--tittr-xza018e",
		"xn--twitr-qsa612c", "xn--tvviter-bkg", "xn--twitr-9db961c",
		"xn--tittr-feg3356b", "xn--witr-xva66bc", "xn--twttr-9za77u",
		"xn--twitr-esa111d", "xn--tvvittr-y8a", "xn--twtte-dta42z",
		"xn--twitr-9ye67f", "xn--ttter-n4a141e", "xn--twir-soda1249b",
		"xn--twlttr-th8b", "xn--wir-lra635acaa", "xn--ttter-y3a961e",
		"xn--twer-mza860ba", "xn--wter-vpa474bca", "xn--ttter-qbe50w",
		"xn--wite-p6ac651c", "xn--titte-rwe13a", "xn--vvier-ldeda",
		"xn--titte-6rc674a", "xn--twtte-z3a09u", "xn--twter-z3a557b",
		"xn--twitr-2we4g", "xn--twttr-o4a737c", "xn--w1tter-h1e",
		"xn--wltter-h1e", "xn--vviter-okbd", "xn--ier-7edba44m",
		"xn--twer-cnb411aa", "xn--twer-68b69ra", "xn--itter-6db862e",
		"xn--twter-0sa680d", "xn--iter-podb11i", "xn--wittr-6db421d",
		"xn--ier-tzcba62v", "xn--tier-r6aa987d", "xn--twttr-49d176z",
		"xn--twter-xwb516a", "xn--w1ter-6yec", "xn--twite-9db651c",
		"xn--twttr-0sa699c", "xn--twter-dta72f", "xn--witte-6ye3275b",
		"xn--wittr-lza418c", "xn--vvier-6yeda", "xn--twtter-j2c",
		"xn--w1ter-6dbc", "xn--wittr-xza819b", "xn--twitr-lza539b",
		"xn--witer-9db552c", "xn--wtter-n4a26a", "xn--twir-ipa394ba",
		"xn--tvvitte-t51c", "xn--twttr-m0a518b", "xn--wie-jxb57tcaa",
		"xn--twer-r6aa57e", "xn--wir-tzcca75o", "xn--twttr-qsa58x",
		"xn--twttr-lza43k", "xn--wter-0ya648bca", "xn--wer-zma763bcaa",
		"xn--wite-podc39e", "xn--titter-3zea", "xn--twir-s6aa583c",
		"xn--titer-nde11i", "xn--twttr-esa78b", "xn--twitt-n0a72w",
		"xn--witte-6db8680c", "xn--twitt-n0a99v", "xn--twitt-n0a60w",
		"xn--tvvitter-p0eb", "xn--wler-podca", "xn--twter-9ye586y",
		"xn--twter-xwb694b", "xn--wir-1ra17aca", "xn--twttr-lza3p",
		"xn--wie-vxb06tcaa", "xn--tittr-esa811f", "xn--tier-r6aa562e",
		"xn--wir-tra651bcaa", "xn--wir-cma667acaa", "xn--wtter-6ye45410a",
		"xn--wer-rma20dca", "xn--ttter-zsa716e", "xn--twitr-xza85c",
		"xn--itter-6db674c", "xn--twttr-z3a567c", "xn--wittr-esa801d",
		"xn--tittr-esa246e", "xn--tvvtter-pf1z", "xn--twie-s6aa98q",
		"xn--tw1ttr-t3a", "xn--tvvtter-5gc", "xn--twite-beb4680c",
		"xn--tltter-wxf", "xn--twitr-49d7h", "xn--twie-s6da9175b",
		"xn--twir-dva300ca", "xn--twitr-ksa06f", "xn--twitr-bze9358b",
		"xn--twitr-m0a677c", "xn--twter-nde865z", "xn--twie-soda1066b",
		"xn--twtte-o4a43u", "xn--twtte-o4a14u", "xn--wter-5pa812cca",
		"xn--tvvittr-oya", "xn--twer-r6da586y", "xn--twite-beb551c",
		"xn--twitr-ode32h", "xn--wittr-6ye97f", "xn--ttter-n4a528c",
		"xn--twitt-rwe4u", "xn--wter-0ya090bca", "xn--tw1ttr-fva",
		"xn--twer-1ya290ba", "xn--wtter-lde075z", "xn--twter-z3a736c",
		"xn--ier-7edba8s", "xn--twtte-rbe17f", "xn--twter-o4a906c",
		"xn--twtte-5nc2309b", "xn--twtte-z3a26u", "xn--twtte-z3a96u",
		"xn--witr-dva09bc", "xn--titte-irc405a", "xn--titte-orc794a",
		"xn--wtter-cta781c", "xn--wter-58b05zca", "xn--twter-8db74p",
		"xn--twttr-5nc882b", "xn--twitt-lsa0171c", "xn--twter-8ye686y",
		"xn--ttter-feg4449z", "xn--twttr-q51b5432h", "xn--twitt-orc391a",
		"xn--twttr-y0a3334a", "tvvltter", "xn--wler-p6aca",
		"xn--twitt-fsa7171c", "xn--tittr-e2e1g", "xn--twer-1ya83aa",
		"xn--witr-dpa24ec", "xn--twitr-9db683c", "xn--ttter-qbe180a",
		"xn--twitt-z0a31w", "xn--tier-roda01i", "xn--witr-p6ac121d",
		"xn--twttr-z3a4604c", "xn--twter-z3a836c", "xn--titer-9ye8s",
		"xn--tw1er-odea", "xn--wier-roda29f", "xn--twter-0sa302c",
		"xn--twer-wpa232ca", "xn--wier-rod76ec", "xn--titer-9db374c",
		"xn--twer-6md0ea", "xn--twtte-irc72640a", "xn--twtte-orc02640a",
		"xn--twtte-0sa01z", "xn--twtte-0sa71z", "xn--itter-6db297d",
		"xn--wter-p6ab992b", "xn--twitt-flf7584b", "xn--witer-6dbc",
		"xn--twitr-ksa322c", "xn--ttter-feg876w", "xn--witer-6dbd",
		"xn--twitr-9za709b", "xn--twir-dpa54ea", "xn--witr-p6dc67f",
		"xn--twier-9db852c", "xn--witr-dpa252cca", "xn--wter-p6ab1383a",
		"xn--witr-ipa094bca", "xn--wier-r6aa373b", "xn--wer-7zaba69560a",
		"xn--wittr-m0a098b", "xn--ttter-wwb405d", "xn--wir-sdd6dca",
		"xn--wie-7zaca01r", "xn--twir-s6da0k", "xn--tvvtter-l95a",
		"xn--wtter-6db3383a", "xn--wir-4qa11bca", "xn--witr-dva000cca",
		"xn--twitt-b0a5931c", "xn--tittr-9za023e", "xn--twttr-esa51c",
		"xn--w1tter-okb", "xn--twter-5nc34v", "xn--twitt-rwe3468b",
		"xn--wir-cma12dca", "xn--twlttr-f5a", "xn--tw1ter-l1e",
		"xn--tw1ttr-th8b", "xn--twttr-9za60k", "xn--titte-irc028b",
		"xn--titte-orc318b", "xn--twlttr-tva", "xn--twttr-dta8w",
		"xn--twitr-9za54c", "xn--tittr-xza433e", "xn--twttr-lza55170a",
		"xn--twitr-y0a978b", "xn--twir-dva742ba", "xn--twitt-mza417c",
		"xn--twitr-9za987c", "xn--twitr-bze57f", "xn--twitt-6rc983a",
		"xn--twttr-esa94470a", "xn--wier-rod76eda", "xn--titte-rwe32s",
		"xn--tittr-lza428e", "xn--witr-p6ac5863c", "xn--wite-jdc90yca",
		"xn--twite-rwe6h", "xn--twitt-orc014a", "xn--twitt-irc714a",
		"xn--twer-r6aa0383a", "xn--twttr-5nc355a", "xn--iter-p6db12r",
		"xn--twlttr-7of", "xn--twer-6md63fa", "xn--wter-vpa032cca",
		"xn--tw1ttr-m4a", "xn--wir-gma167acaa", "xn--twttr-qsa41n",
		"xn--twie-s6aa01r", "xn--ttter-wwb217b", "xn--w1er-p6dca",
		"xn--twter-8db992b", "xn--wie-7zaca98q", "xn--wie-7zaca28q",
		"xn--wter-p6db686y", "xn--twite-9db5680c", "xn--twitt-rsa3071c",
		"xn--tier-r6da02r", "xn--titer-8ye12r", "xn--twter-z3a29a",
		"xn--wir-tra18aca", "xn--twttr-xwb426b", "xn--twtte-5nc8n",
		"xn--twlter-l1e", "xn--twter-9db0383a", "xn--wittr-6ye4k",
		"xn--twtte-0sa83z", "xn--witr-podc60f", "xn--twitr-ode2249b",
		"xn--wtter-6ye886y", "xn--twtte-6rc99540a", "xn--tier-roda62v",
		"xn--ttter-r1f0299z", "xn--wittr-lde90f", "xn--tw1ter-skb",
		"xn--witr-podc2249b", "xn--twter-0sa580d", "xn--tvvittr-v9a",
		"xn--twtte-xwb3010c", "xn--twttr-y0a94u", "xn--twite-orc10u",
		"xn--twttr-2we577y", "xn--twter-9db69560a", "xn--ier-tzcba200a",
		"xn--twitt-z0a58v", "xn--twitt-z0a29v", "xn--twttr-esa3b",
		"xn--twitr-xza229b", "xn--twttr-dta6z", "xn--twlttr-73a",
		"xn--wtter-6db87e", "xn--twitt-rsa499c", "xn--tittr-qsa213d",
		"xn--wir-tzcca50f", "xn--witer-9db373b", "xn--witer-nde39f",
		"xn--w1tter-orf", "xn--tw1er-9yea", "xn--twtte-5nc7l",
		"xn--wltter-okb", "xn--tvviter-9qb", "xn--twttr-xza7q",
		"xn--iter-p6db9s", "xn--wier-p6ac773b", "xn--wtter-y3a09a",
		"xn--twttr-ksa08b", "xn--witte-6rc67t", "xn--twttr-m0a1l",
		"xn--twir-dpa552ca", "xn--twir-soda50f", "xn--wittr-6db271c",
		"xn--wlter-6dbc", "xn--wir-4qa681bcaa", "xn--twtte-dta969c",
		"xn--twter-rbe47g", "xn--tittr-2we43n", "xn--tittr-ksa111f",
		"xn--tw1ttr-04a", "xn--wittr-2we0g", "xn--wtter-y3a536c",
		"xn--twie-soda29e", "xn--wier-p6dca", "xn--wie-tzcca29e",
		"xn--twitr-beb272b", "xn--twitr-lza818c", "xn--wter-p6db25410a",
		"xn--twttr-y0a87j", "xn--twir-xva322ba", "xn--twitt-b0a686c",
		"xn--twitr-lza718c", "xn--wir-lra19aca", "xn--wier-pod27ea",
		"xn--twitr-lza37c", "xn--twier-ode59f", "xn--twler-9dba",
		"xn--twttr-5nc2581c", "xn--wter-0ya63ac", "xn--wittr-6db672b",
		"xn--wter-p6ab74p", "xn--twitr-qsa890d", "xn--twitr-9db5863c",
		"xn--twter-o4a016c", "xn--wittr-ksa101d", "xn--twttr-9za72170a",
		"xn--witer-6yec", "xn--itter-6ye32r", "xn--twter-8db1383a",
		"xn--twer-mza428ba", "xn--twtte-o4a26u", "xn--tvvittr-gya",
		"xn--titte-6rc297b", "xn--twtte-irc1681a", "xn--twtte-orc4581a",
		"xn--tittr-qsa826e", "xn--wite-o4d9fc", "xn--twlttr-fvf",
		"xn--twitr-qsa35f", "xn--wittr-xza55c", "xn--twier-9db673b",
		"xn--tittr-y0a679c", "xn--twttr-rbe10i", "xn--tittr-49d601a",
		"xn--twtte-o69aw797h", "xn--twltte-85c", "xn--twltte-15c",
		"xn--witr-xva679bca", "xn--wtter-6db99560a", "xn--witer-lde89f",
		"xn--twir-iwa559ba", "xn--twie-s6aa551c", "xn--wer-2ta750bcaa",
		"xn--twier-9dba", "xn--wir-7zaca4863c", "xn--ttter-r1f451x",
		"xn--twttr-ksa6a", "xn--wer-jua730bcaa", "xn--titte-feg3173b",
		"xn--twter-rbe2f", "xn--twter-9db892b", "xn--twitt-6rc421b",
		"xn--titer-9db987d", "xn--twter-rbe37g", "xn--twter-8db79560a",
		"xn--twtte-xwb493b", "xn--twitr-esa922c", "xn--wittr-m0a277c",
		"xn--twitr-ode85o", "xn--twir-iwa54ba", "xn--witte-irc40u",
		"xn--witte-orc79t", "xn--tvvtter-xnf", "xn--wittr-y0a678b",
		"xn--ttter-cta782d", "xn--tittr-elf85i", "xn--twier-9db952c",
		"xn--twttr-esa3864a", "xn--wer-rma773bcaa", "xn--wie-7zaca4680c",
		"xn--wir-7zaca583c", "xn--twter-o4a56a", "xn--tvvtter-tfb",
		"xn--twite-6rc97t", "xn--twler-odea", "xn--twttr-xwb115a",
		"xn--tittr-xza810d", "xn--twitr-qsa25f", "xn--twir-ipa942ca",
		"xn--twtte-xwb94j", "t-witter", "tw-itter",
		"twi-tter", "twit-ter", "twitt-er",
		"twitte-r", "twiztter", "tw3itter",
		"twittedr", "tweitter", "twjitter",
		"twitt4er", "twirtter", "twitgter",
		"twittrer", "twuitter", "t3witter",
		"twittger", "twityter", "twitt6er",
		"twitte3r", "twi8tter", "twiktter",
		"twitrter", "tawitter", "twittder",
		"twxitter", "txwitter", "twitzter",
		"tw9itter", "twi5tter", "twiotter",
		"twitt3er", "twittezr", "twittyer",
		"twqitter", "twsitter", "twitt5er",
		"twkitter", "twittzer", "tewitter",
		"twittewr", "twittfer", "twaitter",
		"twiytter", "twit6ter", "tswitter",
		"twitte4r", "twoitter", "tqwitter",
		"twigtter", "twijtter", "twittesr",
		"t2witter", "twit5ter", "twiftter",
		"tw2itter", "twitfter", "tw8itter",
		"twi6tter", "twi9tter", "twiutter",
		"twittwer", "twittser", "twittr",
		"witter", "twitte", "twiter",
		"twtter", "titter", "twittter",
		"twiitter", "ttwitter", "twitteer",
		"twwitter", "twitrer", "ywitter",
		"t3itter", "twi6ter", "txitter",
		"twittet", "tqitter", "t2itter",
		"twjtter", "zwitter", "twittee",
		"twittef", "twittwr", "5witter",
		"twitted", "twizter", "twigter",
		"twitger", "twutter", "twit5er",
		"twitt4r", "tw9tter", "teitter",
		"6witter", "rwitter", "twitzer",
		"twit6er", "twitte5", "twotter",
		"twirter", "twitfer", "gwitter",
		"twi5ter", "twityer", "twitte4",
		"taitter", "twittsr", "twittzr",
		"tw8tter", "fwitter", "twittrr",
		"twiyter", "twifter", "twitt3r",
		"wtitter", "tiwtter", "twtiter",
		"twitetr", "twittre", "twetter",
		"twittor", "twittir", "twittercom",
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Whitelist:  whitelist,
		Suspicious: suspicious,
	}
}
