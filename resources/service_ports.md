| service       | api  | 依赖                                                         | rpc  | 依赖            |
| :------------ | ---- | ------------------------------------------------------------ | :--- | --------------- |
| user          | 8000 | UserRpc UserTokenRpc GoodsInfoRpc UserAddressRpc ShoppingCartRpc OrderRpc | 9000 | UserTokenRpc    |
| usertoken     |      |                                                              | 9001 |                 |
| admin         | 8002 | AdminRpc AdminTokenRpc GoodsInfoRpc GoodsCategoryRpc CarouselRpc IndexConfigRpc OrderRpc | 9002 | AdminTokenRpc   |
| admintoken    |      |                                                              | 9003 |                 |
| goodsinfo     | 8004 | GoodsInfoRpc                                                 | 9004 |                 |
| goodscategory | 8005 | GoodsCategoryRpc                                             | 9005 |                 |
| carousel      | 8006 | CarouselRpc IndexConfigRpc                                   | 9006 |                 |
| indexconfig   |      |                                                              | 9007 |                 |
| useraddress   |      |                                                              | 9008 |                 |
| shoppingcart  |      |                                                              | 9009 |                 |
| order         |      |                                                              | 9010 | ShoppingCartRpc |
| mqueu-job     |      |                                                              | 9011 | OrderRpc        |
|               |      |                                                              |      |                 |

