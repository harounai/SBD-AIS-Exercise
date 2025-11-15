# EX5 Summary:

I synced my fork with the professor’s repo and opened the Exc_5 skeleton in my IDE so I could start working on the assignment. I added the sws service to the docker compose file, mounted the frontend folder, and set it up so the frontend could run at [http://localhost](http://localhost). After that, I configured Traefik with the needed options, networks, and labels, and connected everything correctly to the web network.

I also set up the orderservice with the right environment variables, healthcheck, and routing labels for Traefik. While doing this, I ran into a bunch of small issues like wrong YAML indentation, an incorrect Postgres volume path, leftover labels, docker volume errors, and some Git problems. I fixed all of those along the way and managed to get the services running without errors.

In the end, the frontend worked fine, Postgres was healthy, and the orderservice itself was running. The only thing I could not solve was getting [http://orders.localhost](http://orders.localhost) to work through Traefik, so I could not actually see or test the orderservice in the browser. That is the point where I stopped.

**Note to the professor:**
We talked about this already in class, but here's the note as you requested. The submission is late because I wasn’t fully aware of the deadline. It did not show up in my Moodle dashboard until the morning of the due date. Thanks for understandin!
