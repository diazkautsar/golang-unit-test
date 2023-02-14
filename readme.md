unit file harus berakhiran _test
penggunaan nama function harus ada nama Test nya. contoh TestHelloWorld

masuk di rectory package
go test
go test -v
go test -v -run=NamaFunctionTestnya

go test ./... 
untuk run di root project. semua unit test

....
t.Fail() => menggagalkan unit test tapi akan tetap melanjutkan code program selanjutnya, ketika selesai maka unit test tersebut dianggap gagal
t.FailNow() => menggagalkan saat itu juga

t.Error() => bisa kasih message, terus akan panggil t.Fail()
t.Fatal() => bisa kasih message, terus akan panggil t.Now()