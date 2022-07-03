# Available Routes

## Track [/track]

Represents the consolidated data on the upper Blockchain

### New (Create) [POST]

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
                "message": "Track Created"
            }

## Segment [/segment]

Represents the data on the lower Blockchain

### New (Create) [POST]

+ Attributes (object)

  + Reference (string, required)
  + Identification (string, required)
  + Data (string, required)

+ Request (application/json)

  + Body

            {
                "Data": "data",
                "Reference": "Track1",
                "Identification": "95bd8c8b-98ac-48e6-883d-1bcf0afe6fbd"
            }

+ Response 200 (application/json)

  + Body

            {
                "message": "Segment Created"
            }
