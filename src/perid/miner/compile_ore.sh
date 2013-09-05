#! /bin/sh
ragel -Z ore.rl
ragel -Vp ore.rl > ore.dot && dot -Tpng ore.dot > ore.png && firefox ore.png
go test

