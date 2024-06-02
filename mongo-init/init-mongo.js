db.alerts.insertMany([
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844568"),
        "alertRisk": "medium",
        "alertType": "Data Leakage",
        "ipAddress": "192.168.0.2",
        "createdAt": ISODate("2024-10-08T11:15:00Z"),
        "lastUpdate": ISODate("2024-10-08T11:15:00Z")
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844567"),
        "alertRisk": "high",
        "alertType": "Unauthorized Access",
        "ipAddress": "192.168.0.1",
        "createdAt": ISODate("2024-10-07T13:20:00Z"),
        "lastUpdate": ISODate("2024-10-07T13:20:00Z")
    }
]);
