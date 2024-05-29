# example-of-loosely-coupled-microservices

These are general examples of loosely coupled services. Here are three services: user, account, and company. The account passes data to the company service as a value instead of a body.
Whenever data passes from the other two services at the company via API hits, it will be saved.