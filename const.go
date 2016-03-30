package main

var kStopWords = []string{"RT", "http", "co", "jp", "org", "com", "net", "https", "www", "フォロー", "リフォロー", "js"}
var kSymbols = "^彡－Ｏ＝≡◔=☆∧∩*⁄⌑⁄⁄⁄)⁄／`┏┗┐━┏┓☞？-‐'●､･⊂⊃〈〉！≧│➦#[]\"|;:/0123456789＿,.|｜・￣´_\\⌒‿()（）ー…@＠$＄%％&*!+=~?<>＼°。.✦✧┈┼╋【】「」＞＜～♪①②"
var kAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNLOPQRSTUVWXYZ"

const kUrlRegexp = `#(?i)\b((?:[a-z][\w-]+:(?:/{1,3}|[a-z0-9%])|www\d{0,3}[.]|[a-z0-9.\-]+[.][a-z]{2,4}/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s` + "`" + `!()\[\]{};:'".,<>?«»“”‘’]))#iS`
