// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const MIPSStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/cannon/MIPS.sol:MIPS\",\"label\":\"oracle\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_contract(IPreimageOracle)1001\"}],\"types\":{\"t_contract(IPreimageOracle)1001\":{\"encoding\":\"inplace\",\"label\":\"contract IPreimageOracle\",\"numberOfBytes\":\"20\"}}}"

var MIPSStorageLayout = new(solc.StorageLayout)

var MIPSDeployedBin = "0x608060405234801561001057600080fd5b50600436106100415760003560e01c8063155633fe146100465780637dc0d1d014610067578063f8e0cb9614610098575b600080fd5b61004e61016c565b6040805163ffffffff9092168252519081900360200190f35b61006f610174565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61015a600480360360408110156100ae57600080fd5b8101906020810181356401000000008111156100c957600080fd5b8201836020820111156100db57600080fd5b803590602001918460018302840111640100000000831117156100fd57600080fd5b91939092909160208101903564010000000081111561011b57600080fd5b82018360208201111561012d57600080fd5b8035906020019184600183028401116401000000008311171561014f57600080fd5b509092509050610190565b60408051918252519081900360200190f35b634000000081565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600061019a611a6a565b608081146101a757600080fd5b604051610600146101b757600080fd5b606486146101c457600080fd5b61016684146101d257600080fd5b6101ef565b8035602084810360031b9190911c8352920192910190565b8560806101fe602082846101d7565b9150915061020e602082846101d7565b9150915061021e600482846101d7565b9150915061022e600482846101d7565b9150915061023e600482846101d7565b9150915061024e600482846101d7565b9150915061025e600482846101d7565b9150915061026e600482846101d7565b9150915061027e600182846101d7565b9150915061028e600182846101d7565b9150915061029e600882846101d7565b6020810190819052909250905060005b60208110156102d0576102c3600483856101d7565b90935091506001016102ae565b505050806101200151156102ee576102e6610710565b915050610708565b6101408101805160010167ffffffffffffffff1690526060810151600090610316908261081e565b9050603f601a82901c16600281148061033557508063ffffffff166003145b15610382576103788163ffffffff1660021461035257601f610355565b60005b60ff16600261036b856303ffffff16601a6108e9565b63ffffffff16901b61095c565b9350505050610708565b6101608301516000908190601f601086901c81169190601587901c16602081106103a857fe5b602002015192508063ffffffff851615806103c957508463ffffffff16601c145b156103fa578661016001518263ffffffff16602081106103e557fe5b6020020151925050601f600b86901c166104b1565b60208563ffffffff16101561045d578463ffffffff16600c148061042457508463ffffffff16600d145b8061043557508463ffffffff16600e145b15610446578561ffff169250610458565b6104558661ffff1660106108e9565b92505b6104b1565b60288563ffffffff1610158061047957508463ffffffff166022145b8061048a57508463ffffffff166026145b156104b1578661016001518263ffffffff16602081106104a657fe5b602002015192508190505b60048563ffffffff16101580156104ce575060088563ffffffff16105b806104df57508463ffffffff166001145b156104fe576104f0858784876109c7565b975050505050505050610708565b63ffffffff60006020878316106105635761051e8861ffff1660106108e9565b9095019463fffffffc861661053481600161081e565b915060288863ffffffff161015801561055457508763ffffffff16603014155b1561056157809250600093505b505b600061057189888885610b50565b63ffffffff9081169150603f8a16908916158015610596575060088163ffffffff1610155b80156105a85750601c8163ffffffff16105b15610687578063ffffffff16600814806105c857508063ffffffff166009145b156105ff576105ed8163ffffffff166008146105e457856105e7565b60005b8961095c565b9b505050505050505050505050610708565b8063ffffffff16600a1415610620576105ed858963ffffffff8a1615611216565b8063ffffffff16600b1415610642576105ed858963ffffffff8a161515611216565b8063ffffffff16600c1415610659576105ed6112fb565b60108163ffffffff16101580156106765750601c8163ffffffff16105b15610687576105ed81898988611778565b8863ffffffff1660381480156106a2575063ffffffff861615155b156106d15760018b61016001518763ffffffff16602081106106c057fe5b63ffffffff90921660209290920201525b8363ffffffff1663ffffffff146106ee576106ee8460018461195c565b6106fa85836001611216565b9b5050505050505050505050505b949350505050565b6000610728565b602083810382015183520192910190565b60806040518061073a60208285610717565b9150925061074a60208285610717565b9150925061075a60048285610717565b9150925061076a60048285610717565b9150925061077a60048285610717565b9150925061078a60048285610717565b9150925061079a60048285610717565b915092506107aa60048285610717565b915092506107ba60018285610717565b915092506107ca60018285610717565b915092506107da60088285610717565b60209091019350905060005b6020811015610808576107fb60048386610717565b90945091506001016107e6565b506000815281810382a081900390209150505b90565b60008061082a836119f8565b9050600384161561083a57600080fd5b602081019035610857565b60009081526020919091526040902090565b8460051c8160005b601b8110156108af5760208501943583821c60011680156108875760018114610898576108a5565b6108918285610845565b93506108a5565b6108a28483610845565b93505b505060010161085f565b5060805191508181146108ca57630badf00d60005260206000fd5b5050601f94909416601c0360031b9390931c63ffffffff169392505050565b600063ffffffff8381167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80850183169190911c821615159160016020869003821681901b830191861691821b92911b0182610946576000610948565b815b90861663ffffffff16179250505092915050565b6000610966611a6a565b5060e08051610100805163ffffffff90811690935284831690526080918516156109b657806008018261016001518663ffffffff16602081106109a557fe5b63ffffffff90921660209290920201525b6109be610710565b95945050505050565b60006109d1611a6a565b5060806000600463ffffffff881614806109f157508663ffffffff166005145b15610a675760008261016001518663ffffffff1660208110610a0f57fe5b602002015190508063ffffffff168563ffffffff16148015610a3757508763ffffffff166004145b80610a5f57508063ffffffff168563ffffffff1614158015610a5f57508763ffffffff166005145b915050610ae4565b8663ffffffff1660061415610a855760008460030b13159050610ae4565b8663ffffffff1660071415610aa25760008460030b139050610ae4565b8663ffffffff1660011415610ae457601f601087901c1680610ac85760008560030b1291505b8063ffffffff1660011415610ae25760008560030b121591505b505b606082018051608084015163ffffffff169091528115610b2a576002610b0f8861ffff1660106108e9565b63ffffffff90811690911b8201600401166080840152610b3c565b60808301805160040163ffffffff1690525b610b44610710565b98975050505050505050565b6000603f601a86901c81169086166020821015610f245760088263ffffffff1610158015610b845750600f8263ffffffff16105b15610c2b578163ffffffff1660081415610ba057506020610c26565b8163ffffffff1660091415610bb757506021610c26565b8163ffffffff16600a1415610bce5750602a610c26565b8163ffffffff16600b1415610be55750602b610c26565b8163ffffffff16600c1415610bfc57506024610c26565b8163ffffffff16600d1415610c1357506025610c26565b8163ffffffff16600e1415610c26575060265b600091505b63ffffffff8216610e7457601f600688901c16602063ffffffff83161015610d485760088263ffffffff1610610c6657869350505050610708565b63ffffffff8216610c865763ffffffff86811691161b9250610708915050565b8163ffffffff1660021415610caa5763ffffffff86811691161c9250610708915050565b8163ffffffff1660031415610cd5576103788163ffffffff168763ffffffff16901c826020036108e9565b8163ffffffff1660041415610cf9575050505063ffffffff8216601f84161b610708565b8163ffffffff1660061415610d1d575050505063ffffffff8216601f84161c610708565b8163ffffffff1660071415610d48576103788763ffffffff168763ffffffff16901c886020036108e9565b8163ffffffff1660201480610d6357508163ffffffff166021145b15610d75578587019350505050610708565b8163ffffffff1660221480610d9057508163ffffffff166023145b15610da2578587039350505050610708565b8163ffffffff1660241415610dbe578587169350505050610708565b8163ffffffff1660251415610dda578587179350505050610708565b8163ffffffff1660261415610df6578587189350505050610708565b8163ffffffff1660271415610e12575050505082821719610708565b8163ffffffff16602a1415610e45578560030b8760030b12610e35576000610e38565b60015b60ff169350505050610708565b8163ffffffff16602b1415610e6e578563ffffffff168763ffffffff1610610e35576000610e38565b50610f1f565b8163ffffffff16600f1415610e975760108563ffffffff16901b92505050610708565b8163ffffffff16601c1415610f1f578063ffffffff1660021415610ec057505050828202610708565b8063ffffffff1660201480610edb57508063ffffffff166021145b15610f1f578063ffffffff1660201415610ef3579419945b60005b6380000000871615610f15576401fffffffe600197881b169601610ef6565b9250610708915050565b6111af565b60288263ffffffff16101561108e578163ffffffff1660201415610f7157610f688660031660080260180363ffffffff168563ffffffff16901c60ff1660086108e9565b92505050610708565b8163ffffffff1660211415610fa757610f688660021660080260100363ffffffff168563ffffffff16901c61ffff1660106108e9565b8163ffffffff1660221415610fd85750505063ffffffff60086003851602811681811b198416918316901b17610708565b8163ffffffff1660231415610ff1578392505050610708565b8163ffffffff1660241415611025578560031660080260180363ffffffff168463ffffffff16901c60ff1692505050610708565b8163ffffffff166025141561105a578560021660080260100363ffffffff168463ffffffff16901c61ffff1692505050610708565b8163ffffffff1660261415610f1f5750505063ffffffff60086003851602601803811681811c198416918316901c17610708565b8163ffffffff16602814156110c65750505060ff63ffffffff60086003861602601803811682811b9091188316918416901b17610708565b8163ffffffff16602914156110ff5750505061ffff63ffffffff60086002861602601003811682811b9091188316918416901b17610708565b8163ffffffff16602a14156111305750505063ffffffff60086003851602811681811c198316918416901c17610708565b8163ffffffff16602b1415611149578492505050610708565b8163ffffffff16602e141561117d5750505063ffffffff60086003851602601803811681811b198316918416901b17610708565b8163ffffffff1660301415611196578392505050610708565b8163ffffffff16603814156111af578492505050610708565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f696e76616c696420696e737472756374696f6e00000000000000000000000000604482015290519081900360640190fd5b6000611220611a6a565b506080602063ffffffff86161061129857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f76616c6964207265676973746572000000000000000000000000000000000000604482015290519081900360640190fd5b63ffffffff8516158015906112aa5750825b156112d857838161016001518663ffffffff16602081106112c757fe5b63ffffffff90921660209290920201525b60808101805163ffffffff808216606085015260049091011690526109be610710565b6000611305611a6a565b506101e051604081015160808083015160a084015160c09094015191936000928392919063ffffffff8616610ffa141561137d5781610fff81161561134f57610fff811661100003015b63ffffffff84166113735760e08801805163ffffffff838201169091529550611377565b8395505b5061172b565b8563ffffffff16610fcd1415611399576340000000945061172b565b8563ffffffff1661101814156113b2576001945061172b565b8563ffffffff1661109614156113ea57600161012088015260ff83166101008801526113dc610710565b97505050505050505061081b565b8563ffffffff16610fa314156115a95763ffffffff831661140a576115a4565b63ffffffff83166005141561158157600061142c8363fffffffc16600161081e565b6000805460208b01516040808d015181517fe03110e1000000000000000000000000000000000000000000000000000000008152600481019390935263ffffffff16602483015280519495509293849373ffffffffffffffffffffffffffffffffffffffff9093169263e03110e19260448082019391829003018186803b1580156114b657600080fd5b505afa1580156114ca573d6000803e3d6000fd5b505050506040513d60408110156114e057600080fd5b508051602090910151909250905060038516600481900382811015611503578092505b5081851015611510578491505b8260088302610100031c9250826008828460040301021b9250600180600883600403021b036001806008858560040303021b039150811981169050838119861617945050506115678563fffffffc1660018561195c565b60408a018051820163ffffffff16905296506115a4915050565b63ffffffff831660031415611598578094506115a4565b63ffffffff9450600993505b61172b565b8563ffffffff16610fa4141561167d5763ffffffff8316600114806115d4575063ffffffff83166002145b806115e5575063ffffffff83166004145b156115f2578094506115a4565b63ffffffff8316600614156115985760006116148363fffffffc16600161081e565b6020890151909150600384166004038381101561162f578093505b83900360089081029290921c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600193850293841b0116911b176020880152600060408801529350836115a4565b8563ffffffff16610fd7141561172b578163ffffffff166003141561171f5763ffffffff831615806116b5575063ffffffff83166005145b806116c6575063ffffffff83166003145b156116d457600094506115a4565b63ffffffff8316600114806116ef575063ffffffff83166002145b80611700575063ffffffff83166006145b80611711575063ffffffff83166004145b1561159857600194506115a4565b63ffffffff9450601693505b6101608701805163ffffffff808816604090920191909152905185821660e09091015260808801805180831660608b0152600401909116905261176c610710565b97505050505050505090565b6000611782611a6a565b5060806000601063ffffffff881614156117a1575060c08101516118f9565b8663ffffffff16601114156117c15763ffffffff861660c08301526118f9565b8663ffffffff16601214156117db575060a08101516118f9565b8663ffffffff16601314156117fb5763ffffffff861660a08301526118f9565b8663ffffffff16601814156118305763ffffffff600387810b9087900b02602081901c821660c08501521660a08301526118f9565b8663ffffffff16601914156118625763ffffffff86811681871602602081901c821660c08501521660a08301526118f9565b8663ffffffff16601a14156118ad578460030b8660030b8161188057fe5b0763ffffffff1660c0830152600385810b9087900b8161189c57fe5b0563ffffffff1660a08301526118f9565b8663ffffffff16601b14156118f9578463ffffffff168663ffffffff16816118d157fe5b0663ffffffff90811660c0840152858116908716816118ec57fe5b0463ffffffff1660a08301525b63ffffffff84161561192e57808261016001518563ffffffff166020811061191d57fe5b63ffffffff90921660209290920201525b60808201805163ffffffff80821660608601526004909101169052611951610710565b979650505050505050565b6000611967836119f8565b9050600384161561197757600080fd5b6020810190601f8516601c0360031b83811b913563ffffffff90911b1916178460051c60005b601b8110156119ed5760208401933582821c60011680156119c557600181146119d6576119e3565b6119cf8286610845565b94506119e3565b6119e08583610845565b94505b505060010161199d565b505060805250505050565b60ff81166103800261016681019036906104e601811015611a64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611af56023913960400191505060405180910390fd5b50919050565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101919091526101608101611ad0611ad5565b905290565b604051806104000160405280602090602082028036833750919291505056fe636865636b207468617420746865726520697320656e6f7567682063616c6c64617461a164736f6c6343000706000a"

var MIPSDeployedSourceMap = "1075:33529:0:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1655:45;;;:::i;:::-;;;;;;;;;;;;;;;;;;;2081:29;;;:::i;:::-;;;;;;;;;;;;;;;;;;;22315:5721;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;22315:5721:0;;-1:-1:-1;22315:5721:0;-1:-1:-1;22315:5721:0;:::i;:::-;;;;;;;;;;;;;;;;1655:45;1690:10;1655:45;:::o;2081:29::-;;;;;;:::o;22315:5721::-;22393:7;22412:18;;:::i;:::-;22547:4;22540:5;22537:15;22527:2;;22616:1;22614;22607:11;22527:2;22664:4;22658:11;22671;22655:28;22645:2;;22737:1;22735;22728:11;22645:2;22797:3;22779:16;22776:25;22766:2;;22871:1;22869;22862:11;22766:2;22927:3;22913:12;22910:21;22900:2;;23000:1;22998;22991:11;22900:2;23030:416;;;23264:24;;23252:2;23248:13;;;23245:1;23241:21;23237:52;;;;23306:20;;23360:21;;;23414:18;;;23108:338::o;:::-;23523:16;23581:4;23633:18;23648:2;23645:1;23642;23633:18;:::i;:::-;23625:26;;;;23683:18;23698:2;23695:1;23692;23683:18;:::i;:::-;23675:26;;;;23737:17;23752:1;23749;23746;23737:17;:::i;:::-;23729:25;;;;23794:17;23809:1;23806;23803;23794:17;:::i;:::-;23786:25;;;;23839:17;23854:1;23851;23848;23839:17;:::i;:::-;23831:25;;;;23888:17;23903:1;23900;23897;23888:17;:::i;:::-;23880:25;;;;23933:17;23948:1;23945;23942;23933:17;:::i;:::-;23925:25;;;;23978:17;23993:1;23990;23987;23978:17;:::i;:::-;23970:25;;;;24025:17;24040:1;24037;24034;24025:17;:::i;:::-;24017:25;;;;24076:17;24091:1;24088;24085;24076:17;:::i;:::-;24068:25;;;;24125:17;24140:1;24137;24134;24125:17;:::i;:::-;24234:2;24227:10;;24217:21;;;;24117:25;;-1:-1:-1;24227:10:0;-1:-1:-1;24322:1:0;24307:105;24332:2;24329:1;24326:9;24307:105;;;24381:17;24396:1;24393;24390;24381:17;:::i;:::-;24373:25;;-1:-1:-1;24373:25:0;-1:-1:-1;24350:1:0;24343:9;24307:105;;;24311:14;;;24478:5;:12;;;24474:63;;;24513:13;:11;:13::i;:::-;24506:20;;;;;24474:63;24547:10;;;:15;;24561:1;24547:15;;;;;24624:8;;;;-1:-1:-1;;24616:20:0;;-1:-1:-1;24616:7:0;:20::i;:::-;24602:34;-1:-1:-1;24662:10:0;24670:2;24662:10;;;;24731:1;24721:11;;;:26;;;24736:6;:11;;24746:1;24736:11;24721:26;24717:332;;;24974:64;24985:6;:11;;24995:1;24985:11;:20;;25003:2;24985:20;;;24999:1;24985:20;24974:64;;25036:1;25007:25;25010:4;25017:10;25010:17;25029:2;25007;:25::i;:::-;:30;;;;24974:10;:64::i;:::-;24967:71;;;;;;;24717:332;25274:15;;;;25085:9;;;;25214:4;25208:2;25200:10;;;25199:19;;;25274:15;25299:2;25291:10;;;25290:19;25274:36;;;;;;;;;;;;-1:-1:-1;25335:5:0;25355:11;;;;;:29;;;25370:6;:14;;25380:4;25370:14;25355:29;25351:756;;;25439:5;:15;;;25455:5;25439:22;;;;;;;;;;;;;;-1:-1:-1;;25498:4:0;25492:2;25484:10;;;25483:19;25351:756;;;25532:4;25523:6;:13;;;25519:588;;;25641:6;:13;;25651:3;25641:13;:30;;;;25658:6;:13;;25668:3;25658:13;25641:30;:47;;;;25675:6;:13;;25685:3;25675:13;25641:47;25637:229;;;25743:4;25750:6;25743:13;25738:18;;25637:229;;;25830:21;25833:4;25840:6;25833:13;25848:2;25830;:21::i;:::-;25825:26;;25637:229;25519:588;;;25896:4;25886:6;:14;;;;:32;;;;25904:6;:14;;25914:4;25904:14;25886:32;:50;;;;25922:6;:14;;25932:4;25922:14;25886:50;25882:225;;;25998:5;:15;;;26014:5;25998:22;;;;;;;;;;;;;25993:27;;26091:5;26083:13;;25882:225;26132:1;26122:6;:11;;;;:25;;;;;26146:1;26137:6;:10;;;26122:25;26121:42;;;;26152:6;:11;;26162:1;26152:11;26121:42;26117:117;;;26186:37;26199:6;26207:4;26213:5;26220:2;26186:12;:37::i;:::-;26179:44;;;;;;;;;;;26117:117;26263:13;26244:16;26399:4;26389:14;;;;26385:400;;26460:19;26463:4;26468:6;26463:11;26476:2;26460;:19::i;:::-;26454:25;;;;26512:10;26507:15;;26542:16;26507:15;26556:1;26542:7;:16::i;:::-;26536:22;;26586:4;26576:6;:14;;;;:32;;;;;26594:6;:14;;26604:4;26594:14;;26576:32;26572:203;;;26665:4;26653:16;;26759:1;26751:9;;26572:203;26385:400;;26810:10;26823:26;26831:4;26837:2;26841;26845:3;26823:7;:26::i;:::-;26852:10;26823:39;;;;-1:-1:-1;26944:4:0;26937:11;;;26972;;;:24;;;;;26995:1;26987:4;:9;;;;26972:24;:39;;;;;27007:4;27000;:11;;;26972:39;26968:711;;;27031:4;:9;;27039:1;27031:9;:22;;;;27044:4;:9;;27052:1;27044:9;27031:22;27027:116;;;27091:37;27102:4;:9;;27110:1;27102:9;:21;;27118:5;27102:21;;;27114:1;27102:21;27125:2;27091:10;:37::i;:::-;27084:44;;;;;;;;;;;;;;;27027:116;27161:4;:11;;27169:3;27161:11;27157:93;;;27207:28;27216:5;27223:2;27227:7;;;;27207:8;:28::i;27157:93::-;27267:4;:11;;27275:3;27267:11;27263:93;;;27313:28;27322:5;27329:2;27333:7;;;;;27313:8;:28::i;27263:93::-;27418:4;:11;;27426:3;27418:11;27414:72;;;27456:15;:13;:15::i;27414:72::-;27577:4;27569;:12;;;;:27;;;;;27592:4;27585;:11;;;27569:27;27565:104;;;27623:31;27634:4;27640:2;27644;27648:5;27623:10;:31::i;27565:104::-;27731:6;:14;;27741:4;27731:14;:28;;;;-1:-1:-1;27749:10:0;;;;;27731:28;27727:85;;;27800:1;27775:5;:15;;;27791:5;27775:22;;;;;;;;;:26;;;;:22;;;;;;:26;27727:85;27850:9;:26;;27863:13;27850:26;27846:84;;27892:27;27901:9;27912:1;27915:3;27892:8;:27::i;:::-;28003:26;28012:5;28019:3;28024:4;28003:8;:26::i;:::-;27996:33;;;;;;;;;;;;;22315:5721;;;;;;;:::o;2605:1791::-;2646:12;2791:206;;;2891:2;2887:13;;;2877:24;;2871:31;2860:43;;2931:13;;2970;;;2842:155::o;:::-;3068:4;3152;3146:11;3180:5;3252:21;3270:2;3266;3260:4;3252:21;:::i;:::-;3240:33;;;;3310:21;3328:2;3324;3318:4;3310:21;:::i;:::-;3298:33;;;;3372:20;3390:1;3386:2;3380:4;3372:20;:::i;:::-;3360:32;;;;3437:20;3455:1;3451:2;3445:4;3437:20;:::i;:::-;3425:32;;;;3490:20;3508:1;3504:2;3498:4;3490:20;:::i;:::-;3478:32;;;;3547:20;3565:1;3561:2;3555:4;3547:20;:::i;:::-;3535:32;;;;3600:20;3618:1;3614:2;3608:4;3600:20;:::i;:::-;3588:32;;;;3653:20;3671:1;3667:2;3661:4;3653:20;:::i;:::-;3641:32;;;;3708:20;3726:1;3722:2;3716:4;3708:20;:::i;:::-;3696:32;;;;3767:20;3785:1;3781:2;3775:4;3767:20;:::i;:::-;3755:32;;;;3824:20;3842:1;3838:2;3832:4;3824:20;:::i;:::-;3885:2;3875:13;;;;-1:-1:-1;3812:32:0;-1:-1:-1;3983:1:0;3968:112;3993:2;3990:1;3987:9;3968:112;;;4046:20;4064:1;4060:2;4054:4;4046:20;:::i;:::-;4034:32;;-1:-1:-1;4034:32:0;-1:-1:-1;4011:1:0;4004:9;3968:112;;;3972:14;4143:1;4139:2;4132:13;4238:5;4234:2;4230:14;4223:5;4218:27;4344:14;;;4327:32;;;-1:-1:-1;;2605:1791:0;;:::o;18530:1741::-;18603:11;18686:14;18703:24;18715:11;18703;:24::i;:::-;18686:41;;18823:1;18816:5;18812:13;18809:2;;;18854:1;18851;18844:12;18809:2;18987;18975:15;;;18932:20;19085:141;;;;19132:12;;;19168:2;19161:13;;;;19209:2;19196:16;;;19114:112::o;:::-;19381:5;19378:1;19374:13;19412:4;19444:1;19429:375;19454:2;19451:1;19448:9;19429:375;;;19569:2;19557:15;;;19510:20;19600:12;;;19614:1;19596:20;19633:78;;;;19717:1;19712:78;;;;19589:201;;19633:78;19670:23;19685:7;19679:4;19670:23;:::i;:::-;19662:31;;19633:78;;19712;19749:23;19767:4;19758:7;19749:23;:::i;:::-;19741:31;;19589:201;-1:-1:-1;;19472:1:0;19465:9;19429:375;;;19433:14;19906:4;19900:11;19885:26;;19984:7;19978:4;19975:17;19965:2;;20022:10;20019:1;20012:21;20060:2;20057:1;20050:13;19965:2;-1:-1:-1;;20196:2:0;20185:14;;;;20173:10;20169:31;20166:1;20162:39;20226:16;;;;20244:10;20222:33;;18747:1518;-1:-1:-1;;;18747:1518:0:o;2209:288::-;2270:6;2305:18;;;;2314:8;;;;2305:18;;;;;;2304:25;;;;;2321:1;2364:2;:9;;;2358:16;;;;;2357:22;;2356:32;;;;;;;2414:9;;2413:15;2304:25;2467:21;;2487:1;2467:21;;;2478:6;2467:21;2452:11;;;;;:37;;-1:-1:-1;;;2209:288:0;;;;:::o;16019:624::-;16088:12;16147:18;;:::i;:::-;-1:-1:-1;16302:8:0;;;16331:12;;;16320:23;;;;;;;16353:20;;;;;16207:4;;16477:13;;;16473:82;;16534:6;16543:1;16534:10;16506:5;:15;;;16522:8;16506:25;;;;;;;;;:38;;;;:25;;;;;;:38;16473:82;16623:13;:11;:13::i;:::-;16616:20;16019:624;-1:-1:-1;;;;;16019:624:0:o;11451:1713::-;11548:12;11606:18;;:::i;:::-;-1:-1:-1;11666:4:0;11690:17;11789:1;11778:12;;;;;:28;;;11794:7;:12;;11805:1;11794:12;11778:28;11774:859;;;11822:9;11834:5;:15;;;11850:6;11834:23;;;;;;;;;;;;;11822:35;;11894:2;11887:9;;:3;:9;;;:25;;;;;11900:7;:12;;11911:1;11900:12;11887:25;11886:58;;;;11925:2;11918:9;;:3;:9;;;;:25;;;;;11931:7;:12;;11942:1;11931:12;11918:25;11871:73;;11774:859;;;;12044:7;:12;;12055:1;12044:12;12040:593;;;12101:1;12093:3;12087:15;;;;12072:30;;12040:593;;;12193:7;:12;;12204:1;12193:12;12189:444;;;12249:1;12242:3;12236:14;;;12221:29;;12189:444;;;12358:7;:12;;12369:1;12358:12;12354:279;;;12438:4;12432:2;12423:11;;;12422:20;12461:8;12457:76;;12517:1;12510:3;12504:14;;;12489:29;;12457:76;12550:3;:8;;12557:1;12550:8;12546:77;;;12607:1;12599:3;12593:15;;;;12578:30;;12546:77;12354:279;;12701:8;;;;;12771:12;;;;12760:23;;;;;12915:162;;;;13002:1;12976:22;12979:5;12987:6;12979:14;12995:2;12976;:22::i;:::-;:27;;;;;;;12962:42;;12971:1;12962:42;12947:57;:12;;;:57;12915:162;;;13050:12;;;;;13065:1;13050:16;13035:31;;;;12915:162;13144:13;:11;:13::i;:::-;13137:20;11451:1713;-1:-1:-1;;;;;;;;11451:1713:0:o;28082:6520::-;28169:6;28203:10;28211:2;28203:10;;;;;;28250:11;;28354:4;28345:13;;28341:6215;;;28473:1;28463:6;:11;;;;:27;;;;;28487:3;28478:6;:12;;;28463:27;28459:532;;;28514:6;:11;;28524:1;28514:11;28510:431;;;-1:-1:-1;28536:4:0;28510:431;;;28584:6;:11;;28594:1;28584:11;28580:361;;;-1:-1:-1;28606:4:0;28580:361;;;28650:6;:13;;28660:3;28650:13;28646:295;;;-1:-1:-1;28674:4:0;28646:295;;;28715:6;:13;;28725:3;28715:13;28711:230;;;-1:-1:-1;28739:4:0;28711:230;;;28781:6;:13;;28791:3;28781:13;28777:164;;;-1:-1:-1;28805:4:0;28777:164;;;28846:6;:13;;28856:3;28846:13;28842:99;;;-1:-1:-1;28870:4:0;28842:99;;;28910:6;:13;;28920:3;28910:13;28906:35;;;-1:-1:-1;28934:4:0;28906:35;28975:1;28966:10;;28459:532;29044:11;;;29040:3190;;29104:4;29099:1;29091:9;;;29090:18;29137:4;29091:9;29130:11;;;29126:1203;;;29221:4;29213;:12;;;29209:1102;;29260:2;29253:9;;;;;;;29209:1102;29362:12;;;29358:953;;29409:11;;;;;;;;-1:-1:-1;29402:18:0;;-1:-1:-1;;29402:18:0;29358:953;29521:4;:12;;29529:4;29521:12;29517:794;;;29568:11;;;;;;;;-1:-1:-1;29561:18:0;;-1:-1:-1;;29561:18:0;29517:794;29683:4;:12;;29691:4;29683:12;29679:632;;;29730:27;29739:5;29733:11;;:2;:11;;;;29751:5;29746:2;:10;29730:2;:27::i;29679:632::-;29867:4;:12;;29875:4;29867:12;29863:448;;;-1:-1:-1;;;;29914:17:0;;;29926:4;29921:9;;29914:17;29907:24;;29863:448;30042:4;:12;;30050:4;30042:12;30038:273;;;-1:-1:-1;;;;30089:17:0;;;30101:4;30096:9;;30089:17;30082:24;;30038:273;30220:4;:12;;30228:4;30220:12;30216:95;;;30267:21;30276:2;30270:8;;:2;:8;;;;30285:2;30280;:7;30267:2;:21::i;30216:95::-;30473:4;:12;;30481:4;30473:12;:28;;;;30489:4;:12;;30497:4;30489:12;30473:28;30469:1025;;;30537:2;30532;:7;30525:14;;;;;;;30469:1025;30615:4;:12;;30623:4;30615:12;:28;;;;30631:4;:12;;30639:4;30631:12;30615:28;30611:883;;;30679:2;30674;:7;30667:14;;;;;;;30611:883;30749:4;:12;;30757:4;30749:12;30745:749;;;30797:2;30792;:7;30785:14;;;;;;;30745:749;30866:4;:12;;30874:4;30866:12;30862:632;;;30915:2;30910;:7;30902:16;;;;;;;30862:632;30986:4;:12;;30994:4;30986:12;30982:512;;;31035:2;31030;:7;31022:16;;;;;;;30982:512;31106:4;:12;;31114:4;31106:12;31102:392;;;-1:-1:-1;;;;31151:7:0;;;31149:10;31142:17;;31102:392;31250:4;:12;;31258:4;31250:12;31246:248;;;31309:2;31291:21;;31297:2;31291:21;;;:29;;31319:1;31291:29;;;31315:1;31291:29;31284:36;;;;;;;;;31246:248;31421:4;:12;;31429:4;31421:12;31417:77;;;31465:2;31462:5;;:2;:5;;;:13;;31474:1;31462:13;;31417:77;29040:3190;;;;31571:6;:13;;31581:3;31571:13;31567:663;;;31617:2;31611;:8;;;;31604:15;;;;;;31567:663;31680:6;:14;;31690:4;31680:14;31676:554;;;31741:4;:9;;31749:1;31741:9;31737:92;;;-1:-1:-1;;;31788:21:0;;;31774:36;;31737:92;31873:4;:12;;31881:4;31873:12;:28;;;;31889:4;:12;;31897:4;31889:12;31873:28;31869:347;;;31929:4;:12;;31937:4;31929:12;31925:75;;;31974:3;;;31925:75;32021:8;32055:113;32065:10;32062:13;;:18;32055:113;;32137:8;32108:3;32137:8;;;;;32108:3;32055:113;;;32196:1;-1:-1:-1;32189:8:0;;-1:-1:-1;;32189:8:0;31869:347;28341:6215;;;32267:4;32258:6;:13;;;32254:2302;;;32309:6;:14;;32319:4;32309:14;32305:1088;;;32350:42;32368:2;32373:1;32368:6;32378:1;32367:12;32362:2;:17;32354:26;;:3;:26;;;;32384:4;32353:35;32390:1;32350:2;:42::i;:::-;32343:49;;;;;;32305:1088;32447:6;:14;;32457:4;32447:14;32443:950;;;32488:45;32506:2;32511:1;32506:6;32516:1;32505:12;32500:2;:17;32492:26;;:3;:26;;;;32522:6;32491:37;32530:2;32488;:45::i;32443:950::-;32589:6;:14;;32599:4;32589:14;32585:808;;;-1:-1:-1;;;32636:21:0;32655:1;32650;32645:6;;32644:12;32636:21;;32689:36;;;32756:5;32751:10;;32636:21;;;;;32750:18;32743:25;;32585:808;32823:6;:14;;32833:4;32823:14;32819:574;;;32864:3;32857:10;;;;;;32819:574;32923:6;:14;;32933:4;32923:14;32919:474;;;32979:2;32984:1;32979:6;32989:1;32978:12;32973:2;:17;32965:26;;:3;:26;;;;32995:4;32964:35;32957:42;;;;;;32919:474;33055:6;:14;;33065:4;33055:14;33051:342;;;33111:2;33116:1;33111:6;33121:1;33110:12;33105:2;:17;33097:26;;:3;:26;;;;33127:6;33096:37;33089:44;;;;;;33051:342;33189:6;:14;;33199:4;33189:14;33185:208;;;-1:-1:-1;;;33236:26:0;33260:1;33255;33250:6;;33249:12;33244:2;:17;33236:26;;33294:41;;;33366:5;33361:10;;33236:26;;;;;33360:18;33353:25;;32254:2302;33435:6;:14;;33445:4;33435:14;33431:1125;;;-1:-1:-1;;;33484:4:0;33478:34;33510:1;33505;33500:6;;33499:12;33494:2;:17;33478:34;;33560:27;;;33540:48;;;33610:10;;33479:9;;;33478:34;;33609:18;33602:25;;33431:1125;33670:6;:14;;33680:4;33670:14;33666:890;;;-1:-1:-1;;;33719:6:0;33713:36;33747:1;33742;33737:6;;33736:12;33731:2;:17;33713:36;;33797:29;;;33777:50;;;33849:10;;33714:11;;;33713:36;;33848:18;33841:25;;33666:890;33910:6;:14;;33920:4;33910:14;33906:650;;;-1:-1:-1;;;33953:20:0;33971:1;33966;33961:6;;33960:12;33953:20;;34001:36;;;34065:5;34059:11;;33953:20;;;;;34058:19;34051:26;;33906:650;34120:6;:14;;34130:4;34120:14;34116:440;;;34157:2;34150:9;;;;;;34116:440;34203:6;:14;;34213:4;34203:14;34199:357;;;-1:-1:-1;;;34246:25:0;34269:1;34264;34259:6;;34258:12;34253:2;:17;34246:25;;34299:41;;;34368:5;34362:11;;34246:25;;;;;34361:19;34354:26;;34199:357;34423:6;:14;;34433:4;34423:14;34419:137;;;34460:3;34453:10;;;;;;34419:137;34506:6;:14;;34516:4;34506:14;34502:54;;;34543:2;34536:9;;;;;;34502:54;34566:29;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;16924:688;17010:12;17069:18;;:::i;:::-;-1:-1:-1;17129:4:0;17224:2;17212:14;;;;17204:41;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;17333:14;;;;;;;:30;;;17351:12;17333:30;17329:94;;;17408:4;17379:5;:15;;;17395:9;17379:26;;;;;;;;;:33;;;;:26;;;;;;:33;17329:94;17470:12;;;;;17459:23;;;;:8;;;:23;17522:1;17507:16;;;17492:31;;;17592:13;:11;:13::i;4437:6628::-;4480:12;4538:18;;:::i;:::-;-1:-1:-1;4696:15:0;;:18;;;;4598:4;4840:18;;;;4880;;;;4920;;;;;4598:4;;4676:17;;;;4840:18;4880;5002;;;5016:4;5002:18;4998:5795;;;5048:2;5071:4;5068:7;;:12;5064:112;;5156:4;5153:7;;5145:4;:16;5139:22;5064:112;5193:7;;;5189:141;;5225:10;;;;;5253:16;;;;;;;;5225:10;-1:-1:-1;5189:141:0;;;5313:2;5308:7;;5189:141;4998:5795;;;;5434:10;:18;;5448:4;5434:18;5430:5363;;;1690:10;5468:14;;5430:5363;;;5554:10;:18;;5568:4;5554:18;5550:5243;;;5593:1;5588:6;;5550:5243;;;5706:10;:18;;5720:4;5706:18;5702:5091;;;5755:4;5740:12;;;:19;5773:26;;;:14;;;:26;5820:13;:11;:13::i;:::-;5813:20;;;;;;;;;;;5702:5091;5944:10;:18;;5958:4;5944:18;5940:4853;;;6083:14;;;6079:2210;;;;;6237:22;;;1923:1;6237:22;6233:2056;;;6354:10;6367:27;6375:2;6380:10;6375:15;6392:1;6367:7;:27::i;:::-;6453:11;6484:6;;6504:17;;;;6523:20;;;;;6484:60;;;;;;;;;;;;;;;;;;;;6354:40;;-1:-1:-1;6453:11:0;;;;6484:6;;;;;:19;;:60;;;;;;;;;;;:6;:60;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;6484:60:0;;;;;;;;;-1:-1:-1;6484:60:0;-1:-1:-1;6755:1:0;6747:10;;6845:1;6841:17;;;6916;;;6913:2;;;6946:5;6936:15;;6913:2;;7025:6;7021:2;7018:14;7015:2;;;7045;7035:12;;7015:2;7147:3;7142:1;7134:6;7130:14;7125:3;7121:24;7117:34;7110:41;;7243:3;7239:1;7227:9;7218:6;7215:1;7211:14;7207:30;7203:38;7199:48;7192:55;;7363:1;7359;7355;7343:9;7340:1;7336:17;7332:25;7328:33;7324:41;7486:1;7482;7478;7469:6;7457:9;7454:1;7450:17;7446:30;7442:38;7438:46;7434:54;7416:72;;7582:10;7578:15;7572:4;7568:26;7560:34;;7694:3;7686:4;7682:9;7677:3;7673:19;7670:28;7663:35;;;;7828:33;7837:2;7842:10;7837:15;7854:1;7857:3;7828:8;:33::i;:::-;7879:20;;;:38;;;;;;;;;-1:-1:-1;6233:2056:0;;-1:-1:-1;;6233:2056:0;;8020:18;;;1842:1;8020:18;8016:273;;;8178:2;8173:7;;8016:273;;;8236:10;8231:15;;1998:3;8264:10;;8016:273;5940:4853;;;8401:10;:18;;8415:4;8401:18;8397:2396;;;8543:15;;;1769:1;8543:15;;:34;;-1:-1:-1;8562:15:0;;;1804:1;8562:15;8543:34;:57;;;-1:-1:-1;8581:19:0;;;1881:1;8581:19;8543:57;8539:1505;;;8625:2;8620:7;;8539:1505;;;8739:23;;;1966:1;8739:23;8735:1309;;;8782:10;8795:27;8803:2;8808:10;8803:15;8820:1;8795:7;:27::i;:::-;8894:17;;;;8782:40;;-1:-1:-1;9121:1:0;9113:10;;9211:1;9207:17;9282:13;;;9279:2;;;9304:5;9298:11;;9279:2;9578:14;;;9392:1;9574:22;;;9570:32;;;;9471:26;9495:1;9384:10;;;9475:18;;;9471:26;9566:43;9380:20;;9670:12;9786:17;;;:23;9850:1;9827:20;;;:24;9388:2;-1:-1:-1;9388:2:0;8735:1309;;8397:2396;10122:10;:18;;10136:4;10122:18;10118:675;;;10208:2;:7;;10214:1;10208:7;10204:579;;;10277:14;;;;;:40;;-1:-1:-1;10295:22:0;;;1923:1;10295:22;10277:40;:62;;;-1:-1:-1;10321:18:0;;;1842:1;10321:18;10277:62;10273:376;;;10368:1;10363:6;;10273:376;;;10410:15;;;1769:1;10410:15;;:34;;-1:-1:-1;10429:15:0;;;1804:1;10429:15;10410:34;:61;;;-1:-1:-1;10448:23:0;;;1966:1;10448:23;10410:61;:84;;;-1:-1:-1;10475:19:0;;;1881:1;10475:19;10410:84;10406:243;;;10523:1;10518:6;;10406:243;;10204:579;10692:10;10687:15;;2032:4;10720:11;;10204:579;10860:15;;;;;:23;;;;:18;;;;:23;;;;10893:15;;:23;;;:18;;;;:23;-1:-1:-1;10974:12:0;;;;10963:23;;;:8;;;:23;11026:1;11011:16;10996:31;;;;;11045:13;:11;:13::i;:::-;11038:20;;4437:6628;;;;;;;;:::o;13505:2222::-;13599:12;13657:18;;:::i;:::-;-1:-1:-1;13717:4:0;13741:10;13850:4;13841:13;;;;13837:1545;;;-1:-1:-1;13876:8:0;;;;13837:1545;;;13983:5;:13;;13992:4;13983:13;13979:1403;;;14012:14;;;:8;;;:14;13979:1403;;;14130:5;:13;;14139:4;14130:13;14126:1256;;;-1:-1:-1;14165:8:0;;;;14126:1256;;;14272:5;:13;;14281:4;14272:13;14268:1114;;;14301:14;;;:8;;;:14;14268:1114;;;14430:5;:13;;14439:4;14430:13;14426:956;;;14549:9;14499:17;14479;;;14499;;;;14479:37;14556:2;14549:9;;;;;14531:8;;;:28;14573:22;:8;;;:22;14426:956;;;14720:5;:13;;14729:4;14720:13;14716:666;;;14783:11;14769;;;14783;;;14769:25;14834:2;14827:9;;;;;14809:8;;;:28;14851:22;:8;;;:22;14716:666;;;15012:5;:13;;15021:4;15012:13;15008:374;;;15078:3;15059:23;;15065:3;15059:23;;;;;;;;15041:42;;:8;;;:42;15115:23;;;;;;;;;;;;;;15097:42;;:8;;;:42;15008:374;;;15288:5;:13;;15297:4;15288:13;15284:98;;;15334:3;15328:9;;:3;:9;;;;;;;;15317:20;;;;:8;;;:20;15362:9;;;;;;;;;;;;15351:20;;:8;;;:20;15284:98;15467:14;;;;15463:77;;15526:3;15497:5;:15;;;15513:9;15497:26;;;;;;;;;:32;;;;:26;;;;;;:32;15463:77;15586:12;;;;;15575:23;;;;:8;;;:23;15638:1;15623:16;;;15608:31;;;15707:13;:11;:13::i;:::-;15700:20;13505:2222;-1:-1:-1;;;;;;;13505:2222:0:o;20607:1584::-;20755:14;20772:24;20784:11;20772;:24::i;:::-;20755:41;;20892:1;20885:5;20881:13;20878:2;;;20923:1;20920;20913:12;20878:2;21062;21244:15;;;21081:2;21070:14;;21058:10;21054:31;21051:1;21047:39;21204:16;;;21001:20;;21189:10;21178:22;;;21174:27;21164:38;21161:60;21650:5;21647:1;21643:13;21713:1;21698:375;21723:2;21720:1;21717:9;21698:375;;;21838:2;21826:15;;;21779:20;21869:12;;;21883:1;21865:20;21902:78;;;;21986:1;21981:78;;;;21858:201;;21902:78;21939:23;21954:7;21948:4;21939:23;:::i;:::-;21931:31;;21902:78;;21981;22018:23;22036:4;22027:7;22018:23;:::i;:::-;22010:31;;21858:201;-1:-1:-1;;21741:1:0;21734:9;21698:375;;;-1:-1:-1;;22164:4:0;22157:18;-1:-1:-1;;;;20816:1369:0:o;17816:500::-;18106:20;;;18130:7;18106:32;18099:3;:40;;;18188:14;;18227:17;;18221:24;;;18213:72;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;18295:14;17816:500;;;:::o;-1:-1:-1:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o"

func init() {
	if err := json.Unmarshal([]byte(MIPSStorageLayoutJSON), MIPSStorageLayout); err != nil {
		panic(err)
	}

	layouts["MIPS"] = MIPSStorageLayout
	deployedBytecodes["MIPS"] = MIPSDeployedBin
}
