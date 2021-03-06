package main

import (
	"fmt"
	"strings"

	"github.com/mbark/advent-of-code-2020/util"
)

const input = `
swswswswswnwswswswsweswsesw
wneswseseneswnweneswwswseswwnwseswswe
ewswswswwswnwnwsweswwwwwswwwswwsw
senwseseswseseswswesewseseseseseswsese
seswnwseseswseseswswseseswswseswesese
swnwnenwnenwnwnwswnwnwnwnwnwnwnwenwnwnw
seseeseneswnwseseneeeseeswsewseseese
senenwnwnewnwnwnwnenwnenwnwnwnenwswnenw
eneneneweeneneeneneneeneneenesew
nenesenewwenwseseseseswsesewneseweswse
swswweswswswswswswswseswnenwswswsweswsw
nwswnwnwnwnwnwnenenenwnwnw
seseeneseeesenesewseswsewseeeese
wwseswwnwneswsewswwswswswwswswswswne
wwewwnwwwswsenwnewnwwnwsewwww
wnwnwnwenwnwnwnenwnweswnwnwwswnwnwnw
neeneneswnenenenenwnenenwnenenenenewnwe
nenwsenwnwnwnwnwswswnwnwnwenwnwnwnwnwnw
swenwsenweenwswnwswseseeswswesenwne
ewseswsewswneewneneneeneseneeswnwsw
nwswnwswewswswswesesweswswse
swswswswenwswswswwswneswneswswswsesww
seswswseseswswsenwswse
wwwwwwwswsewnewwnwwwwwew
swwwswewswswswswswnwswswswswseswsww
sesewsewneseneseseseswswseeseseswnesw
wwswswwwwswwewwswwwsewnwnww
nwnenwnenenwnenwnwnwsenwnwne
neewwwneswseseeeneeseesesesewee
senwneeneneenewneneneneneneneswnenenene
swswswneswwswswwswswsewneswesw
nenenenenwnenesewnenenenenenenwnwnenwne
nwnwnwsesenwnwwsewnwnenwnewnwnwwwnwnw
nwnenwwnweswseseeswenenwweneneeswswsw
eeesweeeenweseseweeneeseeese
swswseseswseswnwseseswswswseseseneswswsese
swnwnwewnewwwwnwnwwnwwswnwse
sesenwnenenwnenwnwnwnenwnewnwnwnw
neswneneeneneenenenenewnwneneswnenwewne
nwnenwswnwenwnwnwnwnwnwnwswnwnwnwnwnwnwnw
nwwwewwneewswneswnwwwwwsewsew
wswswswswswweswwwwswseswswswneswswe
wnwwnwnwnewnwnwnwwnwwsenwwnwnwwnw
wswwwswswwwwnesew
nwnwneenenwnwswnwnwnwswnwsenwnwwenenenwnw
swsewseseswswseswneseseswseseswseseseese
swseswseswnwsewneweseseseeswsesesesew
sweswwwneswswwswesenenwsweswswswswse
wswseneeseweeeeweseesenwnwseew
wwnwnwwwnwnwnwnwwnwnwnwnwswwnwwe
swnwnwwseneneneneneneswswswnewenwsese
seseseseseesesenwseseneesewswseswnwnw
neswswsenwesenesenesenesesewwneswnwswnwne
nesesesenenwnwnwnenwnwswnwneenewnenwne
neneneneneeneenewseewnenesenenenenene
eseseseseneseeseseesenwseseseseseseesww
neeenenesweeeeeeweneeeeeene
enwwwwwsewwwnwnwwnwwwwwnwnwnw
neseeeeseseseeseeseneeseseseenwwsw
enenesewneenenenwseseneeweeswenwene
wneneseneneneenenenenenwnenenenenenene
wnwneeswewnwnwseswswsweswsenewswsw
eseeswsenwenweseseeseseseesenesesese
wswswseseseseneenesesesewsesenewsew
wnwwwwwwewwswwwwwwwswsww
swenewwwseswwnesesenewneswseneswswswsw
eswseseseseesewseseneseenwsewseee
swneneswneneneneeneeenwneeenenwwenene
neswnenesenenwnenenenenenenenenenenenwnene
sesesesesewsesesweseseseseseseesesenee
nwnwnwnwnwnwswnwnesewnwnwnwwnwnwnwnenw
eenesewwwwnwnwwwwnwwwsw
swnenwwswswswneseeseneswneseswwseswsw
neneenenwnenenwnwnwseswnwnwnenwnenenenenw
neenenwneneseneeneneeneneeeeneswnee
swswswswneseswwswswseswswweneenwnew
neseewsesesewswsese
enenweneeneneeeeneeswswneneeeneenw
seeeeneeeenwseewswenewnwseesew
sweeeneesenweswewsweseeseenenwee
seesesewswenenwnweswsewseneenwsesw
swwswwswwwneswwsww
wwwwwwwwwwwwewwwweww
neneneseneeneneeswneenesenwnewnenwe
nwnwnwnwnwnenwnwnwnwsenwnwnwwnwnwnwnwnw
eeeeeseseeenweeeesweeenee
nwnenwnwneneseneswnenwnenwwnwnwnenenwnwnw
sewswswwsewwneesewnewnewne
nwnwwnweswnwneswswnwnwnwnenwnweswenwe
nwneenenenenwneneneeneseeneneswnenenene
swwwwswswswswnewwwwwsw
neneneneeneeneeneneeewnwneswneneesene
nenenewnenenenenenenewnenwnwneesenene
swnwswswwwswswwweswnewswswwswnwswsee
neseseeseseeseseseeeseseewesewse
nwnenenesesenwnwswnwwswwnwnwnenwnwsew
seesesewseneseseseseseseswsesesenwswsese
nwsenwwnenwneseswwwnwwnewesewwswnw
sesesesenesesesesesewsewsesesesesesesesee
seswwnwseeswnwwneeneseweeseenee
nwseenwnwnwnewswnw
nwnwnenwsenwnwnwnwnww
wnwwwnwenwnwwsenewnwnwnwwsewwnenw
sesenwsenwseswseseseseeewsenesesesesese
seeneeeeseeenwnwsweeeeneeene
eeeeseeeeswneeneeneesweenew
swseneswswnwseswswseswswswswneswswswsesw
nesesweneseseneeweesenwnwseeswnwswse
senwswnenenwnwneeesw
nweeseeenweesweeeeeseeswese
sweneenwneeneeneesenwnwnwnesweswswnee
swswnwewswnwswwwwswwswwwewswswswe
swswswsenwswwswsweswswnwsweswswesesw
swswsesesenenwewseseswseeswseseswsesesw
seswnewewswnenwenwnenenenenenenwwene
wnwswswwwsewnwnwswweseswsenewwww
wswswweswwwnwsewwswswwswsw
swswseswsweswswswswswswswswswswnwswswsw
nenwnewnenenenwnenenwnwnenwneneneenese
seswswswswseneswsesw
neseneneenewsenwwnwnwwswwnwwee
swswswwswswswwnwnewwwneeneswwnwse
nwswwwwnwswesewsenwseswwswwwwwnew
wsenweeeseeeseeseeeeeeeseese
seseseseseseeswseesesesenesesewnesese
nenesesewsenwneseseseneseswseswwswsesese
sewseseewsenenweewswseseneeswsee
swseewnwsenwnenwswneswnwwswnesenwnewnw
neeeewseseneswnenwneneweneneseseew
nenwnenwswneneenenwnesesenewneneswnenwnw
nenewnenesweneneenesenenenenenenenwneee
wnwsenwnwsenenenesesesenesewnwnwnwnwnenw
nwswswswnewnwsweseswnesenweswenwesww
nwnwnenesenwnwwnwneswnenwnwnwnwnwenwnw
seswseseswswneswnwswseswsesesenwswswswswsw
nenwnwsenwnwnwnenwnenenenwsenwwnwnwnenwnwne
eswwswswwswswswswswswwsww
nenesenwenwsweesweeneenw
nenenesenenenenwnwnenenewnenwswnwnenenwne
eeseeeneseewneeeeseeweeeese
swswwseseswsenwseeswswswsesesw
nwnwsenwsenenenwswseswsewseswswwnwswnew
nwwnwnwnwenwnwwwnwwnwwwnwnwwwsw
wewnwwswswnwwnwnwnwnwewswnwesenwwne
nenenwnwnwnwenewenwsenwnenenwnwnewnwne
seseseenwseseseseseswwseswnwsee
swseswseseseswsesesesenwseseswsese
swswneseswseswwswswswseswsweswswswswswsw
seneneeseseswseseswsenwseeseseswsesee
nwneneseneneneneneenesenewseneewnene
nwwnenwnenenwnwsenwnwsenwnwnwnwnenenenw
wwnwwsewwwwwsewwwnwnwnwnwwnese
seseswnwseswseswsesesenwsesesese
wnweswsenwnwswnenwnwnwswwsewnenwnwwnw
wseswnwseswswseswswsenwswswsesene
neeeneenwneneeeeneeesweenenweesw
wwwnwnwwwwwwwwwnwwwwesew
neseeneneneneneneneneneneenwnenwneneswnene
neeeesenweeswenewne
nwswnwenwwwwnwswwwwwwnwnwenwnwe
wwwwswwwwwwswwwwwewsw
eeseeswswswneeeweewneneneenenene
swswswesesesenwswseswnese
sesewswseseseswseseseseseseseseseswnese
neeeeeneeweeeesweeenenwe
nwwnesenwneneswenenenenenwnwnenwnwwnene
neeenwneneswneneneswneeeneswneene
nwwewwnwnwnwswwswwwnwwwwwwnew
nenwnwnwwswnwwnwnenwesee
wswswswswwseswswswswnewswswnwswwswsww
neneneneswsenenenwne
neeeweeeeeeeseneneneneeesewe
nwswnwnwnwnwenwneeswnwnwwnwnw
swsesenwseswseseswwseneseneseseesesesese
nweeweeeeeseeseeswseeee
nwsesesesesewesenwnenwswneeseneswsesw
nwneneenenewsweeneneneneneneeewnene
swseeseneswwneswnwwswneewnesw
neseseswswseswseswseswsenwsweseswsesesese
seswswwswwwwewswswwnewwwswswnww
senwseesesesesesesesesesesesenwsenwsese
seswnwsesesewseseneseseseesenwsesesesewse
neeenwnesweenwneeeeeeeeeneswene
seenwsenweswnwswwnwnwnwnwnewnwswenw
swswswswswswswswswswwswswseswne
enwswswnwswneswsewnwswseneseenwwnwsw
swneswnenenwnwneeeneneneeswneswnenesew
nwswnwwswwseewseewsenenwnwwwnwwwse
eeeseesenesesewseeseswseneesesesese
swnewswseswwwnwwewew
eeeeeeseeeeeswnwe
nenenenenwnwneeneswwneeneneneswnenenesw
eswwswswswsweswnw
newnwsenwnwnenwnwenwnenenwnwneswnwwswnene
eeneeneneeeneeweeswneeneseeewne
nwwnenewenenwnwnwnwe
neneswwneswneneneswene
swwneneweneeswwneenesenwnwenenese
seesewsesesesesenesesesesewsesesesesese
nesenwesesesenwwswsesewesesesesesesesese
nwsenwneeweenweswnwnwnwwwenwwne
seswswswnwswseswswswseswenwswswswswnesw
swnwsesweswswseswwswswwseeswswsenesesw
ewseseneeseeeneeenwnweeesenwne
neswswswswswswswswseseswwenwswnwswswsw
nwswswwseeswswswswswswswswswswswswswne
nenwnenenenenwnenenenewnenenwneenwnw
senweeeenwseeseeenwseeseseeeese
seseseswseseswswseseseneneseseswswwseswswse
swswwswwswwswswswswwwwswwswwe
sesesewseswseseseseseseseesesesese
nwwnwnwnwnwenwwnwnwnwwnwnw
nwnwnwnwnewnenwnenwnenwnesenenwnenenwnw
nwwseneesenwwswwneseenwseswwwnwe
enwnwswsesenwswsesweeseseseseewnwswse
eeeeseseeesweneeeewewwse
ewneneeseeneseswswnwswwseswesesene
eeeneeeneeweeneeeeeeweese
nwnesenwnwnewnwnenwnenwnwneeswnwnwneswe
swseweeeseseneeeeseneseseeneeew
seewnenenwenesesewnwnwnwseswe
nesenewsesweseseswwsesewneswswswese
neenenesenwneeneeneswswneneenenenene
wneswsenenwsewseeswseswswnesewswnese
senwnwnenenewwwnenwenwnenwseenenenw
nwwwsewneswswwnenwswnewnenwseesenenwnw
eneneneeneseswweeneeewnweneee
seenenwswswswnwnwwseneneseeeeeswsw
nenenesenewnenenwneneneneneneenenenenene
wswwwwwwswwswwnwwswwswewwew
neenweneneseeneswwswnew
eeeenweeseneeneeeeeneeswee
nwsesewsesesewnwseswneseseesesesenwse
wnwnenenenenwwesweeneneswsenwnwswsene
swswwswswswswswswwsweneswnwwnwswesw
wwewwwwnwww
eeeeneswsweneenwnenwneneneeneene
sewseswswswswswseseswseewnwnenwsw
neeneseweeneeneesweeneenweese
sweeeeeeneeeeeeee
wnwnwwnenewwnwnwwsewswwwwnwnwwnww
seswwseseseswnwswseeswswswnwseswseswswse
eseswnenwweeswnwwneeswnese
sesesewnewwwwneneww
nweneesweneesenewnenw
eesesesenwseseeseseseseseese
neswswnewenenenenenenwenenenwne
swnenwnwnwsenwwnwwwneewwnwwwnenwwse
seswswsesewweneenenwsesewswsw
seeseeeswenwswnweeeswnweeeeee
sweswneswwwwswne
seseswseswswseneswswseswseswneswswseswsese
wnwwnwnwnwwwwwnwwwwnwenww
wswseswswswswsweseswswnwseswneswswswsw
enwwnwnwwseeesewnww
nwnwsweseswswswwswswseswswweswnwswe
swseswswswwwswswnewsw
wswnewnewnewenwsewwwseseswnwnwwse
eseneseseeesesweesesewseseesesewse
neewseeseeeswswsesesenenw
newneweenenesewseewwnwwsenesesee
wnenwwswwswswnwnweeenwneswseenwnwnw
swnwnwnwnwenwnenwnwnwnwnwnw
neneesewnweseeneeneseewsewnwswne
wwwwwenwwwwwwswwwwwnwsew
seesweneseeseweeseneseeseseeee
nwnwnwsewwnwwnwnwnwnwnwnwnwnwnwneswwse
newnwswnwwnwnwnwnwwwnwnwnwwwnwnwenw
seweeeseeeseeseeeeeseswnweseene
eseseeeeeswesesweeseneeesenesee
eseswwsesenwseseswseseseeswsenese
wnwnwnwnwnenwnwenenw
neneenesweswsenwseneneswnenwwwneswneww
wswneswewswseswswswswwsw
seseeneeseswsewneswnwseneswswneenwnee
nwwnesenwsenwnenwnwnesewnwswswnwneswwnw
enweneeeneeneneeneneswseneseenwne
nwnenwnenenwnwnwsenwnwnewwnwnenwene
wwnwneswwewnwweswwneseswwswnesew
enewneewsewenenenenesenenewene
nwsenwnwnwnenenenenenenenwnwnwnenw
nenenenenenwnwneneneswne
enwneseeneneneneswwneeneneswe
swswwswsewsewwwwneeeswwwswnwww
swnwseenenenwenwnenwnwwnwswneswnwnwnenw
wswswwswwwswswswswswswwswweswnwsw
seseswwneswwswswswswswswneswswswswswsww
seseswnwnwnwnwnwnwnwnwwewnwnw
wwnwnwwwnwnewnwsewwnwnwnwnwnwewsw
wseswnwwneswnenewwseewseswswwwne
wswwwwwwwwwnwewwwnw
senwswseneseseseswswswsweswswseswseswsw
wwwswwwswwwsewnwswwnewwswswswsw
swseswswwswseswswseswseseswseswne
ewswswsweswswnewswseneseswseswsenwnwe
wwwwwsenwwwwwwwwwwne
nwsewwwwsewnwwwnwne
swswswswswnewswwwwwswswswswswsw
nwwseenwsenenwwnwnwnwnwenwwnwnwwsw
eswseswswsenenenenwswseswswswseswsenwswse
wswwswwswwsewwwswswswwnwwnene
sesesesesesenewwswwseneseseeswsesesese
swnwnwnwnwnwnwnwnenwnwnwenenwnwnwwnwnwnw
wneenenwneenwwnenwnwnwneneneneeneswnw
swwwwwwwwwwwwwwwnewseww
senwneswseswsesenwsesee
neswwseseeswseseswswsesww
nwswseswseseswseneeseseneseseseswseswnesw
eneeeneneneewenenenenenweenesene
wnwwnenenwnenwswnwnenweenwnwnenese
wswseswswneswswswnweswswsweseseswsene
senwwseneseeesewnesewswsewsesenwnese
esenesesesesesesesesesesewsesesewsesw
neeneneseswnwnesenwwnwnwswnenwne
neenenwnenewneswnewneneneenenenenenwe
swseswswswswseswswnwswsweswsesesesesw
nwenwsewnwnenenwnwwnwnwnwswnwnwnwnwwnw
wsenwwnwswneweswwweewnewesene
swwnwwswwwnwnwewewewwwnwnwnww
wsewwnwswsweswswswswswswswswswswswswsw
seswwswwwewswewwwwnewwwwne
esesesesesewneseeseseseseeswseneesesese
nwnwnenenwneneeneweneneneseseswnenwnenw
neneneeneswseswwnenenenenenwneneneneenene
nwnwnwnwnwnenwsenwnwnwnenwnwwnwnenw
eeeeeeewnweeeseeeeeeee
wwwnwwnwewnwwnwwnwswnwwnwnewwnw
swenwswnwswesesenwswswswswnwwnweseesw
neeneneeneneswneneneswneneneneenenenene
nwsenwseeeesesese
wnwwnwnwwnweesewnwswnwswnwnwseswneww
nwseseswseseeneseneseseseseseswswesesee
neneeenenwswneeeswneneneenwnenenene
seswseeswwswswswswswsweswnwse
neeeneneeneswenenwnenwneeneweswnee
ewewwswnwneewwnwswsewswwenew
swwswswswsweswswswswswswswsenewswswsesw
nwsewnwswnwnwsenwnewwnenwnwnwwwnwewnw
nwseswwwwwswswswwswneswwswwwwsww
nenwnwnwnenenenwnenenenenenesenenenene
eseseswsweseswseswsenesesewwsesesesese
sesesesesesesesenewnwseseseseseseesese
nenenenesenewnesenewwenesenenwnenene
wewewswwswnwwwwwwwswwww
enwwnwnwnwwwswnewnwnwwnwseewnwsenew
swseseenwnweswseseswswnenenwnesesewsw
wwnwwwwsewwwwwwewwwwwww
nwewnwnwwenwswnwnwenwnwnwnwnwsw
wsesesenwsenwneeseseewseeswseseese
enweneeeswnewneswewneneeneeneee
seswwseswseseseswswesesesesesesesenwsese
swwnwswwnewsewesesenwwewnew
nweneeneseneneneenenee
newnenenenenenwnenesenenenenenene
wwwwwwwwswwnwswwwsewwwwwne
swnwswnwenesenwnenew
nwwwwwwwwwwsewwwwwnewww
swwsweswnwnenwnwnenwnenewewswsenwswnese
enenewneswneneneeeeneeseneneneneenee
nenwnenenenenwnenwwnenenenesesenenwnewne
swnenenenesweswnenwneswswnwswnwnwneswswne
wwwwsesewneswwwnewwnewwnewsew
seswseseseseseeseseswsesesesenenwsesesese
wwwswwnwnwewwnwwnwwww
seseseseswsesenwswsese
swnwswswswswswnwswsweswswswswswswswseswswsw
nwnenwswnenwnwnwnwsenenwnwnwenwnwnwnwnwne
seneneneneenwneneswwnwswneneswnesenenw
swswnewwnenwewswwswswswsesewswsww
neeswwnwnwnwnwswnewenenwnwnwnwnwesewne
nwswseswewnewswnwswwsweswseeneseswne
sesenwswesenweenwwnwesw
seeseseseswnwseseseswseseswsesesesesese
esewseeesweseseseseesenenweseese
nwswseenweeneeeeeseenwseenwsesesw
enweeeseeesenwsweee
nenwneneeeneneneeneeseeneeneneneew
swswseswneswseswswswse
enenenenwnwnwneswnenwnwnwnwneeneswnenenw
nenenwneswnenwnwnenwnweswnwnenwnwnenenwne
swseesewseeenwenwnwsenenwweseese
swseenenenenenenenenenwwnwnenwwenenene
nwnwnweneneewwnwsenenenenwswswnesene
seeswswesenwneeneeenweswesesenww
eeeeeneseeeeweeeeeesweee
enwswswneswswswwswswswswseswswswswswswsw
sewsesesesesesenesesesesesesesesenesesw
nesenesenenewnweneeenenenenewnee
swseseseswsweswsesesesweswnwswswseswsesenw
eeeeeeeweeeeeeweeeeee
nwnwnwnwnenwnenwnwesenwwnwnwnenwnwnwnw
wswwnwewnwwwswswwwswswswwswswe
eeenweeneeeeeeeweeeseseee
seswseseseswsesesesesenesesesese
wnwnwwsewswwwewwwnwwneweseswne
nwnwnwnweswnwnwnwnwwnewnwswnwnwnwnwnw
seseeseewseswsewsesesewneeseswsenene
seseneseswseesenwsesesenenwseseswwsesese
wswwnwenwwnwswnewwwwwnwnenwwnw
ewewwwwwwswwwwwewnwwsww
wwsewswwwwnewneswwwswwnenwwne
wsewnwnwnwwwwwnweeswnewnwwnwnw
nwnwwnwnwnwnwnwswnwenwewswnwwwwnwnw
newwwewswsenwnwnwwwsew
swswwseswnwnweswwsweswwwwswnwswnwse
eseseseenweewseseseseseeseeeenwe
eseseeeswesesewnwenwseenwseseesene
wwnwwnwseewnwnwnenwsesenwnwswnwsene
eseeseseewseseseeseese
nwenwneswweneeeswseeseenwneee
wweswwwswwseswwswswswwwswwnew
seenwwseswwseeneseenwsenwswe
nenenenwwneswnenenwnwsesewseneneneenwswnw
nwnenenenenwneneneswnwnenwnwenene
sewnwswesewsewneneswwwsenwne
wnwnwnwnwnenwnenwwswnwnwnwnwnwwwswnw
sweseneseseswswswswseswnwswseswwswsww
wenwwwsenwseswnwnwwwnenwnwnwnwnwww
seneswswswswneswswswswswswswwswswwsw
nwnwnwenwnwnwwnwnwnwnwnwnwnwnwnwnwswnw
nwnwesenenenwnenwnwnewnenewnenwnwswnwsenw
nwwnwwenwnwnwsenwnwnwwewwnwnwwnw
nwnwnwnwnwnwswnwnwnwnenwnwnenwswnenw
nwnwnwnenwnwnwnwnwnesenenenenwnwnwnww
swsenwnwwnesenwnwwnwnwsewnwseenwnwne
swswenwswseswswswwswswswseswswnwseswswse
wneneneweseneswswswsenww
nwneswswnwswnwenenwwneswnwnenenenesenwne
nenenwneseswnenenenenwneneeneenewee
nwwwewwnwnwnewsewswnewesewnww
nwswneswseseseswseseswseseswseseseseswse
swswwwwwswwwnwswsew
swsweswneswswswnweswwswswswswswwsww
swswswswswswswswenwwwswswwswneswswsw
seeeeseweswswenwneseeseenwnwswseww
enweeneeenweneeseseneeweseee
nwswnwenenwnwnwnwswnwnwnwenwsenwwwse
nweewwnwesenewswnwwswsenwwnwsww
enwnwswwseswnewsewneeeswseew
neneenesweeneeneneneneewnwenenese
swwneneseenesewsesesenewesesesesewsenw
wnwwnwwwwwewwwwsww
wnwswwnwsesewwwwwnewnenww
nwwnwnwnenwswnwneneenwnwswnenwnwnenenw
eswswewwswwnenwewneneswswewsesw
neswwswswwswswswwswswswwswwewswww
eseenwweseseeeeneeeseeenwese
neeeeweeeeeneenwesesweswneswse
nwwwnwwwewwnwswnwnwwenwsenwwwse
nwnwnwnwnwnenenwnwnwnwnwnwnwnenwnesenwsene
nwseswswswswseswswseswswswswswswswswneswsw
wnwnenwwenwnwwewwenwwsenwwswwswnw
seseswneswnwswwswwseseswneenenenesww
esesewseeneseswwsesesesesesesesesese
wwnwnwnewnwnwsenwnwwnwenwnwsenwsenw
swwwneswwwwwwwwwewswwwww
seneneeseneewswnwswswnewnenwwnenenwswne
swswseswwwswwwswwswwwwwwwnew
eeseseewseseseenwseeeeeseseesese
enenenenwnwswnenwnenwnenenwneswswnenw
nesewswsesesesesenesesenwseeseneseswswe
neeneesweswneeswneenwenweenenwsw
swnenwwsenwneswneneswnenenenenesweene
wnwewwnwwwwwnwnwsew
nwnwneswwnwnwswneenwnenwneneeenenwnw
neeeneseeewneseswnwseswnw
nwseenenewenwnewneswnenwenwswsenene
wnenwnewnwnwnwnwnwenwnesenwnwnenenwnwnwnw
newwwwswwnwwswwswswewsewweww
neeenenenwneneseneseseenenenwnewnene
wnweseswseseenwnw
senwswseswnweseseswseswswswswswswswenw
nwnwnwnwnwsenwswnwnwnwnenwwnwnwnenesenw
eneneneenesenesweeeeeeeeeewe
nwnenwneneswnwnwnwnwnenwnwnenenwewnwnee
neeenwseeeneeeeneeeeseenenew
wsweeswseswswswswweswewnwneswseswnw
swswnwseswswswnwswswswswswswwwswswenwse
wnweswwswswnwseswswneswswsewswswsw
nwnenenwnenwnwnenenwewnenwnwnwsenwnwnesw
wnenenesenwnenenenenesenenenenewnenenenene
swnwnewswswseseswseswwnwnwsweneswsew
neeeeneenenenenenwneneswweenenenene
nwewswswnenenwnwwswweweneeesenwne
sesewenesesewsewsesenesesenwnesesenwsw
nwewwnwwwwnwwwwswwwwe
eseseesewsesesesewneseneeseeseesesew
neneneeneneeeeeneenenesw
newewswsesewnenwwwswwwnwwsw
sewwnwewwsesenenwnwswswnenenw
nwnwnwsewwswesenee
ewnwseewswnwwswwwwsw
wwwenewwesewwwwwwwwwww
wnwseeswenenwww
neeseseeesesewwneeeeseswneeseswe
wwsewswwswwswwswwswnewswswswwsw
eeseeweseenweseseswsesenesesesese
senenwwnwwwnwswnwewwnwwswnwwnww
nwnwnwswnwnenwnwnwsesenenenwnwnenenenenw
wswnwswswewnwwwnwseneesenewswwew
nweeswesweeeeeeeeeeeenwe
eeneeenenwenweeeeseneesweew
seewswseswsweswwenwswswseswswnweswswsw
seswenewwswswnewnwwsenenewswswwwse
nenewweeneeseneeneeswnenenewnene
neweswnenwnwswneenwwsenwsenwenewnw
nwnwneseswnenweenenenenesewnenwsenenwwe
nesweswenesweneenweenewsenenenwnw
seeeseeeseseseenenweeseeenweseew
nenewneeswneneneeneneseneeeneneenenee
swswswswnwwseseseswnwsesenwseseeesesene
swnwsesenwneseenwsesesesewsesweseeee
nwneneneneseneneneneneneneew
wswwwwwwwsenwnwwnwnwwwwneww
`

const testInput = `
sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew
`

type Tile struct {
	x int
	y int
	z int
}

func main() {
	stop := util.WithProfiling()
	defer stop()

	instructions := util.ReadInput(input, "\n")
	tiles := first(instructions)
	fmt.Printf("first %d\n", len(tiles))
	tiles2 := second(tiles, 100)

	fmt.Printf("second %d\n", len(tiles2))
}

func first(instructions []string) map[int]Tile {
	tiles := make(map[int]Tile)

	for _, ins := range instructions {
		tm := Tile{}

		for len(ins) > 0 {
			tmm, remove := tm.Move(ins)
			tm = tmm
			ins = ins[remove:]
		}

		if _, ok := tiles[tm.Hash()]; ok {
			delete(tiles, tm.Hash())
		} else {
			tiles[tm.Hash()] = tm
		}

	}

	return tiles
}

func second(tiles map[int]Tile, times int) map[int]Tile {
	for i := 0; i < times; i++ {
		newt := make(map[int]Tile, len(tiles))
		adjacentBlack := make(map[Tile]int, len(tiles)*4)

		for hash, tile := range tiles {
			adjblack := 0
			for _, tm := range tile.Neighbors() {
				if _, ok := tiles[tm.Hash()]; ok {
					adjblack += 1
					continue
				}

				adjacentBlack[tm] += 1
			}

			flipToWhite := adjblack == 0 || adjblack > 2
			if !flipToWhite {
				newt[hash] = tile
			}
		}

		for t, black := range adjacentBlack {
			if black == 2 {
				newt[t.Hash()] = t
			}
		}

		tiles = newt
	}

	return tiles
}

func (t Tile) Hash() int {
	return t.z + 1000*t.y + 1000*1000*t.x
}

func (t Tile) Neighbors() []Tile {
	neighbors := make([]Tile, 6)
	for i, m := range []struct {
		x int
		y int
		z int
	}{{z: 1, y: -1}, {x: -1, z: 1}, {x: 1, z: -1}, {z: -1, y: 1}, {x: 1, y: -1}, {x: -1, y: 1}} {
		neighbors[i] = Tile{x: t.x + m.x, y: t.y + m.y, z: t.z + m.z}
	}

	return neighbors
}

func (t Tile) Move(ins string) (Tile, int) {
	x := t.x
	y := t.y
	z := t.z
	remove := 1

	if strings.HasPrefix(ins, "se") {
		remove = 2
		z += 1
		y -= 1
	} else if strings.HasPrefix(ins, "sw") {
		remove = 2
		x -= 1
		z += 1
	} else if strings.HasPrefix(ins, "ne") {
		remove = 2
		x += 1
		z -= 1
	} else if strings.HasPrefix(ins, "nw") {
		remove = 2
		z -= 1
		y += 1
	} else if strings.HasPrefix(ins, "e") {
		x += 1
		y -= 1
	} else if strings.HasPrefix(ins, "w") {
		x -= 1
		y += 1
	}

	return Tile{x: x, y: y, z: z}, remove
}
