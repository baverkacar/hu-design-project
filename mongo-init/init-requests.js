// init-requests.js
db.requests.insertMany([
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844567"),
        "sourceIp": "192.168.6.5",
        "destIp": "192.168.6.4",
        "protocol": "ICMP",
        "time": "22:49",
        "date": ISODate("2024-05-27T22:49:00Z"),
        "permission": "Allowed"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844568"),
        "sourceIp": "192.168.6.4",
        "destIp": "192.168.6.5",
        "protocol": "ICMP",
        "time": "22:49",
        "date": ISODate("2024-05-27T22:49:00Z"),
        "permission": "Allowed"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844569"),
        "sourceIp": "192.168.6.5",
        "destIp": "192.168.6.4",
        "protocol": "TCP",
        "time": "22:50",
        "date": ISODate("2024-05-27T22:50:00Z"),
        "permission": "Denied"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456a"),
        "sourceIp": "192.168.6.5",
        "destIp": "192.168.6.4",
        "protocol": "TCP",
        "time": "21:21",
        "date": ISODate("2024-05-31T21:21:00Z"),
        "permission": "Denied"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456b"),
        "sourceIp": "192.168.6.5",
        "destIp": "192.168.6.4",
        "protocol": "ICMP",
        "time": "16:26",
        "date": ISODate("2024-06-01T16:26:00Z"),
        "permission": "Allowed"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456c"),
        "sourceIp": "192.168.6.4",
        "destIp": "192.168.6.5",
        "protocol": "ICMP",
        "time": "16:26",
        "date": ISODate("2024-06-01T16:26:00Z"),
        "permission": "Allowed"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456d"),
        "sourceIp": "192.168.6.5",
        "destIp": "192.168.6.4",
        "protocol": "TCP",
        "time": "16:27",
        "date": ISODate("2024-06-01T16:27:00Z"),
        "permission": "Denied"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456e"),
        "sourceIp": "192.168.6.5",
        "destIp": "192.168.6.4",
        "protocol": "TCP",
        "time": "16:29",
        "date": ISODate("2024-06-01T16:29:00Z"),
        "permission": "Denied"
    }
]);
