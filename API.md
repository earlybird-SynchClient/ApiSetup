
# ApiSetup API
API for ApiSetup app.

Table of Contents

## invoice get

| Specification | Value |
|-----|-----|
| Resource Path | /invoice |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api/invoice/URL |
| GET Params | Token |
| Example |  http://localhost:8080/api/invoice/qqqqqqqqqqqqqqqqqqqq?Token=4233-1-234 |
{"success":true,"message":"get successfully","data":{"invoicestatus":"enable","vendorcode":"S002-D2","invoicedate":"2018-10-23 00:00","estpaydate":"2018-10-05 00:00","NumberInvoice":"099888765","invoiceamount":"3000.00"}}


## category

| Specification | Value |
|-----|-----|
| Resource Path | /supplier |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api/supplier/URL |
| GET Params | Token |
| Example |  http://localhost:8080/api/supplier/qqqqqqqqqqqqqqqqqqqq?Token=4233-1-234 |
{"success":true,"message":"get successfully","data":{"vendorcode":"S001-D1","supplier":"Change Chen No.3","email":"test3@google.com"}}
