package pkg

import (
    "dbt-rules/RULES/cc"
)

var main = cc.Binary{
    Out: out("main"),
    Srcs: ins("main.c"),
}
