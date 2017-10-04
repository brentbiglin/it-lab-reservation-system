# Technology Learning Studio

* October
	* [October 3rd](#2017-10-03) - Initial Post
	* [October 4th](#2017-10-04) - First Day of Ponzu
 
## Creating a Reservation System

### 2017-10-03

#### Initial Post

This readme acts as a blog for Professor Howison's Technology Learning Studio course at the University of Texas at Austin. I'll document my learning process here.

### 2017-10-04 

#### First Day of Ponzu

Ponzu is a framework for creating a content management system (CMS) with a built-in database and encryption. My goal is to create a reservation system for the devices in the IT Lab. This won't be a user-facing system, but an administrative system for tracking the devices, who has them checked out, when they are due, their associated accessories, and what kind of condition they're in.

The first day of Ponzu, I dealt a lot with the setup of the environment and the dependencies of the CMS. I've never dealt with Docker before, and I alternated between populating the CMS with the appropriate attributes and the corresponding inputs (short text, text block, select) and researching the best way to deploy it while I'm developing it so that IT staff members can populate it and provide feedback.

After that, my goal is to take advantage of the JSON APIs and addons that have been developed for Ponzu so that when a device is checked out by a student, either the student or the IT Lab or both will be warned when the item is due, overdue, or reserved either by email our through Spiceworks tickets. Figuring out how to automate these parts together will be the most difficult and time-consuming part of the process. Deploying it on a server should just be a matter of coordinating with the IT staff and Googling.
