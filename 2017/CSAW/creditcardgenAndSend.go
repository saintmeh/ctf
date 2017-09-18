package main
import (
    "net"
    "os"
    "encoding/json"
    "fmt"
    "strings"
    "strconv"
    "time"
    "math/rand"
)

type creditcard struct{
    CreditCard cc
}
type cc struct{
    IssuingNetwork string
    CardNumber uint64
}

func main() {
    
    strEcho := "-111\n"
    servAddr := "misc.chal.csaw.io:8308"
    visacards := []byte(`[
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716171778596672
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4056568785514317
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716385349433013
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916851015480941
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007148619629
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916650478316877
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716685230429057
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716804079337020
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929027778442908
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007194075395
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539594314795422
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532397541942701
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532222712664423
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556968156124578
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916449381520638
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4963627584461542
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929165160635026
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4500310738957554
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716139209591117
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716379792420700
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556987840866584
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4998651999893957
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929435963204829
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007170817331
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539211102241820
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4920215349328461
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556755438534884
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007109791706
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929593301065862
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4552001100525791
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716706169907218
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007198337536
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929705901470678
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929098437981909
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007171879066
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556164548627771
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539953185663343
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007153319099
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007104111330
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539284119198879
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556837483380304
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4485523724940280
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4485971913873236
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007140827667
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916071267450233
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716173649358621
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4683830523908693
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539328647044731
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716072889728883
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4485572566658203
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007181471136
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532464153119018
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556870085216139
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929511646057766
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929902737645710
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539860276843447
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4485840673225459
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916344577837344
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532116720959892
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929764865183314
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556001432004248
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007151789889
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539678070445403
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916231381637863
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916762727650095
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007166316710
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916748317863685
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4929786222614145
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556255327723201
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4485331149114667
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007132581256
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532543918898873
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007148822652
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532955894269673
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716065259112647
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4485757375531283
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539225437035194
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716974992015746
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007176989621
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539388743565948
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556482278566310
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916543109747877
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4847370131543627
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916906481969528
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539037034834733
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4933207070003897
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532100156522853
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4916428282865672
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007138977144
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539930845093046
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532862467375863
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532205888569609
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532049655268890
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539349201424739
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4024007169853719
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4716370105668895
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4539598276664815
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4556055610286738
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4532076907158074
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Visa",
      "CardNumber": 4485872998536716
    }
  }
]`)
    mastercards :=  []byte(`[
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5404635557914680
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5201343630616043
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5527569357265112
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5562949328404576
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5478012881900408
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5266539001219075
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5246556483435479
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5560610124059038
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5115969190043189
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5234707657232184
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5290399429948585
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5535329182643458
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5281176793383222
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5167436979286378
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5244564244421459
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5536800759791549
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5156645104140630
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5250424327207559
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5295595248430664
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5376553267976846
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5108632430950980
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5474952960172748
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5102490132678524
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5571574131437321
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5363482238445910
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5242519296101062
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5144553311595635
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5413732136800297
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5131804948173162
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5337320278554352
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5188157619664775
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5202020119545377
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5472555092242974
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5283777399615954
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5390638960668659
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5495547897371317
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5402382778820345
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5418364509020306
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5218130418899359
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5313353019799612
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5469938849766609
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5129275396894126
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5415680830192446
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5119455861251687
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5351255279883148
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5434715640074884
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5121807174317942
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5268043099983294
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5126805545743726
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5136885650008931
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5112581044399742
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5311531010828694
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5129533560251089
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5159304449965858
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5210231894698998
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5436098316793642
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5348873258808138
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5358267074435829
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5249094320853701
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5508498170097895
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5470257282667803
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5432010845841440
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5312164358730014
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5501934430685183
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5358799162150498
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5482805528143389
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5201422533212926
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5476677149664665
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5392954218461351
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5262365156031402
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5277046907043416
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5448051001887431
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5429164258734413
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5452579022110774
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5178489383855145
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5305495208016835
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5400009385126269
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5521581555775715
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5155351280773077
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5229093299387091
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5131071275366013
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5249323666139908
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5301256375228437
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5403494298790996
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5420450107071977
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5253100910696963
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5406940668321751
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5177676227632422
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5150141820802279
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5200040415152985
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5307718269067118
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5240628453936426
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5398468656629548
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5592322414130626
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5309743711352916
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5568993326250007
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5368104071145523
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5495400808258753
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5235686705259796
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "MasterCard",
      "CardNumber": 5289721656531375
    }
  }
]`)
    amexcards := []byte(`[
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 374481346455603
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 373612231816969
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 370031559651713
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 341167043213144
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 373746940176990
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 376520018856263
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 349393064479123
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 343361259906339
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 377827484456281
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 342859953293467
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 374911522721329
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 346704341798429
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 373194379331721
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340438389982605
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 349074841283398
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 343467857564899
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 344548901790314
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340715036929778
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 370254223926655
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 371291041468493
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 371091708205893
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 342633412800547
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 378994648409413
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 373350918810738
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 379557176774901
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 342475012865714
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 348761036794624
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 379808513368216
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 376463043838289
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340211224051609
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 349922075743182
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 341351119177348
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 346017087224025
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 341689705990628
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 345864000191844
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340504125284182
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 375307595779014
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340171391089008
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 377765488149744
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 375758868798794
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 345684594424116
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 377459198648043
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 373712263796988
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 343912758772059
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 342816826716561
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 377549476560892
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340651683321854
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 374300198954366
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 371190580437549
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 343627998234724
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 341651738879962
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 374938003299392
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 377038413222352
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 341349125898550
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 347944825829972
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 379955257923654
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 375477597904447
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 344927377941546
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 349405272275869
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 370516408513064
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 376203368434540
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 378027689094255
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 345530159502913
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 376090925936665
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 345884200241884
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 346625460379578
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 372990938003227
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 344626269577588
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 376834609274079
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340372998270146
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 372291493842368
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 378578575645048
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 373087422958740
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 372080765831450
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 371582480139334
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 376880862422657
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 375508979015970
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 341783924451585
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 372856883289659
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 376218342385501
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340398057678986
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 347556125097082
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340483110332563
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 372934925746455
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 347619634668445
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 377504112127614
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 373909484926051
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 377277783118671
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 346353653506309
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 378131196658345
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 343761246210192
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 375996628228252
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340787920096832
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 346279535075975
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 376109036730559
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 340352568551950
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 344279402627207
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 344982228482786
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 372345695178631
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "American Express",
      "CardNumber": 343898652350235
    }
  }
]`)
    discovercards := []byte(`[
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011020999535349
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011668751451319
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011264146980848
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011475079053490
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011180886336704
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011040538078494
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011776610659053
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011688414289864
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011202327666687
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011179639832027
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011150168430145
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011815404673670
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011562622030826
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011116234728705
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011329612550753
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011091513868623
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011776395564338
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011389116126427
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011712070223551
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011431392972775
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011927614894549
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011039739676239
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011963311485570
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011047304463206
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011828191306289
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011342462986191
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011355741446621
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011267580966181
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011762833656241
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011421243735141
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011329497241908
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011669391322662
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011405386699933
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011362321771998
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011860710685281
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011291154087976
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011137454005269
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011828592883720
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011660954639048
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011455584334627
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011078080784374
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011441512545484
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011039153626124
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011375505757906
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011896936015150
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011398828975781
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011129568980867
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011686763042314
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011869623979368
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011733393388111
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011130707029749
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011855362066442
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011498678969577
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011550394143695
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011690683840641
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011380111605788
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011865414477786
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011781818298134
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011612980031583
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011569119888632
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011315307911665
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011240058943222
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011812556324277
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011078254936776
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011576095580208
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011554134267319
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011046099344407
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011283629114618
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011996667711132
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011460462981879
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011702371302428
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011762292830881
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011975993430880
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011584039511875
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011474705320035
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011823027849586
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011133660651707
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011895884534022
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011279333237828
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011779282512196
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011456482853254
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011429216158175
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011971441939895
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011262171522774
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011349551552396
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011079459693519
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011717099423496
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011431591846143
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011000220534279
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011981741249945
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011829272878352
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011575594372125
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011854362316410
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011593248565358
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011298329507952
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011198168732085
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011141003646896
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011665704379746
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011742613455212
    }
  },
  {
    "CreditCard": {
      "IssuingNetwork": "Discover",
      "CardNumber": 6011872289161461
    }
  }
]`)


    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }


    reply := make([]byte, 1024)

    var vcs []creditcard
    err = json.Unmarshal(visacards, &vcs)
    var mcs []creditcard
    err = json.Unmarshal(mastercards, &mcs)
    var acs []creditcard
    err = json.Unmarshal(amexcards, &acs)
    var dcs []creditcard
    err = json.Unmarshal(discovercards, &dcs)

mcount := 0
vcount := 0
dcount := 0
acount := 0

for i:=0; i<10000; i++ {
    _, err = conn.Read(reply)
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }
    println("reply1 from server=", string(reply))

    println("Check Card")
    if strings.Contains(string(reply), "MasterCard") {

        println("writing mastercard:")
	    strEcho = strconv.FormatUint(mcs[mcount].CreditCard.CardNumber,10)+"\n" 
	    _, err = conn.Write( []byte(strEcho) )
	    if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	    }
		mcount++
	    println("write to server = ", strEcho)
    }
    if strings.Contains(string(reply), "Visa") {
        println("writing visa:")
	    strEcho = strconv.FormatUint(vcs[vcount].CreditCard.CardNumber,10)+"\n" 
	    _, err = conn.Write( []byte(strEcho) )
	    if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	    }
		vcount++

	    println("write to server = ", strEcho)
    }
    if strings.Contains(string(reply), "American") {
        println("writing amex:")
	    _, err = conn.Write( []byte(
    strconv.FormatUint(acs[acount].CreditCard.CardNumber,10)+"\n" ) )
	    if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	    }
		acount++

	    println("write to server = ", strEcho)
    }
    if strings.Contains(string(reply), "Discover") {
        println("writing discover:")
	    _, err = conn.Write( []byte(
    strconv.FormatUint(dcs[dcount].CreditCard.CardNumber,10)+"\n" ) )
	    if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	    }
		dcount++

	    println("write to server = ", strEcho)
    }
    
    if strings.Contains(string(reply), "I need a new card that starts with") {
	
	starting := reply[:len(reply)-4]
        starting = starting[len("I need a new card that starts with ")+8:]
	startingstr := strings.Split(string(starting), "!")[0]
        println("writing computed starting with \""+ startingstr+"\"" )
	strEcho = GenerateWithPrefix(16,startingstr )+"\n"
	println(" as ", strEcho, ":")
	    _, err = conn.Write( []byte( strEcho ) )
	    if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	    }
		dcount++

	    println("write to server = ", strEcho)
    }
    if strings.Contains(string(reply), "I need a new card which ends with") {
	targetSuffix := reply[:len(reply)-4]
        targetSuffix = targetSuffix[len("I need a new card which ends with")+8:]
	targetSuffixstr := strings.TrimSpace(strings.Split(string(targetSuffix), "!")[0])
	strEcho=GenerateWithSuffix(16,targetSuffixstr)+"\n";
	println("write to server = ", strEcho)
	_, err = conn.Write( []byte( strEcho ) )

    }

    if strings.Contains(string(reply),"I need to know if ") {
	number := strings.TrimSpace(strings.Split(string(reply), "I need to know if ")[1])
	number = strings.TrimSpace(strings.Split(string(number), "is valid")[0])
	println(".",number,".")
	if Valid(string(number))&& len(string(number)) == 16 {
		strEcho = "1\n"
	} else {
		strEcho = "0\n"
	}
	println("write to server = ", strEcho)
	println("write to server = ", strEcho)
	_, err = conn.Write( []byte( strEcho ) )
    }
}
    conn.Close()
//    fmt.Printf("%+v", cc)
    fmt.Printf("")
    println(mastercards)
    println(amexcards)
    println(discovercards)

    

    

}




// Valid returns a boolean indicating if the argument was valid according to the Luhn algorithm.
func Valid(luhnString string) bool {
	checksumMod := calculateChecksum(luhnString, false) % 10

	return checksumMod == 0
}

// Generate creates and returns a string of the length of the argument targetSize.
// The returned string is valid according to the Luhn algorithm.
func Generate(size int) string {
	random := randomString(size - 1)
	controlDigit := strconv.Itoa(generateControlDigit(random))

	return random + controlDigit
}

// GenerateWithPrefix creates and returns a string of the length of the argument targetSize
// but prefixed with the second argument.
// The returned string is valid according to the Luhn algorithm.
func GenerateWithPrefix(size int, prefix string) string {
	size = size - 1 - len(prefix)

	random := prefix + randomString(size)
	controlDigit := strconv.Itoa(generateControlDigit(random))

	return random + controlDigit
}

// GenerateWithSuffix creates and returns a string of the length of the argument targetSize
// but suffixed with the second argument.
// The returned string is valid according to the Luhn algorithm.
func GenerateWithSuffix(size int, suffix string) string {
	presuffix := suffix[:len(suffix)-1]
	println(presuffix)
	controlDigit := suffix[len(suffix)-1:]
	println(controlDigit)
	digitsToGenerate := size - len(suffix)
	returnValue := "-1"
	for i:=0;i<1000;i++ {
		randomPrefix := randomString(digitsToGenerate)
		calculatedControlDigit := strconv.Itoa(generateControlDigit(randomPrefix+presuffix))
		println(randomPrefix,calculatedControlDigit, controlDigit, suffix, presuffix)
		if randomPrefix[:1]!="0" && calculatedControlDigit == controlDigit &&  Valid(randomPrefix+suffix){
			returnValue=randomPrefix+suffix
			break
		}
	}
	return returnValue
}

func randomString(size int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	source := make([]int, size)

	for i := 0; i < size; i++ {
		source[i] = rand.Intn(9)
	}

	return integersToString(source)
}

func generateControlDigit(luhnString string) int {
	controlDigit := calculateChecksum(luhnString, true) % 10

	if controlDigit != 0 {
		controlDigit = 10 - controlDigit
	}

	return controlDigit
}

func calculateChecksum(luhnString string, double bool) int {
	source := strings.Split(luhnString, "")
	checksum := 0

	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checksum += n
	}

	return checksum
}

func integersToString(integers []int) string {
	result := make([]string, len(integers))

	for i, number := range integers {
		result[i] = strconv.Itoa(number)
	}

	return strings.Join(result, "")
}
