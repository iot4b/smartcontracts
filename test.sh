#!/usr/bin/env bash

go run . device
go run . device new cf5291ae66fc448f6c8cc683753457f5bfa64bd99ea44b17a3da6002f00b5bbd 074321f6dff911894aa791105af3c216a8eb802134a0656501ef64750d64a0e0 '{
                        "node": "0:8de799e17832c2380235e8ff508cca308ca8ff8ec633aae2efbca9abfb3dbe55",
                        "elector": "0:8e18edd847fdc6bdd95640b3ff76a90d1d12d757c92061d0bfb12a03440f759e",
                        "vendor": "0:8de799e17832c2380235e8ff508cca308ca8ff8ec633aae2efbca9abfb3dbe55",
                        "owners": [
                            "0:a61e618c41bc8558efccc5d352acc857b271eb9a29587d8aa95a004a0cb39e8e",
                            "0:8e18edd847fdc6bdd95640b3ff76a90d1d12d757c92061d0bfb12a03440f759e"
                        ],

                        "active": false,
                        "lock": false,
                        "stat":  false,

                        "type": "test-device",
                        "version": "0.1mac",

                        "vendorName": "apple",
                        "vendorData": {
                            "serialNumber": "123311-33V3309",
                            "region": "Russia"
                        }
                    }'

go run . node
go run . node new cf5291ae66fc448f6c8cc683753457f5bfa64bd99ea44b17a3da6002f00b5bbd 074321f6dff911894aa791105af3c216a8eb802134a0656501ef64750d64a0e0  '{
                      "elector": "0:8e18edd847fdc6bdd95640b3ff76a90d1d12d757c92061d0bfb12a03440f759e",

                      "ipPort": "157.245.57.218:5683",
                      "location": "Russia",
                      "contactInfo": "test-device"
                  }'

go run . elector
go run . elector new cf5291ae66fc448f6c8cc683753457f5bfa64bd99ea44b17a3da6002f00b5bbd 074321f6dff911894aa791105af3c216a8eb802134a0656501ef64750d64a0e0 '{"nodes":["0:8de799e17832c2380235e8ff508cca308ca8ff8ec633aae2efbca9abfb3dbe55"]}'

go run . owner
go run . owner new cf5291ae66fc448f6c8cc683753457f5bfa64bd99ea44b17a3da6002f00b5bbd 074321f6dff911894aa791105af3c216a8eb802134a0656501ef64750d64a0e0 0:8e18edd847fdc6bdd95640b3ff76a90d1d12d757c92061d0bfb12a03440f759e

go run . vendor
go run . vendor new cf5291ae66fc448f6c8cc683753457f5bfa64bd99ea44b17a3da6002f00b5bbd 074321f6dff911894aa791105af3c216a8eb802134a0656501ef64750d64a0e0  '{
                        "elector":  "0:8e18edd847fdc6bdd95640b3ff76a90d1d12d757c92061d0bfb12a03440f759e",
                        "vendorName": "apple",
                        "contactInfo":  "owneer@node.com, Alex, +1 222 333 4567",
                        "profitShare": 50
                    }'