# Redd Exchange

Redd Exchange is a cryptocurrency exchange. It is still developing stage.

## Architecture

![Alt text](https://github.com/reddtsai/static/blob/master/Exchange.png)

## API

- POST localhost:5001/order/buy

    Request body
    |field|type|desc|
    |-|-|-|
    |symbol|string|交易對，btcusdt|
    |amount|string|交易數量|
    |price|string|交易金額|

    Response
    ``` json
    {
        "status": "ok",
        "id": "12345678"
    }
    ```
- POST localhost:5001/order/sell

    Request body
    |field|type|desc|
    |-|-|-|
    |symbol|string|交易對，btcusdt|
    |amount|string|交易數量|
    |price|string|交易金額|

    Response
    ``` json
    {
        "status": "ok",
        "id": "12345678"
    }
    ```