## Walking Through a Web Request with MVC

MVC may be easier to understand if we take a normal web request and examine how each part of the request would be handled in a theoretical application. We haven’t written all the pieces of code that we are going to talk about just yet, but seeing how data will flow through our application once we build it should help as we progress with building our application.

We will walk through what happens when a user submits an update to their contact information. For instance, they might be updating their address after moving.
![ ![Alt text](![image.png](mwc-1to2.png))](mwc-1to2.png)

**1. A user submits an update to their contact information**

The first step is pretty straightforward. The user submits their updated contact information to our server by submitting a form. Their browser then sends a web request to our server and our router is pretty much the first thing to interact with the request. Technically other code may run before the router, and we will see this later in the course, but for now we can just assume that all incoming requests are directed to the router as their first step.

There isn’t really a specific place that routing needs to go in an MVC architecture. It could technically fit into the controllers, but many applications separate routing from the controllers, and I have opted to do the same here.

**2. The router routes the request to the UserController**

When the router receives the web request, it looks at several pieces of information to decide how to proceed. In this case, it likely sees something like a PUT request to the /user endpoint and opts to send the request to some code in the UserController source code. This might be a method Update on a UserController type, or perhaps just an UpdateUser function.

```
// Theoretical code. Don't add it to your project
package controllers

// Option A: UserController type with an Update method
type UserController struct {
  // ...
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
  // ...
}

// Option B: UpdateUser function
func UpdateUser(w http.ResponseWriter, r *http.Request) {
  // ...
}
```

!![[Alt text](image.png)](mwc-3to4.png)

**3. The UserController uses the UserStore to update the user’s contact info**

When the incoming web request gets passed along to the user controller code, it will need to do a few things. First, it needs to parse the incoming data. Second, it needs to update the user in the database with the new information provided.

Controllers don’t interact directly with the database, so the user controller uses the UserStore provided by the models package to update the user’s contact information. This allows us to isolate all of our database specific code from our controller code, making both easier to manage.

**4. The UserStore returns the updated data**

After updating the user in the database, the UserStore will return the updated user object to the controller. This sounds complex, but really this is just a function inside the models package returning data after it is called.

![![!\[Alt text\](image.png)](mwc-1to2.png)](mwc-5to6.png)

**5. The UserController uses the ShowUser view to generate HTML**

Once the user has been updated the last thing our controller needs to do is render a response to the user. Controllers don’t create HTML directly, so our code would use something like a ShowUser view to generate an HTML response that shows the newly updated contact information.

Notice that in all of these steps our controller is basically just calling code from the views and models packages. It rarely does any hard work on its own and instead acts more like a coordinator.

**6. The ShowUser view writes and HTML response to the user**

After the user controller uses the ShowUser view, HTML will be returned in a response to the web request. This will likely be something like their contact information page with the updated information and perhaps a message saying something like, “Your information has been updated!”
