package art_link

import (
    "bufio"
    "fmt"
    "strings"
    "time"
    "embed"
    "github.com/nsf/termbox-go"

)
//go:embed *.txt
var artFiles embed.FS



var frame2 =  `
               _-o#&&*''''?d:>b\\_
          _o/\"      ,, dMF9MMMMMHo_
       .o&#'         "MbHMMMMMMMMMMMHo.
     .o\"\" '         vodM*$&&HMMMMMMMMMM?.
    ,'              $M&ood,~' (&##MMMMMMH\\
   /               ,MMMMMMM#b?#bobMMMMHMMML
  &              ?MMMMMMMMMMMMMMMMM7MMM$R*Hk
 ?$.            :MMMMMMMMMMMMMMMMMMM/HMMM| *L
|               |MMMMMMMMMMMMMMMMMMMMbMH'   T,
$H#:             *MMMMMMMMMMMMMMMMMMMMb#}'   ?
]MMH#             \\ \\*\"*#MMMMMMMMMMMMM'    -
MMMMMb_                   |MMMMMMMMMMMP'     :
HMMMMMMMHo                  MMMMMMMMMT       .
?MMMMMMMMP                  9MMMMMMMM}       -
-?MMMMMMM                  |MMMMMMMMM?,d-    '
 :|MMMMMM-                  MMMMMMMT .M|.   :
  .9MMM[                    &MMMMM*'  '    .
   :9MMk                     MMM#\"        -
     &M}                      M          .-
       &.                             .
         ~,   .                     ./
            . _                  .-
              ' --._,dd###pp=\"\"'

`
var frame3 = `
               _v->#H#P? \"':o<>\\_
          .,dP   ''  \"'-o.+H6&MMMHo_
        oHMH9'          ?&bHMHMMMMMMHo.
      oMP\"' '           ooMP*#&HMMMMMMM?.
    ,M*          -      *MSdob// ^&##MMMH\\
   d*'                .,MMMMMMH#o>#ooMMMMMb
  HM-                :HMMMMMMMMMMMMMMM&HM[R\\
 d\"Z\\.               9MMMMMMMMMMMMMMMMM[HMM|:
-H    -              MMMMMMMMMMMMMMMMMMMbMP' :
:??Mb#                9MMMMMMMMMMMMMMMMMMH#! .
: MMMMH#,              \"*MMMMMMMMMMMMMMMMH  -
||MMMMMM6\\.                   \MMMMMMMMMH'  :
:|MMMMMMMMMMHo                  9MMMMMMMM'   .
. HMMMMMMMMMMP'                 !MMMMMMMM
-  #MMMMMMMMM                   HMMMMMMM*,/  :
 :  ?MMMMMMMF                   HMMMMMM',P' :
  .  HMMMMR'                    {MMMMP' ^' -
   :  HMMMT                     iMMH'     .'
    -. HMH                               .
      -:*H                            . '
        -\\,,    .                  .-
          ' .  _                 .-
                \\.__,obb#q==~'''

`

var frame4 =  `
               .ovr:HMM#?:' >b\\_
          .,:&Hi' '   \"' \\\\|&bSMHo_
        oHMMM#*}           ?&dMMMMMMHo.
     .dMMMH\"''''           ,oHH*&&9MMMM?.
    ,MMM*'                 *M\\bd<|\"*&#MH\\
   dHH?'                   :MMMMMM#bd#odMML
  H' |\\                  dMMMMMMMMMMMMMM9Mk
 JL/\"7+,.                MMMMMMMMMMMMMMMH9ML
-Hp     '               |MMMMMMMMMMMMMMMMHH|:
:  \\\\#M#d?                HMMMMMMMMMMMMMMMMH.
.   JMMMMM##,                *\"\"'\"*#MMMMMMMMH
-. ,MMMMMMMM6o_                    |MMMMMMMM':
:  |MMMMMMMMMMMMMb\\                 TMMMMMMT :
.   ?MMMMMMMMMMMMM'                 :MMMMMM|.
-    ?HMMMMMMMMMM:                  HMMMMMM\\|:
 :     9MMMMMMMMH'                  MMMMMP.P.
  .     MMMMMMT''                   HMMM*''-
   -    TMMMMM'                     MM*'  -
    '.   HMM#                            -
      -.  9M:                          .'
        -.  b,,    .                . '
          '-\\   .,               .-
              '-:b~\\\\_,oddq==--\"

`

var frame5 =  `
              _oo##'9MMHb':'-,o_
          .oH\":HH$' \"\"'  \"' -\\7*R&o_
       .oHMMMHMH#9:          \"\\bMMMMHo.
      dMMMMMM*\"\"' '           .oHM\"H9MM?.
    ,MMMMMM'                   \"HLbd<|?&H\\
   JMMH#H'                     |MMMMM#b>bHb
  :MH  .\"\\                    |MMMMMMMMMMMM&
 .:M:d-\"|:b..                 9MMMMMMMMMMMMM+
:  \"*H|      -                &MMMMMMMMMMMMMH:
.     LvdHH#d?                 ?MMMMMMMMMMMMMb
:      iMMMMMMH#b               \"*\"'\"#HMMMMMM
.   . ,MMMMMMMMMMb\\.                   {MMMMMH
-     |MMMMMMMMMMMMMMHb,                MMMMM|
:      |MMMMMMMMMMMMMMH'                &MMMM,
-        #MMMMMMMMMMMM                 |MMMM6-
 :         MMMMMMMMMM+                 ]MMMT/
  .        MMMMMMMP\"                   HMM*
   -       |MMMMMH'                   ,M#'-
    '.     :MMMH|                       .-
      .     |MM                        -
         .    #?..    .             ..'
           -.     _.             .-
              '-|.#qo__,,ob=~~-''

`

var frame6 =  `
            _ooppH[ MMMD::-zzzzz-\\_
         oHMMMR:\"&MZ\\  \"'  \"  |$-_
       ..dMMMMMMMMdMMM#9\\         'HHo.
     . ,dMMMMMMMMMMM\" '            ?MP?.
    . |MMMMMMMMMMM                  \"$b&\\
   -  |MMMMHH##M                      HMMH?
  -   TTMM|    >..                   \\MMMMMH
 :     |MM\\,#-\"\"$~b\\.              MMMMMM+
.         \"H&#        -               &MMMMMM|
:            *\\v,#MHddc.               9MMMMMb
.               MMMMMMMM##\\            \"\":HM
-          .  .HMMMMMMMMMMRo_.               |M
:             |MMMMMMMMMMMMMMMM#\\           :M
-               HMMMMMMMMMMMMMMM'           |T
:                *HMMMMMMMMMMMM'            H'
 :                 MMMMMMMMMMM|            |T
  .                MMMMMMMM?'             ./
   .               MMMMMMH'              ./
    -.            |MMMH#'                .
      .            MM*                . '
        -.         #M: .    .       .-
            .         .,         .-
              '-.-~ooHH__,,v~--

`
var frame7 =  `
              _ood>H&H&Z?#M#b-\\.
          \\HMMMMMR?\\M6b.\"' ''  v.
       .. .MMMMMMMMMMHMMM#&.        ~o.
     .   ,HMMMMMMMMMMMM*\"'-           &b.
    .   .MMMMMMMMMMMMH'                \"&\
   -     RMMMMM#H##R'                   4Mb
  -      |7MMM'    ?::                  |MMb
 /         HMM__#|\"\\>?v..               MMML
.            \"'#Hd|                     9MMM:
-                |\\,\\?HH#bbL              9MM
:                   !MMMMMMMH#b,           \"\"
.              .   ,MMMMMMMMMMMbo.           |
:                  4MMMMMMMMMMMMMMMHo        |
:                   ?MMMMMMMMMMMMMMM?        :
-.                    #MMMMMMMMMMMM:        .-
 :                     |MMMMMMMMMM?         .
  -                    JMMMMMMMT'          :
   .                   MMMMMMH'           -
    -.                |MMM#*             -
      .               HMH'            .
        -.            #H:.          .-
            .           .\\       .-
              '-..-+oodHL_,--/-
`
var frame8 =  `
               _,\\?dZkMHF&$*q#b..
          .//9MMMMMMM?:'HM\\\\\" -'' ..
       ..   :MMMMMMMMMMHMMMMH?_     -\\
     .     .dMMMMMMMMMMMMMM'\"'\"       \\.
    .      |MMMMMMMMMMMMMR              \\\\
   -        T9MMMMMHH##M\"                 ?
  :          (9MMM'    !':.                 &k
 .:            HMM\\_?p \"\":-b\\.           ML
-                \"'\"H&#,       :           |M|
:                     ?\\,\\dMH#b#.           9
:                        |MMMMMMM##,          *
:                   .   +MMMMMMMMMMMo_         -
:                       HMMMMMMMMMMMMMM#,     :
:                        9MMMMMMMMMMMMMH'     .
: .                       *HMMMMMMMMMMP     .'
 :                          MMMMMMMMMH'      .
  -                        :MMMMMMM'       .
   .                       9MMMMM*'        -
    -.                    {MMM#'          :
      -                  |MM\"          .'
        .                &M'..  .   ..'
          ' .             ._     .-
              '-. -voboo#&:,-.-

`

var frame9 =  `
               _oo:\\bk99M[<$$+b\\.
           .$*\"MMMMMMMM[:\"\\Mb\\?^\" .
       . '     HMMMMMMMMMMHMMMM+?.   .
     .        .HMMMMMMMMMMMMMMP\"''     .
    .          MMMMMMMMMMMMMM|         - .
   -            &MMMMMMHH##H:             :
  :              (*MMM}     |\\             |
 :  -              ?MMb__#|\"\" |+v..         \
.                     ''*H#b       -        :|
:                          *\\v,#M#b#,        \
.                             9MMMMMMHb.     :
:                        .   #MMMMMMMMMb\\    -
-                           .HMMMMMMMMMMMMb  :
:                             MMMMMMMMMMMMH  .
-:  .                          #MMMMMMMMMP   '
 :                              ]MMMMMMMH'  :
  -                            ,MMMMMM?'   .
   :                           HMMMMH\"    -
    -.                       .HMM#*     .-
      .                     .HH*'     .
        -.                  &R\".    .-
           -.               ._   .-
              '-. .voodoodc?..-

`
var frame10 =  `
              _\\oo\\?ddk9MRbS>v\\_
          ..:>*\"\"MMMMMMMMM:?|H?$?-.
       ..- -     \"HMMMMMMMMMMHMMMH\\_-.
     .            dMMMMMMMMMMMMMMT\"    .
    .             TMMMMMMMMMMMMMM        .
   -                &HMMMMMM#H#H:         .
  -                  \\7HMMH     |\\.        .
 :                      HMM\\_?c \"\"+?\\..
-                         \"  #&#|      .     -
:                               ?,\\#MHdb.    .
:                                 |MMMMMH#.  :
:                            .   ,HMMMMMMMb, -
: '                              4MMMMMMMMMMH
:   .                             9MMMMMMMMMT-
:.                                 #MMMMMMMH '
 :      '                           HMMMMMH':
  -                                |MMMMH\" -
   :                              |MMMH*' .'
    '?                           dMM#'   .
      \\.                       .dH\"    .'
        -.                    ,M'-  ..'
            .                .. ..-
              '-. .\\ooooboo<^.-

`
var frame11 =  `
                 o,:o?\\?dM&MHcc~,.
          ..^':&#\"\"HMMMMMMMM$:?&&?.
        .   -       'HMMMMMMMMMHMMMp\\.
     . '             |MMMMMMMMMMMMMM\"' .
    .                 9MMMMMMMMMMMMM    -.
   -                    *9MMMMMHH##[      .
  -                      \\Z9MMM     ~\\     .
 :       '|                 ?MMb_?p\"\"-?v..  :
-                              \"'*&#,    -   .
:                                   ?,oHH#?  .
--                                    |MMMMH,:
:                                 .  |MMMMMM6,
:   -                                |MMMMMMMM
?                                     HMMMMMMP
-- . '                                |HMMMMM'
 :.      .  '                          JMMMM+
  \\                                   ,MMMP:
   :                                 |MMH?:
    -:\\.                            dM#\" .
       \\                          ,H*' .'
        -.                       d':..'
            .                  .,.-
              '-.. .\\oooodov~^-

`
var frame12 = `
             _o\\:,??\\??MR9#cb\\_
          .v/''':&#\"\"#HMMMMMMM$?*d\\.
       ..~' - -        \"#MMMMMMMMMMMHv.
     .-'                 HMMMMMMMMMMMR!.
    :                     9MMMMMMMMMMM| -.
   .                        *9MMMMMH##|   .
  -                           (#MMH    :,  .
 :           '|                  HMb_>/\"|\\,.
.'                                 \"'#&b   - .
:                                      ?\\oHH?.
:                                        !MMM&
:  .                                  .  HMMMM
/.      -                               -MMMMM
\\ .                                      9MMMP
:. .  . -                                |MMM'
 \\... '                                  .MMT
  &.                                    .dMP
   \\,                                  .HM*
    \\.  \\.                            ,H&'
      -  | -                        ,&':
        .                         ,/\\ '
          '-..                  _.-
              \"---.._\\o,oov+--'\"

`
var frame13 =  `
             _,d?,:?o?:?HM>#b\\_
          ..H*\"'' 'H#*\"**MMMMMM6$$v_
        v//\"   -          '#MMMMMMMMHo.
      /\"                    |MMMMMMMMMM:.
    ,>                        HMMMMMMMMH:.
   :                            #HMMMMHH\\ -
  '                               Z#MM,   ,:
 :               '\\                 ?HH_>: \\,
:                                     \"'*&|  :
.                                         <\\Hb
:                                           MM
:                                        . iMM
Mb\\.                                       {MM
::. -       -                              !MP
 &.   .  .  -                              :M'
 9H,  \\  '                                 |T
  HM?                                     ,P
   *ML                                   ??
    :&.    o                           .d'
      ':  |T                          /\"
        -.                         .<''
           ...                  ..-
              \" -=.,_,,,oov-~.-

`
var frame14 =  `
            _,oc>?_:b?o?HH#b\\_
          .v/99*\"\"\" '*H#\"\"*HMMMM
        oH* /\"   -   '      \" #MMMMM#
     ./>-                      MMMMMMMb
    ,b/'                         #MMMMMMM\\
   :'                               HMMMMb:
  /-                                 |&MH  \\
 /                    -.               |Hb??\\
,-  '                                    \" &,.
1                                           \\}
!.                                           T
$,.                                        . 1
? M??.                                       M
?.::| '\\        -                            ?
 M?&.    .   .  -                           ,'
 9MMH\\   ..  '                              .
  HMMM#.                                   :'
   9#MMb                                 ..
    -:\"#      b.                        .-
      .      {!                        /
        -                           ,-'
          ' .                    .-
                ^==\\_.,,,ov--\\-

`
var frame15 =  `
               _\\o##??,:io??$#b\\_
          .oH#\"H9*\"\"\" \" #H*\"*\
        oHMM- -'    -  ''       *HMMHo.
      dM#S>-                      ?MMMM?.
    ,&&,/'                         \"#MMMH\\
   d?-\"                               *HMMb
  H?                                   \"ZHb:
 /:                        \\              H?L
|:|   .                                     *:
:?:                                          \\
>\"                                           :
M|\\,_                                        |
!|\":HH?-'.                                   :
:^'_:?\"\\  --         -
- |ML?b      .   ..  -                       -
 :HMMMMH\\    \\                             :
  >MMMMMM#.                                .
   ^M*HMMM|                               -
     .  \"#+      ?v                     .
      .    -    +?'                    -
         .                          ..'
           - .                   .-
              \" \\b=p?.._\\\\vv---

`
var frame16 =  `
             _,o#bH\\??::?o?cbo_
          .o#MH#**SH\"\"' \" *H#\"*#MHo_
        oHMMMH^  ^\"    -         '*HHo.
     .dMMM#\">>-                      HM?.
    ,MH:R_o/                          *MH\\
   dMM' '                               \"ML
  HMR! '                                  #k
 d&'.                          -.           L
:M ::                                        -
/| !|                                        -
k.$-\"                                        :
}9R:!,,_.                                    .
\\::\\': *M#\\-'.                               -
: \"''..:\"! \\  '-          -
-   ,HMb.H|      .    _   -                 .'
 : ,MMMMMMMb.    ..                         .
  . HMMMMMMMM?                             .
   . 9M#*HMMMM                            :
    -.'   \"##*       b,                  .
      .            ,/'                 .'
         .                          ..'
           - .                  ..-
              \" *#d##c.._\\v----

`
var frame17 =  `
             _,o#&oHb?\\o::d?>\\_
          .oHHMMM#**$M\"\"  \" *HH\"#&o_
        oHMMMMMMD' .''    -  '      bo
     .dMMMMMH*'/|-                    \\b.
    ,MMMM?T|_o/                         \\\\
   dMMMMP  ''                              |
  HMMMH& -                                 \\
 /MH7' :                          --         :
-:MM  {.      .                              .
:i?' .!&                                      .
:{, o| '                                     :
-T?9M\\:-'o,_                                 .
: \\?::  \" ?9MHo./..                         -
.  '\" '^ _. \"!\"^.   -         -
-      ,bMM?.M\\       .    .  -      .      .'
 :   .oMMMMMMMMb.    ..                      .
  .   HMMMMMMMMMMb                         -
   -   9MH*#HMMMMH                        .'
    '.  '    \"*##'       b.              :
      .               .d''             .'
        -.                          . '
           -.                    .-
              \" *##H###:._\\--.-

`

var frame18 =  `
            _oo#H&d#b?\\b:_>>\\_
          .oHMMMMMMH*\"*9R\"'-  *#P\\-_
        oHMMMMMMMMM$  .\"       '    ^-
     .dMMMMMMMMH*\",?-                 '\\.
    ,MMMMMMM:?}.,d'                      .
   dMMMMMMMH  /''                         :
  HMMMMMMM&' -                             -
 dPTMMP>' :                           -.    :
|? -MM}  .\\                                  .
J' ::*'  -$L                                 .
:  ?b .,H- '                                 :
-  |6.&MP:: !.,_.                            -
:    \\:: \"' \" :\"MM#,-^,            -     :
-         :' _.:\"?  \\    -                 .
:         .?bMML.]#        -   _         .  .'
 -      .o#MMMMMMMMH\\     \\.          .   .
  -      HMMMMMMMMMMMH                     :
   .      HMM#*#MMMMMH'                   -
    -.     '      ##*'      i+           :
      -             '     v/'          .'
        -                           ..'
          ' .                    .-
              \" *##HMH##:__,-.-
 `

var frame19 =  `
             _oo##Mbb&bo??o_>\\_
          .oHMMMMMMMMM**#?M*' \"?*&..
        oHMMMMMMMMMMMM4   \"       -   .
     .dMMMMMMMMMMMM#\"\\?.-                .
    ,MMMMMMMMMM}\"9:_,d'                   -.
   dMMMMMMMMMMM|  ^''                       .
  &MMMMMMMMMMH\\  -                          .
 :{M*\"MMMPT\"' :                          -. :
.'M'  'MMM.  -T,       .                       .
- k   i:?''  -|&                               .
:    -o&  .,H- \"                              :
-      M: HMP|:'!.o._.                         .
:      \"<:::'<^ '\"  9MH#,-^ .                -
-         '''  ''._. \"? ^|   ^        -       :
:              ?#dMM_.M?       .   .  -     ..'
 :          ,ddMMMMMMMMMb.    ..   '         .
  .         TMMMMMMMMMMMMM,                 :
   -         ?MMH**#MMMMMH'                :
    '.        '     \" ##*'      &.       :
      -.                '    ,~\"       .'
        -.                          ..'
            .                    .-
                *##HMMMH#<:,..-

`
var frame20 =  `
	      _,dd#HMb&dHo?\\?:\\_
          .oHMMMMMMMMMMMH***9P' \ "\\v.
        oHMMMMMMMMMMMMMMM>   '         -.
     .dMMMMMMMMMMMMMMMH*'|~-'            .
    ,MMMMMMMMMMMMM6> H._,&                -.
   dMMMMMMMMMMMMMMM|   \"                   .
  H*MMMMMMMMMMMMMH&. -                       .
 d' HMM\"\"&MMMPT'' :.                       .-
,'  MP    TMMM,   |:        .                  -
|   #:    ? *\"   : &L                         :
!    '   /?H   ,#r  '                          :
.         ?M: HMM^<~->,o._                     :
:           9:::' *-  ': 9MHb,|-,           '  :
.              \"''':' :_ \"\"!\"^.   |        :
 .                 _dbHM6_|H.      .   .  '   .'
 \\              _odHMMMMMMMMH,    ..         :
  -             |MMMMMMMMMMMMM|              :
   .             9MMH**#MMMMMH'             :
    -.            '     \"?##\"      d     :
      .                    '    ,/\"    .'
        ..                          ..'
             .                   .-
              ' \"#HHMMMMM#<>..-

`
var frame21 =  `
              _oo##bHMb&d#bd,>\\_
          .oHMMMMMMMMMMMMMM***9R\"-..
        oHMMMMMMMMMMMMMMMMMH\\  ?    -.
     .dMMMMMMMMMMMMMMMMMMM#\".}-'       .
    ,MMMMMMMMMMMMMMMMM6/ H _o}          -.
   dMMMMMMMMMMMMMMMMMMML   ''             .
  HbP*HMMMMMMMMMMMMMMM*: -                 ,
 dMH'  MMMP'\"HMMMR'T\"  :                   :
|H'   -MR'    ?MMMb    P,       .            .
1&     *|     |. *\"  .- &|                   .
M'      \"    |\\&|  .,#~ \"'                 :
T             :HL.|HMH\\c~ |v,\\_             :
|               \"|:::': -   '\"MM#\\-'.      -
%                   '  ' ' :_ '?' |   .       :
||,                     ,#dMM?.M?      .  .  -
 ?\\                 .,odMMMMMMMMM?    \\    :
  /                 |MMMMMMMMMMMMM:         .'
   .                 TMMH#*9MMMMM*         :
    -.                      \"*#*'    ,:  .
      .                           .v'' .'
        .                           ..'
          '- .                   .-
              \" \\+HHMMMMMMHr~.-

`
var frame22 =  `
              _,,>#b&HMHd&&bb>\\_  __
          _oHMMMMMMMMMMMMMMMMH**H:.   o_
        oHMMMMMMMMMMMMMMMMMMMM#v ?      .
     .dMMMMMMMMMMMMMMMMMMMMMMH* +|       .
    ,MMMMMMMMMMMMMMMMMMMMMb|?+.,H         -.
   ddHMMMMMMMMMMMMMMMMMMMMMb   '            .
  HMMkZ**HMMMMMMMMMMMMMMMMH\\  -   .        :
 dTMMM*   9MMMP'\"*MMMMPT\"  ..               :
|M6H''    4MP'    \"HMMM|   !|.      .         .
1MHp'      #L      $ *\"'  .-:&.               .
MMM'        \"     q:H.  .o#-  '               :
MM'                ?H?.|MMH::::-o,_.           -
M[                   *?:::'|   \" :9MH\\~-.
&M.                     \"\"' '^'.:. ?' . '|  -:
 M|d,                       .dbHM[.1?     .. :
 9||| .                  _obMMMMMMMMH,   .  :
  H.^                    MMMMMMMMMMMM}     -
   \\                     |MMH#*HMMMMH'    .'
    .                            #*'   ,:-
                                 '' .-'.
        .                           .-
          '- .                   .-
              ' \\bqHMMMMMMHHb--

 `

var frame23 =  `
             .,:,#&6dHHHb&##o\\_
          .oHHMMMMMMMMMMMMMMMMMH*\\,.
        oHMMMMMMMMMMMMMMMMMMMMMMHb:'-.
     .dMMMMMMMMMMMMMMMMMMMMMMMMMH|\\/'  .
    ,&HMMMMMMMMMMMMMMMMMMMMMMM/\"&.,d.   -.
   dboMMHMMMMMMMMMMMMMMMMMMMMMML  '       .
  HMHMMM$Z***MMMMMMMMMMMMMMMMMM|.-         .
 dMM}MMMM#'   9MMMH?\" MMMMR'T'  _           :
|MMMbM#''     |MM\"      MMMH.   <_           .
dMMMM#&        *&.     .? *\"   .'&:          .
MMMMMH-          '    -v/H   .dD \"'  '       :
MMMM*                   *M: 4MM*::-!v,_      :
MMMM                      *?::\" \"'  \"?9Mb::. :
&MMM,                        \"'\"'|\"._ \"? | - :
 MMM}.H                          ,#dM[_H   ..:
 9MMi M: .                   .ooHMMMMMMM,  ..
  9Mb  -                     1MMMMMMMMMM|  :
   ?M                        |MM#*#MMMM*  .
    -.                             |#\"' ,'
      .                            -\" v
        -.                          .-
           - .                   .
              '-*#d#HHMMMMHH#\"-'

`

var frame24 =  `
	     ,<_:&S6dHHHb&bb\\_
          .odHMMMMMMMMMMMMMMMMMMM}-_
       .oHMMMMMMMMMMMMMMMMMMMMMMMM#d:.
      ?9MMMMMMMMMMMMMMMMMMMMMMMMMMMH-$ .
    ,::dHMMMMMMMMMMMMMMMMMMMMMMMMH:\\.?? -.
   dMdboHMMHMMMMMMMMMMMMMMMMMMMMMMH, '    .
  HMMMM7MMMb$R***MMMMMMMMMMMMMMMMMH\\ -     .
 dMMMMM/MMMMM*    $MMMM*'\"*MMMM?&'  .       :
|MMMMMMb1H*'       HMP'    '9MMM|   &.    .  .
dMMMMMMM##~         #\\      |. *\"  .-9.      :
9MMMMMMMM*                 |v7?  .,H  '      :
SMMMMMMH'                   '9M_-MMH::-\\v_   :
:HMMMMM                        \\_:\"'|' ':9Mv\\.
-|MMMMM,                         \"\" ' ':. ?\\ \\
 :MMMMM}.d}                         .?bM6,|  |
 :?MMM6  M|  .                   .,oHMMMMM| /
  .?MMM-  '                      &MMMMMMMM|.
   - HM-                         HMH#*MMM?:
    '.                           '    #*:
      -                              -'/
         .                          . '
            .                    .
              '--##HH#HMMMHH#\"\"

`
var frame25 =  `
              _o,d_?dZdoHHHb#b\\_
          .vdMMMMMMMMMMMMMMMMMMMMH\\.
       .,HHMMMMMMMMMMMMMMMMMMMMMMMMH&,.
      /?RMMMMMMMMMMMMMMMMMMMMMMMMMMMMH|..
    ,\\?> T#RMMMMMMMMMMMMMMMMMMMMMMMM6 \\|/
   dMMbd#ooHMMMHMMMMMMMMMMMMMMMMMMMMMH, ' '
  HMMMMMMMTMMMMb$ZP**HMMMMMMMMMMMMMMMM|.   :
 dMMMMMMMM}$MMMMMH'    HMMMH?\" MMMM?T' .    :
|MMMMMMMMMMoMH*''       MM?      MMM|  +\\    .
1MMMMMMMMMMMb#/         ?#?      | #\"  -T:   :
*'HMMMMMMMMMM*'           \"     ~?&  .?} ' ' .
- 4MMMMMMMMP\"                     M? HMTc:\\\\.:
:  MMMMMMM[                       \"#::: > \"?M{
.  |MMMMMMH.                          '  '_ :-
-  |MMMMMMM|.dD                         ,#Mb\\'
 :  *MMMMM: iM|  .                   _oHMMMM:
  .  ?MMMM'  \"'                     ,MMMMMMP
   :   HMH                          JM#*MMT
    -.  '                               #'
      .                                /
        -.            -              .'
           -.                    .
              '--=&&MH##HMHH#\"\"\"

`

var frame26 =  `
              .-:?,Z?:&$dHH##b\\_
           ,:bqRMMMMMMMMMMMMMMMMMHo.
        .?HHHMMMMMMMMMMMMMMMMMMMMMMMHo.
      -o/*M9MMMMMMMMMMMMMMMMMMMMMMMMMMMv
    .:H\\b\\'|?#HHMMMMMMMMMMMMMMMMMMMMMM6?Z\\
   .?MMMHbdbbodMMMMHMMMMMMMMMMMMMMMMMMMM\\':
  :MMMMMMMMMMM7MMMMb?6P**#MMMMMMMMMMMMMMM_ :
 \\MMMMMMMMMMMMb^MMMMMM?    *MMMM*\" MMMR<' . -
.1MMMMMMMMMMMMMb]M#\"\"       9MR'    ?MMb  \\. :
-MMMMMMMMMMMMMMMH##|         *&.     | *' .\\ .
-?\"\"*MMMMMMMMMMMMM'            '    |?b  ,}\" :
:    MMMMMMMMMMH'                     M_|M}r\\?
.     MMMMMMMMM'                       $_: '\"H
-     TMMMMMMMM,                        '\"  ::
:     {MMMMMMMM| oH|                      .#M-
 :     9MMMMMM' .MP   .                 ,oMMT
  .     HMMMMP'   '                    ,MMMP
   -      MMH'                         HH9*
    '.                                  .'
      -                               . '
         .               -          .-
            .                    .-
              ' -==pHMMH##HH#\"\"\"

`

func PrintFileSlowly(fileName string) error {
    fileData, err := artFiles.ReadFile(fileName)
    if err != nil {
        return err
    }

    scanner := bufio.NewScanner(strings.NewReader(string(fileData)))
    for scanner.Scan() {
        line := scanner.Text()
        for _, char := range line {
            fmt.Print(string(char))
            time.Sleep(2 * time.Millisecond) // Adjust the delay here to mimic the baud rate
        }
        fmt.Println()
    }

    return scanner.Err()
}


// Modified loadingScreen function to include a channel for exit signal
func LoadingScreen(done chan bool) {
    frames := []string{frame2, frame3, frame4, frame5, frame6, frame7, frame8, frame9, frame10, frame11, frame12, frame13, frame14, frame15, frame16, frame17, frame18, frame19, frame20, frame21, frame22, frame23, frame24, frame25, frame26}
    ticker := time.NewTicker(50 * time.Millisecond)

    for {
        select {
        case <-ticker.C:
            termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
            w, h := termbox.Size()
            frame := frames[time.Now().UnixMilli()/200%int64(len(frames))] // Cycle through frames

            for y, line := range splitIntoLines(frame) {
                for x, char := range line {
                    termbox.SetCell(w/2-len(line)/2+x, h/2-len(splitIntoLines(frame))/2+y, char, termbox.ColorWhite, termbox.ColorDefault)
                }
            }

            termbox.Flush()

        case <-done: // Exit on receiving a signal on the 'done' channel
            return
        }
    }
}

func splitIntoLines(s string) []string {
    return strings.Split(s, "\n")
}








