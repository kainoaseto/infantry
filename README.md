# Infantry
"Clean" Onion Architecture based high performance Go Rest Framework built on Echo

Currently, this is a WIP as existing progress has been halted by rearchitecting based on 
the knowledge and benefits of using the Onion Architecture described here:
https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This architecture change was made after this frameworks conception so this project is
now in a WIP state. The main reason I am going against Go protocol by merging modules
into one framework is to provide interfaces all in one package for an Onion Architecture
approach. Go seems to already hint that this architecture may be intended. This would
be welcomed as this is mainly a discovery project to see just how well this architecture
works in Go and in other projects such as Pratton.

Benefits include:
- Extremely modular code that allows for seperation between logic and external tools
- Easily testable
- Easily migratable
- Usable in a Monorepo (Single repository - Google style)
