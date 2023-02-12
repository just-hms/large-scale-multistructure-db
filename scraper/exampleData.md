# Scraped data format example

```json
{
    "Roma": [
        {
            "name": "Modafferi Barber Shop: Dal 1970 barbieri a Roma",
            "rating": 4.9,
            "location": "Via dei Cappuccini, 11, 00187 Roma RM, Italy",
            "coordinates": "41.90415580000001 12.4877116",                              #"lat long"
            "imageLink": "",                                                            #A Url if present
            "phone": "+39 06 481 7077",
            "calendar": [                                                               #A place can have more than one hour listing per day. If a day is not included, shop is not open
                {"is_overnight": false, "start": "0830", "end": "1900", "day": 2},
                {"is_overnight": false, "start": "0830", "end": "1900", "day": 3},
                {"is_overnight": false, "start": "0830", "end": "1900", "day": 4}, 
                {"is_overnight": false, "start": "0830", "end": "1900", "day": 5}, 
                {"is_overnight": false, "start": "0830", "end": "1900", "day": 6}
            ],
            "reviewData": {
                "reviews": [
                    {
                        "username": "Ariel Goldberg Afriat", 
                        "rating": 5, 
                        "body": "Amazing barbershop and not a normal one. you pay for a one of kind experience with lots of laughs and professionalism. The crew was very friendly and the way they work was very professional (you also get free drinks with is a great bonus) you can feel the personal treatment a customer receives. They exist for 52 years and by my experience they\u2019re gonna stay for much longer. Thank you Ivano!"
                    }, 
                    {
                        "username": "Kris Green", 
                        "rating": 5, 
                        "body": "This was the first time I went to a barber, it was an amazing surprise from my wife. Ivano is amazing. The whole experience was simply, amazing! If I lived in Rome, I would go back every other week. The staff is extremely friendly. I wish we had this place with the staff in the states. Keep up the great work and making your father proud!"
                    }, 
                    {
                        "username": "Rushal C", 
                        "rating": 5, 
                        "body": "Amazing service and vintage barbershop with great barbers. The place is clean and the haircuts are cleaner.\n\nOnly shop that was open early @8:30am in the area during our stay in Rome. I was there at opening and they got me in the chair right way."
                    },
                    ...
                ]
            }
        },
        {
            "name": "Max&J\u00f2 Barber Shop",
            ...
        },
        ...
    ],
    "Firenze":[],
    ...
}
```