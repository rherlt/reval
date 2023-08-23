export var ratings = {
    "scenarios": [
        {
            "id": "7d1e8d9a-f0c6-4f12-86d4-6956e1b47e01",
            "name": "Llama 2 Szenario",
            "description": "Das ist das erste tolle Szenario mit den Antworten aus Llama-2-13b-chat-hf_responses_rated_by_gpt-3.5-turbo.json",
            "totalResponseCount": 400,
            "progressStatistics": [
                {
                    "name": "rated",
                    "value": 70
                },
                {
                    "name": "unrated",
                    "value": 330
                }
            ],
            "resultStatistics": [
                {
                    "name": "positive",
                    "value": 12
                },
                {
                    "name": "negative",
                    "value": 46
                },
                {
                    "name": "neutral",
                    "value": 12
                }
            ],
            "ratingScore": {
                "min": -1,
                "value": -0.5,
                "max": 1
            }
        },
        {
            "id": "7d1e8d9a-f0c6-4f12-86d4-6956e1b47e01",
            "name": "Vicuna Szenario",
            "description": "Das ist das zweite tolle Szenario mit den Antworten aus Vicuna-33b-chat-hf_responses_rated_by_gpt-3.5-turbo.json",
            "totalResponseCount": 400,
            "progressStatistics": [
                {
                    "name": "rated",
                    "value": 370
                },
                {
                    "name": "unrated",
                    "value": 30
                }
            ],
            "resultStatistics": [
                {
                    "name": "positive",
                    "value": 310
                },
                {
                    "name": "negative",
                    "value": 40
                },
                {
                    "name": "neutral",
                    "value": 20
                }
            ],
            "ratingScore": {
                "min": -1,
                "value": 0.8,
                "max": 1
            }
        }
    ]
}