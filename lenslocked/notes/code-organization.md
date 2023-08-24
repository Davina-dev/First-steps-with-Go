With this code structure nested code like that in psql/user_store.go will often reference types defined in parent packages (eg app.User), but the reverse is typically a code smell and can lead to cyclical dependency bugs.

While this code structure can work incredibly well when you get the hang of it, it is a bit more advanced and confusing, so it is generally not a great choice to initially learn with, and most developers have more luck adopting it after they have some experience with Go.Separate of concerns

Outside of HTML and CSS, Model-View-Controller (MVC) is the most common pattern that falls under separation of concerns.

> - **Controller**: _connect it all. accepts user inout. passes that to models to do stuff, then passes data to views to render thingd; usually handlers_
> - **Models**: _data, logic, rules; usually db_
> - **Views**: _rendering things usually html_

Below is an example of what our application might look like using MVC.

```
myapp/
  controllers/
    user_handler.go
    gallery_handler.go
    ...
  views/           []
    user_templates.go
    ...
  models/
    user_store.go
    ...
```

**Buffalo** uses what I consider to be a variation of MVC, but doesn’t name directories exactly the same. Controllers are stored in an actions directory, and views are stored in a templates directory, and there are several other directories being used in addition to the three core MVC directories. It is important to remember that using MVC does not mean you are limited to only putting code in those three directories.
MVC is very common in frameworks for a variety of reasons, but one of those reasons is that it is predictable. It doesn’t matter if you are building a banking app or time tracking app, any project using MVC will store the various types of files in the same directories making it easier for the framework tooling to guess where the controllers or views are in the source code. With that knowledge, frameworks can often do some magic behind the scenes to speed up initial development.

Examples of popular web frameworks that use MVC include Ruby on Rails (ruby) and Django (python).

## Dependency based structure

Code can also be structured based on dependencies. This is a strategy taught by Ben Johnson on his site **Go Beyond**.

The idea with dependency based structure is to try to organize your code based on what dependencies it has. For instance, if you used a PostgreSQL database, you might have a postgres package that has all of your postgres-related code. If you rely on the Stripe API, you might have a package named stripe with all of your code related to interacting with Stripe.

In order to put all of these dependencies together, code structured this way will often have a common set of interfaces and types that are dependency agnostic. Dependencies then use these interfaces for anything outside of their scope, and we can provide the actual implementations when setting up our application.

Let’s look at an example. Imagine we had the following directories and source files.

**\*Warning**: If some of the code here starts to go over your head don’t worry about it. This organization pattern is a little more advanced and we will likely explore it again later in the course when you have a bit more experience.\*

```
 myapp/
  user.go
  user_store.go
  psql/
    user_store.go # implements the UserStore interface with a Postgresql specific implementation
```

Inside user.go we might declare a User type. This type won’t be specific to a database or anything else, it is just a User type that we can use anywhere in our application.

```
package app

type User struct {
  ID UserID
  Name string
  Email string
}
```

user_store.go might provide us with a UserStore interface that defines some common actions we take with users, but doesn’t actually provide an implementation.

```
package app

type UserStore interface {
  Create(name, email, password string) (*User, error)
  // ...
}
```

At this point our code doesn’t have any real implementation of the UserStore type, so we might provide a Postgres specific implementation inside psql/user_store.go.

```
package psql

type UserStore struct {
  // ...
}

func (us *UserStore) Create(name, email, password string) (*app.User, error) {
  // Create a user in the Postgres database.
}
```

With this code structure nested code like that in psql/user_store.go will often reference types defined in parent packages (eg app.User), but the reverse is typically a code smell and can lead to cyclical dependency bugs.

While this code structure can work incredibly well when you get the hang of it, it is a bit more advanced and confusing, so it is generally not a great choice to initially learn with, and most developers have more luck adopting it after they have some experience with Go.

With this code structure nested code like that in psql/user_store.go will often reference types defined in parent packages (eg app.User), but the reverse is typically a code smell and can lead to cyclical dependency bugs.

## Many more...

- Domain driven desing
- Onion architecture
- ...
