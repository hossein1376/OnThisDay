# On This Day

What has happened on this exact day, in all the history?  
[Let's find out!](https://on-this-day.vercel.app)

## Source Code

Backend is written in Go, and it's fairly straightforward; the Dockerfile is included as well.  
The app retrieves data from Wikipedia and stores it in a Redis cache database for one day.
The reason being that Wikipedia's API takes something about 5-10 seconds to respond, which is inconvenient.

Frontend is written with Next.js and TypeScript.

## Support Wikipedia

Wikipedia is a great resource, and it's free for everyone. If you have the means, please consider making a donation.