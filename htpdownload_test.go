package httpdownload

import (
    "testing"
)

func Test_GetandSave(t *testing.T){
    
    GetandSave("https://github.com/golangframework/String/blob/master/String_test.go","/media/timeloveboy/moedisk/热舞/test.go")
}