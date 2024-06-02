db.requests.insertMany([
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844567"),
        "sourceIp": "192.168.0.1",
        "destIp": "192.168.0.1",
        "protocol": "TCP",
        "time": "13:20",
        "date": ISODate("2024-10-07T13:20:00Z"),
        "permission": "Allowed"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844568"),
        "sourceIp": "192.168.0.1",
        "destIp": "192.168.0.1",
        "protocol": "TCP",
        "time": "13:20",
        "date": ISODate("2024-10-07T13:20:00Z"),
        "permission": "Denied"
    }
]);
