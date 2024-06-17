#Simplified workflows:

##Enterprise:

Front-end requirements:

- User should be able to register:
  - This requires a list of countries, counties, cities and phone codes to be available (done)
- User should be able to recover it's password (done)
- User should be able to update it's profile: (done)
  - Fields that can be updated: profile image, first name, last name, phone, country, county, city (done)
  - This requires a list of countries, counties, cities and phone codes to be available (done)
- User should be able to login and receive a JWTtoken (done)
- User should know beforehand if he/she owns any enterprise ????? maybe this is a design problem
- User should be able to send a complaint (done)
- User should be able to see his/her sent,draft and archived complaints (done)
- User should be able to rate a closed complaint (done)

- When user is an enterprise owner:

  - User should be able to create an enterprise: (done)
    - This requires a list of industries to be available (done)
    - This requires a list of countries, counties, cities and phone codes to be available (done)
  - User should be able to update the enterprise profile:
    - Fields that can be updated: logo image, website, email, phone,
      country, county, city (done)
    - This requires a list of industries to be available (done)
    - This requires a list of countries, counties, cities and phone codes to be available (done)
  - User should be able to see the enterprise profile:
    - He/she should be able to hire,approve and fire employees (not in front-end yet)
    - He/she should be able to see the enterprise complaints
    - He/she should be able to see,review and manage the enterprise employees (done)
    - User should be able to create,see and manage enterprise complaints

- When user is an employee:
  - User should be able to see the enterprise employee profile
  - User should be able to see the enterprise complaints
  - User should be able to see the enterprise employees
  - User should be able to see and manage the enterprise complaints
  - If employeee is a manager:
    - User should be able to hire employees
    - User should be able to feedback assistants

All app services will be translated later to command/queries after the persistence layer update


###Hiring:

1. An enterprise selects a user to hire from the list of available to work
   -An invitation is sent to the user
   -The invitation contains the proposed position
   -It verifies if the user already works there
2. The selected user react to the invitation he can accept it or ignore it
   -if he ignores it, nothing happens
   -if he accepts it it its a new employee waiting for
   the final appproval
3. A list of pending hirings will be disponible for the owner of the enterprise
   -the list contains all user who has accepted the invitation
   -its the final step of hiring an employee
   -if its canceled, its moved to log and you cannot undone this operation from the client side

### Complaint:

1. An user or enterprise writes a complaint to an enterprise, its status starts as OPEN

-

2. An employee or the owner answer the complaint for first time, its status changes to STARTED

-

3. The author of the complaint answer for first time to the employee/owner, the status changes to IN_DISCUSSION

-

4. When the employee or owner consider it's solved, it ask the author to review his/her attention
   the status change to IN_REVIEW
5. The author of the complaint write a review on the attention and the complaint get closed.

-

6. A manager receive the closed complaint, review the discussion and provides a feedback to the assistant
   the status changes to IN_HISTORY.

connection ack message example
{
    	"type":        "connection_ack",
		"operation_id": nil,
		"payload": {
		"query":        "subscription {userNotifications(ID: \"mock-client@user.com\") { hiringInvitation {id,enterpriseName, proposedPosition } }}", //how you obtain them
			"subscription_id": "SubscriptionOperationName-UniqueID"
			"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJlcmNobzAwMUBnbWFpbC5jb20iLCJmdWxsX25hbWUiOiJGZXJuYW5kbyBBZ3VzdGluIFNvcmlhIiwicHJvZmlsZV9pbWciOiIvZGVmYXVsdC5qcGciLCJnZW5kZXIiOiJNQUxFIiwicHJvbm91biI6IkhlIiwiY2xpZW50X2RhdGEiOnsiaXAiOiIiLCJkZXZpY2UiOiIiLCJnZW9sb2NhbGl6YXRpb24iOnsibGF0aXR1ZGUiOjAsImxvbmdpdHVkZSI6MH0sImxvZ2luX2RhdGUiOiIxNzE4MjA4MDY3NzU5In0sInJlbWVtYmVyX21lIjp0cnVlLCJhdXRob3JpdGllcyI6W3siZW50ZXJwcmlzZV9pZCI6IlNwb29uIGNvbXBhbnkiLCJhdXRob3JpdHkiOiJPV05FUiJ9XSwiZXhwIjoxNzE4Mjk0NDY3LCJpYXQiOjE3MTgyMDgwNjd9.I-OAPbUf4PFkV2FxMRCHa0cqyHxwSARcUM43TGaCIao", //user authentication
			"enterprise_id":   "" //enterprise authentication -if required-
		}
}