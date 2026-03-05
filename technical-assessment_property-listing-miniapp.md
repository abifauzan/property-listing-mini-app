# Technical Assessment: Property Listing Mini App

> Converted from the provided PDF into Markdown. Includes page images for visual reference where available.

---

## Page 1

![Page 1](./page-01.png)

Technical Assessment: Property Listing Mini App
1. Overview
The goal of this test is to build a simplified "Property Finder" application. You
are required to develop both the Frontend (using the DANA Mini Program
framework) and a Backend REST API (using GoLang or Node.js) to serve
the data.
Please try to replicate the UI/UX shown in the attached screenshots as
closely as possible.

---

## Page 2

![Page 2](./page-02.png)


---

## Page 3

![Page 3](./page-03.png)

2. Technology Stack
- 
Frontend: DANA Mini Program Framework.
  - 
Documentation: https://mini-program.dana.id/docs/
  - 
Note: The "About" screen in the reference mentions "Ant Design
Mobile". You may use available component libraries compatible with the
mini-program environment or build custom components.
- 
Backend: GoLang (Preferred) or Node.js.

---

## Page 4

![Page 4](./page-04.png)

- 
Database: You may use a persistent database (PostgreSQL,
MongoDB, MySQL) or a simple in-memory data store/JSON file for this
assessment.
3. Functional Requirements
A. Navigation (Tab Bar)
Implement a bottom Tab Bar with two items:
- Navigates to the "About" page.
- Navigates to the listing page.
B. Screen 1: About Page (Home)
- UI: Reference Screenshot 1.
- Content: Display a logo, a title ("About Property Finder"), and a descriptive
text.
- Functionality: This is a static information page.
C. Screen 2: Property Listing
- UI: Reference Screenshots 2, 3, and 4.
- Features:
- Search Bar: Allow users to filter properties by title (e.g., typing "sam"
shows items containing "sample").
- Layout Toggle: Implement a button to switch the list view between
Vertical List (Screenshot 2) and Grid View (Screenshot 3).
- Data Fetching: Fetch property data from your Backend API.
- Display: Show Property Image, Title, and Price (formatted as Rp
X.XXX.XXX).
D. Screen 3: Property Detail
- UI: Reference Screenshot 5.
- Navigation: Clicking a card in the Listing screen should navigate here,
passing the Property ID.
- Content: Display the full details: Large Hero Image, Title, formatted Price,
gallery carousel and a Description text.

---

## Page 5

![Page 5](./page-05.png)

Details
- Action: A fixed "Book Now" button at the bottom (can be a dummy action
that shows a toast/alert).
4. Backend & Data Structure
Please build a RESTful or graphql API to support the frontend.Data Schema
(JSON Example): Please use the following structure as a baseline for your
data:
```json

// Listings
{
"data": {
"propertyListings": [
{
"documentId": "x0zzhpyrox8ixdy75e2zpw1m",
"Banner": {
"url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/
pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg"
},
"Title": "My Villa Sample",
"Price": "750000",
"createdAt": "2025-01-31T08:18:56.778Z"
},
{
"documentId": "fn9yzf1ii9zrdscnn9nw9wk4",
"Banner": {
"url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/
pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg"
},
"Title": "My Villa",
"Price": "750000",
"createdAt": "2024-12-12T12:39:51.290Z"
},
{
"documentId": "l74mb5x4h2hxdhmrx0hm6nyk",
"Banner": {
"url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987633/
pexels_vlad_52729368_8059591_280a9e0982.jpg"
},
"Title": "Villa Sample",
"Price": "500000",
"createdAt": "2024-12-12T12:30:50.940Z"
}
]
}
}
```

---

## Page 6

![Page 6](./page-06.png)

{
"data": {
"propertyListings": [
{
"Banner": {
"url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/
pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg"
},
"Title": "My Villa Sample",
"documentId": "x0zzhpyrox8ixdy75e2zpw1m",
"Description": "Aenean tellus velit, porttitor eget diam vel, pharetra pellentesque ante.
Proin egestas ornare ipsum. Morbi aliquam dui et felis placerat, in posuere libero molestie.
Vivamus ut arcu nulla. Quisque vulputate neque risus, vitae porttitor neque consectetur vel. Duis
in ex sem. Sed rutrum porta scelerisque. Nullam eget turpis sit amet tellus vehicula gravida a in
justo. Quisque eget nisl ut ipsum lacinia pulvinar. Ut egestas, tortor vitae volutpat suscipit,
magna augue semper nunc, vitae sodales est ipsum at turpis. Aenean quis erat et libero
rhoncus sodales. Mauris nunc orci, eleifend ac accumsan id, porta ut magna. Proin at urna
finibus, dapibus lacus vel, convallis tortor.",
"Images": [
{
"url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/
pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg"
},
{
"url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987637/
pexels_koprivakart_3354648_fe963523cc.jpg"
},
{
"url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987636/
pexels_avinashpatel_544542_615f33a344.jpg"
},
{
"url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987633/
pexels_vlad_52729368_8059591_280a9e0982.jpg"
}
],
"Facilities": [
"Kitchen",
"Free Park",
"Bar"
],
"Price": "750000",
"Terms": "Donec gravida, quam at volutpat pharetra, elit ante ullamcorper eros, nec
dictum leo ex a purus. Donec euismod sem non mauris consectetur, vitae consectetur magna
sodales. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis
egestas. Nam nec aliquam libero. Praesent et consequat arcu, sit amet ullamcorper nibh. Sed
laoreet sodales sollicitudin. Nulla massa ipsum, convallis sed lacus sed, maximus ultricies
dolor.",
"Conditions": "Vestibulum sit amet vehicula tellus. Etiam sed pharetra risus. Vivamus
facilisis laoreet ante eget ultrices. Morbi at varius enim, posuere euismod diam. Praesent
congue pretium diam eget sagittis. Sed consequat risus vitae vestibulum sollicitudin. Maecenas
blandit, orci eget placerat mattis, nulla massa iaculis elit, quis malesuada arcu lacus vel nulla.
Cras et pharetra sem. Curabitur efficitur porta lectus in pellentesque."
}
]
}
}

5. Evaluation Criteria

---

## Page 7

![Page 7](./page-07.png)

We will evaluate your submission based on the following:
- Visual Fidelity: How closely the Mini Program matches the provided
screenshots (layout, spacing, typography).
- Framework Knowledge: Correct usage of the DANA Mini Program
lifecycle, AXML, and ACSS.
- Code Quality: Clean, readable, and maintainable code structure (both
Frontend and Backend).
- Functionality: Search works, layout toggle works, and navigation is
smooth.
- Bonus Points:
- Implementation of "Pull to Refresh" on the listing page.
- Unit tests for the Backend logic.
- Handling empty states (e.g., no search results found).
6. Submission
- Upload your source code to a public Git repository (GitHub/GitLab).
- Include a README.md file with:
- Instructions on how to run the Backend.
- Instructions on how to run the Mini Program (IDE requirements, etc.).
- Screenshots or a screen recording of your running application.
- Send the repository link to us maximum five(5) days after you receive
the email

Good Luck!
