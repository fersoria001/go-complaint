# Minimum domain refactor requirements for the new UI:
## Updates:
This domain refactors will require refactor the persistence and application layers too.
### User aggregate: (done)
    . It must have an Id field of type UUID PRIMARY KEY, currently it's and string email address.
### Enterprise aggregate: (done)
    . It must have an Id field of type UUID PRIMARY KEY, currently it's the enterprise name.
### Complaint aggregate(done):
    . It must hold the sender as a value object {id, type string | isEnterprise bool }.
    . The complaint body must be move to act as the first reply of the complaint.
## New:
### Hiring proccess must be an aggregate | entity(done).
    . This probably will need to be constructed by an event pipeline, I will avoid doing this
    if it takes a lot of time.
### Reviews must be an aggregate | entity(done)
    <!-- . This probably will need to be constructed by an event pipeline, I will avoid doing this
    if it takes a lot of time. -->
    . A best solution would be to separate the queries and manage them differently in the GUI
    . Manage them with the complaint status and query for complaints where id and status.
    ----- it was solved by querying for complaints because they got all the required data to fill the GUI ------
### Complaint handler(done): 
    <!-- . Both the list and the complaint should be fetched with WS to notify about new messages. -->
    .The complaint list has been moved to subscriptions, still need the chat handler
### Subscriptions(done):
    . added the complaint list as a subscription
    . Right now the only subscription needed is the notifications subscription, it will need to
    before or after persist the object in the database send the objects to some kind of intermediary
    that can be DI into the graphQL resolver preserving the ID or some method to select it.