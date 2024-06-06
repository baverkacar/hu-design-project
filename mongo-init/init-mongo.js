// init-alerts.js
db.alerts.insertMany([
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844567"),
        "alertRisk": "Medium",
        "alertType": "ICMP Ping",
        "ipAddress": "192.168.6.5",
        "createdAt": ISODate("2024-05-27T22:49:00Z"),
        "lastUpdate": ISODate("2024-05-27T22:49:00Z")
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844568"),
        "alertRisk": "Medium",
        "alertType": "ICMP Ping",
        "ipAddress": "192.168.6.4",
        "createdAt": ISODate("2024-05-27T22:49:00Z"),
        "lastUpdate": ISODate("2024-05-27T22:49:00Z")
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844569"),
        "alertRisk": "High",
        "alertType": "TCP SYN Scan",
        "ipAddress": "192.168.6.5",
        "createdAt": ISODate("2024-05-27T22:50:00Z"),
        "lastUpdate": ISODate("2024-05-27T22:50:00Z")
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456a"),
        "alertRisk": "High",
        "alertType": "TCP SYN Scan",
        "ipAddress": "192.168.6.5",
        "createdAt": ISODate("2024-05-31T21:21:00Z"),
        "lastUpdate": ISODate("2024-05-31T21:21:00Z")
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456b"),
        "alertRisk": "Medium",
        "alertType": "ICMP Ping",
        "ipAddress": "192.168.6.5",
        "createdAt": ISODate("2024-06-01T16:26:00Z"),
        "lastUpdate": ISODate("2024-06-01T16:26:00Z")
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456c"),
        "alertRisk": "Medium",
        "alertType": "ICMP Ping",
        "ipAddress": "192.168.6.4",
        "createdAt": ISODate("2024-06-01T16:26:00Z"),
        "lastUpdate": ISODate("2024-06-01T16:26:00Z")
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456d"),
        "alertRisk": "High",
        "alertType": "SSH Brute Force Attack",
        "ipAddress": "192.168.6.5",
        "createdAt": ISODate("2024-06-01T16:27:00Z"),
        "lastUpdate": ISODate("2024-06-01T16:27:00Z")
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca084456e"),
        "alertRisk": "High",
        "alertType": "SSH Brute Force Attack",
        "ipAddress": "192.168.6.5",
        "createdAt": ISODate("2024-06-01T16:29:00Z"),
        "lastUpdate": ISODate("2024-06-01T16:29:00Z")
    }
]);
