# Point-Of-Sales

Use Tool

    $ go get github.com/jmoiron/sqlx [Lib SQL]
    $ go get -u github.com/gin-gonic/gin [API]
    $ go get github.com/smartystreets/goconvey [Testing]
    $ go get github.com/joho/godotenv
    $ go mod init STACK-GO


Using Hexagonal Architecture

Bisnis Core
- Aktor     :: Admin, Buyer, Kasir, 
- Beli      :: PO -> PI -> INVRECEIPT -> PAYMENT -> BANK/CASH --Buyer
- Jual      :: SO (Pesan Makanan) -> SI (Terbit Inv) -> RECEIVE (CASH/TRANSFER/CREDIT) -> BANK/CASH --Kasir
- Status    :: PI -> Item (Bertambah)
- Status    :: SI -> Item (Berkurang)
