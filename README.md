# Live version of the app [here](https://rss-agg-demo.netlify.app/).

<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="20%" alt="GOBACKEND-logo">
</p>
<p align="center">
    <h1 align="center">GOBACKEND</h1>
</p>
<p align="center">
    <em><code>❯ RSS-feed aggregator</code></em>
</p>
<p align="center">
	<!-- local repository, no metadata badges. --></p>
<p align="center">
		<em>Built with the tools and technologies:</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=default&logo=YAML&logoColor=white" alt="YAML">
	<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=default&logo=Docker&logoColor=white" alt="Docker">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=default&logo=Go&logoColor=white" alt="Go">
	<img src="https://img.shields.io/badge/PostgreSQL-4169E1.svg?style=default&logo=PostgreSQL&logoColor=white" alt="PostgreSQL">
	<img src="https://img.shields.io/badge/Heroku-430098.svg?style=default&logo=Heroku&logoColor=white" alt="Heroku">
</p>

<br>

##### Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Tests](#-tests)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)

---

## Overview

<code>❯ RSS feed aggregator in Go</code>

---

## Features

<code>❯ Add, follow & unfollow, fetch all of the latest posts from the RSS feeds</code>

---

## Repository Structure

```sh
└── gobackend/
    ├── backend-deployment.yaml
    ├── backend-service.yaml
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── gobackend.exe
    ├── handler_err.go
    ├── handler_feed.go
    ├── handler_feed_follows.go
    ├── handler_readiness.go
    ├── handler_user.go
    ├── internal
    │   ├── auth
    │   └── database
    ├── json.go
    ├── main.go
    ├── middleware_auth.go
    ├── models.go
    ├── postgres-deployment.yaml
    ├── postgres-service.yaml
    ├── Procfile
    ├── README.md
    ├── rss.go
    ├── scraper.go
    ├── sql
    │   ├── queries
    │   └── schema
    ├── sqlc.yaml
    └── vendor
        ├── github.com
        └── modules.txt
```

---

## Modules

<details closed><summary>.</summary>

| File                                                           | Summary                                                                                                                                                                                                                                          |
| -------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [backend-deployment.yaml](gobackend\backend-deployment.yaml)   | <code>❯ file contains the configuration for deploying a backend service in a Kubernetes cluster</code>                                                                                                                                           |
| [backend-service.yaml](gobackend\backend-service.yaml)         | <code>❯ Kubernetes configuration defines a Service of type LoadBalancer for the backend application</code>                                                                                                                                       |
| [docker-compose.yml](gobackend\docker-compose.yml)             | <code>❯ defines a multi-service application using Docker Compose (version 3.8) with two services: db and backend</code>                                                                                                                          |
| [Dockerfile](gobackend\Dockerfile)                             | <code>❯ This Dockerfile is used to build and run a Go application</code>                                                                                                                                                                         |
| [go.mod](gobackend\go.mod)                                     | <code>❯ Go module file defines the dependencies for a Go project</code>                                                                                                                                                                          |
| [go.sum](gobackend\go.sum)                                     | <code>❯ This snippet provides the hashes for the Go module dependencies in the go.sum file                                                                                                                                                       |
| [handler_err.go](gobackend\handler_err.go)                     | <code>❯ HTTP handler function in Go called handlerErr</code>                                                                                                                                                                                     |
| [handler_feed.go](gobackend\handler_feed.go)                   | <code>❯ provides API handlers for managing feeds. It uses the encoding/json package for JSON parsing, the google/uuid package for unique identifiers, and the database package for database interactions</code>                                  |
| [handler_feed_follows.go](gobackend\handler_feed_follows.go)   | <code>❯ defines several HTTP handler functions to manage feed follows within an API. It leverages the go-chi/chi router for request routing and google/uuid for generating and parsing UUIDs.</code>                                             |
| [handler_readiness.go](gobackend\handler_readiness.go)         | <code>❯ HTTP handler for a readiness check</code>                                                                                                                                                                                                |
| [handler_user.go](gobackend\handler_user.go)                   | <code>❯ defines HTTP handlers for user-related operations in a web service</code>                                                                                                                                                                |
| [json.go](gobackend\json.go)                                   | <code>❯ contains two helper functions for handling HTTP responses in a web service</code>                                                                                                                                                        |
| [main.go](gobackend\main.go)                                   | <code>❯ sets up an HTTP server using the net/http package along with chi for routing and cors for handling cross-origin requests. It connects to a PostgreSQL database and serves API endpoints related to users, feeds, and feed follows</code> |
| [middleware_auth.go](gobackend\middleware_auth.go)             | <code>❯ defines a middleware function for handling authenticated requests in a web serve</code>                                                                                                                                                  |
| [models.go](gobackend\models.go)                               | <code>❯ provides utility functions for converting database models to API response models. It defines types and conversion functions for User, Feed, FeedFollow, and Post objects</code>                                                          |
| [postgres-deployment.yaml](gobackend\postgres-deployment.yaml) | <code>❯ Kubernetes deployment YAML configuration</code>                                                                                                                                                                                          |
| [postgres-service.yaml](gobackend\postgres-service.yaml)       | <code>❯ Kubernetes Service YAML configuration</code>                                                                                                                                                                                             |
| [Procfile](gobackend\Procfile)                                 | <code>❯ used in Heroku to define the command that should run to start the web server</code>                                                                                                                                                      |
| [rss.go](gobackend\rss.go)                                     | <code>❯ defines a function for fetching and parsing an RSS feed from a given URL</code>                                                                                                                                                          |
| [scraper.go](gobackend\scraper.go)                             | <code>❯ contains functions for periodically scraping RSS feeds, parsing their content, and storing the data in a database</code>                                                                                                                 |
| [sqlc.yaml](gobackend\sqlc.yaml)                               | <code>❯ configuration file is used for generating Go code from SQL schema and queries, specifically for a PostgreSQL database</code>                                                                                                             |

</details>

<details closed><summary>internal.auth</summary>

| File                                       | Summary                                                                                            |
| ------------------------------------------ | -------------------------------------------------------------------------------------------------- |
| [auth.go](gobackend\internal\auth\auth.go) | <code>❯ function for extracting an API key from the Authorization header of an HTTP request</code> |

</details>

<details closed><summary>internal.database</summary>

| File                                                                   | Summary                                                                                                             |
| ---------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| [db.go](gobackend\internal\database\db.go)                             | <code>❯ defines the basic structure for interacting with a database using the SQL commands generated by sqlc</code> |
| [feeds.sql.go](gobackend\internal\database\feeds.sql.go)               | <code>❯ Tulossa...</code>                                                                                           |
| [feed_follows.sql.go](gobackend\internal\database\feed_follows.sql.go) | <code>❯ Tulossa...</code>                                                                                           |
| [models.go](gobackend\internal\database\models.go)                     | <code>❯ Tulossa...</code>                                                                                           |
| [posts.sql.go](gobackend\internal\database\posts.sql.go)               | <code>❯ Tulossa...</code>                                                                                           |
| [users.sql.go](gobackend\internal\database\users.sql.go)               | <code>❯ Tulossa...</code>                                                                                           |

</details>

<details closed><summary>sql.queries</summary>

| File                                                       | Summary                   |
| ---------------------------------------------------------- | ------------------------- |
| [feeds.sql](gobackend\sql\queries\feeds.sql)               | <code>❯ Tulossa...</code> |
| [feed_follows.sql](gobackend\sql\queries\feed_follows.sql) | <code>❯ Tulossa...</code> |
| [posts.sql](gobackend\sql\queries\posts.sql)               | <code>❯ Tulossa...</code> |
| [users.sql](gobackend\sql\queries\users.sql)               | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>sql.schema</summary>

| File                                                                            | Summary                   |
| ------------------------------------------------------------------------------- | ------------------------- |
| [001_users.sql](gobackend\sql\schema\001_users.sql)                             | <code>❯ Tulossa...</code> |
| [002_users_apikey.sql](gobackend\sql\schema\002_users_apikey.sql)               | <code>❯ Tulossa...</code> |
| [003_feeds.sql](gobackend\sql\schema\003_feeds.sql)                             | <code>❯ Tulossa...</code> |
| [004_feed_follows.sql](gobackend\sql\schema\004_feed_follows.sql)               | <code>❯ Tulossa...</code> |
| [005_feeds_lastfetchedat.sql](gobackend\sql\schema\005_feeds_lastfetchedat.sql) | <code>❯ Tulossa...</code> |
| [006_posts.sql](gobackend\sql\schema\006_posts.sql)                             | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>vendor</summary>

| File                                        | Summary                   |
| ------------------------------------------- | ------------------------- |
| [modules.txt](gobackend\vendor\modules.txt) | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>vendor.github.com.go-chi.chi</summary>

| File                                                            | Summary                   |
| --------------------------------------------------------------- | ------------------------- |
| [chain.go](gobackend\vendor\github.com\go-chi\chi\chain.go)     | <code>❯ Tulossa...</code> |
| [chi.go](gobackend\vendor\github.com\go-chi\chi\chi.go)         | <code>❯ Tulossa...</code> |
| [context.go](gobackend\vendor\github.com\go-chi\chi\context.go) | <code>❯ Tulossa...</code> |
| [Makefile](gobackend\vendor\github.com\go-chi\chi\Makefile)     | <code>❯ Tulossa...</code> |
| [mux.go](gobackend\vendor\github.com\go-chi\chi\mux.go)         | <code>❯ Tulossa...</code> |
| [tree.go](gobackend\vendor\github.com\go-chi\chi\tree.go)       | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>vendor.github.com.go-chi.cors</summary>

| File                                                         | Summary                   |
| ------------------------------------------------------------ | ------------------------- |
| [cors.go](gobackend\vendor\github.com\go-chi\cors\cors.go)   | <code>❯ Tulossa...</code> |
| [utils.go](gobackend\vendor\github.com\go-chi\cors\utils.go) | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>vendor.github.com.google.uuid</summary>

| File                                                                 | Summary                   |
| -------------------------------------------------------------------- | ------------------------- |
| [CONTRIBUTORS](gobackend\vendor\github.com\google\uuid\CONTRIBUTORS) | <code>❯ Tulossa...</code> |
| [dce.go](gobackend\vendor\github.com\google\uuid\dce.go)             | <code>❯ Tulossa...</code> |
| [doc.go](gobackend\vendor\github.com\google\uuid\doc.go)             | <code>❯ Tulossa...</code> |
| [hash.go](gobackend\vendor\github.com\google\uuid\hash.go)           | <code>❯ Tulossa...</code> |
| [marshal.go](gobackend\vendor\github.com\google\uuid\marshal.go)     | <code>❯ Tulossa...</code> |
| [node.go](gobackend\vendor\github.com\google\uuid\node.go)           | <code>❯ Tulossa...</code> |
| [node_js.go](gobackend\vendor\github.com\google\uuid\node_js.go)     | <code>❯ Tulossa...</code> |
| [node_net.go](gobackend\vendor\github.com\google\uuid\node_net.go)   | <code>❯ Tulossa...</code> |
| [null.go](gobackend\vendor\github.com\google\uuid\null.go)           | <code>❯ Tulossa...</code> |
| [sql.go](gobackend\vendor\github.com\google\uuid\sql.go)             | <code>❯ Tulossa...</code> |
| [time.go](gobackend\vendor\github.com\google\uuid\time.go)           | <code>❯ Tulossa...</code> |
| [util.go](gobackend\vendor\github.com\google\uuid\util.go)           | <code>❯ Tulossa...</code> |
| [uuid.go](gobackend\vendor\github.com\google\uuid\uuid.go)           | <code>❯ Tulossa...</code> |
| [version1.go](gobackend\vendor\github.com\google\uuid\version1.go)   | <code>❯ Tulossa...</code> |
| [version4.go](gobackend\vendor\github.com\google\uuid\version4.go)   | <code>❯ Tulossa...</code> |
| [version6.go](gobackend\vendor\github.com\google\uuid\version6.go)   | <code>❯ Tulossa...</code> |
| [version7.go](gobackend\vendor\github.com\google\uuid\version7.go)   | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>vendor.github.com.joho.godotenv</summary>

| File                                                                 | Summary                   |
| -------------------------------------------------------------------- | ------------------------- |
| [godotenv.go](gobackend\vendor\github.com\joho\godotenv\godotenv.go) | <code>❯ Tulossa...</code> |
| [LICENCE](gobackend\vendor\github.com\joho\godotenv\LICENCE)         | <code>❯ Tulossa...</code> |
| [parser.go](gobackend\vendor\github.com\joho\godotenv\parser.go)     | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>vendor.github.com.lib.pq</summary>

| File                                                                        | Summary                   |
| --------------------------------------------------------------------------- | ------------------------- |
| [array.go](gobackend\vendor\github.com\lib\pq\array.go)                     | <code>❯ Tulossa...</code> |
| [buf.go](gobackend\vendor\github.com\lib\pq\buf.go)                         | <code>❯ Tulossa...</code> |
| [conn.go](gobackend\vendor\github.com\lib\pq\conn.go)                       | <code>❯ Tulossa...</code> |
| [connector.go](gobackend\vendor\github.com\lib\pq\connector.go)             | <code>❯ Tulossa...</code> |
| [conn_go115.go](gobackend\vendor\github.com\lib\pq\conn_go115.go)           | <code>❯ Tulossa...</code> |
| [conn_go18.go](gobackend\vendor\github.com\lib\pq\conn_go18.go)             | <code>❯ Tulossa...</code> |
| [copy.go](gobackend\vendor\github.com\lib\pq\copy.go)                       | <code>❯ Tulossa...</code> |
| [doc.go](gobackend\vendor\github.com\lib\pq\doc.go)                         | <code>❯ Tulossa...</code> |
| [encode.go](gobackend\vendor\github.com\lib\pq\encode.go)                   | <code>❯ Tulossa...</code> |
| [error.go](gobackend\vendor\github.com\lib\pq\error.go)                     | <code>❯ Tulossa...</code> |
| [krb.go](gobackend\vendor\github.com\lib\pq\krb.go)                         | <code>❯ Tulossa...</code> |
| [notice.go](gobackend\vendor\github.com\lib\pq\notice.go)                   | <code>❯ Tulossa...</code> |
| [notify.go](gobackend\vendor\github.com\lib\pq\notify.go)                   | <code>❯ Tulossa...</code> |
| [rows.go](gobackend\vendor\github.com\lib\pq\rows.go)                       | <code>❯ Tulossa...</code> |
| [ssl.go](gobackend\vendor\github.com\lib\pq\ssl.go)                         | <code>❯ Tulossa...</code> |
| [ssl_permissions.go](gobackend\vendor\github.com\lib\pq\ssl_permissions.go) | <code>❯ Tulossa...</code> |
| [ssl_windows.go](gobackend\vendor\github.com\lib\pq\ssl_windows.go)         | <code>❯ Tulossa...</code> |
| [url.go](gobackend\vendor\github.com\lib\pq\url.go)                         | <code>❯ Tulossa...</code> |
| [user_other.go](gobackend\vendor\github.com\lib\pq\user_other.go)           | <code>❯ Tulossa...</code> |
| [user_posix.go](gobackend\vendor\github.com\lib\pq\user_posix.go)           | <code>❯ Tulossa...</code> |
| [user_windows.go](gobackend\vendor\github.com\lib\pq\user_windows.go)       | <code>❯ Tulossa...</code> |
| [uuid.go](gobackend\vendor\github.com\lib\pq\uuid.go)                       | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>vendor.github.com.lib.pq.oid</summary>

| File                                                        | Summary                   |
| ----------------------------------------------------------- | ------------------------- |
| [doc.go](gobackend\vendor\github.com\lib\pq\oid\doc.go)     | <code>❯ Tulossa...</code> |
| [types.go](gobackend\vendor\github.com\lib\pq\oid\types.go) | <code>❯ Tulossa...</code> |

</details>

<details closed><summary>vendor.github.com.lib.pq.scram</summary>

| File                                                          | Summary                   |
| ------------------------------------------------------------- | ------------------------- |
| [scram.go](gobackend\vendor\github.com\lib\pq\scram\scram.go) | <code>❯ Tulossa...</code> |

</details>

---
