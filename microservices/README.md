# 《Microservices》

Go from monolith to microservices in Go for robust application architecture.

## ch1

### Introduction to Microservices

- Date Intensive Applications
  - Certain business logic cannot be handled by an RDBMS
  - Certain data might not map easily to a table structure like product catalog
  - Certain values might need to be precomputed (similarity algorithms)
  - Others may need to be handled on the fly with precomputed values
  - Memory-intensive applications-personalization
  - These kind of apps can be easily handled by their own specialized service - good
  target for microservices
  
- Microservices Architectures Are a Great Solution to App Bloat.
  - Monolith
    - All functionality lives on the same repo - mono repo
    - They runs on the same process, can be scaled by running various instances of the same process
    - Tightly coupled
    - Same language
    - Deploy all functionalities simultaneously (train model)
  - Characteristics of Microservices
    - Separate some functionality into its own server application
    - Make it accessible with a well-defined contract, generally through REST
    - Generally with its own data store
    - Able to scale independently
    - Able to deploy independently
    - Can be handled by separate teams
    - Use the right tools for the specific job (programming languages, frameworks, and data stores)
  
### Writing Our First Microservices

- Steps:
  - Spinning a web server in Go
  - Adding a handler
  - Accessing QueryString data
  - Returning JSON

#### MongoDB tips

```sudo docker run --name s1v6mongo -d -p 27017:27017 mongo```

```sudo docker exec -it s1v6mongo /bin/bash```

```mongodb
> show dbs
> use dbName
> show collections
> short = db.collectionName
> short.find()
> short.count()
```  

## ch2

- URL Syntax
  - `scheme:[//[user[:password]@]host[:port]][/path][?query][#fragment]`
  - Required - scheme, domain
 
- The Context Package
  - Using the context package, to cancel execution of a group of HTTP requests.
  - to share request-scoped data, among requests.
  
- What Are Contexts?
  - The context type - carries deadlines, cancelation signals, and other request-scoped values,
  across API boundaries and between processes.
  
  
## ch3

- What Is JSON?
  - JSON (JavaScript Object Notation) is a lightweight data interchange format.
  - Very easy for humans to read and write and also for machines to parse and generate.
  - Currently, it is the most common format for using in REST services.

- What Are Protocol Buffers?
  - Protocol buffers are Google's language-neutral, platform-neutral, extensible, mechanism for
  serializing structured data.
  - Binary format rather than a text one (like XML or JSON).
  - Smaller and faster to transmit over the network.
 
 
- Docker Postgres faults:
  - `sudo docker run --name postgres -v s3v6pgdata:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=packt -e POSTGRES_USER=packt -e POSTGRES_DB=wta -d postgres`
  - `Exit(1)`
  - `docker container logs 83e5a189c6ad`
  - `docker pull postgres:10`
  - `sudo docker run --name postgres -v s3v6pgdata:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=packt -e POSTGRES_USER=packt -e POSTGRES_DB=wta -d postgres:10`
  
  
## ch4

- v1
  - Using HTTPS to secure your services
  - Using Let's Encrypt to acquire HTTPS certificates
  - Access tokens for authenticating users of our services
  - Using JWT as access tokens
  - general security guidelines
  
- v2
  - how TLS handshakes use public key cryptography to start communications
  - Why certificate authorities are needed
  - Using HTTPS in our applications
  - Using self-signed certificates to test our applications
  - Customizing out HTTPS configuration to comply with Mozilla guidelines
  
- TLS - Transport Layer Security
  - Recommended protocol for establishing secure communications
  - Superseded SSL
  - Starts with a mechanism known as the TLS handshake
  - The first steps of this mechanism use public key cryptography

- Public Key Cryptography
  - A message encrypted with the public key can only be decrypted with the private key.
  - A message signed with the private key can be verified by anyone with the public key.
  
- `ll /etc/ssl/certs` 

``` 
cd ~
mkdir certs
cd certs

openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key

openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650 -subj "/C=AR/ST=CABA/L=CABA/O=Example Org/OU=IT Department/CN=*"

openssl genrsa -out client.key 2048
openssl ecparam -genkey -name secp384r1 -out client.key

openssl req -new -x509 -sha256 -key client.key -out client.pem -days 3650 -subj "/C=GB/ST=LONDON/L=LONDON/O=Another Org/OU=IT Department/CN=*"

// check: openssl x509 -in server.pem -text -noout
```      

- What Is an Access Token?
  - A string that can be used to authenticate a user when accessing an API.
  
- What Is JWT?
  - A technology that can be used to generate Access Tokens containing information
  that can be used when authenticating request.
  - part of  the JOSE (Javascript Object Signing and Encryption) standard.
  - Parts of a JWT
    - Header
    - Payload (claims)
    - Signature

- General Practices to Secure Your Microservices
  - Use an API Gateway and put your core services out of the public network.
  - Accept traffic only from IPs you trust.
  - Always use HTTPS when authenticating users.
  - Use HTTPS in all communication (from service to service).
  - Use modern configurations for TLS
  - Use CA-signed certificates for TLS.
  - Set appropriate timeouts on your servers.
  
- Security Practices for Datastores
  - Don't store passwords in plain text. Encrypt them.
  - Take measures to prevent SQL injection.
  - Set up appropriate users with restricted access based on their usage of the database.
  - Restrict traffic by IP.
  - Follow all the security guidelines of your particular datastore.
  - Sanitize user output if you are sending data that will be displayed on an HTML page to prevent XSS.


## ch5

- What Is Load Testing?
  - A type of software testing that we do to understand the behavior of an application
  under normal and anticipated peak conditions.
  - It helps to identify the maximum operating capacity of an application as well as any
  bottlenecks and determine which element is causing degradation.