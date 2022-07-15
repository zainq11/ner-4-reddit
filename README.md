# ner-4-reddit
An attempt to perform Named-entity recognition (NER) on Reddit content

## Done
- Get access token
- refresh access token when expired
- Review reddit APIs to call

## TODOs
- Stream of text from a reddit link
- Process data (fuzzy word match)
- Append to list


- Main will supply a sub-reddit to the Client. This can be passed into `InitClient`.
- `InitClient` will return a `Client`
- Main will call `Read` on `Client` to get data filled into a `[]byte`.
- `Client.Read` will invoke reddit api and pull the latest data from the ongoing anchor to the next limit. Ongoing anchor
is described in https://www.reddit.com/dev/api listings section. `Client` keeps track of the state of the reader.

