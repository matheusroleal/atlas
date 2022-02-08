# Atlas
Middleware for Storing Data on Multiple Blockchains

![Atlas Icon](infra/doc/icon.png?raw=true "Icon retrive from https://www.vectorstock.com/royalty-free-vector/atlas-holding-up-world-mascot-vector-20996478")

## What it is?

A blockchain is a distributed database that maintains a continuously growing list of ordered records, called blocks. These blocks are linked using cryptography. Each block contains a cryptographic hash of the previous block, a timestamp, and transaction data. 

Our systems tries to make it simple to create links cross-chains. In other words, it allows blockchains to speak to one another in a standardized way. Cross blockchain compatibility, allows different blockchains to communicate with one another without the help of intermediaries. Instead, people will be able to transact with users from other compatible blockchains. The entire process will take place without any downtime or expensive transaction fees.

## How it works?

![Atlas Diagram](infra/doc/BlockchainMiddleware.png?raw=true "Icon retrive from https://www.vectorstock.com/royalty-free-vector/atlas-holding-up-world-mascot-vector-20996478")

### Track [/track]

Represents the consolidated data on the upper Blockchain

#### New (Create) [POST]

+ Attributes (object)

    + Reference (string, required)
    + Identification (string, required)

+ Request (application/json)

    + Body

            {
                "Reference": "Track1",
                "Identification": "95bd8c8b-98ac-48e6-883d-1bcf0afe6fbd"
            }

+ Response 200 (application/json)

    + Body

            {
                "message": "Segment Created"
            }
