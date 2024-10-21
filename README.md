# Go LIMS

Go LIMS--a lightweight Laboratory Information Management System built primarily to evaluate the DX of Go for enterprise, internal applications.

## The Data Model

This app's data model is relatively simple. I define the following objects: Analyses, Samples, Tests, and Resources.
* Users define one or more analyses, which represent a specific procedure that can be run against a Sample.
* Tests link Samples to Analyses in a **many-to-many relationship** and are where results are stored.
* Each Test has associated Resources in a parent-child relationship. Resources record what Equipment or Media was used to run the test. There are multiple **types** of resources.

## The Stack

**Data Storage**
In-memory storage only

**Backend API**
Standard library, as defined in the [Writing Web Applications](https://go.dev/doc/articles/wiki/) tutorial.

**Frontend Reactivity**
HTMX

## Project Values

As this is an internal application, I am evaluating GoLang for these attributes:

1. *Expressive language features,* to model complex business logic
2. *Explicit control flow,* to identify bugs before they ship
3. *Simple build and deployment pipelines,* to keep maintanence costs low
4. *Solid testing ergonomics*

I am **less concerned** with these features:

1. *Low-level hardware access.*
2. *Beautiful or innovative frontends.* The less I have to do in the frontend, the better.