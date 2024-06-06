// init-devices.js
db.devices.insertMany([
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844567"),
        "osName": "Linux",
        "ipAddress": "192.168.6.5"
    },
    {
        "_id": ObjectId("5f47ac62e7179a6ca0844568"),
        "osName": "Windows",
        "ipAddress": "192.168.6.4"
    }
]);
