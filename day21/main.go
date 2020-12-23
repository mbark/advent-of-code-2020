package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mbark/advent-of-code-2020/util"
)

var input = `
mxc ssdszsn phc llj nhnd fqqcnm njrcfg mrczqmj zfnttf tqkfx zvq rftbr jgbk kmlkx tzdks czvphx qpbl fzlrf zmsdzh qfslcb psqc rfx vnjxs lkgln sqhvzg qzjrtl hfzqf qhtvsr zdthsvl cgflg nggbtk mxq rdkrtr djdhrn nkzqj lcrs mhrlx hpbnj blvfz djhj ssrgt lpvfv fslqkg rzcps njdb jbtlfv cncpbssj lsgqf hbdv gbhjv drtdz flndv gpjr xqgb xvjk zjbgp rjc bfnnnrn gcptp prksl xdcp pdt xtgjslz fgptl kdm njqrhcc nqvnn lxldx spnd nlqhlsn tcclbr gpqgkt jfrlp hddd bmqzgvr rfxgl fgsr hjvzcp mxzb hqd (contains eggs, fish, peanuts)
zcphr pcslhrg strpjp hbbgk nqvnn rtq jqzklv fnlk mhrlx fgsr bvlb gpqgkt fqqcnm dhsssf pdhlzg dchvb njdb rms zppxp pdt nfgr th jbz tktj rtpff xvjk fqtc ldrgj jbtlfv hzmk tqkfx nppxr qzjrtl nlqhlsn qfslcb xbvrx jhv ssrgt xqgb nbgv kmlkx chgjqc nthsf vhjgg mcl fzlrf lzvh pmhhqrk hjzkg lpvfv jgbk xsk blvfz thtlt nfck lftqn mxq flndv lsgqf njqrhcc mrczqmj nnskqnmn rjc zfnttf sqhvzg rfxgl zmsdzh rxd zbxj phc rtmfg bfnnnrn kdm dfcgd (contains peanuts)
vhjpjdr nxrgp pdj njdb fgsr zmsdzh fctmvrs dfcgd nnskqnmn cvbc rxd th fgptl fqqcnm xqgb nkzqj dkg phc qpbl dgxnc fnlk llj tktj nlqhlsn bkzbrm lzvh rfgf fgcd nppxr vxmrk mhrlx pcslhrg czvphx hfzqf dxrsgs jbz dxdrk chgjqc tqkfx nfgr snzxr frxmq hjzkg mgv lsgqf jhbnm rjc xtgjslz sxsxm gvdstsc pdt lpvfv zbxj (contains wheat, dairy)
xdcp njrcfg blvfz kdbxxzv tcclbr jhbnm rlgr njqrhcc nggbtk rntk phc zmsdzh lxldx qfslcb psqc mzrx ggrc hqd mrczqmj jp mxq sqhvzg bggmsj mcl gcptp jgbk tzdks hbbgk llj smdlg tqkfx pdt jqgm jhv vnjxs gcpks cqqcd hrnvd rjc rfgf jlrt bvlb njdb rdkrtr dxrsgs czvphx flndv nppxr gvdstsc nlqhlsn hdtr gpqgkt fgcd lzvh bvcxfp dchvb spnd jsh gbhjv vhjgg bjcvpc prbk rng zvq nthsf kjjst pdj zdthsvl pdhlzg lsgqf mgv vbkb ssrgt jxvx (contains nuts, dairy)
jqgm prksl psqc lftqn jlrt jqzklv fppgp ldrgj cgflg xsk gpqgkt rjc sfkcp zmsdzh njqrhcc nppxr tcclbr snzxr hbdv npftghk phc hbmvpmt bjcvpc czmml bkzbrm ggrc pdt blvfz spnd smdlg mbmtz gcptp hpbnj jbz lsgqf czkfv gqnxlr nhnd chgjqc rzcps jhv fnlk gcpks nlqhlsn fslqkg rsr hjvzcp cvbc rdkrtr zdthsvl nxrgp hjzkg vbkb nfgr fgcd tktj czvphx qhtvsr hcjkd qfslcb hqd fqqcnm nhzthvn sxsxm (contains fish, sesame, soy)
zmsdzh dxrsgs qrpzt tzdks bggmsj chgjqc cgflg sxsxm zppxp jbtlfv bvcxfp pdj hbmvpmt fzlrf lxldx llj dgxnc hfzqf hqzcncv gbhjv spnd phc hjzkg dhsssf fqtc xtgjslz nfck thtlt tcclbr xdcp djdhrn zjbgp nhnd kdm ksrjn mxq rng pdt dxdrk rtpff sqhvzg ldrgj kjjst rjc jp lsgqf flndv ggrc njqrhcc gpqgkt mrczqmj mxc rftbr psqc lzvh nmps (contains soy)
jfrlp hqd mxq fnlk cncpbssj nnskqnmn rzcps hbmvpmt pdt zmsdzh pfrqf zfnttf zjbgp dchvb jp mzbgp phc qpbl qdkk nqvnn rsr xbvrx sfkcp cqqcd cgflg dhsssf rlgr vnbfvkp hrnvd lsgqf ldrgj nxrgp qzjrtl czmml nthsf fqqcnm kmlkx lpm xvjk nbgv rfx fctmvrs qjhq njdb pdhlzg xsk rbjmdn tzdks spnd chgjqc bvlb jqzklv lzvh jxvx rxd kjjst jvhsj vxmrk (contains dairy, peanuts)
njqrhcc mxc mgv phc fzlrf spnd jvhsj kmlkx cldgd fgsr vnjxs ksrjn fqtc qjhq nfgr bmqzgvr pdhlzg frxmq gcptp rdjdq vbkb fgptl npftghk dxrsgs fgcd qdkk qzjrtl lpvfv hfzqf fqqcnm lzvh hpbnj hzmk jqzklv ktnlk tktj vhjpjdr llj fppgp zmsdzh hjzkg jbz cqqcd rtq hbdv njrcfg fslqkg jfrlp smdlg dhsssf qhtvsr rjc lsgqf jhv rftbr xgpdnz (contains wheat, dairy, eggs)
zmsdzh sjpzc tkhzz kdbxxzv tqkfx fslqkg rtpff nbgv prksl nxrgp pdt zfnttf qdkk rjc vnbfvkp nppxr fppgp spnd bkzbrm rftbr mzbgp llj nfgr xtgjslz lpvfv zjbgp fctmvrs zgvtn xfn nvmm rntk jhbnm frxmq strpjp cqqcd phc dfcgd nggbtk kdm sxsxm gbhjv cldgd qjhq th jlrt rdjdq mzrx hqd jsh vhjgg nlqhlsn vxmrk zcphr lzvh mcl dhsssf bggmsj kjjst fqqcnm (contains nuts, dairy)
zbxj jlrt lxldx jqgm zfnttf rdkrtr smdlg mcl fctmvrs vhjgg vbkb fqtc nlqhlsn czmml bxqpgx hpbnj fzlrf spnd pdj flndv phc pdt rzcps mzbgp gcpks ldrgj rtmfg zjbgp lzvh hjvzcp rfxgl czkfv gpqgkt bkzbrm jhbnm dxrsgs xdcp hbdv xvjk lpm lkgln gqnxlr kmlkx nfgr jbtlfv fgsr prksl jfrlp nthsf drtdz dfcgd dgxnc dhsssf rjc tkhzz tktj xgpdnz zmsdzh nhzthvn fslqkg bmqzgvr hrnvd lsgqf bjcvpc mbmtz mxc nbgv th bggmsj nqvnn nnskqnmn mxq rlgr rtpff kjjst pfrqf (contains nuts, soy)
tktj jbz rbjmdn zcphr fqqcnm djdhrn smdlg nkzqj qdkk vnbfvkp rtmfg fnlk jqgm tkhzz rdkrtr hpbnj nhzthvn rjc czkfv lftqn cldgd fgsr rfxgl lkgln bvpxc lsgqf kdm blvfz spnd nlqhlsn xvjk nmps lzvh nvmm snzxr psqc vnjxs sfkcp pdt rfgf gcptp zdthsvl mrczqmj hqzcncv pcslhrg djhj hqd bggmsj nfck hdtr bxqpgx jsh rtpff jqzklv qlkfvk phc jfrlp strpjp (contains dairy, wheat, peanuts)
hbmvpmt psqc pfrqf gvdstsc chgjqc cgflg pdhlzg gcptp nmps tktj frxmq mzbgp pdt mcl fppgp jqzklv mhrlx ssrgt tcclbr rzcps nhzthvn rjc rfxgl fqqcnm gqnxlr gpqgkt fgsr hqzcncv spnd nbgv nhnd mxc xtgjslz nfgr vbkb zjbgp jgbk nppxr rfgf dxrsgs nqvnn hfzqf gcpks jlrt rdkrtr rntk rlgr rxd kjjst flndv lxldx phc jqgm zmsdzh hrnvd zdthsvl drtdz rfx qpbl lzvh xbvrx gpjr hbdv tzdks blvfz smdlg gbhjv pmhhqrk nnskqnmn rtq prksl qfslcb ggrc nthsf bggmsj zppxp (contains eggs, fish)
nthsf pmhhqrk gpjr zmsdzh jhv zfnttf vhjpjdr rtmfg cvbc nvmm vbkb bmqzgvr lftqn prksl rjc bfnnnrn qzjrtl pdj hrnvd jbtlfv mgv kdm xtgjslz nkzqj mhrlx fgcd kmlkx qrpzt xfn rntk fppgp fqqcnm dgxnc nggbtk bjcvpc nppxr qfslcb fctmvrs pdt njqrhcc rdkrtr ldrgj rfxgl jgbk spnd phc rfx lkgln rtq xdcp jhbnm prbk jqhn hfzqf mxc rms mxzb rtpff tqkfx fslqkg jbz hbmvpmt hbdv rfgf lzvh (contains fish)
gcpks mrczqmj rtq gpjr npftghk hbbgk lxldx xtgjslz njrcfg hjvzcp kdm jp sfkcp bvcxfp xdcp nqvnn tkhzz czvphx ssdszsn xvjk zmsdzh chgjqc rdkrtr hqd qrpzt mbmtz frxmq pdt zppxp nmps hdtr dfcgd nbgv gbhjv th rsr lsgqf fqqcnm rjc cldgd flndv bjcvpc zfnttf zgvtn rms fslqkg phc rtpff hjzkg rlgr dxrsgs nvmm zcphr rxd nxrgp zdthsvl jhbnm jqhn qfslcb lftqn psqc czkfv pcslhrg fqtc xsk nnskqnmn mxq qdkk spnd lcrs ksrjn kmlkx ldrgj nfgr nlqhlsn cvbc djhj thtlt bmqzgvr lpvfv qpbl xqgb (contains dairy, soy, peanuts)
prbk thtlt fqqcnm rjc tkhzz hcjkd blvfz xvjk njrcfg nnskqnmn xfn xdcp czkxlg hzmk bvlb ssdszsn xsk mbmtz spnd zjbgp lzvh lsgqf rtq kjjst xgpdnz bxqpgx czmml gpjr ggrc ktnlk jbtlfv pmhhqrk nkzqj phc zbxj fzlrf dgxnc lcrs xbvrx rdkrtr rfxgl hqd gbhjv vbkb pdt nfck qdkk jlrt (contains soy, nuts)
qhtvsr mxc prksl jhv cvbc dkg drtdz gpqgkt jbtlfv rdkrtr qfslcb phc ssdszsn dchvb fppgp djdhrn tktj tkhzz jbz njqrhcc dhsssf lftqn mrczqmj jgbk lcrs hrnvd fqqcnm jqzklv rfx hjvzcp czkxlg pcslhrg kdbxxzv zvq spnd qzjrtl lzvh vxmrk hbmvpmt bxqpgx zmsdzh qdkk pfrqf jsh strpjp th njdb pdt sjpzc bfnnnrn lxldx nbgv hbbgk rjc (contains fish)
sfkcp njqrhcc xvjk lpm rfgf fctmvrs spnd djdhrn nhnd zmsdzh njrcfg gvdstsc zjbgp nthsf prbk njdb jqzklv hdtr dxrsgs phc thtlt nppxr mxc fqqcnm gbhjv lcrs mzrx pdhlzg frxmq hbmvpmt nggbtk rdkrtr zdthsvl smdlg nfck nhzthvn pdt dkg xqgb bmqzgvr sjpzc qfslcb hjzkg llj nqvnn rjc lzvh rtq kmlkx hzmk rfx xbvrx fgptl (contains nuts, fish, peanuts)
xgpdnz hqzcncv tkhzz lpvfv hpbnj fgsr zbxj rjc rms rntk ssdszsn tktj bvlb snzxr hrnvd xbvrx rxd rfxgl spnd mrczqmj gcpks cgflg lzvh zjbgp hzmk rng lpm rdjdq hdtr fqqcnm zmsdzh fctmvrs qzjrtl gbhjv nvmm zdthsvl blvfz xvjk lsgqf jvhsj cvbc phc jqgm jqzklv gpjr nbgv gqnxlr xfn zcphr (contains fish, eggs, sesame)
pmhhqrk vhjpjdr rjc lcrs xvjk lpm dxdrk rtq vhjgg gbhjv lzvh vbkb pdt njdb fgcd bxqpgx npftghk ktnlk sqhvzg rlgr bvpxc ssrgt jbtlfv gpjr hfzqf fgsr xqgb bggmsj nmps lxldx djhj nppxr jhbnm pcslhrg mgv jgbk nlqhlsn fqtc qrpzt jfrlp mzbgp fnlk nvmm cgflg ggrc kdm zppxp nnskqnmn jqgm fctmvrs hrnvd vnbfvkp rtmfg nxrgp phc ssdszsn xbvrx rzcps rftbr ldrgj hdtr mrczqmj rtpff lkgln zgvtn hbdv czvphx spnd qpbl zvq zmsdzh rbjmdn tqkfx rdjdq tkhzz dchvb rfxgl fqqcnm (contains soy, nuts, eggs)
bvpxc chgjqc bggmsj rfgf qpbl fgcd bmqzgvr frxmq tktj lsgqf mxzb mzbgp bxqpgx gvdstsc rzcps hbbgk drtdz vhjgg fqqcnm rntk mbmtz spnd vxmrk thtlt nhzthvn mgv jp cncpbssj sqhvzg lftqn nmps lpm nggbtk hdtr xvjk pdt kdbxxzv qlkfvk zmsdzh fzlrf cldgd jqhn ktnlk npftghk prbk qhtvsr smdlg hqzcncv gqnxlr nthsf nvmm zvq nnskqnmn mhrlx rtq phc xtgjslz qrpzt gcptp rfx mxq hjvzcp rjc lkgln zgvtn xqgb (contains nuts)
mcl xvjk xsk sjpzc rsr zmsdzh zjbgp drtdz zcphr zvq jhv qfslcb fzlrf lcrs tqkfx rfgf prbk xqgb rjc spnd dchvb tcclbr strpjp flndv njqrhcc hfzqf pdt dkg lzvh gcpks hjvzcp lsgqf fgptl mzbgp sxsxm nvmm phc rng frxmq dxrsgs nnskqnmn nhnd bvcxfp (contains nuts, eggs)
rtq jfrlp dchvb fppgp jp xtgjslz sfkcp zppxp rzcps spnd mbmtz llj rbjmdn rsr nbgv lsgqf jqzklv mxc lxldx hdtr njrcfg phc gbhjv nthsf dkg nhnd lcrs fqqcnm hcjkd fgptl dfcgd dxdrk pdt nlqhlsn zmsdzh hbbgk pmhhqrk rjc pdhlzg dgxnc zfnttf jbtlfv nmps bmqzgvr fgcd jhbnm flndv hzmk chgjqc dhsssf nnskqnmn cldgd (contains nuts, peanuts)
sjpzc cvbc zcphr gpqgkt lsgqf ggrc rtmfg hrnvd cldgd ssdszsn tktj chgjqc zjbgp spnd njqrhcc hjvzcp mkxbf dgxnc ldrgj dxdrk nqvnn mzbgp ssrgt mbmtz nfgr jp pdhlzg hbdv gcpks xtgjslz sqhvzg ksrjn mcl zmsdzh fppgp dxrsgs blvfz lftqn jhv czvphx jlrt lpm bjcvpc jvhsj pdt zppxp bvlb zvq djdhrn nggbtk phc qfslcb tqkfx nmps nppxr jxvx rzcps vhjpjdr bvpxc snzxr lzvh fqqcnm qrpzt sxsxm (contains wheat, fish)
strpjp dxrsgs qlkfvk cgflg sxsxm hpbnj lsgqf mgv prbk prksl zmsdzh rtpff xfn czkfv dgxnc nfck ldrgj cncpbssj pfrqf zppxp bggmsj qzjrtl xvjk kjjst mbmtz sjpzc lftqn lkgln mxc rtmfg rdjdq hqzcncv jqgm psqc nxrgp hzmk rntk rms th kdm nthsf hjzkg dkg sqhvzg zvq gpqgkt qrpzt pdt zcphr rfxgl mrczqmj jbtlfv ssdszsn blvfz bvlb dfcgd nmps thtlt fqqcnm jp zfnttf fnlk frxmq rlgr hbbgk vbkb xgpdnz vhjgg czvphx vxmrk fgptl cldgd lpvfv lcrs pdhlzg hddd spnd phc pcslhrg rjc (contains eggs, peanuts, dairy)
zfnttf spnd rfx vbkb sfkcp tkhzz zmsdzh rfxgl dgxnc lsgqf sqhvzg phc bvpxc rlgr dxrsgs pdt vnbfvkp lftqn rfgf hcjkd fgsr nkzqj xgpdnz pmhhqrk tqkfx dfcgd cncpbssj rdkrtr qlkfvk blvfz xsk nfgr rsr czmml rtq bvcxfp kmlkx zdthsvl hzmk rjc xvjk gcpks tktj xfn prbk kdbxxzv fqqcnm mxc hjvzcp jp (contains peanuts)
nfck nnskqnmn rtpff sqhvzg lsgqf dfcgd hcjkd cvbc fzlrf bvpxc rftbr dgxnc jlrt cncpbssj cgflg fslqkg pdt nbgv xdcp nqvnn rsr jxvx hqzcncv mrczqmj gpjr fqqcnm zmsdzh cldgd qfslcb czkfv bfnnnrn prbk fqtc zjbgp qhtvsr rdkrtr bvcxfp tkhzz rfx smdlg nggbtk mzbgp qpbl nvmm th snzxr gbhjv rng vxmrk mbmtz mhrlx blvfz rzcps qrpzt nhnd lxldx gcptp zgvtn njdb kjjst mcl npftghk rjc jqhn cqqcd prksl lzvh phc vnbfvkp psqc jgbk zvq hfzqf (contains nuts, fish)
hcjkd rzcps fnlk fzlrf dgxnc cqqcd hbdv jvhsj lsgqf rxd bvcxfp lpvfv qpbl mrczqmj jqgm mxq jqhn hfzqf zmsdzh rjc lzvh nfgr vhjgg rntk rbjmdn ssdszsn tktj pdhlzg sqhvzg tqkfx czkxlg rfxgl gpqgkt njdb dkg strpjp zdthsvl nbgv lkgln vnbfvkp rfgf pcslhrg fppgp bmqzgvr hqzcncv rng dfcgd xbvrx qdkk hzmk vhjpjdr qlkfvk rms gqnxlr hdtr jlrt kdm sjpzc qhtvsr bvlb nppxr pdt vxmrk fqqcnm jxvx qrpzt thtlt zgvtn phc ksrjn fslqkg (contains soy, peanuts, fish)
rsr hfzqf fctmvrs zjbgp qdkk qjhq rfx hqzcncv nfgr nhzthvn xbvrx jhbnm pdt rng rdjdq pdj nnskqnmn jhv bfnnnrn kdm cldgd nppxr nthsf lzvh czvphx vhjgg hdtr mbmtz mgv jbz mhrlx fppgp rlgr ssrgt jvhsj tktj fqqcnm zbxj nvmm qzjrtl czmml ggrc zmsdzh tkhzz dxdrk jp mrczqmj phc kmlkx jqzklv nhnd rftbr thtlt pcslhrg hbmvpmt jsh nmps nggbtk rtq nqvnn lkgln snzxr xdcp fqtc lsgqf spnd mxzb llj gbhjv (contains fish)
spnd rng lsgqf nkzqj gbhjv mhrlx mcl hzmk sfkcp bfnnnrn qpbl dkg dxdrk nmps kdm djdhrn xbvrx rlgr tcclbr drtdz fqqcnm dhsssf gcptp kdbxxzv zcphr vnjxs mrczqmj rms kjjst bvlb fgsr bkzbrm jhv gpqgkt ssdszsn qjhq th zmsdzh kmlkx rjc hqzcncv rxd mxc jhbnm nggbtk bjcvpc njrcfg lftqn jbz jqhn jbtlfv hbmvpmt rsr nqvnn pdhlzg rzcps hdtr nnskqnmn rntk ssrgt qrpzt llj thtlt mbmtz mxzb pdt hfzqf nxrgp zdthsvl lxldx lzvh (contains peanuts)
rfgf hddd rng ssrgt tzdks zjbgp blvfz spnd lzvh ldrgj fqqcnm fnlk bvcxfp gcptp dkg bfnnnrn zfnttf kdbxxzv gpjr rntk jsh nxrgp njdb xtgjslz dhsssf dxdrk lpm nnskqnmn lcrs mzrx nvmm jp prksl rzcps xgpdnz czmml rbjmdn mxq qhtvsr dgxnc bxqpgx mbmtz nmps bmqzgvr dfcgd pdt xbvrx njrcfg rfx fppgp fgptl fzlrf qrpzt zmsdzh cncpbssj tkhzz phc mkxbf rjc njqrhcc cqqcd tqkfx mxc nhzthvn hzmk pdj (contains fish)
rms qrpzt hbbgk sxsxm dchvb jqhn gcptp xqgb psqc fctmvrs nlqhlsn tktj rntk qfslcb xdcp rfxgl fqqcnm jlrt zgvtn fgcd bvlb cvbc xbvrx rdkrtr rlgr prbk zmsdzh ssrgt phc tcclbr lpvfv djhj rsr rfx blvfz lzvh gpqgkt hpbnj bfnnnrn dxdrk nmps cncpbssj hjzkg hrnvd hbdv dgxnc hdtr rjc rdjdq jxvx qpbl zvq rtmfg jhbnm vhjpjdr nvmm lsgqf czkxlg pdt prksl xgpdnz rbjmdn gcpks fgsr zjbgp frxmq (contains soy, fish, peanuts)
drtdz nbgv hcjkd mxq sqhvzg xfn nfck phc gcptp hdtr nnskqnmn tqkfx rfxgl jqhn lpvfv blvfz czkxlg djhj lsgqf bggmsj flndv kmlkx mxzb czmml vnjxs spnd dchvb lzvh fzlrf bkzbrm xbvrx prbk ksrjn rxd fqqcnm xgpdnz pdt jfrlp bvpxc fgptl ktnlk njdb zmsdzh (contains nuts, wheat, fish)
vxmrk fqqcnm rlgr prbk fgsr tktj gqnxlr ssrgt phc rng gcpks mzrx dxdrk njqrhcc czmml cgflg qjhq vnjxs prksl ktnlk tcclbr qhtvsr sjpzc bggmsj nhzthvn zmsdzh ggrc jp rfgf rtpff zcphr lsgqf spnd blvfz rjc tkhzz sxsxm xfn pdt xqgb mrczqmj zdthsvl jhbnm bvcxfp ssdszsn djhj npftghk (contains nuts, eggs, sesame)
sfkcp bvpxc rfgf bjcvpc bggmsj jbz zppxp gpqgkt mgv lzvh qjhq hbbgk mrczqmj zmsdzh rlgr cvbc qzjrtl rbjmdn mxc czmml fqqcnm hfzqf jfrlp nhzthvn fctmvrs lftqn tcclbr djdhrn rjc cgflg spnd sjpzc phc pmhhqrk fslqkg flndv hzmk nvmm jp prksl dkg gcptp pdt (contains fish)
dchvb nppxr njdb fqqcnm hbbgk strpjp bvcxfp nhnd rjc hpbnj tkhzz jqzklv spnd vbkb jqgm phc pdt bxqpgx kdbxxzv zmsdzh xgpdnz rfxgl kdm drtdz nhzthvn ldrgj rdkrtr thtlt cvbc mhrlx bggmsj hjvzcp sjpzc nbgv gcptp hqd lcrs hddd djhj ssdszsn tqkfx xtgjslz mzbgp rdjdq nnskqnmn hqzcncv hjzkg jfrlp mxq nmps pmhhqrk pdj th fgcd psqc lsgqf gcpks xfn fgptl nfck mkxbf mbmtz rxd blvfz ggrc zppxp mcl ksrjn zbxj nkzqj hbdv lpm dxdrk zcphr qhtvsr pcslhrg tktj (contains sesame)
lzvh psqc rjc nhzthvn tqkfx lcrs pfrqf lkgln spnd nhnd pmhhqrk lxldx hqd dchvb zvq vnbfvkp mcl rsr rdkrtr hpbnj rfgf fzlrf ssrgt phc djdhrn fppgp rftbr xqgb hbdv vbkb tktj jqgm sfkcp pdt bggmsj bmqzgvr zfnttf jbtlfv nxrgp flndv sxsxm lsgqf zbxj dfcgd nmps mxc bvpxc rbjmdn gbhjv dxrsgs sjpzc jp fgcd zmsdzh jhbnm vxmrk bjcvpc zcphr gqnxlr njdb nvmm jqzklv kdbxxzv jsh xtgjslz gcpks rntk qpbl djhj lftqn ldrgj xfn hdtr tzdks tcclbr gcptp nnskqnmn (contains eggs, soy)
gpqgkt lftqn frxmq dchvb hbdv gvdstsc djdhrn ssrgt drtdz rjc gpjr jvhsj rfxgl hqd spnd rtpff mzrx qfslcb cgflg jqhn strpjp zcphr xtgjslz lsgqf pcslhrg smdlg czmml nqvnn fqtc qpbl dkg phc xfn fgsr fppgp dgxnc cqqcd qzjrtl ggrc lkgln ksrjn nvmm qdkk rxd jlrt jhbnm nhnd lpm hjvzcp bfnnnrn pmhhqrk rdjdq hcjkd czvphx xbvrx zgvtn dxrsgs fqqcnm zmsdzh njqrhcc lzvh rzcps fzlrf (contains sesame, soy)
jgbk gvdstsc spnd rtmfg psqc njdb drtdz thtlt dxrsgs cldgd hjvzcp nggbtk dfcgd nppxr mbmtz gcpks rdjdq lzvh czvphx rjc vnbfvkp mhrlx frxmq bfnnnrn kmlkx qpbl lsgqf ldrgj qrpzt jp bvpxc phc zmsdzh jbtlfv nhzthvn rsr jxvx th rzcps kdbxxzv xgpdnz nqvnn rdkrtr pmhhqrk pdt prbk rbjmdn fgptl gpjr zbxj czmml tqkfx xdcp njqrhcc cncpbssj chgjqc fgcd (contains nuts, dairy, eggs)
`

var testInput = `
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`

type Food struct {
	Ingredients []string
	Allergens   []string
}

func main() {
	var food []Food

	for _, l := range util.ReadInput(input, "\n") {
		s := strings.Split(l, " (")
		ing := strings.Split(s[0], " ")
		all := strings.Split(s[1][len("contains "):len(s[1])-1], ", ")

		food = append(food, Food{
			Ingredients: ing,
			Allergens:   all,
		})
	}

	fmt.Printf("first %d\n", first(food))
	fmt.Printf("second %s\n", second(food))
}

func first(food []Food) int {
	ingredients := make(map[string]bool)

	for _, f := range food {
		for _, ing := range f.Ingredients {
			ingredients[ing] = true
		}
	}

	alwaysListed := findAllergenMapping(food)

	others := make(map[string]bool)
	for ing := range ingredients {
		isListed := false
		for _, ings := range alwaysListed {
			for _, in := range ings {
				if in == ing {
					isListed = true
					break
				}
			}
		}

		if !isListed {
			others[ing] = true
		}
	}

	contained := 0
	for ing := range others {
		for _, f := range food {
			for _, in := range f.Ingredients {
				if ing == in {
					contained += 1
				}
			}
		}
	}

	return contained
}

func second(food []Food) string {
	mapping := findAllergenMapping(food)
	changesMade := true
	fm := make(map[string]string)
	var toRemove []string

	for changesMade {
		for a, f := range mapping {
			if len(f) == 1 {
				fm[f[0]] = a
				toRemove = append(toRemove, f[0])
			}
		}

		newm := make(map[string][]string)
		for a, fs := range mapping {
			var nf []string
			for _, f := range fs {
				remove := false
				for _, r := range toRemove {
					if f == r {
						remove = true
						break
					}
				}

				if !remove {
					nf = append(nf, f)
				}
			}

			newm[a] = nf
		}

		mapping = newm
		changesMade = len(toRemove) > 0
		toRemove = []string{}
	}

	var fs []string
	for ing := range fm {
		fs = append(fs, ing)
	}

	sort.Slice(fs, func(i, j int) bool {
		return fm[fs[i]] < fm[fs[j]]
	})

	return strings.Join(fs, ",")
}

func findAllergenMapping(food []Food) map[string][]string {
	ingredients := make(map[string]bool)
	allergenMap := make(map[string][]Food)

	for _, f := range food {
		for _, a := range f.Allergens {
			allergenMap[a] = append(allergenMap[a], f)
		}

		for _, ing := range f.Ingredients {
			ingredients[ing] = true
		}
	}

	alwaysListed := make(map[string][]string)

	for a, foods := range allergenMap {
		contains := make(map[string]int)

		for _, f := range foods {
			for _, ing := range f.Ingredients {
				contains[ing] += 1
			}
		}

		for ing, n := range contains {
			if n == len(foods) {
				alwaysListed[a] = append(alwaysListed[a], ing)
			}
		}
	}

	return alwaysListed
}
