# Task

> Enclose a README with a short description of your product, the thoughts behind the design of your code, a list of the stories you selected and why you chose the features that you did.
> It should also contain instructions on how to install, build and run the application.

# Short description 

TODO

# Thoughts behind the code

TODO

# List of stories selected, and why they were chosen

> We want at least 2 points to be implemented from each of the requirements areas in the matrix further below in this document, and the rest is up to you!

Looking at the feature list, I decided to look at the features and how many points each of it will bring, as well as which features I'm most comfortable with or/and have most experience with, as well as something that can showcase my skills appropriately

Since it's at least 2 points from each column, I will describe each of the column and decisions below

`_design/accessibility`
- A modern design (1 pt)
- No UI framework used (such as Bootstrap, Ant) (1 pt)

I decide to start with these first, since I will take the example design provided in the code, and try to make it nice, simple and flat design, which shouldn't need too many additions from the provided example
The second one about the UI framework should be solved with just using plain React

An extra task I decided to go for if I'll have time is
- Responsive design, works on mobile and tablets (2 pt)

I had experience working with SASS, and creating an application from a given design, that provided all the different dimensions and sizes for different screens. I'm in no way experienced or know how to properly do it, but taking that as a reference point I decided that I can still make it more mobile friendly, since the search bar would pretty much take almost all the width of the screen, so that needs minimum adjustments if any, and the cards with workers could be changed from a row of 4 to a row of 1/2 cards for example to make it nicer to read on a mobile, but this is only if I'll have time remaning and add extra 2 points to keep it safe

`_functionality`
- Don’t use our API, build your own API by scraping the current page and setting up a local backend server. How/If you persist is up to you. (4 pt)

This one is pretty straightforward, since we talked with Andreas and agreed that it makes sense for me to showcase my backend code with Golang, and this one will be the best way to do so

`_testing/QA`
- End-to-end testing (with an existing framework) (2 pt)

I decided to take this one. It won't be a very extensive testing, since it will be a simple GET endpoint that will not take any parameters, but I decided it's still a good idea to do it, since I can showcase how I worked with testing in Golang, as well as add a mock for fetcher of the API, which will allow us to have consistent test results, and not send an API request every time the tests are running

In sum it will be 8 points (or +2 points if I'll have time to do responsive design)

I decided not to choose the other features since some of them felt a bit confusing and needed more discussion. For example Fancy animations (1 pt) or Keyboard only function (1 pt). I emailed about the first one already and the answer I got was "use your skills to deliver a 'fancy animation' or a 'modern design' such as you would, given a client that doesn't quite know how to define it themselves but rather trust in your judgement.". Regarding the keyboard input it's also not very clear what that means. The only keyboard input I could think of was the search functionality, so I could use "Enter" to submit my filter for names and cities, but since I'm not going to do filtering in the features I chose, and if I'd do search I would filter the cards differently than pressing enter (so filter them every time a new letter comes in the search field). Some other features I didn't have much experience in doing, and decided not to pursue them

# Questions

1.
- Don’t use our API, build your own API by scraping the current page and setting up a local backend server. How/If you persist is up to you. (4 pt)

When working with the response, I saw that it has fields like "published", and "highlighted". For now from experience, I'm discarding employees that have "published" set to `false`. Not sure about the "highlighted" field. I will send an email to clarify this later, this shouldn't be hard to filter out if we need to (I added a function to filter the raw response, so it can be extended)

2.
- End-to-end testing (with an existing framework) (2 pt)

When working on the end-to end test (provided in backend/api/employees_test.go), I also decided to add a small unit tests as well for the filtering function (in backend/pkg/employee/employee_test.go), since that is pretty much the only function that exists in the API end point, that will filter out employees by criteria. It's not a big function, but to showcase Golang table testing and unit testing in general, I decided to throw it in as well, don't know if it's enough to cover this feature request
- Unit tests for existing functionality (reasonable coverage) (2 pt)