
### Design Choices

#### Database & Data Ingestion

The instructions on the test were written in a manner that left room for interpretation, as it was requested to `read data from a CSV file`, to that end it was decided to create an endpoint to receive a CSV file, validate it and then store it in a database for relational data.

Sqlite was chosen for the purposes of ease of use for the test, however in production a combination of a relational database system (Postgres, MySQL, etc) and an in-memory database (Redis) would be the likely choice.

[Ent](https://entgo.io/) was used as a means to generate the helper functions for interacting with the database, it is a framework I have often used in Golang when building APIs and applications that need to interact with relational databases.

#### Hybrid Functional Programming

There are certain aspects of the code that does and does not adhere to Function Programming (FP) methodology, this is because Golang (unlike a language like Haskell) is not designed with that methodology in mind.

Aspects such as not using `for` `while` loops can be done, but are inefficient (nevertheless, as it was requested, I did make my best effort to try and implement those limitations where possible), and can potentially overflow the call stack.

While I understand the concerns that FP is meant to address, I strongly believe that is it just another tool in a programmer's toolkit and it has it's appropriates places to be used, but one should evaluate if such a tool even need to be used. 

#### Daemon

Whilst working on the weather data fetching, I was presented with a problem of API limitations, the free plan for OpenWeather would only allow 60 calls per minute, and there was more than 60 stadiums being ingested. 

To get around this, I came up with a solution of building a daemon that could run in the background on a separate thread when the API starts up, it would run in a for loop with a cancellable context, that every minute it would check for stadiums' weather data to see if there was any that either did not have weather data or had outdated weather data and would fetch and upsert new data.

The daemon would keep track of how many calls it made per minute and once it hit 60, it would sleep for a minute before resuming, that way it could ensure the data never went over the limit.

It was however pointed out to me that a solution like this was problematic for two reasons:

- What if there were thousands of stadiums?
- In the real world, is there actually going to be anyone who needs weather data for all stadiums?

Due to those questions, after a discussion with the team, I made the daemon optional (with an environment variable/flag) and instead updated the weather stadium endpoint to perform the same task (fetch weather for stadium) but on demand and to follow the same rule of using cached data if the data was new enough.

#### Endpoints

Originally there was going to be one endpoint (which would read the CSV from disk), but as stated above, a decision was made to introduce ingesting the CSV from a request and then storing it into a database.

From there that meant there would be two endpoints, one for ingesting and the other for listing, but then another discussion occurred (about the daemon) and real world use, which split the list endpoint into two different endpoints, one for listing all the stadiums (name and location) and one for listing the weather for the stadium.

That being said, much like how the daemon was made to be optional, I introduced an optional URL query parameter that allows the weather to be returned along with the stadium information, essentially adhering to the provided requirements while also allowing a more `real world` approach to be used as well.

#### Structure

Normally my folder structure for APIs in Golang would group by purpose, models would go in `models` handlers (aka controllers) would go in `handlers` and utilities would go in `utils`; however, due to the request that the methodology of domain driven design (DDD) be used, a new different structure was used, one where entities and their related functions/files (controller, route, service, model etc) were to be stored in a folder pertaining to the entity (stadium in stadium, weather in weather, location in location).

The exception being the CLI flags model, and the database schema and ent generated helper functions.

#### Tests

Due to time constraints, tests were not able to be written in time for submission, but the application has been thoroughly tested manually through means of using tools such as Postman and using the flutter front end for integration testing


### Team Work

This was developed in parallel with Katherine and Scott working on their aspects of the overall project, Scott worked on a front end application that could demonstrate real world usage of the API and Katherine developed the main approach in PHP/Laravel, this code base was developed to show that other languages can be utilized by our team.

Collaboration on this project was done in the form of discussions and review of the code each one did, to identify any potential issues (and their solutions).
